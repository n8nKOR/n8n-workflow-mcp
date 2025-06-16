# Mailjet Node

## Overview

Consume Mailjet API

## Credentials

- Name: mailjetEmailApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: email

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| From Email | string | The title for the email | Yes |  |
| To Email | string | Email address of the recipient. Multiple ones can be separated by comma. | Yes |  |
| Text | string | Plain text message of email | No |  |
| HTML | string | HTML text message of email | No |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | collection | Bcc Email address of the recipient. Multiple ones can be separated by comma. | No |  |
| Variables (JSON) | string | HTML text message of email | No |  |
| Variables | fixedCollection |  | No |  |

#### Operation: sendTemplate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| From Email | string | The title for the email | Yes |  |
| To Email | string | Email address of the recipient. Multiple ones can be separated by comma. | Yes |  |
| Template Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | collection | BCC Recipients of the email separated by , | No |  |
| Variables | fixedCollection |  | No |  |
| Variables (JSON) | string | HTML text message of email | No |  |

### Resource: sms

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| From | string | Customizable sender name. Should be between 3 and 11 characters in length, only alphanumeric characters are allowed. | Yes |  |
| To | string | Message recipient. Should be between 3 and 15 characters in length. The number always starts with a plus sign followed by a country code, followed by the number. Phone numbers are expected to comply with the E.164 format. | Yes |  |
| Text | string |  | Yes |  |

## UseCases

- **Transactional Emails** : Send order confirmations, receipts, and transaction notifications
- **Marketing Campaigns** : Execute email marketing campaigns and newsletters
- **User Registration** : Send welcome emails and account verification messages
- **Password Reset** : Send password reset links and security notifications
- **Event Notifications** : Send event reminders, updates, and confirmations
- **Customer Support** : Send support responses and ticket updates via email
- **SMS Notifications** : Send SMS alerts, reminders, and important notifications
- **Multi-channel Communication** : Combine email and SMS for comprehensive messaging
- **Template-based Messaging** : Use predefined templates for consistent communication
- **Personalized Content** : Send personalized messages using dynamic variables
- **Automated Workflows** : Trigger automated email and SMS sequences
- **Customer Onboarding** : Guide new customers through onboarding processes
- **Abandoned Cart Recovery** : Send cart abandonment emails to recover sales
- **Survey Distribution** : Send surveys and collect customer feedback
- **Appointment Reminders** : Send appointment confirmations and reminders
- **Invoice Delivery** : Send invoices and billing information to customers
- **Newsletter Subscription** : Manage newsletter subscriptions and content delivery
- **Emergency Alerts** : Send urgent notifications via SMS and email
- **Two-Factor Authentication** : Send SMS codes for authentication purposes
- **Status Updates** : Send project updates and status notifications to stakeholders

