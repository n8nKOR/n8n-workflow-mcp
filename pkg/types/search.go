package types

// WorkflowTemplate represents a workflow template from community
type WorkflowTemplate struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	URL              string   `json:"url"`
	Category         string   `json:"category"`
	Tags             []string `json:"tags,omitempty"`
	TokenizedContent []string `json:"-"` // For indexing purposes
}

// WorkflowSearchResult represents a single workflow search result with score
type WorkflowSearchResult struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags,omitempty"`
	Score       float64  `json:"score"`
}

// WorkflowSearchResponse represents the complete search response
type WorkflowSearchResponse struct {
	Query      string                 `json:"query"`
	Results    []WorkflowSearchResult `json:"results"`
	TotalHits  int                    `json:"totalHits"`
	MaxResults int                    `json:"maxResults"`
}

// WorkflowSearchRequest represents a workflow search request
type WorkflowSearchRequest struct {
	Query      string `json:"query"`
	MaxResults int    `json:"maxResults,omitempty"`
}
