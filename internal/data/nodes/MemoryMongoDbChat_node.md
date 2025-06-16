# MemoryMongoDbChat

## Overview

The MemoryMongoDbChat node provides MongoDB-based conversation memory storage within n8n's LangChain ecosystem. This node enables persistent, scalable conversation memory using MongoDB as the backend storage system, offering document-based storage with flexible schema design, automatic indexing, and high-performance queries. It's ideal for production environments that require reliable, distributed memory storage with MongoDB's enterprise features like replication, sharding, and advanced security.

## Credentials

The MemoryMongoDbChat node requires MongoDB database credentials:

- **MongoDB Connection**: Database connection credentials
- **Authentication**: Username/password or connection string authentication
- **SSL/TLS**: Secure connection configuration for production environments

### Credential Configuration
To configure MongoDB credentials:
1. Set up MongoDB connection (local, cloud, or Atlas)
2. Create database and collection for memory storage
3. Configure authentication and access permissions
4. Enable SSL/TLS for secure connections
5. Set up appropriate indexes for performance optimization

## Inputs

The MemoryMongoDbChat node does not accept direct inputs as it serves as a memory provider:

- **No Direct Inputs**: Memory nodes provide storage capabilities to other components
- **MongoDB Operations**: Manages memory operations through MongoDB database

## Outputs

The node outputs a memory connection:

- **AI Memory**: LangChain memory for use with chat models and agents
- **Connection Type**: `NodeConnectionTypes.AiMemory`
- **Usage**: Connect to chat models, agents, or chains that require persistent MongoDB memory

## Properties

### Resource: Memory

#### Operation: MongoDB Chat Memory

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Database Name | String | MongoDB database name for memory storage | Yes | - |
| Collection Name | String | MongoDB collection name for memory documents | Yes | chat_memory |
| Session Key | String | Unique identifier for memory session | No | default |
| Session ID | String | Session identifier for memory isolation | No | - |
| Max Messages | Number | Maximum number of messages to retain per session | No | 100 |
| TTL (Time to Live) | Number | Automatic document expiration in seconds | No | 0 |
| Index Creation | Boolean | Automatically create performance indexes | No | true |

### Field Details

**Database Configuration:**
- **Database Name**: Target MongoDB database for memory storage
- **Collection Name**: Collection to store conversation documents
- **Document Structure**: Flexible JSON document schema for conversation data

**Session Management:**
- **Session Key**: Creates unique memory namespaces for different workflows
- **Session ID**: Further subdivides memory within a session key namespace
- **Document Isolation**: Uses MongoDB queries for session isolation

**Performance Options:**
- **Max Messages**: Controls document size and query performance
- **TTL**: Automatic cleanup using MongoDB TTL indexes
- **Indexing**: Compound indexes on session fields for fast retrieval

## Technical Details

### MongoDB Integration
- **Document Storage**: Conversation data stored as JSON documents
- **Flexible Schema**: Dynamic document structure for various message types
- **Indexing Strategy**: Optimized indexes for session-based queries
- **Connection Pooling**: Efficient database connection management

### Performance Optimization
- **Compound Indexes**: Multi-field indexes for fast session queries
- **TTL Indexes**: Automatic document cleanup for memory management
- **Query Optimization**: Efficient aggregation pipelines for memory operations
- **Connection Reuse**: Persistent connections for reduced latency

### Scalability Features
- **Horizontal Scaling**: Support for MongoDB sharding across multiple servers
- **Replication**: Read scaling through MongoDB replica sets
- **Load Balancing**: Distribute queries across MongoDB cluster nodes
- **Auto-sharding**: Automatic data distribution for large-scale deployments

### Security Features
- **Authentication**: MongoDB user authentication and authorization
- **Encryption**: In-transit and at-rest encryption support
- **Access Control**: Role-based access control for memory operations
- **Audit Logging**: MongoDB audit logs for compliance requirements

### Error Handling
- **Connection Resilience**: Automatic reconnection and failover handling
- **Write Concerns**: Configurable write acknowledgment for data durability
- **Read Preferences**: Handle replica set read operations and failures
- **Transaction Support**: ACID transactions for critical memory operations

## UseCases

- **Enterprise Chat Applications** : Production-grade conversation memory with MongoDB's enterprise features and reliability
- **Multi-tenant SaaS Platforms** : Isolated memory storage for different customers using MongoDB's flexible document model
- **High-volume Conversational AI** : Scalable memory storage for applications handling thousands of concurrent conversations
- **Global Chat Applications** : Distributed memory storage across multiple geographic regions using MongoDB Atlas
- **Compliance-required Systems** : Audit-compliant memory storage with MongoDB's security and logging features
- **Real-time Analytics** : Conversation analytics using MongoDB's aggregation framework for real-time insights
- **Microservices Architecture** : Shared memory storage across distributed microservices using MongoDB as central data store
- **Development and Testing** : Flexible development environments with MongoDB's schema-less design for rapid iteration
- **Backup and Recovery** : Enterprise backup and recovery strategies using MongoDB's built-in backup solutions
- **Memory Migration** : Migrate conversation data from other systems using MongoDB's import/export tools
- **Custom Memory Queries** : Advanced memory queries using MongoDB's rich query language and aggregation pipeline
- **Memory Archival** : Long-term storage and archival of conversation data using MongoDB's compression features
- **Cross-platform Integration** : Integration with existing MongoDB-based systems and applications
- **Performance Monitoring** : Memory usage monitoring and optimization using MongoDB's built-in profiling tools
- **Data Replication** : Automatic data replication across multiple data centers for high availability 