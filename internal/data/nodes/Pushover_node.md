# Pushover

## Overview

The Pushover node allows you to send real-time push notifications to mobile devices and desktop applications using the Pushover API. Pushover is a notification service that delivers instant notifications across iOS, Android, and desktop platforms. This node enables you to integrate push notification capabilities into your n8n workflows for alerts, monitoring, emergency notifications, and instant messaging scenarios.

## Credentials

This node requires Pushover API credentials:
- **Credential Name**: `pushoverApi`
- **Required Fields**: 
  - API Token: Your application's API token from Pushover
  - User Key: Your user key from Pushover dashboard

To obtain API credentials:
1. Create a Pushover account at https://pushover.net
2. Register a new application to get an API token
3. Copy your User Key from your dashboard
4. Add both credentials to n8n

## Inputs

- **Main**: Notification data including message content, priority settings, and delivery options

## Outputs

- **Main**: Returns JSON response with notification delivery status, receipt ID, and any error information

## Properties

### Resource: Message

#### Operation: Push

| Field Name | Type | Description | Required |
|---|---|---|---|
| User Key | String | The recipient's user key from Pushover dashboard | Yes |
| Message | String | The notification message content | Yes |
| Priority | Options | Notification priority level (-2 to 2) | No |
| Retry (Seconds) | Number | Retry interval for emergency priority (minimum 30 seconds) | No* |
| Expire (Seconds) | Number | Expiration timeout for emergency priority (maximum 10800 seconds) | No* |
| Additional Fields | Collection | Advanced notification configuration options | No |

*Required only when Priority is set to Emergency (2)

##### Additional Fields Options:

| Field Name | Type | Description | Default |
|---|---|---|---|
| Title | String | Notification title | App name |
| Sound | Options | Notification sound | Default |
| Timestamp | String | Custom timestamp | Current time |
| HTML | Boolean | Enable HTML formatting | false |
| URL | String | Supplementary URL | None |
| URL Title | String | Title for URL | None |
| Device | String | Target device name | All devices |
| Attachment | Binary | Image attachment | None |



## UseCases

- **System Monitoring**: Send critical alerts for server failures and security breaches
- **Emergency Notifications**: Deliver urgent alerts requiring immediate attention with acknowledgment
- **Application Alerts**: Notify about errors, deployments, and performance issues
- **IoT Device Monitoring**: Alert about sensor readings and device malfunctions
- **Business Process Alerts**: Notify about workflow completions and approval requests
- **Security Notifications**: Send alerts for login attempts and unauthorized access
- **Backup and Maintenance**: Notify about backup completions and system updates
- **E-commerce Alerts**: Send notifications for new orders and payment failures
- **Development Notifications**: Alert about build failures and code deployments
- **Personal Reminders**: Send custom reminders and task notifications
- **Financial Alerts**: Notify about payment confirmations and billing issues
- **Content Management**: Alert about new submissions and publication schedules
