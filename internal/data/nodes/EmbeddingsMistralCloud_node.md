# EmbeddingsMistralCloud

## Overview

Text embedding using Mistral AI cloud-based embedding services

## Credentials

- Name: Mistral Cloud API, Required: Yes

## Inputs

- No direct inputs (serves as embedding provider)

## Outputs

- AI Embedding

## Properties

### Resource: Embeddings

#### Operation: Generate Embeddings

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Options | The model which will compute the embeddings | Yes |
| Options | Collection | Additional configuration options | No |

#### Options Collection

| Field Name | Type | Description | Required |
|---|---|---|---|
| Batch Size | Number | Maximum number of documents to send in each request (max: 2048) | No |
| Strip New Lines | Boolean | Whether to strip new lines from the input text | No |

## UseCases

- **European AI Compliance** : Build AI applications with European data sovereignty requirements
- **Privacy-Sensitive Applications** : Process confidential content with strong privacy protections
- **Semantic Search Systems** : Create intelligent search with Mistral's embedding quality
- **Document Classification** : Automatically categorize documents with semantic understanding
- **Content Recommendation** : Build recommendation engines with semantic similarity
- **Knowledge Management** : Develop enterprise knowledge bases with semantic search
- **Customer Support** : Power intelligent support systems with semantic matching
- **Academic Research** : Process research documents with advanced embedding models
- **Legal Document Analysis** : Analyze legal content with privacy-preserving AI
- **Financial Applications** : Process financial documents with regulatory compliance
- **Healthcare Systems** : Handle medical data with strong privacy protections
- **Government Applications** : Deploy in public sector with European AI standards
- **Enterprise RAG Systems** : Build retrieval-augmented generation with semantic search
- **Multilingual Applications** : Process content in European languages effectively
- **Content Moderation** : Identify similar or policy-violating content
- **Social Media Analytics** : Understand social content themes and sentiment