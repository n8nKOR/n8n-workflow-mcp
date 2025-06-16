# ServiceNow Node

## Overview

Consume ServiceNow API

## Credentials

- Name: serviceNowOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: attachment

#### Operation: upload

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table Record ID | string | Sys_id of the record in the table specified in Table Name that you want to attach the file to | Yes |  |
| Input Data Field Name | string | Name of the binary property that contains the data to upload | Yes |  |
| Options | collection | Name to give the attachment | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attachment ID | string | Sys_id value of the attachment to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attachment ID | string | Sys_id value of the attachment to delete | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: businessService

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

### Resource: configurationItems

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

### Resource: department

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

### Resource: dictionary

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

### Resource: incident

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Short Description | string | Short description of the incident | Yes |  |
| Additional Fields | collection | Which user is the incident assigned to. Requires the selection of an assignment group. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Incident ID | string | Unique identifier of the incident | Yes |  |
| Update Fields | collection | Which user is the incident assigned to. Requires the selection of an assignment group. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: tableRecord

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Data to Send | options | Use when node input names match destination field names | No |  |
| Inputs to Ignore | string | List of input properties to avoid sending, separated by commas. Leave empty to send all inputs. | No |  |
| Fields to Send | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Table Record ID | string | Unique identifier of the record | Yes |  |
| Data to Send | options | Use when node input names match destination field names | No |  |
| Inputs to Ignore | string | List of input properties to avoid sending, separated by commas. Leave empty to send all inputs. | No |  |
| Fields to Send | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Short Description | string | Short description of the user | Yes |  |
| Additional Fields | collection | Whether to activate the user | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Retrieve Identifier | options | Unique identifier of the user | Yes |  |
| Username | string | Unique identifier of the user | Yes |  |
| User ID | string | Unique identifier of the user | Yes |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | Unique identifier of the user | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | Unique identifier of the user | Yes |  |
| Update Fields | collection | Whether to activate the user | No |  |

### Resource: userGroup

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

### Resource: userRole

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to exclude Table API links for reference fields | No |  |

## UseCases

- [IT Service Management] : Manage IT service requests and incident resolution workflows
- [Incident Management] : Track and resolve IT incidents with automated escalation
- [Problem Management] : Identify and resolve root causes of recurring IT issues
- [Change Management] : Manage IT changes with approval workflows and risk assessment
- [Asset Management] : Track and manage IT assets and configuration items
- [Service Request Fulfillment] : Automate fulfillment of standard IT service requests
- [Knowledge Management] : Maintain knowledge base for IT support and self-service
- [Service Catalog] : Provide self-service portal for IT services and requests
- [Workflow Automation] : Automate IT processes and approval workflows
- [Integration Management] : Integrate ServiceNow with other enterprise systems
- [Reporting and Analytics] : Generate IT service reports and performance metrics
- [User Management] : Manage user accounts and access provisioning
- [Service Level Management] : Monitor and enforce service level agreements
- [Configuration Management] : Maintain configuration management database (CMDB)
- [Security Operations] : Manage security incidents and vulnerability response

