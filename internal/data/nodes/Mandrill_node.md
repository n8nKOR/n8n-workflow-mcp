# Mandrill Node

## Overview

Send transactional emails using the Mandrill API. This node enables sending of HTML emails and template-based messages through Mailchimp's transactional email service, with comprehensive tracking and delivery options.

## Credentials

- Name: mandrillApi, Required: Yes

Requires a Mandrill API key from your Mailchimp account:
1. Log into your Mailchimp account
2. Navigate to Settings > API Keys
3. Generate a Mandrill API key
4. Add the key to n8n credentials

## Inputs

- Main: Input data for email sending (recipient lists, template variables, etc.)

## Outputs

- Main: Delivery status and message information

## Properties

### Resource: Mandrill Email

#### Operation: Send Template

| Field Name | Type | Description | Required |
|---|---|---|---|
| Template Name or ID | Options | Mandrill template to use | Yes |
| From Email | String | Sender email address | Yes |
| To Email | String | Recipient email addresses (comma-separated) | Yes |
| JSON Parameters | Boolean | Use JSON for advanced configuration | No |

#### Operation: Send HTML

| Field Name | Type | Description | Required |
|---|---|---|---|
| From Email | String | Sender email address | Yes |
| To Email | String | Recipient email addresses (comma-separated) | Yes |
| HTML | String | HTML email content | Yes |
| Subject | String | Email subject line | Yes |
| JSON Parameters | Boolean | Use JSON for advanced configuration | No |
## UseCases

- **Mandrill Integration** : Use Mandrill for workflow automation and data processing
- **Data Processing** : Process and transform data using Mandrill capabilities
- **Workflow Enhancement** : Enhance workflows with Mandrill functionality
- **Automation Tasks** : Automate repetitive tasks using Mandrill features

## Operations

### Send Template

Send emails using Mandrill templates with dynamic content replacement.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Template Name or ID | Options | Mandrill template to use | Yes | - |
| From Email | String | Sender email address | Yes | - |
| To Email | String | Recipient email addresses (comma-separated) | Yes | - |
| JSON Parameters | Boolean | Use JSON for advanced configuration | No | false |

**Template Selection:**
- Dynamically loads available templates from your Mandrill account
- Choose from dropdown or specify template ID using expressions
- Templates must be created in your Mailchimp/Mandrill dashboard

### Send HTML

Send custom HTML emails with full content control.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| From Email | String | Sender email address | Yes | - |
| To Email | String | Recipient email addresses (comma-separated) | Yes | - |
| JSON Parameters | Boolean | Use JSON for advanced configuration | No | false |

## Advanced Options

### Email Configuration

| Field Name | Type | Description | Default |
|---|---|---|---|
| From Name | String | Sender display name | - |
| HTML | String | HTML email content (for Send HTML) | - |
| Subject | String | Email subject line | - |
| Text | String | Plain text version | - |

### Delivery Options

| Field Name | Type | Description | Default |
|---|---|---|---|
| Important | Boolean | Mark email as important | false |
| Track Opens | Boolean | Track email opens | true |
| Track Clicks | Boolean | Track link clicks | true |
| Auto Text | Boolean | Auto-generate text from HTML | false |
| Auto HTML | Boolean | Auto-generate HTML from text | false |
| Async | Boolean | Enable background sending for bulk emails | false |

### Advanced Settings

| Field Name | Type | Description | Default |
|---|---|---|---|
| BCC Address | String | Blind copy address | - |
| Subaccount | String | Mandrill subaccount | - |
| IP Pool | String | Dedicated IP pool | - |
| Tags | String | Message tags for organization | - |
| Send At | String | Schedule delivery time | - |

### Google Analytics Integration

| Field Name | Type | Description | Default |
|---|---|---|---|
| Google Analytics Campaign | String | UTM campaign parameter | - |
| Google Analytics Domains | String | Domains for automatic UTM tracking | - |

### Domain Configuration

| Field Name | Type | Description | Default |
|---|---|---|---|
| Tracking Domain | String | Custom tracking domain | - |
| Signing Domain | String | DKIM signing domain | - |
| Return Path Domain | String | Bounce handling domain | - |

### Content Processing

| Field Name | Type | Description | Default |
|---|---|---|---|
| Inline CSS | Boolean | Inline CSS styles | false |
| URL Strip QS | Boolean | Remove query strings from URLs | false |
| Preserve Recipients | Boolean | Show all recipients in headers | false |
| View Content Link | Boolean | Add web view link | false |

## Template Features

### Merge Variables
- **Global Merge Variables**: Apply to all recipients
- **Per-Recipient Variables**: Unique data for each recipient
- **Handlebars Syntax**: Use `{{variable_name}}` in templates
- **Default Values**: Specify fallback values for missing data

### Content Types
- **HTML Templates**: Rich content with styling
- **Text Templates**: Plain text alternatives
- **Subject Templates**: Dynamic subject lines
- **Header Templates**: Custom email headers

### Template Management
- **Dynamic Loading**: Available templates loaded from Mandrill
- **Version Control**: Template versioning support
- **Preview**: Template preview capabilities
- **Testing**: Send test emails before deployment

## Recipient Management

### Email Formats
- **Simple Email**: `user@example.com`
- **Named Email**: `John Doe <user@example.com>`
- **Multiple Recipients**: Comma-separated list
- **CC/BCC Support**: Additional recipient types

### Recipient Types
- **To**: Primary recipients
- **CC**: Carbon copy recipients  
- **BCC**: Blind carbon copy recipients
- **Reply-To**: Custom reply address

## Tracking and Analytics

### Email Tracking
- **Open Tracking**: Monitor email opens with pixel tracking
- **Click Tracking**: Track link clicks with redirect URLs
- **Bounce Tracking**: Handle bounced emails
- **Spam Complaints**: Monitor spam reports

### Google Analytics
- **UTM Parameters**: Automatic UTM parameter injection
- **Campaign Tracking**: Link to Google Analytics campaigns
- **Domain Filtering**: Specify domains for tracking
- **Custom Parameters**: Additional tracking parameters

### Delivery Status
- **Sent**: Successfully sent to Mandrill
- **Queued**: Queued for delivery (async mode)
- **Scheduled**: Scheduled for future delivery
- **Rejected**: Rejected due to invalid data

## Error Handling

### Common Errors
- **Invalid API Key**: Check Mandrill credentials
- **Template Not Found**: Verify template exists and is published
- **Invalid Email**: Check email format and validity
- **Rate Limits**: Handle API rate limiting

### Validation
- **Email Format**: Automatic email validation
- **Template Validation**: Check template availability
- **Content Validation**: Validate HTML and text content
- **Merge Variable Validation**: Check required variables

## Use Cases

### Transactional Emails
- **Order Confirmations**: E-commerce purchase confirmations
- **Account Notifications**: Password resets, welcome emails
- **System Alerts**: Service notifications and updates
- **Receipts**: Payment and subscription receipts

### Marketing Communications
- **Promotional Emails**: Product announcements and offers
- **Newsletter**: Regular content distribution
- **Event Invitations**: Webinar and event notifications
- **Follow-up Sequences**: Automated email sequences

### Business Automation
- **Workflow Notifications**: Internal team notifications
- **Customer Support**: Automated response emails
- **Lead Nurturing**: Educational email sequences
- **Abandoned Cart**: E-commerce recovery emails

### Integration Scenarios
- **CRM Integration**: Sync with customer relationship management
- **E-commerce Platforms**: Order and shipping notifications
- **User Onboarding**: Multi-step welcome sequences
- **Event-Driven**: Trigger emails based on user actions

### Advanced Implementations
- **A/B Testing**: Template and subject line testing
- **Personalization**: Dynamic content based on user data
- **Localization**: Multi-language email support
- **Compliance**: GDPR and privacy-compliant communications

