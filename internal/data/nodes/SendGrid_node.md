# SendGrid Node

## Overview

Consume SendGrid API

## Credentials

- Name: sendGridApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: list

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the list | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | ID of the list | Yes |  |
| Delete Contacts | boolean | Whether to delete all contacts on the list | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | ID of the list | Yes |  |
| Contact Sample | boolean | Whether to return the contact sample | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | ID of the list | Yes |  |
| Name | string | Name of the list | Yes |  |

### Resource: contact

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | The query field accepts valid <a href= | No |  |

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | Primary email for the contact | Yes |  |
| Additional Fields | collection | Adds a custom field to set also values which have not been predefined | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact IDs | string | ID of the contact. Multiple can be added separated by comma. | No |  |
| Delete All | boolean | Whether all contacts will be deleted | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| By | options | Search the user by ID or email | Yes |  |
| Contact ID | string | ID of the contact | Yes |  |
| Email | string | Email of the contact | Yes |  |

### Resource: mail

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sender Email | string | Email address of the sender of the email | No |  |
| Sender Name | string | Name of the sender of the email | No |  |
| Recipient Email | string | Comma-separated list of recipient email addresses | No |  |
| Subject | string | Subject of the email to send | No |  |
| Dynamic Template | boolean | Whether this email will contain a dynamic template | Yes |  |
| MIME Type | options | MIME type of the email to send | No |  |
| Message Body | string | Message body of the email to send | Yes |  |
| Template Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| Dynamic Template Fields | fixedCollection | Key of the dynamic template field | No |  |
| Additional Fields | collection | Comma-separated list of binary properties | No |  |

## UseCases

- **Transactional Email Automation** : Send automated transactional emails like order confirmations and receipts
- **Marketing Campaign Management** : Execute email marketing campaigns and newsletter distributions
- **Customer Onboarding** : Send welcome emails and onboarding sequences to new customers
- **Notification Systems** : Send system notifications and alerts to users and administrators
- **E-commerce Communications** : Send shipping notifications, delivery updates, and promotional offers
- **Event-Driven Messaging** : Trigger emails based on user actions and system events
- **Customer Support** : Send support ticket updates and resolution notifications
- **Subscription Management** : Manage email subscriptions and unsubscribe requests
- **A/B Testing** : Conduct email A/B tests to optimize campaign performance
- **Personalized Communications** : Send personalized emails based on user preferences and behavior

