# MemoryRedisChat

## Overview

The MemoryRedisChat node provides Redis-based conversation memory storage within n8n's LangChain ecosystem. This node leverages Redis's high-performance in-memory data structure store for ultra-fast conversation memory operations, offering sub-millisecond response times, automatic expiration, pub/sub capabilities, and horizontal scaling through Redis Cluster. Redis's rich data types and atomic operations make it ideal for real-time conversational AI applications requiring immediate memory access and updates.

## Credentials

The MemoryRedisChat node requires Redis connection credentials:

- **Redis Connection**: Database connection credentials
- **Authentication**: Password or ACL-based authentication
- **SSL/TLS**: Secure connection configuration for production environments
- **Cluster Configuration**: Redis Cluster or Sentinel setup for high availability

### Credential Configuration
To configure Redis credentials:
1. Set up Redis server (local, cloud, or managed service like Redis Cloud)
2. Configure authentication (password or Redis ACL)
3. Enable SSL/TLS for secure connections in production
4. Set up Redis Cluster or Sentinel for high availability
5. Configure memory policies and expiration settings
6. Optimize Redis configuration for conversation workloads

## Inputs

The MemoryRedisChat node does not accept direct inputs as it serves as a memory provider:

- **No Direct Inputs**: Memory nodes provide storage capabilities to other components
- **Redis Operations**: Manages memory operations through Redis database

## Outputs

The node outputs a memory connection:

- **AI Memory**: LangChain memory for use with chat models and agents
- **Connection Type**: `NodeConnectionTypes.AiMemory`
- **Usage**: Connect to chat models, agents, or chains that require high-speed Redis memory

## Properties

### Resource: Memory

#### Operation: Redis Chat Memory

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Database Index | Number | Redis database index (0-15) | No | 0 |
| Key Prefix | String | Prefix for Redis keys to namespace memory | No | chat: |
| Session Key | String | Unique identifier for memory session | No | default |
| Session ID | String | Session identifier for memory isolation | No | - |
| Max Messages | Number | Maximum number of messages to retain per session | No | 100 |
| TTL (Time to Live) | Number | Automatic key expiration in seconds | No | 3600 |
| Compression | Boolean | Enable message compression for memory efficiency | No | false |
| Atomic Operations | Boolean | Use Redis transactions for atomic updates | No | true |

### Field Details

**Redis Configuration:**
- **Database Index**: Select Redis logical database (0-15 for single instance)
- **Key Prefix**: Namespace Redis keys to avoid collisions with other applications
- **Data Structure**: Uses Redis Hashes for efficient message storage

**Session Management:**
- **Session Key**: Creates unique memory namespaces using Redis key patterns
- **Session ID**: Hierarchical key structure for memory organization
- **Key Isolation**: Complete separation of conversation data using key prefixes

**Performance Options:**
- **Max Messages**: Controls memory usage per conversation session
- **TTL**: Automatic cleanup using Redis expiration for memory management
- **Compression**: Optional message compression for reduced memory footprint
- **Atomic Operations**: MULTI/EXEC transactions for data consistency

## Technical Details

### Redis Integration
- **In-Memory Storage**: All conversation data stored in Redis memory for ultra-fast access
- **Data Structures**: Optimized use of Redis Hashes, Lists, and Sets for conversation data
- **Key Design**: Hierarchical key naming for efficient memory organization
- **Connection Pooling**: Efficient connection management for high-throughput operations

### Performance Optimization
- **Sub-millisecond Access**: Redis's in-memory architecture for instant memory operations
- **Pipeline Operations**: Batch multiple Redis commands for improved throughput
- **Lua Scripting**: Server-side scripts for complex atomic memory operations
- **Memory Optimization**: Efficient data encoding and compression options

### Scalability Features
- **Redis Cluster**: Horizontal scaling across multiple Redis nodes
- **Sharding**: Automatic data distribution using Redis Cluster's hash slots
- **Read Replicas**: Scale read operations across Redis replica nodes
- **Sentinel**: High availability with automatic failover using Redis Sentinel

### Advanced Features
- **Pub/Sub**: Real-time memory updates using Redis publish/subscribe
- **Streams**: Use Redis Streams for conversation event logging
- **Geospatial**: Location-based memory queries using Redis geospatial commands
- **JSON Support**: Store complex conversation metadata using RedisJSON module

### Memory Management
- **Expiration Policies**: Flexible TTL settings for automatic memory cleanup
- **Memory Policies**: Configure Redis eviction policies for memory-constrained environments
- **Persistence Options**: Optional RDB/AOF persistence for memory durability
- **Memory Monitoring**: Built-in Redis metrics for memory usage tracking

### Error Handling
- **Connection Resilience**: Automatic reconnection and retry mechanisms
- **Cluster Failover**: Seamless failover in Redis Cluster deployments
- **Data Validation**: Validate Redis responses and handle connection errors
- **Monitoring Integration**: Integration with Redis monitoring tools and alerts

## UseCases

- **Real-time Chat Applications** : Ultra-fast memory operations for applications requiring sub-millisecond response times
- **High-frequency Trading Bots** : Immediate memory access for financial conversation AI requiring instant decision-making
- **Gaming and Interactive AI** : Low-latency memory for real-time AI characters and interactive gaming experiences
- **Live Streaming Platforms** : Fast memory operations for AI moderation and real-time chat analysis
- **IoT and Edge Computing** : Lightweight memory storage for edge-deployed conversational AI systems
- **Mobile Applications** : Optimized memory operations for mobile chat applications with limited bandwidth
- **Microservices Architecture** : Shared memory store across distributed microservices with Redis as central cache
- **Session-based Applications** : Temporary memory storage with automatic expiration for session-based interactions
- **Real-time Analytics** : Immediate conversation insights using Redis's data structures and real-time operations
- **Load Testing and Benchmarking** : High-performance memory testing for conversational AI performance optimization
- **Caching Layer** : Memory caching for frequently accessed conversation data with automatic invalidation
- **Pub/Sub Communication** : Real-time memory synchronization across multiple AI agents using Redis pub/sub
- **Development and Prototyping** : Fast iteration and testing with Redis's simple setup and immediate feedback
- **Memory Clustering** : Distributed memory across multiple Redis nodes for large-scale conversation handling
- **Event-driven Architecture** : Memory operations integrated with Redis Streams for event-driven AI workflows