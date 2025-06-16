# Mattermost Node

## Overview

The Mattermost node allows you to send data to Mattermost instances. Mattermost is an open-source messaging platform designed for team collaboration, providing secure messaging, file sharing, and integration capabilities. This node enables you to manage channels, send messages, handle reactions, and manage users within your Mattermost workspace through n8n workflows.

## Credentials

This node requires Mattermost API credentials:
- **Server URL**: Your Mattermost server URL (e.g., https://your-mattermost.com)
- **Access Token**: Personal access token or bot token with appropriate permissions

To obtain API credentials:
1. Log into your Mattermost instance as an admin or user with appropriate permissions
2. Go to **Account Settings** → **Security** → **Personal Access Tokens**
3. Create a new token with the required permissions
4. Copy the token for use in n8n (this is shown only once)

## Inputs

- **Main**: JSON data for creating/updating resources

## Outputs

- **Main**: Returns JSON data from Mattermost API responses

## Properties

### Resource: Channel

#### Operation: Add User
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel to add user to | Yes |
| User ID | String | ID of the user to add to channel | Yes |

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Team ID | String | ID of the team to create channel in | Yes |
| Display Name | String | Display name for the channel | Yes |
| Name | String | Channel name (URL-friendly) | Yes |
| Type | Options | Channel type (public, private) | Yes |
| Additional Fields | Collection | Optional channel properties | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel to delete | Yes |

#### Operation: Member
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel to get members from | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Restore
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel to restore | Yes |

#### Operation: Search
| Field Name | Type | Description | Required |
|---|---|---|---|
| Team ID | String | ID of the team to search in | Yes |
| Term | String | Search term | Yes |

#### Operation: Statistics
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel to get statistics for | Yes |

### Resource: Message

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel containing the message | Yes |
| Message ID | String | ID of the message to delete | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Post
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel to post to | Yes |
| Message | String | Message content to post | Yes |
| Additional Fields | Collection | Optional message properties | No |

#### Operation: Post Ephemeral
| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel ID | String | ID of the channel to post ephemeral message to | Yes |
| User ID | String | ID of the user who will see the ephemeral message | Yes |
| Message | String | Ephemeral message content (visible only to specified user) | Yes |
| Additional Fields | Collection | Optional ephemeral message properties | No |

### Resource: Reaction

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | String | ID of the message to add reaction to | Yes |
| Emoji Name | String | Name of the emoji reaction | Yes |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | String | ID of the message | Yes |
| Emoji Name | String | Name of the emoji reaction to remove | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | String | ID of the message | Yes |

### Resource: User

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Username | String | Username for the user | Yes |
| Email | String | Email address for the user | Yes |
| Password | String | Password for the user | Yes |
| Additional Fields | Collection | Optional user properties | No |

#### Operation: Deactivate
| Field Name | Type | Description | Required |
|---|---|---|---|
| User ID | String | ID of the user to deactivate | Yes |

#### Operation: Get By Email
| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | String | Email address of the user to retrieve | Yes |

#### Operation: Get By ID
| Field Name | Type | Description | Required |
|---|---|---|---|
| User ID | String | ID of the user to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Invite
| Field Name | Type | Description | Required |
|---|---|---|---|
| Team ID | String | ID of the team to invite user to | Yes |
| Email | String | Email address of user to invite | Yes |

## Features

- **Team Collaboration**: Manage team communications and channels
- **Message Management**: Send, retrieve, and delete messages including ephemeral messages
- **User Administration**: Create, manage, and invite users with flexible retrieval options
- **Channel Operations**: Create, search, and manage channels with member management
- **Reaction Handling**: Add and remove emoji reactions
- **Statistics**: Get channel and user statistics
- **File Attachments**: Support for file uploads and attachments

## UseCases

### Team Communication
- **Automated Notifications**: Send workflow status updates to team channels
- **Alert Systems**: Post critical alerts and notifications
- **Bot Integration**: Create chatbots for team assistance
- **Status Updates**: Broadcast system or project status

### Project Management
- **Task Notifications**: Notify team when tasks are completed
- **Meeting Reminders**: Send automated meeting reminders
- **Progress Reports**: Generate and share progress updates
- **Issue Tracking**: Post bug reports and issues to relevant channels

### DevOps Integration
- **Deployment Notifications**: Notify teams of deployments
- **Monitoring Alerts**: Send monitoring and error alerts
- **CI/CD Integration**: Report build and test results
- **Infrastructure Updates**: Notify about infrastructure changes

### Customer Support
- **Support Ticket Updates**: Notify support team of new tickets
- **Escalation Alerts**: Alert management of escalated issues
- **Response Tracking**: Track response times and updates
- **Knowledge Sharing**: Share solutions and knowledge

## Authentication & Security

### Access Tokens
- **Personal Access Tokens**: For individual user actions
- **Bot Tokens**: For automated system actions
- **Session Tokens**: For temporary access (not recommended for automation)

### Permissions
- **Channel Permissions**: Read, write, manage channel permissions
- **Team Permissions**: Team management and member administration
- **System Permissions**: Admin-level operations (user creation, etc.)

### Best Practices
- **Token Rotation**: Regularly rotate access tokens
- **Minimal Permissions**: Grant only necessary permissions
- **Secure Storage**: Store tokens securely in n8n credentials
- **Rate Limiting**: Respect API rate limits

## API Rate Limits

Mattermost API has rate limiting:
- **Per User**: Varies by Mattermost edition and configuration
- **Per App**: Based on server configuration
- **Burst Limits**: Temporary higher limits for short bursts

## Error Handling

Common issues and solutions:
- **Authentication Failed**: Verify token and permissions
- **Channel Not Found**: Check channel ID and access permissions
- **User Not Found**: Verify user ID and team membership
- **Rate Limit Exceeded**: Implement retry logic with delays
- **Invalid Message Format**: Check message content and formatting

## Integration Examples

### Workflow Notifications
```json
{
  "resource": "message",
  "operation": "post",
  "channelId": "{{ $json.channelId }}",
  "message": "Workflow '{{ $workflow.name }}' completed successfully!"
}
```

### Error Alerts
```json
{
  "resource": "message", 
  "operation": "post",
  "channelId": "alerts-channel-id",
  "message": "🚨 Error in {{ $json.system }}: {{ $json.error }}"
}
```

### User Management
```json
{
  "resource": "user",
  "operation": "create",
  "username": "{{ $json.username }}",
  "email": "{{ $json.email }}",
  "password": "{{ $json.temporaryPassword }}"
}
```

## Advanced Features

### Message Formatting
- **Markdown Support**: Rich text formatting with Markdown
- **Mentions**: @user and @channel mentions
- **Emoji Support**: Unicode and custom emoji reactions
- **File Attachments**: Upload and share files

### Webhook Integration
- **Incoming Webhooks**: Receive data from external systems
- **Outgoing Webhooks**: Send data to external systems
- **Slash Commands**: Create custom slash commands
- **Interactive Messages**: Buttons and interactive elements

### Channel Management
- **Public Channels**: Open team communication
- **Private Channels**: Restricted team communication
- **Direct Messages**: One-on-one conversations
- **Group Messages**: Small group conversations

