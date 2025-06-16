# Telegram Node

## Overview

Sends data to Telegram

## Credentials

- Name: telegramApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: message

#### Operation: deleteMessage

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | Unique identifier of the message to delete | Yes |  |

#### Operation: pinChatMessage

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Whether to send a notification to all chat members about the new pinned message | No |  |

#### Operation: editMessageText

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message Type | options | The type of the message to edit | No |  |
| Chat ID | string | Unique identifier for the target chat or username, To find your chat ID ask @get_id_bot | Yes |  |
| Message ID | string | Unique identifier of the message to edit | Yes |  |
| Inline Message ID | string | Unique identifier of the inline message to edit | Yes |  |
| Reply Markup | options | Additional interface options | No |  |

#### Operation: sendAnimation

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Animation | string | Animation to send. Pass a file_id to send an animation that exists on the Telegram servers (recommended), an HTTP URL for Telegram to get an animation from the Internet. | No |  |

#### Operation: sendAudio

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Audio | string | Audio file to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), an HTTP URL for Telegram to get a file from the Internet. | No |  |

#### Operation: sendChatAction

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Action | options | Type of action to broadcast. Choose one, depending on what the user is about to receive. The status is set for 5 seconds or less (when a message arrives from your bot). | No |  |

#### Operation: sendDocument

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Document | string | Document to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), an HTTP URL for Telegram to get a file from the Internet. | No |  |

#### Operation: sendLocation

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Latitude | number | Location latitude | No |  |
| Longitude | number | Location longitude | No |  |

#### Operation: sendMediaGroup

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Media | fixedCollection | The media to add | No |  |

#### Operation: sendPhoto

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Photo | string | Photo to send. Pass a file_id to send a photo that exists on the Telegram servers (recommended), an HTTP URL for Telegram to get a photo from the Internet. | No |  |

#### Operation: sendSticker

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sticker | string | Sticker to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), an HTTP URL for Telegram to get a .webp file from the Internet. | No |  |

#### Operation: sendVideo

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Video | string | Video file to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), an HTTP URL for Telegram to get a file from the Internet. | No |  |

### Resource: chat

#### Operation: member

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | Unique identifier of the target user | Yes |  |

#### Operation: setDescription

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Description | string | New chat description, 0-255 characters | Yes |  |

#### Operation: setTitle

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | New chat title, 1-255 characters | Yes |  |

### Resource: callback

#### Operation: answerQuery

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Query ID | string | Unique identifier for the query to be answered | Yes |  |
| Additional Fields | collection | The maximum amount of time in seconds that the result of the callback query may be cached client-side | No |  |

#### Operation: answerInlineQuery

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Query ID | string | Unique identifier for the answered query | Yes |  |
| Results | string | A JSON-serialized array of results for the inline query | Yes |  |
| Additional Fields | collection | The maximum amount of time in seconds that the result of the callback query may be cached client-side | No |  |

### Resource: file

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File ID | string | The ID of the file | Yes |  |
| Download | boolean | Whether to download the file | No |  |

## UseCases

- [Workflow Status Notifications] : Send real-time notifications about workflow execution status, including success confirmations and error alerts for automated processes
- [Content Publishing Alerts] : Notify administrators when Instagram content is successfully published or when upload errors occur during automated social media posting
- [System Monitoring] : Monitor database operations, API failures, and system health by sending alerts when issues are detected in automated workflows
- [User Interaction Notifications] : Send notifications when users submit forms, complete applications, or trigger important workflow events
- [Daily Recipe Sharing] : Automatically send random recipes or content to Telegram channels on a scheduled basis for community engagement
- [Multimodal AI Assistant] : Create intelligent bots that process text, voice, and image messages, providing AI-powered responses with voice transcription and image analysis capabilities
- [Long-term Memory Chatbot] : Build conversational AI agents that maintain persistent memory of user preferences and conversation history using external storage systems

