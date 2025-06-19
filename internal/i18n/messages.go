package i18n

// Messages는 모든 언어의 메시지를 포함합니다
var Messages = map[string]map[string]string{
	"ko": {
		// 도구 설명
		"tool.ping.description":            "서버 연결 상태를 확인합니다",
		"tool.search.description":          "n8n 노드 문서를 검색합니다",
		"tool.search_workflow.description": "n8n 워크플로우 템플릿을 검색합니다. 키워드나 카테고리에 기반하여 커뮤니티 워크플로우를 찾아 URL과 함께 반환합니다. (영어로 검색해야 정확한 결과를 얻을 수 있습니다)\n\n📝 **검색 후 필수 작업**: 모든 검색 결과를 마크다운 파일(.md)로 저장해야 합니다. URL은 반드시 클릭 가능한 링크 형태 `[워크플로우 제목](URL)`로 변환하여 브라우저에서 바로 접근할 수 있도록 구성하세요.",

		// 매개변수 설명
		"tool.search.param.query.description":                "검색할 키워드 또는 구문 (예: HTTP, webhook, database)",
		"tool.search.param.limit.description":                "반환할 최대 결과 수 (기본: 5, 최대: 20)",
		"tool.search_workflow.param.query.description":       "검색할 키워드 또는 구문 (예: AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "반환할 최대 결과 수 (기본: 10, 최대: 50)",

		// 오류 메시지
		"error.query_required":   "query 파라미터가 필요합니다",
		"error.search_failed":    "검색 실행 중 오류 발생: %v",
		"error.marshal_failed":   "결과 변환 중 오류 발생: %v",
		"error.no_results":       "검색 결과가 없습니다",
		"error.service_not_init": "검색 서비스가 초기화되지 않았습니다",

		// 성공 메시지
		"success.ping": "pong",
	},
	"en": {
		// Tool descriptions
		"tool.ping.description":            "Check server connection status",
		"tool.search.description":          "Search n8n node documentation",
		"tool.search_workflow.description": "Search n8n workflow templates. Find community workflows based on keywords or categories and return them with URLs. (Search must be performed in English for accurate results)\n\n📝 **Required Post-Search Action**: All search results must be saved as a markdown file (.md). URLs must be converted to clickable link format `[Workflow Title](URL)` to enable direct browser access.",

		// Parameter descriptions
		"tool.search.param.query.description":                "Keywords or phrases to search for (e.g., HTTP, webhook, database)",
		"tool.search.param.limit.description":                "Maximum number of results to return (default: 5, max: 20)",
		"tool.search_workflow.param.query.description":       "Keywords or phrases to search for (e.g., AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "Maximum number of results to return (default: 10, max: 50)",

		// Error messages
		"error.query_required":   "query parameter is required",
		"error.search_failed":    "Error occurred during search execution: %v",
		"error.marshal_failed":   "Error occurred during result conversion: %v",
		"error.no_results":       "No search results found",
		"error.service_not_init": "Search service is not initialized",

		// Success messages
		"success.ping": "pong",
	},
	"ja": {
		// ツールの説明
		"tool.ping.description":            "サーバー接続状態を確認します",
		"tool.search.description":          "n8nノードドキュメントを検索します",
		"tool.search_workflow.description": "n8nワークフローテンプレートを検索します。キーワードやカテゴリに基づいてコミュニティワークフローを見つけ、URLと共に返します。（正確な結果を得るには英語で検索する必要があります）\n\n📝 **検索後必須作業**: すべての検索結果をマークダウンファイル（.md）として保存する必要があります。URLは必ずクリック可能なリンク形式 `[ワークフロータイトル](URL)` に変換して、ブラウザから直接アクセスできるように構成してください。",

		// パラメータの説明
		"tool.search.param.query.description":                "検索するキーワードまたはフレーズ（例：HTTP、webhook、database）",
		"tool.search.param.limit.description":                "返す最大結果数（デフォルト：5、最大：20）",
		"tool.search_workflow.param.query.description":       "検索するキーワードまたはフレーズ（例：AI、automation、chatbot）",
		"tool.search_workflow.param.max_results.description": "返す最大結果数（デフォルト：10、最大：50）",

		// エラーメッセージ
		"error.query_required":   "queryパラメータが必要です",
		"error.search_failed":    "検索実行中にエラーが発生しました：%v",
		"error.marshal_failed":   "結果変換中にエラーが発生しました：%v",
		"error.no_results":       "検索結果が見つかりません",
		"error.service_not_init": "検索サービスが初期化されていません",

		// 成功メッセージ
		"success.ping": "pong",
	},
}

// getKoreanMessages는 한국어 메시지를 반환합니다
func (i *I18n) getKoreanMessages() map[string]string {
	return map[string]string{
		// 기본 도구 설명
		"tool.ping.description":             "서버가 정상 작동하는지 확인하는 기본 ping 테스트",
		"tool.search_n8n_nodes.description": "n8n 노드를 검색합니다. 키워드에 기반하여 관련 노드들을 찾고 점수와 함께 반환합니다.",
		"tool.search_workflow.description":  "n8n 워크플로우 템플릿을 검색합니다. 키워드나 카테고리에 기반하여 커뮤니티 워크플로우를 찾아 URL과 함께 반환합니다. (영어로 검색해야 정확한 결과를 얻을 수 있습니다)\n\n📝 **검색 후 필수 작업**: 모든 검색 결과를 마크다운 파일(.md)로 저장해야 합니다. URL은 반드시 클릭 가능한 링크 형태 `[워크플로우 제목](URL)`로 변환하여 브라우저에서 바로 접근할 수 있도록 구성하세요.",

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

		// 워크플로우 검색 파라미터
		"tool.search_workflow.param.query.description":       "검색할 키워드 또는 구문 (예: AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "반환할 최대 결과 수 (기본: 5, 최대: 15)",
		"tool.search_workflow.param.category.description":    "워크플로우 카테고리로 필터링 (예: airtable, blog, common, crm, email)",

		// 다음 단계 가이드
		"search_workflow.next_step": "검색된 워크플로우들을 분석하여 각각의 기능과 용도를 설명하고, **모든 URL을 클릭 가능한 마크다운 링크 형태**로 변환하여 문서를 생성한 후 `.md` 파일로 저장해주세요. 사용자가 링크를 클릭하면 바로 브라우저에서 워크플로우를 확인할 수 있도록 하고, 각 워크플로우의 제목, 설명, 활용 방법을 포함하여 구성해주세요.",

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
		"tool.search_workflow.description":  "Search n8n workflow templates. Find community workflows based on keywords or categories and return them with URLs. (Search must be performed in English for accurate results)\n\n📝 **Required Post-Search Action**: All search results must be saved as a markdown file (.md). URLs must be converted to clickable link format `[Workflow Title](URL)` to enable direct browser access.",

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

		// 워크플로우 검색 파라미터
		"tool.search_workflow.param.query.description":       "Keywords or phrases to search for (e.g., AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "Maximum number of results to return (default: 5, max: 15)",
		"tool.search_workflow.param.category.description":    "Filter by workflow category (e.g., airtable, blog, common, crm, email)",

		// Next step guide
		"search_workflow.next_step": "Analyze the searched workflows to explain each one's functionality and purpose, then **convert all URLs into clickable markdown link format** and generate a document saved as a `.md` file. Ensure users can click the links to directly open workflows in their browser, and include each workflow's title, description, and usage instructions.",

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
		"tool.search_workflow.description":  "n8nワークフローテンプレートを検索します。キーワードやカテゴリに基づいてコミュニティワークフローを見つけ、URLと共に返します。（正確な結果を得るには英語で検索する必要があります）\n\n📝 **検索後必須作業**: すべての検索結果をマークダウンファイル（.md）として保存する必要があります。URLは必ずクリック可能なリンク形式 `[ワークフロータイトル](URL)` に変換して、ブラウザから直接アクセスできるように構成してください。",

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

		// 워크플로우 검색 파라미터
		"tool.search_workflow.param.query.description":       "検索するキーワードまたはフレーズ（例：AI、automation、chatbot）",
		"tool.search_workflow.param.max_results.description": "返す結果の最大数（デフォルト: 5、最大: 15）",
		"tool.search_workflow.param.category.description":    "ワークフローカテゴリでフィルタリング（例：airtable、blog、common、crm、email）",

		// 次のステップガイド
		"search_workflow.next_step": "検索されたワークフローを分析して、それぞれの機能と用途を説明し、**すべてのURLをクリック可能なマークダウンリンク形式**に変換して文書を生成し`.md`ファイルとして保存してください。ユーザーがリンクをクリックするとブラウザで直接ワークフローを確認できるようにし、各ワークフローのタイトル、説明、活用方法を含めて構成してください。",

		"param.file_path.description":            "検証するn8nワークフローJSONファイルの絶対/相対パス（例：'./workflows/my_workflow.json'）",
		"param.include_line_numbers.description": "行番号別詳細エラー情報とコンテキストを含めるかどうか（デフォルト: true）",
		"param.auto_fix_suggestions.description": "自動修正提案を提供するかどうか（デフォルト: true）",
	}
}
