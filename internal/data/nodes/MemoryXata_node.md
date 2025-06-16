# MemoryXata

## Overview

The MemoryXata node provides serverless, edge-distributed conversation memory storage using Xata's modern database platform within n8n's LangChain ecosystem. Xata combines the flexibility of NoSQL with the reliability of PostgreSQL, offering automatic scaling, built-in search capabilities, and global edge distribution for ultra-low latency memory access. This node enables sophisticated conversation memory management with zero-maintenance database operations and advanced query capabilities.

## Credentials

The MemoryXata node requires Xata database credentials:

### Required Credentials
- **Xata API**: Authentication and connection details for Xata database

### Credential Configuration

| Field | Type | Description | Required |
|---|---|---|---|
| API Key | Password | Xata API authentication key | Yes |
| Database Endpoint | String | Xata database endpoint URL | Yes |
| Branch | String | Xata database branch name | Yes |

### Authentication Method

The node uses API key authentication with Xata's serverless database platform for secure, scalable memory storage.

## Inputs

The MemoryXata node does not accept direct inputs as it serves as a memory provider:

- **No Direct Inputs**: Memory nodes provide storage capabilities to other components
- **Configuration-Based**: All configuration is done through node properties and credentials

## Outputs

The node outputs a memory connection:

- **AI Memory**: LangChain memory for use with chat models and agents
- **Connection Type**: `NodeConnectionTypes.AiMemory`
- **Usage**: Connect to chat models, agents, or chains that require serverless memory

## Properties

### Resource: Memory

#### Operation: Xata Memory

| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String/Expression | Unique identifier for conversation session | No |
| Context Window Length | Number | Maximum number of message pairs to remember | No |

### Configuration Details

#### Session ID
- **Purpose**: Unique identifier for conversation sessions in Xata
- **Generation**: Auto-generated from workflow context or manually specified
- **Format**: String value that identifies the conversation thread
- **Versioning**: Different approaches across node versions for flexibility

#### Context Window Length
- **Purpose**: Number of message pairs to maintain in active memory
- **Availability**: Version 1.3 and later
- **Range**: 1 to unlimited (practical limit: 100)
- **Performance**: Balanced context depth with query efficiency

## UseCases

- **Serverless Memory Storage** : Zero-maintenance conversation memory with automatic scaling and global edge distribution
- **Enterprise Chat Applications** : Production-grade memory for high-availability chat systems with built-in redundancy
- **Global Applications** : Low-latency memory access worldwide through Xata's edge-distributed infrastructure
- **Development and Testing** : Flexible memory management with branch-based development workflows
- **Analytics and Monitoring** : Built-in conversation analytics and monitoring through Xata's dashboard
- **Real-time Collaboration** : Support for concurrent access and real-time collaboration features
- **Scalable AI Systems** : Automatic scaling for applications with varying memory requirements
