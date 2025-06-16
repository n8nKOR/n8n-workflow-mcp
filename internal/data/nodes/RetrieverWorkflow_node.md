# RetrieverWorkflow

## Overview

Custom retrieval using n8n workflows as retriever logic

## Credentials

- No specific credentials required (depends on workflow content)

## Inputs

- No direct inputs (serves as retriever provider)

## Outputs

- AI Retriever

## Properties

### Resource: Workflow Retrieval

#### Operation: Execute Workflow

| Field Name | Type | Description | Required |
|---|---|---|---|
| Source | Options | Where to get the workflow (database/parameter) | Yes |
| Workflow ID | String | ID of workflow to execute (v1.0) | Yes |
| Workflow | Workflow Selector | Workflow selector (v1.1+) | Yes |
| Workflow JSON | JSON | JSON definition of workflow (parameter mode) | Yes |
| Workflow Values | Collection | Values to pass to the workflow | No |

## UseCases

- **Custom Retrieval Logic** : Implement complex, business-specific retrieval algorithms
- **API Integration** : Integrate external APIs and services into retrieval workflows
- **Data Transformation** : Apply custom data processing and transformation during retrieval
- **Multi-Source Retrieval** : Combine multiple data sources in a single retrieval operation
- **Business Rule Implementation** : Apply business rules and logic during document retrieval
- **Complex Query Processing** : Handle sophisticated query processing and routing
- **External Service Integration** : Connect to external databases and services for retrieval
- **Dynamic Retrieval Logic** : Implement retrieval logic that adapts based on query parameters
- **Custom Scoring** : Apply custom relevance scoring and ranking algorithms
- **Hybrid Retrieval Systems** : Combine multiple retrieval strategies in one workflow 