# Manual Trigger Node

## Overview

Entry point for manually triggered workflow execution. Activates when the 'Execute workflow' button is clicked in n8n interface.

## Credentials

None

## Inputs

None

## Outputs

- Main

## Properties

### Resource: Manual Trigger

#### Operation: Trigger Workflow

| Field Name | Type | Description | Required |
|---|---|---|---|
| Notice | notice | Informational text about manual trigger usage | No |

### Trigger Behavior
This node serves as the entry point for workflow execution when manually triggered:

- **Trigger Type**: Manual execution via 'Execute workflow' button
- **Execution Context**: Starts workflow when test button is clicked on canvas
- **Data Output**: Emits empty JSON array `[{}]` to initialize workflow data
- **Single Instance**: Maximum of 1 Manual Trigger node per workflow

### User Interface Elements
- **Notice Display**: Shows informational message about trigger usage
- **Alternative Triggers**: Provides links to explore other trigger types (schedules, webhooks)
- **Node Creator Integration**: Direct access to node creation panel

### Configuration
- **Default Name**: "When clicking 'Execute workflow'"
- **Default Color**: #909298 (gray)
- **Icon**: Mouse pointer (fa:mouse-pointer)
- **Group**: Trigger nodes

### Workflow Integration
- **Execution Starting Point**: Primary entry point for manual workflow testing
- **Development Tool**: Essential for workflow development and testing
- **No Data Requirements**: Works without external data inputs
- **Immediate Response**: Triggers instantly when activated

## UseCases

- **Web Scraping Testing**: Manually execute Paul Graham essay scraping and AI summarization workflows to test and verify results.
- **OpenAI Model Testing**: Manually execute workflows integrating various OpenAI models (ChatGPT, DALL-E 2, Whisper) to validate each model's behavior.
- **Workflow Development and Debugging**: Execute complex automation workflows step-by-step during development to verify each node's operation and debug issues.
- **One-time Data Processing**: Manually execute workflows needed for non-regular data processing tasks or special circumstances.

