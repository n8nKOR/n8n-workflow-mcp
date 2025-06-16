# MySQL Node

## Overview

Get, add, update, and manage data in MySQL databases. MySQL is one of the world's most popular open-source relational database management systems, widely used for web applications, data warehousing, and enterprise applications. This node provides comprehensive database operations for integrating MySQL with n8n workflows.

## Credentials

- Name: mySql, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: database

#### Operation: deleteTable

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table | resource | The table to operate on | Yes |  |
| Delete Type | options | Type of delete operation (deleteRows, deleteTable) | Yes | deleteRows |
| Where | fixedCollection | Data used as WHERE clause to identify which rows to delete | No |  |

#### Operation: executeQuery

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Query | string | The SQL query to execute (SELECT, INSERT, UPDATE, DELETE, etc.) | Yes |  |
| Options | collection | Additional query options (parameters, formatting) | No |  |

#### Operation: insert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table | resource | The table to operate on | Yes |  |
| Columns | string | Columns to insert data into with their values | Yes |  |
| Options | collection | Additional insert options (ignore duplicates, on duplicate key) | No |  |

#### Operation: select

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table | resource | The table to operate on | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 50 |
| Options | collection | Additional select options (WHERE, ORDER BY, GROUP BY) | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table | resource | The table to operate on | Yes |  |
| Columns | string | Columns to update with new values | Yes |  |
| Where | fixedCollection | Data used as WHERE clause to identify which rows to update | No |  |
| Options | collection | Additional update options | No |  |

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Table | resource | The table to operate on | Yes |  |
| Columns | string | Columns to insert or update data (INSERT ... ON DUPLICATE KEY UPDATE) | Yes |  |
| Options | collection | Additional upsert options | No |  |

## UseCases

- **Web Application Data Management**: Store and manage user accounts, content, settings, and application data for web applications and websites
- **E-commerce Database Operations**: Handle product catalogs, inventory management, order processing, and customer data for online stores
- **Content Management Systems**: Manage articles, media files, categories, and user-generated content in CMS platforms
- **Customer Relationship Management**: Store customer information, interaction history, sales data, and support tickets in CRM systems
- **Analytics and Reporting**: Extract data for business intelligence, generate reports, and perform data analysis for decision-making
- **Inventory and Asset Management**: Track inventory levels, asset locations, maintenance schedules, and procurement data
- **Financial Data Processing**: Manage financial transactions, accounting records, billing information, and payment processing data
- **User Authentication and Authorization**: Handle user credentials, permissions, roles, and session management for secure applications
- **Log and Event Management**: Store application logs, system events, audit trails, and monitoring data for analysis
- **Data Migration and Synchronization**: Transfer data between systems, synchronize databases, and perform ETL operations for data integration

