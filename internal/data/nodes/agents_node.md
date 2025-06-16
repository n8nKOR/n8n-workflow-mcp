# Agents

## Overview

Agents are intelligent components in n8n's LangChain ecosystem that can reason, plan, and execute actions to accomplish complex tasks. These nodes combine language models with tools and memory to create autonomous AI systems capable of multi-step problem solving and decision making.

## Credentials

Agents typically inherit credentials from their connected components:

- **Language Model Credentials**: Required for the underlying AI model
- **Tool Credentials**: Required for any connected tools
- **Memory Credentials**: Required for connected memory components

## Inputs

- **Main**: Input data from previous nodes containing tasks or queries for the agent
- **Model**: Connection to an AI Language Model component (required)
- **Tools**: Connections to various tool components (optional)
- **Memory**: Connection to a memory component (optional)

## Outputs

- **Main**: Returns the agent's response and execution results

## Properties

### Resource: AI Agent

#### Operation: Task Execution

| Field Name | Type | Description | Required |
|---|---|---|---|
| Agent Type | Options | Select the type of agent (ReAct, OpenAI Functions, etc.) | Yes |
| System Message | String | System prompt that defines the agent's role and behavior | No |
| Max Iterations | Number | Maximum number of steps the agent can take | No |
| Return Intermediate Steps | Boolean | Whether to return the agent's reasoning steps | No |
| Tool Selection | Collection | Configuration for tool selection and usage | No |
| Memory Configuration | Collection | Settings for agent memory management | No |

## UseCases

- **Task Automation** : Automate complex multi-step tasks using AI reasoning
- **Customer Support** : Create intelligent customer service agents
- **Data Analysis** : Build agents that can analyze and interpret data
- **Content Creation** : Develop agents for automated content generation
- **Research Assistance** : Create research agents that can gather and synthesize information 