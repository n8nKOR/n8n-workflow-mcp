package workflow

import (
	"fmt"
)

// NodeHelper는 노드 정보 관리를 담당합니다
type NodeHelper struct{}

// NewNodeHelper는 새로운 노드 헬퍼를 생성합니다
func NewNodeHelper() *NodeHelper {
	return &NodeHelper{}
}

// GetDisplayName은 노드 타입에 해당하는 표시 이름을 반환합니다
func (nh *NodeHelper) GetDisplayName(nodeType string) string {
	displayNames := map[string]string{
		"ManualTrigger":  "수동 트리거",
		"Schedule":       "스케줄 트리거",
		"Webhook":        "웹훅 트리거",
		"GoogleSheets":   "Google Sheets",
		"Slack":          "Slack",
		"Discord":        "Discord",
		"EmailSend":      "이메일 발송",
		"HttpRequest":    "HTTP 요청",
		"Code":           "코드 실행",
		"If":             "조건 분기",
		"Notion":         "Notion",
		"Airtable":       "Airtable",
		"MySql":          "MySQL",
		"Postgres":       "PostgreSQL",
		"Telegram":       "Telegram",
		"Jira":           "Jira",
		"Trello":         "Trello",
		"Asana":          "Asana",
		"Github":         "GitHub",
		"Gitlab":         "GitLab",
		"SplitInBatches": "배치 분할",
		"Files":          "파일 처리",
		"ReadPdf":        "PDF 읽기",
	}

	if displayName, exists := displayNames[nodeType]; exists {
		return displayName
	}
	return nodeType
}

// GetCategory는 노드 타입에 해당하는 카테고리를 반환합니다
func (nh *NodeHelper) GetCategory(nodeType string) string {
	categories := map[string]string{
		"ManualTrigger":  "trigger",
		"Schedule":       "trigger",
		"Webhook":        "trigger",
		"GoogleSheets":   "data",
		"Slack":          "communication",
		"Discord":        "communication",
		"EmailSend":      "communication",
		"HttpRequest":    "http",
		"Code":           "transform",
		"If":             "logic",
		"Notion":         "data",
		"Airtable":       "data",
		"MySql":          "database",
		"Postgres":       "database",
		"Telegram":       "communication",
		"Jira":           "productivity",
		"Trello":         "productivity",
		"Asana":          "productivity",
		"Github":         "developer",
		"Gitlab":         "developer",
		"SplitInBatches": "transform",
		"Files":          "file",
		"ReadPdf":        "file",
	}

	if category, exists := categories[nodeType]; exists {
		return category
	}
	return "general"
}

// GetDescription은 노드 타입과 사용자 요구사항을 기반으로 설명을 생성합니다
func (nh *NodeHelper) GetDescription(nodeType, userRequirement string) string {
	descriptions := map[string]string{
		"ManualTrigger":  "수동으로 워크플로우를 시작합니다",
		"Schedule":       "지정된 시간에 워크플로우를 자동 실행합니다",
		"Webhook":        "외부 시스템에서 호출되는 웹훅을 수신합니다",
		"GoogleSheets":   "Google Sheets에서 데이터를 읽거나 쓰는 작업을 수행합니다",
		"Slack":          "Slack 채널이나 사용자에게 메시지를 전송합니다",
		"Discord":        "Discord 서버에 메시지를 전송하거나 상호작용합니다",
		"EmailSend":      "이메일을 발송합니다",
		"HttpRequest":    "외부 API에 HTTP 요청을 보내고 응답을 받습니다",
		"Code":           "JavaScript 또는 Python 코드를 실행하여 데이터를 처리합니다",
		"If":             "조건에 따라 워크플로우를 다른 경로로 분기합니다",
		"Notion":         "Notion 데이터베이스에서 데이터를 읽거나 쓰는 작업을 수행합니다",
		"Airtable":       "Airtable 베이스에서 데이터를 관리합니다",
		"MySql":          "MySQL 데이터베이스에서 데이터를 조회하거나 수정합니다",
		"Postgres":       "PostgreSQL 데이터베이스에서 데이터를 조회하거나 수정합니다",
		"Telegram":       "Telegram 봇을 통해 메시지를 전송합니다",
		"Jira":           "Jira 이슈를 생성하거나 관리합니다",
		"Trello":         "Trello 보드에서 카드를 관리합니다",
		"Asana":          "Asana 프로젝트에서 작업을 관리합니다",
		"Github":         "GitHub 저장소와 상호작용합니다",
		"Gitlab":         "GitLab 저장소와 상호작용합니다",
		"SplitInBatches": "대용량 데이터를 작은 배치로 분할하여 처리합니다",
		"Files":          "파일을 읽거나 쓰는 작업을 수행합니다",
		"ReadPdf":        "PDF 파일에서 텍스트를 추출합니다",
	}

	if description, exists := descriptions[nodeType]; exists {
		return description
	}
	return fmt.Sprintf("%s 노드를 사용하여 작업을 수행합니다", nodeType)
}

// GetPurpose는 노드 타입의 목적을 반환합니다
func (nh *NodeHelper) GetPurpose(nodeType string) string {
	purposes := map[string]string{
		"ManualTrigger":  "사용자의 직접적인 개입으로 워크플로우를 시작",
		"Schedule":       "정기적인 자동화 실행",
		"Webhook":        "외부 시스템과의 실시간 연동",
		"GoogleSheets":   "스프레드시트 데이터 관리",
		"Slack":          "팀 커뮤니케이션 및 알림",
		"Discord":        "커뮤니티 알림 및 상호작용",
		"EmailSend":      "이메일 기반 알림 및 커뮤니케이션",
		"HttpRequest":    "외부 API와의 데이터 교환",
		"Code":           "커스텀 로직 및 데이터 변환",
		"If":             "조건부 로직 구현",
		"Notion":         "문서 및 데이터베이스 관리",
		"Airtable":       "구조화된 데이터 관리",
		"MySql":          "관계형 데이터베이스 연동",
		"Postgres":       "고급 관계형 데이터베이스 연동",
		"Telegram":       "개인 메시징 및 봇 서비스",
		"Jira":           "이슈 추적 및 프로젝트 관리",
		"Trello":         "칸반 스타일 작업 관리",
		"Asana":          "팀 프로젝트 관리",
		"Github":         "소스 코드 관리 및 협업",
		"Gitlab":         "DevOps 및 CI/CD 연동",
		"SplitInBatches": "대용량 데이터 처리 최적화",
		"Files":          "파일 시스템 연동",
		"ReadPdf":        "문서 내용 추출 및 분석",
	}

	if purpose, exists := purposes[nodeType]; exists {
		return purpose
	}
	return fmt.Sprintf("%s 관련 작업 수행", nodeType)
}

// EstimateInputData는 노드 타입에 따른 예상 입력 데이터를 반환합니다
func (nh *NodeHelper) EstimateInputData(nodeType string) string {
	inputData := map[string]string{
		"ManualTrigger":  "사용자 입력",
		"Schedule":       "스케줄 정보",
		"Webhook":        "웹훅 페이로드",
		"GoogleSheets":   "행 데이터 또는 스프레드시트 ID",
		"Slack":          "메시지 내용 및 채널 정보",
		"Discord":        "메시지 내용 및 채널 정보",
		"EmailSend":      "이메일 제목, 내용, 수신자",
		"HttpRequest":    "API 엔드포인트, 헤더, 파라미터",
		"Code":           "처리할 데이터",
		"If":             "조건 판단용 데이터",
		"Notion":         "페이지 또는 데이터베이스 항목",
		"Airtable":       "레코드 데이터",
		"MySql":          "SQL 쿼리 파라미터",
		"Postgres":       "SQL 쿼리 파라미터",
		"Telegram":       "메시지 내용 및 채팅 ID",
		"Jira":           "이슈 정보",
		"Trello":         "카드 정보",
		"Asana":          "작업 정보",
		"Github":         "저장소 정보 또는 이슈 데이터",
		"Gitlab":         "프로젝트 정보 또는 머지 리퀘스트",
		"SplitInBatches": "전체 데이터 배열",
		"Files":          "파일 경로 또는 내용",
		"ReadPdf":        "PDF 파일 경로",
	}

	if data, exists := inputData[nodeType]; exists {
		return data
	}
	return "입력 데이터"
}

// EstimateOutputData는 노드 타입에 따른 예상 출력 데이터를 반환합니다
func (nh *NodeHelper) EstimateOutputData(nodeType string) string {
	outputData := map[string]string{
		"ManualTrigger":  "트리거 정보",
		"Schedule":       "실행 시간 정보",
		"Webhook":        "수신된 데이터",
		"GoogleSheets":   "처리 결과 또는 읽은 데이터",
		"Slack":          "메시지 전송 결과",
		"Discord":        "메시지 전송 결과",
		"EmailSend":      "이메일 발송 결과",
		"HttpRequest":    "API 응답 데이터",
		"Code":           "처리된 데이터",
		"If":             "조건 분기 결과",
		"Notion":         "처리 결과 또는 읽은 데이터",
		"Airtable":       "처리 결과 또는 읽은 데이터",
		"MySql":          "쿼리 실행 결과",
		"Postgres":       "쿼리 실행 결과",
		"Telegram":       "메시지 전송 결과",
		"Jira":           "이슈 처리 결과",
		"Trello":         "카드 처리 결과",
		"Asana":          "작업 처리 결과",
		"Github":         "저장소 작업 결과",
		"Gitlab":         "프로젝트 작업 결과",
		"SplitInBatches": "배치 데이터",
		"Files":          "파일 처리 결과",
		"ReadPdf":        "추출된 텍스트",
	}

	if data, exists := outputData[nodeType]; exists {
		return data
	}
	return "출력 데이터"
}
