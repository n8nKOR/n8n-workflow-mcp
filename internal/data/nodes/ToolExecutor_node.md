# ToolExecutor

## Overview

The ToolExecutor node provides powerful tool execution capabilities within n8n's LangChain ecosystem, enabling direct invocation of AI tools without requiring a full AI agent context. This specialized node serves as a bridge between n8n workflows and LangChain tools, allowing developers to execute individual tools, toolkits, or tool collections programmatically. It supports both structured and unstructured tool inputs, automatic schema conversion, and comprehensive error handling, making it essential for testing tools, building custom workflows, and integrating tool functionality into broader automation processes.

**Key Features:**
- Direct tool execution without AI agent requirements
- Support for individual tools and comprehensive toolkits
- Automatic schema conversion and parameter validation
- Flexible input handling with JSON and string support
- Comprehensive error handling and execution monitoring
- Integration with all LangChain-compatible tools
- Tool name-based filtering for toolkit execution
- Seamless workflow integration and data processing

## Credentials

This node does not require specific credentials as it executes tools that may have their own credential requirements. The connected tools may require their own authentication depending on their functionality.

## Inputs

The ToolExecutor node accepts two types of inputs:

### Main Input
- **Connection Type**: `NodeConnectionTypes.Main`
- **Purpose**: Workflow data and execution context
- **Usage**: Standard n8n workflow connection for data flow

### Tool Input
- **Connection Type**: `NodeConnectionTypes.AiTool`
- **Purpose**: LangChain tools or toolkits to execute
- **Usage**: Connect to tool nodes (ToolCode, ToolCalculator, etc.) or toolkits

## Outputs

The node outputs execution results:

- **Main Output**: Execution results and tool responses
- **Connection Type**: `NodeConnectionTypes.Main`
- **Usage**: Connect to subsequent workflow nodes for result processing

## Properties

### Resource: Tool Execution

#### Operation: Execute Tools

| Field Name | Type | Description | Required |
|---|---|---|---|
| Query | JSON/String | Parameters to pass to the tool | Yes |
| Tool Name | String | Name of specific tool to execute (for toolkits) | No |

### Configuration Details

#### Query Parameter
- **Purpose**: Provides input data and parameters for tool execution
- **Format**: JSON object or string (automatically parsed)
- **Usage**: Contains all necessary parameters for tool functionality
- **Examples**:
  - Simple string: `"Calculate 2 + 2"`
  - JSON object: `{"operation": "add", "a": 2, "b": 2}`
  - Complex data: `{"query": "weather", "location": "New York", "date": "today"}`

#### Tool Name
- **Purpose**: Specifies which tool to execute when using toolkits
- **Usage**: Required only when executing specific tools from toolkits
- **Format**: Exact tool name as defined in the toolkit
- **Examples**: `"calculator"`, `"web_search"`, `"code_interpreter"`

### Tool Execution Types

#### Individual Tool Execution
- **Usage**: Execute single tools directly connected to the node
- **Behavior**: Tool Name parameter is optional or ignored
- **Example**: Calculator tool for mathematical operations

#### Toolkit Execution
- **Usage**: Execute specific tools from connected toolkits
- **Behavior**: Tool Name parameter specifies which tool to run
- **Example**: Execute "search" tool from a web toolkit

## UseCases

- **Tool Development and Testing**: Test custom tools during development and validation
- **Automated Tool Execution**: Execute tools as part of automated workflows and processes
- **Tool Performance Testing**: Benchmark and performance test tool execution
- **Custom Integration Workflows**: Integrate tool execution into custom business workflows
- **Tool Documentation**: Generate documentation and examples for tool usage
- **Debugging and Troubleshooting**: Debug tool behavior and troubleshoot issues
- **Batch Data Processing**: Process multiple data items through tools efficiently
- **Tool Comparison**: Compare different tools for similar functionality
- **Workflow Automation**: Automate complex processes using multiple tools
- **API Testing**: Test tool-based API integrations and connections
- **Data Transformation**: Transform data using specialized tools and algorithms
- **Business Process Integration**: Integrate tools into existing business processes
- **Quality Assurance**: Validate tool outputs and ensure quality standards
- **Performance Monitoring**: Monitor tool performance and execution metrics
- **Tool Chain Orchestration**: Orchestrate complex tool chains and dependencies