# Zulip Node

## Overview

Consume Zulip API

## Credentials

- Name: zulipApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: message

#### Operation: sendPrivate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| To | multiOptions | The destination stream, or a comma-separated list containing the usernames (emails) of the recipients. Choose from the list, or specify IDs using an <a href= | Yes |  |
| Content | string | The content of the message | Yes |  |

#### Operation: sendStream

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Stream Name or ID | options | The destination stream, or a comma-separated list containing the usernames (emails) of the recipients. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Topic Name or ID | options | The topic of the message. Only required if type is stream, ignored otherwise. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Content | string | The content of the message | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | Unique identifier for the message | Yes |  |
| Update Fields | collection | The content of the message | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | Unique identifier for the message | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | Unique identifier for the message | Yes |  |

#### Operation: updateFile

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Put Output File in Field | string |  | Yes |  |

### Resource: stream

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | JSON format parameters for stream creation | No |  |
| Subscriptions | fixedCollection | A list of dictionaries containing the the key name and value specifying the name of the stream to subscribe. If the stream does not exist a new stream is created. | Yes |  |
| Additional Fields | collection | If announce is True and one of the streams specified in subscriptions has to be created (i.e. doesnt exist to begin with), an announcement will be made notifying that a new stream was created. | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Whether to include all active streams. The user must have administrative privileges to use this parameter. | No |  |

#### Operation: getSubscribed

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Whether each returned stream object should include a subscribers field containing a list of the user IDs of its subscribers | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Stream ID | string | ID of stream to update | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | JSON format parameters for stream creation | No |  |
| Additional Fields | collection | Whether the stream is limited to announcements | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Stream ID | string | ID of stream to delete | Yes |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | The email address of the new user | Yes |  |
| Full Name | string | The full name of the new user | Yes |  |
| Password | string | The password of the new user | Yes |  |
| Short Name | string | The short name of the new user. Not user-visible. | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | The ID of user to get | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | The ID of user to update | Yes |  |
| Additional Fields | collection | The users full name | Yes |  |

#### Operation: deactivate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | The ID of user to deactivate | Yes |  |

## UseCases

- [Team Communication] : Facilitate organized team communication with topics and streams
- [Project Collaboration] : Coordinate project activities and discussions in dedicated streams
- [Automated Notifications] : Send automated alerts and notifications to team channels
- [Integration Automation] : Connect Zulip with other tools for workflow automation
- [Customer Support] : Provide customer support through organized chat channels
- [Meeting Coordination] : Schedule and coordinate team meetings and discussions
- [Status Updates] : Share project status updates and progress reports with teams
- [Knowledge Sharing] : Share knowledge and documentation through searchable chat history
- [Incident Management] : Coordinate incident response and communication during outages
- [Code Review] : Discuss code changes and reviews in development-focused streams
- [Cross-Team Coordination] : Facilitate communication between different departments
- [Event Management] : Coordinate events and activities through dedicated channels
- [Bot Integration] : Deploy chatbots for automated responses and information sharing
- [Archive Management] : Access and search historical conversations and decisions
- [Remote Work] : Enable effective communication for distributed and remote teams

