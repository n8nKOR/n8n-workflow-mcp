# ToolCode

## Overview

The ToolCode node provides powerful custom code execution capabilities within n8n's LangChain ecosystem, enabling developers to create sophisticated, personalized tools for AI agents using JavaScript or Python. This versatile node allows the creation of custom tools that can process complex logic, integrate with external APIs, perform specialized calculations, and execute domain-specific operations that extend AI agent capabilities beyond standard tools. With support for both structured and unstructured inputs, schema validation, and sandboxed execution environments, it offers unlimited flexibility for building enterprise-grade AI solutions.

**Key Features:**
- Custom JavaScript and Python code execution for AI tools
- Structured input schema support with automatic validation
- Sandboxed execution environment for secure code processing
- Dynamic tool creation with flexible input/output handling
- Integration with AI agents and LangChain tool ecosystem
- Advanced error handling and execution monitoring
- Support for complex business logic and external integrations
- Enterprise-grade security and resource management

## Inputs

The ToolCode node does not accept direct inputs as it serves as a tool provider:

- **No Direct Inputs**: Tool nodes provide capabilities to AI agents
- **Agent-Driven**: AI agents call the tool with query parameters
- **Schema-Based**: Optionally accepts structured inputs based on defined schema

## Outputs

The node outputs a tool connection:

- **AI Tool**: Custom code-based tool for AI agent usage
- **Connection Type**: `NodeConnectionTypes.AiTool`
- **Usage**: Connect to AI agents or other components requiring custom tool functionality

## Credentials

This node does not require specific credentials as it executes code in sandboxed environments. However, the executed code may access external services that require their own authentication.

## Properties

### Resource: Tool Code

#### Operation: Execute Custom Code

| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Name of the tool (alphanumeric and underscores only) | Yes |
| Description | String | Description of tool functionality for AI agents | Yes |
| Language | Options | Programming language (JavaScript/Python) | Yes |
| Code | String | Custom code to execute | Yes |
| Specify Input Schema | Boolean | Enable structured input validation | No |
| Schema Type | Options | Schema definition method (fromJson/manual) | No |
| JSON Schema Example | String | Example JSON for automatic schema generation | No |
| Input Schema | String | Manual JSON schema definition | No |

### Configuration Details

#### Tool Name
- **Purpose**: Identifies the tool for AI agent usage
- **Format**: Alphanumeric characters and underscores only
- **Example**: `weather_lookup`, `data_processor`, `api_connector`
- **Best Practices**: Use descriptive, snake_case naming

#### Tool Description
- **Purpose**: Guides AI agents on when and how to use the tool
- **Format**: Clear, descriptive text explaining functionality
- **Example**: "Call this tool to get weather information for a specific city. Input should be the city name."
- **Guidelines**: Be specific about input requirements and expected outputs

#### Language Selection
- **JavaScript**: Full-featured, with access to standard JavaScript libraries
- **Python (Beta)**: Python execution with standard library support
- **Considerations**: Choose based on existing expertise and required functionality

#### Schema Configuration
- **Unstructured**: Simple string input via `query` parameter
- **Structured**: JSON schema-validated input with typed parameters
- **Benefits**: Type safety, input validation, better AI agent integration

## UseCases

- **Custom Business Rules**: Implement complex business logic and decision trees
- **API Integration**: Connect AI agents to external APIs and services
- **Data Transformation**: Transform and validate data according to business requirements
- **Specialized Calculations**: Perform domain-specific calculations and analysis
- **Workflow Integration**: Integrate with existing enterprise systems and workflows
- **Prototype Development**: Rapidly develop and test new AI capabilities
- **Custom Validators**: Create custom validation logic for data processing
- **External Service Connectors**: Build connectors to specialized external services
- **Real-time Processing**: Process real-time data streams and events
- **Custom Parsers**: Parse and process specialized data formats