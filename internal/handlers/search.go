package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/i18n"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/services"
)

// SearchHandler는 검색 관련 MCP 도구 핸들러를 관리합니다
type SearchHandler struct {
	searchService         *services.SearchService
	workflowSearchService *services.WorkflowSearchService
}

// NewSearchHandler는 새로운 검색 핸들러 인스턴스를 생성합니다
func NewSearchHandler(searchService *services.SearchService, workflowSearchService *services.WorkflowSearchService) *SearchHandler {
	return &SearchHandler{
		searchService:         searchService,
		workflowSearchService: workflowSearchService,
	}
}

// WorkflowSearchResultItem represents a single workflow result with URL and description
type WorkflowSearchResultItem struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

// WorkflowSearchResponseData represents the complete search response with next_step
type WorkflowSearchResponseData struct {
	Results  []WorkflowSearchResultItem `json:"results"`
	NextStep string                     `json:"next_step"`
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

// HandleSearchWorkflows는 워크플로우 검색을 처리합니다
func (h *SearchHandler) HandleSearchWorkflows(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 필수 파라미터 추출
	query, err := request.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError("query 파라미터가 필요합니다"), nil
	}

	// max_results 파라미터 추출 (선택적)
	maxResults := request.GetInt("max_results", 5)
	if maxResults > 15 {
		maxResults = 15 // 최대값 제한
	}

	// 워크플로우 검색 서비스 상태 확인
	if h.workflowSearchService == nil {
		return mcp.NewToolResultError("워크플로우 검색 서비스가 초기화되지 않았습니다"), nil
	}

	// 로드된 워크플로우 수 확인
	workflowCount := h.workflowSearchService.GetWorkflowCount()
	if workflowCount == 0 {
		return mcp.NewToolResultError(fmt.Sprintf("워크플로우 데이터가 로드되지 않았습니다. 현재 로드된 워크플로우 수: %d. 시스템 로그를 확인해주세요.", workflowCount)), nil
	}

	// 검색 수행
	response, err := h.workflowSearchService.SearchWorkflows(query, maxResults)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("워크플로우 검색 실행 중 오류 발생: %v", err)), nil
	}

	// URL과 description 추출
	results := make([]WorkflowSearchResultItem, len(response.Results))
	for i, result := range response.Results {
		results[i] = WorkflowSearchResultItem{
			URL:         result.URL,
			Description: result.Description,
		}
	}

	// i18n 인스턴스 생성 (기본적으로 한국어 사용)
	i18nInstance := i18n.New("ko")

	// next_step 메시지 가져오기
	nextStepMessage := i18nInstance.T("search_workflow.next_step")

	// 최종 응답 데이터 구성
	responseData := WorkflowSearchResponseData{
		Results:  results,
		NextStep: nextStepMessage,
	}

	// 결과를 JSON으로 변환
	resultJSON, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("결과 변환 중 오류 발생: %v", err)), nil
	}

	return mcp.NewToolResultText(string(resultJSON)), nil
}
