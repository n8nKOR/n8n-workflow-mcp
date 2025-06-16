# Mocean Node

## Overview

Send SMS and voice messages via Mocean API for communication automation workflows. Mocean provides reliable SMS and voice messaging services for businesses to reach customers through automated notifications, alerts, and marketing campaigns.

## Credentials

- Name: moceanApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: voice

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| To | string | Recipient phone number in international format | Yes |  |
| From | string | Sender phone number or caller ID | Yes |  |
| Message | string | Text message to convert to speech | Yes |  |
| Language | options | Voice language for text-to-speech conversion | No | en-US |

### Resource: sms

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| To | string | Recipient phone number in international format | Yes |  |
| From | string | Sender phone number or short code | Yes |  |
| Message | string | SMS message content (up to 160 characters) | Yes |  |
| Options | collection | Additional SMS delivery options | No |  |

## UseCases

- **Customer Notification Systems**: Send automated SMS notifications for order confirmations, shipping updates, and delivery alerts to customers
- **Authentication and Security**: Implement two-factor authentication (2FA) by sending SMS verification codes for secure user login processes
- **Marketing Campaigns**: Execute SMS marketing campaigns with promotional messages, special offers, and event announcements to customer lists
- **Emergency Alerts**: Send urgent voice and SMS alerts for critical system notifications, security breaches, or emergency situations
- **Appointment Reminders**: Automate appointment reminder SMS messages for healthcare, service, and consultation scheduling systems
- **Customer Support Integration**: Send SMS updates on support ticket status, resolution notifications, and follow-up communications
- **Voice Broadcasting**: Deliver automated voice messages for announcements, surveys, and important communications to large audiences
- **Payment and Transaction Alerts**: Send SMS notifications for payment confirmations, transaction alerts, and billing reminders
- **System Monitoring Alerts**: Integrate with monitoring systems to send SMS and voice alerts when system issues or downtime occur
- **Event-Driven Messaging**: Trigger SMS and voice messages based on workflow events, user actions, or system conditions

