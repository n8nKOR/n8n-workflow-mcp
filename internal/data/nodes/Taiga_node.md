# Taiga Node

## Overview

Consume Taiga API

## Credentials

- Name: taigaApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: epic

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to which the epic belongs. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Subject | string |  | Yes |  |
| Additional Fields | collection | ID of the user to assign the epic to. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Epic ID | string | ID of the epic to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Epic ID | string | ID of the epic to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to which the epic belongs. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | ID of the user whom the epic is assigned to. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to set the epic to. Choose from the list, or specify an ID using an <a href= | No |  |
| Epic ID | string | ID of the epic to update | Yes |  |
| Update Fields | collection | ID of the user to whom the epic is assigned. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: issue

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to which the issue belongs. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Subject | string |  | Yes |  |
| Additional Fields | collection | ID of the user to whom the issue is assigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue ID | string | ID of the issue to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue ID | string | ID of the issue to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to which the issue belongs. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | ID of the user to assign the issue to. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to set the issue to. Choose from the list, or specify an ID using an <a href= | No |  |
| Issue ID | string | ID of the issue to update | Yes |  |
| Update Fields | collection | ID of the user whom the issue is assigned to. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to which the task belongs. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Subject | string |  | Yes |  |
| Additional Fields | collection | ID of the user to whom the task is assigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of the task to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of the task to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to which the task belongs. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | ID of the user whom the task is assigned to. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to set the task to. Choose from the list, or specify an ID using an <a href= | No |  |
| Task ID | string | ID of the task to update | Yes |  |
| Update Fields | collection | ID of the user to assign the task to. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: userStory

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to which the user story belongs. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Subject | string |  | Yes |  |
| Additional Fields | collection | ID of the user to whom the user story is assigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User Story ID | string | ID of the user story to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User Story ID | string | ID of the user story to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to which the user story belongs. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | ID of the user whom the user story is assigned to. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | ID of the project to set the user story to. Choose from the list, or specify an ID using an <a href= | No |  |
| User Story ID | string | ID of the user story to update | Yes |  |
| Update Fields | collection | ID of the user to assign the the user story to. Choose from the list, or specify an ID using an <a href= | No |  |

## UseCases

- [Agile Project Management] : Manage agile software development projects with Scrum and Kanban
- [Sprint Planning] : Plan and organize development sprints with user stories and tasks
- [User Story Management] : Create and manage user stories with acceptance criteria
- [Task Tracking] : Track development tasks and monitor progress across team members
- [Bug Tracking] : Report and track software bugs through resolution
- [Team Collaboration] : Enable development team collaboration and communication
- [Backlog Management] : Maintain and prioritize product backlog items
- [Progress Monitoring] : Monitor project progress with burndown charts and metrics
- [Issue Management] : Manage project issues and impediments systematically
- [Release Planning] : Plan software releases and track milestone achievements
- [Code Integration] : Integrate with version control systems for development tracking
- [Client Communication] : Communicate project status and updates to clients
- [Resource Management] : Manage team resources and capacity planning
- [Quality Assurance] : Track quality assurance activities and testing progress
- [Project Reporting] : Generate project reports and performance analytics

