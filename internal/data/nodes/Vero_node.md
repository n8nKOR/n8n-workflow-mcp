# Vero Node

## Overview

Consume Vero API

## Credentials

- Name: veroApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier of the customer | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | collection | The table to create the row in | No |  |
| Data | fixedCollection | Key value pairs that represent the custom user properties you want to update | No |  |
| Data | json | Key value pairs that represent the custom user properties you want to update | No |  |

#### Operation: alias

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The old unique identifier of the user | Yes |  |
| New ID | string | The new unique identifier of the user | Yes |  |

#### Operation: unsubscribe

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier of the user | Yes |  |

#### Operation: resubscribe

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier of the user | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier of the user | Yes |  |

#### Operation: addTags

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier of the user | Yes |  |
| Tags | string | Tags to add separated by  | Yes |  |

#### Operation: removeTags

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier of the user | Yes |  |
| Tags | string | Tags to remove separated by  | Yes |  |

### Resource: event

#### Operation: track

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The unique identifier of the customer | Yes |  |
| Email | string |  | Yes |  |
| Event Name | string | The name of the event tracked | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Data | fixedCollection | Key value pairs that represent any properties you want to track with this event | No |  |
| Extra | fixedCollection | Key value pairs that represent reserved, Vero-specific operators. Refer to the note on "deduplication" below. | No |  |
| Data | json | Key value pairs that represent the custom user properties you want to update | No |  |
| Extra | json | Key value pairs that represent reserved, Vero-specific operators. Refer to the note on "deduplication" below. | No |  |

## UseCases

- **Email Marketing Automation** : Create and manage automated email campaigns based on user behavior and events
- **Customer Journey Mapping** : Track user interactions and create personalized customer journey workflows
- **User Segmentation** : Segment users based on properties and behaviors for targeted messaging
- **Event-Based Triggers** : Set up automated responses based on user actions and events
- **User Profile Management** : Maintain comprehensive user profiles with custom properties and tags
- **Behavioral Analytics** : Track and analyze user behavior patterns for marketing insights
- **Subscription Management** : Handle user subscription preferences and unsubscribe requests
- **Tag-Based Organization** : Organize users with tags for efficient campaign targeting
- **User Lifecycle Management** : Manage user lifecycle from onboarding to retention campaigns
- **Cross-Platform Integration** : Integrate Vero with other platforms for unified customer data management

