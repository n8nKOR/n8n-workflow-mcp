# Clockify Node

## Overview

The Clockify node allows you to consume the Clockify REST API for time tracking and project management. Clockify is a time tracking tool that helps teams and individuals track work hours across projects, clients, and tasks. This node enables you to manage workspaces, clients, projects, tasks, tags, users, and time entries within your n8n workflows, making it easy to automate time tracking and reporting processes.

## Credentials

This node requires Clockify API credentials:
- **API Key**: Your Clockify API key

To obtain API credentials:
1. Log into your Clockify account
2. Navigate to Settings → Profile Settings
3. Generate an API key in the API section
4. Copy the API key for use in n8n

## Inputs

- **Main**: JSON data for creating/updating time tracking records

## Outputs

- **Main**: Returns JSON data from Clockify API responses

## Properties

### Resource: workspace

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: client

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Client Name | string | Name of client being created | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Client ID | string | ID of the client to delete | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Client ID | string | ID of the client to retrieve | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | If provided, clients will be filtered by name | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Client ID | string | ID of the client to update | No |  |
| Name | string | New name for the client | Yes |  |
| Update Fields | collection | Address of client being created/updated | No |  |

### Resource: project

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Name | string | Name of project being created | Yes |  |
| Additional Fields | collection | Additional fields to configure for the project | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project ID | string | ID of the project | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project ID | string | ID of the project | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Additional fields to filter projects | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project ID | string | ID of the project | Yes |  |
| Update Fields | collection | Fields to update for the project | No |  |

### Resource: tag

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of tag being created | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag ID | string | ID of the tag | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Additional fields to configure for tags | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag ID | string | ID of the tag | Yes |  |
| Update Fields | collection | Fields to update for the tag | No |  |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task Name | string | Name of task to create | Yes |  |
| Additional Fields | collection | Additional configuration fields for the task | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of task to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of task to get | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Text to match in the task name | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of task to update | Yes |  |
| Update Fields | collection | Fields to update for the task | No |  |

### Resource: user

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Additional fields to filter users | No |  |

### Resource: timeEntry

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start | dateTime | Start time for the time entry | Yes |  |
| Additional Fields | collection | Additional configuration fields for the time entry | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Time Entry ID | string | ID of the time entry to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Time Entry ID | string | ID of the time entry to retrieve | Yes |  |
| Additional Fields | collection | Additional fields to include in the response | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Time Entry ID | string | ID of the time entry to update | Yes |  |
| Update Fields | collection | Fields to update for the time entry | No |  |

## UseCases

- **Automated Time Tracking**: Automatically create time entries based on calendar events, task completion, or project milestones
- **Project Budget Monitoring**: Track project time and costs against budgets, sending alerts when thresholds are exceeded
- **Client Billing Automation**: Generate billable hour reports and automate invoice creation based on tracked time
- **Team Productivity Analytics**: Analyze team performance by aggregating time tracking data across projects and clients
- **Task Management Integration**: Sync time entries with project management tools and automatically track time spent on specific tasks
- **Timesheet Automation**: Automatically populate timesheets from various sources and ensure accurate time reporting

