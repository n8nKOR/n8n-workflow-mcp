# Document Loaders

## Overview

Document Loaders are specialized nodes in n8n's LangChain ecosystem that handle the ingestion and processing of various document formats and data sources. These nodes convert different file types, web content, and structured data into standardized LangChain Document objects that can be used by AI models, vector stores, and other LangChain components.

## Credentials

Document Loaders support various authentication methods depending on the source:

- **No Credentials**: Required for basic file processing operations
- **GitHub API**: Required for DocumentGithubLoader to access repositories

## Inputs

Document Loaders accept data connections for processing documents and text splitting from previous workflow nodes.

## Outputs

Document Loaders provide processed documents in LangChain Document format with metadata to subsequent workflow nodes.

## Properties

### Resource: Document Loading

#### Operation: Load JSON Data

| Field Name | Type | Description | Required |
|---|---|---|---|
| Data Type | Options | Type of data to load (JSON Data, Binary Data) | Yes |
| JSON Mode | Options | How to process JSON input data | Yes |
| JSON Data | String | JSON data to process when using expression mode | No |
| Metadata | Collection | Additional metadata fields to include | No |

#### Operation: Load Binary Data

| Field Name | Type | Description | Required |
|---|---|---|---|
| Data Type | Options | Type of data to load (JSON Data, Binary Data) | Yes |
| Binary Mode | Options | How to process binary input data | Yes |
| Data Format | Options | Binary file format (Auto-detect, PDF, CSV, DOCX, etc.) | Yes |
| Binary Property | String | Property name containing binary data | No |

#### Operation: Load GitHub Repository

| Field Name | Type | Description | Required |
|---|---|---|---|
| Repository | String | GitHub repository (owner/repo format) | Yes |
| Branch | String | Git branch to load from | No |
| Recursive | Boolean | Load files recursively from subdirectories | No |
| Ignore Paths | String | Comma-separated paths to ignore | No |
| Max File Size | Number | Maximum file size in MB | No |

#### Operation: Text Splitting Configuration

| Field Name | Type | Description | Required |
|---|---|---|---|
| Text Splitting Mode | Options | Mode for text splitting (Simple, Custom) | No |
| Chunk Size | Number | Size of text chunks in characters | No |
| Chunk Overlap | Number | Overlap between chunks in characters | No |
| Separator | String | Custom separator for text splitting | No |

## UseCases

- **Document Processing** : Convert PDF, DOCX, and other document formats into processable text
- **Code Repository Analysis** : Load and process entire GitHub repositories for AI analysis
- **Data Ingestion** : Import JSON and CSV data for LangChain processing workflows
- **Content Extraction** : Extract text content from various binary file formats
- **Batch Document Processing** : Process multiple documents simultaneously with consistent formatting
- **Metadata Enrichment** : Add custom metadata to documents for enhanced retrieval
- **Text Chunking** : Split large documents into manageable chunks for AI processing
- **Knowledge Base Creation** : Build comprehensive knowledge bases from multiple document sources
- **RAG System Preparation** : Prepare documents for Retrieval-Augmented Generation workflows
- **Automated Documentation Processing** : Process and index technical documentation automatically 