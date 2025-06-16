# GetResponse Node

## Overview

Consume GetResponse API

## Credentials

- Name: getResponseApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: contact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string |  | No |  |
| Campaign Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| Additional Fields | collection | The end user specified key of the user defined data. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of contact to delete | Yes |  |
| Options | collection | This makes it possible to pass the IP from which the contact unsubscribed. Used only if the messageId was send. | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | Unique identifier for a particular contact | Yes |  |
| Options | collection | List of fields that should be returned. ID is always returned. Fields should be separated by comma | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Search contacts by campaign ID | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | Unique identifier for a particular contact | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

## UseCases

- **Email Marketing Automation**: Create and manage email marketing campaigns with automated contact management and segmentation
- **Lead Generation and Capture**: Automatically add leads to email campaigns from websites, forms, and other lead generation sources
- **Customer Journey Automation**: Build automated email sequences based on customer behavior, preferences, and engagement patterns
- **E-commerce Integration**: Sync customer data from online stores for targeted email campaigns based on purchase behavior
- **Newsletter Management**: Automate newsletter subscriptions, unsubscribes, and content distribution to subscriber lists
- **Event-driven Marketing**: Trigger email campaigns based on user actions, website interactions, or milestone achievements
- **Customer Segmentation**: Organize contacts into targeted groups for personalized marketing campaigns and messaging
- **A/B Testing Automation**: Automate email campaign testing with different subject lines, content, and send times
- **CRM Integration**: Synchronize contact data between GetResponse and customer relationship management systems
- **Marketing Analytics**: Track email campaign performance, engagement metrics, and conversion rates for optimization

