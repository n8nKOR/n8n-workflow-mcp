package search

// NodeDocument 구조체는 n8n 노드 문서의 구조를 표현합니다.
type NodeDocument struct {
	NodeName         string           `json:"nodeName"`
	Overview         string           `json:"overview"`
	Credentials      []CredentialInfo `json:"credentials"`
	Inputs           []string         `json:"inputs"`
	Outputs          []string         `json:"outputs"`
	Resources        []ResourceInfo   `json:"resources"`
	FilePath         string           `json:"filePath"`
	Content          string           `json:"content"`
	UseCases         []string         `json:"useCases"`
	TokenizedContent []string         `json:"-"` // 인덱싱용 토큰화된 콘텐츠

	// 섹션별 토큰화된 콘텐츠
	TokenizedNodeName  []string `json:"-"`
	TokenizedOverview  []string `json:"-"`
	TokenizedUseCases  []string `json:"-"`
	TokenizedOperation []string `json:"-"`
}

// CredentialInfo 구조체는 인증 정보를 표현합니다.
type CredentialInfo struct {
	Name     string `json:"name"`
	Required bool   `json:"required"`
}

// ResourceInfo 구조체는 리소스 정보를 표현합니다.
type ResourceInfo struct {
	Name       string          `json:"name"`
	Operations []OperationInfo `json:"operations"`
}

// OperationInfo 구조체는 오퍼레이션 정보를 표현합니다.
type OperationInfo struct {
	Name   string      `json:"name"`
	Fields []FieldInfo `json:"fields"`
}

// FieldInfo 구조체는 필드 정보를 표현합니다.
type FieldInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

// SearchResult 구조체는 검색 결과를 표현합니다.
type SearchResult struct {
	NodeName  string         `json:"nodeName"`
	Score     float64        `json:"score"`
	Overview  string         `json:"overview"`
	Resources []ResourceInfo `json:"resources"`
}

// SearchResponse 구조체는 검색 API 응답을 표현합니다.
type SearchResponse struct {
	Query     string         `json:"query"`
	Results   []SearchResult `json:"results"`
	TotalHits int            `json:"totalHits"`
}

// SearchWeights 구조체는 검색 시 각 섹션별 가중치를 설정합니다.
type SearchWeights struct {
	NodeName  float64 `json:"nodeName"`
	Overview  float64 `json:"overview"`
	UseCases  float64 `json:"useCases"`
	Operation float64 `json:"operation"`
}

// DefaultSearchWeights 함수는 기본 검색 가중치를 반환합니다.
func DefaultSearchWeights() SearchWeights {
	return SearchWeights{
		NodeName:  3.0, // 노드명이 가장 중요
		Overview:  2.0, // 개요가 두 번째로 중요
		UseCases:  1.5, // 사용 사례가 세 번째로 중요
		Operation: 1.0, // 오퍼레이션이 네 번째로 중요
	}
}
