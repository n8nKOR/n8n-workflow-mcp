# AI Agent

## Overview

The AI Agent node generates action plans and executes them using external tools. It implements LangChain-based conversational AI agents that can automatically plan complex tasks and execute them step by step. This node provides intelligent agents capable of reasoning, planning, and using various tools to accomplish goals.

## Credentials

The AI Agent node itself does not require separate authentication credentials, but the connected Chat Model and Tools may each require their own authentication credentials.

## Inputs

| Input Type | Description | Required |
|---|---|---|
| Main | Data input from previous node | No |
| Chat Model | Conversational language model (OpenAI, Anthropic, Azure OpenAI, etc.) | Yes |
| Memory | Memory for storing conversation history | No |
| Tool | Tools that the agent can use | Yes |
| Output Parser | Parser for specifying output format | No |

### Supported Chat Models
- Anthropic Claude
- Azure OpenAI
- AWS Bedrock
- Mistral Cloud
- Ollama
- OpenAI
- Groq
- Google Vertex AI
- Google Gemini
- DeepSeek
- OpenRouter
- xAI Grok
- Code (for custom implementations)

## Outputs

| Output Type | Description |
|---|---|
| Main | Agent execution results and responses |

## Properties

### Resource: AI Agent

#### Operation: Execute Agent

| Field Name | Type | Description | Required |
|---|---|---|---|
| Prompt | String | The task or question for the agent to process | Yes |
| Agent Type | Options | Type of agent to use (ReAct, Plan and Execute, etc.) | Yes |
| Max Iterations | Number | Maximum number of iterations the agent can perform | No |
| Early Stopping Method | Options | Method to use for early stopping | No |

#### Options

| Field Name | Type | Description | Required |
|---|---|---|---|
| System Message | String | System message to set agent behavior | No |
| Human Message Template | String | Template for human messages | No |
| Return Intermediate Steps | Boolean | Whether to return intermediate steps | No |
## UseCases

- **Automated Data Analysis**: Automatically perform complex tasks to analyze data and derive insights
- **Multi-step Workflows**: Plan and execute multiple steps of work sequentially
- **API Integration Automation**: Automatically process complex business logic by combining various APIs
- **Customer Support Automation**: Analyze customer inquiries and automatically respond using appropriate tools
- **Content Generation and Editing**: Automate tasks for generating and editing content such as text and images
- **Research and Information Gathering**: Automatically collect and organize information through web searches, database queries, etc.
- **Code Generation and Debugging**: Generate code solutions and debug issues using programming tools
- **Document Processing**: Extract, analyze, and process information from various document types
- **Decision Making**: Make complex decisions based on multiple data sources and criteria
- **Task Orchestration**: Coordinate multiple tools and services to accomplish complex objectives 