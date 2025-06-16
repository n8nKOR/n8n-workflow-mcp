# MemoryManager

## Overview

The MemoryManager node provides comprehensive conversation memory management within n8n's LangChain ecosystem. This node acts as a central memory coordinator, enabling sophisticated memory operations including storage, retrieval, and management of conversation history across multiple sessions and workflows. It supports various memory backends and provides advanced features like memory persistence, session isolation, and automated cleanup operations.

## Credentials

No specific credentials required - operates within existing LangChain memory context and uses configured memory backend credentials.

## Inputs

The MemoryManager node does not accept direct inputs as it serves as a memory provider:

- **No Direct Inputs**: Memory manager nodes provide management capabilities to other components
- **Memory Operations**: Manages memory operations for connected chat models and agents

## Outputs

The node outputs a memory connection:

- **AI Memory**: LangChain memory for use with chat models and agents
- **Connection Type**: `NodeConnectionTypes.AiMemory`
- **Usage**: Connect to chat models, agents, or chains that require sophisticated memory management

## Properties

### Resource: Memory

#### Operation: Memory Management

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Memory Backend | Options | Type of memory backend (buffer, file, database) | Yes | buffer |
| Session Key | String | Unique identifier for memory session | No | default |
| Session ID | String | Session identifier for memory isolation | No | - |
| Max Messages | Number | Maximum number of messages to retain | No | 10 |
| Memory Persistence | Boolean | Enable persistent memory storage | No | false |
| Auto Cleanup | Boolean | Enable automatic cleanup of old sessions | No | true |
| Cleanup Interval | Number | Hours between cleanup operations | No | 24 |

### Field Details

**Memory Backend:**
- **Buffer**: In-memory storage with configurable size limits
- **File**: File-based persistence for local storage requirements
- **Database**: Database-backed storage for production environments

**Session Management:**
- **Session Key**: Creates unique memory namespaces for different workflows
- **Session ID**: Further subdivides memory within a session key namespace
- **Isolation**: Ensures memory separation between different conversations

**Memory Control:**
- **Max Messages**: Controls memory size and performance impact
- **Persistence**: Enables memory survival across workflow restarts
- **Auto Cleanup**: Prevents memory leaks and manages storage space

## Technical Details

### Memory Backend Support
- **Buffer Memory**: Fast in-memory operations with configurable limits
- **File-based Memory**: Persistent storage using local file system
- **Database Memory**: Scalable storage using database backends
- **Custom Backends**: Support for custom memory implementations

### Session Isolation
- **Workflow Separation**: Isolate memory across different workflows
- **User Separation**: Maintain separate memory for different users
- **Context Switching**: Dynamic switching between memory contexts
- **Memory Namespacing**: Hierarchical memory organization

### Performance Optimization
- **Memory Pooling**: Efficient memory allocation and reuse
- **Lazy Loading**: Load memory data only when needed
- **Compression**: Compress stored conversation data
- **Indexing**: Fast retrieval through memory indexing

### Error Handling
- **Memory Validation**: Validate memory operations and data integrity
- **Fallback Mechanisms**: Handle memory backend failures gracefully
- **Recovery Procedures**: Automatic recovery from memory corruption
- **Monitoring**: Built-in memory usage monitoring and alerts

## UseCases

- **Enterprise Chat Applications** : Centralized memory management for large-scale conversational AI systems
- **Multi-tenant Applications** : Isolated memory management for different customers or organizations
- **Development and Testing** : Flexible memory management for AI application development workflows
- **Session Management** : Advanced session handling for complex conversational scenarios
- **Memory Optimization** : Efficient memory usage in resource-constrained environments
- **Conversation Analytics** : Memory-based analysis and insights into conversation patterns
- **Compliance and Auditing** : Persistent memory storage for regulatory compliance requirements
- **Cross-workflow Memory** : Share memory context across multiple related workflows
- **Memory Backup and Recovery** : Implement backup and recovery strategies for conversation data
- **Performance Monitoring** : Monitor and optimize memory usage across AI applications
- **Custom Memory Solutions** : Implement specialized memory requirements using custom backends
- **Memory Migration** : Migrate memory data between different storage systems
- **Real-time Memory Operations** : High-performance memory operations for real-time applications
- **Memory Scaling** : Scale memory operations across distributed systems
- **Memory Security** : Implement secure memory handling with encryption and access controls 