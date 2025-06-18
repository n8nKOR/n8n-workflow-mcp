package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/data"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/search"
	"github.com/n8nKOR/n8n-workflow-mcp/pkg/types"
)

// WorkflowSearchService provides workflow search functionality
type WorkflowSearchService struct {
	workflows   []types.WorkflowTemplate
	workflowMap map[string]*types.WorkflowTemplate
	bm25        *search.BM25Scorer
	tokenizer   func(string) []string
	mutex       sync.RWMutex
}

// NewWorkflowSearchService creates a new workflow search service
func NewWorkflowSearchService() *WorkflowSearchService {
	return &WorkflowSearchService{
		workflows:   make([]types.WorkflowTemplate, 0),
		workflowMap: make(map[string]*types.WorkflowTemplate),
		tokenizer:   tokenizeText,
	}
}

// LoadWorkflowsFromFile loads workflow templates from a JSON file
func (ws *WorkflowSearchService) LoadWorkflowsFromFile(filePath string) error {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	if !filepath.IsAbs(filePath) {
		wd, _ := os.Getwd()
		filePath = filepath.Join(wd, filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read workflows file: %w", err)
	}

	var workflows []types.WorkflowTemplate
	if err := json.Unmarshal(data, &workflows); err != nil {
		return fmt.Errorf("failed to parse workflows JSON: %w", err)
	}

	ws.workflows = workflows
	ws.workflowMap = make(map[string]*types.WorkflowTemplate)

	// Tokenize content and build map
	for i := range ws.workflows {
		ws.workflows[i].TokenizedContent = ws.tokenizer(ws.workflows[i].Description)
		ws.workflowMap[ws.workflows[i].Name] = &ws.workflows[i]
	}

	// Initialize BM25 scorer
	ws.initializeBM25()

	return nil
}

// LoadEmbeddedWorkflows loads workflow templates from embedded data
func (ws *WorkflowSearchService) LoadEmbeddedWorkflows() error {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	workflowData, err := data.ReadEmbeddedWorkflowsFile()
	if err != nil {
		return fmt.Errorf("failed to read embedded workflows file: %w", err)
	}

	var workflows []types.WorkflowTemplate
	if err := json.Unmarshal(workflowData, &workflows); err != nil {
		return fmt.Errorf("failed to parse workflows JSON: %w", err)
	}

	ws.workflows = workflows
	ws.workflowMap = make(map[string]*types.WorkflowTemplate)

	// Tokenize content and build map
	for i := range ws.workflows {
		ws.workflows[i].TokenizedContent = ws.tokenizer(ws.workflows[i].Description)
		ws.workflowMap[ws.workflows[i].Name] = &ws.workflows[i]
	}

	// Initialize BM25 scorer
	ws.initializeBM25()

	return nil
}

// initializeBM25 initializes the BM25 scorer for workflows
func (ws *WorkflowSearchService) initializeBM25() {
	if len(ws.workflows) == 0 {
		return
	}

	documents := make([][]string, len(ws.workflows))
	for i, workflow := range ws.workflows {
		documents[i] = workflow.TokenizedContent
	}

	ws.bm25 = search.NewBM25Scorer(documents, search.DefaultBM25Parameters())
}

// SearchWorkflows searches for workflows matching the query
func (ws *WorkflowSearchService) SearchWorkflows(query string, maxResults int) (*types.WorkflowSearchResponse, error) {
	ws.mutex.RLock()
	defer ws.mutex.RUnlock()

	if len(ws.workflows) == 0 {
		return nil, fmt.Errorf("no workflows loaded")
	}

	// Default max results
	if maxResults <= 0 {
		maxResults = 10
	}

	// Tokenize query
	tokenizedQuery := ws.tokenizer(query)

	type indexScore struct {
		Index int
		Score float64
	}

	var indexScores []indexScore

	// Score all workflows
	for i := range ws.workflows {
		score := ws.bm25.Score(tokenizedQuery, i)
		if score > 0 {
			indexScores = append(indexScores, indexScore{Index: i, Score: score})
		}
	}

	// Sort by score (descending)
	sort.Slice(indexScores, func(i, j int) bool {
		return indexScores[i].Score > indexScores[j].Score
	})

	// Limit results
	resultCount := maxResults
	if resultCount > len(indexScores) {
		resultCount = len(indexScores)
	}

	var results []types.WorkflowSearchResult
	for i := 0; i < resultCount; i++ {
		workflow := ws.workflows[indexScores[i].Index]
		results = append(results, types.WorkflowSearchResult{
			Name:        workflow.Name,
			Description: workflow.Description,
			URL:         workflow.URL,
			Category:    workflow.Category,
			Tags:        workflow.Tags,
			Score:       indexScores[i].Score,
		})
	}

	return &types.WorkflowSearchResponse{
		Query:      query,
		Results:    results,
		TotalHits:  len(indexScores),
		MaxResults: maxResults,
	}, nil
}

// GetWorkflowCount returns the total number of loaded workflows
func (ws *WorkflowSearchService) GetWorkflowCount() int {
	ws.mutex.RLock()
	defer ws.mutex.RUnlock()
	return len(ws.workflows)
}

// tokenizeText tokenizes text into terms for search indexing
func tokenizeText(text string) []string {
	// Convert to lowercase
	text = strings.ToLower(text)

	// Replace special characters with spaces
	reg := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	text = reg.ReplaceAllString(text, " ")

	// Split camelCase (e.g., ActionNetwork -> Action Network)
	camelCase := regexp.MustCompile(`([a-z])([A-Z])`)
	text = camelCase.ReplaceAllString(text, "$1 $2")
	text = strings.ToLower(text)

	// Tokenize
	tokens := strings.Fields(text)

	// Remove stop words
	stopWords := map[string]bool{
		"a": true, "an": true, "the": true,
		"and": true, "or": true, "but": true,
		"in": true, "on": true, "at": true,
		"to": true, "for": true, "of": true,
		"with": true, "by": true, "from": true,
		"is": true, "are": true, "was": true,
		"were": true, "be": true, "been": true,
		"have": true, "has": true, "had": true,
		"do": true, "does": true, "did": true,
	}

	filteredTokens := make([]string, 0, len(tokens))
	for _, token := range tokens {
		token = strings.ToLower(token)
		if !stopWords[token] && len(token) > 1 {
			filteredTokens = append(filteredTokens, token)
		}
	}

	return filteredTokens
}
