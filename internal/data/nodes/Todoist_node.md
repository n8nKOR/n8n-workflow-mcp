# Todoist Node

## Overview

Consume Todoist API

## Credentials

- Name: todoistApi (API Key), Required: Yes
- Name: todoistOAuth2Api (OAuth2), Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Task

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|
| authentication | options | Choose authentication method | Yes |
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| content | string | Task content | Yes |
| project | resourceLocator | The destination project | Yes |
| labels | multiOptions | Optional labels for the task | No |
| description | string | Task description | No |
| dueDateTime | dateTime | Specific due date in RFC3339 format | No |
| dueLang | string | Language code for due string | No |
| dueString | string | Human-readable due date | No |
| parentId | options | Parent task | No |
| priority | number | Task priority from 1 (normal) to 4 (urgent) | No |
| section | options | Task section | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|
| authentication | options | Choose authentication method | Yes |
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| taskId | string | ID of the task to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|
| authentication | options | Choose authentication method | Yes |
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| taskId | string | ID of the task to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|
| authentication | options | Choose authentication method | Yes |
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| returnAll | boolean | Whether to return all results or only up to limit | No |
| limit | number | Maximum number of results to return (1-500) | No |
| filter | string | Filter by supported filter syntax | No |
| ids | string | Comma-separated list of task IDs | No |
| labelId | options | Filter by label | No |
| lang | string | Language tag for filter | No |
| parentId | options | Filter by parent task | No |
| projectId | options | Filter by project | No |
| sectionId | options | Filter by section | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|
| authentication | options | Choose authentication method | Yes |
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| taskId | string | ID of the task to update | Yes |
| content | string | Task content | No |
| description | string | Task description | No |
| dueDateTime | dateTime | Due date and time | No |
| dueLang | string | Language code | No |
| dueString | string | Human-readable due date | No |
| labels | multiOptions | Task labels | No |
| priority | number | Task priority (1-4) | No |

#### Operation: Close

| Field Name | Type | Description | Required |
|---|---|---|
| authentication | options | Choose authentication method | Yes |
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| taskId | string | ID of the task to close | Yes |

#### Operation: Reopen

| Field Name | Type | Description | Required |
|---|---|---|
| authentication | options | Choose authentication method | Yes |
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| taskId | string | ID of the task to reopen | Yes |

#### Operation: Move

| Field Name | Type | Description | Required |
|---|---|---|
| authentication | options | Choose authentication method | Yes |
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| taskId | string | ID of the task to move | Yes |
| project | resourceLocator | The destination project | Yes |
| section | options | Destination section (v2.1+) | No |
| parent | options | Destination parent task (v2.1+) | No |

## UseCases

- [Task Management]: Create, update, and organize tasks in Todoist projects with due dates, priorities, and labels for personal and team productivity
- [Project Planning]: Manage project tasks with hierarchical structure using parent-child relationships and section organization
- [Workflow Integration]: Automatically create tasks from external triggers like emails, forms, or calendar events to streamline task capture

