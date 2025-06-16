# Embeddings

## Overview

Embeddings nodes are specialized components in n8n's LangChain ecosystem that convert text into high-dimensional vector representations. These nodes provide the foundation for semantic search, similarity matching, and various AI applications that require numerical representations of text data.

## Credentials

Embeddings nodes typically require credentials specific to their provider:

- **Provider-Specific Credentials**: Each embedding provider requires its own authentication
- **API Keys**: Most providers require API keys for access
- **Service Configuration**: Provider-specific configuration parameters

## Inputs

Embeddings nodes do not accept input connections as they provide embedding functionality to other components.

## Outputs

- **Embeddings**: Returns an embeddings component that can be connected to vector stores and other AI components

## Properties

### Resource: Text Embeddings

#### Operation: Embedding Generation

| Field Name | Type | Description | Required |
|---|---|---|---|
| Provider | Options | Select the embedding provider (OpenAI, Azure OpenAI, Cohere, etc.) | Yes |
| Model | String | The specific model to use for generating embeddings | Yes |
| Batch Size | Number | Number of texts to process in each batch | No |
| Dimensions | Number | Number of dimensions for the output embeddings | No |
| Additional Options | Collection | Provider-specific configuration options | No |

## UseCases

- **Semantic Search** : Enable semantic search capabilities across text documents
- **Recommendation Systems** : Build content-based recommendation engines
- **Text Similarity** : Compare and match similar text content
- **Vector Databases** : Create vector representations for storage in vector databases
- **Content Classification** : Classify and categorize text content using embeddings

## Supported Models

### OpenAI
- text-embedding-3-small (1536 dimensions)
- text-embedding-3-large (3072 dimensions)
- text-embedding-ada-002 (1536 dimensions)

### Google
- textembedding-gecko (768 dimensions)
- textembedding-gecko-multilingual (768 dimensions)
- text-embedding-004 (768 dimensions)

### Cohere
- embed-english-v3.0
- embed-multilingual-v3.0
- embed-english-light-v3.0

### AWS Bedrock
- amazon.titan-embed-text-v1
- cohere.embed-english-v3
- cohere.embed-multilingual-v3

## UseCases

- **Vector Search**: Build similarity-based search systems for documents or text
- **Semantic Similarity Calculation**: Measure semantic similarity between texts
- **Document Clustering**: Group similar documents together
- **Recommendation Systems**: Implement content-based recommendation algorithms
- **Question-Answer Systems**: Search for documents related to questions
- **Duplicate Detection**: Identify similar or duplicate text
- **Classification Systems**: Extract features for text classification
- **Knowledge Graphs**: Model relationships between concepts
- **Multilingual Processing**: Compare semantic similarity across languages
- **Content Analysis**: Analyze large volumes of text data and discover patterns 