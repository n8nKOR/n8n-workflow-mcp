package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/n8nKOR/n8n-workflow-mcp/pkg/types"
)

// HandleValidateWorkflow는 JSON 파일 경로로 n8n 워크플로우를 검증합니다
func (h *Handler) HandleValidateWorkflow(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 매개변수 추출
	filePath, err := request.RequireString("file_path")
	if err != nil {
		return mcp.NewToolResultError("file_path parameter is required"), nil
	}

	includeLineNumbers := request.GetBool("include_line_numbers", true)

	// 파일 존재 여부 및 유효성 검증
	fileInfo, err := h.workflowService.GetFileInfo(filePath)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("File validation failed: %v", err)), nil
	}

	// 파일에서 워크플로우 검증 실행
	result, err := h.workflowService.ValidateWorkflowFromFile(filePath, includeLineNumbers)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Validation failed: %v", err)), nil
	}

	// 응답 형식 지정
	response := types.ValidateWorkflowResponse{
		WorkflowID:       fileInfo.Name,
		Validated:        result.IsValid,
		ValidationResult: *result,
		Success:          result.IsValid,
		Message:          "Workflow validation completed",
	}

	if result.IsValid {
		response.Message = fmt.Sprintf("✅ Workflow validation passed for file: %s", fileInfo.Name)
	} else {
		response.Message = fmt.Sprintf("❌ Workflow validation failed with %d errors for file: %s", len(result.Errors), fileInfo.Name)
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to format response: %v", err)), nil
	}

	return mcp.NewToolResultText(string(responseJSON)), nil
}
