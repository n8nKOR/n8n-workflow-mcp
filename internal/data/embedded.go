package data

import (
	"embed"
	"io/fs"
	"path/filepath"
	"strings"
)

// EmbeddedNodes는 빌드 시점에 internal/data/nodes 디렉토리의 모든 .md 파일을 임베딩합니다
//
//go:embed nodes/*.md
var EmbeddedNodes embed.FS

// EmbeddedWorkflows는 빌드 시점에 workflows.json 파일을 임베딩합니다
//
//go:embed workflows/workflows.json
var EmbeddedWorkflows embed.FS

// GetEmbeddedNodeFiles는 임베딩된 노드 파일들의 목록을 반환합니다
func GetEmbeddedNodeFiles() ([]string, error) {
	var files []string

	err := fs.WalkDir(EmbeddedNodes, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// ReadEmbeddedNodeFile은 임베딩된 노드 파일의 내용을 읽어 반환합니다
func ReadEmbeddedNodeFile(filename string) ([]byte, error) {
	return EmbeddedNodes.ReadFile(filename)
}

// GetEmbeddedNodeName은 파일 경로에서 노드명을 추출합니다
func GetEmbeddedNodeName(filePath string) string {
	fileName := filepath.Base(filePath)
	return strings.TrimSuffix(fileName, "_node.md")
}

// ReadEmbeddedWorkflowsFile은 임베딩된 워크플로우 JSON 파일의 내용을 읽어 반환합니다
func ReadEmbeddedWorkflowsFile() ([]byte, error) {
	return EmbeddedWorkflows.ReadFile("workflows/workflows.json")
}
