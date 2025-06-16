# RetrieverVectorStore

## Overview

The RetrieverVectorStore node provides seamless vector store retrieval capabilities within n8n's LangChain ecosystem, serving as a bridge between vector databases and retrieval-augmented generation (RAG) workflows. This node converts any compatible vector store into a standardized retriever interface, enabling semantic similarity search, document retrieval, and advanced querying capabilities. It supports both basic vector store retrieval and enhanced contextual compression retrieval with reranking, making it a fundamental component for building intelligent document discovery and knowledge retrieval systems.

## Credentials

This node does not require any credentials as it operates as a retriever provider within the LangChain ecosystem. Authentication is handled by the connected vector store.

## Inputs

The RetrieverVectorStore node requires one essential input:

### Required Inputs

| Input | Connection Type | Description | Required |
|---|---|---|---|
| Vector Store | AiVectorStore | Vector database for document retrieval | Yes |

### Input Details

#### Vector Store Input
- **Purpose**: Provides vector database capabilities for semantic search
- **Requirement**: Any LangChain-compatible vector store
- **Types Supported**: Basic VectorStore instances (Pinecone, Weaviate, Chroma, etc.) and enhanced vector stores with reranking capabilities
- **Usage**: Performs similarity search based on query embeddings

## Outputs

The node outputs a retriever connection:

- **AI Retriever**: Standardized retriever interface for document retrieval
- **Connection Type**: `NodeConnectionTypes.AiRetriever`
- **Usage**: Connect to chains, agents, or other components requiring document retrieval

## Properties

### Resource: Vector Store Retriever

#### Operation: Retrieve Documents

| Field Name | Type | Description | Required |
|---|---|---|---|
| Limit | Number | Maximum number of results to return (default: 4) | Yes |

### Configuration Details

#### Limit (topK)
- **Purpose**: Controls the maximum number of documents returned from vector search
- **Range**: 1 to unlimited (practical limit varies by vector store)
- **Default**: 4 documents
- **Performance Impact**: 
  - Lower values improve response time and reduce noise
  - Higher values provide more comprehensive results but may include less relevant documents
  - Optimal range is typically 3-10 for most use cases

## UseCases

- **Vector Database Integration** : Connect any vector store to retrieval workflows for semantic similarity search
- **Semantic Search** : Perform semantic similarity-based document retrieval across large knowledge bases
- **RAG Enablement** : Enable retrieval-augmented generation workflows with vector-based document discovery
- **Knowledge Discovery** : Facilitate intelligent document and knowledge discovery through similarity matching
- **Enterprise Search** : Build enterprise-grade search systems with vector-based relevance ranking
- **Content Recommendation** : Implement content recommendation systems based on semantic similarity
- **Research and Analysis** : Support research workflows with intelligent document retrieval and analysis
