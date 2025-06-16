# n8n Workflow MCP Server

[![한국어](https://img.shields.io/badge/README-한국어-blue)](README.md) | [![English](https://img.shields.io/badge/README-English-red)](README.en.md) | [![日本語](https://img.shields.io/badge/README-日本語-green)](README.jp.md)

n8nワークフローの検証と管理のためのMCP (Model Context Protocol) サーバーです。

## 🚀 主要機能

### 🌍 多言語サポート
- **韓国語 (ko)** - デフォルト言語
- **英語 (en)** - English support
- **日本語 (jp)** - 日本語サポート
- すべてのツールの説明とパラメータの説明が選択した言語で表示されます

### 📝 ワークフロー管理ツール
- **ping** - サーバー接続状態の確認
- **validate_workflow** - JSONファイルのn8nワークフロー構文検証

### 🔍 n8nノード検索ツール
- **search_n8n_nodes** - キーワードによるn8nノード検索

## 🛠 インストールとビルド

### 必要条件
- Go 1.23以上

### ビルド方法

#### 🚀 スタンドアロンモード（推奨）
ビルド時にすべてのn8nノードデータをバイナリに埋め込み、外部ファイルなしで実行可能：

```bash
# 現在のプラットフォーム用スタンドアロンビルド
make build-standalone
# または
make build

# Linux 64bitスタンドアロンビルド
make build-standalone-linux
# または
make build-linux

# Windows 64bitスタンドアロンビルド
make build-standalone-windows
# または
make build-windows
```

#### 特定プラットフォーム用ビルド

**macOS (Intel):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/n8n-mcp-server-darwin ./cmd/n8n-mcp-server
```

**macOS (Apple Silicon M1/M2):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/n8n-mcp-server-darwin-arm64 ./cmd/n8n-mcp-server
```

## 🎯 実行方法

### 🚀 スタンドアロンモード（推奨）
外部ファイルなしでバイナリのみで実行（デフォルトモード）：

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone

# Windows
build\n8n-mcp-server.exe -standalone

# Makefileを使用
make run
# または
make run-standalone
```

### 📁 ファイルシステムモード（開発用）
外部データファイルを使用する方式：

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone=false

# Windows
build\n8n-mcp-server.exe -standalone=false
```

### Goコマンドでの直接実行
```bash
# スタンドアロンモード（デフォルト）
go run ./cmd/n8n-mcp-server -standalone

# ファイルシステムモード
go run ./cmd/n8n-mcp-server -standalone=false
```

### 🌐 多言語サポート
サーバーは韓国語(ko)、英語(en)、日本語(jp)をサポートしています。`--lang`フラグで言語を設定できます：

```bash
# スタンドアロンモード
./build/n8n-mcp-server -standalone --lang ko  # 韓国語（デフォルト）
./build/n8n-mcp-server -standalone --lang en  # 英語
./build/n8n-mcp-server -standalone --lang jp  # 日本語

# ファイルシステムモード
./build/n8n-mcp-server -standalone=false --lang ko  # 韓国語（デフォルト）
./build/n8n-mcp-server -standalone=false --lang en  # 英語
./build/n8n-mcp-server -standalone=false --lang jp  # 日本語
```

### ⚙️ 検索重み調整
ノード検索時の各要素の重みを調整できます：

```bash
# スタンドアロンモードでの検索重み調整
./build/n8n-mcp-server -standalone \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0

# ファイルシステムモードでの検索重み調整とデータパス変更
./build/n8n-mcp-server -standalone=false \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0 \
  --data-path /custom/path/to/nodes
```

### 🆘 その他のオプション
```bash
# ヘルプを表示
./build/n8n-mcp-server --help

# バージョン確認
./build/n8n-mcp-server --version
```

## 🔧 Claude Code & Cursor設定

### macOS設定
設定ファイルの場所：`~/Library/Application Support/Claude/claude_desktop_config.json`

**スタンドアロンモード（推奨）：**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["--standalone", "--lang", "jp"]
    }
  }
}
```

**ファイルシステムモード（開発用）：**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["-standalone=false", "--lang", "jp"]
    }
  }
}
```

**Go実行モード（開発用）：**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "go",
      "args": ["run", "./cmd/n8n-mcp-server", "--standalone", "--lang", "jp"],
      "cwd": "/Users/{username}/{clone_path}"
    }
  }
}
```

### Windows設定
設定ファイルの場所：`%APPDATA%\Claude\claude_desktop_config.json`

**スタンドアロンモード（推奨）：**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["--standalone", "--lang", "jp"]
    }
  }
}
```

**ファイルシステムモード（開発用）：**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["-standalone=false", "--lang", "jp"]
    }
  }
}
```

**Go実行モード（開発用）：**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "go",
      "args": ["run", ".\\cmd\\n8n-mcp-server", "--standalone", "--lang", "jp"],
      "cwd": "C:\\path\\to\\n8n-workflow-mcp"
    }
  }
}
```

## 📁 プロジェクト構造

```
n8n-workflow-mcp/
├── cmd/
│   └── n8n-mcp-server/           # メインアプリケーションエントリーポイント
├── internal/
│   ├── config/                   # 設定管理
│   ├── handlers/                 # MCPリクエストハンドラー
│   ├── i18n/                     # 多言語サポート
│   ├── search/                   # 検索エンジン
│   ├── server/                   # MCPサーバー実装
│   ├── services/                 # ビジネスロジック
│   ├── store/                    # データストア
│   ├── validator/                # ワークフロー検証
│   └── workflow/                 # ワークフロー管理
├── pkg/
│   └── types/                    # 共通型定義
├── examples/
│   └── config/                   # Claude Code & Cursor設定例
├── test/                         # テストファイル
└── build/                        # ビルド成果物
```

## 🧪 テスト

```bash
# すべてのテストを実行
make test

# カバレッジ付きでテストを実行
make test-coverage
```

## 🛠 開発

```bash
# コードフォーマット
make fmt

# リンティング
make lint

# モジュール整理
make mod-tidy

# ビルドファイル削除
make clean
```

## 📊 モード比較

| 機能 | スタンドアロンモード | ファイルシステムモード |
|------|---------------------|------------------------|
| **デプロイ容易性** | ✅ 単一バイナリ | ❌ 追加ファイル必要 |
| **実行速度** | ✅ 高速 | ✅ 高速 |
| **メモリ使用量** | ⚠️ やや高い | ✅ 低い |
| **開発容易性** | ⚠️ リビルド必要 | ✅ リアルタイム反映 |
| **本番環境** | ✅ 推奨 | ❌ 非推奨 |
| **開発環境** | ✅ 使用可能 | ✅ 推奨 |

## 🔗 関連リンク

- [n8n公式ドキュメント](https://docs.n8n.io/)
- [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go)
- [Claude Code & Cursor](https://claude.ai/download) 