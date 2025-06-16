# OpenAI Assistant

## Overview

The OpenAI Assistant node utilizes OpenAI's Assistant API to create and interact with AI assistants for advanced conversational AI workflows. This LangChain integration node enables you to leverage OpenAI's powerful assistant capabilities with custom instructions, native tools like Code Interpreter and Knowledge Retrieval, and seamless integration with external n8n tools for comprehensive AI-powered automation.

## Credentials

- **Name**: openAiApi
- **Required**: Yes

## Inputs

- **Main**: Primary data input from previous node
- **AI Tools** (optional): Connect custom n8n tools for enhanced functionality
- **AI Memory** (optional): Memory connector for conversation history

## Outputs

- **Main**: Returns JSON data containing the assistant's response and execution results

## Properties

### Resource: assistant

#### Operation: useNewAssistant

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | String | The name of the assistant to create | Yes |  |
| Instructions | String | Instructions that tell the assistant how it should behave and respond | No |  |
| Model | Options | OpenAI model to use for the assistant | Yes | gpt-3.5-turbo-1106 |
| Text Input | String | The message or query to send to the assistant | Yes | {{ $json.chatInput }} |
| OpenAI Tools | Multi-Options | Built-in OpenAI tools to enable (Code Interpreter, Knowledge Retrieval) | No | [] |
| Options | Collection | Additional configuration options (Base URL, Max Retries, Timeout) | No |  |

#### Operation: useExistingAssistant

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Assistant | Resource Locator | ID of the existing OpenAI assistant to use | Yes |  |
| Text Input | String | The message or query to send to the assistant | Yes | {{ $json.chatInput }} |
| OpenAI Tools | Multi-Options | Native OpenAI tools to enable (Code Interpreter, Knowledge Retrieval) | No | [] |
| Options | Collection | Additional configuration options (Base URL, Max Retries, Timeout) | No |  |

## UseCases

- **Customer Support Automation** : Create intelligent customer service assistants that can handle inquiries, access customer data, provide product information, and escalate complex issues appropriately with full conversation history
- **Data Analysis and Reporting** : Build assistants that analyze datasets, generate insights, create visualizations, and produce automated reports with code interpreter capabilities for complex calculations and data processing
- **Content Creation and Writing** : Develop writing assistants for generating blog posts, documentation, marketing copy, and technical content with specific style guidelines and brand voice consistency
- **Technical Documentation** : Create assistants that help generate API documentation, user guides, troubleshooting manuals, and knowledge base articles with retrieval capabilities for existing documentation
- **Code Review and Development** : Build programming assistants that review code, suggest improvements, generate test cases, provide development guidance, and execute code for validation
- **Research and Information Gathering** : Deploy assistants with knowledge retrieval to search through documents, compile research findings, answer questions from uploaded materials, and synthesize information from multiple sources
- **Educational and Training Support** : Create tutoring assistants that provide personalized learning experiences, answer questions, guide users through complex topics, and adapt to different learning styles
- **Business Process Automation** : Integrate assistants into workflows for decision-making, process optimization, automated business rule execution, and workflow orchestration
- **Creative Project Management** : Develop assistants for creative workflows including brainstorming, project planning, content strategy, creative brief generation, and collaborative ideation
- **Multi-Modal Task Execution** : Combine conversation, code execution, and document retrieval for complex workflows requiring multiple AI capabilities and tool interactions
- **Quality Assurance and Testing** : Create assistants that can test workflows, validate data, perform quality checks, and generate comprehensive test reports
- **Knowledge Management** : Build assistants that can organize, categorize, and retrieve organizational knowledge, maintaining up-to-date information repositories
- **Workflow Optimization** : Deploy assistants that can analyze workflow performance, suggest improvements, and automatically optimize process efficiency
- **Compliance and Audit** : Create assistants that can check compliance requirements, generate audit reports, and ensure adherence to regulatory standards
- **Integration Management** : Build assistants that can manage complex integrations, handle API interactions, and coordinate between multiple systems and services
