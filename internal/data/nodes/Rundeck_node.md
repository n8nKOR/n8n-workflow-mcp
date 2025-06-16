# Rundeck

## Overview

The Rundeck node allows you to manage and execute jobs using the Rundeck automation platform API. Rundeck is a runbook automation and job scheduling platform that enables you to safely automate operations procedures in production environments. This node enables you to integrate Rundeck job execution and management into your n8n workflows for DevOps automation, scheduled tasks, and infrastructure management.

## Credentials

This node requires Rundeck API credentials:
- **Credential Name**: `rundeckApi`
- **Required Fields**: 
  - URL: Your Rundeck server URL
  - Token: Your Rundeck API authentication token

To obtain API credentials:
1. Log into your Rundeck server
2. Navigate to your User Profile
3. Go to "User API Tokens" section
4. Generate a new API token
5. Copy the token and server URL for use in n8n

## Inputs

- **Main**: Job execution parameters, job IDs, and configuration data

## Outputs

- **Main**: Returns JSON response with job execution status, metadata, execution ID, and detailed job information

## Properties

### Resource: Job

#### Operation: Execute

| Field Name | Type | Description | Required |
|---|---|---|---|
| Job ID | String | The Rundeck job ID to execute | Yes |
| Arguments | Fixed Collection | Key-value pairs for job parameters | No |
| Filter | String | Filter Rundeck nodes by name for targeted execution | No |

##### Arguments Configuration:
- **Key**: Parameter name as defined in the Rundeck job
- **Value**: Parameter value to pass to the job
- Supports multiple arguments for complex job configurations
- Arguments are converted to Rundeck's expected argument string format

#### Operation: Get Metadata

| Field Name | Type | Description | Required |
|---|---|---|---|
| Job ID | String | The job ID to retrieve metadata for | Yes |

**Metadata Information Returned:**
- Job name, description, and group
- Project information
- Job configuration details
- Execution options and settings
- Node filter configurations
- Scheduled execution information
- Plugin configurations

## UseCases

- **Infrastructure Automation** : Execute deployment scripts, server maintenance, and configuration management tasks
- **CI/CD Pipeline Integration** : Trigger deployment jobs as part of continuous integration workflows
- **Scheduled Maintenance** : Automate routine maintenance tasks like backups, updates, and cleanup operations
- **Incident Response** : Execute predefined runbooks for system recovery and troubleshooting procedures
- **Environment Management** : Manage development, staging, and production environment deployments
- **Monitoring Integration** : Trigger remediation jobs based on monitoring alerts and system events
- **Data Processing** : Execute ETL jobs, data backups, and database maintenance procedures
- **Security Operations** : Run security scans, patch management, and compliance checking procedures
- **Application Lifecycle** : Manage application starts, stops, restarts, and health checks
- **Resource Scaling** : Execute auto-scaling procedures based on load metrics and demand
- **Backup and Recovery** : Automate backup procedures and disaster recovery workflows
- **Compliance Automation** : Execute compliance checks and generate audit reports on schedule
