# E2E Test Node

## Overview

Dummy node used for e2e testing

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Test

#### Operation: Remote Options
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Remote Options | options | Remote options to load with dynamic loading | Yes | [] |
| Field ID | string | ID for field dependency | No | - |
| Other Non Important Field | string | Additional field for testing | No | - |

#### Operation: Resource Locator
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resource Locator | resourceLocator | Test resource locator with multiple modes | Yes | {mode: 'list', value: ''} |

**Resource Locator Modes:**
- **From List**: Searchable list mode with pagination
- **By URL**: URL format with validation (https://example.com/user/[uuid])
- **ID**: Direct UUID input with validation

#### Operation: Resource Mapping
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resource Mapping Component | resourceMapper | Test resource mapper with column mapping | Yes | {mappingMode: 'defineBelow', value: null} |

## UseCases

- **E2E Testing**: End-to-end testing of n8n UI components
- **Component Testing**: Test resource locators and mapping components
- **Validation Testing**: Test URL and UUID validation patterns
- **Dynamic Loading**: Test remote option loading functionality
- **Development Testing**: Mock node for development scenarios

