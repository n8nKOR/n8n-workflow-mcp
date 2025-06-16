# Spontit Node

## Overview

Consume Spontit API

## Credentials

- Name: spontitApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: push

#### Operation: create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Content | string | To provide text in a push, supply one of either "content" or "pushContent" (or both). Limited to 2500 characters | No |
| Additional Fields | collection | Additional options for push notification | No |

#### Additional Fields Options:

| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel Name | string | The name of a channel you created. If you have not yet created a channel, simply don't provide this value and the push will be sent to your main account | No |
| Expiration Stamp | dateTime | A Unix timestamp. When to automatically expire your push notification. The default is 10 days after pushing | No |
| iOS DeepLink | string | An iOS deep link. Use this to deep link into other apps | No |
| Link | string | A link that can be attached to your push. Must be a valid URL | No |
| Open In Home Feed | boolean | Whether the notification opens to the home feed or to a standalone page with the notification | No |
| Open Link In App | boolean | Whether to open the provided link within the iOS app or in Safari | No |
| Push To Emails | string | Emails (strings) to whom to send the notification | No |
| Push To Followers | string | User IDs (strings) to whom to send the notification | No |
| Push To Phone Numbers | string | Phone numbers (strings) to whom to send the notification | No |
| Schedule | dateTime | A Unix timestamp. Schedule a push to be sent at a later date and time | No |
| Subtitle | string | The subtitle of your push. Limited to 20 characters. Only appears on iOS devices | No |
| Title | string | The title of push. Appears in bold at the top. Limited to 100 characters | No |

## UseCases

- **Push Notification Campaigns** : Send targeted push notifications to mobile app users
- **Real-time Alerts** : Send real-time alerts and notifications for urgent events
- **User Engagement** : Engage users with personalized notifications and messages
- **Event Notifications** : Notify users about events, updates, and important information
- **Marketing Automation** : Automate marketing campaigns through push notifications
- **System Monitoring Alerts** : Send system monitoring alerts and status updates
- **Customer Communication** : Communicate with customers through mobile notifications
- **Promotional Messaging** : Send promotional offers and deals to app users
- **News and Updates** : Distribute news updates and announcements to subscribers
- **Emergency Notifications** : Send emergency alerts and critical notifications

