# DocumentJsonInputLoader

## Overview

Deprecated document loader for processing JSON data from workflow inputs

## Credentials

- No credentials required

## Inputs

- AI Text Splitter (optional)

## Outputs

- Documents

## Properties

### Resource: Document Processing

#### Operation: JSON Input Loading

| Field Name | Type | Description | Required |
|---|---|---|---|
| Pointers | String | JSONPath pointers to extract from JSON data (e.g. "/text" or "/text, /meta/title") | No |
| Metadata | Collection | Additional metadata to add to each document for filtering during retrieval | No |

## UseCases

- **JSON Data Processing** : Extract and process text content from structured JSON data in workflows
- **Document Conversion** : Convert JSON objects into LangChain Document format for AI processing
- **Selective Data Extraction** : Use JSONPath pointers to extract specific fields from complex JSON structures
- **Metadata Enhancement** : Add custom metadata to documents for improved retrieval and filtering
- **Legacy Workflow Support** : Maintain compatibility with existing workflows using deprecated JSON input processing
- **API Response Processing** : Extract content from API responses in JSON format
- **Database Query Results** : Process JSON results from database queries
- **Structured Data Extraction** : Extract specific fields from complex JSON structures
- **Content Aggregation** : Combine multiple JSON fields into document content
- **Data Transformation** : Convert structured JSON data to searchable documents 