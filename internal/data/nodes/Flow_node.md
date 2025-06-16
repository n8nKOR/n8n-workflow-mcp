# Flow Node

## Overview

Consume Flow API

## Credentials

- Name: flowApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace ID | string | Create resources under the given workspace | Yes |  |
| Name | string | The title of the task | Yes |  |
| Additional Fields | collection | The ID of the account to whom this task will be assigned | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace ID | string | Create resources under the given workspace | Yes |  |
| Task ID | string |  | Yes |  |
| Update Fields | collection | The title of the task | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |
| Filters | collection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Create resources under the given workspace | No |  |

## UseCases

- **Project Task Management**: Create, update, and track tasks across different projects and workspaces
- **Team Collaboration**: Automate task assignments and updates for team coordination and project management
- **Workflow Automation**: Integrate task management with other business processes for streamlined operations
- **Resource Planning**: Manage task assignments and workload distribution across team members
- **Project Tracking**: Monitor project progress through automated task status updates and reporting
- **Client Project Management**: Organize client work into structured tasks and deliverables
- **Agile Development**: Support agile methodologies with automated sprint planning and task management
- **Deadline Management**: Automate task creation and reminders for time-sensitive projects
- **Cross-platform Integration**: Sync tasks with other project management and productivity tools
- **Performance Analytics**: Track team productivity and project completion metrics through task data

