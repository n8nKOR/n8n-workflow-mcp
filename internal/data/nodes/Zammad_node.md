# Zammad Node

## Overview

Consume the Zammad API

## Credentials

- Name: zammadBasicAuthApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: group

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Group Name | string |  | Yes |  |
| Additional Fields | collection | Name of the custom field to set. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Group ID | string | Group to update. Specify an ID using an <a href= | Yes |  |
| Update Fields | collection | Name of the custom field to set. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Group ID | string | Group to delete. Specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Group ID | string | Group to retrieve. Specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: organization

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name | string |  | Yes |  |
| Additional Fields | collection | Whether the organization is shared with other instances | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | string | Organization to update. Specify an ID using an <a href= | Yes |  |
| Update Fields | collection | Whether the organization is shared with other instances | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | string | Organization to delete. Specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | string | Organization to retrieve. Specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: ticket

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the ticket to create | Yes |  |
| Group Name or ID | options | Group that will own the ticket to create. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Customer Email Name or ID | options | Email address of the customer concerned in the ticket to create. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Article | fixedCollection | Visible to customers | Yes |  |
| Additional Fields | collection | Name of the custom field to set. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string | Ticket to retrieve. Specify an ID using an <a href= | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string | Ticket to delete. Specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| First Name | string |  | Yes |  |
| Last Name | string |  | Yes |  |
| Additional Fields | collection | Name of the custom field to set. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | User to update. Specify an ID using an <a href= | Yes |  |
| Update Fields | collection | Name of the custom field to set. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | User to delete. Specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | User to retrieve. Specify an ID using an <a href= | Yes |  |

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Query | string |  | Yes |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Query to filter results by | No |  |

## UseCases

- **Customer Support Automation** : Automatically create tickets from customer inquiries received via email or forms
- **User Management Integration** : Sync user accounts between Zammad and other systems like LDAP or CRM
- **Ticket Escalation Workflows** : Automatically escalate tickets based on priority, time, or customer status
- **Organization Management** : Sync organization data between Zammad and external business systems
- **Support Analytics** : Extract ticket data for reporting and performance analysis
- **Multi-Channel Support** : Create tickets from various sources like chat, email, or social media
- **Customer Onboarding** : Automatically create user accounts and organizations for new customers
- **SLA Monitoring** : Track and alert on service level agreement compliance
- **Knowledge Base Integration** : Link tickets to relevant knowledge base articles or solutions
- **Team Assignment Automation** : Automatically assign tickets to appropriate groups based on content or customer type

