# SIGNL4 Node

## Overview

Consume SIGNL4 API

## Credentials

- Name: signl4Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: alert

#### Operation: send

| Field Name | Type | Description | Required |
|---|---|---|---|
| Message | string | A more detailed description for the alert | No |
| Additional Fields | collection | Additional options for the alert | No |

#### Send Additional Fields Options:

| Field Name | Type | Description | Required |
|---|---|---|---|
| Alerting Scenario | options | Single ACK or Multi ACK - whether only one person or multiple people need to confirm | No |
| Attachments | fixedCollection | Binary attachments to include with the alert | No |
| External ID | string | Unique ID from a 3rd party system for correlation/synchronization | No |
| Filtering | boolean | Whether to apply event filtering for this event | No |
| Location | fixedCollection | Transmit location information (latitude, longitude) to display a map | No |
| Service | string | Assigns the alert to the service/system category with the specified name | No |
| Title | string | The title or subject of this alert | No |

#### Operation: resolve

| Field Name | Type | Description | Required |
|---|---|---|---|
| External ID | string | The same External ID as used in the original alert to resolve/close it | No |

## UseCases

- [Incident Management] : Manage and escalate critical incidents with automated alerting
- [On-Call Management] : Coordinate on-call schedules and escalation procedures
- [Mobile Alerting] : Send critical alerts to mobile devices for immediate response
- [Team Collaboration] : Enable team collaboration during incident response
- [Emergency Response] : Coordinate emergency response procedures and communication
- [System Monitoring] : Alert teams about system outages and performance issues
- [IoT Device Alerts] : Receive alerts from IoT devices and sensor networks
- [Field Service Management] : Coordinate field service teams and technician dispatching
- [Safety Monitoring] : Monitor safety systems and alert response teams
- [Maintenance Scheduling] : Schedule and coordinate maintenance activities
- [Quality Assurance] : Alert quality teams about production issues and defects
- [Security Incidents] : Manage security incident response and team coordination
- [Facility Management] : Monitor building systems and coordinate facility maintenance
- [Manufacturing Alerts] : Alert production teams about equipment failures and issues
- [Healthcare Coordination] : Coordinate healthcare teams and emergency medical response

