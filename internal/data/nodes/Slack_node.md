# Slack Node

## Overview

Consume Slack API

⚠️ **Version Information**: This node has multiple versions (V1, V2). The current default version is V2 (2.3), which provides enhanced authentication options, improved channel management, and expanded message formatting capabilities. Some features may differ between versions.

## Credentials

- Name: slackApi (for Access Token authentication), Required: Yes
- Name: slackOAuth2Api (for OAuth2 authentication), Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Channel

#### Operation: Archive

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID or channel name to archive | Yes |

#### Operation: Close

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to close | Yes |

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Name | string | Channel name to create | Yes |
| Is Private | boolean | Whether the channel should be private | No |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to retrieve information for | Yes |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Return All | boolean | Whether to return all results | No |
| Limit | number | Maximum number of channels to return | No |

#### Operation: History

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to get history for | Yes |
| Return All | boolean | Whether to return all messages | No |
| Limit | number | Maximum number of messages to return | No |

#### Operation: Invite

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to invite user to | Yes |
| User | string | User ID to invite | Yes |

#### Operation: Join

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID or name to join | Yes |

#### Operation: Kick

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to kick user from | Yes |
| User | string | User ID to kick | Yes |

#### Operation: Leave

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to leave | Yes |

#### Operation: Member

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to get members for | Yes |
| Return All | boolean | Whether to return all members | No |
| Limit | number | Maximum number of members to return | No |

#### Operation: Open

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to open | Yes |

#### Operation: Rename

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to rename | Yes |
| Name | string | New channel name | Yes |

#### Operation: Replies

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID containing the thread | Yes |
| TS | string | Timestamp of the thread message | Yes |
| Return All | boolean | Whether to return all replies | No |
| Limit | number | Maximum number of replies to return | No |

#### Operation: Set Purpose

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to set purpose for | Yes |
| Purpose | string | Channel purpose | Yes |

#### Operation: Set Topic

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to set topic for | Yes |
| Topic | string | Channel topic | Yes |

#### Operation: Unarchive

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to unarchive | Yes |

### Resource: Message

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID containing the message | Yes |
| Timestamp | string | Timestamp of message to delete | Yes |

#### Operation: Get Permalink

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID containing the message | Yes |
| Timestamp | string | Timestamp of the message | Yes |

#### Operation: Post

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID or name to post message to | Yes |
| Text | string | Message text | Yes |
| As User | boolean | Whether to post as authenticated user | No |
| Username | string | Username to post as (bot mode) | No |
| Attachments | fixedCollection | Message attachments | No |
| Blocks | json | Slack blocks for rich message formatting | No |
| Icon Emoji | string | Emoji to use as icon | No |
| Icon URL | string | URL of image to use as icon | No |
| Link Names | boolean | Whether to linkify channel and user names | No |
| Markdown | boolean | Whether to enable markdown formatting | No |
| Parse | options | How to parse the message | No |
| Reply Broadcast | boolean | Whether to broadcast reply to channel | No |
| Thread TS | string | Timestamp of thread to reply to | No |
| Unfurl Links | boolean | Whether to unfurl links | No |
| Unfurl Media | boolean | Whether to unfurl media | No |

#### Operation: Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Query | string | Search query | Yes |
| Sort | options | How to sort results | No |
| Sort Direction | options | Sort direction (asc, desc) | No |
| Return All | boolean | Whether to return all results | No |
| Limit | number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID containing the message | Yes |
| Text | string | Updated message text | Yes |
| Timestamp | string | Timestamp of message to update | Yes |
| As User | boolean | Whether to update as authenticated user | No |
| Attachments | fixedCollection | Updated message attachments | No |
| Blocks | json | Updated Slack blocks | No |
| Link Names | boolean | Whether to linkify channel and user names | No |
| Parse | options | How to parse the message | No |

### Resource: File

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID to get files from | No |
| Return All | boolean | Whether to return all files | No |
| Limit | number | Maximum number of files to return | No |

#### Operation: Upload

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| File | string | Binary property containing file data | Yes |
| Channels | multiOptions | Channels to share file in | No |
| File Comment | string | Initial comment for the file | No |
| File Name | string | Name of the file | No |
| Title | string | Title of the file | No |

### Resource: Reaction

#### Operation: Add

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID containing the message | Yes |
| Name | string | Reaction emoji name | Yes |
| Timestamp | string | Timestamp of message to add reaction to | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID containing the message | Yes |
| Timestamp | string | Timestamp of message to get reactions for | Yes |

#### Operation: Remove

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Channel | string | Channel ID containing the message | Yes |
| Name | string | Reaction emoji name to remove | Yes |
| Timestamp | string | Timestamp of message to remove reaction from | Yes |

### Resource: Star

#### Operation: Add

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Target | options | What to star (message, file, comment) | Yes |
| Channel | string | Channel ID (for messages) | No |
| File | string | File ID (for files) | No |
| File Comment | string | File comment ID (for comments) | No |
| Timestamp | string | Message timestamp (for messages) | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Target | options | What to unstar (message, file, comment) | Yes |
| Channel | string | Channel ID (for messages) | No |
| File | string | File ID (for files) | No |
| File Comment | string | File comment ID (for comments) | No |
| Timestamp | string | Message timestamp (for messages) | No |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Return All | boolean | Whether to return all starred items | No |
| Limit | number | Maximum number of starred items to return | No |

### Resource: User

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| User | string | User ID to get information for | Yes |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Return All | boolean | Whether to return all users | No |
| Limit | number | Maximum number of users to return | No |

#### Operation: Get Presence

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| User | string | User ID to get presence for | Yes |

#### Operation: Update Profile

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Options | fixedCollection | Profile fields to update | Yes |

### Resource: User Group

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Name | string | User group name | Yes |
| Handle | string | User group handle | Yes |
| Description | string | User group description | No |
| Channels | multiOptions | Channels to add the group to | No |

#### Operation: Disable

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| User Group | string | User group ID to disable | Yes |

#### Operation: Enable

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| User Group | string | User group ID to enable | Yes |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| Return All | boolean | Whether to return all user groups | No |
| Limit | number | Maximum number of user groups to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method (Access Token, OAuth2) | Yes |
| User Group | string | User group ID to update | Yes |
| Update Fields | fixedCollection | Fields to update | Yes |

## UseCases

- [AI Slack Bot with Google Gemini] : Implement an AI personal assistant bot that receives messages from Slack and uses Google Gemini AI model to provide automation-related help and advice. Maintains context by storing conversation history in memory and sends AI responses back to Slack channels
- [Customer Support Ticketing System] : Monitor Slack channels for messages with ticket emojis, search for existing support requests, and automatically create structured tickets in Linear with AI-generated titles, summaries, and priority levels
- [AI-Powered Information Broadcasting] : Send AI-summarized articles and news updates to Slack channels with proper Slack markdown formatting, including clickable links, bullet points, and structured content for team information sharing
- [Automated Content Monitoring] : Search Slack messages for specific patterns, keywords, or support requests, and trigger automated workflows for ticket creation, escalation, or response generation

