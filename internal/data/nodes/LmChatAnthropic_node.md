# LmChatAnthropic

## Overview

The LmChatAnthropic node provides access to Anthropic's Claude language models within n8n's LangChain ecosystem. This node enables integration with Claude's advanced AI capabilities, including the latest Claude 4 Sonnet model, offering features like advanced reasoning, thinking mode, and robust conversation handling.

## Credentials

The LmChatAnthropic node requires Anthropic API credentials:

- **Anthropic API**: Authentication for Anthropic Claude services

## Inputs

The LmChatAnthropic node does not accept direct inputs as it serves as a language model provider.

## Outputs

- **AI Language Model**: LangChain language model for use with other AI components

## Properties

### Resource: Anthropic Chat Model

#### Operation: Configure Model

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Resource Locator/Options | Claude model selection | Yes |
| Max Tokens to Sample | Number | Maximum response length | No |
| Temperature | Number | Response randomness (0-1) | No |
| Top K | Number | Low probability cutoff | No |
| Top P | Number | Nucleus sampling threshold (0-1) | No |
| Enable Thinking | Boolean | Activate thinking mode | No |
| Thinking Budget | Number | Tokens for thinking process (minimum 1024) | No |

## UseCases

- **Advanced Reasoning Tasks**: Complex problem-solving with thinking mode and enhanced reasoning capabilities
- **Content Creation**: High-quality writing and creative content generation using Claude's advanced language skills
- **Code Analysis**: Advanced code review, generation, and debugging with Claude's programming expertise
- **Research Assistance**: Information synthesis and analysis for academic and business research
- **Conversational AI**: Sophisticated chatbots and virtual assistants with natural conversation flow
- **Document Processing**: Analysis and summarization of complex documents and reports 