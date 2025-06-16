package types

// MCP 도구 요청/응답 타입들

// ValidateWorkflowRequest는 validate_workflow 도구의 요청을 나타냅니다
type ValidateWorkflowRequest struct {
	WorkflowID         string `json:"workflow_id" validate:"required"`
	IncludeLineNumbers bool   `json:"include_line_numbers,omitempty"`
}

// ValidateWorkflowResponse는 validate_workflow 도구의 응답을 나타냅니다
type ValidateWorkflowResponse struct {
	WorkflowID       string           `json:"workflow_id"`
	Validated        bool             `json:"validated"`
	ValidationResult ValidationResult `json:"validation_result"`
	Success          bool             `json:"success"`
	Message          string           `json:"message,omitempty"`
}

// MCPError는 표준화된 오류 응답을 나타냅니다
type MCPError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// 공통 오류 코드들
const (
	ErrorCodeInvalidArguments    = "INVALID_ARGUMENTS"
	ErrorCodeWorkflowNotFound    = "WORKFLOW_NOT_FOUND"
	ErrorCodeWorkflowNotApproved = "WORKFLOW_NOT_APPROVED"
	ErrorCodeAIServiceError      = "AI_SERVICE_ERROR"
	ErrorCodeValidatorError      = "VALIDATOR_ERROR"
	ErrorCodeInternalError       = "INTERNAL_ERROR"
)
