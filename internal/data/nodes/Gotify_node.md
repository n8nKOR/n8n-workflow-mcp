# Gotify Node

## Overview

Consume Gotify API

## Credentials

- Name: gotifyApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: message

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message | string | The message to send, If using Markdown add the Content Type option | Yes |  |
| Additional Fields | collection | The priority of the message | No |  |
| Options | collection | The message content type | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **System Notifications**: Send automated system alerts, server status updates, and infrastructure monitoring notifications
- **Application Alerts**: Integrate application monitoring with real-time push notifications for critical events and errors
- **DevOps Notifications**: Automate deployment notifications, build status alerts, and CI/CD pipeline updates
- **Security Alerts**: Send immediate security notifications for intrusion detection, authentication failures, and system breaches
- **Service Monitoring**: Monitor service uptime, performance metrics, and automatically notify teams of outages or issues
- **Task Automation**: Send notifications for completed automation tasks, scheduled job results, and workflow status updates
- **Team Communication**: Broadcast important updates, maintenance windows, and operational announcements to teams
- **Mobile App Integration**: Deliver real-time notifications to mobile devices for critical business and operational events
- **Custom Alerting**: Create custom notification workflows for business processes, customer events, and data pipeline status
- **Message Management**: Organize, prioritize, and manage notification history for audit trails and communication tracking

