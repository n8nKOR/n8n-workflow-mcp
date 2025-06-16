# Discord Node

## Overview

Sends data to Discord

## Credentials

- Name: discordBotApi, Required: Yes (for Bot Token)
- Name: discordOAuth2Api, Required: Yes (for OAuth2)  
- Name: discordWebhookApi, Required: Yes (for Webhook)

## Inputs

- Main

## Outputs

- Main

## Properties

### Authentication Options

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Connection Type | options | Type of authentication to use | Yes | botToken |

#### Connection Type Options:
- **Bot Token**: Manage messages, channels, and members on a server
- **OAuth2**: Same features as 'Bot Token' with easier Bot installation
- **Webhook**: Send messages to a specific channel

### Resource: channel

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Guild ID | string | The ID of the Discord server | Yes |  |
| Channel ID | string | The ID of the channel to get | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Guild ID | string | The ID of the Discord server | Yes |  |
| Channel Name | string | Name of the channel to create | Yes |  |
| Channel Type | options | Type of channel to create | Yes | text |
| Additional Fields | collection | Additional channel options | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Guild ID | string | The ID of the Discord server | Yes |  |
| Channel ID | string | The ID of the channel to update | Yes |  |
| Update Fields | collection | Fields to update | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Guild ID | string | The ID of the Discord server | Yes |  |
| Channel ID | string | The ID of the channel to delete | Yes |  |

### Resource: message

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Channel ID | string | The ID of the channel to send the message to | Yes |  |
| Message Content | string | The content of the message | No |  |
| Embeds | fixedCollection | Rich embeds to include with the message | No |  |
| Additional Fields | collection | Additional message options | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Channel ID | string | The ID of the channel | Yes |  |
| Message ID | string | The ID of the message to get | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Channel ID | string | The ID of the channel | Yes |  |
| Message ID | string | The ID of the message to update | Yes |  |
| Message Content | string | The new content of the message | No |  |
| Embeds | fixedCollection | Rich embeds to include with the message | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Channel ID | string | The ID of the channel | Yes |  |
| Message ID | string | The ID of the message to delete | Yes |  |

### Resource: member

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Guild ID | string | The ID of the Discord server | Yes |  |
| User ID | string | The ID of the user to get member information for | Yes |  |

#### Operation: ban

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Guild ID | string | The ID of the Discord server | Yes |  |
| User ID | string | The ID of the user to ban | Yes |  |
| Additional Fields | collection | Additional ban options | No |  |

#### Operation: kick

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Guild ID | string | The ID of the Discord server | Yes |  |
| User ID | string | The ID of the user to kick | Yes |  |
| Reason | string | Reason for the kick | No |  |

### Webhook Resource

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message Content | string | The content of the message | No |  |
| Username | string | Username to use for the webhook | No |  |
| Avatar URL | string | Avatar URL to use for the webhook | No |  |
| Embeds | fixedCollection | Rich embeds to include with the message | No |  |

## UseCases

- [AI-Powered Customer Support Routing] : Automatically analyze user feedback messages received through Discord webhooks and route them to appropriate departments (Customer Success, IT, Helpdesk) based on AI classification of message content and urgency level
- [YouTube Content Sharing with AI Summaries] : Share YouTube videos with AI-generated summaries to Discord channels, providing quick content overviews for community members without requiring them to watch full videos
- [Daily Automated Content Broadcasting] : Send daily translated comics, news updates, or scheduled content to Discord channels using webhooks for consistent community engagement and information sharing
- [Multi-Department Notification System] : Create automated notification workflows that send structured messages to different Discord channels based on business events, alerts, or system status updates

