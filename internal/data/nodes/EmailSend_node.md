# Send Email Node

## Overview

Sends an email using SMTP protocol

## Credentials

- Name: smtp, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Email

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| From Email | String | The email address to send from | Yes |
| To Email | String | Email addresses of the recipients (comma separated) | Yes |
| Subject | String | Subject line of the email | Yes |
| Message | String | Plain text message body | No |
| HTML | String | HTML message body | No |
| CC | String | Carbon copy recipients (comma separated) | No |
| BCC | String | Blind carbon copy recipients (comma separated) | No |
| Reply To | String | Reply-to email address | No |
| Attachments | FixedCollection | File attachments to include | No |

#### Operation: Send and Wait for Response

| Field Name | Type | Description | Required |
|---|---|---|---|
| From Email | String | The email address to send from | Yes |
| To Email | String | Email addresses of the recipients (comma separated) | Yes |
| Subject | String | Subject line of the email | Yes |
| Message | String | Plain text message body | No |
| HTML | String | HTML message body | No |
| CC | String | Carbon copy recipients (comma separated) | No |
| BCC | String | Blind carbon copy recipients (comma separated) | No |
| Reply To | String | Reply-to email address | No |
| Attachments | FixedCollection | File attachments to include | No |
| Webhook URL | String | URL to receive responses | Yes |
| Response Timeout | Number | Time to wait for response (minutes) | No |

### Additional Options

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| CC | string | Carbon copy recipients (comma separated) | No | - |
| BCC | string | Blind carbon copy recipients (comma separated) | No | - |
| Reply To | string | Reply-to email address | No | - |
| Attachments | fixedCollection | File attachments to include | No | - |

### Send and Wait for Response Options

This operation extends the basic send functionality with webhook capabilities to receive responses.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webhook URL | string | URL to receive responses | Yes | - |
| Response Timeout | number | Time to wait for response (minutes) | No | 60 |

### Usage Notes
- Requires SMTP credentials to be configured
- Supports both plain text and HTML email formats
- Can send to multiple recipients using comma-separated addresses
- Supports file attachments through binary data
- The "Send and Wait" operation creates a webhook to capture email responses
- HTML emails can include inline images and styling
- CC and BCC functionality available for broader distribution

## UseCases

- [Customer Notifications] : Send order confirmations, shipping updates, and account notifications to customers
- [Internal Communications] : Send reports, updates, and announcements to team members
- [Invoice and Billing] : Email invoices, receipts, and payment reminders to customers
- [Support Ticketing] : Send support ticket updates and responses to customers
- [Welcome Emails] : Send onboarding and welcome messages to new users
- [Newsletter Distribution] : Send regular newsletters and marketing content to subscribers
- [Promotional Campaigns] : Send special offers and promotional emails to customers
- [Event Invitations] : Send event announcements and invitations to attendees
- [Product Updates] : Notify customers about new features or products via email
- [Survey Requests] : Send customer satisfaction surveys and feedback requests
- [Alert Systems] : Send system alerts and monitoring notifications to administrators
- [Report Distribution] : Automatically send generated reports to stakeholders
- [Scheduled Communications] : Send time-based messages and reminders
- [Workflow Notifications] : Notify team members about workflow status changes
- [Data Backup Notifications] : Send backup completion and status reports
- [CRM Integration] : Send emails based on CRM triggers and customer actions
- [E-commerce Automation] : Send cart abandonment emails and purchase confirmations
- [Form Processing] : Send confirmation emails for form submissions
- [User Registration] : Send verification and activation emails to new users
- [Password Reset] : Send password reset links and security notifications
- [Customer Service] : Send emails and wait for customer responses
- [Approval Workflows] : Send approval requests and capture responses via email
- [Survey Collection] : Send surveys and collect responses via email
- [Feedback Loops] : Create interactive email workflows with response handling
- [Confirmation Processes] : Send confirmation requests and process replies
- [HTML Newsletters] : Send rich HTML emails with branding and styling
- [Document Delivery] : Send PDFs, reports, and other attachments via email
- [Multi-Recipient Broadcasting] : Send emails to multiple recipients with CC/BCC functionality
- [Personalized Content] : Send customized emails based on user data and preferences
- [Template-Based Emails] : Use dynamic templates for consistent messaging across campaigns

