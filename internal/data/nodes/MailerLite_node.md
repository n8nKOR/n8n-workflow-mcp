# MailerLite Node

## Overview

The MailerLite node allows you to consume the MailerLite API for email marketing automation. MailerLite is a popular email marketing platform that helps businesses create, send, and track email campaigns. This node enables you to manage subscribers, handle email campaigns, and automate your email marketing workflows through n8n.

## Credentials

This node requires MailerLite API credentials:
- **API Key**: Your MailerLite API key with appropriate permissions

To obtain API credentials:
1. Log into your MailerLite account
2. Navigate to **Settings** → **API**
3. Generate a new API key or copy an existing one
4. Copy the API key for use in n8n

## Inputs

- **Main**: JSON data for creating/updating subscribers and resources

## Outputs

- **Main**: Returns JSON data from MailerLite API responses

## Properties

### Resource: Subscriber

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | String | Email address of the subscriber | Yes |
| Additional Fields | Collection | Optional subscriber properties | No |

**Additional Fields for Creation:**
| Field Name | Type | Description |
|---|---|---|
| Name | String | Full name of the subscriber |
| Fields | Object | Custom fields for the subscriber |
| Groups | Array | Groups to add subscriber to |
| Status | Options | Subscriber status (active, unsubscribed, etc.) |
| Signup IP | String | IP address used for signup |
| Signup Timestamp | DateTime | When the subscriber signed up |
| Confirmation IP | String | IP address used for confirmation |
| Confirmation Timestamp | DateTime | When subscription was confirmed |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Subscriber ID | String | ID of the subscriber to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Additional Options | Collection | Filtering and pagination options | No |

**Get Many Options:**
| Field Name | Type | Description |
|---|---|---|
| Filter | Object | Filters to apply to the query |
| Sort | String | Sort field and direction |
| Cursor | String | Pagination cursor for next page |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Subscriber ID | String | ID of the subscriber to update | Yes |
| Update Fields | Collection | Subscriber properties to update | No |

**Update Fields:**
| Field Name | Type | Description |
|---|---|---|
| Email | String | New email address |
| Name | String | Updated name |
| Fields | Object | Updated custom fields |
| Groups | Array | Updated group memberships |
| Status | Options | Updated subscriber status |

## Features

- **Subscriber Management**: Create, retrieve, update, and manage email subscribers
- **Custom Fields**: Support for custom subscriber fields and data
- **Group Management**: Organize subscribers into groups and segments
- **Status Tracking**: Track subscriber status (active, unsubscribed, bounced, etc.)
- **Bulk Operations**: Handle multiple subscribers efficiently
- **API v2 Support**: Uses the latest MailerLite API version
- **Real-time Sync**: Keep subscriber data synchronized

## UseCases

### Email Marketing Automation
- **Welcome Series**: Automatically add new users to welcome email sequences
- **Lead Nurturing**: Add leads to appropriate email campaigns based on behavior
- **Customer Onboarding**: Create onboarding email sequences for new customers
- **Re-engagement**: Manage re-engagement campaigns for inactive subscribers

### E-commerce Integration
- **Abandoned Cart**: Add customers to abandoned cart email sequences
- **Purchase Follow-up**: Send post-purchase email sequences
- **Product Recommendations**: Segment customers for targeted product emails
- **Customer Retention**: Create retention email campaigns

### Event Management
- **Event Registration**: Add event registrants to specific email lists
- **Event Reminders**: Send automated event reminder sequences
- **Post-event Follow-up**: Manage post-event email communications
- **Waitlist Management**: Handle event waitlists and notifications

### Content Marketing
- **Newsletter Signup**: Add blog subscribers to newsletter lists
- **Content Delivery**: Automate content delivery based on interests
- **Course Enrollment**: Manage email sequences for online courses
- **Content Segmentation**: Segment subscribers based on content preferences

### Integration Workflows
- **New User Registration**: Automatically add new website registrants to MailerLite with custom fields and tags
- **E-commerce Customer Sync**: Update subscriber profiles with purchase data and customer behavior for targeted campaigns

## Subscriber Status Types

### Active Status
- **Active**: Confirmed and engaged subscribers
- **Unconfirmed**: Subscribers pending email confirmation
- **Subscribed**: Recently subscribed, pending confirmation

### Inactive Status
- **Unsubscribed**: Subscribers who opted out
- **Bounced**: Email addresses that bounced
- **Junk**: Marked as spam or invalid
- **Cancelled**: Administratively cancelled subscriptions

## Custom Fields Management

### Standard Fields
- **Name**: Subscriber's full name
- **Email**: Email address (required)
- **Phone**: Phone number
- **Company**: Company name
- **City**: City location
- **State**: State/Province
- **Country**: Country
- **ZIP**: Postal code

### Custom Fields
- **Text Fields**: Custom text input fields
- **Number Fields**: Numeric data fields
- **Date Fields**: Date and time fields
- **Boolean Fields**: True/false fields
- **Dropdown Fields**: Pre-defined option lists

### Field Configuration
Custom fields must be configured in MailerLite before use:
1. Create custom fields in MailerLite dashboard
2. Note the field IDs or names
3. Use field references in n8n workflows

## Groups and Segmentation

### Group Management
- **Add to Groups**: Assign subscribers to specific groups
- **Remove from Groups**: Unassign subscribers from groups
- **Group Hierarchy**: Organize groups by campaign or purpose
- **Dynamic Grouping**: Automatically assign based on criteria

### Segmentation Strategies
- **Behavioral Segmentation**: Based on email interactions
- **Demographic Segmentation**: Based on subscriber data
- **Purchase Segmentation**: Based on buying behavior
- **Engagement Segmentation**: Based on email engagement levels

## API Rate Limits

MailerLite API has rate limiting:
- **Standard Plan**: 1,000 requests per hour
- **Professional Plan**: 3,000 requests per hour  
- **Enterprise Plan**: 10,000 requests per hour
- **Burst Limits**: Short-term higher limits for sudden spikes

## Error Handling

Common issues and solutions:
- **Invalid Email**: Verify email format and domain validity
- **Duplicate Subscriber**: Handle existing subscriber scenarios
- **Rate Limit Exceeded**: Implement retry logic with delays
- **Invalid API Key**: Verify API key and account permissions
- **Field Validation**: Ensure custom fields exist and are valid

## Authentication & Security

### API Key Management
- **Key Generation**: Create API keys in MailerLite dashboard
- **Key Rotation**: Regularly rotate API keys for security
- **Permission Levels**: Use appropriate permission levels
- **Secure Storage**: Store API keys securely in n8n credentials

### Data Privacy
- **GDPR Compliance**: Handle subscriber data according to GDPR
- **Data Retention**: Manage subscriber data retention policies
- **Consent Tracking**: Track subscriber consent and preferences
- **Data Export**: Export subscriber data when requested


## Best Practices

### Data Management
- **Clean Data**: Validate email addresses before adding
- **Avoid Duplicates**: Check for existing subscribers
- **Custom Fields**: Use meaningful custom field names
- **Group Organization**: Organize groups logically

### Performance Optimization
- **Batch Operations**: Process multiple subscribers efficiently
- **Rate Limiting**: Respect API rate limits
- **Error Handling**: Implement robust error handling
- **Monitoring**: Monitor API usage and performance

### Compliance
- **Double Opt-in**: Use double opt-in for better deliverability
- **Unsubscribe Handling**: Properly handle unsubscribe requests
- **Data Protection**: Follow data protection regulations
- **Consent Management**: Track and manage subscriber consent

## Advanced Features

### Automation Integration
- **Webhook Support**: Real-time subscriber updates
- **Campaign Triggers**: Trigger campaigns based on subscriber actions
- **Behavioral Tracking**: Track subscriber behavior and preferences
- **A/B Testing**: Segment subscribers for testing campaigns

### Analytics Integration
- **Engagement Tracking**: Track email open and click rates
- **Conversion Tracking**: Track subscriber conversion metrics
- **Revenue Attribution**: Track revenue from email campaigns
- **Reporting**: Generate detailed subscriber and campaign reports

