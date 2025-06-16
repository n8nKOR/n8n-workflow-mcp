# OutputParserAutofixing

## Overview

⚠️ **DEPRECATED NODE**: This node is deprecated and users are advised to use the "Structured Output Parser" instead for new implementations.

The OutputParserAutofixing node provided intelligent error recovery for output parsing within n8n's LangChain ecosystem. This wrapper node enhanced other output parsers by automatically attempting to fix malformed or incorrectly formatted outputs using a language model as a correction mechanism. When initial parsing failed, it would invoke an LLM with specific error information and formatting instructions to generate a corrected version of the output, providing robust fallback parsing capabilities.

## Credentials

This node does not require any credentials as it operates as an output parser provider within the LangChain ecosystem. Authentication is handled by the connected language model and output parser.

## Inputs

The OutputParserAutofixing node accepts two inputs:

| Input | Connection Type | Description | Required |
|---|---|---|---|
| Model | AiLanguageModel | LLM for error correction | Yes |
| Output Parser | AiOutputParser | Parser to wrap with auto-fixing | Yes |

## Outputs

The node outputs an enhanced output parser:

- **AI Output Parser**: Enhanced parser with auto-fixing capabilities
- **Connection Type**: `AiOutputParser`
- **Usage**: Connect to chains, agents, or other components requiring output parsing

## Properties

### Resource: Auto-fixing Parser

#### Operation: Configure

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Options | Collection | Additional configuration options | No | - |

### Options Configuration

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Retry Prompt | String | Custom prompt template for error correction | No | NAIVE_FIX_PROMPT |

### Retry Prompt Field Details

**Field Configuration:**
- **Type**: String with multi-line support
- **Default Value**: Uses built-in NAIVE_FIX_PROMPT template
- **Required Placeholders**: Must include "{error}", "{instructions}", and "{completion}" placeholders
- **Validation**: Node validates that prompt contains '{error}' placeholder
- **Field Hint**: "Should include \"{error}\", \"{instructions}\", and \"{completion}\" placeholders"

**Prompt Template Structure:**
The retry prompt template allows customization of how the LLM corrects parsing errors. The default template includes:
- **{error}**: The original parsing error message
- **{instructions}**: The expected format instructions from the parser
- **{completion}**: The malformed output that needs correction

**Default NAIVE_FIX_PROMPT:**
```
Instructions:
{instructions}

Completion:
{completion}

Above, the Completion did not satisfy the constraints given in the Instructions.
Error:
{error}

Please try again. Please only respond with an answer that satisfies the constraints laid out in the Instructions:
```

### Prompt Validation

**Required Elements:**
- **Error Placeholder**: '{error}' must be present in the prompt
- **Instructions Placeholder**: '{instructions}' should be included for parser context
- **Completion Placeholder**: '{completion}' should be included for original output

**Validation Logic:**
- Node validates prompt contains '{error}' placeholder before execution
- Throws validation error if required placeholder is missing
- Provides clear error messages for invalid prompt configurations

## Connection Requirements

### Connection Hint Notice
The node includes a connection hint notice for proper usage context:
- **Purpose**: Explains how to properly connect the auto-fixing parser
- **Usage Context**: Provides guidance for LangChain integration
- **Field Position**: Displayed prominently in the node configuration

## Error Handling Process

### 1. Initial Parsing Attempt
The node first attempts to parse the output using the connected output parser.

### 2. Error Detection
If parsing fails, the node captures the error information including:
- Original output text
- Error message from the parser
- Expected format specification

### 3. LLM Correction
The language model is invoked with:
- The original malformed output
- Error details and context
- Instructions on the expected format
- Custom correction prompts (if provided)

### 4. Retry Logic
The corrected output is parsed again, with retry attempts up to the configured limit.

### 5. Final Validation
The successfully parsed output is validated and returned, or an error is thrown if all attempts fail.

## Validation Logic

### Input Validation
- Verifies that both language model and output parser connections are established
- Validates that the connected components are compatible
- Checks for required configuration parameters
- Validates retry prompt contains required placeholders

### Output Validation
- Ensures the corrected output meets the expected format requirements
- Validates data types and structure consistency
- Performs semantic validation when applicable

### Error Recovery Patterns
- **Format Correction**: Fixes common formatting issues (missing brackets, quotes, etc.)
- **Structure Repair**: Corrects JSON/XML structure problems
- **Data Type Conversion**: Handles type mismatches and conversions
- **Content Enhancement**: Adds missing required fields or properties

## Deprecation Information

### Migration Path
- **Replacement Node**: Use "Structured Output Parser" for new implementations
- **Enhanced Features**: Structured Output Parser provides better reliability and performance
- **Migration Benefits**: Improved error handling and more robust parsing capabilities

### Support Status
- **Current Status**: Deprecated but functional
- **Maintenance**: Limited maintenance, no new features
- **Recommendation**: Migrate to Structured Output Parser for production use

## UseCases

⚠️ **Note**: These use cases are provided for legacy reference only. Use "Structured Output Parser" for new implementations.

- **Unreliable LLM Outputs**: Handle inconsistent formatting from various language models with automatic correction
- **Format Recovery**: Automatically correct minor formatting errors in structured outputs without manual intervention
- **Parser Robustness**: Enhance existing parsers with intelligent error recovery for improved reliability
- **Workflow Reliability**: Prevent workflow failures due to parsing errors in production environments
- **Data Quality Improvement**: Automatically improve data quality from AI outputs through correction mechanisms
- **Legacy System Integration**: Handle outputs from older or less reliable AI systems with modern correction techniques
- **Batch Processing**: Process large volumes of potentially malformed outputs with automated correction
- **Error Tolerance**: Provide fault-tolerant parsing for production systems requiring high availability
- **Development Support**: Aid in development and testing of parsing workflows with automatic error correction
- **Format Migration**: Help migrate between different output formats with intelligent conversion