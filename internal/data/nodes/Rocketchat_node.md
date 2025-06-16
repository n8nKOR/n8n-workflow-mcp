# RocketChat

## Overview

The RocketChat node allows you to send messages to channels and direct messages using the RocketChat API. RocketChat is an open-source team communication platform that provides messaging, file sharing, voice and video calling capabilities. This node enables you to integrate RocketChat messaging into your n8n workflows for team notifications, alerts, automated messages, and rich media communications.

## Credentials

This node requires RocketChat API credentials:
- **Credential Name**: `rocketchatApi`
- **Required Fields**: 
  - User: Your RocketChat username
  - Password: Your RocketChat password
  - Domain: Your RocketChat server URL

To obtain API credentials:
1. Access your RocketChat server
2. Log in with your account credentials
3. Ensure your account has API access permissions
4. Use your login credentials for the node

## Inputs

- **Main**: Message data including channel information, text content, and attachment details

## Outputs

- **Main**: Returns JSON response with message delivery status, message ID, and timestamp information

## Properties

### Resource: Chat

#### Operation: Post Message

| Field Name | Type | Description | Required |
|---|---|---|---|
| Channel | String | Channel name with prefix (#channel, @user) | Yes |
| Text | String | Message text content | No |
| JSON Parameters | Boolean | Use JSON format for attachments | No |
| Additional Fields | Collection | Message customization options | No |

##### Additional Fields Options:

| Field Name | Type | Description | Default |
|---|---|---|---|
| Alias | String | Custom display name for the message sender | None |
| Avatar | String | Custom avatar image URL | None |
| Emoji | String | Custom emoji for sender display | None |
| Attachments | Collection | Rich message attachments | None |

##### Attachment Options:

| Field Name | Type | Description |
|---|---|---|
| Color | String | Attachment border color (hex) |
| Text | String | Attachment text content |
| Timestamp | String | Message timestamp |
| Thumbnail URL | String | Small thumbnail image |
| Message Link | String | Link to related message |
| Author Name | String | Author display name |
| Author Link | String | Author profile link |
| Author Icon | String | Author avatar URL |
| Title | String | Attachment title |
| Title Link | String | Title hyperlink |
| Title Link Download | Boolean | Enable title download |
| Image URL | String | Full-size image attachment |
| Audio URL | String | Audio file attachment |
| Video URL | String | Video file attachment |
| Fields | Array | Custom field data |

## UseCases

- **System Monitoring** : Send alerts about server status, application errors, and performance issues to IT channels
- **DevOps Notifications** : Notify teams about deployments, build completions, and pipeline status updates
- **Customer Support** : Alert support teams about new tickets, escalations, and urgent customer issues
- **Sales Alerts** : Notify sales teams about new leads, deal updates, and quota achievements
- **Project Management** : Send updates about task completions, deadline reminders, and milestone achievements
- **E-commerce Notifications** : Alert teams about new orders, payment issues, and inventory alerts
- **Content Publishing** : Notify content teams about new submissions, approvals, and publication schedules
- **Event Management** : Send reminders about meetings, events, and important announcements
- **Security Alerts** : Notify security teams about login attempts, breaches, and suspicious activities
- **Marketing Campaigns** : Send campaign updates, performance metrics, and A/B test results
- **Financial Reporting** : Share daily/weekly reports, budget alerts, and financial milestone notifications
- **Team Collaboration** : Facilitate automated workflows that keep teams informed about relevant business processes
