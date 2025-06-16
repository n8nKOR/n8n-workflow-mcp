# Pushcut

## Overview

The Pushcut node allows you to send custom notifications to iOS devices using the Pushcut app and API. Pushcut is an iOS app that enables you to create custom notifications, widgets, and automation shortcuts on your iPhone or iPad. This node integrates with Pushcut's server-side capabilities to trigger notifications remotely and pass custom data to your iOS devices as part of your n8n workflows.

## Credentials

This node requires Pushcut API credentials:
- **Credential Name**: `pushcutApi`
- **Required Fields**: 
  - API Key: Your Pushcut API key

To obtain API credentials:
1. Install the Pushcut app on your iOS device from the App Store
2. Open the Pushcut app and create an account
3. Navigate to "Account" → "API" in the app
4. Generate an API key
5. Copy the API key for use in n8n

## Inputs

- **Main**: Data to be sent with the notification, including custom parameters and values

## Outputs

- **Main**: Returns JSON response confirming notification delivery status and any response data from the target devices

## Properties

### Resource: Notification

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| Notification Name or ID | Options | Select from available notifications or specify custom ID | Yes |
| Additional Fields | Collection | Optional notification customization settings | No |

##### Additional Fields Options:

| Field Name | Type | Description | Default |
|---|---|---|---|
| Device Names or IDs | Multi-Select | Target devices to receive notification | All devices |
| Input | String | Custom value passed to the notification action | None |
| Text | String | Override notification text content | App-defined |
| Title | String | Override notification title | App-defined |

**Notification Configuration:**
- **Dynamic Loading**: Available notifications are loaded directly from your Pushcut account
- **Device Targeting**: Choose specific devices or send to all registered devices
- **Custom Content**: Override default notification title and text
- **Action Input**: Pass custom data to Pushcut actions and shortcuts

**Device Management:**
- Automatically loads all devices registered to your Pushcut account
- Supports targeting specific devices by name or ID
- Default behavior sends to all available devices
- Real-time device status and availability

## UseCases

- **Home Automation Control** : Send notifications to control smart home devices through iOS shortcuts
- **Personal Task Management** : Send reminders and task notifications to your personal devices
- **System Monitoring Alerts** : Notify mobile devices about server status, backups, or system issues
- **IoT Device Notifications** : Alert about sensor readings, device status changes, or maintenance needs
- **Workflow Status Updates** : Send progress notifications for long-running processes or batch jobs
- **Emergency Alerts** : Critical system alerts that require immediate attention on mobile devices
- **Content Publishing Notifications** : Alert content creators about new comments, uploads, or reviews
- **E-commerce Order Alerts** : Notify about new orders, payment confirmations, or inventory updates
- **Development Deployment Notifications** : Alert about build completions, deployment status, or errors
- **Meeting and Event Reminders** : Send contextual reminders with meeting details and quick actions
- **Data Backup Confirmations** : Confirm successful backups or alert about backup failures
- **Custom Dashboard Updates** : Send personalized metrics and KPI updates to executive devices
