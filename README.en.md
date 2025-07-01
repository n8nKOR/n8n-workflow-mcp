# n8n Workflow MCP Server

[![한국어](https://img.shields.io/badge/README-한국어-blue)](README.md) | [![English](https://img.shields.io/badge/README-English-red)](README.en.md) | [![日本語](https://img.shields.io/badge/README-日本語-green)](README.jp.md) | [![Tiếng Việt](https://img.shields.io/badge/README-Tiếng%20Việt-yellow)](README.vn.md) | [![繁體中文](https://img.shields.io/badge/README-繁體中文-purple)](README.tw.md) | [![ไทย](https://img.shields.io/badge/README-ไทย-orange)](README.th.md)

## 🎬 Demo Video

Check out the video demonstrating the workflow search functionality:

<video src="https://github.com/user-attachments/assets/1ec10869-be02-4ac7-a962-a6fdf865fffa" width="352" height="720"></video>

An MCP (Model Context Protocol) server for n8n workflow validation and management.

## 🚀 Key Features

### 🌍 Multi-language Support
- **Korean (ko)** - Default language
- **English (en)** - English support
- **Japanese (jp)** - 日本語サポート
- **Vietnamese (vn)** - Hỗ trợ tiếng Việt
- **Traditional Chinese (tw)** - 繁體中文支援
- **Thai (th)** - การสนับสนุนภาษาไทย
- All tool descriptions and parameter descriptions are displayed in the selected language

### 📝 Workflow Management Tools
- **ping** - Check server connection status
- **validate_workflow** - Validate n8n workflow syntax in JSON files

### 🔍 Search Tools
- **search_n8n_nodes** - Search n8n nodes by keywords
- **search_workflow** - Search community n8n workflow templates

## 🛠 Installation and Build

### Requirements
- Go 1.23 or higher

### Build Instructions

#### 🚀 Standalone Mode (Recommended)
Embed all n8n node data into the binary at build time for execution without external files:

```bash
# Build standalone for current platform
make build-standalone
# or
make build

# Build standalone for Linux 64bit
make build-standalone-linux
# or
make build-linux

# Build standalone for Windows 64bit
make build-standalone-windows
# or
make build-windows
```

#### Platform-specific Build

**macOS (Intel):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/n8n-mcp-server-darwin ./cmd/n8n-mcp-server
```

**macOS (Apple Silicon M1/M2):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/n8n-mcp-server-darwin-arm64 ./cmd/n8n-mcp-server
```

## 🎯 How to Run

### 🚀 Standalone Mode (Recommended)
Run with binary only, no external files required (default mode):

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone

# Windows
build\n8n-mcp-server.exe -standalone

# Using Makefile
make run
# or
make run-standalone
```

### 📁 File System Mode (For Development)
Use external data files:

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone=false

# Windows
build\n8n-mcp-server.exe -standalone=false
```

### Direct Execution with Go
```bash
# Standalone mode (default)
go run ./cmd/n8n-mcp-server -standalone

# File system mode
go run ./cmd/n8n-mcp-server -standalone=false
```

### 🌐 Multi-language Support
The server supports Korean (ko), English (en), Japanese (jp), Vietnamese (vn), Traditional Chinese (tw), and Thai (th). You can set the language using the `--lang` flag:

```bash
# Standalone mode
./build/n8n-mcp-server -standalone --lang ko  # Korean (default)
./build/n8n-mcp-server -standalone --lang en  # English
./build/n8n-mcp-server -standalone --lang jp  # Japanese
./build/n8n-mcp-server -standalone --lang vn  # Vietnamese
./build/n8n-mcp-server -standalone --lang tw  # Traditional Chinese
./build/n8n-mcp-server -standalone --lang th  # Thai

# File system mode
./build/n8n-mcp-server -standalone=false --lang ko  # Korean (default)
./build/n8n-mcp-server -standalone=false --lang en  # English
./build/n8n-mcp-server -standalone=false --lang jp  # Japanese
./build/n8n-mcp-server -standalone=false --lang vn  # Vietnamese
./build/n8n-mcp-server -standalone=false --lang tw  # Traditional Chinese
./build/n8n-mcp-server -standalone=false --lang th  # Thai
```

### ⚙️ Search Weight Adjustment
You can adjust the weight of each element when searching nodes:

```bash
# Adjust search weights in standalone mode
./build/n8n-mcp-server -standalone \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0

# Adjust search weights and change data path in file system mode
./build/n8n-mcp-server -standalone=false \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0 \
  --data-path /custom/path/to/nodes
```

### 🆘 Other Options
```bash
# Show help
./build/n8n-mcp-server --help

# Check version
./build/n8n-mcp-server --version
```

## 🔧 Claude Code & Cursor Configuration

### macOS Configuration
Configuration file location: `~/Library/Application Support/Claude/claude_desktop_config.json`

**Standalone Mode (Recommended):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["--standalone", "--lang", "en"]
    }
  }
}
```

**File System Mode (For Development):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["-standalone=false", "--lang", "en"]
    }
  }
}
```

**Go Execution Mode (For Development):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "go",
      "args": ["run", "./cmd/n8n-mcp-server", "--standalone", "--lang", "en"],
      "cwd": "/Users/{username}/{clone_path}"
    }
  }
}
```

### Windows Configuration
Configuration file location: `%APPDATA%\Claude\claude_desktop_config.json`

**Standalone Mode (Recommended):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["--standalone", "--lang", "en"]
    }
  }
}
```

**File System Mode (For Development):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["-standalone=false", "--lang", "en"]
    }
  }
}
```

**Go Execution Mode (For Development):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "go",
      "args": ["run", ".\\cmd\\n8n-mcp-server", "--standalone", "--lang", "en"],
      "cwd": "C:\\path\\to\\n8n-workflow-mcp"
    }
  }
}
```

## 📁 Project Structure

```
n8n-workflow-mcp/
├── cmd/
│   └── n8n-mcp-server/           # Main application entry point
├── internal/
│   ├── config/                   # Configuration management
│   ├── handlers/                 # MCP request handlers
│   ├── i18n/                     # Multi-language support
│   ├── search/                   # Search engine
│   ├── server/                   # MCP server implementation
│   ├── services/                 # Business logic
│   ├── store/                    # Data store
│   ├── validator/                # Workflow validation
│   └── workflow/                 # Workflow management
├── pkg/
│   └── types/                    # Common type definitions
├── examples/
│   └── config/                   # Claude Code & Cursor configuration examples
├── test/                         # Test files
└── build/                        # Build artifacts
```

## 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

## 🛠 Development

```bash
# Code formatting
make fmt

# Linting
make lint

# Clean modules
make mod-tidy

# Clean build files
make clean
```

## 📊 Mode Comparison

| Feature | Standalone Mode | File System Mode |
|---------|-----------------|-------------------|
| **Deployment Ease** | ✅ Single binary | ❌ Additional files required |
| **Execution Speed** | ✅ Fast | ✅ Fast |
| **Memory Usage** | ⚠️ Slightly higher | ✅ Lower |
| **Development Ease** | ⚠️ Rebuild required | ✅ Real-time reflection |
| **Production Environment** | ✅ Recommended | ❌ Not recommended |
| **Development Environment** | ✅ Available | ✅ Recommended |

## 🔗 Related Links

- [n8n Official Documentation](https://docs.n8n.io/)
- [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go)
- [Claude Code](https://claude.ai/download)
- [Cursor](https://cursor.com) 