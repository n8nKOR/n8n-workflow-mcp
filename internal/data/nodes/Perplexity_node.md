# Perplexity Node

## Overview

Interact with the Perplexity API to generate AI responses with citations

## Credentials

- Name: perplexityApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: AI

#### Operation: Generate Response
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Model | options | Perplexity AI model to use | Yes | - |
| Messages | collection | Chat messages for conversation context | Yes | - |
| Max Tokens | number | Maximum number of tokens to generate | No | - |
| Temperature | number | Creativity/randomness of responses (0.0-2.0) | No | 0.7 |
| Top P | number | Nucleus sampling parameter (0.0-1.0) | No | 1.0 |
| Return Citations | boolean | Whether to include citations in the response | No | true |
| Return Images | boolean | Whether to include relevant images | No | false |
| Return Related Questions | boolean | Whether to return related questions | No | false |

**Available Models:**
- **llama-3.1-sonar-small-128k-online**: Fast online model
- **llama-3.1-sonar-large-128k-online**: Large online model with better quality
- **llama-3.1-sonar-huge-128k-online**: Highest quality online model

## UseCases

- **Research Assistant**: Get AI-powered research with citations
- **Content Creation**: Generate content with factual backing
- **Question Answering**: Answer questions with real-time information
- **Fact Checking**: Verify information with cited sources
- **Knowledge Synthesis**: Combine information from multiple sources
- **Educational Support**: Provide learning assistance with references
