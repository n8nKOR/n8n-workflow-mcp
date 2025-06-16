# EmbeddingsOllama

## Overview

The EmbeddingsOllama node provides text embedding capabilities using locally hosted Ollama models within n8n's LangChain ecosystem. This node enables privacy-focused, offline embedding generation by connecting to a local Ollama server. It offers complete data sovereignty, cost-effective processing, and the flexibility to use open-source embedding models.

## Credentials

The EmbeddingsOllama node requires Ollama server connection credentials:

- **Ollama**: Connection configuration for local Ollama server

## Inputs

The EmbeddingsOllama node does not accept direct inputs as it serves as an embedding provider.

## Outputs

- **AI Embedding**: LangChain embedding provider for use with other AI components

## Properties

### Resource: Ollama Embeddings

#### Operation: Generate Embeddings

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Resource Locator | Ollama embedding model to use | Yes |
| Base URL | String | Ollama server endpoint URL | No |
| Request Options | Collection | Additional request configuration | No |

## UseCases

- **Privacy-Sensitive Applications**: Process confidential data without external transmission
- **Offline AI Processing**: Generate embeddings without internet connectivity
- **Cost-Effective Solutions**: Avoid per-request API fees with local processing
- **Enterprise Data Security**: Keep sensitive information within organizational boundaries
- **Research and Development**: Experiment with various models without usage costs
- **Edge Computing**: Deploy embedding generation at edge locations 