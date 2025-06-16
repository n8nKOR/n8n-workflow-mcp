# Quick Base Node

## Overview

Integrate with the Quick Base RESTful API

## Credentials

- Name: quickbaseApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: field

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table ID | string | The table identifier | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to get back the custom permissions for the field(s) | No |  |

### Resource: file

#### Operation: download

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Input Binary Field | string |  | Yes |  |

### Resource: record

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table ID | string | The table identifier | Yes |  |
| Columns | string | Comma-separated list of the properties which should used as columns for the new rows | Yes |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |
| Options | collection | Specify an array of field IDs that will return data for any updates or added record. Record ID (FID 3) is always returned if any field ID is requested. Choose from the list, or specify IDs using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table ID | string | The table identifier | Yes |  |
| Where | string | The filter to delete records. To delete all records specify a filter that will include all records, for example {3.GT.0} where 3 is the ID of the Record ID field. | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table ID | string | The table identifier | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | An array of field IDs for the fields that should be returned in the response. If empty, the default columns on the table will be returned. Choose from the list, or specify IDs using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table ID | string | The table identifier | Yes |  |
| Columns | string | Comma-separated list of the properties which should used as columns for the new rows | Yes |  |
| Update Key | string | Update can use the key field on the table, or any other supported unique field | No |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |
| Options | collection | Specify an array of field IDs that will return data for any updates or added record. Record ID (FID 3) is always returned if any field ID is requested. Choose from the list, or specify IDs using an <a href= | No |  |

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table ID | string | The table identifier | Yes |  |
| Columns | string | Comma-separated list of the properties which should used as columns for the new rows | Yes |  |
| Update Key | string | Update can use the key field on the table, or any other supported unique field | No |  |
| Merge Field Name or ID | options | <p>You\ | No |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |
| Options | collection | Specify an array of field IDs that will return data for any updates or added record. Record ID (FID 3) is always returned if any field ID is requested. Choose from the list, or specify IDs using an <a href= | No |  |

### Resource: report

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table ID | string | The table identifier | Yes |  |
| Report ID | string | The identifier of the report, unique to the table | Yes |  |

#### Operation: run

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table ID | string | The table identifier | Yes |  |
| Report ID | string | The identifier of the report, unique to the table | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- [Database Management] : Create and manage custom databases for business applications
- [Workflow Automation] : Automate business processes and workflows with low-code solutions
- [Project Management] : Build custom project management applications tailored to specific needs
- [Inventory Tracking] : Create inventory management systems with real-time tracking capabilities
- [Customer Relationship Management] : Build custom CRM solutions for customer data management
- [Reporting and Analytics] : Generate custom reports and dashboards for business insights
- [Team Collaboration] : Create collaborative workspaces for team projects and data sharing
- [Asset Management] : Track and manage company assets, equipment, and resources
- [Document Management] : Organize and manage documents with custom metadata and workflows
- [Quality Management] : Build quality control systems for manufacturing and service processes
- [Human Resources] : Create HR applications for employee management and tracking
- [Financial Tracking] : Build custom financial applications for budgeting and expense tracking
- [Compliance Management] : Create compliance tracking systems for regulatory requirements
- [Task Management] : Build custom task management systems for teams and departments
- [Data Integration] : Integrate QuickBase with other business systems and applications

