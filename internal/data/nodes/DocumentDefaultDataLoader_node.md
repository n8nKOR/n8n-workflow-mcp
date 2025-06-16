# Default Data Loader

## Overview

The Default Data Loader node loads data from previous steps in the workflow for use with LangChain document processing. This node can process both JSON and binary data, supporting various file formats including PDF, DOCX, CSV, and plain text files.

## Credentials

This node does not require credentials as it processes data from previous workflow steps.

## Inputs

- **Text Splitter**: Optional connection for custom text splitting (when using custom mode)

## Outputs

- **Document**: Returns processed documents for use with other LangChain components

## Properties

### Resource: Document Processing

#### Operation: Default Data Loading

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type of Data | Options | Whether to process JSON or Binary data | Yes |
| Mode | Options | How to load the data (all input data or specific data) | Yes |
| Data Format | Options | Format of binary data (auto-detect, CSV, DOCX, PDF, etc.) | No |
| Data | Expression | Specific data to load when using expression mode | No |
| Binary Property | String | Name of binary property to load | No |
| Text Splitting Mode | Options | How to split text (automatic or custom) | No |

## UseCases

- **Document Processing**: Load documents from previous workflow steps for AI processing
- **Data Transformation**: Convert various file formats into document format for LangChain
- **Text Extraction**: Extract text content from binary files for analysis
- **Workflow Integration**: Bridge between file processing and AI document workflows
- **Multi-format Support**: Handle multiple document formats in a single workflow 