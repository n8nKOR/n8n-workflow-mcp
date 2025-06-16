# Matrix Node

## Overview

Consume Matrix API

## Credentials

- Name: matrixApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: account

#### Operation: me

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| (No parameters) | - | Get current user's account information | - |  |

### Resource: event

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room ID | string | The room related to the event | Yes |  |
| Event ID | string | The event ID to retrieve | Yes |  |

### Resource: media

#### Operation: upload

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room Name or ID | options | Room ID to upload media to. Choose from the list, or specify an ID using expression | Yes |  |
| Input Binary Field | string | Binary field containing the media file to upload | Yes |  |
| Media Type | options | Type of media being uploaded | Yes | General file |
| Additional Fields | collection | Additional media upload options | No |  |

#### Media Upload Additional Fields
| Field Name | Type | Description |
|---|---|---|
| Filename | string | Name of the file being uploaded |
| Content Type | string | MIME type of the media file |

### Resource: message

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room Name or ID | options | The channel to send the message to. Choose from the list, or specify an ID using expression | Yes |  |
| Text | string | The text content to send | No |  |
| Message Type | options | Type of message to send | No | Text message |
| Message Format | options | Format of the message content | No | Text only |
| Fallback Text | string | A plain text message to display in case the HTML cannot be rendered by the Matrix client | No |  |

#### Operation: Get Many Messages

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room Name or ID | options | The room to retrieve messages from. Choose from the list, or specify an ID using expression | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | Yes |  |
| Limit | number | Max number of results to return | No |  |
| Other Options | collection | Advanced filtering and pagination options | No |  |

#### Message Retrieval Options
| Field Name | Type | Description |
|---|---|---|
| From Token | string | Token to start returning events from (for pagination) |
| To Token | string | Token to stop returning events at |
| Direction | options | Direction to return events (backwards/forwards) |
| Filter | string | JSON RoomEventFilter to filter returned events |

### Resource: room

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room Name | string | Display name for the new room | Yes |  |
| Preset | options | Room visibility and permissions preset | Yes | Open and public chat |
| Room Alias | string | Room alias (e.g., #roomname:domain.com) | No |  |

#### Operation: join

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room ID or Alias | string | Room identifier or alias to join | Yes |  |

#### Operation: leave

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room Name or ID | options | Room to leave. Choose from the list, or specify an ID using expression | Yes |  |

#### Operation: invite

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room Name or ID | options | Room to invite user to. Choose from the list, or specify an ID using expression | Yes |  |
| User ID | string | The fully qualified user ID of the invitee (e.g., @user:domain.com) | Yes |  |

#### Operation: kick

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room Name or ID | options | Room to kick user from. Choose from the list, or specify an ID using expression | Yes |  |
| User ID | string | The fully qualified user ID to kick | Yes |  |
| Reason | string | Reason for removing the user | No |  |

### Resource: roomMember

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Room Name or ID | options | Room to get members from. Choose from the list, or specify an ID using expression | Yes |  |
| Filters | collection | Member filtering options | No |  |

#### Member Filtering Options
| Field Name | Type | Description |
|---|---|---|
| Membership | options | Filter by membership state (join, invite, leave, ban) |
| Not Membership | options | Exclude specific membership states |

## UseCases

- **Team Communication Hub**: Create automated team communication channels for project updates, announcements, and real-time collaboration across distributed teams
- **Incident Response Coordination**: Set up automated incident response channels that create rooms, invite responders, and share critical information during emergencies
- **Meeting Room Automation**: Automatically create temporary meeting rooms, invite participants, and share meeting materials for virtual conferences and discussions
- **Customer Support Integration**: Create dedicated support rooms for customer inquiries, automatically route messages, and manage customer interactions
- **Content Broadcasting**: Distribute announcements, news updates, and important information to multiple Matrix rooms simultaneously
- **File Sharing and Collaboration**: Upload and share files, documents, and media content with team members in organized channels
- **Event-Driven Notifications**: Send automated notifications to Matrix rooms based on external triggers like system alerts, deployment status, or workflow completions
- **Bot Integration and Automation**: Deploy chatbots and automated assistants that can respond to queries, execute commands, and provide information
- **Cross-Platform Messaging**: Bridge Matrix with other communication platforms to unify messaging across different tools and services
- **Privacy-Focused Communication**: Leverage Matrix's decentralized and end-to-end encrypted messaging for secure organizational communication

