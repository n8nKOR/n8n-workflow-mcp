# Wekan Node

## Overview

Consume Wekan API

## Credentials

- Name: wekanApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: board

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | The title of the board | Yes |  |
| Owner Name or ID | options | The user ID in Wekan. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Additional Fields | collection | Whether to set the board active | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board to get | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User Name or ID | options | The ID of the user that boards are attached. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: card

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that list belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list to create card in. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Title | string | The title of the card | Yes |  |
| Swimlane Name or ID | options | The swimlane ID of the new card. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Author Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Additional Fields | collection | The new list of assignee IDs attached to the card. Choose from the list, or specify IDs using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that list belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that list belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card ID | string | The ID of the card to get | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that list belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| From Object | options |  | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Swimlane Name or ID | options | The ID of the swimlane that card belongs to. Choose from the list, or specify an ID using an <a href= | No |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that list belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card to update. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Update Fields | collection | Update the owner of the card. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: cardComment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Author Name or ID | options | The user who posted the comment. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Comment | string | The comment text | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Comment Name or ID | options | The ID of the comment to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board that card belongs to | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Comment ID | string | The ID of the comment to get | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: checklist

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board where the card is in. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card to add checklist to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Title | string | The title of the checklist to add | Yes |  |
| Items | string | Items to be added to the checklist | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card that checklist belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Checklist Name or ID | options | The ID of the checklist to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card that checklist belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Checklist Name or ID | options | The ID of the checklist to get. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that list belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card to get checklists. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: checklistItem

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card that checklistItem belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Checklist Name or ID | options | The ID of the checklistItem that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Checklist Item Name or ID | options | The ID of the checklistItem item to get. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card that checklistItem belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Checklist ID | string | The ID of the checklistItem that card belongs to | Yes |  |
| Checklist Item Name or ID | options | The ID of the checklistItem item to get. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Card Name or ID | options | The ID of the card that checklistItem belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| CheckList Name or ID | options | The ID of the checklistItem that card belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Checklist Item Name or ID | options | The ID of the checklistItem item to update. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Update Fields | collection | The new title for the checklistItem item | No |  |

### Resource: list

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board the list should be created in. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Title | string | The title of the list | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that list belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | The ID of the list to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The ID of the board that list belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| List ID | string | The ID of the list to get | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | ID of the board where the lists are in. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **Project Management** : Manage projects using Kanban boards, lists, and task cards
- **Task Tracking** : Track task progress, assignments, and completion status
- **Team Collaboration** : Facilitate team collaboration through card comments and updates
- **Agile Development** : Implement agile development workflows using Wekan boards
- **Workflow Automation** : Automate workflow processes by integrating Wekan with other systems
- **Sprint Planning** : Organize and manage sprints using Wekan's board structure
- **Bug Tracking** : Track software bugs and issues using cards and checklists
- **Content Planning** : Plan and organize content creation using Wekan boards
- **Event Management** : Organize events and track preparation tasks
- **Personal Productivity** : Manage personal tasks and projects using Kanban methodology
- **Resource Management** : Track resource allocation and team capacity
- **Quality Assurance** : Manage QA processes and testing workflows
- **Customer Support** : Track support tickets and customer issues
- **Documentation Management** : Organize documentation projects and writing tasks
- **Release Management** : Plan and track software releases and deployments
- **Marketing Campaigns** : Manage marketing campaigns and promotional activities
- **Educational Projects** : Organize educational content and course development
- **Research Projects** : Track research tasks and academic projects
- **Construction Projects** : Manage construction projects and task dependencies
- **Manufacturing Workflows** : Organize manufacturing processes and quality control

