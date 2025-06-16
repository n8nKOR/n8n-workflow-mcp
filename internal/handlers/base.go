package handlers

import (
	"github.com/n8nKOR/n8n-workflow-mcp/internal/workflow"
)

// Handler는 모든 MCP 도구 핸들러를 관리합니다
type Handler struct {
	workflowService *workflow.Service
}

// NewHandler는 새로운 핸들러 인스턴스를 생성합니다
func NewHandler(workflowService *workflow.Service) *Handler {
	return &Handler{
		workflowService: workflowService,
	}
}
