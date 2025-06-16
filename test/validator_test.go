package test

import (
	"strings"
	"testing"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestN8nValidator(t *testing.T) {
	v := validator.NewN8nValidator(true, 50)

	t.Run("ValidWorkflow", func(t *testing.T) {
		validJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "node1",
					"name": "Start",
					"type": "n8n-nodes-base.start",
					"position": [100, 200],
					"parameters": {}
				},
				{
					"id": "node2", 
					"name": "HTTP Request",
					"type": "n8n-nodes-base.httpRequest",
					"position": [300, 200],
					"parameters": {
						"url": "https://api.example.com"
					}
				}
			],
			"connections": {
				"Start": {
					"main": [
						[
							{
								"node": "HTTP Request",
								"type": "main",
								"index": 0
							}
						]
					]
				}
			},
			"active": true
		}`

		result, err := v.ValidateWorkflow(validJSON)
		require.NoError(t, err)
		assert.True(t, result.IsValid)
		assert.False(t, result.IsError)
		assert.Equal(t, 0, result.ErrorCount)
		assert.Empty(t, result.Errors)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		invalidJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "node1",
					"name": "Start"
					"type": "n8n-nodes-base.start"
				}
			]
		}`

		result, err := v.ValidateWorkflow(invalidJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)
		assert.Equal(t, 1, result.ErrorCount)
		assert.Len(t, result.Errors, 1)
		assert.Equal(t, "INVALID_JSON_SYNTAX", result.Errors[0].ErrorType)
		assert.Contains(t, result.Errors[0].Message, "JSON syntax error")
	})

	t.Run("MissingNodes", func(t *testing.T) {
		emptyNodesJSON := `{
			"name": "Test Workflow",
			"nodes": [],
			"connections": {}
		}`

		result, err := v.ValidateWorkflow(emptyNodesJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)

		// Should have error about missing nodes
		found := false
		for _, err := range result.Errors {
			if err.ErrorType == "MISSING_REQUIRED_FIELD" && err.Field == "nodes" {
				found = true
				assert.Contains(t, err.Message, "at least one node")
				break
			}
		}
		assert.True(t, found, "Expected error about missing nodes")
	})

	t.Run("MissingConnections", func(t *testing.T) {
		noConnectionsJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "node1",
					"name": "Start",
					"type": "n8n-nodes-base.start",
					"position": [100, 200]
				}
			]
		}`

		result, err := v.ValidateWorkflow(noConnectionsJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)

		// Should have error about missing connections
		found := false
		for _, err := range result.Errors {
			if err.ErrorType == "MISSING_REQUIRED_FIELD" && err.Field == "connections" {
				found = true
				assert.Contains(t, err.Message, "connections object")
				break
			}
		}
		assert.True(t, found, "Expected error about missing connections")
	})

	t.Run("InvalidNodeStructure", func(t *testing.T) {
		invalidNodeJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "",
					"name": "",
					"type": "",
					"position": [100]
				}
			],
			"connections": {}
		}`

		result, err := v.ValidateWorkflow(invalidNodeJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)
		assert.Greater(t, result.ErrorCount, 0)

		// Check for various node validation errors
		errorTypes := make(map[string]bool)
		for _, err := range result.Errors {
			errorTypes[err.ErrorType] = true
		}

		assert.True(t, errorTypes["MISSING_REQUIRED_FIELD"], "Expected missing required field errors")
		assert.True(t, errorTypes["INVALID_PARAMETER"], "Expected invalid parameter errors")
	})

	t.Run("DuplicateNodeIDs", func(t *testing.T) {
		duplicateIDJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "duplicate_id",
					"name": "Node 1",
					"type": "n8n-nodes-base.start",
					"position": [100, 200]
				},
				{
					"id": "duplicate_id", 
					"name": "Node 2",
					"type": "n8n-nodes-base.httpRequest",
					"position": [300, 200]
				}
			],
			"connections": {}
		}`

		result, err := v.ValidateWorkflow(duplicateIDJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)

		// Should have duplicate ID error
		found := false
		for _, err := range result.Errors {
			if err.ErrorType == "DUPLICATE_NODE_ID" {
				found = true
				assert.Contains(t, err.Message, "Duplicate node ID")
				assert.Equal(t, "duplicate_id", err.NodeID)
				break
			}
		}
		assert.True(t, found, "Expected duplicate node ID error")
	})

	t.Run("ErrorSummary", func(t *testing.T) {
		complexErrorJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "",
					"name": "",
					"type": ""
				},
				{
					"id": "node2",
					"name": "",
					"type": ""
				}
			],
			"connections": {}
		}`

		result, err := v.ValidateWorkflow(complexErrorJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)

		// Check summary
		assert.Greater(t, result.Summary.TotalLines, 0)
		assert.Greater(t, len(result.Summary.ErrorLines), 0)
		assert.Greater(t, len(result.Summary.ErrorTypes), 0)

		// Should have multiple missing required field errors
		assert.Greater(t, result.Summary.ErrorTypes["MISSING_REQUIRED_FIELD"], 0)
	})

	t.Run("MaxErrorsLimit", func(t *testing.T) {
		// Create validator with low max errors
		limitedValidator := validator.NewN8nValidator(true, 3)

		// JSON that would generate many errors
		manyErrorsJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{"id": "", "name": "", "type": ""},
				{"id": "", "name": "", "type": ""},
				{"id": "", "name": "", "type": ""},
				{"id": "", "name": "", "type": ""},
				{"id": "", "name": "", "type": ""}
			],
			"connections": {}
		}`

		result, err := limitedValidator.ValidateWorkflow(manyErrorsJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)

		// Should be limited to 3 errors
		assert.LessOrEqual(t, len(result.Errors), 3)
	})

	t.Run("DisabledLineNumbers", func(t *testing.T) {
		// Create validator with line numbers disabled
		noLineValidator := validator.NewN8nValidator(false, 50)

		invalidJSON := `{
			"name": "Test Workflow",
			"nodes": [],
			"connections": {}
		}`

		result, err := noLineValidator.ValidateWorkflow(invalidJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)

		// All errors should have line 1 when line numbers are disabled
		for _, err := range result.Errors {
			assert.Equal(t, 1, err.Line)
		}
	})

	t.Run("ValidTags", func(t *testing.T) {
		validTagsJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "node1",
					"name": "Start",
					"type": "n8n-nodes-base.start",
					"position": [100, 200],
					"parameters": {}
				}
			],
			"connections": {
				"Start": {
					"main": [[]]
				}
			},
			"tags": [
				{
					"id": "ZZS1BBlFeeZV8Kg0",
					"name": "AI",
					"createdAt": "2025-05-11T14:48:25.304Z",
					"updatedAt": "2025-05-11T14:48:25.304Z"
				}
			]
		}`

		result, err := v.ValidateWorkflow(validTagsJSON)
		require.NoError(t, err)
		assert.True(t, result.IsValid)
		assert.False(t, result.IsError)
		assert.Equal(t, 0, result.ErrorCount)
	})

	t.Run("InvalidTags", func(t *testing.T) {
		invalidTagsJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "node1",
					"name": "Start",
					"type": "n8n-nodes-base.start",
					"position": [100, 200],
					"parameters": {}
				}
			],
			"connections": {
				"Start": {
					"main": [[]]
				}
			},
			"tags": [
				{
					"id": "",
					"name": "",
					"createdAt": "",
					"updatedAt": ""
				}
			]
		}`

		result, err := v.ValidateWorkflow(invalidTagsJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)
		assert.Greater(t, result.ErrorCount, 0)

		// Check for tag validation errors
		tagErrors := 0
		for _, err := range result.Errors {
			if strings.HasPrefix(err.Field, "tags.") {
				tagErrors++
			}
		}
		assert.Greater(t, tagErrors, 0, "Expected tag validation errors")
	})

	t.Run("DuplicateTagIDs", func(t *testing.T) {
		duplicateTagJSON := `{
			"name": "Test Workflow",
			"nodes": [
				{
					"id": "node1",
					"name": "Start",
					"type": "n8n-nodes-base.start",
					"position": [100, 200],
					"parameters": {}
				}
			],
			"connections": {
				"Start": {
					"main": [[]]
				}
			},
			"tags": [
				{
					"id": "duplicate_tag",
					"name": "Tag 1",
					"createdAt": "2025-05-11T14:48:25.304Z",
					"updatedAt": "2025-05-11T14:48:25.304Z"
				},
				{
					"id": "duplicate_tag",
					"name": "Tag 2",
					"createdAt": "2025-05-11T14:48:25.304Z",
					"updatedAt": "2025-05-11T14:48:25.304Z"
				}
			]
		}`

		result, err := v.ValidateWorkflow(duplicateTagJSON)
		require.NoError(t, err)
		assert.False(t, result.IsValid)
		assert.True(t, result.IsError)

		// Should have duplicate tag ID error
		found := false
		for _, err := range result.Errors {
			if err.ErrorType == "DUPLICATE_TAG_ID" {
				found = true
				assert.Contains(t, err.Message, "Duplicate tag ID")
				break
			}
		}
		assert.True(t, found, "Expected duplicate tag ID error")
	})
}
