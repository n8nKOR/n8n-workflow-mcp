# Help Scout Node

## Overview

Consume Help Scout API

## Credentials

- Name: helpScoutOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: conversation

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Mailbox Name or ID | options | ID of a mailbox where the conversation is being created. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Status | options | Conversation status | Yes |  |
| Subject | string | Conversation’s subject | Yes |  |
| Type | options | Conversation type | Yes |  |
| Resolve Data | boolean | By default the response only contain the ID to resource. If this option gets activated, it will resolve the data automatically. | No |  |
| Additional Fields | collection | The Help Scout user assigned to the conversation | No |  |
| Threads | fixedCollection | The message text | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation ID | string |  | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Filters conversations by assignee ID | No |  |

### Resource: customer

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resolve Data | boolean | By default the response only contain the ID to resource. If this option gets activated, it will resolve the data automatically. | No |  |
| Additional Fields | collection | Customer’s age | No |  |
| Address | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |
| Chat Handles | fixedCollection | Chat type | No |  |
| Emails | fixedCollection | Location for this email address | No |  |
| Phones | fixedCollection | Location for this phone | No |  |
| Social Profiles | fixedCollection | Type of social profile | No |  |
| Websites | fixedCollection | Website URL | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Filters customers by first name | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string |  | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string |  | No |  |
| Update Fields | collection | Customer’s age | No |  |

### Resource: mailbox

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Mailbox ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: thread

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation ID | string |  | Yes |  |
| Type | options |  | Yes |  |
| Text | string | The chat text | Yes |  |
| Additional Fields | collection | Whether a draft reply is created | No |  |
| Attachments | fixedCollection | Attachment’s file name | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation ID | string |  | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **Customer Support Automation**: Automate customer support ticket creation, conversation management, and customer communication workflows
- **Help Desk Operations**: Streamline help desk processes with automated conversation routing, status updates, and response management
- **Customer Communication**: Manage customer interactions across multiple channels with unified conversation threads and messaging
- **Support Team Collaboration**: Coordinate support team responses, assignments, and internal communication for customer issues
- **Knowledge Base Integration**: Connect customer conversations with knowledge base articles and support documentation
- **Multi-mailbox Management**: Organize customer support across different departments, products, or service lines
- **Customer Relationship Management**: Maintain comprehensive customer profiles with interaction history and contact information
- **Performance Analytics**: Track support metrics, response times, and customer satisfaction for team performance optimization
- **Escalation Management**: Automate escalation processes for complex customer issues and priority support cases
- **Customer Satisfaction Tracking**: Monitor customer feedback, satisfaction scores, and support quality metrics

