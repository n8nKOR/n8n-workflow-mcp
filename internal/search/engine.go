package search

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/data"
)

// SearchEngine 구조체는 검색 엔진을 표현합니다.
type SearchEngine struct {
	Documents     []NodeDocument
	DocumentMap   map[string]*NodeDocument
	Tokenizer     func(string) []string
	SearchWeights SearchWeights
	mutex         sync.RWMutex

	// 섹션별 BM25 점수 계산기
	bm25NodeName  *BM25Scorer
	bm25Overview  *BM25Scorer
	bm25UseCases  *BM25Scorer
	bm25Operation *BM25Scorer
}

// NewSearchEngine 함수는 새로운 검색 엔진을 생성합니다.
func NewSearchEngine() *SearchEngine {
	return &SearchEngine{
		Documents:     make([]NodeDocument, 0),
		DocumentMap:   make(map[string]*NodeDocument),
		Tokenizer:     tokenize,
		SearchWeights: DefaultSearchWeights(),
	}
}

// SetSearchWeights 함수는 검색 가중치를 설정합니다.
func (se *SearchEngine) SetSearchWeights(weights SearchWeights) {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	se.SearchWeights = weights
}

// tokenize 함수는 텍스트를 토큰화합니다.
func tokenize(text string) []string {
	// 소문자 변환
	text = strings.ToLower(text)

	// 특수문자 제거 및 공백으로 대체
	reg := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	text = reg.ReplaceAllString(text, " ")

	// 카멜케이스 분리 (예: ActionNetwork -> Action Network)
	camelCase := regexp.MustCompile(`([a-z])([A-Z])`)
	text = camelCase.ReplaceAllString(text, "$1 $2")

	// 토큰화
	tokens := strings.Fields(text)

	// 불용어 제거
	stopWords := map[string]bool{
		"a": true, "an": true, "the": true,
		"and": true, "or": true, "but": true,
		"is": true, "are": true, "was": true,
		"were": true, "be": true, "been": true,
		"to": true, "of": true, "in": true,
		"for": true, "with": true, "by": true,
		"at": true, "on": true, "from": true,
	}

	filteredTokens := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if !stopWords[token] && len(token) > 1 {
			filteredTokens = append(filteredTokens, token)
		}
	}

	return filteredTokens
}

// LoadEmbeddedDocuments 함수는 임베딩된 마크다운 문서를 로드하고 파싱합니다.
func (se *SearchEngine) LoadEmbeddedDocuments() error {
	se.mutex.Lock()
	defer se.mutex.Unlock()

	files, err := data.GetEmbeddedNodeFiles()
	if err != nil {
		return fmt.Errorf("임베딩된 파일 목록 조회 실패: %w", err)
	}

	for _, filePath := range files {
		content, err := data.ReadEmbeddedNodeFile(filePath)
		if err != nil {
			log.Printf("임베딩된 파일 읽기 실패 %s: %v", filePath, err)
			continue
		}

		doc, err := parseMarkdown(string(content), filePath)
		if err != nil {
			log.Printf("마크다운 파싱 실패 %s: %v", filePath, err)
			continue
		}

		se.Documents = append(se.Documents, doc)
		se.DocumentMap[doc.NodeName] = &se.Documents[len(se.Documents)-1]
	}

	// 섹션별 토큰화된 콘텐츠 생성
	for i, doc := range se.Documents {
		// 기존 전체 콘텐츠 토큰화
		se.Documents[i].TokenizedContent = se.Tokenizer(doc.Content)

		// 섹션별 토큰화
		se.Documents[i].TokenizedNodeName = se.Tokenizer(doc.NodeName)
		se.Documents[i].TokenizedOverview = se.Tokenizer(doc.Overview)

		// UseCases 토큰화
		useCasesText := strings.Join(doc.UseCases, " ")
		se.Documents[i].TokenizedUseCases = se.Tokenizer(useCasesText)

		// Operation 토큰화 (Resources의 모든 Operation 정보 포함)
		var operationText strings.Builder
		for _, resource := range doc.Resources {
			for _, operation := range resource.Operations {
				operationText.WriteString(operation.Name)
				operationText.WriteString(" ")
				for _, field := range operation.Fields {
					operationText.WriteString(field.Name)
					operationText.WriteString(" ")
					operationText.WriteString(field.Description)
					operationText.WriteString(" ")
				}
			}
		}
		se.Documents[i].TokenizedOperation = se.Tokenizer(operationText.String())
	}

	// 섹션별 BM25 점수 계산기 초기화
	se.initializeSectionBM25()

	return nil
}

// LoadDocuments 함수는 마크다운 문서를 로드하고 파싱합니다. (기존 파일 시스템 방식)
func (se *SearchEngine) LoadDocuments(dirPath string) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		filePath := filepath.Join(dirPath, entry.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filePath, err)
			continue
		}

		doc, err := parseMarkdown(string(content), filePath)
		if err != nil {
			log.Printf("Error parsing markdown %s: %v", filePath, err)
			continue
		}

		se.Documents = append(se.Documents, doc)
		se.DocumentMap[doc.NodeName] = &se.Documents[len(se.Documents)-1]
	}

	// 섹션별 토큰화된 콘텐츠 생성
	for i, doc := range se.Documents {
		// 기존 전체 콘텐츠 토큰화
		se.Documents[i].TokenizedContent = se.Tokenizer(doc.Content)

		// 섹션별 토큰화
		se.Documents[i].TokenizedNodeName = se.Tokenizer(doc.NodeName)
		se.Documents[i].TokenizedOverview = se.Tokenizer(doc.Overview)

		// UseCases 토큰화
		useCasesText := strings.Join(doc.UseCases, " ")
		se.Documents[i].TokenizedUseCases = se.Tokenizer(useCasesText)

		// Operation 토큰화 (Resources의 모든 Operation 정보 포함)
		var operationText strings.Builder
		for _, resource := range doc.Resources {
			for _, operation := range resource.Operations {
				operationText.WriteString(operation.Name)
				operationText.WriteString(" ")
				for _, field := range operation.Fields {
					operationText.WriteString(field.Name)
					operationText.WriteString(" ")
					operationText.WriteString(field.Description)
					operationText.WriteString(" ")
				}
			}
		}
		se.Documents[i].TokenizedOperation = se.Tokenizer(operationText.String())
	}

	// 섹션별 BM25 점수 계산기 초기화
	se.initializeSectionBM25()

	return nil
}

// initializeSectionBM25 섹션별 BM25 점수 계산기를 초기화합니다.
func (se *SearchEngine) initializeSectionBM25() {
	if len(se.Documents) == 0 {
		return
	}

	params := DefaultBM25Parameters()

	// NodeName 토큰화된 문서들
	nodeNameDocs := make([][]string, len(se.Documents))
	for i, doc := range se.Documents {
		nodeNameDocs[i] = doc.TokenizedNodeName
	}
	se.bm25NodeName = NewBM25Scorer(nodeNameDocs, params)

	// Overview 토큰화된 문서들
	overviewDocs := make([][]string, len(se.Documents))
	for i, doc := range se.Documents {
		overviewDocs[i] = doc.TokenizedOverview
	}
	se.bm25Overview = NewBM25Scorer(overviewDocs, params)

	// UseCases 토큰화된 문서들
	useCasesDocs := make([][]string, len(se.Documents))
	for i, doc := range se.Documents {
		useCasesDocs[i] = doc.TokenizedUseCases
	}
	se.bm25UseCases = NewBM25Scorer(useCasesDocs, params)

	// Operation 토큰화된 문서들
	operationDocs := make([][]string, len(se.Documents))
	for i, doc := range se.Documents {
		operationDocs[i] = doc.TokenizedOperation
	}
	se.bm25Operation = NewBM25Scorer(operationDocs, params)
}

// calculateWeightedBM25 함수는 섹션별 가중치를 적용한 BM25 점수를 계산합니다.
func (se *SearchEngine) calculateWeightedBM25(queryTokens []string, docIndex int) float64 {
	var totalScore float64

	if se.bm25NodeName != nil {
		totalScore += se.bm25NodeName.Score(queryTokens, docIndex) * se.SearchWeights.NodeName
	}
	if se.bm25Overview != nil {
		totalScore += se.bm25Overview.Score(queryTokens, docIndex) * se.SearchWeights.Overview
	}
	if se.bm25UseCases != nil {
		totalScore += se.bm25UseCases.Score(queryTokens, docIndex) * se.SearchWeights.UseCases
	}
	if se.bm25Operation != nil {
		totalScore += se.bm25Operation.Score(queryTokens, docIndex) * se.SearchWeights.Operation
	}

	return totalScore
}

// Search 함수는 쿼리에 맞는 문서를 검색합니다.
func (se *SearchEngine) Search(query string, topN int) (SearchResponse, error) {
	se.mutex.RLock()
	defer se.mutex.RUnlock()

	if len(se.Documents) == 0 {
		return SearchResponse{}, fmt.Errorf("search engine not initialized")
	}

	tokenizedQuery := se.Tokenizer(query)

	type indexScore struct {
		Index int
		Score float64
	}

	indexScores := make([]indexScore, len(se.Documents))
	for i := range se.Documents {
		indexScores[i] = indexScore{Index: i, Score: se.calculateWeightedBM25(tokenizedQuery, i)}
	}

	sort.Slice(indexScores, func(i, j int) bool {
		return indexScores[i].Score > indexScores[j].Score
	})

	resultCount := topN
	if resultCount > len(indexScores) {
		resultCount = len(indexScores)
	}

	var results []SearchResult
	for i := 0; i < resultCount; i++ {
		if indexScores[i].Score <= 0 {
			continue
		}
		doc := se.Documents[indexScores[i].Index]
		results = append(results, SearchResult{
			NodeName:  doc.NodeName,
			Score:     indexScores[i].Score,
			Overview:  doc.Overview,
			Resources: doc.Resources,
		})
	}

	return SearchResponse{
		Query:     query,
		Results:   results,
		TotalHits: len(results),
	}, nil
}
