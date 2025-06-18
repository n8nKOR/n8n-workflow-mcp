# Claude Code & Cursor 設定例

このディレクトリには、Claude Code & Cursor アプリで n8n-workflow-mcp サーバーを使用するための設定ファイル例が含まれています。

## 📁 ファイル一覧

### 🚀 Standalone モード（推奨）

#### 1. `claude_desktop_config_standalone.json`
- **用途**: Standalone モードでビルドされたバイナリを使用（推奨）
- **利点**: 
  - 外部ファイルなしで単一バイナリとして実行
  - 高速な起動時間
  - 簡単なデプロイ
- **対象**: 一般ユーザー、本番環境

#### 2. `claude_desktop_config_standalone_en.json`
- **用途**: 英語インターフェースで Standalone モードを実行
- **特徴**: すべてのツール説明が英語で表示
- **対象**: 英語圏ユーザー

#### 3. `claude_desktop_config_standalone_jp.json`
- **用途**: 日本語インターフェースで Standalone モードを実行
- **特徴**: すべてのツール説明が日本語で表示
- **対象**: 日本語ユーザー

### 🛠 開発モード

#### 4. `claude_desktop_config_dev_filesystem.json`
- **用途**: ファイルシステムモードで開発用実行
- **利点**: 
  - ソースコード変更時にデータファイル修正が即座に反映
  - 開発中のデバッグが便利
- **欠点**: 外部データファイルへの依存
- **対象**: 開発者

#### 5. `claude_desktop_config_dev_go.json`
- **用途**: Go コマンドで直接実行（開発用）
- **利点**: 
  - ソースコード変更時の自動コンパイル
  - 開発中のリアルタイム反映
- **欠点**: Go 環境が必要、実行時間の増加
- **対象**: 開発者

### 🪟 Windows 専用

#### 6. `claude_desktop_config_windows_standalone.json`
- **用途**: Windows で Standalone モードを実行
- **特徴**: Windows パス形式を使用
- **対象**: Windows ユーザー

#### 7. `claude_desktop_config_windows_dev.json`
- **用途**: Windows で Go コマンドによる開発用実行
- **特徴**: Windows パス形式と開発オプション
- **対象**: Windows 開発者

## 🎯 推奨設定

### 一般ユーザー
1. **macOS/Linux**: `claude_desktop_config_standalone.json`
2. **Windows**: `claude_desktop_config_windows_standalone.json`

### 開発者
1. **開発時**: `claude_desktop_config_dev_go.json`
2. **テスト時**: `claude_desktop_config_dev_filesystem.json`

### 多言語ユーザー
1. **英語**: `claude_desktop_config_standalone_en.json`
2. **日本語**: `claude_desktop_config_standalone_jp.json`
3. **韓国語**: `claude_desktop_config_standalone.json`（デフォルト）

## 🛠️ 利用可能なツール

### 1. `search_n8n_nodes`
- **用途**: n8nノード検索
- **機能**: キーワードに基づいてn8nノードを検索し、スコア付きで結果を返す
- **パラメータ**:
  - `query` (必須): 検索するキーワード
  - `limit` (オプション): 返す結果の最大数 (デフォルト: 5、最大: 20)

### 2. `search_workflow` ✨ **新機能追加！**
- **用途**: n8nワークフローテンプレート検索
- **機能**: コミュニティワークフローテンプレートを検索し、ダウンロードURLを提供
- **パラメータ**:
  - `query` (必須): 検索するキーワード (例: "AI"、"automation"、"chatbot")
  - `max_results` (オプション): 返す結果の最大数 (デフォルト: 5、最大: 15)
  - `category` (オプション): カテゴリフィルター (例: airtable、blog、common、crm、email)
- **使用例**:
  - AIチャットボットワークフローを探す: `query: "AI chatbot"`
  - メール自動化ワークフロー: `query: "email automation", category: "email"`
  - Airtable統合ワークフロー: `category: "airtable"`

### 3. `validate_workflow`
- **用途**: n8nワークフローJSON検証
- **機能**: ワークフローJSONファイルの構文と構造を検証
- **パラメータ**:
  - `file_path` (必須): 検証するJSONファイルのパス
  - `include_line_numbers` (オプション): 詳細な行番号情報を含める (デフォルト: true)
  - `auto_fix_suggestions` (オプション): 自動修正提案を提供 (デフォルト: true)

### 4. `ping`
- **用途**: サーバーステータス確認
- **機能**: MCPサーバーが正常に動作しているかを確認する基本テスト

## 🔧 設定方法

### 1. 設定ファイルの場所を確認
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

### 2. 設定ファイルの選択と適用
1. 上記のファイル一覧から用途に合った設定ファイルを選択します。
2. 選択したファイルの内容をコピーします。
3. Claude Code & Cursor 設定ファイルに貼り付けます。
4. ファイルパスを自分の環境に合わせて修正します。

### 3. パス修正ガイド

**macOS/Linux 例:**
```json
"command": "/Users/john/projects/n8n-workflow-mcp/build/n8n-mcp-server"
```

**Windows 例:**
```json
"command": "C:\\Users\\john\\projects\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe"
```

### 4. Claude Code & Cursor の再起動
設定を適用した後、Claude Code & Cursor アプリを完全に終了してから再起動します。

## ⚙️ 高度な設定オプション

### 検索重み調整
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/path/to/n8n-mcp-server",
      "args": [
        "--standalone",
        "--lang", "jp",
        "--name-weight", "3.0",
        "--overview-weight", "2.0",
        "--case-weight", "1.5",
        "--op-weight", "1.0"
      ]
    }
  }
}
```

### カスタムデータパス（ファイルシステムモード）
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/path/to/n8n-mcp-server",
      "args": [
        "-standalone=false",
        "--data-path", "/custom/path/to/nodes",
        "--lang", "jp"
      ]
    }
  }
}
```

## 🔍 トラブルシューティング

### よくある問題

1. **サーバーが起動しない場合**
   - ファイルパスが正しいか確認
   - 実行権限があるか確認（`chmod +x`）
   - ビルドされたバイナリが存在するか確認

2. **権限エラーが発生する場合**
   ```bash
   chmod +x /path/to/n8n-mcp-server
   ```

3. **パス関連エラー**
   - 絶対パスの使用を推奨
   - Windows ではバックスラッシュのエスケープが必要: `\\`

### ログの確認
Claude Code & Cursor のログを確認するには:
- **macOS**: Console.app で「Claude」を検索
- **Windows**: イベントビューアーまたは Claude ログファイルを確認

## 📝 注意事項

- すべてのパスは絶対パスを使用する必要があります
- Windows でパスにバックスラッシュ（`\`）を使用する場合は、二重に入力してください: `\\`
- またはスラッシュ（`/`）を使用することもできます
- 設定変更後は必ず Claude Code & Cursor を再起動してください
- 開発モードは Go 環境がインストールされている必要があります

## 🆘 ヘルプ

追加のヘルプが必要な場合:
1. プロジェクトのメイン README.md を参照
2. GitHub Issues でお問い合わせ
3. 設定ファイル構文検証: [JSON Validator](https://jsonlint.com/) 