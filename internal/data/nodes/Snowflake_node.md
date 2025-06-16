# Snowflake Node

## Overview

Get, add and update data in Snowflake

## Credentials

- Name: snowflake, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Database Operations

#### Operation: Execute Query

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform. Options: "Execute Query", "Insert", "Update" | Yes |
| Query | String | The SQL query to execute | Yes |

#### Operation: Insert

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Table | String | Name of the table in which to insert data to | Yes |
| Columns | String | Comma-separated list of properties to use as columns for new rows | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Table | String | Name of the table in which to update data in | Yes |
| Update Key | String | Name of the property which decides which rows should be updated (normally "id") | Yes |
| Columns | String | Comma-separated list of properties to use as columns for rows to update | No |

## UseCases

- **Data Warehouse Operations** : Execute complex analytical queries on large datasets stored in Snowflake
- **ETL Pipeline Integration** : Extract, transform, and load data as part of automated data pipelines
- **Real-time Analytics** : Query real-time data for business intelligence and reporting dashboards
- **Data Migration** : Migrate data between different systems using Snowflake as an intermediary
- **Automated Reporting** : Generate automated reports by querying Snowflake data warehouses
- **Data Quality Monitoring** : Monitor data quality by running validation queries on datasets
- **Business Intelligence Integration** : Connect BI tools with Snowflake for advanced analytics
- **Machine Learning Data Preparation** : Prepare and transform data for machine learning workflows
- **Compliance Reporting** : Generate compliance reports by querying regulated data in Snowflake
- **Performance Monitoring** : Monitor database performance and query execution metrics

