# Brevo Node

## Overview

The Brevo node (formerly SendinBlue) allows you to consume the Brevo API for email marketing and automation. Brevo is a comprehensive digital marketing platform providing email marketing, SMS, chat, and CRM tools. This node enables email sending, contact management, and marketing automation within n8n workflows.

## Credentials

This node requires Brevo API credentials:
- **API Key**: Your Brevo API key

To obtain API credentials:
1. Log into your Brevo account
2. Navigate to Settings → API Keys
3. Generate a new API key
4. Copy the key for use in n8n

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Contact

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | string | Contact email address | Yes |
| Attributes | collection | Contact attributes and values | No |
| ListIds | multiOptions | Lists to add contact to | No |
| UpdateEnabled | boolean | Update if contact exists | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | string | Contact email address to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | string | Contact email address to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Limit | number | Maximum number of contacts to return | No |
| Offset | number | Number of contacts to skip | No |
| ModifiedSince | dateTime | Filter by modification date | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | string | Contact email address to update | Yes |
| Attributes | collection | Updated contact attributes | No |
| ListIds | multiOptions | Lists to add/remove contact | No |

### Resource: Email

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| To | collection | Recipient email addresses | Yes |
| Subject | string | Email subject line | Yes |
| HtmlContent | string | HTML email content | No |
| TextContent | string | Plain text email content | No |
| Sender | object | Sender name and email | Yes |
| ReplyTo | object | Reply-to address | No |
| Attachment | collection | Email attachments | No |

#### Operation: Send Template

| Field Name | Type | Description | Required |
|---|---|---|---|
| TemplateId | number | Email template ID | Yes |
| To | collection | Recipient email addresses | Yes |
| Sender | object | Sender name and email | Yes |
| Params | object | Template parameters | No |

### Resource: Contact Attribute

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | string | Attribute name | Yes |
| Type | options | Attribute type (text, number, date, etc.) | Yes |
| Category | options | Attribute category | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | string | Attribute name to delete | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Category | options | Filter by category | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | string | Attribute name to update | Yes |
| Value | string | New attribute value | Yes |

### Resource: Sender

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | string | Sender name | Yes |
| Email | string | Sender email address | Yes |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Id | number | Sender ID to delete | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | string | Filter by domain | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Id | number | Sender ID to update | Yes |
| Name | string | Updated sender name | No |
| Email | string | Updated sender email | No |

## UseCases

- **Email Marketing Campaigns** : Send targeted email campaigns to segmented contact lists
- **Transactional Emails** : Automate order confirmations, password resets, and notifications
- **Contact Management** : Sync customer data between CRM systems and Brevo
- **Marketing Automation** : Create workflows based on user behavior and engagement
- **A/B Testing** : Test different email templates and content variations
- **Newsletter Distribution** : Send regular newsletters to subscriber lists
- **Customer Onboarding** : Automate customer welcome series and onboarding emails
- **Event Notifications** : Send event invitations and reminders to attendees
- **Lead Nurturing** : Develop leads through automated email sequences
- **Customer Support** : Send support updates and communication via email
