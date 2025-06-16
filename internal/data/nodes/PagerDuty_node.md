# PagerDuty Node

## Overview

Consume PagerDuty API for incident management and operational alerting workflows. PagerDuty is a digital operations management platform that provides incident response, alerting, and on-call management capabilities. This node enables integration with PagerDuty's comprehensive incident lifecycle management system for automated DevOps and IT operations.

## Credentials

- Name: pagerDutyApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: incident

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | A succinct description of the nature, symptoms, cause, or effect of the incident | Yes |  |
| Service Name or ID | options | The incident will be created on this service. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Email | string | The email address of a valid user associated with the account making the request | Yes |  |
| Additional Fields | collection | Delegate this incident to the specified escalation policy. Cannot be specified if an assignee is given. Choose from the list, or specify an ID using an <a href= | No |  |
| Conference Bridge | fixedCollection | Phone numbers should be formatted like +1 415-555-1212,,,,1234#, where a comma (,) represents a one-second wait and pound (#) completes access code input | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Incident ID | string | Unique identifier for the incident | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | When set to all, the since and until parameters and defaults are ignored | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Incident ID | string | Unique identifier for the incident | Yes |  |
| Email | string | The email address of a valid user associated with the account making the request | Yes |  |
| Update Fields | collection | Escalate the incident to this level in the escalation policy | No |  |
| Conference Bridge | fixedCollection | Phone numbers should be formatted like +1 415-555-1212,,,,1234#, where a comma (,) represents a one-second wait and pound (#) completes access code input | No |  |

### Resource: incidentNote

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Incident ID | string | Unique identifier for the incident | Yes |  |
| Content | string | The note content | Yes |  |
| Email | string | The email address of a valid user associated with the account making the request | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Incident ID | string | Unique identifier for the incident | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: logEntry

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Log Entry ID | string | Unique identifier for the log entry | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Additional details to include | No |  |

### Resource: user

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | Unique identifier for the user | Yes |  |

## UseCases

- **Automated Incident Management**: Create and manage incidents automatically based on system alerts, monitoring triggers, and application errors
- **DevOps Pipeline Integration**: Integrate PagerDuty with CI/CD pipelines to create incidents for deployment failures and automatically resolve them on successful fixes
- **System Monitoring Alerts**: Connect monitoring systems to automatically create PagerDuty incidents when system metrics exceed thresholds or services become unavailable
- **Incident Response Automation**: Automate incident escalation, team notifications, and response coordination based on severity levels and response times
- **On-Call Schedule Management**: Manage on-call rotations, schedule changes, and team availability through automated workflows and notifications
- **Service Status Communication**: Update incident status, add notes, and communicate progress to stakeholders throughout the incident resolution process
- **Escalation Policy Automation**: Implement automated escalation workflows that route incidents to appropriate teams based on service type, severity, and time of day
- **Incident Post-Mortem Integration**: Automatically gather incident data, timeline information, and response metrics for post-incident analysis and reporting
- **Multi-System Alert Correlation**: Correlate alerts from multiple monitoring systems to reduce noise and create meaningful incidents instead of alert storms
- **Customer Communication Workflows**: Automatically update status pages, send customer notifications, and manage external communication during service disruptions

