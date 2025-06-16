package validator

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/n8nKOR/n8n-workflow-mcp/pkg/types"
)

// N8nValidator는 n8n 워크플로우 검증을 구현합니다
type N8nValidator struct {
	enableLineNumbers bool
	maxErrors         int
}

// NewN8nValidator는 새로운 n8n 워크플로우 검증기를 생성합니다
func NewN8nValidator(enableLineNumbers bool, maxErrors int) *N8nValidator {
	if maxErrors <= 0 {
		maxErrors = 50 // 기본 최대 오류 수
	}

	return &N8nValidator{
		enableLineNumbers: enableLineNumbers,
		maxErrors:         maxErrors,
	}
}

// ValidateWorkflow는 n8n 워크플로우 JSON을 검증하고 상세 오류를 반환합니다
func (v *N8nValidator) ValidateWorkflow(jsonStr string) (*types.ValidationResult, error) {
	result := &types.ValidationResult{
		IsValid: true,
		IsError: false,
		Errors:  []types.ValidationError{},
		Summary: types.ValidationSummary{
			ErrorTypes: make(map[string]int),
		},
	}

	// 라인 추적을 포함한 JSON 파싱
	var workflow types.N8nWorkflow
	lines := strings.Split(jsonStr, "\n")
	result.Summary.TotalLines = len(lines)

	// 기본 JSON 문법 검증
	if err := json.Unmarshal([]byte(jsonStr), &workflow); err != nil {
		validationErr := v.parseJSONError(err, lines)
		result.Errors = append(result.Errors, validationErr)
		result.IsValid = false
		result.IsError = true
		result.ErrorCount = 1
		result.Summary.ErrorTypes["INVALID_JSON_SYNTAX"] = 1
		result.Summary.ErrorLines = []int{validationErr.Line}
		return result, nil
	}

	// n8n 워크플로우 구조 검증
	errors := v.validateN8nStructure(&workflow, lines)

	if len(errors) > 0 {
		result.IsValid = false
		result.IsError = true
		result.Errors = errors
		result.ErrorCount = len(errors)

		// 오류 유형과 라인 수 계산
		errorLines := make(map[int]bool)
		errorTypes := make(map[string]int)

		for _, err := range errors {
			errorLines[err.Line] = true
			errorTypes[err.ErrorType]++
		}

		// 슬라이스로 변환
		for line := range errorLines {
			result.Summary.ErrorLines = append(result.Summary.ErrorLines, line)
		}
		result.Summary.ErrorTypes = errorTypes
	}

	return result, nil
}

// parseJSONError는 JSON 파싱 오류에서 라인 정보를 추출합니다
func (v *N8nValidator) parseJSONError(err error, lines []string) types.ValidationError {
	errStr := err.Error()

	// 오류 메시지에서 라인 번호 추출 시도
	lineRegex := regexp.MustCompile(`line (\d+)`)
	matches := lineRegex.FindStringSubmatch(errStr)

	line := 1
	if len(matches) > 1 {
		if parsedLine, parseErr := json.Number(matches[1]).Int64(); parseErr == nil {
			line = int(parsedLine)
		}
	}

	// 열 정보 추출 시도
	columnRegex := regexp.MustCompile(`column (\d+)`)
	columnMatches := columnRegex.FindStringSubmatch(errStr)

	column := 1
	if len(columnMatches) > 1 {
		if parsedColumn, parseErr := json.Number(columnMatches[1]).Int64(); parseErr == nil {
			column = int(parsedColumn)
		}
	}

	context := ""
	if line <= len(lines) && line > 0 {
		context = strings.TrimSpace(lines[line-1])
	}

	return types.ValidationError{
		Line:       line,
		Column:     column,
		ErrorType:  "INVALID_JSON_SYNTAX",
		Message:    fmt.Sprintf("JSON syntax error: %s", err.Error()),
		Suggestion: "Please check JSON syntax and fix any malformed JSON",
		Context:    context,
	}
}

// validateN8nStructure는 n8n 워크플로우 구조를 검증합니다
func (v *N8nValidator) validateN8nStructure(workflow *types.N8nWorkflow, lines []string) []types.ValidationError {
	var errors []types.ValidationError

	// 필수 필드 확인
	if len(workflow.Nodes) == 0 {
		errors = append(errors, types.ValidationError{
			Line:       v.findFieldLine(lines, "nodes"),
			Column:     1,
			Field:      "nodes",
			ErrorType:  "MISSING_REQUIRED_FIELD",
			Message:    "Workflow must contain at least one node",
			Suggestion: "Add at least one node to the workflow",
			Context:    "\"nodes\": []",
		})
	}

	if workflow.Connections == nil {
		errors = append(errors, types.ValidationError{
			Line:       v.findFieldLine(lines, "connections"),
			Column:     1,
			Field:      "connections",
			ErrorType:  "MISSING_REQUIRED_FIELD",
			Message:    "Workflow must have a connections object",
			Suggestion: "Add a connections object to define node relationships",
			Context:    "\"connections\": {}",
		})
	}

	// 각 노드 검증
	nodeIDs := make(map[string]bool)
	nodeNames := make(map[string]bool)
	for i, node := range workflow.Nodes {
		nodeErrors := v.validateNode(&node, i, lines, nodeIDs)
		errors = append(errors, nodeErrors...)

		// ID와 이름 모두 추적
		if node.ID != "" {
			nodeIDs[node.ID] = true
		}
		if node.Name != "" {
			nodeNames[node.Name] = true
		}

		if len(errors) >= v.maxErrors {
			break
		}
	}

	// 연결 검증
	if workflow.Connections != nil {
		connectionErrors := v.validateConnections(workflow.Connections, nodeNames, lines)
		errors = append(errors, connectionErrors...)
	}

	// 태그 검증
	if len(workflow.Tags) > 0 {
		tagErrors := v.validateTags(workflow.Tags, lines)
		errors = append(errors, tagErrors...)
	}

	// 오류를 maxErrors로 제한
	if len(errors) > v.maxErrors {
		errors = errors[:v.maxErrors]
	}

	return errors
}

// validateNode는 단일 n8n 노드를 검증합니다
func (v *N8nValidator) validateNode(node *types.N8nNode, index int, lines []string, nodeIDs map[string]bool) []types.ValidationError {
	var errors []types.ValidationError

	// 필수 노드 필드 확인
	if node.ID == "" {
		errors = append(errors, types.ValidationError{
			Line:       v.findNodeLine(lines, index, "id"),
			Column:     1,
			NodeID:     "",
			NodeName:   node.Name,
			Field:      "id",
			ErrorType:  "MISSING_REQUIRED_FIELD",
			Message:    "Node must have an ID",
			Suggestion: "Add a unique ID to the node",
			Context:    "\"id\": \"\"",
		})
	} else {
		// 중복 ID 확인
		if nodeIDs[node.ID] {
			errors = append(errors, types.ValidationError{
				Line:       v.findNodeLine(lines, index, "id"),
				Column:     1,
				NodeID:     node.ID,
				NodeName:   node.Name,
				Field:      "id",
				ErrorType:  "DUPLICATE_NODE_ID",
				Message:    fmt.Sprintf("Duplicate node ID: %s", node.ID),
				Suggestion: "Use a unique ID for each node",
				Context:    fmt.Sprintf("\"id\": \"%s\"", node.ID),
			})
		}
		nodeIDs[node.ID] = true
	}

	if node.Name == "" {
		errors = append(errors, types.ValidationError{
			Line:       v.findNodeLine(lines, index, "name"),
			Column:     1,
			NodeID:     node.ID,
			NodeName:   "",
			Field:      "name",
			ErrorType:  "MISSING_REQUIRED_FIELD",
			Message:    "Node must have a name",
			Suggestion: "Add a name to the node",
			Context:    "\"name\": \"\"",
		})
	}

	if node.Type == "" {
		errors = append(errors, types.ValidationError{
			Line:       v.findNodeLine(lines, index, "type"),
			Column:     1,
			NodeID:     node.ID,
			NodeName:   node.Name,
			Field:      "type",
			ErrorType:  "MISSING_REQUIRED_FIELD",
			Message:    "Node must have a type",
			Suggestion: "Specify the node type (e.g., 'n8n-nodes-base.httpRequest')",
			Context:    "\"type\": \"\"",
		})
	}

	// 위치 배열 검증
	if len(node.Position) != 2 {
		errors = append(errors, types.ValidationError{
			Line:       v.findNodeLine(lines, index, "position"),
			Column:     1,
			NodeID:     node.ID,
			NodeName:   node.Name,
			Field:      "position",
			ErrorType:  "INVALID_PARAMETER",
			Message:    "Node position must be an array of 2 numbers [x, y]",
			Suggestion: "Set position as [x, y] coordinates",
			Context:    "\"position\": [x, y]",
		})
	}

	return errors
}

// validateConnections는 워크플로우 연결을 검증합니다
func (v *N8nValidator) validateConnections(connections map[string]interface{}, nodeNames map[string]bool, lines []string) []types.ValidationError {
	var errors []types.ValidationError

	for sourceNodeName, connectionData := range connections {
		// 소스 노드 존재 여부 확인
		if !nodeNames[sourceNodeName] {
			errors = append(errors, types.ValidationError{
				Line:       v.findConnectionLine(lines, sourceNodeName),
				Column:     1,
				NodeID:     sourceNodeName,
				Field:      "connections",
				ErrorType:  "INVALID_NODE_CONNECTION",
				Message:    fmt.Sprintf("Connection references non-existent source node: %s", sourceNodeName),
				Suggestion: "Ensure the source node exists in the workflow",
				Context:    fmt.Sprintf("\"%s\": [...]", sourceNodeName),
			})
		}

		// 기본 검증 - 연결은 nil이 아니어야 함
		if connectionData == nil {
			errors = append(errors, types.ValidationError{
				Line:       v.findConnectionLine(lines, sourceNodeName),
				Column:     1,
				NodeID:     sourceNodeName,
				Field:      "connections",
				ErrorType:  "INVALID_NODE_CONNECTION",
				Message:    "Connection cannot be null",
				Suggestion: "Provide valid connection configuration",
				Context:    "null connection found",
			})
		}
	}

	return errors
}

// validateTags는 워크플로우 태그를 검증합니다
func (v *N8nValidator) validateTags(tags []types.N8nTag, lines []string) []types.ValidationError {
	var errors []types.ValidationError

	tagIDs := make(map[string]bool)
	for i, tag := range tags {
		// 필수 필드 검증
		if tag.ID == "" {
			errors = append(errors, types.ValidationError{
				Line:       v.findTagLine(lines, i, "id"),
				Column:     1,
				Field:      "tags.id",
				ErrorType:  "MISSING_REQUIRED_FIELD",
				Message:    "Tag must have an ID",
				Suggestion: "Add a unique ID to the tag",
				Context:    "\"id\": \"\"",
			})
		} else {
			// 중복 ID 확인
			if tagIDs[tag.ID] {
				errors = append(errors, types.ValidationError{
					Line:       v.findTagLine(lines, i, "id"),
					Column:     1,
					Field:      "tags.id",
					ErrorType:  "DUPLICATE_TAG_ID",
					Message:    fmt.Sprintf("Duplicate tag ID: %s", tag.ID),
					Suggestion: "Use a unique ID for each tag",
					Context:    fmt.Sprintf("\"id\": \"%s\"", tag.ID),
				})
			}
			tagIDs[tag.ID] = true
		}

		if tag.Name == "" {
			errors = append(errors, types.ValidationError{
				Line:       v.findTagLine(lines, i, "name"),
				Column:     1,
				Field:      "tags.name",
				ErrorType:  "MISSING_REQUIRED_FIELD",
				Message:    "Tag must have a name",
				Suggestion: "Add a name to the tag",
				Context:    "\"name\": \"\"",
			})
		}

		if tag.CreatedAt == "" {
			errors = append(errors, types.ValidationError{
				Line:       v.findTagLine(lines, i, "createdAt"),
				Column:     1,
				Field:      "tags.createdAt",
				ErrorType:  "MISSING_REQUIRED_FIELD",
				Message:    "Tag must have a createdAt timestamp",
				Suggestion: "Add a createdAt timestamp in ISO format",
				Context:    "\"createdAt\": \"2025-05-11T14:48:25.304Z\"",
			})
		}

		if tag.UpdatedAt == "" {
			errors = append(errors, types.ValidationError{
				Line:       v.findTagLine(lines, i, "updatedAt"),
				Column:     1,
				Field:      "tags.updatedAt",
				ErrorType:  "MISSING_REQUIRED_FIELD",
				Message:    "Tag must have an updatedAt timestamp",
				Suggestion: "Add an updatedAt timestamp in ISO format",
				Context:    "\"updatedAt\": \"2025-05-11T14:48:25.304Z\"",
			})
		}
	}

	return errors
}

// JSON에서 라인 번호를 찾는 헬퍼 함수들
func (v *N8nValidator) findFieldLine(lines []string, fieldName string) int {
	if !v.enableLineNumbers {
		return 1
	}

	pattern := fmt.Sprintf(`"%s"`, fieldName)
	for i, line := range lines {
		if strings.Contains(line, pattern) {
			return i + 1
		}
	}
	return 1
}

func (v *N8nValidator) findNodeLine(lines []string, nodeIndex int, fieldName string) int {
	if !v.enableLineNumbers {
		return 1
	}

	// 노드 라인을 찾는 간단한 휴리스틱 - 실제 구현에서는
	// 위치 추적을 포함한 더 정교한 JSON 파싱이 필요함
	nodeCount := 0
	for i, line := range lines {
		if strings.Contains(line, `"nodes"`) {
			// 노드 배열을 찾았으므로 계산 시작
			for j := i; j < len(lines); j++ {
				if strings.Contains(lines[j], `{`) {
					if nodeCount == nodeIndex {
						// 우리 노드를 찾았으므로 필드 찾기
						for k := j; k < len(lines) && k < j+20; k++ {
							if strings.Contains(lines[k], fmt.Sprintf(`"%s"`, fieldName)) {
								return k + 1
							}
						}
						return j + 1
					}
					nodeCount++
				}
			}
		}
	}
	return 1
}

func (v *N8nValidator) findConnectionLine(lines []string, nodeName string) int {
	if !v.enableLineNumbers {
		return 1
	}

	pattern := fmt.Sprintf(`"%s"`, nodeName)
	inConnections := false
	for i, line := range lines {
		if strings.Contains(line, `"connections"`) {
			inConnections = true
			continue
		}
		if inConnections && strings.Contains(line, pattern) {
			return i + 1
		}
		// 연결 섹션을 지나쳤으면 중단
		if inConnections && strings.Contains(line, `}`) && !strings.Contains(line, pattern) {
			break
		}
	}
	return 1
}

func (v *N8nValidator) findTagLine(lines []string, tagIndex int, fieldName string) int {
	if !v.enableLineNumbers {
		return 1
	}

	// 태그 라인을 찾는 간단한 휴리스틱
	tagCount := 0
	for i, line := range lines {
		if strings.Contains(line, `"tags"`) {
			// 태그 배열을 찾았으므로 계산 시작
			for j := i; j < len(lines); j++ {
				if strings.Contains(lines[j], `{`) {
					if tagCount == tagIndex {
						// 우리 태그를 찾았으므로 필드 찾기
						for k := j; k < len(lines) && k < j+10; k++ {
							if strings.Contains(lines[k], fmt.Sprintf(`"%s"`, fieldName)) {
								return k + 1
							}
						}
						return j + 1
					}
					tagCount++
				}
			}
		}
	}
	return 1
}
