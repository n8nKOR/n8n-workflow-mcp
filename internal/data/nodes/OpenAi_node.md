# OpenAI (Legacy)

## Overview

**⚠️ Legacy Implementation**: This node references the legacy nodes-base implementation which is marked as hidden. The documentation describes modern OpenAI features that may not be available in this legacy implementation. This implementation only supports basic chat, image, and text resources without advanced features.

**For Modern OpenAI Features**: Consider using the newer LangChain OpenAI nodes or HTTP Request nodes for access to the latest OpenAI API capabilities including GPT-4, assistants, and advanced features.

## Credentials

This node requires OpenAI API credentials:
- **Credential Name**: `openAiApi`
- **Required Fields**: 
  - API Key: Your OpenAI API key

To obtain API credentials:
1. Sign up for an OpenAI account at https://platform.openai.com
2. Navigate to API Keys section
3. Generate a new API key
4. Copy the API key for use in n8n

## Inputs

- **Main**: Text input or parameters for OpenAI operations

## Outputs

- **Main**: Returns JSON data containing the response from OpenAI services

## Properties

### Resource: text

#### Operation: completion

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Model | Options | Text completion model to use | Yes | text-davinci-003 |
| Prompt | String | Text prompt for completion | Yes | - |
| Maximum Number of Tokens | Number | Maximum tokens to generate | No | 16 |
| Temperature | Number | Sampling temperature (0-2) | No | 1 |
| Additional Options | Collection | Advanced completion settings | No | - |

**Additional Options:**
- **Top P**: Nucleus sampling parameter
- **Frequency Penalty**: Frequency penalty for repetition
- **Presence Penalty**: Presence penalty for new topics
- **Stop Sequences**: Sequences where API will stop generating

### Resource: image

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Prompt | String | Description of the image to generate | Yes | - |
| Size | Options | Image size (256x256, 512x512, 1024x1024) | No | 1024x1024 |
| Response Format | Options | Format of response (url or b64_json) | No | url |
| Number of Images | Number | Number of images to generate (1-10) | No | 1 |

### Resource: chat

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Model | Options | Chat model to use | Yes | gpt-3.5-turbo |
| Messages | Collection | Array of chat messages | Yes | - |
| Temperature | Number | Sampling temperature (0-2) | No | 1 |
| Maximum Number of Tokens | Number | Maximum tokens to generate | No | - |
| Additional Options | Collection | Advanced chat settings | No | - |

**Message Format:**
- **Role**: System, user, or assistant
- **Content**: Message content

**Additional Options:**
- **Top P**: Nucleus sampling parameter
- **Frequency Penalty**: Frequency penalty for repetition
- **Presence Penalty**: Presence penalty for new topics
- **Stop Sequences**: Sequences where API will stop generating

## Supported Models

### Text Models (Legacy)
- **text-davinci-003**: Most capable text completion model
- **text-curie-001**: Faster, less capable than Davinci
- **text-babbage-001**: Capable of straightforward tasks
- **text-ada-001**: Fastest model for simple tasks

### Chat Models
- **gpt-3.5-turbo**: Most cost-effective chat model
- **gpt-3.5-turbo-16k**: Extended context version
- **gpt-4**: More capable but slower and more expensive
- **gpt-4-32k**: Extended context version of GPT-4

### Image Models
- **DALL-E 2**: Image generation model

## Limitations

- **Legacy API Support**: Uses older OpenAI API endpoints
- **Limited Features**: Does not support newer features like assistants, function calling, or advanced tools
- **Deprecated Models**: Some models may be deprecated by OpenAI
- **Basic Functionality**: Provides basic text, chat, and image generation without advanced features

## Migration Recommendations

For modern OpenAI features, consider:
1. **LangChain OpenAI Nodes**: For advanced chat, embeddings, and tool integration
2. **HTTP Request Node**: For direct access to latest OpenAI API endpoints
3. **Newer OpenAI Nodes**: Check for updated OpenAI nodes in n8n community

## UseCases

- **Basic Text Generation**: Generate simple text completions for content creation
- **Legacy System Integration**: Maintain compatibility with existing workflows using older API versions
- **Simple Chatbots**: Create basic conversational interfaces with chat completion
- **Image Generation**: Generate images for content creation and visual assets
- **Content Automation**: Automate content generation for blogs, descriptions, and marketing copy
- **Prototype Development**: Quickly prototype AI-powered features with simple API calls
- **Educational Projects**: Learn OpenAI integration concepts with straightforward implementation
- **Text Processing**: Perform basic text analysis, summarization, and transformation
- **Creative Writing**: Generate creative content, stories, and artistic text
- **Data Augmentation**: Generate synthetic text data for training and testing purposes

