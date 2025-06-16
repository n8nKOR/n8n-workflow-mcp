# Language Models (LLMs)

## Overview

Language Models (LLMs) nodes are a collection of nodes that enable the use of conversational language models from various AI service providers in n8n workflows. These nodes support the latest language models from OpenAI, Anthropic, Google, AWS Bedrock, Azure, Mistral, Ollama, and more, capable of performing tasks such as text generation, conversation, and analysis.

## Credentials

Each language model node requires authentication information from their respective service providers:
- OpenAI API Key
- Anthropic API Key
- Google Cloud/Vertex AI authentication
- AWS Bedrock authentication
- Azure OpenAI authentication
- Mistral Cloud API Key
- Groq API Key
- DeepSeek API Key
- Cohere API Key
- Hugging Face API Token
- Ollama (no authentication required for local installation)

## Inputs

| Input Type | Description | Required |
|---|---|---|
| (None) | Language model nodes have no input connections | No |

## Outputs

| Output Type | Description |
|---|---|
| Language Model | Language model output |

## Properties

### Resource: Language Model Collection

#### Operation: Provide Language Model

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model Provider | Options | Language model provider (OpenAI, Anthropic, Google, etc.) | Yes |
| Model | Resource Locator | Specific model to use | Yes |

#### Options

| Field Name | Type | Description | Required |
|---|---|---|---|
| Temperature | Number | Controls randomness in responses (0-1) | No |
| Max Tokens | Number | Maximum number of tokens to generate | No |
| Top P | Number | Controls diversity via nucleus sampling (0-1) | No |
| Frequency Penalty | Number | Penalty for repeating tokens | No |
| Presence Penalty | Number | Penalty for new topics | No |
## Supported Models

### OpenAI
- GPT-4o (latest multimodal model)
- GPT-4o-mini (lightweight version)
- GPT-4 Turbo
- GPT-3.5 Turbo

### Anthropic
- Claude 3.5 Sonnet
- Claude 3 Opus
- Claude 3 Haiku

### Google
- Gemini 1.5 Pro
- Gemini 1.5 Flash
- Gemini 1.0 Pro

### AWS Bedrock
- Claude 3 (Anthropic)
- Llama 2/3 (Meta)
- Titan (Amazon)
- Command (Cohere)

### Others
- Mistral Large/Medium/Small
- Llama models via Ollama
- Various models via Groq
- DeepSeek models
- Cohere Command models

## Advanced Features

### Tool Calling
Most modern models support function calling (Tool Calling) capabilities to interact with external tools.

### JSON Mode
Models like OpenAI and Google can generate structured JSON responses.

### Vision Capabilities
Multimodal models like GPT-4o and Gemini provide image analysis capabilities.

### Streaming
Support real-time response streaming to receive long responses progressively.

## UseCases

- **Content Generation**: Automatically generate blog posts, marketing copy, technical documentation
- **Customer Support**: Automated customer inquiry response systems
- **Code Generation**: Automatic programming code generation and review
- **Translation Services**: Multilingual translation and localization
- **Data Analysis**: Text data analysis and insight extraction
- **Educational Tools**: Generate personalized learning content
- **Creative Support**: Support writing novels, screenplays, advertising copy
- **Summarization Services**: Automatic summarization of long documents or articles
- **Question-Answer**: Knowledge base-based question-answer systems
- **Sentiment Analysis**: Analyze emotions or intentions in text
- **Classification Tasks**: Automatic classification of text or content
- **Conversational AI**: Implement chatbots or virtual assistants 