# QuestDB Node

## Overview

Get, add and update data in QuestDB. QuestDB is a high-performance time-series database designed for real-time analytics on time-series and event data. This node allows you to execute SQL queries and insert data into QuestDB tables, making it ideal for handling large volumes of time-series data with low latency.

## Credentials

- Name: questDb, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Database Operations

#### Operation: Execute Query and Insert

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Operation to perform (Execute Query/Insert) | Yes |
| Query | String | SQL query to execute | Conditional |
| Table | String | Name of the table for insert operations | Conditional |
| Columns | String | Comma-separated list of columns for insert | Conditional |
| Return Fields | String | Fields to return from insert operation | No |
| Mode | Options | Query execution mode (Independently/Transaction) | No |
| Query Parameters | String | Parameters for parameterized queries | No |

## UseCases

- **Time Series Data**: Store and query time series data efficiently
- **Real-time Analytics**: Perform real-time analytics on streaming data
- **IoT Data Processing**: Process and store IoT sensor data
- **Financial Data**: Handle high-frequency financial data and trading information
- **Log Analysis**: Store and analyze application logs and metrics


