# LmChatOllama

## Overview

The LmChatOllama node provides access to locally-hosted language models through Ollama within n8n's LangChain ecosystem. Ollama enables running large language models locally on your own hardware, providing complete data privacy, offline capabilities, and full control over your AI infrastructure.

## Credentials

The LmChatOllama node requires Ollama API credentials:

- **Ollama API**: Connection configuration for local Ollama instance

## Inputs

The LmChatOllama node does not accept direct inputs as it serves as a language model provider.

## Outputs

- **AI Language Model**: LangChain language model for use with other AI components

## Properties

### Resource: Ollama Chat Model

#### Operation: Chat Completion

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Resource Locator | Ollama model to use for chat completion | Yes |
| Base URL | String | Ollama server endpoint URL | No |
| Temperature | Number | Controls randomness in responses (0-1) | No |
| Top P | Number | Controls diversity via nucleus sampling (0-1) | No |
| Top K | Number | Limits vocabulary for each step | No |
| Repeat Penalty | Number | Penalty for repeating tokens | No |
| Context Length | Number | Maximum context length for the model | No |
| Request Options | Collection | Additional request configuration | No |

## UseCases

- **Private AI Applications**: Complete data privacy with local processing
- **Offline AI Processing**: Air-gapped environments and offline capabilities
- **Cost-Effective AI**: No API costs, one-time hardware investment
- **Custom Model Deployment**: Use fine-tuned or specialized models
- **Development and Testing**: Local development without external dependencies
- **Enterprise On-Premises**: Regulatory compliance and data sovereignty
- **Edge Computing**: AI processing at the edge without cloud connectivity
- **Research and Experimentation**: Academic research with full model control 