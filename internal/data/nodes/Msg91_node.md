# MSG91

## Overview

The MSG91 node allows you to send transactional SMS messages using the MSG91 SMS gateway service. MSG91 is a cloud-based communication platform that provides reliable SMS delivery services globally. This node enables you to integrate SMS capabilities into your n8n workflows for notifications, alerts, OTP delivery, and marketing communications.

## Credentials

This node requires MSG91 API credentials:
- **Credential Name**: `msg91Api`
- **Required Fields**: 
  - Authentication Key: Your MSG91 API authentication key

To obtain API credentials:
1. Sign up for a MSG91 account at https://msg91.com
2. Log into your MSG91 dashboard
3. Navigate to "API" or "Settings" section
4. Generate or copy your Authentication Key (Auth Key)
5. Copy the authentication key for use in n8n

## Inputs

- **Main**: SMS message data including recipient numbers, sender ID, and message content

## Outputs

- **Main**: Returns JSON response with SMS delivery status, request ID, and API response details

## Properties

### Resource: SMS

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| Sender ID | String | The sender ID or number from which to send the message | Yes |
| To | String | The recipient number with country code | Yes |
| Message | String | The SMS message content to send | Yes |

**Parameter Details:**

**Sender ID:**
- Must be approved by MSG91 for your account
- Can be alphanumeric (up to 6 characters) or numeric (up to 15 digits)
- Alphanumeric sender IDs don't support replies
- Numeric sender IDs support two-way messaging

**To (Recipient Number):**
- Must include country code (e.g., +1234567890)
- Supports international numbers
- Format: +[country_code][phone_number]
- Example: +911234567890 for India

**Message:**
- Maximum 160 characters for standard SMS
- Longer messages will be split into multiple SMS units
- Supports Unicode characters for international languages
- Special characters may affect character count

## UseCases

- **OTP Verification** : Send one-time passwords for user authentication and account verification
- **Transaction Alerts** : Notify customers about payment confirmations, order updates, and account activities
- **System Monitoring** : Send critical system alerts and server status notifications to administrators
- **Appointment Reminders** : Send healthcare, service, or meeting reminders to customers
- **Marketing Campaigns** : Deliver promotional messages and special offers to subscriber lists
- **Emergency Notifications** : Send urgent alerts for security breaches, system failures, or critical events
- **E-commerce Updates** : Notify customers about order status, shipping updates, and delivery confirmations
- **Customer Support** : Send support ticket updates, resolution notifications, and follow-up messages
- **Event Notifications** : Alert attendees about event changes, cancellations, or important updates
- **Backup and Maintenance Alerts** : Notify IT teams about scheduled maintenance, backup completion, or system issues
- **Lead Nurturing** : Send follow-up messages to potential customers and qualified leads
- **Survey and Feedback Requests** : Send survey links and feedback requests after service completion
