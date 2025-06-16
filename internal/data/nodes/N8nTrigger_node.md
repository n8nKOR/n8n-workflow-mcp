# N8n Trigger Node

## Overview

The N8n Trigger node monitors internal n8n instance events and automatically triggers workflows when specific system events occur. This trigger enables automation based on n8n's internal lifecycle events such as instance startup, workflow activation, and workflow updates, making it perfect for system monitoring, audit logging, and administrative automation.

## Credentials

No credentials required - this node monitors internal n8n events and does not require external authentication.

## Inputs

This node is a trigger node and does not accept input connections.

## Outputs

- **Main**: Returns event data when n8n system events occur

## Properties

### Resource: N8n Instance Events

#### Operation: Monitor System Events

| Field Name | Type | Description | Required |
|---|---|---|---|
| Events | Multi Options | System events that will trigger the workflow | Yes |

#### Available Events

| Event | Description | When It Triggers |
|---|---|---|
| `n8n.instance.started` | N8n instance startup event | When the n8n instance starts or restarts |
| `n8n.workflow.activated` | Workflow activation event | When the workflow containing this trigger is activated |
| `n8n.workflow.updated` | Active Workflow Updated event | When the workflow containing this trigger is updated |

## Event Output Data Structure

### Instance Started Event
```json
{
  "eventName": "n8n.instance.started",
  "timestamp": "2024-01-15T10:30:00.000Z",
  "instanceId": "n8n-instance-001",
  "version": "1.0.0",
  "environment": "production"
}
```

### Workflow Activated Event
```json
{
  "eventName": "n8n.workflow.activated",
  "timestamp": "2024-01-15T10:30:00.000Z",
  "workflowId": "123",
  "workflowName": "Customer Processing",
  "userId": "user-456",
  "active": true
}
```

### Workflow Updated Event
```json
{
  "event": "Workflow updated",
  "timestamp": "2024-01-15T10:30:00.000Z",
  "workflow_id": "123"
}
```

## UseCases

### System Administration
- **Instance Health Monitoring**: Monitor n8n instance startup and status
- **Backup Automation**: Trigger backup processes on instance restart
- **System Maintenance**: Schedule maintenance tasks on startup
- **Resource Monitoring**: Track system resource usage on events

### Workflow Management
- **Deployment Tracking**: Log workflow activations and updates
- **Change Management**: Track and approve workflow changes
- **Version Control**: Integrate with version control systems
- **Rollback Procedures**: Implement automated rollback triggers

### DevOps Integration
- **CI/CD Pipeline Integration**: Trigger deployment pipelines
- **Environment Synchronization**: Sync workflows across environments
- **Automated Testing**: Run tests when workflows are updated
- **Performance Monitoring**: Track workflow performance metrics

### Audit and Compliance
- **Audit Logging**: Maintain detailed logs of system changes
- **Compliance Reporting**: Generate compliance reports on changes
- **Change Approval**: Implement change approval workflows
- **Security Monitoring**: Monitor for unauthorized changes

### Notification and Alerting
- **Team Notifications**: Notify teams of system events
- **Stakeholder Updates**: Update stakeholders on deployments
- **Error Alerting**: Send alerts for system issues
- **Status Dashboard**: Update status dashboards in real-time

### Integration Workflows
- **Third-party Synchronization**: Sync with external systems
- **Data Pipeline Triggers**: Trigger data processing pipelines
- **Service Orchestration**: Coordinate microservice deployments
- **API Gateway Updates**: Update API configurations on changes

## Troubleshooting

### Event Not Triggering
- Verify event selection matches expected trigger condition
- Check workflow activation status
- Ensure n8n instance has proper permissions
- Review execution logs for error messages

### Multiple Triggers
- Events may fire multiple times during rapid changes
- Implement deduplication logic if needed
- Use delays or queues for high-frequency events

### Performance Issues
- Monitor workflow execution times
- Optimize heavy operations triggered by events
- Consider async processing for non-critical tasks

## Related Nodes

- **Webhook Trigger**: For external system events
- **Schedule Trigger**: For time-based automation
- **Switch**: For event-specific logic branching
- **Set**: For event data manipulation
- **HTTP Request**: For external system notifications

