# MemoryBufferWindow

## Overview

The MemoryBufferWindow node provides efficient conversation memory management within n8n's LangChain ecosystem using a sliding window approach. This memory node maintains a configurable buffer of recent conversation history, automatically discarding older messages when the buffer reaches its maximum capacity. It's designed for AI applications that need context awareness but want to manage memory usage efficiently by keeping only the most recent interactions.

## Credentials

This node does not require direct credentials as it operates within the n8n workflow context:

- **Workflow Access**: Standard n8n workflow execution permissions
- **Node Dependencies**: Inherits credentials from connected nodes when applicable
- **System Access**: Uses n8n's built-in security and access control mechanisms

## Inputs

The MemoryBufferWindow node does not accept direct inputs as it serves as a memory provider:

- **No Direct Inputs**: Memory nodes provide storage capabilities to other components
- **Configuration-Based**: All configuration is done through node properties

## Outputs

The node outputs a memory connection:

- **AI Memory**: LangChain memory for use with chat models and agents
- **Connection Type**: `NodeConnectionTypes.AiMemory`
- **Usage**: Connect to chat models, agents, or chains that require conversation memory

## Properties

### Resource: Memory

#### Operation: Buffer Window Memory

| Field Name | Type | Description | Required |
|---|---|---|---|
| Session Key | String | Unique identifier for memory session | No |
| Session ID | String | Session identifier for memory isolation | No |
| Context Window Length | Number | Maximum number of message pairs to remember | No |

### Configuration Details

#### Session Key
- **Purpose**: Creates unique memory namespaces for different workflows
- **Default**: Uses workflow ID for automatic isolation
- **Custom Values**: Can be customized for cross-workflow memory sharing
- **Format**: Any string value that uniquely identifies the session

#### Session ID
- **Purpose**: Further subdivides memory within a session key namespace
- **Default**: "default" for single-session scenarios
- **Use Cases**: Multiple concurrent conversations within the same workflow
- **Format**: Alphanumeric string identifier

#### Context Window Length
- **Purpose**: Controls how many message pairs (user + AI) to retain
- **Range**: 1 to 100 message pairs
- **Memory Impact**: Higher values consume more memory
- **Performance**: Lower values provide faster access

## UseCases

- **Conversation Context** : Maintain recent conversation history for AI interactions with configurable memory management
- **Memory Efficiency** : Control memory usage through configurable window sizes for resource optimization
- **Session Management** : Isolate conversation memory across different sessions and workflows
- **Real-time Chat** : Provide immediate access to recent conversation context for chat applications
- **Multi-user Applications** : Support multiple concurrent conversations with isolated memory buffers
- **Resource Optimization** : Automatic cleanup of unused memory buffers to prevent resource leaks
- **Development and Testing** : Efficient memory management for AI application development and testing
