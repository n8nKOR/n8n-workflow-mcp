package test

import (
	"testing"

	"github.com/n8nKOR/n8n-workflow-mcp/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkflowTypes(t *testing.T) {
	t.Run("NewWorkflow", func(t *testing.T) {
		description := "Test workflow description"
		workflow := types.NewWorkflow(description)

		require.NotNil(t, workflow)
		assert.Equal(t, description, workflow.Description)
		assert.Equal(t, types.StatusDraft, workflow.Status)
		assert.NotEmpty(t, workflow.ID)
		assert.False(t, workflow.CreatedAt.IsZero())
		assert.False(t, workflow.UpdatedAt.IsZero())
	})

	t.Run("UpdateStatus", func(t *testing.T) {
		workflow := types.NewWorkflow("Test")
		originalStatus := workflow.Status

		workflow.UpdateStatus(types.StatusApproved)

		assert.Equal(t, types.StatusApproved, workflow.Status)
		assert.NotEqual(t, originalStatus, workflow.Status)
		// UpdatedAt이 업데이트되었는지 확인 (0이 아닌 시간)
		assert.False(t, workflow.UpdatedAt.IsZero())
	})

	t.Run("AddError", func(t *testing.T) {
		workflow := types.NewWorkflow("Test")

		err := types.ValidationError{
			Line:      10,
			Column:    5,
			ErrorType: "MISSING_FIELD",
			Message:   "Required field missing",
		}

		workflow.AddError(err)

		assert.Equal(t, types.StatusError, workflow.Status)
		assert.Len(t, workflow.Errors, 1)
		assert.Equal(t, err.Line, workflow.Errors[0].Line)
	})

	t.Run("ClearErrors", func(t *testing.T) {
		workflow := types.NewWorkflow("Test")
		workflow.AddError(types.ValidationError{Message: "Test error"})

		workflow.ClearErrors()

		assert.Empty(t, workflow.Errors)
		assert.Equal(t, types.StatusDraft, workflow.Status)
	})

	t.Run("ToJSON", func(t *testing.T) {
		workflow := types.NewWorkflow("Test workflow")

		jsonStr, err := workflow.ToJSON()

		require.NoError(t, err)
		assert.Contains(t, jsonStr, "Test workflow")
		assert.Contains(t, jsonStr, "draft")
	})
}
