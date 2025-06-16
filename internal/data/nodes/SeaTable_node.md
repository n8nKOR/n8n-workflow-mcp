# SeaTable Node

## Overview

Consume the SeaTable API

⚠️ **Version Information**: This node has multiple versions (V1, V2). The current default version is V2, which provides enhanced functionality and improved field handling. Some features may differ between versions.

## Credentials

- Name: seaTableApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Row

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Data to Send | Options | Whether to insert the input data this node receives in the new row | Yes |
| Apply Column Default Values | Boolean | Whether to use the column default values to populate new rows during creation | No |
| Inputs to Ignore | String | List of input properties to avoid sending, separated by commas | No |
| Columns to Send | Collection | Add destination column with its value | No |
| Save to "Big Data" Backend | Boolean | Whether write to Big Data backend or not | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Row ID | String | ID of the row to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Row ID | String | ID of the row to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| View Name | String | Name of the view to get rows from | No |
| Convert Link ID | Boolean | Whether to convert link column IDs to readable names | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Row ID | String | ID of the row to update | Yes |
| Data to Send | Options | Whether to insert the input data this node receives in the row | Yes |
| Columns to Send | Collection | Add destination column with its value | No |

### Resource: Base

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | The ID of the SeaTable base | Yes |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

### Resource: Link

#### Operation: Add

| Field Name | Type | Description | Required |
|---|---|---|---|
| Link Column | String | Name of the link column | Yes |
| Row ID | String | ID of the row to link from | Yes |
| Other Row ID | String | ID of the row to link to | Yes |

#### Operation: List

| Field Name | Type | Description | Required |
|---|---|---|---|
| Link Column | String | Name of the link column | Yes |
| Row ID | String | ID of the row to get links from | Yes |

#### Operation: Remove

| Field Name | Type | Description | Required |
|---|---|---|---|
| Link Column | String | Name of the link column | Yes |
| Row ID | String | ID of the row to unlink from | Yes |
| Other Row ID | String | ID of the row to unlink | Yes |

### Resource: Asset

#### Operation: Upload

| Field Name | Type | Description | Required |
|---|---|---|---|
| File Path | String | Path to the file to upload | Yes |
| File Name | String | Name for the uploaded file | No |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Asset URL | String | URL of the asset to download | Yes |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Asset URL | String | URL of the asset to delete | Yes |

## UseCases

- **Database Management** : Transfer data between SeaTable and other systems, create backups, and manage data import/export
- **Workflow Automation** : Automatically process form submissions, manage tasks, and track inventory
- **CRM Integration** : Sync contacts and leads between SeaTable and CRM systems
- **Project Management** : Link tasks and projects across different platforms
- **Reporting and Analytics** : Export data for custom dashboards and business intelligence
- **Team Collaboration** : Share and update team project information and resources



- **HTTP Request**: For custom SeaTable API calls
- **JSON**: For processing SeaTable API responses
- **Code**: For custom data transformation logic
- **Schedule Trigger**: For automated SeaTable operations

