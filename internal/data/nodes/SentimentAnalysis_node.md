# SentimentAnalysis

## Overview

⚠️ **Node Type**: This is a **LangChain AI node** that provides sentiment analysis capabilities. LangChain nodes in n8n are specialized AI-powered components that require connection to language models and are part of the LangChain ecosystem for AI workflow automation.

**Node Category**: AI Sub-Nodes (LangChain)
**Dependencies**: Requires AI Language Model connection
**Ecosystem**: Part of n8n's LangChain integration package
**Version Management**: This node is fully standardized and integrated with n8n's LangChain ecosystem

The SentimentAnalysis node provides text sentiment analysis capabilities within n8n's LangChain ecosystem. This node analyzes text input to determine emotional tone, polarity (positive, negative, neutral), and sentiment intensity, enabling automated understanding of customer feedback, social media content, and other text-based communications.

## Credentials

The SentimentAnalysis node requires AI language model connection:

- **AI Language Model**: Connection to any supported language model node for sentiment analysis

## Inputs

- **Text Data**: Input containing text content for sentiment analysis
- **AI Language Model**: Required connection to language model for processing

## Outputs

- **Sentiment Results**: Simple sentiment classification (positive, negative, neutral)
- **Processed Data**: Original data enhanced with basic sentiment analysis

## Properties

### Resource: Sentiment Analysis

#### Operation: Analyze Text Sentiment

| Field Name | Type | Description | Required |
|---|---|---|---|
| Text Field | String | Field name containing text to analyze | Yes |
| Custom Prompt | String | Custom prompt for sentiment analysis | No |

## UseCases

- **Customer Feedback Analysis**: Analyze customer reviews, surveys, and feedback for sentiment trends
- **Social Media Monitoring**: Monitor brand sentiment across social media platforms and posts
- **Support Ticket Prioritization**: Prioritize customer support tickets based on sentiment urgency
- **Content Moderation**: Identify negative or toxic content for moderation workflows
- **Market Research**: Analyze market sentiment from news articles, reports, and discussions
- **Product Reviews**: Automatically categorize product reviews by sentiment for insights
- **Email Analysis**: Analyze customer emails and communications for sentiment and urgency
- **Chat Analytics**: Monitor live chat conversations for customer satisfaction and sentiment
- **Survey Processing**: Process open-ended survey responses for sentiment insights
- **Brand Monitoring**: Track brand perception and sentiment across various channels 