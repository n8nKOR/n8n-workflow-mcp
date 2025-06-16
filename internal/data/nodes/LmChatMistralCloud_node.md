# LmChatMistralCloud

## Overview

The LmChatMistralCloud node provides access to Mistral AI's advanced language models through their cloud platform within n8n's LangChain ecosystem. Mistral AI specializes in creating efficient, high-performance language models with excellent multilingual capabilities and strong reasoning skills. This node enables seamless integration with Mistral's cloud API while maintaining compatibility with LangChain's comprehensive AI workflow tools.

## Credentials

The LmChatMistralCloud node requires Mistral Cloud API credentials:

### Required Credentials
- **Mistral Cloud API**: Authentication for Mistral AI's cloud services

### Credential Configuration

| Field | Type | Description | Required |
|---|---|---|---|
| API Key | Password | Mistral AI API access token | Yes |

### Authentication Method

The node uses Bearer token authentication for secure access to Mistral AI's cloud services.

## Inputs

The LmChatMistralCloud node does not accept direct inputs as it serves as a language model provider:

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
| Model | Options | The model which will generate the completion | Yes |
| Options | Collection | Additional configuration options | No |

#### Options Collection

| Field Name | Type | Description | Required |
|---|---|---|---|
| Maximum Number of Tokens | Number | Maximum tokens to generate | No |
| Random Seed | Number | Seed for deterministic output | No |
| Safe Mode | Boolean | Enable content safety filtering | No |
| Sampling Temperature | Number | Randomness in output (0.0 to 1.0) | No |
| Top P | Number | Nucleus sampling parameter (0.0 to 1.0) | No |

## UseCases

- **Multilingual Applications** : Excellent support for European languages with high-quality translation capabilities
- **Content Creation** : High-quality writing and creative tasks with sophisticated language understanding
- **Code Generation** : Software development assistance and analysis with strong programming language support
- **Research and Analysis** : Academic and scientific computing with advanced reasoning capabilities
- **Business Intelligence** : Strategic analysis and decision support for enterprise applications
- **Educational Applications** : Tutoring and learning assistance with multilingual support
- **Deterministic Testing** : Reproducible AI outputs for validation using random seed control
