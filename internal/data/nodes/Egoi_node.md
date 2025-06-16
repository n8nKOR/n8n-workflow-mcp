# E-goi Node

## Overview

Consume E-goi API

## Credentials

- Name: egoiApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: contact

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | Contact ID of the subscriber | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Birth date of a subscriber | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| By | options | Search by | No |  |
| Contact ID | string | Contact ID of the subscriber | No |  |
| Email | string | Email address for subscriber | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **Email Marketing Automation**: Create, update, and manage email subscribers for automated marketing campaigns
- **Subscriber Management**: Maintain comprehensive contact databases with custom fields and subscriber information
- **Multi-channel Marketing**: Coordinate email, SMS, and push notification campaigns for unified customer engagement
- **Customer Segmentation**: Organize contacts based on demographics, behavior, and engagement patterns for targeted campaigns
- **Lead Generation**: Capture and manage leads from various sources with automated contact creation
- **Newsletter Management**: Build and maintain newsletter subscriber lists with automated sign-up processes
- **E-commerce Integration**: Sync customer data from online stores for personalized marketing campaigns
- **Customer Lifecycle Marketing**: Create automated campaigns for different stages of the customer journey
- **Re-engagement Campaigns**: Target inactive subscribers with win-back campaigns and re-engagement strategies
- **Marketing Analytics**: Track subscriber engagement, campaign performance, and conversion metrics for data-driven decisions

