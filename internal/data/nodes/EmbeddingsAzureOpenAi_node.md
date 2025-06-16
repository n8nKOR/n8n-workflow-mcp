# EmbeddingsAzureOpenAi

## Overview

The Embeddings Azure OpenAI node provides integration with Azure OpenAI's embedding models, enabling the generation of vector embeddings from text data. This node is essential for building semantic search, recommendation systems, and other AI applications that require text similarity and vector operations.

## Credentials

- **Credential Name**: azureOpenAiApi
- **Required**: Yes
- **Configuration**: 
  - API Key: Your Azure OpenAI API key
  - Resource Name: Azure OpenAI resource name
  - API Version: Azure OpenAI API version
  - Endpoint: Custom endpoint URL (optional)

## Inputs

This node does not accept input connections as it provides embedding functionality to other components.

## Outputs

- **Embeddings**: Returns an embeddings component that can be connected to vector stores and other AI components

## Properties

### Resource: Text Embeddings

#### Operation: Embedding Generation

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model (Deployment) Name | String | The name of the model deployment to use for embeddings | Yes |
| Batch Size | Number | Maximum number of documents to send in each request (max: 2048) | No |
| Strip New Lines | Boolean | Whether to strip new lines from the input text | No |
| Timeout | Number | Maximum request timeout in seconds (-1 for no timeout) | No |
| Dimensions | Options | Number of dimensions for output embeddings (256, 512, 1024, 1536, 3072) | No |

## UseCases

- **Semantic Search**: Generate embeddings for text documents to enable semantic search capabilities
- **Recommendation Systems**: Create vector representations for content-based recommendations
- **Text Similarity**: Compare text similarity using vector embeddings
- **Knowledge Base**: Build vector databases for AI-powered knowledge retrieval
- **Content Classification**: Use embeddings for automated content categorization 