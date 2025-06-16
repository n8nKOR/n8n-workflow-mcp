# Embeddings OpenAI

## Overview

The Embeddings OpenAI node provides text embedding capabilities using OpenAI's embedding models. This node converts text into high-dimensional vector representations that can be used for semantic search, similarity comparisons, and other AI applications within the LangChain ecosystem.

## Credentials

- **Credential Name**: openAiApi
- **Required**: Yes
- **Configuration**: OpenAI API credentials with appropriate permissions

## Inputs

This node does not accept input connections as it provides embedding functionality to other components.

## Outputs

- **Embeddings**: Returns an embeddings component that can be connected to vector stores and other AI components

## Properties

### Resource: Text Embeddings

#### Operation: OpenAI Text Embeddings

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Options | The model which will generate the embeddings | Yes |
| Dimensions | Options | The number of dimensions the resulting output embeddings should have | No |
| Base URL | String | Override the default base URL for the API | No |
| Batch Size | Number | Maximum number of documents to send in each request | No |
| Strip New Lines | Boolean | Whether to strip new lines from the input text | No |
| Timeout | Number | Maximum amount of time a request is allowed to take in seconds | No |

## UseCases

- **Semantic Search**: Create semantic search capabilities using text embeddings
- **Document Similarity**: Compare document similarity using vector representations
- **Content Recommendation**: Build recommendation systems based on content similarity
- **Text Classification**: Use embeddings for text classification tasks
- **Vector Database Storage**: Store text embeddings in vector databases for retrieval 