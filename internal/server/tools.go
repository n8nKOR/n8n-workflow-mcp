package server

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/handlers"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/i18n"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/server/tools"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/services"
)

// SetupTools configures all available tools for the MCP server
func SetupTools(s *server.MCPServer, handler *handlers.Handler, searchService *services.SearchService, i18nInstance *i18n.I18n) error {
	// ping 도구 추가
	if err := tools.SetupPingTool(s, i18nInstance); err != nil {
		return fmt.Errorf("ping 도구 설정 실패: %w", err)
	}

	// validate_workflow 도구 추가
	if err := tools.SetupValidateWorkflowTool(s, handler, i18nInstance); err != nil {
		return fmt.Errorf("validate_workflow 도구 설정 실패: %w", err)
	}

	// 검색 핸들러 생성
	searchHandler := handlers.NewSearchHandler(searchService)

	// n8n 노드 검색 도구 추가
	if err := tools.SetupSearchTool(s, searchHandler, i18nInstance); err != nil {
		return fmt.Errorf("search_n8n_nodes 도구 설정 실패: %w", err)
	}

	return nil
}
