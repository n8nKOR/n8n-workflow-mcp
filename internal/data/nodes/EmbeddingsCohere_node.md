# EmbeddingsCohere

## Overview

The EmbeddingsCohere node provides text embedding capabilities using Cohere's industry-leading AI models within n8n's LangChain ecosystem. This node converts text into high-dimensional vectors using Cohere's specialized embedding models, supporting both English and multilingual content.

## Credentials

The EmbeddingsCohere node requires Cohere API credentials for authentication:

- **Cohere API**: Authentication for Cohere embedding services

## Inputs

The EmbeddingsCohere node does not accept direct inputs as it serves as an embedding provider.

## Outputs

- **AI Embedding**: LangChain embedding provider for use with other AI components

## Properties

### Resource: Cohere Embeddings

#### Operation: Generate Embeddings

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Resource Locator | Cohere embedding model to use | Yes |
| Input Type | Options | Type of input text (search_document, search_query, classification, clustering) | No |
| Truncate | Options | How to handle text that exceeds model limits | No |
| Request Options | Collection | Additional request configuration | No |

## UseCases

- **Semantic Search**: Build intelligent search systems with semantic understanding
- **Content Recommendation**: Create personalized content recommendation engines
- **Document Classification**: Automatically categorize and organize documents
- **Multilingual Applications**: Process content in 100+ languages with consistent quality
- **Customer Support**: Power intelligent customer service with semantic matching
- **Knowledge Management**: Build corporate knowledge bases with semantic search 