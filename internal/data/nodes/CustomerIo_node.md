# Customer.io Node

## Overview

Consume Customer.io API

## Credentials

- Name: customerIoApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: campaign

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign ID | number | The unique identifier for the campaign | Yes |  |

#### Operation: getMetrics

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign ID | number | The unique identifier for the campaign | Yes |  |
| Period | options | Specify metric period | No |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | collection | Integer specifying how many steps to return. Defaults to the maximum number of timeperiods available, or 12 when using the months period. Maximum timeperiods available are 24 hours, 45 days, 12 weeks and 120 months | No |  |

### Resource: customer

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier for the customer | Yes |  |

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier for the customer | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | Object of values to set as described <a href= | No |  |
| Additional Fields | collection | Property name | Yes |  |

### Resource: event

#### Operation: track

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | The unique identifier for the customer | Yes |  |
| Event Name | string | Name of the event to track | No |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | Object of values to set as described <a href= | No |  |
| Additional Fields | collection | Custom Properties | Yes |  |

#### Operation: trackAnonymous

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event Name | string | The unique identifier for the customer | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | Object of values to set as described <a href= | No |  |
| Additional Fields | collection | Custom Properties | Yes |  |

## UseCases

- **Behavioral Email Marketing**: Trigger personalized email sequences based on customer actions like purchases, sign-ups, or product usage
- **Customer Lifecycle Automation**: Create automated messaging flows for onboarding, engagement, retention, and win-back campaigns
- **E-commerce Automation**: Send abandoned cart emails, purchase confirmations, shipping updates, and product recommendations
- **Customer Segmentation**: Dynamically segment customers based on behavior, preferences, and engagement patterns for targeted campaigns
- **Multi-channel Campaigns**: Coordinate messaging across email, SMS, push notifications, and in-app messages for unified customer experiences
- **Event-Driven Marketing**: Automatically trigger marketing campaigns based on real-time customer behaviors and milestone achievements
- **Customer Analytics and Reporting**: Track campaign performance, customer engagement metrics, and conversion rates for data-driven marketing decisions
- **Subscription Management**: Automate renewal reminders, upgrade notifications, and churn prevention campaigns for subscription businesses
- **Lead Nurturing**: Create sophisticated drip campaigns to guide leads through the sales funnel with targeted content and offers
- **Customer Support Integration**: Trigger follow-up communications after support interactions to ensure customer satisfaction and gather feedback

