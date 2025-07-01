# n8n Workflow MCP Server

[![한국어](https://img.shields.io/badge/README-한국어-blue)](README.md) | [![English](https://img.shields.io/badge/README-English-red)](README.en.md) | [![日本語](https://img.shields.io/badge/README-日本語-green)](README.jp.md) | [![Tiếng Việt](https://img.shields.io/badge/README-Tiếng%20Việt-yellow)](README.vn.md) | [![繁體中文](https://img.shields.io/badge/README-繁體中文-purple)](README.tw.md) | [![ไทย](https://img.shields.io/badge/README-ไทย-orange)](README.th.md)

## 🎬 Video Demo

Xem video demo tính năng tìm kiếm workflow:

<video src="https://github.com/user-attachments/assets/1ec10869-be02-4ac7-a962-a6fdf865fffa" width="352" height="720"></video>

Máy chủ MCP (Model Context Protocol) để xác thực và quản lý workflow n8n.

## 🚀 Tính năng chính

### 🌍 Hỗ trợ đa ngôn ngữ
- **Tiếng Hàn (ko)** - Ngôn ngữ mặc định
- **Tiếng Anh (en)** - Hỗ trợ tiếng Anh
- **Tiếng Nhật (jp)** - サポート日本語
- **Tiếng Việt (vn)** - Hỗ trợ tiếng Việt
- **Tiếng Trung Phồn thể (tw)** - 繁體中文支援
- **Tiếng Thái (th)** - การสนับสนุนภาษาไทย
- Tất cả các mô tả công cụ và tham số sẽ được hiển thị bằng ngôn ngữ đã chọn

### 📝 Công cụ quản lý workflow
- **ping** - Kiểm tra trạng thái kết nối máy chủ
- **validate_workflow** - Xác thực cú pháp workflow n8n trong file JSON

### 🔍 Công cụ tìm kiếm
- **search_n8n_nodes** - Tìm kiếm node n8n bằng từ khóa
- **search_workflow** - Tìm kiếm template workflow n8n của cộng đồng

## 🛠 Cài đặt và Build

### Yêu cầu
- Go 1.23 trở lên

### Cách build

#### 🚀 Chế độ Standalone (Khuyến nghị)
Bao gồm tất cả dữ liệu node n8n vào binary tại thời điểm build, có thể chạy mà không cần file bên ngoài:

```bash
# Build standalone cho platform hiện tại
make build-standalone
# hoặc
make build

# Build standalone cho Linux 64bit
make build-standalone-linux
# hoặc
make build-linux

# Build standalone cho Windows 64bit
make build-standalone-windows
# hoặc
make build-windows
```

#### Build cho platform cụ thể

**macOS (Intel):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/n8n-mcp-server-darwin ./cmd/n8n-mcp-server
```

**macOS (Apple Silicon M1/M2):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/n8n-mcp-server-darwin-arm64 ./cmd/n8n-mcp-server
```

## 🎯 Cách chạy

### 🚀 Chế độ Standalone (Khuyến nghị)
Chạy chỉ với binary mà không cần file bên ngoài (chế độ mặc định):

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone

# Windows
build\n8n-mcp-server.exe -standalone

# Sử dụng Makefile
make run
# hoặc
make run-standalone
```

### 📁 Chế độ File System (Cho phát triển)
Sử dụng file dữ liệu bên ngoài:

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone=false

# Windows
build\n8n-mcp-server.exe -standalone=false
```

### Chạy trực tiếp bằng lệnh Go
```bash
# Chế độ Standalone (mặc định)
go run ./cmd/n8n-mcp-server -standalone

# Chế độ File System
go run ./cmd/n8n-mcp-server -standalone=false
```

### 🌐 Hỗ trợ đa ngôn ngữ
Máy chủ hỗ trợ tiếng Hàn(ko), tiếng Anh(en), tiếng Nhật(jp), tiếng Việt(vn), tiếng Trung phồn thể(tw), tiếng Thái(th). Bạn có thể đặt ngôn ngữ bằng flag `--lang`:

```bash
# Chế độ Standalone
./build/n8n-mcp-server -standalone --lang ko  # Tiếng Hàn (mặc định)
./build/n8n-mcp-server -standalone --lang en  # Tiếng Anh
./build/n8n-mcp-server -standalone --lang jp  # Tiếng Nhật
./build/n8n-mcp-server -standalone --lang vn  # Tiếng Việt
./build/n8n-mcp-server -standalone --lang tw  # Tiếng Trung phồn thể
./build/n8n-mcp-server -standalone --lang th  # Tiếng Thái

# Chế độ File System
./build/n8n-mcp-server -standalone=false --lang ko  # Tiếng Hàn (mặc định)
./build/n8n-mcp-server -standalone=false --lang en  # Tiếng Anh
./build/n8n-mcp-server -standalone=false --lang jp  # Tiếng Nhật
./build/n8n-mcp-server -standalone=false --lang vn  # Tiếng Việt
./build/n8n-mcp-server -standalone=false --lang tw  # Tiếng Trung phồn thể
./build/n8n-mcp-server -standalone=false --lang th  # Tiếng Thái
```

### ⚙️ Điều chỉnh trọng số tìm kiếm
Bạn có thể điều chỉnh trọng số của từng yếu tố khi tìm kiếm node:

```bash
# Điều chỉnh trọng số tìm kiếm trong chế độ Standalone
./build/n8n-mcp-server -standalone \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0

# Điều chỉnh trọng số tìm kiếm trong chế độ File System và thay đổi đường dẫn dữ liệu
./build/n8n-mcp-server -standalone=false \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0 \
  --data-path /custom/path/to/nodes
```

### 🆘 Các tùy chọn khác
```bash
# Xem trợ giúp
./build/n8n-mcp-server --help

# Kiểm tra phiên bản
./build/n8n-mcp-server --version
```

## 🔧 Cấu hình Claude Code & Cursor

### Cấu hình macOS
Vị trí file cấu hình: `~/Library/Application Support/Claude/claude_desktop_config.json`

**Chế độ Standalone (Khuyến nghị):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["--standalone", "--lang", "vn"]
    }
  }
}
```

**Chế độ File System (Cho phát triển):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["-standalone=false", "--lang", "vn"]
    }
  }
}
```

**Chế độ Go (Cho phát triển):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "go",
      "args": ["run", "./cmd/n8n-mcp-server", "--standalone", "--lang", "vn"],
      "cwd": "/Users/{username}/{clone_path}"
    }
  }
}
```

### Cấu hình Windows
Vị trí file cấu hình: `%APPDATA%\Claude\claude_desktop_config.json`

**Chế độ Standalone (Khuyến nghị):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["--standalone", "--lang", "vn"]
    }
  }
}
```

**Chế độ File System (Cho phát triển):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["-standalone=false", "--lang", "vn"]
    }
  }
}
```

**Chế độ Go (Cho phát triển):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "go",
      "args": ["run", ".\\cmd\\n8n-mcp-server", "--standalone", "--lang", "vn"],
      "cwd": "C:\\path\\to\\n8n-workflow-mcp"
    }
  }
}
```

## 📁 Cấu trúc dự án

```
n8n-workflow-mcp/
├── cmd/
│   └── n8n-mcp-server/           # Entry point ứng dụng chính
├── internal/
│   ├── config/                   # Quản lý cấu hình
│   ├── handlers/                 # Handler yêu cầu MCP
│   ├── i18n/                     # Hỗ trợ đa ngôn ngữ
│   ├── search/                   # Engine tìm kiếm
│   ├── server/                   # Triển khai server MCP
│   ├── services/                 # Logic nghiệp vụ
│   ├── store/                    # Kho dữ liệu
│   ├── validator/                # Xác thực workflow
│   └── workflow/                 # Quản lý workflow
├── pkg/
│   └── types/                    # Định nghĩa kiểu chung
├── examples/
│   └── config/                   # Ví dụ cấu hình Claude Code & Cursor
├── test/                         # File test
└── build/                        # Kết quả build
```

## 🧪 Testing

```bash
# Chạy tất cả test
make test

# Chạy test với coverage
make test-coverage
```

## 🛠 Phát triển

```bash
# Format code
make fmt

# Linting
make lint

# Dọn dẹp module
make mod-tidy

# Xóa file build
make clean
```

## 📊 So sánh chế độ

| Tính năng | Chế độ Standalone | Chế độ File System |
|-----------|-------------------|-------------------|
| **Dễ triển khai** | ✅ Binary đơn | ❌ Cần file bổ sung |
| **Tốc độ chạy** | ✅ Nhanh | ✅ Nhanh |
| **Sử dụng bộ nhớ** | ⚠️ Hơi cao | ✅ Thấp |
| **Dễ phát triển** | ⚠️ Cần rebuild | ✅ Phản ánh realtime |
| **Môi trường sản xuất** | ✅ Khuyến nghị | ❌ Không khuyến nghị |
| **Môi trường phát triển** | ✅ Có thể sử dụng | ✅ Khuyến nghị |

## 🔗 Liên kết liên quan

- [Tài liệu chính thức n8n](https://docs.n8n.io/)
- [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go)
- [Claude Code](https://claude.ai/download)
- [Cursor](https://cursor.com) 