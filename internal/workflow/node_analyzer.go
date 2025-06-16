package workflow

import (
	"strings"
)

// NodeAnalyzer는 사용자 요구사항을 분석하여 필요한 노드들을 추천합니다
type NodeAnalyzer struct{}

// NewNodeAnalyzer는 새로운 노드 분석기를 생성합니다
func NewNodeAnalyzer() *NodeAnalyzer {
	return &NodeAnalyzer{}
}

// AnalyzeProgramAndSuggestNodes는 사용자 요구사항을 분석하여 필요한 노드들을 추천합니다
func (na *NodeAnalyzer) AnalyzeProgramAndSuggestNodes(userRequirement string) []string {
	requirement := strings.ToLower(userRequirement)
	var suggestedNodes []string

	// 트리거 노드 분석
	if strings.Contains(requirement, "매일") || strings.Contains(requirement, "정기적") ||
		strings.Contains(requirement, "스케줄") || strings.Contains(requirement, "시간") {
		suggestedNodes = append(suggestedNodes, "Schedule")
	}
	if strings.Contains(requirement, "웹훅") || strings.Contains(requirement, "webhook") ||
		strings.Contains(requirement, "api 호출") {
		suggestedNodes = append(suggestedNodes, "Webhook")
	}
	if strings.Contains(requirement, "수동") || strings.Contains(requirement, "버튼") {
		suggestedNodes = append(suggestedNodes, "ManualTrigger")
	}

	// 데이터 소스 노드 분석
	if strings.Contains(requirement, "google sheets") || strings.Contains(requirement, "구글 시트") ||
		strings.Contains(requirement, "스프레드시트") {
		suggestedNodes = append(suggestedNodes, "GoogleSheets")
	}
	if strings.Contains(requirement, "notion") || strings.Contains(requirement, "노션") {
		suggestedNodes = append(suggestedNodes, "Notion")
	}
	if strings.Contains(requirement, "airtable") || strings.Contains(requirement, "에어테이블") {
		suggestedNodes = append(suggestedNodes, "Airtable")
	}
	if strings.Contains(requirement, "mysql") || strings.Contains(requirement, "postgres") ||
		strings.Contains(requirement, "데이터베이스") || strings.Contains(requirement, "db") {
		if strings.Contains(requirement, "mysql") {
			suggestedNodes = append(suggestedNodes, "MySql")
		} else {
			suggestedNodes = append(suggestedNodes, "Postgres")
		}
	}

	// HTTP 요청 관련
	if strings.Contains(requirement, "api") || strings.Contains(requirement, "http") ||
		strings.Contains(requirement, "rest") || strings.Contains(requirement, "데이터 가져오기") ||
		strings.Contains(requirement, "외부") {
		suggestedNodes = append(suggestedNodes, "HttpRequest")
	}

	// 메시징/알림 노드 분석
	if strings.Contains(requirement, "slack") || strings.Contains(requirement, "슬랙") {
		suggestedNodes = append(suggestedNodes, "Slack")
	}
	if strings.Contains(requirement, "discord") || strings.Contains(requirement, "디스코드") {
		suggestedNodes = append(suggestedNodes, "Discord")
	}
	if strings.Contains(requirement, "email") || strings.Contains(requirement, "이메일") ||
		strings.Contains(requirement, "메일") {
		suggestedNodes = append(suggestedNodes, "EmailSend")
	}
	if strings.Contains(requirement, "telegram") || strings.Contains(requirement, "텔레그램") {
		suggestedNodes = append(suggestedNodes, "Telegram")
	}

	// 프로젝트 관리 도구
	if strings.Contains(requirement, "jira") || strings.Contains(requirement, "지라") {
		suggestedNodes = append(suggestedNodes, "Jira")
	}
	if strings.Contains(requirement, "trello") || strings.Contains(requirement, "트렐로") {
		suggestedNodes = append(suggestedNodes, "Trello")
	}
	if strings.Contains(requirement, "asana") || strings.Contains(requirement, "아사나") {
		suggestedNodes = append(suggestedNodes, "Asana")
	}

	// 개발 도구
	if strings.Contains(requirement, "github") || strings.Contains(requirement, "깃허브") ||
		strings.Contains(requirement, "git") {
		suggestedNodes = append(suggestedNodes, "Github")
	}
	if strings.Contains(requirement, "gitlab") || strings.Contains(requirement, "깃랩") {
		suggestedNodes = append(suggestedNodes, "Gitlab")
	}

	// 데이터 처리 노드
	if strings.Contains(requirement, "변환") || strings.Contains(requirement, "처리") ||
		strings.Contains(requirement, "가공") || strings.Contains(requirement, "계산") ||
		strings.Contains(requirement, "분석") {
		suggestedNodes = append(suggestedNodes, "Code")
	}
	if strings.Contains(requirement, "조건") || strings.Contains(requirement, "분기") ||
		strings.Contains(requirement, "if") || strings.Contains(requirement, "만약") {
		suggestedNodes = append(suggestedNodes, "If")
	}
	if strings.Contains(requirement, "분할") || strings.Contains(requirement, "배치") {
		suggestedNodes = append(suggestedNodes, "SplitInBatches")
	}

	// 파일 처리
	if strings.Contains(requirement, "파일") || strings.Contains(requirement, "업로드") ||
		strings.Contains(requirement, "다운로드") {
		suggestedNodes = append(suggestedNodes, "Files")
	}
	if strings.Contains(requirement, "pdf") {
		suggestedNodes = append(suggestedNodes, "ReadPdf")
	}

	// 기본 노드들 (항상 필요)
	if len(suggestedNodes) == 0 {
		suggestedNodes = append(suggestedNodes, "ManualTrigger", "HttpRequest", "Code")
	}

	return na.removeDuplicateNodes(suggestedNodes)
}

// removeDuplicateNodes는 중복된 노드를 제거합니다
func (na *NodeAnalyzer) removeDuplicateNodes(nodes []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, node := range nodes {
		if !keys[node] {
			keys[node] = true
			result = append(result, node)
		}
	}

	return result
}
