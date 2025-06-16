# OutputParserStructured

## Overview

The OutputParserStructured node provides advanced structured output parsing capabilities within n8n's LangChain ecosystem, designed to convert unstructured language model responses into precisely defined JSON formats. This parser ensures reliable data extraction by enforcing strict schema validation and offering intelligent auto-fixing mechanisms when outputs don't conform to expected structures. It's the cornerstone for building robust AI workflows that require consistent, structured data output from language models.

## Credentials

This node does not require specific credentials as it operates on data structures and parsing logic. However, when using the auto-fix feature, the connected language model may require its own credentials.

## Inputs

The OutputParserStructured node accepts optional language model input for auto-fixing:

### Optional Inputs

| Input | Connection Type | Description | Required |
|---|---|---|---|
| Language Model | AiLanguageModel | LLM for auto-fixing malformed outputs | No |

### Input Details

#### Language Model Input (Auto-Fix)
- **Purpose**: Provides LLM for automatic error correction when parsing fails
- **Requirement**: Only required when auto-fix feature is enabled
- **Compatibility**: Any LangChain-compatible language model
- **Usage**: Invoked when primary parsing fails to correct output format

## Outputs

The node outputs a structured output parser:

- **AI Output Parser**: Enhanced parser with schema validation and auto-fixing
- **Connection Type**: `NodeConnectionTypes.AiOutputParser`
- **Usage**: Connect to chains, agents, or other components requiring structured output

## Properties

### Resource: Output Parser

#### Operation: Parse Structured Output

| Field Name | Type | Description | Required |
|---|---|---|---|
| Schema Type | Options | Method for defining schema (fromJson/manual) | Yes |
| JSON Example | JSON | Example JSON for schema generation | No |
| JSON Schema | JSON | Manual JSON schema definition | No |
| Auto-Fix | Boolean | Enable automatic error correction | No |
| Retry Prompt | String | Custom prompt for auto-fix attempts | No |

### Configuration Details

#### Schema Type Options

##### From JSON Example (fromJson)
- **Purpose**: Generate schema automatically from provided JSON example
- **Benefits**: Simple setup, intuitive configuration
- **Usage**: Provide a representative JSON example
- **Schema Generation**: Automatically infers types and structure

##### Manual JSON Schema (manual)
- **Purpose**: Define precise schema using JSON Schema specification
- **Benefits**: Fine-grained control, complex validation rules
- **Usage**: Write complete JSON Schema definition
- **Validation**: Full JSON Schema v7 specification support

#### Auto-Fix Configuration
- **Purpose**: Automatically correct malformed outputs using LLM
- **Requirement**: Language model connection required when enabled
- **Process**: Re-invoke LLM with error details and correction instructions
- **Retry Logic**: Configurable retry prompts and attempts

## UseCases

- **Data Extraction** : Extract structured data from unstructured language model outputs with reliable schema validation
- **API Integration** : Ensure consistent JSON formats for API calls and data processing workflows
- **Form Processing** : Parse natural language responses into structured form data with validation
- **Database Operations** : Convert AI responses into database-ready structured formats
- **Workflow Automation** : Enable reliable downstream processing with guaranteed data structures
- **Quality Assurance** : Implement auto-fixing mechanisms to handle malformed outputs gracefully
- **Enterprise Applications** : Build production-grade AI systems requiring consistent data formats
