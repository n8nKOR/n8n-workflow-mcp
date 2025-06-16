# Trello Node

## Overview

Consume Trello API to manage boards, lists, cards, and team collaboration features for project management and task organization workflows.

## Credentials

- Name: trelloApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: attachment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Source URL | string | The URL of the attachment to add | Yes |  |
| Additional Fields | collection | The MIME type of the attachment to add | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attachment ID | string | The ID of the attachment to delete | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attachment ID | string | The ID of the attachment to get | Yes |  |
| Additional Fields | collection | Fields to return. Either  | No |  |

### Resource: board

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The name of the board | Yes |  |
| Description | string | The description of the board | No |  |
| Additional Fields | collection | Determines the type of card aging that should take place on the board if card aging is enabled | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Fields | collection | Whether the board is closed | No |  |

### Resource: boardMember

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board to get members from | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board to add member to | Yes |  |
| Member ID | string | The ID of the member to add to the board | Yes |  |
| Type | options | Invite as normal member | Yes |  |
| Additional Fields | collection | Whether to allow organization admins to add multi-board guests onto a board | No |  |

#### Operation: invite

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board to invite member to | Yes |  |
| Email | string | The ID of the board to update | Yes |  |
| Additional Fields | collection | Invite as normal member | No |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board to remove member from | Yes |  |
| Member ID | string | The ID of the member to remove from the board | Yes |  |

### Resource: card

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | The ID of the list to create card in | Yes |  |
| Name | string | The name of the card | Yes |  |
| Description | string | The description of the card | No |  |
| Additional Fields | collection | A due date for the card | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Fields | collection | The ID of the image attachment the card should use as its cover, or null for none | No |  |

### Resource: cardComment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Text | string | Text of the comment | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comment ID | string | The ID of the comment to delete | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comment ID | string | The ID of the comment to delete | Yes |  |
| Text | string | Text of the comment | Yes |  |

### Resource: checklist

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The URL of the checklist to add | Yes |  |
| Additional Fields | collection | The ID of a source checklist to copy into the new one | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checklist ID | string | The ID of the checklist to delete | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checklist ID | string | The ID of the checklist to get | Yes |  |
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: createCheckItem

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checklist ID | string | The ID of the checklist to update | Yes |  |
| Name | string | The name of the new check item on the checklist | Yes |  |
| Additional Fields | collection | Whether the check item is already checked when created | No |  |

#### Operation: deleteCheckItem

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| CheckItem ID | string | The ID of the checklist item to delete | Yes |  |

#### Operation: getCheckItem

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| CheckItem ID | string | The ID of the checklist item to get | Yes |  |
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: updateCheckItem

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| CheckItem ID | string | The ID of the checklist item to update | Yes |  |
| Additional Fields | collection | The new name for the checklist item | No |  |

#### Operation: completedCheckItems

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Fields to return. Either  | No |  |

### Resource: label

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name for the label | Yes |  |
| Color | options | The color for the label | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Label ID | string | The ID of the label to delete | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Label ID | string | Get information about a label by ID | Yes |  |
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: addLabel

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Label ID | string | The ID of the label to add | Yes |  |

#### Operation: removeLabel

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Label ID | string | The ID of the label to remove | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Label ID | string | The ID of the label to update | Yes |  |
| Update Fields | collection | Name of the label | No |  |

### Resource: list

#### Operation: archive

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | The ID of the list to archive or unarchive | Yes |  |
| Archive | boolean | Whether the list should be archived or unarchived | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board the list should be created in | Yes |  |
| Name | string | The name of the list | Yes |  |
| Additional Fields | collection | ID of the list to copy into the new list | No |  |

#### Operation: getCards

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | The ID of the list to get cards | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | The ID of the list to get | Yes |  |
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board ID | string | The ID of the board | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Fields to return. Either  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | The ID of the list to update | Yes |  |
| Update Fields | collection | ID of a board the list should be moved to | No |  |

## UseCases

- **Project Management Automation**: Automate project workflows by creating boards, lists, and cards based on project requirements and team assignments
- **Task Management Integration**: Sync tasks from external systems to Trello boards, automatically creating cards with due dates, assignments, and priorities
- **Team Collaboration Enhancement**: Automate team member assignments, board sharing, and collaboration workflows for distributed team management
- **Bug Tracking and Issue Management**: Create automated bug tracking workflows with card creation, labeling, and assignment based on issue reports
- **Content Planning and Editorial Calendars**: Manage content creation workflows with automated card creation for editorial calendars and publication schedules
- **Customer Support Ticket Management**: Convert customer support requests into Trello cards with automatic assignment and progress tracking
- **Event and Campaign Management**: Organize events and marketing campaigns with automated board setup, task assignment, and milestone tracking
- **Client Project Coordination**: Manage client projects with automated board creation, progress updates, and client communication workflows
- **Personal Productivity Automation**: Create personal task management systems with automated card creation from various input sources
- **Workflow Status Reporting**: Generate project status reports and notifications based on Trello board activities and card progress

