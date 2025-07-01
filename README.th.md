# n8n Workflow MCP Server

[![한국어](https://img.shields.io/badge/README-한국어-blue)](README.md) | [![English](https://img.shields.io/badge/README-English-red)](README.en.md) | [![日本語](https://img.shields.io/badge/README-日本語-green)](README.jp.md) | [![Tiếng Việt](https://img.shields.io/badge/README-Tiếng%20Việt-yellow)](README.vn.md) | [![繁體中文](https://img.shields.io/badge/README-繁體中文-purple)](README.tw.md) | [![ไทย](https://img.shields.io/badge/README-ไทย-orange)](README.th.md)

## 🎬 วิดีโอสาธิต

ดูวิดีโอสาธิตฟีเจอร์การค้นหาเวิร์กโฟลว์:

<video src="https://github.com/user-attachments/assets/1ec10869-be02-4ac7-a962-a6fdf865fffa" width="352" height="720"></video>

เซิร์ฟเวอร์ MCP (Model Context Protocol) สำหรับการตรวจสอบและจัดการเวิร์กโฟลว์ n8n

## 🚀 ฟีเจอร์หลัก

### 🌍 รองรับหลายภาษา
- **เกาหลี (ko)** - ภาษาเริ่มต้น
- **อังกฤษ (en)** - English support
- **ญี่ปุ่น (jp)** - 日本語サポート
- **เวียดนาม (vn)** - Hỗ trợ tiếng Việt
- **จีนดั้งเดิม (tw)** - 繁體中文支援
- **ไทย (th)** - การสนับสนุนภาษาไทย
- คำอธิบายเครื่องมือและพารามิเตอร์ทั้งหมดจะแสดงในภาษาที่เลือก

### 📝 เครื่องมือจัดการเวิร์กโฟลว์
- **ping** - ตรวจสอบสถานะการเชื่อมต่อเซิร์ฟเวอร์
- **validate_workflow** - ตรวจสอบไวยากรณ์เวิร์กโฟลว์ n8n ในไฟล์ JSON

### 🔍 เครื่องมือค้นหา
- **search_n8n_nodes** - ค้นหาโหนด n8n ด้วยคำหลัก
- **search_workflow** - ค้นหาเทมเพลตเวิร์กโฟลว์ n8n ของชุมชน

## 🛠 การติดตั้งและบิลด์

### ข้อกำหนด
- Go 1.23 ขึ้นไป

### วิธีการบิลด์

#### 🚀 โหมด Standalone (แนะนำ)
รวมข้อมูลโหนด n8n ทั้งหมดไว้ในไบนารีขณะบิลด์ สามารถรันได้โดยไม่ต้องใช้ไฟล์ภายนอก:

```bash
# บิลด์ standalone สำหรับแพลตฟอร์มปัจจุบัน
make build-standalone
# หรือ
make build

# บิลด์ standalone สำหรับ Linux 64bit
make build-standalone-linux
# หรือ
make build-linux

# บิลด์ standalone สำหรับ Windows 64bit
make build-standalone-windows
# หรือ
make build-windows
```

#### บิลด์สำหรับแพลตฟอร์มเฉพาะ

**macOS (Intel):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/n8n-mcp-server-darwin ./cmd/n8n-mcp-server
```

**macOS (Apple Silicon M1/M2):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/n8n-mcp-server-darwin-arm64 ./cmd/n8n-mcp-server
```

## 🎯 วิธีการรัน

### 🚀 โหมด Standalone (แนะนำ)
รันด้วยไบนารีเพียงอย่างเดียวโดยไม่ต้องใช้ไฟล์ภายนอก (โหมดเริ่มต้น):

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone

# Windows
build\n8n-mcp-server.exe -standalone

# ใช้ Makefile
make run
# หรือ
make run-standalone
```

### 📁 โหมด File System (สำหรับการพัฒนา)
ใช้ไฟล์ข้อมูลภายนอก:

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone=false

# Windows
build\n8n-mcp-server.exe -standalone=false
```

### รันโดยตรงด้วยคำสั่ง Go
```bash
# โหมด Standalone (เริ่มต้น)
go run ./cmd/n8n-mcp-server -standalone

# โหมด File System
go run ./cmd/n8n-mcp-server -standalone=false
```

### 🌐 รองรับหลายภาษา
เซิร์ฟเวอร์รองรับเกาหลี(ko), อังกฤษ(en), ญี่ปุ่น(jp), เวียดนาม(vn), จีนดั้งเดิม(tw), ไทย(th) คุณสามารถตั้งค่าภาษาด้วยแฟล็ก `--lang`:

```bash
# โหมด Standalone
./build/n8n-mcp-server -standalone --lang ko  # เกาหลี (เริ่มต้น)
./build/n8n-mcp-server -standalone --lang en  # อังกฤษ
./build/n8n-mcp-server -standalone --lang jp  # ญี่ปุ่น
./build/n8n-mcp-server -standalone --lang vn  # เวียดนาม
./build/n8n-mcp-server -standalone --lang tw  # จีนดั้งเดิม
./build/n8n-mcp-server -standalone --lang th  # ไทย

# โหมด File System
./build/n8n-mcp-server -standalone=false --lang ko  # เกาหลี (เริ่มต้น)
./build/n8n-mcp-server -standalone=false --lang en  # อังกฤษ
./build/n8n-mcp-server -standalone=false --lang jp  # ญี่ปุ่น
./build/n8n-mcp-server -standalone=false --lang vn  # เวียดนาม
./build/n8n-mcp-server -standalone=false --lang tw  # จีนดั้งเดิม
./build/n8n-mcp-server -standalone=false --lang th  # ไทย
```

### ⚙️ ปรับค่าน้ำหนักการค้นหา
คุณสามารถปรับค่าน้ำหนักของแต่ละองค์ประกอบเมื่อค้นหาโหนด:

```bash
# ปรับค่าน้ำหนักการค้นหาในโหมด Standalone
./build/n8n-mcp-server -standalone \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0

# ปรับค่าน้ำหนักการค้นหาในโหมด File System และเปลี่ยนพาธข้อมูล
./build/n8n-mcp-server -standalone=false \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0 \
  --data-path /custom/path/to/nodes
```

### 🆘 ตัวเลือกอื่นๆ
```bash
# ดูความช่วยเหลือ
./build/n8n-mcp-server --help

# ตรวจสอบเวอร์ชัน
./build/n8n-mcp-server --version
```

## 🔧 การตั้งค่า Claude Code & Cursor

### การตั้งค่า macOS
ตำแหน่งไฟล์การตั้งค่า: `~/Library/Application Support/Claude/claude_desktop_config.json`

**โหมด Standalone (แนะนำ):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["--standalone", "--lang", "th"]
    }
  }
}
```

**โหมด File System (สำหรับการพัฒนา):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["-standalone=false", "--lang", "th"]
    }
  }
}
```

**โหมด Go (สำหรับการพัฒนา):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "go",
      "args": ["run", "./cmd/n8n-mcp-server", "--standalone", "--lang", "th"],
      "cwd": "/Users/{username}/{clone_path}"
    }
  }
}
```

### การตั้งค่า Windows
ตำแหน่งไฟล์การตั้งค่า: `%APPDATA%\Claude\claude_desktop_config.json`

**โหมด Standalone (แนะนำ):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["--standalone", "--lang", "th"]
    }
  }
}
```

**โหมด File System (สำหรับการพัฒนา):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["-standalone=false", "--lang", "th"]
    }
  }
}
```

**โหมด Go (สำหรับการพัฒนา):**
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "go",
      "args": ["run", ".\\cmd\\n8n-mcp-server", "--standalone", "--lang", "th"],
      "cwd": "C:\\path\\to\\n8n-workflow-mcp"
    }
  }
}
```

## 📁 โครงสร้างโปรเจ็กต์

```
n8n-workflow-mcp/
├── cmd/
│   └── n8n-mcp-server/           # จุดเข้าแอพพลิเคชันหลัก
├── internal/
│   ├── config/                   # การจัดการการตั้งค่า
│   ├── handlers/                 # ตัวจัดการคำขอ MCP
│   ├── i18n/                     # รองรับหลายภาษา
│   ├── search/                   # เอ็นจินค้นหา
│   ├── server/                   # การใช้งานเซิร์ฟเวอร์ MCP
│   ├── services/                 # ตรรกะทางธุรกิจ
│   ├── store/                    # ที่เก็บข้อมูล
│   ├── validator/                # การตรวจสอบเวิร์กโฟลว์
│   └── workflow/                 # การจัดการเวิร์กโฟลว์
├── pkg/
│   └── types/                    # การกำหนดประเภททั่วไป
├── examples/
│   └── config/                   # ตัวอย่างการตั้งค่า Claude Code & Cursor
├── test/                         # ไฟล์ทดสอบ
└── build/                        # ผลลัพธ์การบิลด์
```

## 🧪 การทดสอบ

```bash
# รันการทดสอบทั้งหมด
make test

# รันการทดสอบพร้อม coverage
make test-coverage
```

## 🛠 การพัฒนา

```bash
# จัดรูปแบบโค้ด
make fmt

# การตรวจสอบโค้ด
make lint

# จัดระเบียบโมดูล
make mod-tidy

# ลบไฟล์บิลด์
make clean
```

## 📊 เปรียบเทียบโหมด

| ฟีเจอร์ | โหมด Standalone | โหมด File System |
|---------|----------------|------------------|
| **ความสะดวกในการปรับใช้** | ✅ ไบนารีเดียว | ❌ ต้องใช้ไฟล์เพิ่มเติม |
| **ความเร็วในการรัน** | ✅ เร็ว | ✅ เร็ว |
| **การใช้หน่วยความจำ** | ⚠️ สูงเล็กน้อย | ✅ ต่ำ |
| **ความสะดวกในการพัฒนา** | ⚠️ ต้องบิลด์ใหม่ | ✅ สะท้อนแบบเรียลไทม์ |
| **สภาพแวดล้อมการผลิต** | ✅ แนะนำ | ❌ ไม่แนะนำ |
| **สภาพแวดล้อมการพัฒนา** | ✅ สามารถใช้ได้ | ✅ แนะนำ |

## 🔗 ลิงก์ที่เกี่ยวข้อง

- [เอกสารอย่างเป็นทางการของ n8n](https://docs.n8n.io/)
- [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go)
- [Claude Code](https://claude.ai/download)
- [Cursor](https://cursor.com) 