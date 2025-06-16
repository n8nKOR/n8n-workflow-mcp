package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// WorkflowStatus는 워크플로우의 상태를 나타냅니다
type WorkflowStatus string

const (
	StatusDraft     WorkflowStatus = "draft"
	StatusApproved  WorkflowStatus = "approved"
	StatusGenerated WorkflowStatus = "generated"
	StatusValidated WorkflowStatus = "validated"
	StatusError     WorkflowStatus = "error"
)

// Workflow는 워크플로우 인스턴스를 나타냅니다
type Workflow struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	Nodes       []string          `json:"nodes"`
	Mermaid     string            `json:"mermaid,omitempty"`
	JSON        string            `json:"json,omitempty"`
	Status      WorkflowStatus    `json:"status"`
	Errors      []ValidationError `json:"errors,omitempty"`
	FilePath    string            `json:"file_path,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// ValidationError는 줄 번호 정보가 포함된 검증 오류를 나타냅니다
type ValidationError struct {
	Line       int    `json:"line"`
	Column     int    `json:"column"`
	NodeID     string `json:"node_id,omitempty"`
	NodeName   string `json:"node_name,omitempty"`
	Field      string `json:"field,omitempty"`
	ErrorType  string `json:"error_type"`
	Message    string `json:"message"`
	Suggestion string `json:"suggestion,omitempty"`
	Context    string `json:"context,omitempty"`
}

// ValidationResult는 워크플로우 검증 결과를 나타냅니다
type ValidationResult struct {
	IsValid    bool              `json:"is_valid"`
	IsError    bool              `json:"is_error"`
	ErrorCount int               `json:"error_count"`
	Errors     []ValidationError `json:"errors,omitempty"`
	Summary    ValidationSummary `json:"summary,omitempty"`
}

// ValidationSummary는 검증 결과의 요약을 제공합니다
type ValidationSummary struct {
	TotalLines int            `json:"total_lines"`
	ErrorLines []int          `json:"error_lines"`
	ErrorTypes map[string]int `json:"error_types"`
}

// N8nNode는 n8n 워크플로우 노드를 나타냅니다
type N8nNode struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Position   []float64              `json:"position"`
	Parameters map[string]interface{} `json:"parameters"`
}

// N8nConnection은 노드 간의 연결을 나타냅니다
type N8nConnection struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	SourceIndex int    `json:"sourceIndex"`
	DestIndex   int    `json:"destIndex"`
}

// N8nTag는 n8n 워크플로우 태그를 나타냅니다
type N8nTag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// N8nWorkflow는 완전한 n8n 워크플로우를 나타냅니다
type N8nWorkflow struct {
	Name        string                 `json:"name"`
	Nodes       []N8nNode              `json:"nodes"`
	Connections map[string]interface{} `json:"connections"`
	Active      bool                   `json:"active"`
	Settings    map[string]interface{} `json:"settings,omitempty"`
	StaticData  map[string]interface{} `json:"staticData,omitempty"`
	Tags        []N8nTag               `json:"tags,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
}

// FileInfo는 파일 정보를 나타냅니다
type FileInfo struct {
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	Size     int64     `json:"size"`
	ModTime  time.Time `json:"mod_time"`
	IsExists bool      `json:"is_exists"`
}

// NewWorkflow는 새로운 워크플로우 인스턴스를 생성합니다
func NewWorkflow(description string) *Workflow {
	return &Workflow{
		ID:          uuid.New().String(),
		Description: description,
		Status:      StatusDraft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// UpdateStatus는 워크플로우 상태와 타임스탬프를 업데이트합니다
func (w *Workflow) UpdateStatus(status WorkflowStatus) {
	w.Status = status
	w.UpdatedAt = time.Now()
}

// AddError는 워크플로우에 검증 오류를 추가합니다
func (w *Workflow) AddError(err ValidationError) {
	w.Errors = append(w.Errors, err)
	w.UpdateStatus(StatusError)
}

// ClearErrors는 모든 검증 오류를 지웁니다
func (w *Workflow) ClearErrors() {
	w.Errors = nil
	if w.Status == StatusError {
		w.UpdateStatus(StatusDraft)
	}
}

// ToJSON은 워크플로우를 JSON 문자열로 변환합니다
func (w *Workflow) ToJSON() (string, error) {
	data, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
