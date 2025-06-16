# Stackby Node

## Overview

Read, write, and delete data in Stackby

## Credentials

- Name: stackbyApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Records

#### Operation: Append

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform. Options: "Append", "Delete", "List", "Read" | Yes |
| Stack ID | String | The ID of the stack to access | Yes |
| Table | String | Name of the table to work with | Yes |
| Columns | String | Comma-separated list of properties to use as columns for new rows (Append only) | Yes |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Stack ID | String | The ID of the stack to access | Yes |
| Table | String | Name of the table to work with | Yes |
| ID | String | ID of the record to delete | Yes |

#### Operation: List

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Stack ID | String | The ID of the stack to access | Yes |
| Table | String | Name of the table to work with | Yes |
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return (when Return All is false) | No |
| View | String | Name or ID of a view to filter records | No |

#### Operation: Read

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Stack ID | String | The ID of the stack to access | Yes |
| Table | String | Name of the table to work with | Yes |
| ID | String | ID of the record to return | Yes |

## UseCases

- [Project Management] : Manage projects with collaborative spreadsheets and task tracking
- [Database Management] : Create and manage relational databases with spreadsheet interface
- [Inventory Tracking] : Track inventory levels and manage stock across multiple locations
- [Customer Relationship Management] : Manage customer data and interactions in organized tables
- [Content Planning] : Plan and organize content calendars for marketing and publishing
- [Team Collaboration] : Enable team collaboration on shared data and workflows
- [Data Collection] : Collect and organize data from various sources and forms
- [Reporting and Analytics] : Create reports and analyze data with built-in visualization tools
- [Workflow Automation] : Automate repetitive tasks and data processing workflows
- [Resource Management] : Track and manage company resources and asset allocation
- [Event Management] : Organize and manage events with attendee tracking and logistics
- [Sales Pipeline] : Track sales opportunities and manage customer acquisition process
- [HR Management] : Manage employee data, recruiting, and HR processes
- [Budget Planning] : Create and track budgets with collaborative planning tools
- [Task Management] : Organize tasks and assignments with team accountability features

