package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/i18n"
)

// SetupPingTool configures the ping tool
func SetupPingTool(s *server.MCPServer, i18nInstance *i18n.I18n) error {
	pingTool := mcp.NewTool("ping",
		mcp.WithDescription(i18nInstance.T("tool.ping.description")),
	)

	s.AddTool(pingTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("pong - 기본 기능 활성화"), nil
	})

	return nil
}
