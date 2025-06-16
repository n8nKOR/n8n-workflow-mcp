# Monday.com Node

## Overview

Consume Monday.com API

## Credentials

- Name: mondayComApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: board

#### Operation: archive

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Board unique identifiers. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The board | Yes |  |
| Kind | options | The board | Yes |  |
| Additional Fields | collection | Optional board template ID | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Board unique identifiers. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: boardColumn

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Title | string |  | Yes |  |
| Column Type | options |  | Yes |  |
| Additional Fields | collection | The new column | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

### Resource: boardGroup

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string | The group name | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Group Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

### Resource: boardItem

#### Operation: addUpdate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Item ID | string | The unique identifier of the item to add update to | Yes |  |
| Update Text | string | The update text to add | Yes |  |

#### Operation: changeColumnValue

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The unique identifier of the board. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Item ID | string | The unique identifier of the item to change column of | Yes |  |
| Column Name or ID | options | The column\ | Yes |  |
| Value | json | The column value in JSON format. Documentation can be found <a href= | Yes |  |

#### Operation: changeMultipleColumnValues

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The unique identifier of the board. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Item ID | string | Item | Yes |  |
| Column Values | json | The column fields and values in JSON format. Documentation can be found <a href= | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Group Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string | The new item | Yes |  |
| Additional Fields | collection | The column values of the new item | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Item ID | string | Item | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Item ID | string | Item | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Group Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: getByColumnValue

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | The unique identifier of the board. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Column Name or ID | options | The column\ | Yes |  |
| Column Value | string | The column value to search items by | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: move

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Board Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Item ID | string | The item | Yes |  |
| Group Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

## UseCases

- **Project Management** : Manage projects, tasks, and team collaboration using Monday.com boards
- **Task Tracking** : Track task progress, assignments, and deadlines across teams
- **Team Collaboration** : Facilitate team communication and collaboration through board updates
- **Workflow Automation** : Automate workflow processes by integrating Monday.com with other systems
- **Resource Management** : Manage resources, assignments, and capacity planning
- **Client Management** : Track client projects, communications, and deliverables
- **Sales Pipeline** : Manage sales opportunities, leads, and customer relationships
- **Marketing Campaigns** : Track marketing campaigns, activities, and performance metrics
- **Bug Tracking** : Manage software bugs, issues, and development tasks
- **Event Planning** : Organize events, track preparations, and manage timelines
- **Content Calendar** : Manage content creation, publishing schedules, and editorial workflows
- **HR Management** : Track employee onboarding, performance, and HR processes
- **Inventory Management** : Monitor inventory levels, orders, and supply chain activities
- **Customer Support** : Manage support tickets, customer issues, and resolution tracking
- **Financial Tracking** : Track budgets, expenses, and financial project data
- **Quality Assurance** : Manage QA processes, testing schedules, and quality metrics
- **Vendor Management** : Track vendor relationships, contracts, and performance
- **Training Programs** : Manage training schedules, progress, and certification tracking
- **Compliance Monitoring** : Track compliance requirements, audits, and regulatory tasks
- **Data Integration** : Integrate Monday.com data with other business systems and databases

