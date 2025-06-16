# Typeform Trigger Node

## Overview

The Typeform Trigger node starts workflows when form submissions are received in your Typeform account. This webhook-based trigger node automatically responds to form completions, enabling real-time processing of form responses for lead generation, survey analysis, customer onboarding, and other data collection workflows. The node provides flexible data formatting options to simplify response processing.

## Credentials

- Name: typeformApi, Required: Yes

Authentication options:
- **Access Token**: Personal Access Token from Typeform
- **OAuth2**: OAuth2 authentication for enhanced security

### Credential Configuration

#### Access Token Method
1. Log into your Typeform account
2. Go to Account Settings > Personal tokens
3. Create a new personal access token
4. Copy the token for use in n8n credentials

#### OAuth2 Method
1. Create a Typeform app in your account
2. Configure OAuth2 redirect URIs
3. Use Client ID and Client Secret in n8n OAuth2 credentials

## Inputs

None - This is a trigger node that starts workflows based on Typeform submissions.

## Outputs

- Main: Outputs form submission data when triggered by Typeform webhooks

## Properties

### Resource: Form Submissions

#### Operation: Webhook Trigger

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Form | options | Typeform to monitor for submissions | Yes | - |
| Simplify Answers | boolean | Convert responses to simplified key:value format | No | true |
| Only Answers | boolean | Return only form answers without metadata | No | false |

### Configuration Options

#### Form Selection
Choose from available forms in your Typeform account:
- **Dynamic Loading**: Forms are loaded from your Typeform account
- **Expression Support**: Use expressions to specify form ID dynamically
- **Multi-form Support**: Configure different triggers for different forms

#### Data Processing Options

##### Simplify Answers
When enabled (default), converts complex Typeform response structure to simplified format:
- **Original Format**: Complex nested objects with field IDs
- **Simplified Format**: Key-value pairs using field titles
- **Example**: `"What is your name?": "John Doe"` instead of field IDs

##### Only Answers
When enabled, returns only the form answers without metadata:
- **Excluded Data**: Submission timestamps, form metadata, webhook info
- **Included Data**: Only the actual form responses
- **Use Case**: When you only need the answers for processing

### Webhook Management

#### Automatic Webhook Lifecycle
The node handles complete webhook management:

1. **Creation**: Automatically creates webhooks when workflow is activated
2. **Validation**: Checks for existing webhooks to prevent duplicates  
3. **Configuration**: Sets up proper webhook endpoints and settings
4. **Deletion**: Removes webhooks when workflow is deactivated

#### Webhook Security
- **HTTPS Endpoints**: All webhooks use secure HTTPS connections
- **Automatic Validation**: Typeform validates webhook delivery
- **Unique URLs**: Each workflow gets a unique webhook URL
- **Retry Logic**: Built-in retry handling for failed deliveries

### Response Data Structure

#### Raw Response Format
Without simplification, responses include full Typeform structure:

```json
{
  "event_id": "01FWAZ7H7SWF3VEFCD4E9ZBFA6",
  "event_type": "form_response",
  "form_response": {
    "form_id": "AbCdEf123",
    "token": "response_token_123",
    "submitted_at": "2023-01-01T12:00:00Z",
    "answers": [
      {
        "field": {
          "id": "field_id_123",
          "type": "short_text",
          "ref": "question_1"
        },
        "type": "text",
        "text": "John Doe"
      }
    ]
  }
}
```

#### Simplified Response Format
With simplification enabled:

```json
{
  "What is your name?": "John Doe",
  "Your email address": "john@example.com",
  "How did you hear about us?": "Google Search",
  "form_response_token": "response_token_123",
  "form_response_submitted_at": "2023-01-01T12:00:00Z"
}
```

#### Only Answers Format
With both simplification and "only answers" enabled:

```json
{
  "What is your name?": "John Doe",
  "Your email address": "john@example.com",
  "How did you hear about us?": "Google Search"
}
```

### Supported Field Types

#### Text Fields
- **Short Text**: Single-line text responses
- **Long Text**: Multi-line text responses
- **Email**: Email address validation
- **Website**: URL validation
- **Number**: Numeric responses

#### Choice Fields
- **Multiple Choice**: Single selection from options
- **Picture Choice**: Image-based selection
- **Dropdown**: Dropdown selection
- **Yes/No**: Boolean responses

#### Scale Fields
- **Rating**: Star or numeric ratings
- **Opinion Scale**: Linear scale responses
- **Ranking**: Ordered preferences

#### Date and Time
- **Date**: Date selection
- **Time**: Time selection

#### File Upload
- **File Upload**: Uploaded file references

#### Advanced Fields
- **Payment**: Payment integration responses
- **Legal**: Legal agreement confirmations
- **Hidden Fields**: Pre-populated data

### Field Processing

#### Question Title Mapping
The node automatically maps field IDs to human-readable question titles:
- **Form Definition**: Retrieves form structure from Typeform API
- **Field Mapping**: Maps each field ID to its question title
- **Dynamic Processing**: Handles forms with changing questions

#### Answer Processing
Different field types are processed appropriately:
- **Text Answers**: Direct string values
- **Choice Answers**: Selected option labels
- **Scale Answers**: Numeric values
- **Date Answers**: ISO date strings
- **File Answers**: File URLs and metadata

### API Integration

#### Typeform API v1
- **Base URL**: `https://api.typeform.com/`
- **Authentication**: Bearer token or OAuth2
- **Rate Limiting**: Respects Typeform API rate limits
- **Error Handling**: Comprehensive error handling and logging

#### Webhook API
- **Endpoint**: `/webhooks/{form_id}`
- **Events**: `form_response` events
- **Security**: HTTPS-only endpoints
- **Validation**: Automatic webhook validation

### Advanced Features

#### Version Compatibility
- **v1.0**: Original implementation
- **v1.1**: Enhanced data processing and field mapping
- **Backward Compatibility**: Maintains compatibility with existing workflows

#### Expression Support
- **Dynamic Form Selection**: Use expressions for form ID
- **Conditional Processing**: Process different forms differently
- **Variable Integration**: Integrate with workflow variables

#### Error Handling
- **Webhook Failures**: Graceful handling of webhook delivery failures
- **API Errors**: Proper error mapping and reporting
- **Data Validation**: Validation of incoming webhook data
- **Continue on Fail**: Option to continue workflow despite errors

### Integration Patterns

#### CRM Integration
1. **Lead Capture**: Automatically add form submissions to CRM
2. **Contact Updates**: Update existing contacts with form data
3. **Lead Scoring**: Score leads based on form responses

#### Marketing Automation
1. **Email Sequences**: Trigger email sequences based on form responses
2. **Segmentation**: Segment contacts based on form answers
3. **Personalization**: Use form data for personalized content

#### Customer Support
1. **Ticket Creation**: Create support tickets from contact forms
2. **Case Routing**: Route cases based on form responses
3. **Follow-up Automation**: Automated follow-up based on inquiry type

### Best Practices

#### Form Design
- **Clear Questions**: Use clear, descriptive question titles
- **Logical Structure**: Organize questions in logical order
- **Required Fields**: Mark essential fields as required
- **Validation**: Use appropriate field types for validation

#### Data Processing
- **Simplification**: Use simplified answers for easier processing
- **Validation**: Validate incoming data before processing
- **Error Handling**: Implement proper error handling
- **Logging**: Log form submissions for audit trails

#### Performance Optimization
- **Selective Processing**: Only process necessary data
- **Batch Operations**: Group related operations together
- **Caching**: Cache form definitions when possible
- **Monitoring**: Monitor webhook delivery success

### Security Considerations

#### Data Protection
- **PII Handling**: Properly handle personally identifiable information
- **Data Encryption**: Ensure data is encrypted in transit
- **Access Control**: Limit access to form data
- **Retention Policies**: Implement data retention policies

#### Webhook Security
- **HTTPS Only**: All webhooks use HTTPS encryption
- **Validation**: Validate webhook authenticity
- **Rate Limiting**: Implement rate limiting for webhooks
- **Monitoring**: Monitor for suspicious webhook activity

### Troubleshooting

#### Common Issues
- **Missing Webhooks**: Check workflow activation status
- **Authentication Errors**: Verify API credentials
- **Data Format Issues**: Check simplification settings
- **Webhook Delivery Failures**: Check endpoint accessibility

#### Debugging Tools
- **Typeform Activity**: Check webhook activity in Typeform dashboard
- **n8n Execution Log**: Review execution logs for errors
- **API Testing**: Test API connectivity and permissions
- **Webhook Testing**: Send test webhooks from Typeform

### Migration and Updates

#### Version Migration
- **v1.0 to v1.1**: Automatic migration of existing configurations
- **Testing**: Test workflows after version updates
- **Backup**: Backup workflow configurations before updates

#### API Changes
- **API Monitoring**: Monitor Typeform API changes
- **Compatibility**: Maintain backward compatibility
- **Updates**: Update node when API changes occur

## UseCases

- **Lead Generation**: Automatically capture and process leads from contact forms and landing pages
- **Customer Onboarding**: Trigger onboarding sequences when new customers complete registration forms
- **Survey Data Processing**: Process survey responses for market research and customer feedback analysis
- **Event Registration**: Handle event registrations and trigger confirmation emails and calendar invites
- **Support Ticket Creation**: Automatically create support tickets from customer inquiry forms
- **Newsletter Signup**: Add subscribers to email lists when they complete signup forms
- **Order Forms**: Process product orders and trigger fulfillment workflows
- **Job Applications**: Handle job application submissions and trigger hiring process workflows
- **Feedback Collection**: Collect and process customer feedback for product improvement initiatives
- **Quiz and Assessment**: Process quiz responses and trigger personalized result delivery
- **Contact Form Processing**: Handle website contact forms and route inquiries to appropriate teams
- **Registration Workflows**: Process event, course, or service registrations and trigger follow-up actions
- **Data Collection**: Collect market research data and trigger analysis workflows
- **Appointment Booking**: Process appointment requests and trigger calendar integration
- **Contest Entries**: Handle contest or giveaway entries and trigger winner selection processes
- **Membership Applications**: Process membership applications and trigger approval workflows
- **Product Feedback**: Collect product feedback and trigger product team notifications
- **Customer Surveys**: Process customer satisfaction surveys and trigger response workflows
- **Lead Qualification**: Qualify leads based on form responses and trigger sales team notifications
- **Integration Synchronization**: Sync form data with CRM, email marketing, and other business systems

### Usage Notes
- The Typeform Trigger node only receives form submissions - it cannot create or modify forms
- Webhooks are automatically managed when workflows are activated/deactivated
- Simplify Answers option converts complex response structures to easy-to-use key-value pairs
- Only Answers option excludes metadata for workflows that only need the form responses
- Form selection can be done via dropdown or expression for dynamic form handling
- All webhook endpoints use HTTPS for secure data transmission
- The node supports multiple authentication methods including OAuth2 for enhanced security
- Field types are automatically detected and processed appropriately for different response formats