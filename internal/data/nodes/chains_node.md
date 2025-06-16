# LangChain Chains

## Overview

LangChain Chains is a collection of various chain nodes that utilize conversational language models. These nodes use LLM (Large Language Model) to perform tasks such as text processing, analysis, and generation. Major chain nodes include Basic LLM Chain, Summarization Chain, Retrieval QA Chain, Information Extractor, Sentiment Analysis, Text Classifier, and more.

## Credentials

Each chain node requires authentication information for the connected language model (OpenAI, Anthropic, Azure OpenAI, etc.).

## Inputs

| Input Type | Description | Required |
|---|---|---|
| Main | Data input from previous node | Yes |
| Language Model | Language model connection (OpenAI, Anthropic, etc.) | Yes |
| Output Parser | Parser for specifying output format (optional) | No |

## Outputs

| Output Type | Description |
|---|---|
| Main | Chain execution results |

## Properties

### Resource: Chain Operations

#### Operation: Execute Chain

| Field Name | Type | Description | Required |
|---|---|---|---|
| Chain Type | Options | Type of chain to execute | Yes |
| Input Variables | Collection | Variables to pass to the chain | No |
| Chain Configuration | Object | Configuration options for the chain | No |
| Output Format | Options | Format of the output data | No |

## UseCases

- **Text Summarization**: Automatically summarize long documents or articles
- **Question-Answer System**: Build document-based question-answer systems
- **Information Extraction**: Extract structured information from unstructured text
- **Sentiment Analysis**: Analyze sentiment in customer reviews, social media posts
- **Text Classification**: Automatically classify emails, documents, messages
- **Content Generation**: Automatically generate prompt-based content
- **Translation and Language Processing**: Process and translate multilingual text
- **Data Analysis**: Analyze patterns in text data and derive insights 