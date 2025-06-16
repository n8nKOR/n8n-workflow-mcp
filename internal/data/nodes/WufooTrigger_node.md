# Wufoo Trigger Node

## Overview

The Wufoo Trigger node starts workflows when form submissions are received in your Wufoo account. This webhook-based trigger node automatically responds to form entries, enabling real-time processing of form data for lead generation, customer feedback collection, event registration, and other data capture workflows. The node provides flexible data formatting options and sophisticated field type processing for comprehensive form automation.

## Credentials

- Name: wufooApi, Required: Yes

### Credential Configuration
To configure Wufoo API credentials:
1. Log into your Wufoo account
2. Navigate to Account Settings > API Information
3. Copy your API Key and subdomain
4. The subdomain is the part before ".wufoo.com" in your Wufoo URL
5. Use the API Key and subdomain in n8n credentials

## Inputs

None - This is a trigger node that starts workflows based on Wufoo form submissions.

## Outputs

- Main: Outputs form submission data when triggered by Wufoo webhooks

## Properties

### Resource: Form Submissions

#### Operation: Webhook Trigger

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Forms Name or ID | options | Wufoo form to monitor for submissions | Yes | - |
| Only Answers | boolean | Return only form answers without metadata | No | true |

### Configuration Options

#### Form Selection
Choose from available forms in your Wufoo account:
- **Dynamic Loading**: Forms are loaded from your Wufoo API
- **Form Identification**: Uses Wufoo form hash for identification
- **Multi-form Support**: Configure different triggers for different forms

#### Data Processing Options

##### Only Answers
Controls the format of returned data:
- **Enabled (default)**: Returns only the form field answers
- **Disabled**: Returns complete webhook payload including metadata
- **Use Case**: Enable for simplified processing, disable for full context

### Webhook Management

#### Automatic Webhook Lifecycle
The node handles complete webhook management:

1. **Creation**: Automatically creates webhooks when workflow is activated
2. **Security**: Implements handshake key validation for webhook security
3. **Configuration**: Sets up proper webhook endpoints and field processing
4. **Deletion**: Removes webhooks when workflow is deactivated

#### Webhook Security
- **Handshake Key**: Uses Wufoo's handshake key for webhook validation
- **HTTPS Endpoints**: All webhooks use secure HTTPS connections
- **API Validation**: Wufoo validates webhook delivery authenticity
- **Unique URLs**: Each workflow gets a unique webhook URL

### Field Type Processing

#### Basic Fields
- **Text Fields**: Single and multi-line text inputs
- **Number Fields**: Numeric input validation
- **Email Fields**: Email address validation
- **Phone Fields**: Phone number formatting
- **Website Fields**: URL validation

#### Advanced Fields

##### File Upload Fields
Returns file information:
```json
{
  "FieldId": "https://wufoo.com/cabinet/uploaded-file.pdf"
}
```

##### Address Fields
Combines address components into structured object:
```json
{
  "FieldId": {
    "street": "123 Main St",
    "city": "San Francisco",
    "state": "CA",
    "zip": "94102",
    "country": "US"
  }
}
```

##### Checkbox Fields
Returns array of selected values:
```json
{
  "FieldId": ["Option 1", "Option 3", "Option 5"]
}
```

##### Likert Scale Fields
Returns scale responses as structured object:
```json
{
  "FieldId": {
    "Statement 1": "Agree",
    "Statement 2": "Strongly Agree",
    "Statement 3": "Neutral"
  }
}
```

##### Name Fields (Shortname)
Returns name components:
```json
{
  "FieldId": {
    "first": "John",
    "last": "Doe"
  }
}
```

### Response Data Structure

#### Only Answers Format (Default)
Simplified format with just form responses:
```json
{
  "Field1": "John Doe",
  "Field2": "john@example.com",
  "Field3": "Marketing",
  "Field4": {
    "street": "123 Main St",
    "city": "San Francisco",
    "state": "CA"
  }
}
```

#### Full Webhook Format
Complete webhook payload with metadata:
```json
{
  "FormId": "form123",
  "FormStructure": {
    "Fields": [
      {
        "ID": "Field1",
        "Title": "Full Name",
        "Type": "text",
        "Required": true
      }
    ]
  },
  "FieldData": {
    "Field1": "John Doe",
    "Field2": "john@example.com"
  },
  "DateCreated": "2023-01-01 12:00:00",
  "CreatedBy": "192.168.1.1"
}
```

### Supported Field Types

#### Input Fields
- **Text**: Single-line text input
- **Textarea**: Multi-line text input
- **Number**: Numeric input with validation
- **Email**: Email address with validation
- **Phone**: Phone number formatting
- **Website**: URL validation
- **Password**: Secure password input

#### Selection Fields
- **Radio**: Single selection from options
- **Dropdown**: Dropdown selection
- **Checkbox**: Multiple selection from options
- **Likert**: Scale-based responses

#### Complex Fields
- **Name**: First/last name components
- **Address**: Complete address with components
- **Date**: Date selection
- **Time**: Time selection
- **File Upload**: File attachment handling

#### Special Fields
- **Section Break**: Form organization (no data)
- **Page Break**: Multi-page forms (no data)
- **Captcha**: Spam protection (no data)
- **Hidden**: Pre-populated data

### API Integration

#### Wufoo API v3
- **Base URL**: `https://{subdomain}.wufoo.com/api/v3/`
- **Authentication**: API Key authentication
- **Rate Limiting**: Respects Wufoo API rate limits
- **Error Handling**: Comprehensive error handling and logging

#### Webhook Endpoints
- **Creation**: `forms/{formHash}/webhooks.json`
- **Deletion**: `forms/{formHash}/webhooks/{webhookId}.json`
- **Form Loading**: `forms.json`
- **Security**: Handshake key validation

### Advanced Features

#### Form Management
- **Dynamic Form Loading**: Automatically loads available forms
- **Form Hash Identification**: Uses Wufoo's form hash system
- **Multi-form Monitoring**: Support for multiple form triggers
- **Form Validation**: Validates form existence and accessibility

#### Data Processing
- **Field Type Detection**: Automatically detects and processes field types
- **Complex Field Handling**: Sophisticated processing for complex field types
- **Data Normalization**: Standardizes data formats across field types
- **Metadata Extraction**: Optional metadata extraction for advanced processing

#### Error Handling
- **Webhook Validation**: Validates incoming webhook authenticity
- **API Error Handling**: Proper handling of Wufoo API errors
- **Data Validation**: Validates form data structure and content
- **Graceful Degradation**: Continues processing despite individual field errors

### Integration Patterns

#### Lead Management
1. **Lead Capture**: Automatically capture leads from contact forms
2. **CRM Integration**: Add leads directly to CRM systems
3. **Lead Scoring**: Score leads based on form responses
4. **Follow-up Automation**: Trigger follow-up sequences

#### Customer Support
1. **Ticket Creation**: Create support tickets from inquiry forms
2. **Case Routing**: Route cases based on form responses
3. **Auto-responses**: Send automated acknowledgment emails
4. **Escalation**: Escalate urgent inquiries automatically

#### Event Management
1. **Registration Processing**: Process event registrations
2. **Confirmation Emails**: Send registration confirmations
3. **Payment Integration**: Trigger payment processing workflows
4. **Attendee Management**: Manage attendee lists and communications

### Best Practices

#### Form Design
- **Clear Field Names**: Use descriptive field titles for easy processing
- **Required Fields**: Mark essential fields as required
- **Field Types**: Use appropriate field types for validation
- **Logical Structure**: Organize fields in logical order

#### Data Processing
- **Field Mapping**: Map form fields to appropriate system fields
- **Data Validation**: Validate incoming data before processing
- **Error Handling**: Implement proper error handling for failed processing
- **Logging**: Log form submissions for audit trails

#### Performance Optimization
- **Selective Processing**: Only process necessary form data
- **Batch Operations**: Group related operations for efficiency
- **Caching**: Cache form structure for repeated processing
- **Monitoring**: Monitor webhook delivery success rates

### Security Considerations

#### Data Protection
- **PII Handling**: Properly handle personally identifiable information
- **Data Encryption**: Ensure data is encrypted in transit
- **Access Control**: Limit access to form data
- **Retention Policies**: Implement appropriate data retention policies

#### Webhook Security
- **Handshake Validation**: Always validate webhook handshake keys
- **HTTPS Only**: All webhooks use HTTPS encryption
- **Rate Limiting**: Implement rate limiting for webhook endpoints
- **Monitoring**: Monitor for suspicious webhook activity

### Troubleshooting

#### Common Issues
- **Missing Webhooks**: Check workflow activation status and form configuration
- **Authentication Errors**: Verify API credentials and subdomain
- **Data Format Issues**: Check "Only Answers" setting for expected format
- **Webhook Delivery Failures**: Check endpoint accessibility and handshake key

#### Debugging Tools
- **Wufoo Manager**: Check webhook activity in Wufoo manager
- **n8n Execution Log**: Review execution logs for errors
- **API Testing**: Test API connectivity and form access
- **Webhook Testing**: Send test submissions to validate processing

### Migration and Updates

#### Version Compatibility
- **Wufoo API v3**: Compatible with current Wufoo API
- **Webhook Format**: Supports current Wufoo webhook format
- **Field Types**: Supports all current Wufoo field types

#### Configuration Updates
- **Form Changes**: Automatically adapts to form field changes
- **Webhook Recreation**: May require webhook recreation for major changes
- **Testing**: Test webhook functionality after form updates

## UseCases

- **Lead Generation**: Automatically capture and process leads from contact and inquiry forms
- **Customer Feedback Collection**: Process customer feedback surveys and trigger response workflows
- **Event Registration**: Handle event registrations and trigger confirmation and payment workflows
- **Support Ticket Creation**: Create support tickets from customer inquiry and help request forms
- **Newsletter Signup**: Add subscribers to email lists when they complete signup forms
- **Job Application Processing**: Handle job applications and trigger hiring process workflows
- **Contact Form Processing**: Process website contact forms and route inquiries to appropriate teams
- **Order Form Processing**: Handle product orders and trigger fulfillment and payment workflows
- **Survey Data Collection**: Collect survey responses and trigger analysis and reporting workflows
- **Appointment Booking**: Process appointment requests and trigger calendar integration
- **Member Registration**: Handle membership applications and trigger approval workflows
- **Contest Entry Processing**: Process contest entries and trigger winner selection workflows
- **Feedback Form Automation**: Process feedback forms and trigger improvement initiative workflows
- **Quote Request Processing**: Handle quote requests and trigger sales team notifications
- **Product Registration**: Process product registrations and trigger warranty and support workflows
- **Volunteer Signup**: Handle volunteer registrations and trigger coordination workflows
- **Training Registration**: Process training course registrations and trigger enrollment workflows
- **Donation Processing**: Handle donation forms and trigger acknowledgment and receipt workflows
- **Partnership Applications**: Process partnership applications and trigger review workflows
- **Customer Onboarding**: Handle customer onboarding forms and trigger welcome sequences

### Usage Notes
- The Wufoo Trigger node only receives form submissions - it cannot create or modify forms
- Webhooks are automatically managed when workflows are activated/deactivated
- Only Answers option simplifies response data for easier processing
- Field type processing automatically handles complex fields like addresses and checkboxes
- Handshake key validation ensures webhook security and authenticity
- All webhook endpoints use HTTPS for secure data transmission
- The node supports all Wufoo field types including file uploads and complex structured fields
- Form selection is dynamically loaded from your Wufoo account for easy configuration