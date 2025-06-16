# LmChatAzureOpenAi

## Overview

The LmChatAzureOpenAi node provides access to Azure OpenAI's language models within n8n's LangChain ecosystem. This node enables integration with OpenAI models hosted on Microsoft Azure infrastructure, offering enterprise-grade security, compliance, and regional deployment options.

## Credentials

The LmChatAzureOpenAi node supports two authentication methods:

- **Azure OpenAI API**: Traditional API key authentication  
- **Azure Entra ID Cognitive Services OAuth2**: Enterprise OAuth2 authentication

## Inputs

The LmChatAzureOpenAi node does not accept direct inputs as it serves as a language model provider.

## Outputs

- **AI Language Model**: LangChain language model for use with other AI components

## Properties

### Resource: Azure OpenAI Chat Model

#### Operation: Configure Model

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | Authentication method selection (API Key, Azure Entra ID OAuth2) | Yes |
| Model (Deployment) Name | String | The name of the model deployment to use | Yes |
| Frequency Penalty | Number | Positive values penalize new tokens based on their existing frequency | No |
| Maximum Number of Tokens | Number | The maximum number of tokens to generate in the completion | No |
| Response Format | Options | Output format (Text, JSON) | No |
| Presence Penalty | Number | Positive values penalize new tokens based on whether they appear in the text | No |
| Sampling Temperature | Number | Controls randomness in completions | No |
| Timeout (Ms) | Number | Maximum amount of time a request is allowed to take in milliseconds | No |
| Max Retries | Number | Maximum number of retries to attempt on failure | No |
| Top P | Number | Controls diversity via nucleus sampling | No |

## UseCases

- **Enterprise AI Applications**: Build secure AI applications with Azure's enterprise-grade infrastructure and compliance frameworks
- **Multi-language Support**: Create chatbots and AI assistants that support multiple languages and regions
- **Content Generation**: Generate high-quality content, documentation, and marketing materials using GPT models
- **Code Analysis and Generation**: Analyze code, generate documentation, and assist with software development tasks
- **Customer Service Automation**: Build intelligent customer service systems with natural language understanding
- **Data Analysis and Insights**: Process and analyze large datasets to extract meaningful insights and recommendations 