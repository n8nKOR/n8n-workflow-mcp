# Zendesk Node

## Overview

Consume Zendesk API

## Credentials

- Name: zendeskApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: ticket

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Description | string | The first comment on the ticket | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | collection | Custom field ID. Choose from the list, or specify an ID using an <a href= | No |  |
| Additional Fields | json | Object of values to set as described <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string |  | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Update Fields | collection | The e-mail address of the assignee | No |  |
| Update Fields | json | Object of values to update as described <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string |  | Yes |  |
| Suspended Ticket ID | string | Ticket ID | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | The group to search. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string |  | Yes |  |
| Suspended Ticket ID | string | Ticket ID | Yes |  |

#### Operation: recover

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Suspended Ticket ID | string |  | Yes |  |

### Resource: ticketField

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket Field ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The user | Yes |  |
| Additional Fields | collection | An alias displayed to end users | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | Yes |  |
| Update Fields | collection | An alias displayed to end users | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | Yes |  |

#### Operation: getRelatedData

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | Yes |  |

#### Operation: getOrganizations

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | Yes |  |

### Resource: organization

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Additional Fields | collection | Details about the organization, such as the address | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | string |  | Yes |  |
| Update Fields | collection | Details about the organization, such as the address | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | string |  | Yes |  |

#### Operation: getRelatedData

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | string |  | Yes |  |

## UseCases

- [Customer Support] : Manage customer support tickets and resolve customer issues efficiently
- [Ticket Management] : Track and organize support tickets with priority and status management
- [Knowledge Base] : Maintain self-service knowledge base for customers and support agents
- [Agent Productivity] : Improve support agent productivity with automated workflows
- [Customer Communication] : Communicate with customers through multiple channels (email, chat, phone)
- [Service Level Management] : Monitor and maintain service level agreements with customers
- [Escalation Management] : Automatically escalate tickets based on priority and response time
- [Reporting and Analytics] : Generate support metrics and analyze customer satisfaction
- [Multi-Channel Support] : Provide support across email, chat, social media, and phone
- [Customer Satisfaction] : Track and improve customer satisfaction scores and feedback
- [Team Collaboration] : Enable support team collaboration and knowledge sharing
- [Integration Workflows] : Integrate Zendesk with CRM and other business systems
- [Automation Rules] : Automate ticket routing, assignment, and response workflows
- [User Management] : Manage customer and agent accounts with role-based permissions
- [Performance Monitoring] : Monitor support team performance and identify improvement areas

