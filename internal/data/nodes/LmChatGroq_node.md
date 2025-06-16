# LmChatGroq

## Overview

The LmChatGroq node provides access to Groq's ultra-fast language models within n8n's LangChain ecosystem. Groq specializes in providing lightning-fast inference for large language models through their custom-designed Language Processing Units (LPUs). This node enables seamless integration with Groq's high-performance API while maintaining compatibility with LangChain's comprehensive AI workflow tools.

## Credentials

The LmChatGroq node requires Groq API credentials:

### Required Credentials
- **Groq API**: Authentication for Groq's inference service

### Credential Configuration

| Field | Type | Description | Required |
|---|---|---|---|
| API Key | Password | Groq API access token | Yes |

### Authentication Method

The node uses Bearer token authentication for secure access to Groq's API services.

## Inputs

The LmChatGroq node does not accept direct inputs as it serves as a language model provider:

- **No Direct Inputs**: Language model nodes provide AI capabilities to other components
- **Configuration-Based**: All configuration is done through node properties

## Outputs

The node outputs a language model connection:

- **AI Language Model**: LangChain language model for use with other AI components
- **Connection Type**: `NodeConnectionTypes.AiLanguageModel`
- **Usage**: Connect to chains, agents, tools, or other LangChain components

## Properties

### Resource: Language Model

#### Operation: Chat Completion

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Options | The model which will generate the completion (default: llama3-8b-8192) | Yes |
| Maximum Number of Tokens | Number | The maximum number of tokens to generate in the completion (default: 4096) | No |
| Sampling Temperature | Number | Controls randomness: Lowering results in less random completions (default: 0.7) | No |

## UseCases

- **Real-time Applications** : Ultra-fast responses for interactive applications requiring immediate feedback
- **High-throughput Processing** : Batch processing with lightning-fast inference for large-scale data operations
- **Cost-effective AI** : High-quality AI processing at competitive prices with excellent performance-to-cost ratio
- **Prototype Development** : Rapid AI application prototyping with fast iteration cycles
- **Open Source Model Access** : Access to popular open-source models like Llama, Mixtral, and Gemma
- **Gaming Applications** : Real-time NPC conversations and responses with minimal latency
- **Customer Support** : Real-time customer service automation with instant response capabilities
