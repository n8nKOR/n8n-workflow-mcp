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
	"vi": {
		// Mô tả công cụ
		"tool.ping.description":            "Kiểm tra trạng thái kết nối máy chủ",
		"tool.search.description":          "Tìm kiếm tài liệu node n8n",
		"tool.search_workflow.description": "Tìm kiếm mẫu workflow n8n. Tìm các workflow cộng đồng dựa trên từ khóa hoặc danh mục và trả về kèm URL. (Cần tìm kiếm bằng tiếng Anh để có kết quả chính xác)\n\n📝 **Hành động bắt buộc sau tìm kiếm**: Tất cả kết quả tìm kiếm phải được lưu dưới dạng file markdown (.md). URL phải được chuyển đổi thành định dạng liên kết có thể nhấp `[Tiêu đề Workflow](URL)` để có thể truy cập trực tiếp từ trình duyệt.",

		// Mô tả tham số
		"tool.search.param.query.description":                "Từ khóa hoặc cụm từ cần tìm kiếm (ví dụ: HTTP, webhook, database)",
		"tool.search.param.limit.description":                "Số lượng kết quả tối đa trả về (mặc định: 5, tối đa: 20)",
		"tool.search_workflow.param.query.description":       "Từ khóa hoặc cụm từ cần tìm kiếm (ví dụ: AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "Số lượng kết quả tối đa trả về (mặc định: 10, tối đa: 50)",

		// Thông báo lỗi
		"error.query_required":   "Tham số query là bắt buộc",
		"error.search_failed":    "Lỗi xảy ra trong quá trình tìm kiếm: %v",
		"error.marshal_failed":   "Lỗi xảy ra trong quá trình chuyển đổi kết quả: %v",
		"error.no_results":       "Không tìm thấy kết quả",
		"error.service_not_init": "Dịch vụ tìm kiếm chưa được khởi tạo",

		// Thông báo thành công
		"success.ping": "pong",
	},
	"th": {
		// คำอธิบายเครื่องมือพื้นฐาน
		"tool.ping.description":            "การทดสอบ ping พื้นฐานเพื่อตรวจสอบว่าเซิร์ฟเวอร์ทำงานอย่างถูกต้อง",
		"tool.search.description":          "ค้นหาเอกสาร node n8n",
		"tool.search_workflow.description": "ค้นหาเทมเพลต workflow n8n ค้นหา workflow ชุมชนตามคำหลักหรือหมวดหมู่และส่งคืนพร้อม URL (ต้องค้นหาด้วยภาษาอังกฤษเพื่อให้ได้ผลลัพธ์ที่แม่นยำ)\n\n📝 **การดำเนินการที่จำเป็นหลังการค้นหา**: ผลลัพธ์การค้นหาทั้งหมดต้องบันทึกเป็นไฟล์ markdown (.md) URL ต้องถูกแปลงเป็นรูปแบบลิงก์ที่คลิกได้ `[ชื่อ Workflow](URL)` เพื่อให้สามารถเข้าถึงได้โดยตรงจากเบราว์เซอร์",

		// เครื่องมือที่เกี่ยวข้องกับ workflow
		"tool.validate_workflow.description": `⚠️ **เครื่องมือตรวจสอบ Workflow**: ตรวจสอบ JSON workflow n8n อย่างครอบคลุม

**รายการตรวจสอบ:**
- ความถูกต้องของไวยากรณ์ JSON และความถูกต้องของ schema
- ตรวจสอบโครงสร้าง node และคุณลักษณะที่จำเป็นที่ขาดหายไป
- ความถูกต้องของสถานะการเชื่อมต่อ node และการไหลของข้อมูล
- ตรวจสอบ ID node ที่ซ้ำกันและการอ้างอิงแบบวงกลม
- การปฏิบัติตามมาตรฐาน n8n

**สำคัญ:** ควรตรวจสอบด้วยเครื่องมือนี้เสมอหลังจากสร้างหรือแก้ไข JSON workflow การปรับใช้กับ n8n โดยไม่ตรวจสอบอาจทำให้เกิดข้อผิดพลาดขณะรันไทม์

**เมื่อการตรวจสอบล้มเหลว:** ให้ข้อมูลข้อผิดพลาดโดยละเอียดพร้อมคำแนะนำการแก้ไขทีละบรรทัด`,

		// คำอธิบายพารามิเตอร์
		"param.query.description": "คำหลักหรือวลีที่ต้องการค้นหา (ต้องเป็นภาษาอังกฤษเพื่อให้ได้ผลลัพธ์ที่แม่นยำ)",
		"param.limit.description": "จำนวนผลลัพธ์สูงสุดที่จะส่งคืน (ค่าเริ่มต้น: 5, สูงสุด: 20)",

		// พารามิเตอร์การค้นหา workflow
		"tool.search_workflow.param.query.description":       "คำหลักหรือวลีที่ต้องการค้นหา (เช่น: AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "จำนวนผลลัพธ์สูงสุดที่จะส่งคืน (ค่าเริ่มต้น: 5, สูงสุด: 15)",
		"tool.search_workflow.param.category.description":    "กรองตามหมวดหมู่ workflow (เช่น: airtable, blog, common, crm, email)",

		// คำแนะนำขั้นตอนถัดไป
		"search_workflow.next_step": "วิเคราะห์ workflow ที่ค้นพบเพื่ออธิบายฟังก์ชันและวัตถุประสงค์ของแต่ละอัน จากนั้น **แปลง URL ทั้งหมดเป็นรูปแบบลิงก์ markdown ที่คลิกได้** และสร้างเอกสารบันทึกเป็นไฟล์ `.md` ตรวจสอบให้แน่ใจว่าผู้ใช้สามารถคลิกลิงก์เพื่อเปิด workflow โดยตรงในเบราว์เซอร์ รวมทั้งชื่อเรื่อง คำอธิบาย และคำแนะนำการใช้งานของแต่ละ workflow",

		// คำอธิบายพารามิเตอร์เพิ่มเติม
		"tool.search.param.query.description": "คำหลักหรือวลีที่ต้องการค้นหา (เช่น: HTTP, webhook, database)",
		"tool.search.param.limit.description": "จำนวนผลลัพธ์สูงสุดที่จะส่งคืน (ค่าเริ่มต้น: 5, สูงสุด: 20)",

		// ข้อความแสดงข้อผิดพลาด
		"error.query_required":   "จำเป็นต้องระบุพารามิเตอร์ query",
		"error.search_failed":    "เกิดข้อผิดพลาดระหว่างการค้นหา: %v",
		"error.marshal_failed":   "เกิดข้อผิดพลาดระหว่างการแปลงผลลัพธ์: %v",
		"error.no_results":       "ไม่พบผลลัพธ์การค้นหา",
		"error.service_not_init": "บริการค้นหายังไม่ได้เริ่มต้น",

		// ข้อความแสดงความสำเร็จ
		"success.ping": "pong",

		"param.file_path.description":            "เส้นทางสัมบูรณ์/สัมพันธ์ของไฟล์ JSON workflow n8n ที่จะตรวจสอบ (เช่น: './workflows/my_workflow.json')",
		"param.include_line_numbers.description": "รวมข้อมูลข้อผิดพลาดโดยละเอียดและบริบทตามหมายเลขบรรทัดหรือไม่ (ค่าเริ่มต้น: true)",
		"param.auto_fix_suggestions.description": "ให้คำแนะนำการแก้ไขอัตโนมัติหรือไม่ (ค่าเริ่มต้น: true)",
	},
	"tw": {
		// 工具說明
		"tool.ping.description":            "檢查伺服器連線狀態",
		"tool.search.description":          "搜尋 n8n 節點文件",
		"tool.search_workflow.description": "搜尋 n8n 工作流程範本。根據關鍵字或類別尋找社群工作流程並連同 URL 一起回傳。(必須使用英文搜尋才能取得準確結果)\n\n📝 **搜尋後必要動作**: 所有搜尋結果必須儲存為 markdown 檔案 (.md)。URL 必須轉換為可點擊的連結格式 `[工作流程標題](URL)` 以便能從瀏覽器直接存取。",

		// 參數說明
		"tool.search.param.query.description":                "要搜尋的關鍵字或詞組 (例如: HTTP, webhook, database)",
		"tool.search.param.limit.description":                "回傳結果的最大數量 (預設: 5, 最大: 20)",
		"tool.search_workflow.param.query.description":       "要搜尋的關鍵字或詞組 (例如: AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "回傳結果的最大數量 (預設: 10, 最大: 50)",

		// 錯誤訊息
		"error.query_required":   "query 參數為必要項目",
		"error.search_failed":    "搜尋執行過程中發生錯誤: %v",
		"error.marshal_failed":   "結果轉換過程中發生錯誤: %v",
		"error.no_results":       "找不到搜尋結果",
		"error.service_not_init": "搜尋服務尚未初始化",

		// 成功訊息
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
	}
}

// getVietnameseMessages는 베트남어 메시지를 반환합니다
func (i *I18n) getVietnameseMessages() map[string]string {
	return map[string]string{
		// Mô tả công cụ cơ bản
		"tool.ping.description":             "Kiểm tra ping cơ bản để xác minh máy chủ hoạt động bình thường",
		"tool.search_n8n_nodes.description": "Tìm kiếm node n8n. Tìm các node liên quan dựa trên từ khóa và trả với điểm số.",
		"tool.search_workflow.description":  "Tìm kiếm mẫu workflow n8n. Tìm workflow cộng đồng dựa trên từ khóa hoặc danh mục và trả về cùng với URL. (Cần tìm kiếm bằng tiếng Anh để có kết quả chính xác)\n\n📝 **Hành động bắt buộc sau tìm kiếm**: Tất cả kết quả tìm kiếm phải được lưu dưới dạng file markdown (.md). URL phải được chuyển đổi thành định dạng liên kết có thể nhấp `[Tiêu đề Workflow](URL)` để có thể truy cập trực tiếp từ trình duyệt.",

		// Công cụ liên quan đến workflow
		"tool.validate_workflow.description": `⚠️ **Công cụ Xác thực Workflow**: Xác thực toàn diện JSON workflow n8n.

**Các mục xác thực:**
- Tính chính xác cú pháp JSON và tính hợp lệ của schema
- Kiểm tra cấu trúc node và thuộc tính bắt buộc còn thiếu
- Tính hợp lệ của trạng thái kết nối node và luồng dữ liệu
- Kiểm tra ID node trùng lặp và tham chiếu vòng tròn
- Tuân thủ tiêu chuẩn n8n

**Quan trọng:** Luôn xác thực bằng công cụ này sau khi tạo hoặc sửa đổi JSON workflow. Triển khai lên n8n mà không xác thực có thể gây lỗi thời gian chạy.

**Khi xác thực thất bại:** Cung cấp thông tin lỗi chi tiết cùng với hướng dẫn sửa chữa theo từng dòng.`,

		// Mô tả tham số
		"param.query.description": "Từ khóa hoặc cụm từ cần tìm kiếm (phải bằng tiếng Anh để có kết quả chính xác)",
		"param.limit.description": "Số lượng kết quả tối đa trả về (mặc định: 5, tối đa: 20)",

		// Tham số tìm kiếm workflow
		"tool.search_workflow.param.query.description":       "Từ khóa hoặc cụm từ cần tìm kiếm (ví dụ: AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "Số lượng kết quả tối đa trả về (mặc định: 5, tối đa: 15)",
		"tool.search_workflow.param.category.description":    "Lọc theo danh mục workflow (ví dụ: airtable, blog, common, crm, email)",

		// Hướng dẫn bước tiếp theo
		"search_workflow.next_step": "Phân tích các workflow đã tìm kiếm để giải thích chức năng và mục đích của từng workflow, sau đó **chuyển đổi tất cả URL thành định dạng liên kết markdown có thể nhấp** và tạo tài liệu lưu dưới dạng file `.md`. Đảm bảo người dùng có thể nhấp vào liên kết để mở workflow trực tiếp trong trình duyệt, bao gồm tiêu đề, mô tả và hướng dẫn sử dụng của từng workflow.",

		"param.file_path.description":            "Đường dẫn tuyệt đối/tương đối của file JSON workflow n8n cần xác thực (ví dụ: './workflows/my_workflow.json')",
		"param.include_line_numbers.description": "Có bao gồm thông tin lỗi chi tiết và ngữ cảnh theo số dòng hay không (mặc định: true)",
		"param.auto_fix_suggestions.description": "Có cung cấp đề xuất sửa chữa tự động hay không (mặc định: true)",
	}
}

// getThaiMessages는 태국어 메시지를 반환합니다
func (i *I18n) getThaiMessages() map[string]string {
	return map[string]string{
		// คำอธิบายเครื่องมือพื้นฐาน
		"tool.ping.description":             "การทดสอบ ping พื้นฐานเพื่อตรวจสอบว่าเซิร์ฟเวอร์ทำงานอย่างถูกต้อง",
		"tool.search_n8n_nodes.description": "ค้นหา node n8n ค้นหา node ที่เกี่ยวข้องตามคำหลักและส่งคืนพร้อมคะแนน",
		"tool.search_workflow.description":  "ค้นหาเทมเพลต workflow n8n ค้นหา workflow ชุมชนตามคำหลักหรือหมวดหมู่และส่งคืนพร้อม URL (ต้องค้นหาด้วยภาษาอังกฤษเพื่อให้ได้ผลลัพธ์ที่แม่นยำ)\n\n📝 **การดำเนินการที่จำเป็นหลังการค้นหา**: ผลลัพธ์การค้นหาทั้งหมดต้องบันทึกเป็นไฟล์ markdown (.md) URL ต้องถูกแปลงเป็นรูปแบบลิงก์ที่คลิกได้ `[ชื่อ Workflow](URL)` เพื่อให้สามารถเข้าถึงได้โดยตรงจากเบราว์เซอร์",

		// เครื่องมือที่เกี่ยวข้องกับ workflow
		"tool.validate_workflow.description": `⚠️ **เครื่องมือตรวจสอบ Workflow**: ตรวจสอบ JSON workflow n8n อย่างครอบคลุม

**รายการตรวจสอบ:**
- ความถูกต้องของไวยากรณ์ JSON และความถูกต้องของ schema
- ตรวจสอบโครงสร้าง node และคุณลักษณะที่จำเป็นที่ขาดหายไป
- ความถูกต้องของสถานะการเชื่อมต่อ node และการไหลของข้อมูล
- ตรวจสอบ ID node ที่ซ้ำกันและการอ้างอิงแบบวงกลม
- การปฏิบัติตามมาตรฐาน n8n

**สำคัญ:** ควรตรวจสอบด้วยเครื่องมือนี้เสมอหลังจากสร้างหรือแก้ไข JSON workflow การปรับใช้กับ n8n โดยไม่ตรวจสอบอาจทำให้เกิดข้อผิดพลาดขณะรันไทม์

**เมื่อการตรวจสอบล้มเหลว:** ให้ข้อมูลข้อผิดพลาดโดยละเอียดพร้อมคำแนะนำการแก้ไขทีละบรรทัด`,

		// คำอธิบายพารามิเตอร์
		"param.query.description": "คำหลักหรือวลีที่ต้องการค้นหา (ต้องเป็นภาษาอังกฤษเพื่อให้ได้ผลลัพธ์ที่แม่นยำ)",
		"param.limit.description": "จำนวนผลลัพธ์สูงสุดที่จะส่งคืน (ค่าเริ่มต้น: 5, สูงสุด: 20)",

		// พารามิเตอร์การค้นหา workflow
		"tool.search_workflow.param.query.description":       "คำหลักหรือวลีที่ต้องการค้นหา (เช่น: AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "จำนวนผลลัพธ์สูงสุดที่จะส่งคืน (ค่าเริ่มต้น: 5, สูงสุด: 15)",
		"tool.search_workflow.param.category.description":    "กรองตามหมวดหมู่ workflow (เช่น: airtable, blog, common, crm, email)",

		// คำแนะนำขั้นตอนถัดไป
		"search_workflow.next_step": "วิเคราะห์ workflow ที่ค้นพบเพื่ออธิบายฟังก์ชันและวัตถุประสงค์ของแต่ละอัน จากนั้น **แปลง URL ทั้งหมดเป็นรูปแบบลิงก์ markdown ที่คลิกได้** และสร้างเอกสารบันทึกเป็นไฟล์ `.md` ตรวจสอบให้แน่ใจว่าผู้ใช้สามารถคลิกลิงก์เพื่อเปิด workflow โดยตรงในเบราว์เซอร์ รวมทั้งชื่อเรื่อง คำอธิบาย และคำแนะนำการใช้งานของแต่ละ workflow",

		"param.file_path.description":            "เส้นทางสัมบูรณ์/สัมพันธ์ของไฟล์ JSON workflow n8n ที่จะตรวจสอบ (เช่น: './workflows/my_workflow.json')",
		"param.include_line_numbers.description": "รวมข้อมูลข้อผิดพลาดโดยละเอียดและบริบทตามหมายเลขบรรทัดหรือไม่ (ค่าเริ่มต้น: true)",
		"param.auto_fix_suggestions.description": "ให้คำแนะนำการแก้ไขอัตโนมัติหรือไม่ (ค่าเริ่มต้น: true)",
	}
}

// getTraditionalChineseMessages는 번체중국어(대만어) 메시지를 반환합니다
func (i *I18n) getTraditionalChineseMessages() map[string]string {
	return map[string]string{
		// 基本工具說明
		"tool.ping.description":             "基本 ping 測試以檢查伺服器是否正常運作",
		"tool.search_n8n_nodes.description": "搜尋 n8n 節點。根據關鍵字尋找相關節點並回傳分數。",
		"tool.search_workflow.description":  "搜尋 n8n 工作流程範本。根據關鍵字或類別尋找社群工作流程並連同 URL 一起回傳。(必須使用英文搜尋才能取得準確結果)\n\n📝 **搜尋後必要動作**: 所有搜尋結果必須儲存為 markdown 檔案 (.md)。URL 必須轉換為可點擊的連結格式 `[工作流程標題](URL)` 以便能從瀏覽器直接存取。",

		// 工作流程相關工具
		"tool.validate_workflow.description": `⚠️ **工作流程驗證工具**: 全面驗證 n8n 工作流程 JSON。

**驗證項目:**
- JSON 語法準確性和架構有效性
- 節點結構和缺少必要屬性檢查
- 節點連接狀態和資料流有效性
- 重複節點 ID 和循環參考檢查
- n8n 標準合規性

**重要:** 在創建或修改工作流程 JSON 後，請務必使用此工具進行驗證。未經驗證就部署到 n8n 可能會導致執行時錯誤。

**驗證失敗時:** 提供詳細的錯誤訊息和逐行修正指南。`,

		// 參數說明
		"param.query.description": "要搜尋的關鍵字或詞組 (必須使用英文以取得準確結果)",
		"param.limit.description": "回傳結果的最大數量 (預設: 5, 最大: 20)",

		// 工作流程搜尋參數
		"tool.search_workflow.param.query.description":       "要搜尋的關鍵字或詞組 (例如: AI, automation, chatbot)",
		"tool.search_workflow.param.max_results.description": "回傳結果的最大數量 (預設: 5, 最大: 15)",
		"tool.search_workflow.param.category.description":    "依工作流程類別篩選 (例如: airtable, blog, common, crm, email)",

		// 下一步指南
		"search_workflow.next_step": "分析搜尋到的工作流程以解釋每個工作流程的功能和用途，然後 **將所有 URL 轉換為可點擊的 markdown 連結格式** 並產生文件儲存為 `.md` 檔案。確保使用者可以點擊連結直接在瀏覽器中開啟工作流程，並包含每個工作流程的標題、說明和使用說明。",

		"param.file_path.description":            "要驗證的 n8n 工作流程 JSON 檔案的絕對/相對路徑 (例如: './workflows/my_workflow.json')",
		"param.include_line_numbers.description": "是否包含依行號的詳細錯誤訊息和內容 (預設: true)",
		"param.auto_fix_suggestions.description": "是否提供自動修正建議 (預設: true)",
	}
}
