package services

import (
	"path/filepath"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/search"
)

// SearchService 검색 서비스 구조체
type SearchService struct {
	engine *search.SearchEngine
}

// NewSearchService 새로운 검색 서비스를 생성합니다.
func NewSearchService() *SearchService {
	return &SearchService{
		engine: search.NewSearchEngine(),
	}
}

// InitializeEmbedded 임베딩된 데이터로 검색 서비스를 초기화합니다.
func (s *SearchService) InitializeEmbedded(weights search.SearchWeights) error {
	// 가중치 설정
	s.engine.SetSearchWeights(weights)

	// 임베딩된 문서 로드
	return s.engine.LoadEmbeddedDocuments()
}

// InitializeWithWeights 가중치를 설정하고 문서를 로드합니다.
func (s *SearchService) InitializeWithWeights(weights search.SearchWeights, dataPath string) error {
	// 가중치 설정
	s.engine.SetSearchWeights(weights)

	// 문서 로드
	return s.engine.LoadDocuments(dataPath)
}

// Initialize 기본 가중치로 문서를 로드합니다.
func (s *SearchService) Initialize(dataPath string) error {
	return s.engine.LoadDocuments(dataPath)
}

// SearchNodes 노드를 검색합니다.
func (s *SearchService) SearchNodes(query string, limit int) (search.SearchResponse, error) {
	if limit <= 0 {
		limit = 5 // 기본값
	}
	if limit > 20 {
		limit = 20 // 최대값 제한
	}

	return s.engine.Search(query, limit)
}

// SetWeights 검색 가중치를 업데이트합니다.
func (s *SearchService) SetWeights(weights search.SearchWeights) {
	s.engine.SetSearchWeights(weights)
}

// GetDefaultDataPath 기본 데이터 경로를 반환합니다.
func GetDefaultDataPath() string {
	return filepath.Join("internal", "data", "nodes")
}
