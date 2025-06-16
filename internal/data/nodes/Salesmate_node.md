# Salesmate Node

## Overview

Consume Salesmate API

## Credentials

- Name: salesmateApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: company

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Owner Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |
| Additional Fields | collection |  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string |  | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string |  | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| JSON Parameters | boolean |  | No |  |
| Options | collection | Comma-separated list of fields to return | No |  |
| Filters | json |  | No |  |
| Filters | fixedCollection | Value of the property to set | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | If more than one company add them separated by , | Yes |  |

### Resource: activity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string |  | Yes |  |
| Owner Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Type | string | This field displays activity type such as call, meeting etc | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |
| Additional Fields | collection | This field contains details related to the activity | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | string |  | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |
| Update Fields | collection | This field contains details related to the activity | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | string |  | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| JSON Parameters | boolean |  | No |  |
| Options | collection | Comma-separated list of fields to return | No |  |
| Filters | json |  | No |  |
| Filters | fixedCollection | Value of the property to set | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | string | If more than one activity add them separated by , | Yes |  |

### Resource: deal

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string |  | Yes |  |
| Owner Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Primary Contact Name or ID | options | Primary contact for the deal. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Pipeline | options |  | Yes |  |
| Status | options |  | Yes |  |
| Stage | options |  | Yes |  |
| Currency | string |  | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |
| Additional Fields | collection | This field contains details related to the deal | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | string |  | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | string |  | Yes |  |
| RAW Data | boolean | Whether the data should include the fields details | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| JSON Parameters | boolean |  | No |  |
| Options | collection | Comma-separated list of fields to return | No |  |
| Filters | json |  | No |  |
| Filters | fixedCollection | Value of the property to set | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | string | If more than one deal add them separated by , | Yes |  |

## UseCases

- [Sales Pipeline Management] : Track and manage sales opportunities through the entire sales cycle
- [Contact Management] : Maintain comprehensive contact databases with interaction history
- [Lead Generation] : Capture and manage leads from various marketing sources
- [Deal Tracking] : Monitor deal progress and forecast sales revenue
- [Customer Relationship Management] : Build and maintain strong customer relationships
- [Sales Activity Tracking] : Track sales calls, emails, and meetings with prospects
- [Team Collaboration] : Enable sales team collaboration and information sharing
- [Sales Reporting] : Generate sales performance reports and analytics
- [Email Integration] : Integrate email communication with CRM for better tracking
- [Task Management] : Manage sales tasks and follow-up activities
- [Sales Automation] : Automate repetitive sales processes and workflows
- [Customer Support] : Provide customer support through integrated ticketing system
- [Mobile Sales] : Access CRM data and functionality on mobile devices
- [Integration Management] : Connect Salesmate with other business tools and platforms
- [Performance Analytics] : Analyze sales team performance and identify improvement areas

