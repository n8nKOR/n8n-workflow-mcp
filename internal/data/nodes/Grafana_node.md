# Grafana Node

## Overview

Consume the Grafana API

## Credentials

- Name: grafanaApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: dashboard

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the dashboard to create | Yes |  |
| Additional Fields | collection | Folder to create the dashboard in - if the folder is unspecified, the dashboard will be saved to the General folder. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Dashboard UID or URL | string | Unique alphabetic identifier or URL of the dashboard to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Dashboard UID or URL | string | Unique alphabetic identifier or URL of the dashboard to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Dashboard UID or URL | string | Unique alphabetic identifier or URL of the dashboard to update | Yes |  |
| Update Fields | collection | Folder to move the dashboard into - if the folder is unspecified, the dashboard will be saved to the General folder. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: team

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the team to create | Yes |  |
| Additional Fields | collection | Email of the team to create | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team ID | string | ID of the team to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team ID | string | ID of the team to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Name of the team to filter by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team ID | string | ID of the team to update | Yes |  |
| Update Fields | collection | Email of the team to update | No |  |

### Resource: teamMember

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User Name or ID | options | User to add to a team. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Team Name or ID | options | Team to add the user to. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User Name or ID | options | User to remove from the team. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Team Name or ID | options | Team to remove the user from. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Team to retrieve all members from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: user

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | ID of the user to update | Yes |  |
| Update Fields | collection | New role for the user | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | ID of the user to delete | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **Dashboard Management**: Automate Grafana dashboard creation, updates, and deployment across multiple environments
- **Monitoring Automation**: Set up automated monitoring dashboards for infrastructure, applications, and business metrics
- **Team and Access Management**: Manage Grafana teams, user permissions, and access control for dashboard security
- **DevOps Integration**: Integrate dashboard management with CI/CD pipelines for automated monitoring deployment
- **Multi-tenant Operations**: Manage dashboards and teams across multiple organizations and business units
- **Infrastructure as Code**: Implement infrastructure as code practices for Grafana configuration and dashboard management
- **Alerting Integration**: Connect Grafana dashboards with alerting systems and notification workflows
- **Performance Monitoring**: Automate performance dashboard creation for application and system monitoring
- **Business Intelligence**: Create and manage business intelligence dashboards for operational and executive reporting
- **Compliance and Auditing**: Maintain audit trails for dashboard changes and ensure compliance with monitoring standards

