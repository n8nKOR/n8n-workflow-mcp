# Mailgun

## Overview

The Mailgun node enables sending emails through the Mailgun email delivery service. Mailgun is a powerful email automation service that provides reliable email delivery, detailed analytics, and advanced features for transactional and marketing emails. This node supports sending emails with HTML/text content, attachments, and multiple recipients, making it ideal for automated email workflows.

## Credentials

- **Credential Name**: `mailgunApi`
- **Required**: Yes
- **Configuration**: 
  - API Key: Your Mailgun API key
  - API Domain: Your Mailgun API domain (e.g., api.mailgun.net)
  - Email Domain: Your verified sending domain

## Inputs

- **Main**: Input data from previous nodes, which can be used to dynamically populate email content and recipients

## Outputs

- **Main**: Returns the response from Mailgun API, including message ID and delivery status

## Properties

### Resource: Email

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| From Email | String | Email address of the sender, optionally with display name | Yes |
| To Email | String | Email address(es) of the recipient(s). Multiple addresses can be separated by comma | Yes |
| Subject | String | Subject line of the email | Yes |
| Text | String | Plain text content of the email | No |
| HTML | String | HTML content of the email (supports rich formatting) | No |
| Cc Email | String | Carbon copy recipients. Multiple addresses can be separated by comma | No |
| Bcc Email | String | Blind carbon copy recipients. Multiple addresses can be separated by comma | No |
| Attachments | String | Names of binary properties containing attachment data. Multiple properties can be comma-separated | No |

## UseCases

- **Transactional Emails** : Send automated order confirmations, password resets, and account notifications
- **Marketing Communications** : Deliver newsletters, promotional campaigns, and event invitations
- **System Notifications** : Send error alerts, monitoring reports, and security notifications
- **Customer Communication** : Send welcome emails, support responses, and appointment reminders
- **E-commerce Integration** : Integrate with shopping carts for order notifications and customer updates
- **CRM Integration** : Send follow-up emails based on customer interactions and sales activities

## Email Format Examples

- **From Email**: 
  - Simple: `admin@example.com`
  - With name: `Admin <admin@example.com>`
- **Multiple Recipients**: `user1@example.com, user2@example.com, user3@example.com`

## Attachment Handling

The node supports file attachments through binary data properties:
1. Binary data must be available from previous nodes
2. Specify the binary property names in the "Attachments" field
3. Multiple attachments can be included by comma-separating property names
4. Original filenames are preserved when available

## API Integration

The node uses Mailgun's REST API v3 for email delivery:
- **Endpoint**: `https://{apiDomain}/v3/{emailDomain}/messages`
- **Method**: POST
- **Authentication**: API key-based authentication
- **Format**: Multipart form data for attachments, JSON for metadata

## Error Handling

The node implements comprehensive error handling:
- **API Errors**: Detailed error messages from Mailgun API
- **Authentication Errors**: Clear indication of credential issues
- **Validation Errors**: Input validation for required fields
- **Continue on Fail**: Option to continue workflow execution even if email sending fails

## Response Data

Successful email sending returns:
- **Message ID**: Unique identifier for the sent message
- **Status**: Delivery status information
- **Queued**: Confirmation that the message was queued for delivery

## Best Practices

- **Domain Verification**: Ensure your sending domain is properly verified in Mailgun
- **Template Usage**: Use HTML templates for consistent branding
- **List Management**: Maintain clean recipient lists to improve deliverability
- **Testing**: Test email content and formatting before production deployment
- **Monitoring**: Monitor delivery rates and bounce statistics in Mailgun dashboard
- **Compliance**: Follow email marketing regulations (CAN-SPAM, GDPR, etc.)

## Related Nodes

- **HTTP Request**: For custom Mailgun API calls
- **JSON**: For processing Mailgun API responses
- **Code**: For custom email content generation
- **Schedule Trigger**: For automated email campaigns

