# EmbeddingsAwsBedrock

## Overview

The EmbeddingsAwsBedrock node provides text embedding capabilities using AWS Bedrock foundation models within n8n's LangChain ecosystem. This node enables conversion of text to high-dimensional vectors using AWS's managed foundation models, supporting various embedding models available through the Bedrock service.

## Credentials

The EmbeddingsAwsBedrock node requires AWS credentials with appropriate permissions for Bedrock service access:

- **AWS Credentials**: Standard AWS authentication with Bedrock permissions

## Inputs

The EmbeddingsAwsBedrock node does not accept direct inputs as it serves as an embedding provider.

## Outputs

- **AI Embedding**: LangChain embedding provider for use with other AI components

## Properties

### Resource: AWS Bedrock Embeddings

#### Operation: Generate Embeddings

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Resource Locator | AWS Bedrock embedding model to use | Yes |
| Region | String | AWS region for Bedrock service | No |
| Max Retries | Number | Maximum number of retry attempts | No |
| Request Options | Collection | Additional request configuration | No |

## UseCases

- **Enterprise Document Search**: Build enterprise search systems with AWS-managed embeddings
- **Content Recommendation**: Create content recommendation engines using Bedrock embeddings
- **Semantic Search**: Implement semantic search across large document collections
- **Knowledge Base Systems**: Power AI knowledge bases with enterprise-grade embeddings
- **Customer Support**: Enable intelligent customer support with document understanding
- **Compliance Documentation**: Process regulatory and compliance documents with secure embeddings 