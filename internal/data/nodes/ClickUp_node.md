# ClickUp Node

## Overview

Consume ClickUp API (Beta) for comprehensive project management and task organization. Manage tasks, lists, folders, time tracking, goals, and team collaboration within ClickUp workspaces.

## Credentials

- Name: clickUpApi, Required: Yes
- Authentication Methods:
  - Access Token: Personal or app-specific access token
  - OAuth2: Full OAuth2 flow with authorization code grant type
- API Base URL: https://api.clickup.com/api/v2/

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: checklist

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |
| Name | string |  | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checklist ID | string |  | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checklist ID | string |  | Yes |  |
| Update Fields | collection |  | No |  |

### Resource: checklistItem

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checklist ID | string |  | Yes |  |
| Name | string |  | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checklist ID | string |  | Yes |  |
| Checklist Item ID | string |  | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checklist ID | string |  | Yes |  |
| Checklist Item ID | string |  | Yes |  |
| Update Fields | collection | Checklist item that you want to nest the target checklist item underneath | No |  |

### Resource: comment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comment On | options |  | No |  |
| ID | string |  | Yes |  |
| Comment Text | string |  | No |  |
| Additional Fields | collection | Whether creation notifications will be sent to everyone including the creator of the comment | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comment ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comments On | options |  | No |  |
| ID | string |  | Yes |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comment ID | string |  | Yes |  |
| Update Fields | collection |  | No |  |

### Resource: folder

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string |  | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Update Fields | collection |  | No |  |

### Resource: goal

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string |  | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Goal ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Goal ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Goal ID | string |  | Yes |  |
| Update Fields | collection |  | No |  |

## UseCases

- **Project Management Automation** : Automate task creation, assignment, and status updates for team projects
- **Task Tracking and Reporting** : Monitor project progress and generate status reports from ClickUp data
- **Team Collaboration** : Enhance team collaboration through automated notifications and task assignments
- **Goal Management** : Track and manage team goals and key results through automated workflows
- **Time Tracking Integration** : Integrate time tracking data with project management and billing systems
- **Client Project Management** : Manage client projects, deliverables, and communication through ClickUp
- **Resource Planning** : Optimize team resource allocation based on project requirements and availability
- **Quality Assurance** : Automate QA processes with checklist management and testing workflows
- **Sprint Management** : Manage agile development sprints and track sprint progress
- **Cross-Platform Integration** : Integrate ClickUp with other business tools and platforms for seamless workflows

## UseCases

- **Project Management** : Manage complex projects with tasks, subtasks, and dependencies
- **Team Collaboration** : Facilitate team collaboration through comments, mentions, and shared workspaces
- **Agile Development** : Implement agile methodologies with sprints, epics, and story tracking
- **Task Automation** : Automate task creation, updates, and status changes based on triggers
- **Time Tracking** : Track time spent on tasks and projects for billing and productivity analysis
- **Goal Management** : Set and track organizational goals with key results and progress metrics
- **Resource Planning** : Plan and allocate resources across teams and projects
- **Client Management** : Manage client projects with dedicated spaces and permission controls
- **Bug Tracking** : Track and manage software bugs with priority levels and assignments
- **Content Planning** : Plan content creation workflows with editorial calendars and review processes
- **Marketing Campaigns** : Organize marketing campaigns with tasks, timelines, and deliverables
- **Product Development** : Manage product development cycles from ideation to release
- **Quality Assurance** : Organize QA processes with test cases and bug reporting
- **Event Planning** : Plan events with detailed task breakdowns and timeline management
- **Sales Pipeline** : Track sales opportunities and customer relationship management
- **HR Processes** : Manage HR workflows including onboarding and performance reviews
- **Financial Planning** : Track budgets, expenses, and financial project milestones
- **Compliance Management** : Ensure compliance with regulatory requirements through structured workflows
- **Vendor Management** : Manage vendor relationships and contract deliverables
- **Educational Programs** : Organize educational content development and course management

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Key Result ID | string |  | Yes |  |
| Update Fields | collection | Only matters for type Number and Currency. For Currency the unit must be a valid currency code. | No |  |

### Resource: guest

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Email | string |  | No |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Guest ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Guest ID | string |  | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Guest ID | string |  | Yes |  |
| Update Fields | collection |  | No |  |

### Resource: spaceTag

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| New Name | string | New name to set for the tag | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string | The first name on the task | Yes |  |
| Additional Fields | collection | Choose from the list, or specify IDs using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |
| Update Fields | collection | Assignees IDs. Multiple ca be added separated by comma. | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Choose from the list, or specify IDs using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |

#### Operation: member

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: setCustomField

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | The ID of the task to add custom field to | Yes |  |
| Field ID | string | The ID of the field to add custom field to | Yes |  |
| Value Is JSON | boolean | The value is JSON and will be parsed as such. Is needed if for example needed for labels which expects the value to be an array. | No |  |
| Value | string | The value to set on custom field | Yes |  |

### Resource: taskDependency

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |
| Depends On Task ID | string |  | Yes |  |
| Task ID | string |  | Yes |  |
| Depends On Task ID | string |  | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |
| Depends On Task ID | string |  | Yes |  |
| Task ID | string |  | Yes |  |
| Depends On Task ID | string |  | Yes |  |

### Resource: timeEntry

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Running | boolean | Whether to return just the current running time entry | No |  |
| Time Entry ID | string |  | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Start | dateTime |  | Yes |  |
| Duration (Minutes) | number | Duration in minutes | Yes |  |
| Task Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: start

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Task ID | string |  | Yes |  |
| Additional Fields | collection | Description of the time entry | No |  |

#### Operation: stop

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Time Entry ID | string |  | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Archived | boolean |  | Yes |  |
| Time Entry ID | string |  | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: timeEntryTag

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Time Entry IDs | string |  | Yes |  |
| Tags | fixedCollection |  | No |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Time Entry IDs | string |  | Yes |  |
| Tag Names or IDs | multiOptions | Choose from the list, or specify IDs using an <a href= | Yes |  |

### Resource: list

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string |  | Yes |  |
| Additional Fields | collection | Integer mapping as 1 : Urgent, 2 : High, 3 : Normal, 4 : Low | No |  |

#### Operation: member

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | Task ID | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: customFields

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Space Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Folderless List | boolean |  | Yes |  |
| Folder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| List ID | string |  | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

