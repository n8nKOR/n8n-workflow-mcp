# Bitbucket Trigger

## Overview

The Bitbucket Trigger node handles Bitbucket events via webhooks. This trigger node automatically starts workflows when specific events occur in your Bitbucket repositories or workspaces, such as code pushes, pull requests, issue creation, and other repository activities. It enables real-time automation based on Bitbucket activities.

## Credentials

This node requires Bitbucket API credentials:
- **Credential Name**: `bitbucketApi`
- **Required Fields**: 
  - Username: Your Bitbucket username
  - App Password: Your Bitbucket app password

To obtain API credentials:
1. Log into your Bitbucket account
2. Navigate to Personal settings > App passwords
3. Create a new app password with appropriate permissions
4. Use your username and the generated app password

## Inputs

This node is a trigger node and has no inputs.

## Outputs

- **Main**: Returns JSON data containing the webhook event information from Bitbucket

## Properties

### Resource: Repository

#### Operation: Listen for Repository Events

| Field Name | Type | Description | Required |
|---|---|---|---|
| Workspace | String | Workspace name or ID for the repository | Yes |
| Repository | String | Repository name or slug to monitor | Yes |
| Events | MultiSelect | Repository events to listen for (from API) | Yes |

### Resource: Workspace

#### Operation: Listen for Workspace Events

| Field Name | Type | Description | Required |
|---|---|---|---|
| Workspace | String | Workspace name or ID to monitor | Yes |
| Events | MultiSelect | Workspace events to listen for (from API) | Yes |

## API Details

### Webhook Management
- **Webhook URL**: Automatically generated n8n webhook endpoint
- **HTTP Method**: POST
- **Response Mode**: On received (immediate response)
- **Authentication**: Uses Bitbucket API credentials for webhook management

### Event Categories

#### Repository Events
- **Push Events**: Code pushes, branch creation/deletion
- **Pull Request Events**: PR creation, updates, merges, comments
- **Issue Events**: Issue creation, updates, comments, state changes
- **Build Events**: Build status changes, pipeline events
- **Fork Events**: Repository forks and fork updates

#### Workspace Events
- **Repository Events**: Repository creation, deletion, access changes
- **Team Events**: Team member additions, removals, permission changes
- **Project Events**: Project creation, updates, deletions
- **Webhook Events**: Webhook management events

### Webhook Security
- **UUID Validation**: Verifies webhook UUID in headers (`x-hook-uuid`)
- **Active Status**: Only active webhooks trigger workflows
- **Automatic Cleanup**: Webhooks deleted when workflow is deactivated

## UseCases

### DevOps Automation
- **CI/CD Pipeline Triggers**: Start build pipelines when code is pushed to specific branches
- **Automated Testing**: Trigger test suites on pull request creation or updates
- **Deployment Automation**: Deploy applications when pull requests are merged to main
- **Code Quality Checks**: Run linting and security scans on new commits
- **Release Management**: Automate release creation and changelog generation

### Development Workflow Automation
- **Pull Request Management**: Auto-assign reviewers based on file changes or team rules
- **Issue Tracking**: Create tickets in external systems when issues are created
- **Code Review Notifications**: Send Slack/email alerts for pull request events
- **Branch Protection**: Enforce branch policies and merge requirements
- **Commit Message Validation**: Validate commit message formats and standards

### Project Management Integration
- **Task Synchronization**: Sync Bitbucket issues with Jira, Asana, or other project tools
- **Progress Tracking**: Update project boards when pull requests are merged
- **Time Tracking**: Log development time based on commit and PR activities
- **Sprint Planning**: Track code delivery against sprint commitments
- **Reporting**: Generate development velocity and contribution reports

### Security and Compliance
- **Access Monitoring**: Track repository access and permission changes
- **Security Scanning**: Trigger security scans on new code pushes
- **Compliance Auditing**: Log all code changes for compliance requirements
- **Vulnerability Management**: Alert security teams about potential vulnerabilities
- **Access Control**: Enforce repository access policies and approvals

### Communication and Collaboration
- **Team Notifications**: Send customized alerts to team communication channels
- **Documentation Updates**: Auto-update documentation when code changes
- **Stakeholder Reporting**: Generate reports for non-technical stakeholders
- **Change Announcements**: Broadcast important code changes to relevant teams
- **Knowledge Sharing**: Share code insights and best practices across teams

## Error Handling

### Common Error Scenarios
- **Authentication Failures**: Invalid username or app password
- **Permission Errors**: Insufficient permissions for webhook management
- **Network Issues**: Connectivity problems with Bitbucket API
- **Webhook Conflicts**: Existing webhooks with same URL
- **Rate Limiting**: API rate limit exceeded

### Troubleshooting Steps
1. **Verify Credentials**: Test API credentials in node configuration
2. **Check Permissions**: Ensure user has admin access to repository/workspace
3. **Validate Events**: Confirm selected events are valid for the resource type
4. **Monitor Webhook Status**: Check webhook is active in Bitbucket settings
5. **Review Headers**: Verify webhook UUID matches in incoming requests

### Error Recovery
- **Automatic Reconnection**: Webhook recreated if connection is lost
- **Duplicate Prevention**: Checks for existing webhooks before creation
- **Graceful Cleanup**: Removes webhooks when workflow is deactivated
- **Validation**: Validates webhook payload and UUID before processing

## Advanced Configuration

### Event Selection Strategy
- **Granular Control**: Select specific events to avoid unnecessary triggers
- **Performance Optimization**: Limit events to reduce webhook volume
- **Conditional Logic**: Use event data to determine workflow actions
- **Event Filtering**: Filter events based on branch, author, or other criteria

### Webhook Management
- **Resource Types**: Choose between repository-level and workspace-level webhooks
- **Event Scoping**: Understand event hierarchy and inheritance
- **Payload Structure**: Access event data structure for conditional processing
- **Header Analysis**: Use webhook headers for additional context

### Security Best Practices
- **Credential Security**: Store app passwords securely and rotate regularly
- **Access Limitation**: Use minimal required permissions for API access
- **Webhook Validation**: Always validate webhook UUID and payload structure
- **Network Security**: Ensure n8n webhook endpoints are properly secured

### Integration Patterns
- **Event Chaining**: Chain multiple workflow triggers for complex automation
- **Conditional Workflows**: Use event data to control workflow execution paths
- **Error Handling**: Implement robust error handling for failed operations
- **Monitoring**: Set up monitoring for webhook health and performance

### Resource: Bitbucket Webhook

#### Operation: Listen for Events

| Field Name | Type | Description | Required |
|---|---|---|---|
| Events | Multi-Options | Bitbucket events to listen for | Yes |
| Repository | String | Repository to monitor (optional for workspace events) | No |
| Workspace | String | Workspace to monitor | No |

#### Options

| Field Name | Type | Description | Required |
|---|---|---|---|
| Filter | String | Additional event filtering criteria | No |
| Include Headers | Boolean | Whether to include webhook headers | No |
## UseCases

- **CI/CD Pipeline Triggers**: Automatically start build and deployment processes when code is pushed
- **Code Review Automation**: Automatically assign reviewers and send notifications when pull requests are created
- **Issue Management Automation**: Automatically assign team members and update project boards when new issues are created
- **Notification Systems**: Send automatic alerts to team channels when important repository events occur
- **Backup and Synchronization**: Automatically sync code changes to backup systems
- **Quality Management**: Automatically run code quality checks and tests when commits are made
- **Release Management**: Trigger release workflows when tags are created or branches are merged
- **Security Monitoring**: Monitor for security-related events and trigger security workflows
- **Documentation Updates**: Automatically update documentation when code changes are made
- **Team Collaboration**: Enhance team collaboration by automating routine tasks and notifications 