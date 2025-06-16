# RerankerCohere

## Overview

The RerankerCohere node provides advanced document reranking capabilities using Cohere's state-of-the-art reranking models within n8n's LangChain ecosystem. This node is designed to improve the relevance and ordering of document results after initial retrieval from vector stores or search systems. By leveraging Cohere's specialized reranking models, it significantly enhances the quality of retrieved documents by reordering them based on semantic relevance to the user's query, making it essential for building high-performance retrieval-augmented generation (RAG) systems.

## Credentials

The RerankerCohere node requires Cohere API credentials:

### Required Credentials
- **Cohere API**: Authentication credentials for Cohere's reranking services

### Credential Configuration

| Field | Type | Description | Required |
|---|---|---|---|
| API Key | String | Cohere API key for authentication | Yes |
| Host | String | Cohere API endpoint URL | No (uses default) |

### Authentication Method

The node uses API key authentication with Cohere's reranking service for secure access to advanced neural ranking models.

## Inputs

The RerankerCohere node does not accept direct inputs as it serves as a reranker provider:

- **No Direct Inputs**: Reranker nodes provide reranking capabilities to other components
- **Configuration-Based**: All configuration is done through node properties and credentials

## Outputs

The node outputs a reranker connection:

- **AI Reranker**: LangChain reranker for document relevance improvement
- **Connection Type**: `NodeConnectionTypes.AiReranker`
- **Usage**: Connect to retrieval chains or other components requiring document reranking

## Properties

### Resource: Reranker

#### Operation: Rerank Documents

| Field Name | Type | Description | Required |
|---|---|---|---|
| Model | Options | Cohere reranking model to use | Yes |

### Available Models

#### rerank-v3.5 (Default)
- **Purpose**: Latest general-purpose reranking model
- **Performance**: Highest accuracy and relevance scoring
- **Language Support**: Multilingual with optimized English performance
- **Use Cases**: General reranking for most applications

#### rerank-english-v3.0
- **Purpose**: Specialized English-only reranking model
- **Performance**: Optimized for English language content
- **Language Support**: English only
- **Use Cases**: English-only applications requiring optimal performance

#### rerank-multilingual-v3.0
- **Purpose**: Dedicated multilingual reranking model
- **Performance**: Balanced performance across multiple languages
- **Language Support**: 100+ languages supported
- **Use Cases**: Multilingual applications and global content

## UseCases

- **Document Reranking** : Improve document ordering after initial retrieval from vector stores or search systems
- **RAG Optimization** : Optimize retrieval-augmented generation workflows with better document relevance
- **Search Enhancement** : Enhance search result quality by reordering based on semantic relevance
- **Multilingual Applications** : Support global applications with multilingual content reranking capabilities
- **Quality Improvement** : Significantly improve the quality of retrieved documents for downstream processing
- **Performance Optimization** : Reduce noise and improve precision in document retrieval systems
- **Enterprise Search** : Build enterprise-grade search and retrieval systems with advanced relevance ranking
