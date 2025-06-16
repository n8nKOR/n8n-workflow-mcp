# Line

## Overview

The Line node allows you to send notifications through the LINE Notify API to LINE users and groups. This node enables you to integrate LINE messaging capabilities into your n8n workflows, supporting text messages, image attachments, and stickers. 

⚠️ **Important Notice**: LINE Notify service will be discontinued from April 1st, 2025. Plan accordingly for alternative notification methods.

## Credentials

This node requires LINE Notify OAuth2 API credentials:
- **Credential Name**: `lineNotifyOAuth2Api`
- **Authentication Type**: OAuth2
- **Required Fields**: 
  - Client ID: Your LINE Notify application client ID
  - Client Secret: Your LINE Notify application client secret
  - Access Token: OAuth2 access token for API access

To obtain API credentials:
1. Visit LINE Notify website at https://notify-bot.line.me/
2. Log in with your LINE account
3. Go to "My page" and create a new service
4. Configure OAuth2 settings and get your client credentials
5. Complete OAuth2 flow to obtain access token

## Inputs

- **Main**: Data for notification content and configuration
- **Binary Data**: Optional image files to attach to notifications

## Outputs

- **Main**: Returns JSON response with notification delivery status and API response details

## Properties

### Resource: Notification

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| Message | String | Text message content to send | Yes |
| Additional Fields | Collection | Optional enhancement settings | No |

##### Additional Fields Options:

| Field Name | Type | Description | Default |
|---|---|---|---|
| Image | String/Binary | Image to attach (URL or binary data) | None |
| Image Thumbnail | String/Binary | Thumbnail image (max 240×240px JPEG) | None |
| Full Size Image | String/Binary | Full-size image (max 2048×2048px JPEG) | None |
| Sticker Package ID | Number | LINE sticker package identifier | None |
| Sticker ID | Number | Specific sticker identifier within package | None |
| Notification Disabled | Boolean | Disable push notification (silent delivery) | false |

**Image Requirements:**
- **Format**: JPEG only
- **Full Size**: Maximum 2048×2048 pixels
- **Thumbnail**: Maximum 240×240 pixels
- **Source**: Binary data upload or HTTP/HTTPS URL

**Sticker Integration:**
- Use official LINE sticker packages
- Requires both package ID and sticker ID
- Available sticker packages can be found in LINE documentation

## API Details

**Endpoint**: `https://notify-api.line.me/api/notify`
**Method**: POST
**Authentication**: OAuth2 Bearer token
**Content Type**: Form data (for file uploads)

**Response Format:**
```json
{
  "status": 200,
  "message": "ok"
}
```

## UseCases

- **System Monitoring Alerts**: Send critical system alerts and monitoring notifications to LINE groups
- **Customer Service Notifications**: Notify customer service teams about new tickets or urgent issues
- **E-commerce Order Updates**: Send order confirmation, shipping, and delivery notifications to customers
- **Task Management Alerts**: Notify team members about task assignments, deadlines, and project updates
- **Marketing Campaign Updates**: Send promotional messages and campaign updates to subscribers
- **IoT Device Notifications**: Alert users about smart home device status changes or sensor readings
- **Backup and Maintenance Alerts**: Notify IT teams about backup completion, system maintenance, or failures
- **Sales Lead Notifications**: Alert sales teams about new leads or important customer interactions
- **Content Publishing Alerts**: Notify content teams when new articles are published or require review
- **Error and Exception Reporting**: Send application error reports and exception notifications to development teams
- **Event Reminder System**: Send meeting reminders, appointment confirmations, and event notifications
- **Financial Transaction Alerts**: Notify users about payment confirmations, invoice generation, or billing issues
