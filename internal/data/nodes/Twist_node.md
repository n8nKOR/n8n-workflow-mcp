# Twist Node

## Overview

Consume Twist API

## Credentials

- Name: twistOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: channel

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | The ID of the workspace. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string | The name of the channel | Yes |  |
| Additional Fields | collection | The color of the channel | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | The ID of the workspace. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether only archived conversations are returned | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Channel ID | string | The ID of the channel | Yes |  |
| Update Fields | collection | The color of the channel | No |  |

### Resource: comment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Thread ID | string | The ID of the thread | Yes |  |
| Content | string | The content of the comment | Yes |  |
| Additional Fields | collection | The action of the button | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Thread ID | string | The ID of the channel | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether only the IDs of the comments are returned | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comment ID | string | The ID of the comment | Yes |  |
| Update Fields | collection | The action of the button | No |  |

### Resource: messageConversation

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | The ID of the workspace. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Conversation Name or ID | options | The ID of the conversation. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Content | string | The content of the new message. Mentions can be used as <code>[Name](twist-mention://user_id)</code> for users or <code>[Group name](twist-group-mention://group_id)</code> for groups. | No |  |
| Additional Fields | collection | Other options to set | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | The ID of the workspace. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Conversation Name or ID | options | The ID of the conversation. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Additional Fields | collection | Other options to set | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation Message ID | string | The ID of the conversation message | Yes |  |
| Update Fields | collection | Other options to set | No |  |

### Resource: thread

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Channel ID | string | The ID of the channel | Yes |  |
| Title | string | The title of the new thread (1 < length < 300) | Yes |  |
| Content | string | The content of the thread | Yes |  |
| Additional Fields | collection | The action of the button | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Channel ID | string | The ID of the channel | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether only the IDs of the threads are returned | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Thread ID | string | The ID of the thread | Yes |  |
| Update Fields | collection | The action of the button | No |  |

## UseCases

- [Team Communication] : Facilitate organized team communication with threaded conversations
- [Project Collaboration] : Coordinate project activities and discussions in dedicated channels
- [Asynchronous Communication] : Enable asynchronous team communication across time zones
- [Knowledge Management] : Organize and preserve team knowledge through searchable threads
- [Remote Work] : Support remote teams with structured communication tools
- [Thread Management] : Manage conversation threads for better organization and clarity
- [Team Coordination] : Coordinate team activities and share updates efficiently
- [Integration Automation] : Connect Twist with other productivity and project tools
- [Notification Management] : Manage notifications and stay updated on important discussions
- [Cross-Team Communication] : Facilitate communication between different departments
- [Decision Making] : Use structured threads for collaborative decision making
- [Status Updates] : Share project status updates and progress reports
- [Meeting Coordination] : Coordinate meetings and share agendas through team channels
- [Customer Support] : Provide customer support through organized team communication
- [Documentation] : Document processes and decisions through persistent conversation threads

