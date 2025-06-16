# MultiQuery Retriever

## Overview

The MultiQuery Retriever node automates prompt tuning by generating diverse queries and expanding the document pool for enhanced retrieval. This node uses a language model to generate multiple variations of a query to improve retrieval results by capturing different perspectives and phrasings.

## Credentials

This node does not require credentials as it uses connected language models and retrievers.

## Inputs

- **Model**: Language model connection for generating query variations
- **Retriever**: Base retriever connection for document retrieval

## Outputs

- **Retriever**: Enhanced retriever component that can be connected to chains and agents

## Properties

### Resource: Document Retrieval

#### Operation: Multi-Query Enhancement

| Field Name | Type | Description | Required |
|---|---|---|---|
| Query Count | Number | Number of different versions of the given question to generate (default: 3) | No |

## UseCases

- **Enhanced Search Results** : Generate multiple query variations to capture more relevant documents
- **Query Optimization** : Automatically improve search queries without manual tuning
- **Comprehensive Retrieval** : Expand document pool by using different query perspectives
- **Semantic Search Enhancement** : Improve retrieval accuracy through query diversification
- **RAG System Improvement** : Enhance retrieval-augmented generation with better document selection