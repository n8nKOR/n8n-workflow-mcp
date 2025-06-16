# LmOllama

## Overview

⚠️ **DEPRECATED**: Use LmChatOllama node instead. The LmOllama node was the original implementation for accessing Ollama language models, superseded by the LmChatOllama node which provides enhanced features and better integration.

## Credentials

The LmOllama node requires Ollama API credentials:

- **Ollama API**: Connection configuration for local Ollama instance

## Inputs

The LmOllama node does not accept direct inputs as it serves as a language model provider.

## Outputs

- **AI Language Model**: LangChain language model for use with other AI components

## Properties

### Resource: Ollama Language Model

#### Operation: Configure Model

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | String | Ollama model name to use | Yes |
| Base URL | String | Ollama server endpoint URL | No |
| Temperature | Number | Controls randomness in responses | No |
| Request Options | Collection | Additional request configuration | No |

## UseCases

- **Legacy Workflows**: Maintain existing workflows while planning migration to LmChatOllama
- **Local AI Processing**: Complete data privacy with local model processing
- **Offline Capabilities**: Air-gapped environments without external dependencies
- **Cost-Effective AI**: No API costs with local model hosting
- **Custom Models**: Use fine-tuned or specialized local models
- **Development Testing**: Local development without external service dependencies 