# AgentV1

## Overview

Legacy AI agent implementation for backwards compatibility in n8n workflows

## Credentials

- Name: Language Model Credentials, Required: Yes (inherited from connected LLM node)
- Name: Database Credentials, Required: Conditional (for SQL Agent operations)

## Inputs

- Language Model
- Tools
- Memory (optional)
- Output Parser (optional)
- Database Connection (optional)

## Outputs

- Main

## Properties

### Resource: AI Agent

#### Operation: Execute Agent

| Field Name | Type | Description | Required |
|---|---|---|---|
| Agent Type | Options | Type of agent to use (Tools, SQL, etc.) | Yes |
| Prompt | String | System prompt for the agent | No |
| Options | Collection | Additional configuration options | No |

## UseCases

- **Legacy Workflow Maintenance** : Maintaining existing workflows built with older agent patterns
- **Backward Compatibility** : Supporting deprecated agent configurations during migration periods
- **Multi-Version Support** : Managing workflows across different n8n versions
- **Tool Integration Legacy** : Supporting older tool calling patterns
- **Custom Agent Patterns** : Implementing specific agent behaviors available only in legacy versions
- **Database Query Agents** : Using SQL Agent for database interactions
- **Conversational AI** : Simple chat interfaces using conversational agent pattern
- **Function Calling** : OpenAI-specific function calling implementations
- **Planning Systems** : Complex task planning using Plan and Execute pattern
- **Reasoning Workflows** : Step-by-step problem solving with ReAct pattern
- **Migration Testing** : Testing agent behavior across different versions
- **Compatibility Validation** : Ensuring workflow functionality across n8n versions