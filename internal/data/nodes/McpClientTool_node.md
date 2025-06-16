# McpClientTool

## Overview

The McpClientTool node enables integration with Model Context Protocol (MCP) servers within n8n's LangChain ecosystem. This node allows AI agents and language models to access external tools and resources through standardized MCP protocols, extending their capabilities beyond built-in functions.

## Credentials

The McpClientTool node may require credentials depending on the MCP server configuration:

- **MCP Server Credentials**: Authentication required by specific MCP servers
- **Connection Configuration**: Server endpoint and protocol settings

## Inputs

- **AI Agent**: Connection from AI agent or language model requiring tool access
- **Tool Parameters**: Dynamic parameters based on available MCP tools

## Outputs

- **AI Tool**: Tool interface for use with AI agents and language models

## Properties

### Resource: MCP Client

#### Operation: Tool Integration

| Field Name | Type | Description | Required |
|---|---|---|---|
| MCP Server URL | String | Endpoint URL for the MCP server | Yes |
| Tool Name | String | Name of the specific tool to access | Yes |
| Connection Type | Options | Protocol type (HTTP, WebSocket, etc.) | No |
| Timeout | Number | Request timeout in milliseconds | No |
| Retry Attempts | Number | Maximum number of retry attempts | No |
| Authentication | Collection | Authentication configuration for MCP server | No |

## UseCases

- **External API Integration**: Connect AI agents to external APIs and services through MCP
- **Database Access**: Enable AI agents to query and manipulate databases via MCP tools
- **File System Operations**: Allow AI agents to read/write files through MCP file tools
- **Web Browsing**: Provide AI agents with web browsing capabilities via MCP web tools
- **Custom Tool Integration**: Integrate specialized tools and services into AI workflows
- **Multi-Modal Processing**: Access image, audio, and video processing tools through MCP
- **Real-Time Data**: Connect AI agents to real-time data feeds and streaming sources
- **Enterprise Integration**: Access enterprise systems and internal tools via MCP protocols
- **Third-Party Services**: Integrate with third-party services and SaaS platforms
- **Development Tools**: Provide AI agents with code compilation, testing, and deployment tools 