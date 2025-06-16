# MISP Node

## Overview

Consume the MISP API

## Credentials

- Name: mispApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: attribute

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event UUID | string | UUID of the event to attach the attribute to | Yes |  |
| Type | options |  | Yes |  |
| Value | string |  | Yes |  |
| Additional Fields | collection | Who will be able to see this event once published | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attribute ID | string | UUID or numeric ID of the attribute | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attribute ID | string | UUID or numeric ID of the attribute | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attribute ID | string | ID of the attribute to update | Yes |  |
| Update Fields | collection | Who will be able to see this event once published | No |  |

### Resource: event

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Information | string | Information on the event - max 65535 characters | Yes |  |
| Additional Fields | collection | Analysis maturity level of the event | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event ID | string | UUID or numeric ID of the event | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event ID | string | UUID or numeric ID of the event | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: publish

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event ID | string | UUID or numeric ID of the event | Yes |  |

#### Operation: unpublish

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event ID | string | UUID or numeric ID of the event | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event ID | string | UUID or numeric ID of the event | Yes |  |
| Update Fields | collection | Analysis maturity level of the event | No |  |

### Resource: eventTag

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event ID | string | UUID or numeric ID of the event | Yes |  |
| Tag Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event ID | string | UUID or numeric ID of the event | Yes |  |
| Tag Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

### Resource: feed

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Provider | string |  | Yes |  |
| URL | string |  | Yes |  |
| Additional Fields | collection | Who will be able to see this event once published | No |  |

#### Operation: disable

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Feed ID | string | UUID or numeric ID of the feed | Yes |  |

#### Operation: enable

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Feed ID | string | UUID or numeric ID of the feed | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Feed ID | string | UUID or numeric ID of the feed | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Feed ID | string | ID of the feed to update | Yes |  |
| Update Fields | collection | Who will be able to see this event once published | No |  |

### Resource: galaxy

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Galaxy ID | string | UUID or numeric ID of the galaxy | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Galaxy ID | string | UUID or numeric ID of the galaxy | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: noticelist

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Noticelist ID | string | Numeric ID of the noticelist | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: organisation

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organisation ID | string | UUID or numeric ID of the organisation | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organisation ID | string | UUID or numeric ID of the organisation | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organisation ID | string | ID of the organisation to update | Yes |  |
| Update Fields | collection |  | No |  |

### Resource: tag

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Additional Fields | collection | Hex color code for the tag | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag ID | string | Numeric ID of the attribute | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag ID | string | ID of the tag to update | Yes |  |
| Update Fields | collection | Hex color code for the tag | No |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string |  | Yes |  |
| Role ID | string | Role IDs are available in the MISP dashboard at /roles/index | Yes |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | Numeric ID of the user | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | Numeric ID of the user | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | ID of the user to update | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: warninglist

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Warninglist ID | string | Numeric ID of the warninglist | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- [Threat Intelligence Sharing] : Share and exchange threat intelligence with security communities
- [IOC Management] : Manage indicators of compromise for threat detection and prevention
- [Incident Response] : Support incident response with threat intelligence and attribution data
- [Threat Attribution] : Attribute threats to specific actors and campaign groups
- [Malware Analysis] : Share and analyze malware samples and behavioral indicators
- [Security Event Correlation] : Correlate security events with known threat intelligence
- [Threat Hunting] : Support proactive threat hunting with comprehensive threat data
- [SIEM Integration] : Feed threat intelligence into SIEM systems for enhanced detection
- [Vulnerability Intelligence] : Share vulnerability information and exploitation indicators
- [Campaign Tracking] : Track and analyze threat actor campaigns and tactics
- [Feed Management] : Manage and consume threat intelligence feeds from multiple sources
- [Community Collaboration] : Collaborate with security researchers and threat analysts
- [Attribution Analysis] : Analyze threat attribution and actor relationships
- [Threat Landscape Monitoring] : Monitor evolving threat landscape and emerging threats
- [Security Awareness] : Support security awareness programs with threat intelligence

