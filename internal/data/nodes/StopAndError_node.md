# Stop and Error

## Overview

The Stop and Error node allows you to intentionally throw an error in your workflow to halt execution. This node is useful for error handling, validation, and workflow control scenarios where you need to stop processing and provide specific error information.

## Credentials

This node does not require credentials.

## Inputs

- **Main**: Accepts input data to trigger the error

## Outputs

This node has no outputs as it stops workflow execution by throwing an error.

## Properties

### Resource: Error Control

#### Operation: Workflow Error Handling

| Field Name | Type | Description | Required |
|---|---|---|---|
| Error Type | Options | Type of error to throw (Error Message or Error Object) | Yes |
| Error Message | String | Simple error message to display | Yes (when Error Type is Error Message) |
| Error Object | JSON | Complex error object with additional properties | Yes (when Error Type is Error Object) |

## UseCases

- **Validation Errors** : Stop workflow when data validation fails
- **Conditional Error Handling** : Throw errors based on specific conditions
- **Workflow Testing** : Test error handling paths in your workflows
- **Data Quality Control** : Stop processing when data quality issues are detected
- **Business Logic Enforcement** : Enforce business rules by throwing errors when violated
- **Debug Points** : Stop execution at specific points for debugging purposes
- **Safety Checks** : Stop potentially dangerous operations before they execute
- **Resource Protection** : Stop when resources are unavailable or exhausted
- **Emergency Stops** : Halt critical processes in emergency situations
- **Manual Intervention Required** : Stop workflows that require manual intervention

