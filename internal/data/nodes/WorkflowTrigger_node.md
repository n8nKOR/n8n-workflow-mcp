# Workflow Trigger Node

## Overview

Triggers based on various lifecycle events, like when a workflow is activated

## Credentials

None

## Inputs

None

## Outputs

- Main

## Properties

### Resource: Workflow

#### Operation: Trigger

| Field Name | Type | Description | Required |
|---|---|---|---|
| events | multiOptions | Specifies under which conditions an execution should happen | Yes |

**Available Events:**
- **Active Workflow Updated**: update - Triggers when this workflow is updated
- **Workflow Activated**: activate - Triggers when this workflow is activated

### Note
- **Type**: notice
- **Message**: This node is deprecated and would not be updated in the future. Please use 'n8n Trigger' node instead.

## UseCases

- [Workflow Lifecycle Management]: Trigger actions when workflows are activated or updated for monitoring and notification purposes
- [Development Automation]: Automatically execute setup or cleanup tasks when workflows are deployed or modified in development environments

