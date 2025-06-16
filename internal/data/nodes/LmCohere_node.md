# LmCohere

## Overview

The LmCohere node provides access to Cohere's enterprise-focused language models within n8n's LangChain ecosystem. Cohere specializes in building large language models for enterprise applications, with a strong focus on text generation, retrieval-augmented generation (RAG), and multilingual capabilities. This node enables seamless integration with Cohere's API while maintaining compatibility with LangChain's comprehensive AI workflow tools.

## Credentials

The LmCohere node requires Cohere API credentials:

### Required Credentials
- **Cohere API**: Authentication for Cohere's language model services

### Credential Configuration

| Field | Type | Description | Required |
|---|---|---|---|
| API Key | Password | Cohere API access token | Yes |

### Authentication Method

The node uses Bearer token authentication for secure access to Cohere's API services.

## Inputs

The LmCohere node does not accept direct inputs as it serves as a language model provider:

- **No Direct Inputs**: Language model nodes provide AI capabilities to other components
- **Configuration-Based**: All configuration is done through node properties

## Outputs

The node outputs a language model connection:

- **AI Language Model**: LangChain language model for use with other AI components
- **Connection Type**: `NodeConnectionTypes.AiLanguageModel`
- **Usage**: Connect to chains, agents, tools, or other LangChain components

## Properties

### Resource: Language Model

#### Operation: Text Completion

| Field Name | Type | Description | Required |
|---|---|---|---|
| Maximum Number of Tokens | Number | Maximum number of tokens to generate | No |
| Model | String | The name of the model to use | No |
| Sampling Temperature | Number | Controls randomness of completions | No |

## UseCases

- **Enterprise Text Generation** : Business document creation and automation with enterprise-grade language models
- **Retrieval-Augmented Generation** : Knowledge base integration and Q&A systems with specialized RAG optimization
- **Multilingual Applications** : Global content creation in 100+ languages with strong multilingual support
- **Customer Support** : Intelligent customer service automation with enterprise-focused capabilities
- **Business Intelligence** : Strategic analysis and insights generation for enterprise decision-making
- **Document Processing** : Enterprise document analysis, summarization, and intelligent content processing
- **Workflow Automation** : AI-enhanced business process automation with cost-effective enterprise pricing
