# MessageBird

## Overview

The MessageBird node allows you to send SMS messages and manage account balance using the MessageBird REST API. MessageBird is a cloud communications platform that provides reliable SMS, voice, and chat APIs for businesses worldwide. This node enables you to integrate SMS capabilities into your n8n workflows for notifications, alerts, marketing campaigns, and two-way messaging.

## Credentials

This node requires MessageBird API credentials:
- **Credential Name**: `messageBirdApi`
- **Required Fields**: 
  - Access Key: Your MessageBird API access key

To obtain API credentials:
1. Sign up for a MessageBird account at https://messagebird.com
2. Log into your MessageBird dashboard
3. Navigate to "Developers" → "API access"
4. Generate a new access key or copy an existing one
5. Copy the access key for use in n8n

## Inputs

- **Main**: SMS message data including recipients, sender ID, and message content

## Outputs

- **Main**: Returns JSON response with SMS delivery status, message ID, balance information, and API response details

## Properties

### Resource: SMS

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| From | string | The originator (sender ID or number) from which to send the message | Yes |
| To | string | All recipients separated by commas | Yes |
| Message | string | The message content to be sent | Yes |
| Additional Fields | collection | Advanced message configuration options | No |

##### Additional Fields Options:

| Field Name | Type | Description | Default |
|---|---|---|---|
| Created Date-Time | string | Message creation timestamp (RFC3339 format) | Current time |
| Datacoding | options | Character encoding (auto/plain/unicode) | auto |
| Gateway | number | SMS route ID for sending | Default route |
| Group IDs | string | Comma-separated group IDs (replaces recipients) | None |
| Message Type (mclass) | options | Normal (0) or Flash (1) message | Normal |
| Reference | string | Client reference for tracking | None |
| Report URL | string | Status report webhook URL | None |
| Scheduled Date-Time | string | Scheduled send time (RFC3339 format) | Immediate |
| Type | options | Message type (sms/binary/flash) | sms |
| Type Details | string | Extra information for binary messages | None |
| Validity | number | Message validity period in seconds | Default |

### Resource: Balance

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| (No additional fields) | - | Retrieves current account balance | - |

## UseCases

- **Customer Notifications**: Send order confirmations, shipping updates, and delivery notifications
- **Authentication & Security**: Deliver OTP codes and two-factor authentication messages
- **Marketing Campaigns**: Send promotional messages, offers, and product announcements to customer lists
- **Appointment Reminders**: Notify customers about upcoming appointments, bookings, or scheduled services
- **System Alerts**: Send critical system alerts, server monitoring notifications, and IT incident reports
- **E-commerce Updates**: Alert customers about cart abandonment, restocked items, and flash sales
- **Event Notifications**: Send event reminders, schedule changes, and important announcements
- **Customer Support**: Send support ticket updates, resolution notifications, and follow-up surveys
- **Financial Alerts**: Notify users about payment due dates, transaction confirmations, and account activities
- **Emergency Communications**: Send urgent alerts, safety notifications, and crisis communications
- **Lead Nurturing**: Send follow-up messages to prospects and qualified leads
- **Account Management**: Check MessageBird balance and monitor usage for billing and planning purposes
