# LmChatGoogleVertex

## Overview

The LmChatGoogleVertex node provides access to Google's Vertex AI language models within n8n's LangChain ecosystem. Google Vertex AI is Google Cloud's enterprise AI platform that offers advanced Gemini models with enhanced security, compliance, and scalability features. This node enables seamless integration with Vertex AI's managed AI services while maintaining compatibility with LangChain's comprehensive AI workflow tools.

## Credentials

The LmChatGoogleVertex node requires Google API credentials:

### Required Credentials
- **Google API**: Service account authentication for Google Cloud services

### Credential Configuration

| Field | Type | Description | Required |
|---|---|---|---|
| Service Account Email | String | Google Cloud service account email | Yes |
| Private Key | Password | Google Cloud service account private key | Yes |
| Region | Options | Google Cloud region for model deployment | Yes |

### Authentication Method

The node uses Google Cloud service account authentication with JSON Web Token (JWT) format.

## Inputs

The LmChatGoogleVertex node does not accept direct inputs as it serves as a language model provider:

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
| Project | Resource Locator | Google Cloud project for Vertex AI | Yes |
| Model | String | The model name (must start with 'gemini') | Yes |
| Options | Collection | Additional configuration options | No |

#### Options Collection

| Field Name | Type | Description | Required |
|---|---|---|---|
| Location | String | Google Cloud region for model deployment | No |
| Maximum Number of Tokens | Number | Maximum tokens to generate | No |
| Sampling Temperature | Number | Randomness in output (0.0 to 2.0) | No |
| Top K | Number | Top-K sampling parameter | No |
| Top P | Number | Nucleus sampling parameter (0.0 to 1.0) | No |

## UseCases

- **Enterprise AI Applications** : Production-grade AI solutions with enterprise security and compliance features
- **Multi-region Deployments** : Global AI applications with data residency compliance and regional deployment options
- **Advanced Reasoning Tasks** : Complex problem-solving with Vertex AI's enhanced Gemini models
- **Regulated Industries** : AI applications requiring audit trails, compliance monitoring, and enterprise-grade security
- **Scalable AI Services** : High-volume AI processing with managed infrastructure and automatic scaling
- **Research and Development** : Academic and scientific computing with enterprise support and enhanced capabilities
- **Document Processing** : Enterprise document analysis, summarization, and intelligent content processing 