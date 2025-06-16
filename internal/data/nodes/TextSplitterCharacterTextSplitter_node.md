# TextSplitterCharacterTextSplitter

## Overview

Character-based text splitting for document processing

## Credentials

- No credentials required

## Inputs

- No direct inputs (serves as text splitter provider)

## Outputs

- AI Text Splitter

## Properties

### Resource: Text Processing

#### Operation: Character Text Splitting

| Field Name | Type | Description | Required |
|---|---|---|---|
| Separator | String | Character or string used to split text | No |
| Chunk Size | Number | Maximum size of each text chunk | Yes |
| Chunk Overlap | Number | Number of characters to overlap between chunks | No |

## UseCases

- **Structured Text Processing** : Split documents with clear structural separators like paragraphs or sections
- **Code Splitting** : Divide code files into logical sections using language-specific separators
- **Document Preparation** : Prepare documents for vector storage and retrieval systems
- **Content Chunking** : Create appropriately sized chunks for language model processing
- **Data Preprocessing** : Preprocess text data for machine learning and analysis workflows
- **Memory Optimization** : Reduce memory usage by processing text in smaller, manageable chunks
- **Paragraph-Based Splitting** : Split documents by paragraphs using double newline separators
- **Line-Based Processing** : Process structured data line by line with single newline separators
- **Custom Delimiter Splitting** : Split text using custom separators for specific document formats
- **Context Preservation** : Maintain document context across chunks with overlap settings 