# Item List Output Parser

## Overview

The Item List Output Parser node returns AI-generated results as separate items in the workflow. This parser takes text output from language models and splits it into individual items based on a specified separator, making it easier to process lists and structured data in n8n workflows.

## Credentials

This node does not require credentials as it processes output from connected AI components.

## Inputs

This node does not accept input connections as it provides output parsing functionality to other components.

## Outputs

- **Output Parser**: Returns an output parser component that can be connected to chains and agents

## Properties

### Resource: Output Processing

#### Operation: Item List Parsing

| Field Name | Type | Description | Required |
|---|---|---|---|
| Number Of Items | Number | Maximum number of items to return (-1 for unlimited) | No |
| Separator | String | Separator used to split results into items (default: newline) | No |

## UseCases

- **List Processing** : Convert AI-generated lists into separate workflow items
- **Data Structuring** : Structure unstructured text output into processable items
- **Batch Processing** : Process each list item individually in subsequent workflow steps
- **Content Parsing** : Parse structured content from AI responses into manageable pieces
- **Workflow Integration** : Bridge AI output with n8n's item-based processing model