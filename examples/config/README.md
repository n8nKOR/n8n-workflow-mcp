# Claude Desktop 설정 예제

이 디렉토리에는 Claude Desktop 앱에서 n8n-workflow-mcp 서버를 사용하기 위한 설정 파일 예제들이 포함되어 있습니다.

## 📁 파일 목록

### 🚀 Standalone 모드 (권장)

#### 1. `claude_desktop_config_standalone.json`
- **용도**: Standalone 모드로 빌드된 바이너리 사용 (권장)
- **장점**: 
  - 외부 파일 없이 단일 바이너리로 실행
  - 빠른 시작 시간
  - 배포 간편
- **적용 대상**: 일반 사용자, 운영 환경

#### 2. `claude_desktop_config_standalone_en.json`
- **용도**: 영어 인터페이스로 Standalone 모드 실행
- **특징**: 모든 도구 설명이 영어로 표시
- **적용 대상**: 영어권 사용자

#### 3. `claude_desktop_config_standalone_jp.json`
- **용도**: 일본어 인터페이스로 Standalone 모드 실행
- **특징**: 모든 도구 설명이 일본어로 표시
- **적용 대상**: 일본어 사용자

### 🛠 개발 모드

#### 4. `claude_desktop_config_dev_filesystem.json`
- **용도**: 파일 시스템 모드로 개발용 실행
- **장점**: 
  - 소스 코드 변경 시 데이터 파일 수정 즉시 반영
  - 개발 중 디버깅 편리
- **단점**: 외부 데이터 파일 의존성
- **적용 대상**: 개발자

#### 5. `claude_desktop_config_dev_go.json`
- **용도**: Go 명령어로 직접 실행 (개발용)
- **장점**: 
  - 소스 코드 변경 시 자동 컴파일
  - 개발 중 실시간 반영
- **단점**: Go 환경 필요, 실행 시간 증가
- **적용 대상**: 개발자

### 🪟 Windows 전용

#### 6. `claude_desktop_config_windows_standalone.json`
- **용도**: Windows에서 Standalone 모드 실행
- **특징**: Windows 경로 형식 사용
- **적용 대상**: Windows 사용자

#### 7. `claude_desktop_config_windows_dev.json`
- **용도**: Windows에서 Go 명령어로 개발용 실행
- **특징**: Windows 경로 형식과 개발 옵션
- **적용 대상**: Windows 개발자

## 🎯 추천 설정

### 일반 사용자
1. **macOS/Linux**: `claude_desktop_config_standalone.json`
2. **Windows**: `claude_desktop_config_windows_standalone.json`

### 개발자
1. **개발 시**: `claude_desktop_config_dev_go.json`
2. **테스트 시**: `claude_desktop_config_dev_filesystem.json`

### 다국어 사용자
1. **영어**: `claude_desktop_config_standalone_en.json`
2. **일본어**: `claude_desktop_config_standalone_jp.json`
3. **한국어**: `claude_desktop_config_standalone.json` (기본)

## 🔧 설정 방법

### 1. 설정 파일 위치 확인
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

### 2. 설정 파일 선택 및 적용
1. 위의 파일 목록에서 용도에 맞는 설정 파일을 선택합니다.
2. 선택한 파일의 내용을 복사합니다.
3. Claude Desktop 설정 파일에 붙여넣습니다.
4. 파일 경로를 본인의 환경에 맞게 수정합니다.

### 3. 경로 수정 가이드

**macOS/Linux 예시:**
```json
"command": "/Users/john/projects/n8n-workflow-mcp/build/n8n-mcp-server"
```

**Windows 예시:**
```json
"command": "C:\\Users\\john\\projects\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe"
```

### 4. Claude Desktop 재시작
설정을 적용한 후 Claude Desktop 앱을 완전히 종료한 다음 다시 시작합니다.

## ⚙️ 고급 설정 옵션

### 검색 가중치 조정
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/path/to/n8n-mcp-server",
      "args": [
        "--standalone",
        "--lang", "ko",
        "--name-weight", "3.0",
        "--overview-weight", "2.0",
        "--case-weight", "1.5",
        "--op-weight", "1.0"
      ]
    }
  }
}
```

### 커스텀 데이터 경로 (파일 시스템 모드)
```json
{
  "mcpServers": {
    "n8n-workflow": {
      "command": "/path/to/n8n-mcp-server",
      "args": [
        "-standalone=false",
        "--data-path", "/custom/path/to/nodes",
        "--lang", "ko"
      ]
    }
  }
}
```

## 🔍 문제 해결

### 일반적인 문제들

1. **서버가 시작되지 않는 경우**
   - 파일 경로가 올바른지 확인
   - 실행 권한이 있는지 확인 (`chmod +x`)
   - 빌드된 바이너리가 존재하는지 확인

2. **권한 오류가 발생하는 경우**
   ```bash
   chmod +x /path/to/n8n-mcp-server
   ```

3. **경로 관련 오류**
   - 절대 경로 사용 권장
   - Windows에서는 백슬래시 이스케이프 필요: `\\`

### 로그 확인
Claude Desktop의 로그를 확인하려면:
- **macOS**: Console.app에서 "Claude" 검색
- **Windows**: 이벤트 뷰어 또는 Claude 로그 파일 확인

## 📝 주의사항

- 모든 경로는 절대 경로를 사용해야 합니다
- Windows에서 경로에 백슬래시(`\`)를 사용할 때는 이중으로 입력하세요: `\\`
- 또는 슬래시(`/`)를 사용할 수도 있습니다
- 설정 변경 후에는 반드시 Claude Desktop을 재시작하세요
- 개발 모드는 Go 환경이 설치되어 있어야 합니다

## 🆘 도움말

추가적인 도움이 필요하시면:
1. 프로젝트의 메인 README.md 참고
2. GitHub Issues에 문의
3. 설정 파일 문법 검증: [JSON Validator](https://jsonlint.com/) 