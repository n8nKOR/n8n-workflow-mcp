# CrateDB Node

## Overview

Add and update data in CrateDB

## Credentials

- Name: crateDb, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Database

#### Operation: Execute Query

| Field Name | Type | Description | Required |
|---|---|---|---|
| operation | options | The operation to perform on the CrateDB database | Yes |
| query | string | The SQL query to execute (supports n8n expressions and $1, $2 parameters) | Yes |
| queryParams | string | Comma-separated list of properties for query parameters | No |
| mode | options | How queries should be sent (independently, multiple) | No |

#### Operation: Insert

| Field Name | Type | Description | Required |
|---|---|---|---|
| operation | options | The operation to perform on the CrateDB database | Yes |
| schema | string | Name of the schema the table belongs to | Yes |
| table | string | Name of the table to insert data to | Yes |
| columns | string | Comma-separated list of columns for new rows | No |
| returnFields | string | Comma-separated list of fields to return | No |
| mode | options | How queries should be sent (independently, multiple) | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| operation | options | The operation to perform on the CrateDB database | Yes |
| schema | string | Name of the schema the table belongs to | Yes |
| table | string | Name of the table to update data in | Yes |
| updateKey | string | Properties that decide which rows to update (normally "id") | Yes |
| columns | string | Comma-separated list of columns for rows to update | No |
| returnFields | string | Comma-separated list of fields to return | No |
| mode | options | How queries should be sent (independently, multiple) | No |

## UseCases

- [IoT Data Storage]: Insert and manage time-series data from IoT devices and sensors into CrateDB for real-time analytics and monitoring
- [Log Management]: Store and query application logs, system metrics, and operational data with CrateDB's distributed SQL capabilities
- [Real-time Analytics]: Execute complex SQL queries on large datasets for business intelligence and real-time reporting dashboards

