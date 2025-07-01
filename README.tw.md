# n8n Workflow MCP Server

[![한국어](https://img.shields.io/badge/README-한국어-blue)](README.md) | [![English](https://img.shields.io/badge/README-English-red)](README.en.md) | [![日本語](https://img.shields.io/badge/README-日本語-green)](README.jp.md) | [![Tiếng Việt](https://img.shields.io/badge/README-Tiếng%20Việt-yellow)](README.vn.md) | [![繁體中文](https://img.shields.io/badge/README-繁體中文-purple)](README.tw.md) | [![ไทย](https://img.shields.io/badge/README-ไทย-orange)](README.th.md)

## 🎬 示範影片

觀看工作流程搜尋功能的示範影片：

<video src="https://github.com/user-attachments/assets/1ec10869-be02-4ac7-a962-a6fdf865fffa" width="352" height="720"></video>

用於驗證和管理 n8n 工作流程的 MCP (Model Context Protocol) 伺服器。

## 🚀 主要功能

### 🌍 多語言支援
- **韓語 (ko)** - 預設語言
- **英語 (en)** - English support
- **日語 (jp)** - 日本語サポート
- **越南語 (vn)** - Hỗ trợ tiếng Việt
- **繁體中文 (tw)** - 繁體中文支援
- **泰語 (th)** - การสนับสนุนภาษาไทย
- 所有工具說明和參數說明都會以選定的語言顯示

### 📝 工作流程管理工具
- **ping** - 檢查伺服器連線狀態
- **validate_workflow** - 驗證 JSON 檔案中的 n8n 工作流程語法

### 🔍 搜尋工具
- **search_n8n_nodes** - 使用關鍵字搜尋 n8n 節點
- **search_workflow** - 搜尋社群 n8n 工作流程範本

## 🛠 安裝和建置

### 必要條件
- Go 1.23 或更高版本

### 建置方法

#### 🚀 獨立模式 (建議)
在建置時將所有 n8n 節點資料包含到二進位檔案中，無需外部檔案即可執行：

```bash
# 為目前平台建置獨立版本
make build-standalone
# 或
make build

# 為 Linux 64位元建置獨立版本
make build-standalone-linux
# 或
make build-linux

# 為 Windows 64位元建置獨立版本
make build-standalone-windows
# 或
make build-windows
```

#### 特定平台建置

**macOS (Intel):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/n8n-mcp-server-darwin ./cmd/n8n-mcp-server
```

**macOS (Apple Silicon M1/M2):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/n8n-mcp-server-darwin-arm64 ./cmd/n8n-mcp-server
```

## 🎯 執行方法

### 🚀 獨立模式 (建議)
僅使用二進位檔案執行，無需外部檔案（預設模式）：

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone

# Windows
build\n8n-mcp-server.exe -standalone

# 使用 Makefile
make run
# 或
make run-standalone
```

### 📁 檔案系統模式 (開發用)
使用外部資料檔案的方式：

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone=false

# Windows
build\n8n-mcp-server.exe -standalone=false
```

### 直接使用 Go 指令執行
```bash
# 獨立模式（預設）
go run ./cmd/n8n-mcp-server -standalone

# 檔案系統模式
go run ./cmd/n8n-mcp-server -standalone=false
```

### 🌐 多語言支援
伺服器支援韓語(ko)、英語(en)、日語(jp)、越南語(vn)、繁體中文(tw)、泰語(th)。您可以使用 `--lang` 旗標設定語言：

```bash
# 獨立模式
./build/n8n-mcp-server -standalone --lang ko  # 韓語（預設）
./build/n8n-mcp-server -standalone --lang en  # 英語
./build/n8n-mcp-server -standalone --lang jp  # 日語
./build/n8n-mcp-server -standalone --lang vn  # 越南語
./build/n8n-mcp-server -standalone --lang tw  # 繁體中文
./build/n8n-mcp-server -standalone --lang th  # 泰語

# 檔案系統模式
./build/n8n-mcp-server -standalone=false --lang ko  # 韓語（預設）
./build/n8n-mcp-server -standalone=false --lang en  # 英語
./build/n8n-mcp-server -standalone=false --lang jp  # 日語
./build/n8n-mcp-server -standalone=false --lang vn  # 越南語
./build/n8n-mcp-server -standalone=false --lang tw  # 繁體中文
./build/n8n-mcp-server -standalone=false --lang th  # 泰語
```

### ⚙️ 調整搜尋權重
您可以調整搜尋節點時各個元素的權重：

```bash
# 在獨立模式中調整搜尋權重
./build/n8n-mcp-server -standalone \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0

# 在檔案系統模式中調整搜尋權重並變更資料路徑
./build/n8n-mcp-server -standalone=false \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0 \
  --data-path /custom/path/to/nodes
```

### 🆘 其他選項
```bash
# 查看說明
./build/n8n-mcp-server --help

# 檢查版本
./build/n8n-mcp-server --version
```

## 🔧 Claude Code & Cursor 設定

### macOS 設定
設定檔位置：`~/Library/Application Support/Claude/claude_desktop_config.json`

**獨立模式 (建議):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["--standalone", "--lang", "tw"]
    }
  }
}
```

**檔案系統模式 (開發用):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["-standalone=false", "--lang", "tw"]
    }
  }
}
```

**Go 執行模式 (開發用):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "go",
      "args": ["run", "./cmd/n8n-mcp-server", "--standalone", "--lang", "tw"],
      "cwd": "/Users/{username}/{clone_path}"
    }
  }
}
```

### Windows 設定
設定檔位置：`%APPDATA%\Claude\claude_desktop_config.json`

**獨立模式 (建議):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["--standalone", "--lang", "tw"]
    }
  }
}
```

**檔案系統模式 (開發用):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["-standalone=false", "--lang", "tw"]
    }
  }
}
```

**Go 執行模式 (開發用):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "go",
      "args": ["run", ".\\cmd\\n8n-mcp-server", "--standalone", "--lang", "tw"],
      "cwd": "C:\\path\\to\\n8n-workflow-mcp"
    }
  }
}
```

## 📁 專案結構

```
n8n-workflow-mcp/
├── cmd/
│   └── n8n-mcp-server/           # 主要應用程式進入點
├── internal/
│   ├── config/                   # 設定管理
│   ├── handlers/                 # MCP 請求處理器
│   ├── i18n/                     # 多語言支援
│   ├── search/                   # 搜尋引擎
│   ├── server/                   # MCP 伺服器實作
│   ├── services/                 # 商業邏輯
│   ├── store/                    # 資料儲存
│   ├── validator/                # 工作流程驗證
│   └── workflow/                 # 工作流程管理
├── pkg/
│   └── types/                    # 共用型別定義
├── examples/
│   └── config/                   # Claude Code & Cursor 設定範例
├── test/                         # 測試檔案
└── build/                        # 建置結果
```

## 🧪 測試

```bash
# 執行所有測試
make test

# 執行測試並顯示覆蓋率
make test-coverage
```

## 🛠 開發

```bash
# 程式碼格式化
make fmt

# 程式碼檢查
make lint

# 模組整理
make mod-tidy

# 清理建置檔案
make clean
```

## 📊 模式比較

| 功能 | 獨立模式 | 檔案系統模式 |
|------|----------|--------------|
| **部署便利性** | ✅ 單一二進位檔案 | ❌ 需要額外檔案 |
| **執行速度** | ✅ 快速 | ✅ 快速 |
| **記憶體使用量** | ⚠️ 稍高 | ✅ 較低 |
| **開發便利性** | ⚠️ 需要重新建置 | ✅ 即時反映 |
| **生產環境** | ✅ 建議 | ❌ 不建議 |
| **開發環境** | ✅ 可使用 | ✅ 建議 |

## 🔗 相關連結

- [n8n 官方文件](https://docs.n8n.io/)
- [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go)
- [Claude Code](https://claude.ai/download)
- [Cursor](https://cursor.com) 