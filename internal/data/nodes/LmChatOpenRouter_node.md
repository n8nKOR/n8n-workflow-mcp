# LmChatOpenRouter

## Overview

The LmChatOpenRouter node provides access to multiple AI models through OpenRouter's unified API within n8n's LangChain ecosystem. OpenRouter serves as a comprehensive AI gateway, offering access to models from OpenAI, Anthropic, Google, Meta, Mistral, and many other providers through a single, standardized interface.

## Credentials

The LmChatOpenRouter node requires OpenRouter API credentials:

- **OpenRouter API**: Authentication for OpenRouter's unified AI service

## Inputs

The LmChatOpenRouter node does not accept direct inputs as it serves as a language model provider.

## Outputs

- **AI Language Model**: LangChain language model for use with other AI components

## Properties

### Resource: Language Model

#### Operation: Chat Completion

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Options | The model which will generate the completion | Yes |
| Frequency Penalty | Number | Penalize new tokens based on their existing frequency | No |
| Maximum Number of Tokens | Number | Maximum tokens to generate in completion | No |
| Response Format | Options | Format of the response (Text/JSON) | No |
| Presence Penalty | Number | Penalize new tokens based on whether they appear | No |
| Sampling Temperature | Number | Controls randomness of completions | No |
| Timeout | Number | Maximum request time in milliseconds | No |
| Max Retries | Number | Maximum number of retries to attempt | No |
| Top P | Number | Nucleus sampling parameter | No |

## UseCases

- **Multi-Provider AI Strategy**: Access multiple AI providers through single integration
- **Cost Optimization**: Compare and choose most cost-effective models for tasks
- **Provider Redundancy**: Fallback options when primary provider unavailable
- **Model Comparison**: A/B testing different models and providers
- **Specialized Tasks**: Access to specialized models for specific use cases
- **Research and Development**: Experiment with cutting-edge models
- **Global Deployment**: Access providers with different regional strengths
- **Content Creation**: High-quality writing with provider diversity
- **Code Generation**: Access to specialized coding models like CodeLlama
- **Multilingual Applications**: Leverage different providers' language strengths 