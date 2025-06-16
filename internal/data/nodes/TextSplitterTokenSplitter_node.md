# TextSplitterTokenSplitter Node

## Overview

Split text into chunks by tokens

## Credentials

None

## Inputs

None (Text Splitter Provider)

## Outputs

- Text Splitter

## Properties

### Resource: Text Splitting

#### Operation: Token-Based Splitting

| Field Name | Type | Description | Required |
|---|---|---|---|
| Chunk Size | Number | Maximum number of tokens per chunk | Yes |
| Chunk Overlap | Number | Number of tokens to overlap between chunks | Yes |

## UseCases

- **Document Processing**: Split large documents into manageable chunks for AI processing
- **Text Analysis**: Prepare text for analysis by breaking it into token-based segments
- **Vector Database Preparation**: Create appropriately sized chunks for vector embeddings
- **Content Chunking**: Split content while respecting token limits of language models
- **Text Preprocessing**: Prepare text data for downstream AI operations