# LmOpenHuggingFaceInference

## Overview

The LmOpenHuggingFaceInference node provides access to Hugging Face's vast ecosystem of open-source language models through their Inference API within n8n's LangChain ecosystem. Hugging Face hosts thousands of pre-trained models from the community, offering unprecedented access to diverse AI capabilities including multilingual models, domain-specific models, and cutting-edge research implementations. This node enables seamless integration with Hugging Face's inference infrastructure while maintaining compatibility with LangChain's comprehensive AI workflow tools.

## Credentials

The LmOpenHuggingFaceInference node requires Hugging Face API credentials:

### Required Credentials
- **Hugging Face API**: Authentication for Hugging Face's inference services

### Credential Configuration

| Field | Type | Description | Required |
|---|---|---|---|
| API Key | Password | Hugging Face API access token | Yes |

### Authentication Method

The node uses Bearer token authentication for secure access to Hugging Face's inference services.

## Inputs

The LmOpenHuggingFaceInference node does not accept direct inputs as it serves as a language model provider:

- **No Direct Inputs**: Language model nodes provide AI capabilities to other components
- **Configuration-Based**: All configuration is done through node properties

## Outputs

The node outputs a language model connection:

- **AI Language Model**: LangChain language model for use with other AI components
- **Connection Type**: `NodeConnectionTypes.AiLanguageModel`
- **Usage**: Connect to chains, agents, tools, or other LangChain components

## Properties

### Resource: Language Model

#### Operation: Text Generation

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | String | HuggingFace model name to use | Yes |
| Options | Collection | Additional configuration options | No |

#### Options Collection

| Field Name | Type | Description | Required |
|---|---|---|---|
| Custom Inference Endpoint | String | Custom endpoint URL for dedicated deployments | No |
| Maximum Number of Tokens | Number | Maximum tokens to generate | No |
| Sampling Temperature | Number | Randomness in output (0.0 to 2.0) | No |
| Top K | Number | Top-K sampling parameter | No |
| Top P | Number | Nucleus sampling parameter (0.0 to 1.0) | No |

## UseCases

- **Open Source AI Development** : Access to cutting-edge open-source models from the Hugging Face community
- **Research and Experimentation** : Test latest research models and architectures for academic and commercial projects
- **Multilingual Applications** : Global applications with native language support using specialized multilingual models
- **Domain-Specific AI** : Specialized models for healthcare, legal, finance, and other industry-specific applications
- **Cost-Effective AI** : Budget-friendly access to powerful models with flexible pricing options
- **Custom Model Deployment** : Use custom fine-tuned models through dedicated inference endpoints
- **Educational Applications** : Learning and teaching with diverse model types and community contributions
