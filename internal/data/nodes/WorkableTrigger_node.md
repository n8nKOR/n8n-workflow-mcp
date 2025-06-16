# Workable Trigger Node

## Overview

The Workable Trigger node starts workflows when specific recruiting events occur in your Workable account. This webhook-based trigger node automatically responds to candidate management activities such as new candidate applications and candidate stage movements, enabling real-time hiring process automation, candidate tracking, and recruitment workflow optimization. The node provides filtering options to target specific jobs and pipeline stages.

## Credentials

- Name: workableApi, Required: Yes

### Credential Configuration
To configure Workable API credentials:
1. Log into your Workable account as an admin
2. Navigate to Settings > Integrations > API Access Tokens
3. Create a new access token with appropriate permissions
4. Note your Workable subdomain (e.g., "company" from company.workable.com)
5. Use the subdomain and access token in n8n credentials

## Inputs

None - This is a trigger node that starts workflows based on Workable recruiting events.

## Outputs

- Main: Outputs candidate and event data when triggered by Workable webhooks

## Properties

### Resource: Recruiting Events

#### Operation: Webhook Trigger

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Trigger On | options | Type of event to monitor | Yes | Candidate Created |
| Job | options | Specific job to monitor (optional filter) | No | All Jobs |
| Stage | options | Specific stage to monitor (optional filter) | No | All Stages |

### Event Types

#### Candidate Created
Triggered when a new candidate applies to any job or a specific job.

**Use Cases:**
- Welcome new candidates with automated emails
- Create candidate records in external systems
- Trigger initial screening workflows
- Notify hiring managers of new applications

**Event Data Includes:**
- Candidate profile information
- Application details
- Job information
- Source tracking data

#### Candidate Moved
Triggered when a candidate moves between stages in the hiring pipeline.

**Use Cases:**
- Update candidate status in external systems
- Send stage-specific communications
- Trigger interview scheduling workflows
- Track hiring pipeline metrics

**Event Data Includes:**
- Candidate information
- Previous and current stage details
- Job information
- Movement timestamp

### Filtering Options

#### Job Filter
Filter events to only trigger for specific jobs:
- **All Jobs**: Monitor all jobs in the account (default)
- **Specific Job**: Monitor only candidates for a selected job
- **Dynamic Loading**: Jobs are loaded from your Workable account
- **Job Shortcode**: Uses Workable's job shortcode for identification

#### Stage Filter
Filter events to only trigger for specific pipeline stages:
- **All Stages**: Monitor all stage movements (default)
- **Specific Stage**: Monitor only movements to/from selected stage
- **Dynamic Loading**: Stages are loaded from your Workable account
- **Stage Slug**: Uses Workable's stage slug for identification

### Webhook Management

#### Automatic Webhook Lifecycle
The node handles complete webhook management:

1. **Creation**: Automatically creates webhooks when workflow is activated
2. **Validation**: Checks for existing webhooks to prevent duplicates
3. **Configuration**: Sets up proper event filtering and targeting
4. **Deletion**: Removes webhooks when workflow is deactivated

#### Webhook Security
- **HTTPS Endpoints**: All webhooks use secure HTTPS connections
- **Access Token Authentication**: Workable validates webhook authenticity
- **Unique URLs**: Each workflow gets a unique webhook URL
- **Automatic Retry**: Built-in retry logic for failed webhook deliveries

### Event Data Structure

#### Candidate Created Event
```json
{
  "event": "candidate_created",
  "data": {
    "candidate": {
      "id": "12345678",
      "name": "John Doe",
      "headline": "Software Engineer",
      "account": {
        "subdomain": "company",
        "name": "Company Name"
      },
      "job": {
        "shortcode": "ABC123",
        "title": "Senior Developer"
      },
      "stage": "Applied",
      "disqualified": false,
      "disqualification_reason": null,
      "sourced": false,
      "profile_url": "https://company.workable.com/backend/jobs/ABC123/candidates/12345678",
      "address": "San Francisco, CA",
      "phone": "+1234567890",
      "email": "john@example.com",
      "cover_letter": "Application cover letter text...",
      "resume_url": "https://workable-file-storage.s3.amazonaws.com/..."
    }
  }
}
```

#### Candidate Moved Event
```json
{
  "event": "candidate_moved",
  "data": {
    "candidate": {
      "id": "12345678",
      "name": "John Doe",
      "headline": "Software Engineer",
      "account": {
        "subdomain": "company",
        "name": "Company Name"
      },
      "job": {
        "shortcode": "ABC123", 
        "title": "Senior Developer"
      },
      "stage": "Phone Interview",
      "previous_stage": "Applied",
      "disqualified": false,
      "moved_at": "2023-01-01T12:00:00Z"
    }
  }
}
```

### Candidate Data Fields

#### Core Information
- **ID**: Unique candidate identifier
- **Name**: Candidate's full name
- **Email**: Contact email address
- **Phone**: Phone number
- **Address**: Location information
- **Headline**: Professional headline/title

#### Application Details
- **Job Information**: Job shortcode, title, and details
- **Stage**: Current pipeline stage
- **Previous Stage**: Previous stage (for moved events)
- **Application Date**: When candidate applied
- **Source**: How candidate found the position

#### Profile Data
- **Resume URL**: Link to candidate's resume
- **Cover Letter**: Application cover letter text
- **Profile URL**: Direct link to candidate profile in Workable
- **Portfolio**: Links to candidate portfolio/work samples

#### Status Information
- **Disqualified**: Whether candidate is disqualified
- **Disqualification Reason**: Reason for disqualification
- **Sourced**: Whether candidate was sourced by recruiter
- **Rating**: Candidate rating/score

### API Integration

#### Workable API v3
- **Base URL**: `https://{subdomain}.workable.com/spi/v3`
- **Authentication**: Bearer token with access token
- **Rate Limiting**: Respects Workable API rate limits
- **Error Handling**: Comprehensive error handling and logging

#### Webhook Endpoints
- **Event Subscriptions**: Subscribes to candidate lifecycle events
- **Target URL**: Points to n8n webhook endpoint
- **Event Filtering**: Server-side filtering by job and stage
- **Delivery Guarantees**: Reliable webhook delivery with retry logic

### Advanced Features

#### Multi-Job Monitoring
- **Account-Wide**: Monitor all jobs simultaneously
- **Job-Specific**: Focus on specific positions or departments
- **Dynamic Configuration**: Jobs loaded from Workable account
- **Bulk Processing**: Handle multiple candidate events efficiently

#### Stage-Based Filtering
- **Pipeline Stages**: Monitor specific stages in hiring pipeline
- **Custom Stages**: Support for custom stage configurations
- **Stage Transitions**: Track candidate movement between stages
- **Progress Tracking**: Monitor candidate progression through pipeline

#### Event Correlation
- **Candidate Tracking**: Track individual candidate journey
- **Job Analytics**: Analyze hiring metrics per job
- **Stage Performance**: Monitor stage conversion rates
- **Source Analysis**: Track candidate source effectiveness

### Error Handling

#### Webhook Validation
- **Event Validation**: Validates incoming webhook data structure
- **Authentication Verification**: Verifies webhook authenticity
- **Data Integrity**: Ensures complete and valid event data
- **Error Recovery**: Graceful handling of malformed webhooks

#### API Error Handling
- **Rate Limit Management**: Automatic retry with rate limit respect
- **Authentication Errors**: Clear error messages for credential issues
- **Network Errors**: Retry logic for connectivity problems
- **Service Errors**: Proper handling of Workable service issues

### Integration Patterns

#### ATS Integration
1. **Candidate Sync**: Sync candidates with external ATS systems
2. **Profile Enrichment**: Enhance candidate profiles with external data
3. **Multi-System Tracking**: Track candidates across multiple platforms
4. **Data Consolidation**: Consolidate candidate data from multiple sources

#### CRM Integration
1. **Lead Generation**: Convert candidates to leads in CRM systems
2. **Contact Management**: Maintain candidate contact information
3. **Relationship Tracking**: Track long-term candidate relationships
4. **Pipeline Management**: Sync hiring pipeline with sales pipeline

#### Communication Automation
1. **Welcome Emails**: Send automated welcome messages to new candidates
2. **Status Updates**: Notify candidates of application status changes
3. **Interview Scheduling**: Trigger interview scheduling workflows
4. **Rejection Notifications**: Send personalized rejection communications

### Best Practices

#### Event Processing
- **Selective Monitoring**: Only monitor events relevant to your workflow
- **Efficient Processing**: Process events quickly to avoid webhook timeouts
- **Error Handling**: Implement proper error handling for failed processing
- **Logging**: Maintain audit trails of candidate events

#### Data Management
- **Privacy Compliance**: Handle candidate data according to privacy regulations
- **Data Security**: Protect sensitive candidate information
- **Retention Policies**: Implement appropriate data retention policies
- **Access Control**: Limit access to candidate data

#### Performance Optimization
- **Batch Processing**: Group related operations for efficiency
- **Caching**: Cache frequently accessed job and stage information
- **Resource Management**: Monitor memory and CPU usage
- **Rate Limiting**: Respect API rate limits and webhook processing capacity

### Security Considerations

#### Data Protection
- **PII Handling**: Properly handle personally identifiable information
- **Encryption**: Ensure data is encrypted in transit and at rest
- **Access Logging**: Log access to candidate data for audit purposes
- **Compliance**: Ensure GDPR and other privacy regulation compliance

#### Webhook Security
- **HTTPS Only**: All webhooks use HTTPS encryption
- **Authentication**: Validate webhook authenticity
- **Rate Limiting**: Implement rate limiting for webhook endpoints
- **Monitoring**: Monitor for suspicious webhook activity

### Troubleshooting

#### Common Issues
- **Missing Events**: Check webhook configuration and job/stage filters
- **Authentication Errors**: Verify API credentials and permissions
- **Timeout Issues**: Check webhook processing time and optimize workflows
- **Duplicate Events**: Implement deduplication logic if necessary

#### Debugging Tools
- **Workable Webhooks**: Check webhook activity in Workable settings
- **Event Inspector**: Use n8n execution log to inspect webhook data
- **API Testing**: Test API connectivity and permissions
- **Webhook Testing**: Send test events from Workable dashboard

## UseCases

- **Automated Candidate Communications**: Send welcome emails and status updates to candidates automatically
- **Multi-System Candidate Tracking**: Sync candidate information across ATS, CRM, and other recruiting tools
- **Interview Scheduling Automation**: Trigger interview scheduling workflows when candidates advance to interview stages
- **Hiring Analytics**: Collect and analyze hiring metrics, conversion rates, and pipeline performance
- **Candidate Experience Enhancement**: Improve candidate experience with timely communications and updates
- **Compliance Monitoring**: Track hiring process compliance and generate audit trails
- **Recruiter Notifications**: Alert recruiters and hiring managers to important candidate events
- **Background Check Initiation**: Automatically initiate background checks when candidates reach specific stages
- **Reference Check Automation**: Trigger reference check processes for qualified candidates
- **Offer Management**: Automate offer letter generation and approval workflows
- **Onboarding Preparation**: Prepare onboarding processes for candidates who accept offers
- **Rejection Workflow**: Automate rejection communications and candidate database management
- **Source Attribution**: Track and analyze the effectiveness of different candidate sources
- **Diversity Monitoring**: Monitor hiring diversity metrics and compliance reporting
- **Pipeline Forecasting**: Generate hiring pipeline forecasts and capacity planning data
- **Quality Control**: Monitor application quality and implement screening automation
- **Talent Pool Management**: Maintain and nurture talent pools for future opportunities
- **Integration Synchronization**: Keep multiple recruiting and HR systems synchronized
- **Performance Metrics**: Generate recruiter and hiring manager performance metrics
- **Cost Analysis**: Track cost-per-hire and other recruiting efficiency metrics

### Usage Notes
- The Workable Trigger node only receives recruiting events - it cannot create or modify candidates/jobs
- Webhooks are automatically managed when workflows are activated/deactivated
- Job and stage filters allow targeted monitoring of specific hiring processes
- Event data includes comprehensive candidate and job information for workflow processing
- All webhook endpoints use HTTPS for secure data transmission
- The node supports filtering by both job and stage simultaneously for precise event targeting
- Candidate data includes resume URLs, cover letters, and profile links for comprehensive processing
- Rate limiting and error handling ensure reliable operation with Workable's API