# TimescaleDB Node

## Overview

Add and update data in TimescaleDB

## Credentials

- Name: timescaleDb, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Database

#### Operation: Execute Query

| Field Name | Type | Description | Required |
|---|---|---|---|
| operation | options | The operation to perform on the TimescaleDB database | Yes |
| query | string | SQL query to execute (supports n8n expressions and $1, $2 parameters) | Yes |
| queryParams | string | Comma-separated list of properties for query parameters | No |
| mode | options | How queries should be sent (independently, multiple, transaction) | No |

#### Operation: Insert

| Field Name | Type | Description | Required |
|---|---|---|---|
| operation | options | The operation to perform on the TimescaleDB database | Yes |
| schema | string | Name of the schema the table belongs to | Yes |
| table | string | Name of the table to operate on | Yes |
| columns | string | Comma-separated list of properties to use as columns | No |
| returnFields | string | Comma-separated list of fields to return | No |
| mode | options | How queries should be sent (independently, multiple, transaction) | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| operation | options | The operation to perform on the TimescaleDB database | Yes |
| schema | string | Name of the schema the table belongs to | Yes |
| table | string | Name of the table to operate on | Yes |
| updateKey | string | Property that decides which rows to update (normally "id") | Yes |
| columns | string | Comma-separated list of properties to use as columns | No |
| returnFields | string | Comma-separated list of fields to return | No |
| mode | options | How queries should be sent (independently, multiple, transaction) | No |

## UseCases

- [Time-Series Data Management]: Store and query time-series data from IoT devices, sensors, and monitoring systems with TimescaleDB's optimized time-series capabilities
- [Analytics and Reporting]: Execute complex SQL queries on time-series data for business intelligence, performance monitoring, and historical trend analysis
- [Real-time Data Processing]: Insert streaming data and perform real-time analytics on time-series datasets with automatic partitioning and compression

