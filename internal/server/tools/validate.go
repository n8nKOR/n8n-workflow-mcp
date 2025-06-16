package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/handlers"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/i18n"
)

// SetupValidateWorkflowTool은 validate_workflow 도구를 설정합니다
func SetupValidateWorkflowTool(s *server.MCPServer, handler *handlers.Handler, i18nInstance *i18n.I18n) error {
	validateTool := mcp.NewTool("validate_workflow",
		mcp.WithDescription(i18nInstance.T("tool.validate_workflow.description")),
		mcp.WithString("file_path",
			mcp.Required(),
			mcp.Description(i18nInstance.T("param.file_path.description")),
		),
		mcp.WithBoolean("include_line_numbers",
			mcp.Description(i18nInstance.T("param.include_line_numbers.description")),
		),
		mcp.WithBoolean("auto_fix_suggestions",
			mcp.Description(i18nInstance.T("param.auto_fix_suggestions.description")),
		),
	)
	s.AddTool(validateTool, handler.HandleValidateWorkflow)
	return nil
}
