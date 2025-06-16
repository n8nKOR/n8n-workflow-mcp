# ErrorTrigger

## Overview

The ErrorTrigger node is a specialized trigger that activates workflows when errors occur in other n8n workflows. This powerful error handling mechanism enables automated error monitoring, notification systems, and recovery workflows. It provides comprehensive error context including execution details, error messages, stack traces, and workflow information to facilitate effective error management and debugging.

**Key Features:**
- Automatic activation on workflow errors
- Comprehensive error context and execution details
- Integration with n8n's error workflow system
- Support for both execution and trigger errors
- Manual testing capability with mock data
- Single node per workflow limitation for controlled error handling

## Credentials

The ErrorTrigger node does not require any credentials:

- **No Authentication**: The node operates within n8n's internal system
- **System Integration**: Automatically receives error data from the n8n execution engine
- **No External Services**: No external API connections required

## Inputs

The ErrorTrigger node has no inputs as it is a trigger node:

- **No Direct Inputs**: Trigger nodes initiate workflows rather than receiving data
- **System Triggered**: Activated automatically by n8n's error handling system
- **Error Data Injection**: Receives error data directly from the execution engine

## Outputs

The ErrorTrigger node provides a single output with comprehensive error information:

- **Main Output**: Contains detailed error data and execution context
- **Error Context**: Complete information about the failed workflow and execution
- **Execution Details**: Information about the execution that encountered the error

## Properties

### Resource: Error Trigger

#### Operation: Trigger on Workflow Errors

The ErrorTrigger node has minimal configuration as it operates automatically:

| Field Name | Type | Description | Required |
|---|---|---|---|
| Notice | Information | Displays setup instructions and documentation link | Read-only |

#### Configuration Notice

The node displays an informational message:
> "This node will trigger when there is an error in another workflow, as long as that workflow is set up to do so. [More info](https://docs.n8n.io/integrations/core-nodes/n8n-nodes-base.errortrigger)"

## UseCases

### Error Monitoring and Alerting
- **Slack Notifications**: Send immediate alerts to team channels when workflows fail
- **Email Alerts**: Notify administrators of critical workflow failures
- **SMS Notifications**: Send urgent alerts for business-critical workflow errors
- **Dashboard Updates**: Update monitoring dashboards with error status
- **Incident Creation**: Automatically create tickets in incident management systems

### Error Recovery and Retry Logic
- **Automatic Retries**: Implement custom retry logic for transient failures
- **Fallback Workflows**: Execute alternative workflows when primary ones fail
- **Data Recovery**: Restore or recreate lost data from failed executions
- **Service Restoration**: Restart services or connections that caused failures
- **Manual Intervention**: Queue failed workflows for manual review and correction

### Logging and Audit Systems
- **Error Logging**: Send detailed error information to logging systems
- **Audit Trails**: Maintain comprehensive records of workflow failures
- **Performance Analysis**: Analyze error patterns and workflow reliability
- **Compliance Reporting**: Generate reports for regulatory compliance
- **Debugging Support**: Provide developers with detailed error context

### Business Process Management
- **Escalation Procedures**: Implement escalation workflows for critical failures
- **Approval Workflows**: Route failed operations to supervisors for approval
- **Customer Notifications**: Inform customers about service disruptions
- **SLA Monitoring**: Track service level agreement compliance
- **Quality Assurance**: Monitor and improve workflow quality over time

### DevOps and Operations
- **CI/CD Pipeline Monitoring**: Monitor deployment and integration workflows
- **Infrastructure Monitoring**: Detect and respond to system failures
- **Backup Verification**: Ensure backup processes complete successfully
- **Security Incident Response**: Respond to security-related workflow failures
- **Capacity Management**: Monitor and respond to resource-related failures

