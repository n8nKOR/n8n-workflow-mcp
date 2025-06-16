# MemoryChatRetriever

## Overview

The MemoryChatRetriever node provides a deprecated chat memory retrieval system within n8n's LangChain ecosystem. This node has been superseded by the Chat Memory Manager node, which offers enhanced functionality and better integration with modern LangChain memory systems. The MemoryChatRetriever was designed to retrieve conversation history from memory systems but is no longer actively developed or recommended for new implementations.

**⚠️ DEPRECATION NOTICE**: This node is deprecated and should not be used in new workflows. Use the Chat Memory Manager node instead for all chat memory operations.

## Credentials

No specific credentials required - operates within existing LangChain memory context.

## Inputs

The MemoryChatRetriever node does not accept direct inputs as it serves as a memory retriever:

- **No Direct Inputs**: Memory retriever nodes provide retrieval capabilities from existing memory systems
- **Memory Context**: Connects to existing memory systems for conversation history retrieval

## Outputs

The node outputs retrieved memory data:

- **AI Memory**: Retrieved conversation memory for use with chat models and agents
- **Connection Type**: `NodeConnectionTypes.AiMemory`
- **Usage**: Connect to chat models, agents, or chains that require historical conversation context

## Properties

### Resource: Memory

#### Operation: Chat Retrieval

**Note**: This operation is deprecated. Use Chat Memory Manager instead.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Session Key | String | Memory session identifier (deprecated) | No | - |
| Retrieval Mode | Options | How to retrieve memory (deprecated) | No | recent |

### Migration Guide

**From MemoryChatRetriever to Chat Memory Manager:**

1. **Replace Node**: Remove MemoryChatRetriever and add Chat Memory Manager
2. **Session Configuration**: Transfer session key settings to new memory manager
3. **Connection Updates**: Update memory connections to use new manager
4. **Testing**: Verify conversation context is maintained after migration

## UseCases

**⚠️ Note**: These use cases are provided for legacy reference only. Use Chat Memory Manager for new implementations.

- **Legacy Memory Retrieval** : Maintain existing workflows using deprecated memory retrieval (migration recommended)
- **Conversation Context Recovery** : Retrieve historical conversation context from legacy memory systems
- **Backward Compatibility** : Support for workflows created before Chat Memory Manager availability

## Technical Details

### Deprecation Timeline
- **Status**: Deprecated (Not recommended for new use)
- **Replacement**: Chat Memory Manager node
- **Migration Required**: Yes, for optimal performance and feature access
- **Support**: Legacy support only, no new features

### Known Limitations
- Limited integration with modern LangChain memory systems
- Reduced performance compared to Chat Memory Manager
- No active development or bug fixes
- Missing features available in Chat Memory Manager

### Migration Benefits
- **Enhanced Performance**: Chat Memory Manager offers better memory operations
- **Modern Integration**: Full compatibility with latest LangChain features
- **Active Development**: Ongoing improvements and feature additions
- **Better Error Handling**: More robust error handling and validation 