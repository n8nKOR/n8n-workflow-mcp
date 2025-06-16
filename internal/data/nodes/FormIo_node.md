# Form.io Trigger Node

## Overview

The Form.io Trigger node starts workflows when form submissions are received from Form.io, a powerful form and data management platform. This webhook-based trigger node automatically monitors specified Form.io forms and initiates workflows when submissions are created or updated, enabling real-time form processing and automation. The node supports both cloud-hosted and self-hosted Form.io instances with comprehensive event filtering and data processing options.

## Credentials

- Name: formIoApi, Required: Yes

### Credential Configuration
To configure Form.io API credentials:
1. Create a Form.io account (cloud or self-hosted)
2. Choose your environment type:
   - **Cloud-Hosted**: Use Form.io's cloud service
   - **Self-Hosted**: Use your own Form.io server instance
3. Provide your login credentials:
   - Email address for your Form.io account
   - Password for your Form.io account
4. For self-hosted instances, provide your domain URL

### Authentication Method
The node uses JWT token-based authentication:
- Authenticates with Form.io using email/password
- Receives JWT token for subsequent API calls
- Automatically handles token refresh and management
- Supports both `https://formio.form.io` (cloud) and custom domains

## Inputs

None - This is a trigger node that starts workflows based on Form.io form submissions.

## Outputs

- Main: Outputs form submission data when triggered by Form.io webhooks

## Properties

### Resource: Form Submissions

#### Operation: Webhook Trigger

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name or ID | options | Form.io project to monitor | Yes | - |
| Form Name or ID | options | Specific form within the project | Yes | - |
| Trigger Events | multiOptions | Events to monitor (create, update) | Yes | create |
| Simplify | boolean | Return simplified response format | No | true |

### Configuration Options

#### Project Selection
Choose from available projects in your Form.io account:
- **Dynamic Loading**: Projects are loaded from your Form.io instance
- **Project Support**: Both cloud and self-hosted projects supported
- **Access Control**: Only projects you have access to are shown

#### Form Selection
Select specific forms within the chosen project:
- **Dependent Loading**: Forms load after project selection
- **Form Filtering**: Only active forms are displayed
- **Real-time Updates**: Form list updates when project changes

#### Trigger Events
Configure which form events should trigger the workflow:

##### Submission Created
Triggered when a new form submission is created:
- **Use Case**: Initial form submission processing
- **Data**: Complete submission data
- **Timing**: Immediate upon form submission

##### Submission Updated
Triggered when an existing form submission is modified:
- **Use Case**: Handling form submission edits
- **Data**: Updated submission data
- **Timing**: Upon any submission modification

#### Data Processing Options

##### Simplify Output
Controls the format of returned submission data:

**When Enabled (default):**
- Returns clean, easy-to-process data format
- Removes Form.io metadata and system fields
- Ideal for most workflow automation scenarios

**When Disabled:**
- Returns complete raw Form.io webhook payload
- Includes all metadata, system fields, and technical details
- Useful for advanced integrations requiring full data access

## Webhook Management

### Automatic Webhook Lifecycle
The node handles complete webhook management:

1. **Registration**: Automatically creates webhook actions in Form.io when workflow is activated
2. **Configuration**: Sets up proper event filtering and endpoint targeting
3. **Validation**: Verifies webhook registration and connectivity
4. **Cleanup**: Removes webhook actions when workflow is deactivated

### Webhook Security
- **HTTPS Endpoints**: All webhooks use secure HTTPS connections
- **Authentication**: Form.io validates webhook authenticity
- **Unique URLs**: Each workflow gets a unique webhook URL
- **Reliable Delivery**: Built-in retry logic for failed webhook deliveries

## Data Formats

### Simplified Format (Default)

When `Simplify` is enabled, submission data is returned in clean format:

```json
{
  "submissionId": "60f1e2b3c4d5e6f7a8b9c0d1",
  "formId": "60a1b2c3d4e5f6a7b8c9d0e1",
  "data": {
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@example.com",
    "message": "Hello, this is my submission"
  },
  "metadata": {
    "submittedAt": "2023-01-01T12:00:00Z",
    "ipAddress": "192.168.1.1"
  }
}
```

### Raw Format

When `Simplify` is disabled, complete Form.io webhook payload is returned:

```json
{
  "request": {
    "data": {
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@example.com",
      "submit": true,
      "_id": "60f1e2b3c4d5e6f7a8b9c0d1",
      "form": "60a1b2c3d4e5f6a7b8c9d0e1",
      "owner": "60z1y2x3w4v5u6t7s8r9q0p1",
      "created": "2023-01-01T12:00:00.000Z",
      "modified": "2023-01-01T12:00:00.000Z"
    }
  },
  "response": {},
  "action": "create",
  "form": {
    "_id": "60a1b2c3d4e5f6a7b8c9d0e1",
    "title": "Contact Form",
    "name": "contactForm"
  }
}
```

## Supported Form Types

### Basic Forms
- **Contact Forms**: Customer inquiry and contact forms
- **Registration Forms**: Event, account, and service registration
- **Survey Forms**: Customer feedback and satisfaction surveys
- **Order Forms**: Product and service ordering forms

### Advanced Forms
- **Multi-Page Forms**: Complex forms with multiple steps
- **Conditional Forms**: Forms with dynamic field visibility
- **File Upload Forms**: Forms supporting document and media uploads
- **Data Grid Forms**: Forms with repeatable data sections

### Form Components
- **Text Fields**: Single-line and multi-line text inputs
- **Select Components**: Dropdowns, radio buttons, checkboxes
- **Date/Time**: Date pickers and time selectors
- **File Upload**: Document and image upload components
- **Signatures**: Digital signature capture
- **Data Grid**: Repeatable data entry sections

## API Integration

### Form.io API v1
- **Cloud URL**: `https://api.form.io`
- **Authentication**: JWT token-based authentication
- **Rate Limiting**: Respects Form.io API rate limits
- **Error Handling**: Comprehensive error handling and logging

### Webhook Endpoints
- **Registration**: `POST /project/{projectId}/form/{formId}/action`
- **Management**: CRUD operations on webhook actions
- **Events**: Supports create and update events
- **Delivery**: Reliable webhook delivery with retry logic

## Advanced Features

### Multi-Environment Support
- **Cloud Hosting**: Full support for Form.io cloud service
- **Self-Hosted**: Support for custom Form.io server instances
- **Domain Configuration**: Flexible domain and URL configuration
- **Environment Switching**: Easy switching between environments

### Project Management
- **Multi-Project**: Support for multiple Form.io projects
- **Dynamic Loading**: Real-time project and form discovery
- **Access Control**: Respects Form.io permission system
- **Organization Support**: Works with Form.io organizations

### Event Processing
- **Real-Time**: Immediate processing of form submissions
- **Event Filtering**: Selective event monitoring
- **Batch Support**: Handles multiple concurrent submissions
- **Error Recovery**: Robust error handling and recovery

## Error Handling

### Authentication Errors
- **Invalid Credentials**: Clear error messages for login failures
- **Token Expiration**: Automatic token refresh handling
- **Permission Issues**: Proper handling of access restrictions
- **Environment Errors**: Validation of domain and configuration

### Webhook Errors
- **Registration Failures**: Detailed error reporting for webhook setup
- **Delivery Issues**: Retry logic for failed webhook deliveries
- **Data Validation**: Validation of incoming webhook data
- **Network Errors**: Handling of connectivity issues

### Form Configuration Errors
- **Project Access**: Validation of project permissions
- **Form Availability**: Checking form existence and status
- **Event Configuration**: Validation of event type selections
- **API Limits**: Handling of Form.io API rate limits

## Integration Patterns

### Customer Management
1. **Lead Capture**: Automatically capture leads from contact forms
2. **Customer Onboarding**: Process registration and onboarding forms
3. **Support Tickets**: Create support tickets from inquiry forms
4. **Feedback Processing**: Handle customer feedback and surveys

### E-commerce Integration
1. **Order Processing**: Process product and service orders
2. **Customer Registration**: Handle customer account creation
3. **Payment Integration**: Connect with payment processing workflows
4. **Inventory Management**: Update inventory based on orders

### Event Management
1. **Registration Processing**: Handle event and workshop registrations
2. **Attendee Management**: Manage attendee information and communications
3. **Payment Processing**: Process registration fees and payments
4. **Communication Automation**: Send confirmations and updates

## Best Practices

### Form Design
- **Field Validation**: Use Form.io validation for data quality
- **User Experience**: Design forms for optimal user experience
- **Mobile Optimization**: Ensure forms work well on mobile devices
- **Accessibility**: Follow accessibility best practices

### Webhook Configuration
- **Event Selection**: Choose appropriate events for your workflow needs
- **Data Processing**: Use simplified format unless raw data is required
- **Error Handling**: Implement proper error handling in workflows
- **Testing**: Test webhooks thoroughly before production deployment

### Security Considerations
- **Data Privacy**: Handle form data according to privacy regulations
- **Access Control**: Implement appropriate access controls
- **Encryption**: Ensure data is encrypted in transit and at rest
- **Audit Logging**: Maintain logs of form processing activities

## Performance Optimization

### Webhook Processing
- **Efficient Processing**: Minimize webhook processing time
- **Async Operations**: Use asynchronous processing for heavy operations
- **Queue Management**: Implement queues for high-volume submissions
- **Resource Monitoring**: Monitor CPU and memory usage

### API Usage
- **Connection Pooling**: Implement efficient API connection management
- **Caching**: Cache form definitions and project information
- **Rate Limiting**: Respect Form.io API rate limits
- **Bulk Operations**: Use bulk operations where possible

## Troubleshooting

### Common Issues
- **Missing Submissions**: Check webhook configuration and form settings
- **Authentication Errors**: Verify Form.io credentials and permissions
- **Data Format Issues**: Check simplify setting for expected data format
- **Webhook Delivery Failures**: Check endpoint accessibility and response handling

### Debugging Tools
- **Form.io Admin**: Check webhook actions in Form.io admin interface
- **n8n Execution Log**: Review execution logs for errors and data
- **API Testing**: Test Form.io API connectivity and authentication
- **Webhook Testing**: Send test submissions to validate processing

## Migration and Updates

### Version Compatibility
- **Form.io API**: Compatible with Form.io API v1
- **Node Versions**: Supports n8n node versioning
- **Backward Compatibility**: Maintains compatibility with existing workflows

### Update Procedures
- **Webhook Recreation**: May require webhook recreation for major updates
- **Testing**: Test webhook functionality after Form.io updates
- **Credential Updates**: Update credentials if authentication changes

## UseCases

- **Lead Generation**: Automatically capture and process leads from contact forms and landing pages
- **Customer Support**: Create support tickets automatically when inquiry forms are submitted
- **Event Registration**: Process event registrations and trigger confirmation workflows
- **E-commerce Orders**: Handle product orders and initiate fulfillment processes
- **Survey Processing**: Collect and process customer feedback and survey responses
- **User Registration**: Handle user account creation and onboarding workflows
- **Appointment Booking**: Process appointment requests and trigger scheduling workflows
- **Newsletter Signup**: Add subscribers to email lists when signup forms are submitted
- **Quote Requests**: Process quote requests and trigger sales team notifications
- **Feedback Collection**: Collect customer feedback and trigger analysis workflows
- **Job Applications**: Process job applications and trigger HR workflows
- **Payment Forms**: Handle payment information and trigger billing processes
- **Membership Applications**: Process membership applications and trigger approval workflows
- **Contest Entries**: Handle contest entries and trigger winner selection processes
- **Document Submissions**: Process document uploads and trigger review workflows
- **Reservation Systems**: Handle reservations and booking requests
- **Volunteer Registration**: Process volunteer signups and trigger coordination workflows
- **Training Registration**: Handle training course registrations and trigger enrollment
- **Compliance Forms**: Process compliance forms and trigger audit workflows
- **Data Collection**: Collect research data and trigger analysis processes

### Usage Notes
- The Form.io Trigger node only receives form submissions - it cannot create or modify forms
- Webhooks are automatically managed when workflows are activated/deactivated
- Simplify option provides clean data format for easier processing
- Both cloud-hosted and self-hosted Form.io instances are supported
- Event filtering allows monitoring specific submission activities
- All webhook endpoints use HTTPS for secure data transmission
- The node supports Form.io's complete range of form components and features