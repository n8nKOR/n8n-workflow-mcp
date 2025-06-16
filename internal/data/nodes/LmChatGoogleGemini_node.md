# LmChatGoogleGemini

## Overview

The LmChatGoogleGemini node provides access to Google's advanced Gemini language models within n8n's LangChain ecosystem. Google Gemini represents Google's next-generation multimodal AI models, designed for advanced reasoning, creative tasks, and complex problem-solving.

## Credentials

The LmChatGoogleGemini node requires Google Palm API credentials:

- **Google Palm API**: Authentication for Google Gemini language model services

## Inputs

The LmChatGoogleGemini node does not accept direct inputs as it serves as a language model provider.

## Outputs

- **AI Language Model**: LangChain language model for use with other AI components

## Properties

### Resource: Language Model

#### Operation: Chat Completion

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Options | The model which will generate the completion | Yes |
| Max Output Tokens | Number | Maximum number of tokens to generate | No |
| Temperature | Number | Controls randomness of completions | No |
| Top K | Number | Top-k sampling parameter | No |
| Top P | Number | Top-p sampling parameter | No |
| Safety Settings | Collection | Content safety configuration | No |

## UseCases

- **Conversational AI**: Build chatbots and virtual assistants
- **Content Generation**: Generate creative and informative content
- **Text Analysis**: Analyze and process text data
- **Language Translation**: Translate text between languages
- **Code Generation**: Generate code snippets and explanations
- **Advanced Reasoning**: Complex problem-solving with Gemini's reasoning capabilities
- **Multimodal Analysis**: Image understanding and description with vision models
- **Research and Analysis**: Academic research and data interpretation
- **Educational Applications**: Tutoring, learning assistance, and knowledge transfer 