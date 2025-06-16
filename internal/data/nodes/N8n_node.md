# n8n

## Overview

The n8n node provides access to the n8n API, allowing you to handle events and perform actions on your n8n instance. This node enables you to manage workflows, executions, credentials, and audit logs programmatically within your n8n workflows.

## Credentials

- **Credential Name**: `n8nApi`
- **Required**: Yes
- **Configuration**: n8n API credentials with appropriate permissions

## Inputs

- **Main**: Accepts input data for n8n API operations

## Outputs

- **Main**: Returns results from n8n API operations

## Properties

### Resource: n8n Management

#### Operation: n8n API Access

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | Type of n8n resource to work with | Yes |
| Operation | Options | Operation to perform on the resource | Yes |
| Workflow ID | String | Workflow identifier | No |
| Execution ID | String | Execution identifier | No |
| Credential ID | String | Credential identifier | No |
| Additional Fields | Collection | Additional configuration options | No |

## Resources

- **Audit**: Access audit logs and security events
- **Credential**: Manage workflow credentials
- **Execution**: Monitor and manage workflow executions
- **Workflow**: Create, update, and manage workflows

## UseCases

- **Workflow Management** : Programmatically create, update, and manage n8n workflows
- **Execution Monitoring** : Monitor workflow executions and handle failures
- **Credential Management** : Manage and update workflow credentials
- **Audit Tracking** : Track and analyze audit logs for security and compliance
- **Automation Orchestration** : Build meta-workflows that manage other workflows

