# LmChatAwsBedrock

## Overview

The LmChatAwsBedrock node provides access to Amazon Bedrock's foundation models for conversational AI within n8n's LangChain ecosystem. This node enables integration with enterprise-grade AI models hosted on AWS infrastructure, offering enhanced security, compliance, and scalability.

## Credentials

The LmChatAwsBedrock node requires AWS credentials with Bedrock permissions:

- **AWS**: IAM credentials with Amazon Bedrock access

## Inputs

The LmChatAwsBedrock node does not accept direct inputs as it serves as a language model provider.

## Outputs

- **AI Language Model**: LangChain language model for use with other AI components

## Properties

### Resource: Language Model

#### Operation: Chat Completion

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Options | The model which will generate the completion | Yes |
| Maximum Number of Tokens | Number | Maximum number of tokens to generate | No |
| Sampling Temperature | Number | Controls randomness of completions | No |

## UseCases

- **Enterprise AI**: Deploy AI models in enterprise AWS environments
- **Secure AI Processing**: Process sensitive data with AWS security controls
- **Scalable AI Solutions**: Build scalable AI applications on AWS infrastructure
- **Multi-Model Access**: Access various foundation models through AWS Bedrock
- **Compliance-Ready AI**: Use AI models that meet enterprise compliance requirements
- **Hybrid Cloud Architectures**: Integration with existing AWS infrastructure 