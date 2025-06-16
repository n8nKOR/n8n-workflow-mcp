package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/n8nKOR/n8n-workflow-mcp/pkg/types"
)

// FileStore는 JSON 파일을 직접 처리하는 워크플로우 저장소입니다
type FileStore struct{}

// NewFileStore는 새로운 파일 기반 워크플로우 저장소를 생성합니다
func NewFileStore() *FileStore {
	return &FileStore{}
}

// ValidateFilePath는 파일 경로의 유효성을 검증합니다
func (fs *FileStore) ValidateFilePath(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	// 파일 존재 여부 확인
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	// JSON 파일인지 확인
	if ext := filepath.Ext(filePath); ext != ".json" {
		return fmt.Errorf("file must be a JSON file, got: %s", ext)
	}

	return nil
}

// ReadWorkflowFromFile은 JSON 파일에서 워크플로우를 읽습니다
func (fs *FileStore) ReadWorkflowFromFile(filePath string) (*types.Workflow, error) {
	// 파일 경로 검증
	if err := fs.ValidateFilePath(filePath); err != nil {
		return nil, err
	}

	// 파일 읽기
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// JSON이 비어있는지 확인
	if len(data) == 0 {
		return nil, fmt.Errorf("file is empty: %s", filePath)
	}

	// JSON 유효성 검증
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("invalid JSON in file %s: %w", filePath, err)
	}

	// 절대 경로 구하기
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		absPath = filePath
	}

	// Workflow 객체 생성 (JSON 내용과 파일 정보 포함)
	workflow := &types.Workflow{
		ID:          filepath.Base(filePath), // 파일명을 ID로 사용
		Description: fmt.Sprintf("Workflow from file: %s", filepath.Base(filePath)),
		JSON:        string(data),
		Status:      types.StatusGenerated,
		FilePath:    absPath,
	}

	return workflow, nil
}

// WriteWorkflowToFile은 워크플로우를 JSON 파일에 저장합니다
func (fs *FileStore) WriteWorkflowToFile(filePath string, jsonContent string) error {
	if filePath == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	if jsonContent == "" {
		return fmt.Errorf("JSON content cannot be empty")
	}

	// JSON 유효성 검증
	var jsonData interface{}
	if err := json.Unmarshal([]byte(jsonContent), &jsonData); err != nil {
		return fmt.Errorf("invalid JSON content: %w", err)
	}

	// 디렉토리 생성 (필요한 경우)
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// 파일에 JSON 저장
	if err := os.WriteFile(filePath, []byte(jsonContent), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

// GetFileInfo는 파일의 기본 정보를 반환합니다
func (fs *FileStore) GetFileInfo(filePath string) (*types.FileInfo, error) {
	if err := fs.ValidateFilePath(filePath); err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		absPath = filePath
	}

	return &types.FileInfo{
		Name:     fileInfo.Name(),
		Path:     absPath,
		Size:     fileInfo.Size(),
		ModTime:  fileInfo.ModTime(),
		IsExists: true,
	}, nil
}
