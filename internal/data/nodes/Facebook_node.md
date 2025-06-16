# Facebook Graph API

## Overview

Facebook Graph API 노드는 Facebook의 Graph API를 사용하여 Facebook 플랫폼과 상호작용하는 노드입니다. 사용자 정보 조회, 페이지 관리, 포스트 작성, 광고 관리 등 Facebook의 다양한 기능에 접근할 수 있습니다.

## Credentials

- **Credential Name**: `facebookGraphApi`
- **Required Fields**: 
  - Access Token

## Inputs

- **Main Input**: 처리할 데이터를 입력받습니다.

## Outputs

- **Main Output**: Facebook Graph API 응답 데이터를 출력합니다.

## Properties

### Resource: Facebook Graph API

#### Operation: API Request

| Field Name | Type | Description | Required |
|---|---|---|---|
| Host URL | Options | 요청 호스트 URL (Default/Video Uploads) | Yes |
| HTTP Request Method | Options | HTTP 메서드 (GET/POST/DELETE) | Yes |
| Graph API Version | Options | 사용할 Graph API 버전 | Yes |
| Node | String | 작업할 노드 (예: me, page_id) | Yes |
| Edge | String | 노드의 엣지 (예: posts, photos) | No |
| Fields | String | 반환받을 필드 목록 | No |
| Query Parameters | Collection | URL 쿼리 파라미터 | No |
| Headers | Collection | 추가 HTTP 헤더 | No |
| Body Parameters | Collection | 요청 본문 파라미터 (POST/DELETE) | No |

## UseCases

- **소셜 미디어 자동화** : 페이지에 자동으로 포스트 게시 및 콘텐츠 관리
- **사용자 데이터 분석** : 사용자 프로필 정보 및 인사이트 데이터 수집
- **광고 캠페인 관리** : Facebook 광고 생성, 수정, 성과 분석 자동화
- **페이지 관리** : 페이지 정보 업데이트, 메시지 응답, 리뷰 관리
- **이벤트 관리** : Facebook 이벤트 생성, 업데이트, 참석자 관리
- **콘텐츠 모니터링** : 페이지 댓글, 메시지, 멘션 모니터링 및 자동 응답

## Configuration Options

### Host URL Options:
- **Default**: graph.facebook.com (대부분의 요청)
- **Video Uploads**: graph-video.facebook.com (비디오 업로드 전용)

### Supported API Versions:
- v22.0, v21.0, v20.0, v19.0, v18.0, v17.0, v16.0, v15.0
- v14.0, v13.0, v12.0, v11.0, v10.0, v9.0, v8.0, v7.0
- v6.0, v5.0, v4.0, v3.3, v3.2, v3.1, v3.0

## Related Nodes

- **HTTP Request**: For custom Facebook Graph API calls
- **JSON**: For processing Facebook API responses
- **Code**: For custom data transformation logic
- **Schedule Trigger**: For automated Facebook operations 