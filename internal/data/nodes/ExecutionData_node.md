# Execution Data Node

## Overview

Add execution data for search

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Execution Management

#### Operation: Save Highlight Data

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Operation to perform (save) | Yes |
| Data to Save | Collection | Key-value pairs to save | No |
## UseCases

- **Execution Tracking**: Save key data points for execution monitoring
- **Search Enhancement**: Enable execution filtering by custom data (Pro/Enterprise)
- **Debugging Aid**: Store debug information for troubleshooting
- **Business Metrics**: Track important business data across executions
- **Audit Trails**: Maintain execution context for compliance and review

### Usage Notes
- Saved data is displayed on execution details for easy reference
- Data is saved per execution and can be used for searching (Pro/Enterprise)
- The node passes through all input data unchanged
- Multiple key-value pairs can be saved in a single node
- Data is saved to the execution's custom data store
- Useful for tracking business metrics, debugging, and execution monitoring

