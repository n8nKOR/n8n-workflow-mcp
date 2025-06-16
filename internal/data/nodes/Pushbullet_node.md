# Pushbullet Node

## Overview

Consume Pushbullet API

## Credentials

- Name: pushbulletOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: push

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Type | options |  | Yes |  |
| Title | string | Title of the push | Yes |  |
| Body | string | Body of the push | Yes |  |
| URL | string | URL of the push | Yes |  |
| Input Binary Field | string |  | Yes |  |
| Target | options | Send the push to all subscribers to your channel that has this tag | Yes |  |
| Value | string | The value to be set depending on the target selected. For example, if the target selected is email then this field would take the email address of the person you are trying to send the push to. | Yes |  |
| Value Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Push ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Don | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Push ID | string |  | Yes |  |
| Dismissed | boolean | Whether to mark a push as having been dismissed by the user, will cause any notifications for the push to be hidden if possible | Yes |  |

## UseCases

- **Cross-Device Notifications** : Send notifications and messages across multiple devices and platforms
- **System Alerts** : Send system alerts and monitoring notifications to mobile devices
- **File Sharing** : Share files and documents between devices and team members
- **Remote Notifications** : Send remote notifications for server events and system status
- **Team Communication** : Send quick messages and updates to team members
- **IoT Device Alerts** : Send alerts from IoT devices and sensors to mobile devices
- **Automation Notifications** : Notify users about workflow completions and automation results
- **Emergency Alerts** : Send urgent notifications and emergency alerts to devices
- **Link Sharing** : Share links and URLs between devices for easy access
- **Status Updates** : Send status updates and progress notifications for long-running processes

