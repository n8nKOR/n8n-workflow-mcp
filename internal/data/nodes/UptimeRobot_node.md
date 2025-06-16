# UptimeRobot Node

## Overview

Consume UptimeRobot API

## Credentials

- Name: uptimeRobotApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: monitor

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Friendly Name | string | The friendly name of the monitor | Yes |  |
| Type | options | The type of the monitor | Yes |  |
| URL | string | The URL/IP of the monitor | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether the alert contacts set for the monitor to be returned | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the monitor | Yes |  |
| Update Fields | collection | The friendly name of the monitor | No |  |

### Resource: alertContact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Friendly Name | string | The friendly name of the alert contact | Yes |  |
| Type | options | The type of the alert contact | Yes |  |
| Value | string | The correspondent value for the alert contact type | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Alert contact IDs separated with dash, e.g. 236-1782-4790 | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the alert contact | Yes |  |
| Update Fields | collection | The friendly name of the alert contact | No |  |

### Resource: maintenanceWindow

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Duration (Minutes) | number | The maintenance window activation period (minutes) | Yes |  |
| Friendly Name | string | The friendly name of the maintenance window | Yes |  |
| Type | options | The type of the maintenance window | Yes |  |
| Week Day | options |  | No |  |
| Month Day | number |  | No |  |
| Start Time | dateTime | The maintenance window start datetime | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Maintenance windows IDs separated with dash, e.g. 236-1782-4790 | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the maintenance window | Yes |  |
| Duration (Minutes) | number | The maintenance window activation period (minutes) | Yes |  |
| Update Fields | collection | The friendly name of the maintenance window | No |  |

### Resource: publicStatusPage

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Friendly Name | string | The friendly name of the status page | Yes |  |
| Monitor IDs | string | Monitor IDs to be displayed in status page (the values are separated with a dash (-) or 0 for all monitors) | Yes |  |
| Additional Fields | collection | The domain or subdomain that the status page will run on | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Public status pages IDs separated with dash, e.g. 236-1782-4790 | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the public status page | Yes |  |
| Update Fields | collection | The domain or subdomain that the status page will run on | No |  |

## UseCases

- [Website Monitoring] : Monitor website uptime and availability 24/7
- [Server Monitoring] : Track server health and performance metrics
- [API Monitoring] : Monitor API endpoints for availability and response times
- [Downtime Alerts] : Receive immediate alerts when services go down
- [Performance Tracking] : Track website performance and response time metrics
- [SLA Monitoring] : Monitor service level agreements and uptime commitments
- [Multi-Location Monitoring] : Monitor services from multiple geographic locations
- [SSL Certificate Monitoring] : Track SSL certificate expiration and validity
- [Incident Management] : Manage incident response and downtime tracking
- [Status Page Creation] : Create public status pages for service transparency
- [Maintenance Scheduling] : Schedule maintenance windows and notifications
- [Team Collaboration] : Collaborate on incident response and monitoring
- [Integration Automation] : Integrate monitoring alerts with team communication tools
- [Uptime Reporting] : Generate uptime reports and availability statistics
- [Customer Communication] : Communicate service status to customers and stakeholders

