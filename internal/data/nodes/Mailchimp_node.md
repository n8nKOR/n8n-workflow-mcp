# Mailchimp Node

## Overview

Consume Mailchimp API

## Credentials

- Name: mailchimpApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: member

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List Name or ID | options | List of lists. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Email | string | Email address for a subscriber | Yes |  |
| Status | options | Subscriber | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Options | collection | Type of email this member asked to get | No |  |
| Location | fixedCollection | Subscriber location information.n | Yes |  |
| Merge Fields | fixedCollection | An individual merge var and value for a member | Yes |  |
| Merge Fields | json |  | No |  |
| Location | json |  | No |  |
| Interest Groups | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |
| Interest Groups | json |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List Name or ID | options | List of lists. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Email | string | Member | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List Name or ID | options | List of lists. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Email | string | Member | Yes |  |
| Options | collection | A comma-separated list of fields to return | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List Name or ID | options | List of lists. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Restrict results to subscribers whose information changed before the set timeframe | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List Name or ID | options | List of lists. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Email | string | Email address of the subscriber | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Update Fields | collection | Type of email this member asked to get | Yes |  |
| Merge Fields | json |  | No |  |
| Location | json |  | No |  |
| Interest Groups | json |  | No |  |

### Resource: listGroup

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List Name or ID | options | List of lists. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Group Category Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: campaign

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Restrict the response to campaigns created before the set time | No |  |

## UseCases

- **Email Marketing Automation**: Create and manage comprehensive email marketing campaigns with subscriber management and list segmentation
- **Newsletter Management**: Automate newsletter subscriptions, content distribution, and subscriber engagement tracking
- **E-commerce Integration**: Sync customer data for targeted email campaigns based on purchase behavior and shopping patterns
- **Lead Nurturing**: Develop automated email sequences for lead generation, qualification, and customer acquisition
- **Customer Segmentation**: Organize subscribers into targeted groups based on demographics, interests, and behavior for personalized campaigns
- **Marketing Analytics**: Track email campaign performance, engagement metrics, and ROI for data-driven marketing optimization
- **Subscriber Lifecycle Management**: Manage subscriber onboarding, engagement, retention, and re-activation campaigns
- **Event and Webinar Promotion**: Promote events, webinars, and special offers through targeted email campaigns
- **Cross-channel Marketing**: Integrate email marketing with social media, content marketing, and other digital marketing channels
- **Automated Drip Campaigns**: Create sophisticated drip campaigns for customer education, product launches, and brand engagement

