package workflow

import (
	"github.com/n8nKOR/n8n-workflow-mcp/internal/store"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/validator"
	"github.com/n8nKOR/n8n-workflow-mcp/pkg/types"
)

// Service는 워크플로우 작업을 처리합니다
type Service struct {
	fileStore *store.FileStore
	validator *validator.N8nValidator
}

// NewService는 새로운 워크플로우 서비스를 생성합니다
func NewService(validator *validator.N8nValidator) *Service {
	return &Service{
		fileStore: store.NewFileStore(),
		validator: validator,
	}
}

// ValidateWorkflow는 n8n JSON 워크플로우를 검증합니다
func (s *Service) ValidateWorkflow(jsonStr string, includeLineNumbers bool) (*types.ValidationResult, error) {
	// 적절한 설정으로 검증기 생성
	v := validator.NewN8nValidator(includeLineNumbers, 50)

	// 워크플로우 검증
	return v.ValidateWorkflow(jsonStr)
}

// ValidateWorkflowFromFile은 파일에서 n8n JSON 워크플로우를 읽어서 검증합니다
func (s *Service) ValidateWorkflowFromFile(filePath string, includeLineNumbers bool) (*types.ValidationResult, error) {
	// 파일 경로 검증
	if err := s.fileStore.ValidateFilePath(filePath); err != nil {
		return nil, err
	}

	// 파일에서 워크플로우 읽기
	workflow, err := s.fileStore.ReadWorkflowFromFile(filePath)
	if err != nil {
		return nil, err
	}

	// 워크플로우 검증
	return s.ValidateWorkflow(workflow.JSON, includeLineNumbers)
}

// GetFileInfo는 파일 정보를 반환합니다
func (s *Service) GetFileInfo(filePath string) (*types.FileInfo, error) {
	return s.fileStore.GetFileInfo(filePath)
}
