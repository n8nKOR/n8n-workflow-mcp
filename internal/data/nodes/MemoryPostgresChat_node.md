# MemoryPostgresChat

## Overview

The MemoryPostgresChat node provides PostgreSQL-based conversation memory storage within n8n's LangChain ecosystem. This node leverages PostgreSQL's robust relational database capabilities for conversation memory management, offering ACID transactions, advanced indexing, full-text search, and enterprise-grade reliability. PostgreSQL's JSON support enables flexible conversation data storage while maintaining the benefits of a structured relational database with complex queries and data integrity constraints.

## Credentials

The MemoryPostgresChat node requires PostgreSQL database credentials:

- **PostgreSQL Connection**: Database connection credentials
- **Authentication**: Username/password authentication
- **SSL/TLS**: Secure connection configuration for production environments
- **Connection Pooling**: Optimized connection management settings

### Credential Configuration
To configure PostgreSQL credentials:
1. Set up PostgreSQL server (local, cloud, or managed service)
2. Create database and schema for memory storage
3. Configure user authentication and permissions
4. Enable SSL/TLS for secure connections
5. Set up connection pooling for optimal performance
6. Create indexes for conversation query optimization

## Inputs

The MemoryPostgresChat node does not accept direct inputs as it serves as a memory provider:

- **No Direct Inputs**: Memory nodes provide storage capabilities to other components
- **PostgreSQL Operations**: Manages memory operations through PostgreSQL database

## Outputs

The node outputs a memory connection:

- **AI Memory**: LangChain memory for use with chat models and agents
- **Connection Type**: `NodeConnectionTypes.AiMemory`
- **Usage**: Connect to chat models, agents, or chains that require PostgreSQL-based memory

## Properties

### Resource: Memory

#### Operation: PostgreSQL Chat Memory

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Schema Name | String | PostgreSQL schema name for memory tables | No | public |
| Table Name | String | Table name for conversation storage | Yes | chat_memory |
| Session Key | String | Unique identifier for memory session | No | default |
| Session ID | String | Session identifier for memory isolation | No | - |
| Max Messages | Number | Maximum number of messages to retain per session | No | 100 |
| Full Text Search | Boolean | Enable PostgreSQL full-text search indexing | No | true |
| JSON Storage | Boolean | Use JSONB columns for flexible message storage | No | true |
| Auto Vacuum | Boolean | Enable automatic table maintenance | No | true |

### Field Details

**Database Configuration:**
- **Schema Name**: PostgreSQL schema for organizing memory tables
- **Table Name**: Primary table for conversation message storage
- **Data Types**: Optimized column types for conversation data

**Session Management:**
- **Session Key**: Creates unique memory namespaces using primary keys
- **Session ID**: Compound key structure for hierarchical memory organization
- **Isolation**: Row-level security for multi-tenant memory isolation

**Advanced Features:**
- **Full Text Search**: PostgreSQL's built-in text search for conversation queries
- **JSON Storage**: JSONB columns for flexible, queryable message metadata
- **Auto Vacuum**: Automatic table maintenance for optimal performance

## Technical Details

### PostgreSQL Integration
- **ACID Transactions**: Full transaction support for memory operations
- **Relational Structure**: Normalized tables for efficient storage and queries
- **JSON Support**: JSONB columns for flexible conversation metadata
- **Advanced Indexing**: B-tree, GIN, and text search indexes

### Performance Optimization
- **Query Optimization**: PostgreSQL query planner for efficient memory retrieval
- **Index Strategies**: Compound indexes on session and timestamp columns
- **Connection Pooling**: PgBouncer integration for connection management
- **Partitioning**: Table partitioning for large-scale conversation storage

### Advanced Features
- **Full-Text Search**: PostgreSQL's text search for conversation content
- **JSON Queries**: Complex queries on conversation metadata using JSONB
- **Window Functions**: Advanced analytics on conversation patterns
- **Triggers**: Database triggers for automated memory maintenance

### Scalability Features
- **Read Replicas**: Scale read operations across multiple database instances
- **Partitioning**: Horizontal table partitioning for large datasets
- **Sharding**: Application-level sharding across multiple databases
- **Streaming Replication**: Real-time data replication for high availability

### Security Features
- **Row Level Security**: Fine-grained access control for memory data
- **Encryption**: Column-level encryption for sensitive conversation data
- **Audit Logging**: Complete audit trail for memory operations
- **SSL/TLS**: Encrypted connections for data in transit

### Error Handling
- **Transaction Rollback**: Automatic rollback on memory operation failures
- **Constraint Validation**: Database constraints for data integrity
- **Connection Recovery**: Automatic reconnection and failover handling
- **Backup Integration**: Integration with PostgreSQL backup and recovery tools

## UseCases

- **Enterprise Chat Applications** : Production-grade conversation memory with PostgreSQL's enterprise reliability and ACID guarantees
- **Complex Query Requirements** : Advanced conversation analytics using PostgreSQL's rich SQL capabilities and window functions
- **Multi-tenant SaaS Platforms** : Secure memory isolation using PostgreSQL's row-level security and schema separation
- **Compliance and Auditing** : Complete audit trails and data integrity for regulatory compliance requirements
- **Real-time Analytics** : Conversation insights using PostgreSQL's analytical functions and real-time aggregations
- **High-availability Systems** : Memory replication and failover using PostgreSQL's streaming replication features
- **Data Integration** : Seamless integration with existing PostgreSQL-based systems and enterprise data warehouses
- **Search-enabled Applications** : Full-text search capabilities for conversation content using PostgreSQL's text search features
- **Development and Testing** : Flexible development environments with PostgreSQL's schema management and migration tools
- **Performance-critical Applications** : Optimized memory operations using PostgreSQL's advanced indexing and query optimization
- **Cross-platform Memory** : Shared memory storage accessible from multiple applications and programming languages
- **Backup and Recovery** : Enterprise-grade backup and point-in-time recovery for conversation data
- **Memory Migration** : Data migration and transformation using PostgreSQL's rich SQL and procedural languages
- **Custom Memory Analytics** : Advanced conversation analysis using PostgreSQL's statistical functions and extensions
- **Distributed Systems** : Memory coordination across distributed services using PostgreSQL's pub/sub and notification features 