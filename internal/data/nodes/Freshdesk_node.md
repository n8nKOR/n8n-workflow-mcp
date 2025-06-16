# Freshdesk Node

## Overview

Consume Freshdesk API

## Credentials

- Name: freshdeskApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: ticket

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Requester Identification | options | Email address of the requester. If no contact exists with this email address in Freshdesk, it will be added as a new contact. | Yes |  |
| Value | string | Value of the identification selected | Yes |  |
| Status | options |  | Yes |  |
| Priority | options |  | Yes |  |
| Source | options | The channel through which the ticket was created | Yes |  |
| Options | collection | ID of the agent to whom the ticket has been assigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string |  | Yes |  |
| Update Fields | collection | ID of the agent to whom the ticket has been assigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Order sort attribute ascending or descending | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string |  | Yes |  |

### Resource: contact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the contact | Yes |  |
| Email | string | Primary email address of the contact. If you want to associate additional email(s) with this contact, use the other_emails attribute. | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection |  | No |  |

## UseCases

- **Customer Support Automation**: Automate ticket creation, assignment, and escalation for customer support requests and inquiries
- **Multi-channel Support Integration**: Consolidate support tickets from email, chat, phone, and social media into unified workflows
- **Contact Management**: Maintain comprehensive customer contact databases with automated contact creation and updates
- **Ticket Lifecycle Management**: Automate ticket status updates, priority changes, and resolution tracking throughout the support process
- **Agent Productivity**: Streamline agent workflows with automated ticket routing, bulk operations, and performance tracking
- **Customer Communication**: Automate customer notifications, status updates, and follow-up communications for support tickets
- **SLA Management**: Monitor and enforce service level agreements with automated escalation and deadline tracking
- **Knowledge Base Integration**: Connect support tickets with knowledge base articles for faster resolution and self-service
- **Customer Satisfaction Tracking**: Automate feedback collection, satisfaction surveys, and customer experience monitoring
- **Reporting and Analytics**: Generate support metrics, agent performance reports, and customer satisfaction analytics

