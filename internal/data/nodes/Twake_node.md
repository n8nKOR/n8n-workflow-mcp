# Twake Node

## Overview

Consume Twake API

## Credentials

- Name: twakeCloudApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Message

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | Resource to operate on | Yes |
| operation | options | Operation to perform on the resource | Yes |
| channelId | options | Channel's ID to send message to | Yes |
| content | string | Message content to send | Yes |
| senderIcon | string | URL of the image/icon for the sender | No |
| senderName | string | Name to display as the message sender | No |

## UseCases

- [Team Communication]: Send automated messages to Twake channels for team notifications, alerts, and workflow updates
- [Bot Integration]: Create interactive bots that respond to events and send contextual messages to team channels
- [Alert Systems]: Deliver system alerts, monitoring notifications, and status updates to relevant team channels

