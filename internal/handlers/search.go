package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/services"
)

// SearchHandler는 검색 관련 MCP 도구 핸들러를 관리합니다
type SearchHandler struct {
	searchService *services.SearchService
}

// NewSearchHandler는 새로운 검색 핸들러 인스턴스를 생성합니다
func NewSearchHandler(searchService *services.SearchService) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}

// HandleSearchNodes는 n8n 노드 검색을 처리합니다
func (h *SearchHandler) HandleSearchNodes(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 필수 파라미터 추출
	query, err := request.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError("query 파라미터가 필요합니다"), nil
	}

	// limit 파라미터 추출 (선택적)
	limit := request.GetInt("limit", 5)
	if limit > 20 {
		limit = 20 // 최대값 제한
	}

	// 검색 수행
	response, err := h.searchService.SearchNodes(query, limit)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("검색 실행 중 오류 발생: %v", err)), nil
	}

	// 결과를 JSON으로 변환
	resultJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("결과 변환 중 오류 발생: %v", err)), nil
	}

	return mcp.NewToolResultText(string(resultJSON)), nil
}
