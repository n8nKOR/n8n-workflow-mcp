package i18n

// getKoreanMessages는 한국어 메시지를 반환합니다
func (i *I18n) getKoreanMessages() map[string]string {
	return map[string]string{
		// 기본 도구 설명
		"tool.ping.description":             "서버가 정상 작동하는지 확인하는 기본 ping 테스트",
		"tool.search_n8n_nodes.description": "n8n 노드를 검색합니다. 키워드에 기반하여 관련 노드들을 찾고 점수와 함께 반환합니다.",

		// 워크플로우 관련 도구

		"tool.validate_workflow.description": `⚠️ **워크플로우 검증 도구**: n8n 워크플로우 JSON을 종합적으로 검증합니다.

**검증 항목:**
- JSON 구문 정확성 및 스키마 유효성
- 노드 구조 및 필수 속성 누락 검사
- 노드 간 연결 상태 및 데이터 흐름 유효성
- 중복 노드 ID 및 순환 참조 검사
- n8n 표준 규격 준수 여부

**중요:** 워크플로우 JSON 생성 또는 수정 후 반드시 이 도구로 검증하세요. 검증 없이 n8n에 배포하면 실행 시 오류가 발생할 수 있습니다.

**검증 실패 시:** 상세한 오류 정보와 함께 줄 번호별 수정 가이드를 제공합니다.`,

		// 파라미터 설명
		"param.query.description": "검색할 키워드 또는 구문 (영어로 검색해야 정확한 결과를 얻을 수 있습니다)",
		"param.limit.description": "반환할 최대 결과 수 (기본: 5, 최대: 20)",

		"param.file_path.description":            "검증할 n8n 워크플로우 JSON 파일의 절대/상대 경로 (예: './workflows/my_workflow.json')",
		"param.include_line_numbers.description": "줄 번호별 상세 오류 정보 및 컨텍스트 포함 여부 (기본값: true)",
		"param.auto_fix_suggestions.description": "자동 수정 제안사항 제공 여부 (기본값: true)",
	}
}

// getEnglishMessages는 영어 메시지를 반환합니다
func (i *I18n) getEnglishMessages() map[string]string {
	return map[string]string{
		// 기본 도구 설명
		"tool.ping.description":             "Basic ping test to check if the server is working properly",
		"tool.search_n8n_nodes.description": "Search n8n nodes. Find related nodes based on keywords and return them with scores.",

		// 워크플로우 관련 도구

		"tool.validate_workflow.description": `⚠️ **Workflow Validation Tool**: Comprehensively validate n8n workflow JSON.

**Validation Items:**
- JSON syntax accuracy and schema validity
- Node structure and missing required attributes check
- Node connection status and data flow validity
- Duplicate node ID and circular reference check
- n8n standard compliance

**Important:** Always validate with this tool after creating or modifying workflow JSON. Deploying to n8n without validation may cause runtime errors.

**When Validation Fails:** Provides detailed error information with line-by-line correction guides.`,

		// 파라미터 설명
		"param.query.description": "Keywords or phrases to search for (must be in English for accurate results)",
		"param.limit.description": "Maximum number of results to return (default: 5, max: 20)",

		"param.file_path.description":            "Absolute/relative path of n8n workflow JSON file to validate (e.g., './workflows/my_workflow.json')",
		"param.include_line_numbers.description": "Whether to include detailed error information and context by line number (default: true)",
		"param.auto_fix_suggestions.description": "Whether to provide automatic correction suggestions (default: true)",
	}
}

// getJapaneseMessages는 일본어 메시지를 반환합니다
func (i *I18n) getJapaneseMessages() map[string]string {
	return map[string]string{
		// 기본 도구 설명
		"tool.ping.description":             "サーバーが正常に動作しているかを確認する基本的なpingテスト",
		"tool.search_n8n_nodes.description": "n8nノードを検索します。キーワードに基づいて関連ノードを見つけ、スコアと共に返します。",

		// 워크플로우 관련 도구

		"tool.validate_workflow.description": `⚠️ **ワークフロー検証ツール**: n8nワークフローJSONを包括的に検証します。

**検証項目:**
- JSON構文の正確性とスキーマの有効性
- ノード構造と必須属性の欠落チェック
- ノード間接続状態とデータフローの有効性
- 重複ノードIDと循環参照チェック
- n8n標準規格への準拠

**重要:** ワークフローJSONの作成または修正後は必ずこのツールで検証してください。検証なしでn8nにデプロイすると実行時エラーが発生する可能性があります。

**検証失敗時:** 詳細なエラー情報と行番号別修正ガイドを提供します。`,

		// 파라미터 설명
		"param.query.description": "検索するキーワードまたはフレーズ（正確な結果を得るには英語で検索する必要があります）",
		"param.limit.description": "返す結果の最大数（デフォルト: 5、最大: 20）",

		"param.file_path.description":            "検証するn8nワークフローJSONファイルの絶対/相対パス（例：'./workflows/my_workflow.json'）",
		"param.include_line_numbers.description": "行番号別詳細エラー情報とコンテキストを含めるかどうか（デフォルト: true）",
		"param.auto_fix_suggestions.description": "自動修正提案を提供するかどうか（デフォルト: true）",
	}
}
