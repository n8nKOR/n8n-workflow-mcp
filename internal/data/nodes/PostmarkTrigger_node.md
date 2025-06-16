# Postmark Trigger Node

## Overview

The Postmark Trigger node starts workflows when specific email events occur in your Postmark account. This webhook-based trigger node automatically responds to email delivery events such as bounces, opens, clicks, and spam complaints, enabling real-time email tracking and automated response workflows. The node integrates directly with Postmark's webhook system for immediate event notification.

## Credentials

- Name: postmarkApi, Required: Yes

### Credential Configuration
To configure Postmark API credentials:
1. Log into your Postmark account
2. Navigate to Servers in your account dashboard
3. Select the server you want to monitor
4. Go to Settings > API Tokens
5. Copy your Server API Token
6. Use this token in the n8n Postmark API credential

## Inputs

None - This is a trigger node that starts workflows based on Postmark email events.

## Outputs

- Main: Outputs email event data when triggered by Postmark webhooks

## Properties

### Resource: Email Events

#### Operation: Webhook Trigger

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Events | multiOptions | Email events to monitor | Yes | delivery |
| First Open Only | boolean | For 'open' events, trigger only on first open | No | false |
| Include Content | boolean | For 'bounce' and 'spam complaint' events, include message content | No | false |

### Event Types

#### Delivery
Triggered when an email is successfully delivered to the recipient's inbox.

**Use Cases:**
- Confirm successful email delivery
- Update delivery status in CRM systems
- Trigger follow-up workflows

#### Bounce
Triggered when an email bounces back due to delivery issues.

**Bounce Types Included:**
- Hard bounces (permanent failures)
- Soft bounces (temporary failures)
- Blocked bounces (spam filtering)

**Use Cases:**
- Remove invalid email addresses from lists
- Update contact status in databases
- Trigger alternative communication methods

#### Open
Triggered when a recipient opens an email.

**Configuration Options:**
- **First Open Only**: Trigger only on the first open per email
- **All Opens**: Trigger on every open (including repeated opens)

**Use Cases:**
- Track email engagement
- Trigger follow-up sequences
- Measure campaign effectiveness

#### Click
Triggered when a recipient clicks a tracked link in an email.

**Use Cases:**
- Track link engagement
- Trigger landing page personalization
- Measure content effectiveness
- Start conversion workflows

#### Spam Complaint
Triggered when a recipient marks an email as spam.

**Configuration Options:**
- **Include Content**: Include original message content in webhook data

**Use Cases:**
- Automatically remove complainers from lists
- Alert marketing teams to deliverability issues
- Update suppression lists

#### Subscription Change
Triggered when a recipient's subscription status changes.

**Use Cases:**
- Update subscription preferences
- Sync with marketing automation platforms
- Manage suppression lists

### Webhook Configuration

#### Automatic Webhook Management
The node automatically handles webhook lifecycle:

1. **Creation**: Creates webhooks when workflow is activated
2. **Validation**: Checks for existing webhooks to prevent duplicates
3. **Configuration**: Sets up event filtering and options
4. **Deletion**: Removes webhooks when workflow is deactivated

#### Webhook Security
- Webhook URLs are automatically generated and secured
- Postmark validates webhook authenticity
- HTTPS-only webhook endpoints
- Automatic retry handling for failed deliveries

### Event Data Structure

#### Common Fields
All events include these base fields:

| Field | Type | Description |
|-------|------|-------------|
| RecordType | string | Type of event (e.g., 'Delivery', 'Bounce') |
| MessageID | string | Unique message identifier |
| Recipient | string | Email address of recipient |
| Tag | string | Message tag (if set) |
| DeliveredAt | string | ISO timestamp of event |
| Details | string | Additional event details |
| Metadata | object | Custom metadata attached to email |

#### Delivery Event Data
```json
{
  "RecordType": "Delivery",
  "ServerID": 12345,
  "MessageID": "883953f4-6105-42a2-a16a-77a8eac79483",
  "Recipient": "john@example.com",
  "Tag": "welcome-email",
  "DeliveredAt": "2023-01-01T12:00:00.0000000Z",
  "Details": "Test delivery webhook details"
}
```

#### Bounce Event Data
```json
{
  "RecordType": "Bounce",
  "ID": 42,
  "Type": "HardBounce",
  "TypeCode": 1,
  "Name": "Hard bounce",
  "MessageID": "883953f4-6105-42a2-a16a-77a8eac79483",
  "Description": "The server was unable to deliver your message",
  "Details": "SMTP error details",
  "Email": "john@example.com",
  "From": "sender@example.com",
  "BouncedAt": "2023-01-01T12:00:00.0000000Z"
}
```

#### Open Event Data
```json
{
  "RecordType": "Open",
  "FirstOpen": true,
  "MessageID": "883953f4-6105-42a2-a16a-77a8eac79483",
  "Recipient": "john@example.com",
  "ReceivedAt": "2023-01-01T12:00:00.0000000Z",
  "Tag": "newsletter",
  "UserAgent": "Mozilla/5.0...",
  "Platform": "Desktop",
  "Client": "Gmail"
}
```

#### Click Event Data
```json
{
  "RecordType": "Click",
  "MessageID": "883953f4-6105-42a2-a16a-77a8eac79483",
  "Recipient": "john@example.com",
  "ReceivedAt": "2023-01-01T12:00:00.0000000Z",
  "Tag": "newsletter", 
  "Link": "https://example.com/product",
  "LinkID": "12345",
  "UserAgent": "Mozilla/5.0...",
  "Platform": "Desktop",
  "Client": "Gmail"
}
```

### Advanced Features

#### Event Filtering
- **Multiple Event Selection**: Monitor multiple event types simultaneously
- **Event-Specific Options**: Configure options per event type
- **Tag Filtering**: Filter events by message tags (configured in Postmark)

#### Metadata Integration
- **Custom Metadata**: Access custom metadata attached to emails
- **Tag System**: Use Postmark's tag system for event categorization
- **Batch Processing**: Handle multiple events in sequence

#### Error Handling
- **Webhook Validation**: Automatic validation of webhook authenticity
- **Retry Logic**: Built-in retry for failed webhook deliveries
- **Error Logging**: Detailed logging of webhook failures
- **Graceful Degradation**: Continue processing despite individual event failures

### Integration Patterns

#### CRM Integration
1. **Lead Scoring**: Update lead scores based on email engagement
2. **Contact Status**: Update contact records with delivery status
3. **Engagement Tracking**: Track customer engagement across touchpoints

#### Marketing Automation
1. **Drip Campaigns**: Trigger follow-up emails based on engagement
2. **Segmentation**: Automatically segment contacts by behavior
3. **A/B Testing**: Track performance of different email variants

#### Customer Support
1. **Delivery Issues**: Alert support team to delivery problems
2. **Engagement Alerts**: Notify teams of high-value prospect engagement
3. **Suppression Management**: Automatically manage suppression lists

### Best Practices

#### Webhook Management
- **Test Webhooks**: Use Postmark's webhook testing tools
- **Monitor Performance**: Track webhook delivery success rates
- **Handle Failures**: Implement proper error handling and retry logic
- **Secure Endpoints**: Ensure webhook endpoints are secure

#### Event Processing
- **Deduplicate Events**: Handle potential duplicate webhook deliveries
- **Batch Operations**: Process multiple events efficiently
- **Data Validation**: Validate webhook data before processing
- **Logging**: Maintain audit trails of event processing

#### Performance Optimization
- **Selective Monitoring**: Only monitor necessary event types
- **Efficient Processing**: Minimize processing time for webhook handlers
- **Queue Management**: Use queues for high-volume event processing
- **Resource Management**: Monitor memory and CPU usage

### Security Considerations

#### Webhook Security
- **HTTPS Only**: All webhooks use HTTPS encryption
- **Authentication**: Postmark validates webhook authenticity
- **Rate Limiting**: Implement rate limiting for webhook endpoints
- **IP Filtering**: Consider IP whitelist for additional security

#### Data Protection
- **PII Handling**: Properly handle personally identifiable information
- **Data Retention**: Implement appropriate data retention policies
- **Compliance**: Ensure GDPR and other privacy regulation compliance
- **Encryption**: Encrypt sensitive data in transit and at rest

### Troubleshooting

#### Common Issues
- **Missing Events**: Check webhook configuration and Postmark settings
- **Duplicate Events**: Implement deduplication logic
- **Authentication Errors**: Verify API credentials and permissions
- **Rate Limits**: Monitor and handle Postmark API rate limits

#### Debugging Tools
- **Webhook Logs**: Use Postmark's webhook activity logs
- **Test Webhooks**: Send test webhooks from Postmark dashboard
- **Event Inspector**: Use n8n's execution log to inspect webhook data
- **API Monitoring**: Monitor Postmark API health and status

### Migration and Updates

#### Version Compatibility
- **Postmark API**: Compatible with Postmark API v1
- **Webhook Format**: Supports current Postmark webhook format
- **Backward Compatibility**: Maintains compatibility with existing workflows

#### Update Procedures
- **Webhook Recreation**: May require webhook recreation for major updates
- **Testing**: Test webhook functionality after updates
- **Monitoring**: Monitor webhook delivery after changes

## UseCases

- **Email Deliverability Monitoring**: Track email delivery success rates and identify delivery issues in real-time
- **Automated Bounce Handling**: Automatically remove bounced email addresses from mailing lists and update contact databases
- **Engagement Tracking**: Monitor email opens and clicks to measure campaign effectiveness and customer engagement
- **Lead Scoring**: Update lead scores in CRM systems based on email interaction behavior
- **Follow-up Automation**: Trigger automated follow-up sequences when recipients engage with emails
- **Spam Complaint Management**: Automatically process spam complaints and update suppression lists
- **Customer Support Alerts**: Notify support teams when important emails fail to deliver or bounce
- **Marketing Attribution**: Track which emails lead to conversions and sales for attribution analysis
- **List Hygiene**: Maintain clean email lists by automatically handling bounces and unsubscribes
- **Personalization Triggers**: Trigger personalized content delivery based on email engagement patterns
- **A/B Testing Analysis**: Collect engagement data for A/B testing email campaigns and content variations
- **Compliance Reporting**: Generate reports on email delivery and engagement for compliance requirements
- **Real-time Notifications**: Send instant notifications to teams when high-priority emails are opened or clicked
- **Customer Journey Tracking**: Track customer progression through email sequences and marketing funnels
- **Delivery Optimization**: Identify delivery patterns and optimize sending times and frequencies
- **Integration Synchronization**: Keep multiple systems synchronized with real-time email event data
- **Performance Monitoring**: Monitor email infrastructure performance and identify potential issues
- **Segmentation Automation**: Automatically segment contacts based on email engagement behavior
- **Revenue Attribution**: Connect email engagement events to revenue and conversion tracking systems
- **Quality Assurance**: Monitor email delivery quality and identify potential deliverability issues

### Usage Notes
- The Postmark Trigger node only handles receiving email events via webhooks - it does not send emails
- Webhooks are automatically managed (created/deleted) when workflows are activated/deactivated
- First Open Only option for open events prevents duplicate triggers for the same email
- Include Content option for bounces and spam complaints provides additional context for analysis
- All webhook endpoints use HTTPS for secure data transmission
- Event data includes comprehensive metadata for detailed analysis and processing
- Multiple event types can be monitored simultaneously in a single workflow
- Webhook delivery is handled asynchronously by Postmark with automatic retry logic