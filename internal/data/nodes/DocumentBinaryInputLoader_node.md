# DocumentBinaryInputLoader

## Overview

Deprecated document loader for processing binary file data from workflow inputs

## Credentials

- Name: File System Access, Required: No (inherited from n8n environment)

## Inputs

- AI Text Splitter

## Outputs

- Main

## Properties

### Resource: Document Loader

#### Operation: Load Binary Document

| Field Name | Type | Description | Required |
|---|---|---|---|
| Binary Data Key | String | Key name for binary data in input | Yes |
| Loader | Options | Document loader type to use | Yes |
| Options | Collection | Additional configuration options | No |

## UseCases

- **File Upload Processing** : Process files uploaded through n8n workflows
- **Email Attachment Analysis** : Analyze attachments from email workflows
- **Document Conversion** : Convert various file formats to searchable text
- **Content Extraction** : Extract text content from binary files for analysis
- **Legacy Workflow Support** : Maintain existing workflows using this deprecated node
- **Data Migration** : Process files during data migration workflows
- **Batch File Processing** : Handle multiple files from file system or cloud storage
- **Report Processing** : Extract content from business reports and documents
- **Research Data Processing** : Process academic papers and research documents
- **Compliance Document Review** : Extract content from compliance documents