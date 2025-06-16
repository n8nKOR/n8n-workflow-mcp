# Execute Sub-workflow

## Overview

Execute Sub-workflow 노드는 다른 워크플로우를 실행하는 노드입니다. 복잡한 워크플로우를 모듈화하여 재사용 가능한 서브 워크플로우로 분리하고, 메인 워크플로우에서 필요에 따라 호출할 수 있습니다. 이를 통해 워크플로우의 구조를 단순화하고 유지보수성을 향상시킬 수 있습니다.

## Credentials

이 노드는 별도의 자격 증명이 필요하지 않습니다.

## Inputs

- **Main Input**: 서브 워크플로우로 전달할 데이터를 입력받습니다.

## Outputs

- **Main Output**: 서브 워크플로우의 실행 결과를 출력합니다.

## Properties

### Resource: Workflow Execution

#### Operation: Execute Sub-workflow

| Field Name | Type | Description | Required |
|---|---|---|---|
| Source | Options | How to specify the workflow to execute | Yes |
| Workflow ID | String | ID of the workflow to execute | Yes |
| Options | Collection | Additional configuration options | No |

#### Options Collection

| Field Name | Type | Description | Required |
|---|---|---|---|
| Wait for Sub-Workflow | Boolean | Wait for sub-workflow completion | No |
| Load Workflow | Boolean | Load workflow from database | No |
## UseCases

- **모듈화된 워크플로우**: 공통 기능을 서브 워크플로우로 분리하여 여러 메인 워크플로우에서 재사용
- **복잡한 로직 분할**: 복잡한 비즈니스 로직을 단계별로 분할하여 관리 용이성 향상
- **조건부 실행**: 특정 조건에 따라 다른 서브 워크플로우를 실행하여 동적 처리
- **병렬 처리**: 여러 서브 워크플로우를 병렬로 실행하여 처리 성능 향상
- **테스트 및 디버깅**: 개별 기능을 서브 워크플로우로 분리하여 독립적인 테스트 가능
- **팀 협업**: 팀원별로 서브 워크플로우를 개발하고 메인 워크플로우에서 통합 