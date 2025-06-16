# TheHive 5 Node

## Overview

Consume TheHive 5 API

## Credentials

- Name: theHiveProjectApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Alert

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| title | string | Alert title | Yes |
| description | string | Alert description | Yes |
| severity | options | Alert severity level | Yes |
| type | string | Alert type | Yes |
| source | string | Alert source | Yes |
| tags | multiOptions | Labels for categorization | No |
| tlp | options | Traffic Light Protocol settings | No |
| customFields | collection | Resource-specific attributes | No |
| attachments | collection | File uploads | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| alertId | string | ID of the alert to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| alertId | string | ID of the alert to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| returnAll | boolean | Return all results or limit | No |
| limit | number | Maximum number of results | No |
| filters | collection | Filter options | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| alertId | string | ID of the alert to update | Yes |
| updateFields | collection | Fields to update | No |

### Resource: Case

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| title | string | Case title | Yes |
| description | string | Case description | Yes |
| severity | options | Case severity level | Yes |
| tags | multiOptions | Labels for categorization | No |
| tlp | options | Traffic Light Protocol settings | No |
| assignee | options | User assignments | No |
| customFields | collection | Resource-specific attributes | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| caseId | string | ID of the case to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| caseId | string | ID of the case to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| returnAll | boolean | Return all results or limit | No |
| limit | number | Maximum number of results | No |
| filters | collection | Filter options | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| caseId | string | ID of the case to update | Yes |
| updateFields | collection | Fields to update | No |

### Resource: Task

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| caseId | string | ID of the parent case | Yes |
| title | string | Task title | Yes |
| description | string | Task description | No |
| assignee | options | User assignments | No |
| status | options | Task status | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| taskId | string | ID of the task to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| taskId | string | ID of the task to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| returnAll | boolean | Return all results or limit | No |
| limit | number | Maximum number of results | No |
| filters | collection | Filter options | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| taskId | string | ID of the task to update | Yes |
| updateFields | collection | Fields to update | No |

### Resource: Observable

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| caseId | string | ID of the parent case | Yes |
| dataType | options | Type of observable data | Yes |
| data | string | Observable data value | Yes |
| message | string | Observable message | No |
| tags | multiOptions | Labels for categorization | No |
| tlp | options | Traffic Light Protocol settings | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| observableId | string | ID of the observable to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| observableId | string | ID of the observable to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| returnAll | boolean | Return all results or limit | No |
| limit | number | Maximum number of results | No |
| filters | collection | Filter options | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| observableId | string | ID of the observable to update | Yes |
| updateFields | collection | Fields to update | No |

### Resource: Query

#### Operation: Execute

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| query | string | Query to execute against TheHive database | Yes |
| range | string | Query range specification | No |
| sort | collection | Sort options | No |

## UseCases

- [Security Alert Management]: Create and manage security alerts with various severity levels, track investigation progress, and promote alerts to full investigation cases
- [Investigation Case Tracking]: Manage investigation cases with observables, tasks, and assignments to track complex security incidents from detection to resolution
- [Threat Intelligence Integration]: Create observables from threat intelligence feeds and execute queries to correlate indicators across multiple cases and alerts

