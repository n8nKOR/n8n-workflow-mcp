# Function Item Node

## Overview

⚠️ **DEPRECATED**: Use the newer Code node instead. The Function Item node is a legacy n8n core node that executes custom JavaScript code once per input item. This node enables per-item data transformation, enrichment, and custom business logic implementation within workflows.

## Credentials

The Function Item node does not require any credentials as it operates within n8n's internal JavaScript sandbox.

## Inputs

- **Main**: Accepts data items for individual processing

## Outputs

- **Main**: Returns processed items with transformations applied

## Properties

### Resource: JavaScript Execution

#### Operation: Execute JavaScript Code

| Field Name | Type | Description | Required |
|---|---|---|---|
| JavaScript Code | String | The JavaScript code to execute for each item | Yes |

## UseCases

- **Legacy Data Transformation**: Maintain existing workflows that transform data formats while planning migration to Code node
- **Bulk Item Processing**: Process all workflow items at once for operations like sorting, grouping, or aggregating data
- **Data Filtering**: Remove items from the workflow based on complex conditions
- **Field Manipulation**: Add, modify, or remove fields from all items in a batch operation
- **Array Operations**: Perform operations that require access to all items simultaneously
- **Custom Validation**: Implement complex validation logic that needs to consider all items together 