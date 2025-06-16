# sms77 Node

## Overview

⚠️ **Node Name**: The correct node name is "sms77". This node integrates with the sms77 SMS and voice service platform (formerly known as "seven"). Always use "sms77" when searching for this node in n8n.

**Service Provider**: sms77 (SMS and voice communication platform)
**Node Identifier**: sms77
**Previous Names**: seven (deprecated reference)

Send SMS and make text-to-speech calls

## Credentials

- Name: sms77Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: sms

#### Operation: send

| Field Name | Type | Description | Required |
|---|---|---|---|
| From | string | The caller ID displayed in the receivers display. Max 16 numeric or 11 alphanumeric characters | No |
| To | string | The number of your recipient(s) separated by comma. Can be regular numbers or contact/groups from seven | Yes |
| Message | string | The message to send. Max. 1520 characters | Yes |
| Options | collection | Additional options for SMS sending | No |

#### SMS Options:

| Field Name | Type | Description | Required |
|---|---|---|---|
| Delay | dateTime | Pick a date for time delayed dispatch | No |
| Foreign ID | string | Custom foreign ID returned in DLR callbacks | No |
| Flash | boolean | Send as flash message being displayed directly the receiver's display | No |
| Label | string | Custom label used to group analytics | No |
| Performance Tracking | boolean | Whether to enable performance tracking for URLs found in the message text | No |
| TTL | number | Custom time to live specifying the validity period of a message in minutes | No |

### Resource: voice

#### Operation: send

| Field Name | Type | Description | Required |
|---|---|---|---|
| To | string | The number of your recipient(s) separated by comma. Can be regular numbers or contact/groups from seven | Yes |
| Message | string | The message to send. Max. 1520 characters | Yes |
| Options | collection | Additional options for voice call | No |

#### Voice Options:

| Field Name | Type | Description | Required |
|---|---|---|---|
| From | string | The caller ID. Please use only verified sender IDs, one of your virtual inbound numbers or one of our shared virtual numbers | No |

## UseCases

- **SMS Notifications** : Send SMS alerts, reminders, and important notifications to customers
- **Two-Factor Authentication** : Send SMS verification codes for secure authentication
- **Marketing Campaigns** : Execute SMS marketing campaigns and promotional messages
- **Emergency Alerts** : Send urgent notifications and emergency alerts via SMS
- **Appointment Reminders** : Send appointment confirmations and reminder SMS messages
- **Order Updates** : Send order status updates and delivery notifications
- **Voice Notifications** : Make automated voice calls for important announcements
- **Customer Support** : Send support responses and issue updates via SMS
- **Event Notifications** : Send event reminders and updates to attendees
- **System Alerts** : Send system status updates and technical alerts
- **Payment Reminders** : Send payment due reminders and billing notifications
- **Security Alerts** : Send security-related notifications and alerts
- **Survey Distribution** : Send survey invitations and feedback requests
- **Automated Workflows** : Trigger automated SMS and voice call sequences
- **Multi-language Support** : Send messages in multiple languages for global reach
- **Scheduled Messaging** : Schedule SMS and voice messages for optimal delivery times
- **Flash Messages** : Send urgent flash SMS messages for immediate attention
- **Performance Tracking** : Track message delivery and engagement metrics
- **International Messaging** : Send SMS and make voice calls internationally
- **Custom Labeling** : Organize and track messages with custom labels and analytics

