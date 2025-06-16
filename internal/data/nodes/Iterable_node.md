# Iterable Node

## Overview

Consume Iterable API

## Credentials

- Name: iterableApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: event

#### Operation: track

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The name of the event to track | Yes |  |
| Additional Fields | collection | Campaign tied to conversion | No |  |

### Resource: user

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Identifier | options | Identifier to be used | Yes |  |
| Value | string |  | Yes |  |
| Create If Doesn | boolean | Whether to create a new user if the idetifier does not exist | Yes |  |
| Additional Fields | collection | The end user specified key of the user defined data | No | "false)" |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| By | options | Identifier to be used | Yes |  |
| User ID | string | Unique identifier for a particular user | Yes |  |
| Email | string | Email for a particular user | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| By | options | Identifier to be used | Yes |  |
| User ID | string | Unique identifier for a particular user | Yes |  |
| Email | string | Email for a particular user | Yes |  |

### Resource: userList

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List Name or ID | options | Identifier to be used. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Identifier | options | Identifier to be used | Yes |  |
| Value | string |  | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List Name or ID | options | Identifier to be used. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Identifier | options | Identifier to be used | Yes |  |
| Value | string |  | Yes |  |
| Additional Fields | collection | Attribute unsubscribe to a campaign | No |  |

## UseCases

- **Cross-Channel Marketing Automation**: Orchestrate email, SMS, push notifications, and in-app messaging campaigns across multiple channels
- **Event-Driven Marketing**: Track customer behavior events and trigger personalized marketing campaigns based on user actions
- **Customer Journey Orchestration**: Create sophisticated customer journey workflows with automated touchpoints and personalization
- **User Lifecycle Management**: Manage user onboarding, engagement, retention, and re-activation campaigns
- **List Management and Segmentation**: Organize customers into targeted lists for precise audience segmentation and campaign targeting
- **Behavioral Targeting**: Trigger campaigns based on customer interactions, purchases, and engagement patterns
- **Marketing Analytics**: Track campaign performance, user engagement, and conversion metrics for data-driven marketing decisions
- **E-commerce Marketing**: Automate abandoned cart recovery, product recommendations, and post-purchase follow-up campaigns
- **Mobile App Engagement**: Drive app engagement through push notifications, in-app messaging, and mobile-specific campaigns
- **Customer Data Platform Integration**: Unify customer data across channels for comprehensive user profiles and personalized experiences

