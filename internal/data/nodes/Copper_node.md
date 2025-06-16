# Copper Node

## Overview

Consume the Copper API

## Credentials

- Name: copperApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: company

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the company to create | Yes |  |
| Additional Fields | collection | Description of the company to create | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | ID of the company to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | ID of the company to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Country of the company to filter by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | ID of the company to update | Yes |  |
| Update Fields | collection | Description to set for the company | No |  |

### Resource: customerSource

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: lead

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the lead to create | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | string | ID of the lead to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | string | ID of the lead to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Name of the country to filter by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | string | ID of the lead to update | Yes |  |
| Update Fields | collection | Description to set for the lead | No |  |

### Resource: opportunity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the opportunity to create | Yes |  |
| Customer Source ID | string | ID of the customer source that generated this opportunity | No |  |
| Primary Contact ID | string | ID of the primary company associated with this opportunity | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | string | ID of the opportunity to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | string | ID of the opportunity to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Comma-separated IDs of the primary companies to filter by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | string | ID of the opportunity to update | Yes |  |
| Update Fields | collection | ID of the primary company associated with this opportunity | No |  |

### Resource: person

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the person to create | Yes |  |
| Additional Fields | collection | Description to set for the person | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Person ID | string | ID of the person to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Person ID | string | ID of the person to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Name of the person to filter by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Person ID | string | ID of the person to update | Yes |  |
| Update Fields | collection | Description to set for the person | No |  |

### Resource: project

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the project to create | Yes |  |
| Additional Fields | collection | ID of the user who will own the project to create | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project ID | string | ID of the project to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project ID | string | ID of the project to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Name of the project to filter by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project ID | string | ID of the project to update | Yes |  |
| Update Fields | collection | ID of the user who will own the project | No |  |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Additional Fields | collection | ID of the user who will own the task to create | No |  |

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
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Comma-separated IDs of assignee IDs to filter by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of the task to update | Yes |  |
| Update Fields | collection | ID of the user who will own the task | No |  |

### Resource: user

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **CRM Data Management** : Manage customer relationships and sales data in Copper CRM
- **Sales Pipeline Automation** : Automate sales pipeline processes and opportunity tracking
- **Lead Management** : Manage leads and prospect information throughout the sales cycle
- **Customer Data Synchronization** : Sync customer data between Copper and other business systems
- **Sales Reporting** : Generate sales reports and performance analytics from Copper data
- **Contact Management** : Maintain comprehensive contact databases and interaction histories
- **Opportunity Tracking** : Track sales opportunities and deal progression
- **Activity Management** : Manage sales activities and follow-up tasks
- **Integration with Google Workspace** : Leverage Copper's native Google Workspace integration
- **Sales Team Collaboration** : Enable sales team collaboration and information sharing

