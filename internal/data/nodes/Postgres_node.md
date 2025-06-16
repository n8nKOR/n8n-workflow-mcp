# Postgres Node

## Overview

The Postgres node allows you to get, add, and update data in PostgreSQL databases. PostgreSQL is a powerful, open-source relational database system known for its reliability, robustness, and extensive feature set. This node enables you to execute SQL queries, manage tables, and perform CRUD operations on PostgreSQL databases through n8n workflows.

## Credentials

This node requires PostgreSQL database credentials:
- **Host**: PostgreSQL server hostname or IP address
- **Port**: Database port (default: 5432)
- **Database**: Database name to connect to
- **User**: Database username
- **Password**: Database password
- **SSL**: SSL connection settings (optional)

To configure PostgreSQL credentials:
1. Ensure your PostgreSQL server is accessible from n8n
2. Create a database user with appropriate permissions
3. Test the connection to verify credentials work
4. Configure SSL settings if required by your database

## Inputs

- **Main**: JSON data for database operations

## Outputs

- **Main**: Returns JSON data from database query results

## Properties

### Resource: Database

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Schema | String | Database schema name | Yes |
| Table | String | Table name to delete from | Yes |
| Delete Type | Options | Delete entire table or specific rows | Yes |
| Where Clause | String | SQL WHERE clause for row deletion | Conditional |

**Delete Options:**
- **Delete Table**: Remove the entire table structure
- **Delete Rows**: Delete specific rows based on conditions

#### Operation: Execute Query
| Field Name | Type | Description | Required |
|---|---|---|---|
| Query | String | SQL query to execute | Yes |
| Query Parameters | Collection | Parameterized query values | No |

**Query Types Supported:**
- **SELECT**: Retrieve data from tables
- **INSERT**: Add new records
- **UPDATE**: Modify existing records  
- **DELETE**: Remove records
- **CREATE**: Create tables and structures
- **ALTER**: Modify table structure
- **DROP**: Remove tables and structures

#### Operation: Insert
| Field Name | Type | Description | Required |
|---|---|---|---|
| Schema | String | Database schema name | Yes |
| Table | String | Table name to insert into | Yes |
| Data Mode | Options | How to specify data to insert | Yes |
| Columns | Collection | Column mappings and values | Yes |

**Data Mode Options:**
- **Auto-Map**: Automatically map input data to columns
- **Define Below**: Manually specify column mappings
- **Raw**: Use raw SQL VALUES clause

#### Operation: Insert or Update (Upsert)
| Field Name | Type | Description | Required |
|---|---|---|---|
| Schema | String | Database schema name | Yes |
| Table | String | Table name for upsert operation | Yes |
| Data Mode | Options | How to specify data for upsert | Yes |
| Conflict Columns | Array | Columns to check for conflicts | Yes |
| Update Policy | Options | How to handle conflicts | Yes |

**Upsert Strategies:**
- **Insert Only**: Insert if no conflict, skip if exists
- **Update Only**: Update if exists, skip if new
- **Insert or Update**: Insert new records, update existing ones

#### Operation: Select
| Field Name | Type | Description | Required |
|---|---|---|---|
| Schema | String | Database schema name | Yes |
| Table | String | Table name to select from | Yes |
| Select Mode | Options | How to select data | Yes |
| Columns | Collection | Columns to retrieve | Conditional |
| Where Clause | String | SQL WHERE clause for filtering | No |
| Sort | Collection | Sorting options | No |
| Limit | Number | Maximum number of records | No |
| Offset | Number | Number of records to skip | No |

**Select Mode Options:**
- **Select All**: Retrieve all columns
- **Select Columns**: Choose specific columns
- **Count**: Count matching records

#### Operation: Update  
| Field Name | Type | Description | Required |
|---|---|---|---|
| Schema | String | Database schema name | Yes |
| Table | String | Table name to update | Yes |
| Data Mode | Options | How to specify update data | Yes |
| Columns | Collection | Columns to update | Yes |
| Where Clause | String | SQL WHERE clause for filtering | Yes |

## UseCases

- **Database Operations** : Execute SQL queries for data retrieval, insertion, and modification
- **Data Analytics** : Perform complex analytics and reporting on PostgreSQL databases
- **Application Integration** : Connect applications with PostgreSQL for data storage and management
- **ETL Processes** : Extract, transform, and load data between PostgreSQL and other systems
- **User Management** : Store and manage user accounts and authentication data
- **Content Management** : Handle dynamic content storage and retrieval for applications
- **Inventory Management** : Track products, stock levels, and inventory operations
- **Financial Data Processing** : Manage transactions, accounting, and financial records
- **Audit Logging** : Track system events, changes, and compliance data
- **Real-time Data Processing** : Handle real-time event data and notifications
- [Content Duplication Prevention] : Check database for existing content codes before processing new social media posts to prevent duplicate content generation and posting
- [Trend Data Management] : Store and manage trending topic information with unique identifiers for automated content creation workflows
- [User Activity Tracking] : Log and track user interactions, form submissions, and workflow executions for analytics and monitoring purposes
- [Configuration Storage] : Store application settings, API keys, and workflow parameters in a centralized database for easy management and retrieval