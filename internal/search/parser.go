package search

import (
	"path/filepath"
	"strings"
)

// parseMarkdown 함수는 마크다운 문서를 파싱하여 NodeDocument 구조체로 변환합니다.
func parseMarkdown(content string, filePath string) (NodeDocument, error) {
	doc := NodeDocument{
		FilePath: filePath,
		Content:  content,
	}

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			doc.NodeName = strings.TrimSpace(strings.TrimPrefix(trimmed, "# "))
			break
		}
	}

	if doc.NodeName == "" {
		fileName := filepath.Base(filePath)
		doc.NodeName = strings.TrimSuffix(fileName, "_node.md")
	}

	section := ""
	var currentResource *ResourceInfo
	var currentOperation *OperationInfo

	// 섹션별 콘텐츠를 저장할 임시 변수들
	var overviewContent strings.Builder
	var useCasesContent strings.Builder
	var operationContent strings.Builder

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// 섹션 헤더 확인
		if strings.HasPrefix(line, "## ") {
			section = strings.TrimSpace(strings.TrimPrefix(line, "## "))
			continue
		}

		// 섹션별 처리
		switch section {
		case "Overview":
			if i+1 < len(lines) && len(lines[i+1]) > 0 && !strings.HasPrefix(lines[i+1], "#") {
				doc.Overview = lines[i+1]
				overviewContent.WriteString(lines[i+1])
				overviewContent.WriteString(" ")
			}

		case "Credentials":
			if strings.HasPrefix(line, "- Name:") {
				parts := strings.Split(line, ",")
				if len(parts) >= 2 {
					name := strings.TrimSpace(strings.TrimPrefix(parts[0], "- Name:"))
					required := strings.Contains(parts[1], "Yes")
					doc.Credentials = append(doc.Credentials, CredentialInfo{
						Name:     name,
						Required: required,
					})
				}
			}

		case "Inputs":
			if strings.HasPrefix(line, "- ") {
				input := strings.TrimSpace(strings.TrimPrefix(line, "- "))
				if input != "None" {
					doc.Inputs = append(doc.Inputs, input)
				}
			}

		case "Outputs":
			if strings.HasPrefix(line, "- ") {
				output := strings.TrimSpace(strings.TrimPrefix(line, "- "))
				if output != "None" {
					doc.Outputs = append(doc.Outputs, output)
				}
			}

		case "UseCases":
			if strings.HasPrefix(line, "- ") {
				useCase := strings.TrimSpace(strings.TrimPrefix(line, "- "))
				if useCase != "" {
					doc.UseCases = append(doc.UseCases, useCase)
					useCasesContent.WriteString(useCase)
					useCasesContent.WriteString(" ")
				}
			}

		case "Properties":
			// 리소스 섹션 확인
			if strings.HasPrefix(line, "### Resource:") {
				resourceName := strings.TrimSpace(strings.TrimPrefix(line, "### Resource:"))
				doc.Resources = append(doc.Resources, ResourceInfo{
					Name:       resourceName,
					Operations: make([]OperationInfo, 0),
				})
				currentResource = &doc.Resources[len(doc.Resources)-1]
				currentOperation = nil

			} else if strings.HasPrefix(line, "#### Operation:") && currentResource != nil {
				operationName := strings.TrimSpace(strings.TrimPrefix(line, "#### Operation:"))
				currentResource.Operations = append(currentResource.Operations, OperationInfo{
					Name:   operationName,
					Fields: make([]FieldInfo, 0),
				})
				currentOperation = &currentResource.Operations[len(currentResource.Operations)-1]

				// 오퍼레이션 콘텐츠 수집
				operationContent.WriteString(operationName)
				operationContent.WriteString(" ")

			} else if strings.Contains(line, "|") && currentOperation != nil {
				// 테이블 헤더와 구분선 건너뛰기
				if strings.Contains(line, "Field Name") || strings.Contains(line, "---") {
					continue
				}

				// 필드 정보 파싱
				cells := strings.Split(line, "|")
				if len(cells) >= 5 {
					fieldName := strings.TrimSpace(cells[1])
					fieldType := strings.TrimSpace(cells[2])
					fieldDesc := strings.TrimSpace(cells[3])
					fieldRequired := strings.Contains(strings.TrimSpace(cells[4]), "Yes")

					if fieldName != "" {
						currentOperation.Fields = append(currentOperation.Fields, FieldInfo{
							Name:        fieldName,
							Type:        fieldType,
							Description: fieldDesc,
							Required:    fieldRequired,
						})

						operationContent.WriteString(fieldName)
						operationContent.WriteString(" ")
						operationContent.WriteString(fieldDesc)
						operationContent.WriteString(" ")
					}
				}
			}
		}
	}

	return doc, nil
}
