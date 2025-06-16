# McpTrigger

## Overview

The McpTrigger node transforms n8n workflows into Model Context Protocol (MCP) servers, enabling AI clients and agents to discover and execute n8n tools through the standardized MCP interface. This powerful trigger node acts as a bridge between n8n's workflow automation capabilities and the broader MCP ecosystem, allowing external AI systems to leverage n8n's extensive tool library and custom workflows as if they were native MCP server tools.

## Credentials

This node does not require direct credentials as it operates within the n8n workflow context:

- **Workflow Access**: Standard n8n workflow execution permissions
- **Node Dependencies**: Inherits credentials from connected nodes when applicable
- **System Access**: Uses n8n's built-in security and access control mechanisms

## Inputs

The McpTrigger node accepts AI tool connections:

- **AI Tool Input**: Connect tools that should be exposed via MCP protocol
- **Connection Type**: `NodeConnectionTypes.AiTool`
- **Tool Sources**: Any n8n tool or toolkit that provides AI capabilities

## Outputs

The node outputs webhook trigger data:

- **Main Output**: Standard webhook data when MCP tools are executed
- **Connection Type**: `NodeConnectionTypes.Main`
- **Output Data**: Tool execution results and request metadata
- **Trigger Event**: MCP tool execution triggers workflow execution

## Properties

### Resource: MCP Server

#### Operation: Start MCP Server

| Field Name | Type | Description | Required |
|---|---|---|---|
| Path | String | Webhook base path for MCP server endpoint | Yes |
| Authentication | Options | Authentication method for MCP server | Yes |
| HTTP Response Code | Number | HTTP status code for responses | No |
| Response Mode | Options | How to respond to requests | No |

### Authentication Options

#### None (No Authentication)
- **Description**: Open access without authentication requirements
- **Use Case**: Internal networks or development environments
- **Security**: No credential verification required

#### Bearer Authentication
- **Description**: Token-based authentication using Bearer tokens
- **Use Case**: API key or OAuth2 token scenarios
- **Security**: Requires valid bearer token in Authorization header

#### Header Authentication
- **Description**: Custom header-based authentication
- **Use Case**: Proprietary authentication mechanisms
- **Security**: Requires specific header name and value combination

### Webhook Configuration

#### Webhook Path Setup
The Path field configures webhook endpoints for MCP server functionality:
- **Setup Webhook**: Used for MCP server initialization and configuration
- **Default Webhook**: Handles MCP tool execution requests
- **Version-Specific Handling**: Different path handling based on node version

#### Transport Types
The node supports multiple transport mechanisms:
- **SSE (Server-Sent Events)**: Real-time streaming communication
- **StreamableHTTP**: HTTP-based request/response patterns
- **Automatic Detection**: Chooses appropriate transport based on client capabilities

## UseCases

- **Workflow Exposure** : Make n8n workflows available to external AI systems through standardized MCP protocol
- **Tool Integration** : Bridge n8n tools with MCP-compatible clients and AI agents
- **AI Ecosystem Integration** : Connect n8n to the broader AI tool ecosystem for enhanced capabilities
- **External API Access** : Allow secure external access to n8n capabilities through MCP interface
- **Enterprise AI Solutions** : Enable enterprise AI systems to leverage n8n's automation capabilities
- **Custom Tool Development** : Expose custom n8n workflows as reusable tools for AI applications
- **Protocol Compliance** : Provide MCP-compliant server functionality for standardized AI tool integration
