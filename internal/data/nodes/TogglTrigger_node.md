# Toggl Trigger Node

## Overview

The Toggl Trigger node starts workflows when new time entries are created in your Toggl Track account. This polling-based trigger node monitors your Toggl account for new time tracking activities and automatically executes workflows when time entries are added, enabling automated time tracking workflows, project management automation, and productivity analytics.

## Credentials

- Name: togglApi, Required: Yes

### Credential Configuration
To configure Toggl API credentials:
1. Log into your Toggl Track account
2. Use your Toggl account email address as the username
3. Use your Toggl account password (or generate an API token)
4. For enhanced security, consider using an API token instead of your password
5. Test the connection to verify credentials are working

## Inputs

None - This is a trigger node that starts workflows based on Toggl time entry events.

## Outputs

- Main: Outputs time entry data when new time entries are detected

## Properties

### Resource: Time Entry Events

#### Operation: New Time Entry Trigger

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event | options | Type of event to monitor | Yes | New Time Entry |

### Event Types

#### New Time Entry
Triggered when a new time entry is created in your Toggl account.

**Trigger Conditions:**
- Time entry is newly created
- Time entry is started or stopped
- Time entry is manually added

**Polling Mechanism:**
- Checks for new entries since last execution
- Uses timestamp-based filtering
- Maintains state between executions

### Output Data Structure

When triggered, the node outputs time entry information:

| Field | Type | Description |
|-------|------|-------------|
| id | number | Unique time entry identifier |
| workspace_id | number | Workspace ID where entry belongs |
| project_id | number | Project ID (if assigned) |
| task_id | number | Task ID (if assigned) |
| billable | boolean | Whether entry is billable |
| start | string | Start timestamp (ISO 8601) |
| stop | string | Stop timestamp (ISO 8601) |
| duration | number | Duration in seconds |
| description | string | Time entry description |
| tags | array | Array of tag names |
| tag_ids | array | Array of tag IDs |
| duronly | boolean | Duration-only entry flag |
| at | string | Last modified timestamp |
| server_deleted_at | string | Deletion timestamp (if deleted) |
| user_id | number | User ID who created entry |
| uid | number | User ID (alternative field) |

### Time Entry Data Example

```json
{
  "id": 1894738525,
  "workspace_id": 3134975,
  "project_id": 193791154, 
  "task_id": null,
  "billable": true,
  "start": "2023-05-15T09:00:00+00:00",
  "stop": "2023-05-15T10:30:00+00:00",
  "duration": 5400,
  "description": "Working on client project",
  "tags": ["development", "client-work"],
  "tag_ids": [12345, 67890],
  "duronly": false,
  "at": "2023-05-15T10:30:15+00:00",
  "server_deleted_at": null,
  "user_id": 1234567,
  "uid": 1234567
}
```

### Advanced Features

#### State Management
- **Last Check Tracking**: Maintains timestamp of last check
- **Duplicate Prevention**: Prevents duplicate triggering for same entries
- **Incremental Processing**: Only processes new entries since last execution

#### Polling Optimization
- **Efficient Polling**: Uses timestamp-based filtering to minimize API calls
- **Rate Limiting**: Respects Toggl API rate limits
- **Error Recovery**: Handles API errors gracefully with retry logic

#### Data Processing
- **Complete Entry Data**: Returns full time entry information
- **Metadata Inclusion**: Includes project, task, and tag information
- **Timestamp Formatting**: ISO 8601 formatted timestamps for consistency

### API Integration

#### Toggl API v9
- **Base URL**: `https://api.track.toggl.com/api/v9`
- **Endpoint**: `/me/time_entries`
- **Authentication**: Basic HTTP authentication
- **Rate Limiting**: Respects Toggl's API rate limits

#### Request Parameters
- **since**: Filters entries created after specific timestamp
- **before**: Filters entries created before specific timestamp
- **meta**: Includes additional metadata in response

### Error Handling

#### Authentication Errors
- **Invalid Credentials**: Clear error messages for authentication failures
- **Account Issues**: Handling of suspended or deleted accounts
- **Token Expiration**: Graceful handling of expired tokens

#### API Errors
- **Rate Limiting**: Automatic retry with exponential backoff
- **Network Errors**: Retry logic for connectivity issues
- **Server Errors**: Proper error logging and recovery

#### Data Validation
- **Entry Validation**: Validates time entry data structure
- **Date Handling**: Robust date parsing and timezone handling
- **State Persistence**: Maintains polling state across executions

### Integration Patterns

#### Project Management
1. **Time Tracking**: Automatically track project time in external systems
2. **Billing Integration**: Create invoices based on billable time entries
3. **Resource Allocation**: Monitor team time allocation across projects
4. **Productivity Analytics**: Generate productivity reports from time data

#### Team Management
1. **Attendance Tracking**: Monitor when team members start/stop work
2. **Workload Monitoring**: Track individual and team workloads
3. **Performance Metrics**: Calculate productivity metrics from time data
4. **Capacity Planning**: Use historical time data for capacity planning

#### Client Management
1. **Client Reporting**: Generate client reports from tracked time
2. **Billing Automation**: Automate billing based on time entries
3. **Project Updates**: Send project updates based on time tracking
4. **Invoice Generation**: Create invoices from billable time entries

### Best Practices

#### Polling Configuration
- **Appropriate Intervals**: Set reasonable polling intervals to balance real-time updates with API usage
- **State Management**: Ensure polling state is properly maintained
- **Error Handling**: Implement proper error handling for polling failures

#### Data Processing
- **Filtering**: Process only relevant time entries for your workflow
- **Validation**: Validate time entry data before processing
- **Transformation**: Transform data as needed for downstream systems

#### Performance Optimization
- **Efficient Queries**: Use timestamp filtering to minimize data transfer
- **Batch Processing**: Process multiple time entries efficiently
- **Caching**: Cache project and tag information to reduce API calls

### Security Considerations

#### Authentication
- **Secure Credentials**: Store Toggl credentials securely
- **Token Management**: Use API tokens instead of passwords when possible
- **Access Control**: Limit access to time tracking data

#### Data Privacy
- **Time Data**: Handle time tracking data according to privacy policies
- **Client Information**: Protect client-related time tracking information
- **Team Privacy**: Respect team member privacy in time tracking data

### Troubleshooting

#### Common Issues
- **Missing Events**: Check polling interval and API credentials
- **Duplicate Events**: Verify state management and deduplication logic
- **Authentication Errors**: Validate Toggl credentials and permissions
- **Rate Limits**: Monitor API usage and implement proper rate limiting

#### Debugging Tools
- **Toggl Web Interface**: Check time entries in Toggl web interface
- **API Testing**: Test API connectivity and authentication
- **State Inspection**: Monitor polling state and last check timestamps
- **Log Analysis**: Review execution logs for error patterns

### Migration and Updates

#### API Version Compatibility
- **Toggl API v9**: Compatible with current Toggl Track API
- **Legacy Support**: Maintains compatibility with existing integrations
- **Future Updates**: Prepared for future API version updates

#### Configuration Updates
- **Credential Migration**: Smooth migration of credential configurations
- **State Preservation**: Maintains polling state during updates
- **Backward Compatibility**: Preserves existing workflow configurations

## UseCases

- **Automated Time Reporting**: Generate time reports automatically when team members log time
- **Project Budget Tracking**: Monitor project time against budgets and alert when thresholds are reached
- **Client Billing Automation**: Automatically create invoices or billing records from tracked time
- **Productivity Analytics**: Analyze team productivity patterns and generate insights from time data
- **Attendance Monitoring**: Track when team members start and stop work for attendance records
- **Project Status Updates**: Send project status updates to stakeholders based on time tracking activity
- **Resource Allocation Tracking**: Monitor how team resources are allocated across different projects
- **Overtime Alerts**: Alert managers when team members exceed normal working hours
- **Time Entry Validation**: Validate time entries against project requirements and company policies
- **Integration with Project Management**: Sync time tracking data with project management tools
- **Performance Metrics**: Calculate billability rates and productivity metrics from time data
- **Client Communication**: Automatically update clients on project progress based on time tracking
- **Team Capacity Planning**: Use historical time data for future capacity and resource planning
- **Compliance Tracking**: Ensure time tracking compliance with company or client requirements
- **Work Pattern Analysis**: Analyze work patterns and identify optimization opportunities
- **Task Time Estimation**: Improve task time estimation using historical time tracking data
- **Employee Workload Balancing**: Monitor and balance workloads across team members
- **Revenue Recognition**: Trigger revenue recognition processes based on billable time entries
- **Project Profitability Analysis**: Calculate project profitability using actual time vs. estimates
- **Time Theft Prevention**: Monitor for unusual time tracking patterns or potential time theft

### Usage Notes
- The Toggl Trigger node only monitors for new time entries - it cannot create or modify time entries
- Polling is based on time entry creation timestamps, not start/stop times
- The node maintains state between executions to prevent duplicate triggering
- All time data includes timezone information for accurate time calculations
- Billable time entries can be easily identified using the billable flag
- Project and task IDs can be used to filter or categorize time entries in workflows
- Duration is provided in seconds and includes both manual entries and tracked time