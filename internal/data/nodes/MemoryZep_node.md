# MemoryZep

## Overview

The MemoryZep node provides integration with Zep, a long-term memory store specifically designed for AI assistant applications within n8n's LangChain ecosystem. Zep offers advanced memory capabilities including automatic conversation summarization, entity extraction, intent detection, and semantic search over conversation history. This node enables sophisticated memory management with AI-powered conversation understanding, vector-based similarity search, and intelligent memory retrieval for context-aware conversational AI systems.

## Credentials

The MemoryZep node requires Zep server credentials:

- **Zep API**: Connection credentials for Zep memory server
- **API Key**: Authentication key for secure access
- **Server URL**: Zep server endpoint configuration
- **Collection Configuration**: Memory collection and organization settings

### Credential Configuration
To configure Zep credentials:
1. Deploy Zep server (self-hosted or Zep Cloud)
2. Generate API key for authentication
3. Configure server URL and connection settings
4. Set up memory collections and schemas
5. Configure AI models for summarization and extraction
6. Set up vector embeddings for semantic search

## Inputs

The MemoryZep node does not accept direct inputs as it serves as a memory provider:

- **No Direct Inputs**: Memory nodes provide storage capabilities to other components
- **Zep Operations**: Manages memory operations through Zep server API

## Outputs

The node outputs a memory connection:

- **AI Memory**: LangChain memory for use with chat models and agents
- **Connection Type**: `NodeConnectionTypes.AiMemory`
- **Usage**: Connect to chat models, agents, or chains that require AI-enhanced memory

## Properties

### Resource: Memory

#### Operation: Zep Memory

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Session ID | String | Unique session identifier for memory isolation | Yes | - |
| Collection Name | String | Zep collection name for memory organization | No | default |
| Memory Window | Number | Number of recent messages to maintain in active memory | No | 10 |
| Auto Summarize | Boolean | Enable automatic conversation summarization | No | true |
| Entity Extraction | Boolean | Enable automatic entity extraction from conversations | No | true |
| Intent Detection | Boolean | Enable intent detection for conversation understanding | No | false |
| Semantic Search | Boolean | Enable vector-based semantic search over memory | No | true |
| Relevance Threshold | Number | Similarity threshold for semantic search (0.0-1.0) | No | 0.7 |

### Field Details

**Session Management:**
- **Session ID**: Unique identifier for conversation memory isolation
- **Collection Name**: Organize memory into collections for different use cases
- **Memory Organization**: Hierarchical memory structure with collections and sessions

**AI-Enhanced Features:**
- **Auto Summarize**: Intelligent conversation summarization using AI models
- **Entity Extraction**: Automatic extraction of people, places, and concepts
- **Intent Detection**: Understanding user intentions and goals from conversations

**Semantic Capabilities:**
- **Semantic Search**: Vector-based similarity search over conversation history
- **Relevance Threshold**: Control precision of semantic memory retrieval
- **Vector Embeddings**: Automatic generation of conversation embeddings

## Technical Details

### Zep Integration
- **AI-Powered Memory**: Advanced memory operations using AI models
- **Vector Storage**: Embedding-based storage for semantic memory operations
- **Real-time Processing**: Live conversation analysis and memory updates
- **RESTful API**: Modern HTTP API for efficient memory operations

### Advanced Memory Features
- **Conversation Summarization**: AI-generated summaries of conversation history
- **Entity Recognition**: Automatic extraction and tracking of conversation entities
- **Intent Understanding**: Detection of user goals and conversation purposes
- **Semantic Similarity**: Vector-based similarity matching for contextual memory

### Performance Optimization
- **Efficient Indexing**: Optimized indexes for fast memory retrieval
- **Batch Processing**: Efficient batch operations for memory management
- **Caching Layer**: Multi-level caching for frequently accessed memory
- **Async Operations**: Non-blocking memory operations for high performance

### Scalability Features
- **Horizontal Scaling**: Support for distributed Zep deployments
- **Collection Management**: Organize memory across multiple collections
- **Load Balancing**: Distribute memory operations across multiple servers
- **Auto-scaling**: Dynamic scaling based on memory usage patterns

### AI Model Integration
- **Pluggable Models**: Support for various AI models for memory processing
- **Custom Extractors**: Custom entity and intent extraction models
- **Embedding Models**: Configurable embedding models for semantic search
- **Model Versioning**: Support for model updates and version management

### Error Handling
- **Graceful Degradation**: Fallback to basic memory when AI features fail
- **Retry Mechanisms**: Automatic retry for transient failures
- **Data Validation**: Comprehensive validation of memory operations
- **Monitoring**: Built-in monitoring and alerting for memory operations

## UseCases

- **Intelligent Virtual Assistants** : AI-enhanced memory for sophisticated virtual assistants with entity tracking and intent understanding
- **Customer Support AI** : Advanced memory for customer service with automatic issue tracking and context understanding
- **Educational AI Tutors** : Learning-aware memory that tracks student progress, concepts learned, and knowledge gaps
- **Personal AI Companions** : Long-term memory for personal AI assistants with relationship and preference tracking
- **Research and Analytics** : Memory system for research applications with entity extraction and concept mapping
- **Content Generation** : Context-aware memory for content creation with theme and topic tracking
- **Therapeutic AI Applications** : Sensitive conversation memory with emotion and progress tracking for mental health applications
- **Sales and CRM Integration** : Memory system for sales AI with lead tracking, preference detection, and relationship mapping
- **Knowledge Management** : Organizational memory for knowledge management systems with automatic categorization
- **Multi-modal AI Systems** : Unified memory across text, voice, and visual interactions with cross-modal understanding
- **Compliance and Documentation** : Memory system for regulated industries with automatic compliance tracking and documentation
- **Personalization Engines** : Advanced personalization with deep user understanding and preference evolution tracking
- **Conversation Analytics** : Research and analytics platform for conversation pattern analysis and user behavior insights
- **AI Training Data** : Memory system for collecting and organizing high-quality training data with automatic labeling
- **Cross-platform Continuity** : Unified memory across multiple platforms and devices with intelligent context transfer 