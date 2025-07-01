# n8n Workflow MCP Server

[![한국어](https://img.shields.io/badge/README-한국어-blue)](README.md) | [![English](https://img.shields.io/badge/README-English-red)](README.en.md) | [![日本語](https://img.shields.io/badge/README-日본語-green)](README.jp.md) | [![Tiếng Việt](https://img.shields.io/badge/README-Tiếng%20Việt-yellow)](README.vn.md) | [![繁體中文](https://img.shields.io/badge/README-繁體中文-purple)](README.tw.md) | [![ไทย](https://img.shields.io/badge/README-ไทย-orange)](README.th.md)

## 🎬 데모 비디오

워크플로우 검색 기능을 시연하는 비디오를 확인해보세요:

<video src="https://github.com/user-attachments/assets/1ec10869-be02-4ac7-a962-a6fdf865fffa" width="352" height="720"></video>

n8n 워크플로우 검증 및 관리를 위한 MCP (Model Context Protocol) 서버입니다.

## 🚀 주요 기능

### 🌍 다국어 지원
- **한국어 (ko)** - 기본 언어
- **영어 (en)** - English support
- **일본어 (jp)** - 日本語サポート
- **베트남어 (vn)** - Hỗ trợ tiếng Việt
- **대만어 (tw)** - 繁體中文支援
- **태국어 (th)** - การสนับสนุนภาษาไทย
- 모든 도구 설명과 파라미터 설명이 선택한 언어로 표시됩니다

### 📝 워크플로우 관리 도구
- **ping** - 서버 연결 상태 확인
- **validate_workflow** - JSON 파일의 n8n 워크플로우 문법 검증

### 🔍 검색 도구
- **search_n8n_nodes** - 키워드로 n8n 노드 검색
- **search_workflow**- 커뮤니티 n8n 워크플로우 템플릿 검색

## 🛠 설치 및 빌드

### 필요 조건
- Go 1.23 이상

### 빌드 방법

#### 🚀 Standalone 모드 (권장)
빌드 시점에 모든 n8n 노드 데이터를 바이너리에 포함하여 외부 파일 없이 실행 가능:

```bash
# 현재 플랫폼용 standalone 빌드
make build-standalone
# 또는
make build

# Linux 64bit standalone 빌드
make build-standalone-linux
# 또는
make build-linux

# Windows 64bit standalone 빌드
make build-standalone-windows
# 또는
make build-windows
```

#### 특정 플랫폼용 빌드

**macOS (Intel):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/n8n-mcp-server-darwin ./cmd/n8n-mcp-server
```

**macOS (Apple Silicon M1/M2):**
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/n8n-mcp-server-darwin-arm64 ./cmd/n8n-mcp-server
```

## 🎯 실행 방법

### 🚀 Standalone 모드 (권장)
외부 파일 없이 바이너리만으로 실행 (기본 모드):

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone

# Windows
build\n8n-mcp-server.exe -standalone

# Makefile 사용
make run
# 또는
make run-standalone
```

### 📁 파일 시스템 모드 (개발용)
외부 데이터 파일을 사용하는 방식:

```bash
# macOS/Linux
./build/n8n-mcp-server -standalone=false

# Windows
build\n8n-mcp-server.exe -standalone=false
```

### Go 명령어로 직접 실행
```bash
# Standalone 모드 (기본)
go run ./cmd/n8n-mcp-server -standalone

# 파일 시스템 모드
go run ./cmd/n8n-mcp-server -standalone=false
```

### 🌐 다국어 지원
서버는 한국어(ko), 영어(en), 일본어(jp), 베트남어(vn), 대만어(tw), 태국어(th)를 지원합니다. `--lang` 플래그로 언어를 설정할 수 있습니다:

```bash
# Standalone 모드
./build/n8n-mcp-server -standalone --lang ko  # 한국어 (기본값)
./build/n8n-mcp-server -standalone --lang en  # 영어
./build/n8n-mcp-server -standalone --lang jp  # 일본어
./build/n8n-mcp-server -standalone --lang vn  # 베트남어
./build/n8n-mcp-server -standalone --lang tw  # 대만어
./build/n8n-mcp-server -standalone --lang th  # 태국어

# 파일 시스템 모드
./build/n8n-mcp-server -standalone=false --lang ko  # 한국어 (기본값)
./build/n8n-mcp-server -standalone=false --lang en  # 영어
./build/n8n-mcp-server -standalone=false --lang jp  # 일본어
./build/n8n-mcp-server -standalone=false --lang vn  # 베트남어
./build/n8n-mcp-server -standalone=false --lang tw  # 대만어
./build/n8n-mcp-server -standalone=false --lang th  # 태국어
```

### ⚙️ 검색 가중치 조정
노드 검색 시 각 요소의 가중치를 조정할 수 있습니다:

```bash
# Standalone 모드에서 검색 가중치 조정
./build/n8n-mcp-server -standalone \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0

# 파일 시스템 모드에서 검색 가중치 조정 및 데이터 경로 변경
./build/n8n-mcp-server -standalone=false \
  --name-weight 3.0 \
  --overview-weight 2.0 \
  --case-weight 1.5 \
  --op-weight 1.0 \
  --data-path /custom/path/to/nodes
```

### 🆘 기타 옵션
```bash
# 도움말 보기
./build/n8n-mcp-server --help

# 버전 확인
./build/n8n-mcp-server --version
```

## 🔧 Claude Code & Cursor 설정

### macOS 설정
설정 파일 위치: `~/Library/Application Support/Claude/claude_desktop_config.json`

**Standalone 모드 (권장):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["--standalone", "--lang", "ko"]
    }
  }
}
```

**파일 시스템 모드 (개발용):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "/Users/{username}/{clone_path}/build/n8n-mcp-server",
      "args": ["-standalone=false", "--lang", "ko"]
    }
  }
}
```

**Go 실행 모드 (개발용):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "go",
      "args": ["run", "./cmd/n8n-mcp-server", "--standalone", "--lang", "ko"],
      "cwd": "/Users/{username}/{clone_path}"
    }
  }
}
```

### Windows 설정
설정 파일 위치: `%APPDATA%\Claude\claude_desktop_config.json`

**Standalone 모드 (권장):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["--standalone", "--lang", "ko"]
    }
  }
}
```

**파일 시스템 모드 (개발용):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "C:\\path\\to\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe",
      "args": ["-standalone=false", "--lang", "ko"]
    }
  }
}
```

**Go 실행 모드 (개발용):**
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "go",
      "args": ["run", ".\\cmd\\n8n-mcp-server", "--standalone", "--lang", "ko"],
      "cwd": "C:\\path\\to\\n8n-workflow-mcp"
    }
  }
}
```

## 📁 프로젝트 구조

```
n8n-workflow-mcp/
├── cmd/
│   └── n8n-mcp-server/           # 메인 애플리케이션 엔트리포인트
├── internal/
│   ├── config/                   # 설정 관리
│   ├── handlers/                 # MCP 요청 핸들러
│   ├── i18n/                     # 다국어 지원
│   ├── search/                   # 검색 엔진
│   ├── server/                   # MCP 서버 구현
│   ├── services/                 # 비즈니스 로직
│   ├── store/                    # 데이터 저장소
│   ├── validator/                # 워크플로우 검증
│   └── workflow/                 # 워크플로우 관리
├── pkg/
│   └── types/                    # 공통 타입 정의
├── examples/
│   └── config/                   # Claude Code & Cursor 설정 예제
├── test/                         # 테스트 파일
└── build/                        # 빌드 결과물
```

## 🧪 테스트

```bash
# 모든 테스트 실행
make test

# 커버리지와 함께 테스트 실행
make test-coverage
```

## 🛠 개발

```bash
# 코드 포맷팅
make fmt

# 린팅
make lint

# 모듈 정리
make mod-tidy

# 빌드 파일 정리
make clean
```

## 📊 모드 비교

| 기능 | Standalone 모드 | 파일 시스템 모드 |
|------|-----------------|-------------------|
| **배포 편의성** | ✅ 단일 바이너리 | ❌ 추가 파일 필요 |
| **실행 속도** | ✅ 빠름 | ✅ 빠름 |
| **메모리 사용량** | ⚠️ 약간 높음 | ✅ 낮음 |
| **개발 편의성** | ⚠️ 리빌드 필요 | ✅ 실시간 반영 |
| **운영 환경** | ✅ 권장 | ❌ 비권장 |
| **개발 환경** | ✅ 사용 가능 | ✅ 권장 |

## 🔗 관련 링크

- [n8n 공식 문서](https://docs.n8n.io/)
- [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go)
- [Claude Code](https://claude.ai/download)
- [Cursor](https://cursor.com)
