# WhatsApp Business Cloud Node

## Overview

Access WhatsApp API

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: message

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| MessageType | options | The type of the message | No |  |
| Longitude | number |  | Yes |  |
| Latitude | number |  | Yes |  |
| Additional Fields | fixedCollection |  | No |  |
| Text Body | string | The body of the message (max 4096 characters) | Yes |  |
| Additional Fields | collection | The name of the file | No |  |
| Additional Fields | collection | Whether to display URL previews in text messages | No |  |

#### Operation: sendTemplate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Template | options | Name of the template | Yes |  |
| Language Code | string |  | No |  |
| Components | fixedCollection | Allows your customer to call a phone number and visit a website | No |  |

### Resource: media

#### Operation: mediaUpload

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sender Phone Number (or ID) | options | The ID of the business account | Yes |  |
| Property Name | string | Name of the binary property which contains the data for the file to be uploaded | Yes |  |
| Additional Fields | collection | The name to use for the file | No |  |

#### Operation: mediaUrlGet

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Media ID | string | The ID of the media | Yes |  |

#### Operation: mediaDelete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Media ID | string | The ID of the media | Yes |  |

## UseCases

- [Multimodal AI Chatbot] : Create intelligent WhatsApp bots that process text, images, audio, and video messages, providing AI-powered responses with support for voice transcription and image analysis capabilities
- [Customer Support Automation] : Implement automated customer service through WhatsApp with AI-powered response generation, media handling, and integration with knowledge bases for comprehensive support
- [Business Communication Hub] : Enable businesses to send automated notifications, confirmations, and updates to customers through WhatsApp templates and interactive messaging workflows
- [Media Processing Pipeline] : Automatically download, process, and analyze media files (images, videos, audio) sent through WhatsApp for content moderation, analysis, or storage workflows
- [Lead Generation and Follow-up] : Capture leads through WhatsApp interactions and automatically follow up with personalized messages, appointment scheduling, and CRM integration
- [Multi-language Support] : Provide automated translation and localized responses for global customer communication through WhatsApp Business API integration

