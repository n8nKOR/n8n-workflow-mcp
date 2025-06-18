package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/handlers"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/i18n"
)

// SetupSearchWorkflowTool sets up the workflow search tool
func SetupSearchWorkflowTool(s *server.MCPServer, searchHandler *handlers.SearchHandler, i18nInstance *i18n.I18n) error {
	searchWorkflowTool := mcp.NewTool("search_workflow",
		mcp.WithDescription(i18nInstance.T("tool.search_workflow.description")),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description(i18nInstance.T("tool.search_workflow.param.query.description")),
		),
		mcp.WithNumber("max_results",
			mcp.Description(i18nInstance.T("tool.search_workflow.param.max_results.description")),
		),
	)

	s.AddTool(searchWorkflowTool, searchHandler.HandleSearchWorkflows)
	return nil
}
