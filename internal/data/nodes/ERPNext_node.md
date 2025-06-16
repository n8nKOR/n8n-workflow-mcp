# ERPNext Node

## Overview

Consume ERPNext API

## Credentials

- Name: erpNextApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: document

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| DocType Name or ID | options | DocType whose documents to retrieve. Choose from the list, or specify an ID using an <a href= | No |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Comma-separated list of fields to return. Choose from the list, or specify IDs using an <a href= | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| DocType Name or ID | options | DocType you would like to create. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Properties | fixedCollection | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| DocType Name or ID | options | The type of document you would like to get. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Document Name | string | The name (ID) of document you would like to get | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| DocType Name or ID | options | The type of document you would like to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Document Name | string | The name (ID) of document you would like to get | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| DocType Name or ID | options | The type of document you would like to update. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Document Name | string | The name (ID) of document you would like to get | Yes |  |
| Properties | fixedCollection | Properties of request body | No |  |

## UseCases

- **ERP Data Integration**: Sync business data between ERPNext and external systems for unified business operations
- **Automated Invoice Processing**: Create, update, and manage invoices, purchase orders, and financial documents automatically
- **Customer Data Management**: Maintain customer records, contacts, and account information across multiple systems
- **Inventory Management**: Automate stock updates, item management, and warehouse operations
- **Sales Order Automation**: Process sales orders, quotations, and delivery notes through automated workflows
- **HR and Payroll Integration**: Manage employee records, attendance, and payroll data synchronization
- **Financial Reporting**: Extract financial data for automated reporting and analytics dashboards
- **Supply Chain Automation**: Coordinate purchase orders, supplier management, and procurement processes
- **Project Management**: Sync project data, tasks, and resource allocation between ERPNext and project management tools
- **Manufacturing Operations**: Automate work orders, bill of materials, and production planning workflows

