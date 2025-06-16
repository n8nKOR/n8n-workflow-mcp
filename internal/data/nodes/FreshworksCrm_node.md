# Freshworks CRM Node

## Overview

Consume the Freshworks CRM API

## Credentials

- Name: freshworksCrmApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: account

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the account | Yes |  |
| Additional Fields | collection | Address of the account | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Account ID | string | ID of the account to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Account ID | string | ID of the account to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| View Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Account ID | string | ID of the account to update | Yes |  |
| Update Fields | collection | Address of the account | No |  |

### Resource: appointment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the appointment | Yes |  |
| Start Date | dateTime | Timestamp that denotes the start of appointment. Start date if this is an all-day appointment. | Yes |  |
| End Date | dateTime | Timestamp that denotes the end of appointment. End date if this is an all-day appointment. | Yes |  |
| Attendees | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |
| Additional Fields | collection | ID of the user who created the appointment. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Appointment ID | string | ID of the appointment to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Appointment ID | string | ID of the appointment to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Appointment ID | string | ID of the appointment to update | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: contact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| First Name | string | First name of the contact | Yes |  |
| Last Name | string | Last name of the contact | Yes |  |
| Email Address | string | Email addresses of the contact | Yes |  |
| Additional Fields | collection | Address of the contact | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| View Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to update | Yes |  |
| Update Fields | collection | Address of the contact | No |  |

### Resource: deal

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Amount | number | Value of the deal | Yes |  |
| Name | string | Name of the deal | Yes |  |
| Additional Fields | collection | Value of the deal in base currency | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | string | ID of the deal to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | string | ID of the deal to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| View Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | string | ID of the deal to update | Yes |  |
| Update Fields | collection | Value of the deal | No |  |

### Resource: note

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Content | string | Content of the note | Yes |  |
| Target Type | options | Type of the entity for which the note is created | Yes |  |
| Target ID | string | ID of the entity for which note is created. The type of entity is selected in  | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | string | ID of the note to delete | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | string | ID of the note to update | Yes |  |
| Update Fields | collection | Content of the note | No |  |

### Resource: salesActivity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sales Activity Type Name or ID | options | ID of a sales activity type for which the sales activity is created. Choose from the list, or specify an ID using an <a href= | No |  |
| Title | string | Title of the sales activity to create | Yes |  |
| Owner Name or ID | options | ID of the user who owns the sales activity. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Start Date | dateTime | Timestamp that denotes the end of sales activity | Yes |  |
| End Date | dateTime | Timestamp that denotes the end of sales activity | Yes |  |
| Target Type | options | Type of the entity for which the sales activity is created | Yes |  |
| Target ID | string | ID of the entity for which the sales activity is created. The type of entity is selected in  | Yes |  |
| Additional Fields | collection | ID of the user who created the sales activity. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sales Activity ID | string | ID of the salesActivity to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sales Activity ID | string | ID of the salesActivity to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sales Activity ID | string | ID of the salesActivity to update | Yes |  |
| Update Fields | collection | ID of the user who created the sales activity. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: search

#### Operation: query

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Search Term | string | Enter a term that will be used for searching entities | Yes |  |
| Search on Entities | multiOptions | Enter a term that will be used for searching entities | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: lookup

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Search Field | options | Only allowed custom fields of type  | Yes |  |
| Custom Field Name | string |  | Yes |  |
| Custom Field Value | string |  | Yes |  |
| Field Value | string |  | Yes |  |
| Options | collection | Use  | No |  |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the task | Yes |  |
| Due Date | dateTime | Timestamp that denotes when the task is due to be completed | Yes |  |
| Owner Name or ID | options | ID of the user to whom the task is assigned. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Target Type | options | Type of the entity for which the task is updated | Yes |  |
| Target ID | string | ID of the entity for which the task is created. The type of entity is selected in  | Yes |  |
| Additional Fields | collection | ID of the user who created the task. Choose from the list, or specify an ID using an <a href= | No |  |

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
| Filters | collection |  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of the task to update | Yes |  |
| Update Fields | collection | ID of the user who created the sales activity. Choose from the list, or specify an ID using an <a href= | No |  |

## UseCases

- **CRM Data Synchronization**: Sync customer, contact, and account data between Freshworks CRM and external systems
- **Sales Pipeline Automation**: Automate deal creation, updates, and progression through sales stages
- **Lead Management**: Capture, qualify, and manage leads from various sources with automated contact creation
- **Customer Relationship Management**: Maintain comprehensive customer profiles with contacts, accounts, and interaction history
- **Sales Activity Tracking**: Automate logging of calls, meetings, emails, and other sales activities
- **Appointment Scheduling**: Create and manage sales appointments with automated calendar integration
- **Task Management**: Assign and track sales tasks, follow-ups, and action items for sales teams
- **Note and Documentation**: Automatically capture and organize customer notes and interaction records
- **Sales Reporting**: Extract sales data for analytics, performance tracking, and revenue reporting
- **Multi-channel Customer Search**: Search and lookup customers across multiple data points for comprehensive customer insights

