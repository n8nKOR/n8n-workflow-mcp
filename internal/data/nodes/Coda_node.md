# Coda Node

## Overview

The Coda node allows you to interact with Coda documents through the Coda API. Coda is a collaborative document platform that combines documents, spreadsheets, and databases into powerful all-in-one docs. This node enables you to read, write, and manipulate data in Coda tables, create rows, push buttons, and manage document content programmatically.

## Credentials

- Name: codaApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: table

#### Operation: createRow

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Table Name or ID | options | The table to create the row in. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Options | collection | Whether the API will not attempt to parse the data in any way | No |  |

#### Operation: getRow

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Table Name or ID | options | The table to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Row ID | string | ID or name of the row. Names are discouraged because they | Yes |  |
| Options | collection | Whether to return the data exactly in the way it got received from the API | No |  |

#### Operation: getAllRows

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Table Name or ID | options | The table to get the rows from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Query used to filter returned rows, specified as &lt;column_id_or_name&gt;:&lt;value&gt;. If you\ | No |  |

#### Operation: deleteRow

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Table Name or ID | options | The table to delete the row in. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Row ID | string | Row IDs to delete | Yes |  |

#### Operation: pushButton

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Table Name or ID | options | The table to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Row ID | string | ID or name of the row. Names are discouraged because they | Yes |  |
| Column Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getColumn

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Table Name or ID | options | The table to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Column ID | string | The table to get the row from | Yes |  |

#### Operation: getAllColumns

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Table Name or ID | options | The table to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: formula

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Formula ID | string | The formula to get the row from | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: control

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Control ID | string | The control to get the row from | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: view

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| View ID | string | The view to get the row from | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: getAllViewRows

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| View Name or ID | options | The table to get the rows from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Query used to filter returned rows, specified as &lt;column_id_or_name&gt;:&lt;value&gt;. If you\ | No |  |

#### Operation: getAllViewColumns

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| View Name or ID | options | The table to get the rows from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: deleteViewRow

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| View Name or ID | options | The view to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Row Name or ID | options | The view to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: pushViewButton

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| View Name or ID | options | The view to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Row Name or ID | options | The view to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Column Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: updateViewRow

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Doc Name or ID | options | ID of the doc. Choose from the list, or specify an ID using an <a href= | Yes |  |
| View Name or ID | options | The view to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Row Name or ID | options | The view to get the row from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Key Name | string | The view to get the row from | Yes |  |
| Options | collection | Whether the API will not attempt to parse the data in any way | No |  |

## UseCases

- **Data Management** : Synchronize data between n8n workflows and Coda documents
- **Content Automation** : Automatically create and update rows in Coda tables
- **Interactive Workflows** : Trigger actions by pushing buttons in Coda documents
- **Document Analysis** : Extract and analyze data from Coda tables and views
- **Report Generation** : Populate Coda documents with data from external systems
- **Project Management** : Track tasks, projects, and team activities in Coda
- **CRM Integration** : Manage customer data and interactions through Coda tables
- **Inventory Management** : Track products, stock levels, and purchase orders
- **Content Publishing** : Automate content creation and publishing workflows
- **Database Operations** : Perform CRUD operations on Coda documents as databases

