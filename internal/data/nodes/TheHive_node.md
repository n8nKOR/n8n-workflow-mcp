# TheHive Node

## Overview

Consume TheHive API

## Credentials

- Name: theHiveApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: alert

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection |  | No |  |

#### Operation: merge

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Case ID | string |  | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the alert | Yes |  |
| Description | string | Description of the alert | Yes |  |
| Severity | options | Severity of the alert. Default=Medium. | Yes |  |
| Date | dateTime | Date and time when the alert was raised default=now | Yes |  |
| Tags | string | Case Tags | Yes |  |
| TLP | options | Traffict Light Protocol (TLP). Default=Amber. | Yes |  |
| Status | options | Status of the alert | Yes |  |
| Type | string | Type of the alert | Yes |  |
| Source | string | Source of the alert | Yes |  |
| SourceRef | string | Source reference of the alert | Yes |  |
| Follow | boolean | Whether the alert becomes active when updated default=true | Yes |  |
| Artifacts | fixedCollection | Type of the observable. Choose from the list, or specify an ID using an <a href= | No |  |
| Additional Fields | collection | Case template to use when a case is created from this alert | No |  |

#### Operation: executeResponder

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Responder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: promote

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Case template to use when a case is created from this alert | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Fields | collection | Type of the observable. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Options | collection | Whether to include similar cases | No |  |

### Resource: observable

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Data | string |  | Yes |  |
| Input Binary Field | string | The name of the input binary field that represent the attachment file | Yes |  |
| Message | string | Description of the observable in the context of the case | Yes |  |
| Start Date | dateTime | Date and time of the begin of the case default=now | Yes |  |
| TLP | options | Traffict Light Protocol (TLP). Default=Amber. | Yes |  |
| IOC | boolean | Whether the observable is an IOC (Indicator of compromise) | Yes |  |
| Sighted | boolean | Whether sighted previously | Yes |  |
| Status | options | Status of the observable. Default=Ok. | Yes |  |
| Options | collection |  | No |  |

#### Operation: executeAnalyzer

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Analyzer Names or IDs | multiOptions | Choose from the list, or specify IDs using an <a href= | Yes |  |

#### Operation: executeResponder

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Responder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Fields | collection | Description of the observable in the context of the case | No |  |

### Resource: case

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Specify the sorting attribut, + for asc, - for desc | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the case | Yes |  |
| Description | string | Description of the case | Yes |  |
| Severity | options | Severity of the alert. Default=Medium. | Yes |  |
| Start Date | dateTime | Date and time of the begin of the case default=now | Yes |  |
| Owner | string |  | Yes |  |
| Flag | boolean | Flag of the case default=false | Yes |  |
| TLP | options | Traffict Light Protocol (TLP). Default=Amber. | Yes |  |
| Tags | string |  | Yes |  |
| Options | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: executeResponder

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Responder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Task details | Yes |  |
| Status | options | Status of the task. Default=Waiting. | Yes |  |
| Flag | boolean | Whether to flag the task. Default=false. | Yes |  |
| Options | collection | Task details | No |  |

#### Operation: executeResponder

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Responder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Fields | collection | Task details | No |  |

### Resource: log

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message | string | Content of the Log | Yes |  |
| Start Date | dateTime | Date of the log submission default=now | Yes |  |
| Status | options | Status of the log (Ok or Deleted) default=Ok | Yes |  |
| Options | collection | The name of the input binary field which holds binary data | No |  |

#### Operation: executeResponder

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Responder Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

## UseCases

- [Security Incident Response] : Manage and track security incidents and investigations
- [Threat Intelligence] : Collect and analyze threat intelligence data for security teams
- [Case Management] : Create and manage security cases with structured workflows
- [Alert Investigation] : Investigate security alerts and convert them to cases when needed
- [Observable Analysis] : Analyze observables like IPs, domains, and file hashes for threats
- [Collaboration] : Enable security team collaboration on investigations and incident response
- [Task Management] : Assign and track security tasks within incident response workflows
- [Evidence Collection] : Collect and organize digital evidence for security investigations
- [Reporting] : Generate security incident reports and analysis documentation
- [Integration Automation] : Integrate with SIEM and security tools for automated workflows
- [Threat Hunting] : Support proactive threat hunting activities and investigations
- [Compliance] : Maintain security incident records for compliance and audit requirements
- [Notification Management] : Send notifications and alerts for security incidents and updates
- [Workflow Automation] : Automate security response workflows and playbooks
- [Forensic Analysis] : Support digital forensics and malware analysis activities

