# TextClassifier

## Overview

The TextClassifier node provides text classification capabilities within n8n's LangChain ecosystem. This node categorizes text input into predefined classes or labels using AI language models, enabling automated content organization, routing, and processing based on text content analysis.

## Credentials

The TextClassifier node requires AI language model connection:

- **AI Language Model**: Connection to any supported language model node for text classification

## Inputs

- **Text Data**: Input containing text content for classification
- **AI Language Model**: Required connection to language model for processing

## Outputs

- **Classification Results**: Classification labels, categories, and confidence scores
- **Processed Data**: Original data enhanced with classification information

## Properties

### Resource: Text Classification

#### Operation: Classify Text Content

| Field Name | Type | Description | Required |
|---|---|---|---|
| Text Field | String | Field name containing text to classify | Yes |
| Classification Classes | Collection | List of possible classification categories | Yes |
| Output Format | Options | Result format (Label Only, Label with Confidence, Detailed) | No |
| Multi-Label Classification | Boolean | Allow multiple labels per text | No |
| Confidence Threshold | Number | Minimum confidence score for classification | No |
| Custom Prompt | String | Custom classification prompt template | No |
| Batch Processing | Boolean | Process multiple texts in batch | No |

## UseCases

- **Content Categorization**: Automatically categorize articles, documents, and web content
- **Email Routing**: Route incoming emails to appropriate departments based on content
- **Support Ticket Classification**: Classify support tickets by issue type, priority, or department
- **Product Categorization**: Classify products based on descriptions and specifications
- **News Article Classification**: Categorize news articles by topic, industry, or theme
- **Document Organization**: Organize documents into folders and categories automatically
- **Content Moderation**: Classify content for appropriateness and compliance
- **Lead Qualification**: Classify leads based on inquiry content and characteristics
- **Survey Response Categorization**: Classify open-ended survey responses by topic
- **Spam Detection**: Classify messages as spam or legitimate communications 