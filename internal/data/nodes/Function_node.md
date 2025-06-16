# Function Node

## Overview

⚠️ **DEPRECATED**: Use the newer Code node instead. Legacy JavaScript execution node superseded by Code node with better security and features.

## Credentials

This node requires no authentication credentials as it executes JavaScript code locally within the n8n environment.

## Inputs

- **Main**: Input data from previous nodes

## Outputs

- **Main**: Returns modified data after JavaScript execution

## Properties

### Resource: Legacy JavaScript Executor

#### Operation: Execute JavaScript Code

| Field Name | Type | Description | Required |
|---|---|---|---|
| Notice | Notice | Information about newer 'Code' node availability | No |
| JavaScript Code | String | The JavaScript code to execute | Yes |

## UseCases

- **Legacy Data Transformation**: Maintain existing workflows that transform data formats while planning migration to Code node
- **Bulk Item Processing**: Process all workflow items at once for operations like sorting, grouping, or aggregating data
- **Data Filtering**: Remove items from the workflow based on complex conditions
- **Field Manipulation**: Add, modify, or remove fields from all items in a batch operation
- **Array Operations**: Perform operations that require access to all items simultaneously
- **Custom Validation**: Implement complex validation logic that needs to consider all items together 