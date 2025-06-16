package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/handlers"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/i18n"
)

// SetupSearchTool n8n 노드 검색 도구를 설정합니다.
func SetupSearchTool(s *server.MCPServer, searchHandler *handlers.SearchHandler, i18nInstance *i18n.I18n) error {
	searchTool := mcp.NewTool("search_n8n_nodes",
		mcp.WithDescription(i18nInstance.T("tool.search_n8n_nodes.description")),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description(i18nInstance.T("param.query.description")),
		),
		mcp.WithNumber("limit",
			mcp.Description(i18nInstance.T("param.limit.description")),
		),
	)

	s.AddTool(searchTool, searchHandler.HandleSearchNodes)
	return nil
}
