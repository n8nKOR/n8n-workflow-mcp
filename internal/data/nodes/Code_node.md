# Code

## Overview

The Code node (LangChain) provides a secure JavaScript execution environment with full access to the LangChain library ecosystem within n8n workflows. This powerful node enables custom LangChain component creation, advanced data processing, and dynamic workflow logic using JavaScript. It operates in two primary modes: Execute mode for running code with inputs/outputs, and Supply Data mode for creating reusable LangChain tools and components.

## Credentials

The Code node does not require direct credentials. However, it inherits security permissions and access controls from the n8n environment:

- **Environment Variables**: Access to configured environment variables
- **Module Permissions**: Controlled access to external modules
- **LangChain Components**: Full access to LangChain ecosystem
- **n8n Context**: Access to workflow context and data

## Inputs

The Code node supports dynamic input configuration:

### Static Inputs
- **JavaScript Code**: The code to execute within the LangChain environment

### Dynamic Inputs
Configurable based on code requirements:

| Input Type | Description | Required |
|---|---|---|
| Language Model | Connection to various LLM providers | No |
| Chat Model | Chat-capable language models | No |
| Memory | Conversation memory systems | No |
| Tool | Function tools and utilities | No |
| Embeddings | Text embedding providers | No |
| Vector Store | Vector database connections | No |
| Document Loader | Document ingestion components | No |
| Text Splitter | Content chunking utilities | No |
| Output Parser | Response formatting tools | No |
| Retriever | Information retrieval systems | No |

## Outputs

Output configuration depends on the selected mode:

### Execute Mode
- **Result**: Primary execution result
- **Additional Outputs**: User-defined output ports
- **Error Information**: Execution errors and debugging data

### Supply Data Mode  
- **LangChain Component**: Created tool, chain, or component
- **Component Metadata**: Type, description, and usage information

## Properties

### Resource: Code

#### Operation: Execute

| Field Name | Type | Description | Required |
|---|---|---|---|
| Mode | Options | Execution mode (Execute) | Yes |
| Input Configuration | Array | Dynamic input connectors | No |
| Output Configuration | Array | Dynamic output connectors | No |
| JavaScript Code | String | The code to execute within the LangChain environment | Yes |

#### Operation: Supply Data

| Field Name | Type | Description | Required |
|---|---|---|---|
| Mode | Options | Execution mode (Supply Data) | Yes |
| JavaScript Code | String | The code to create LangChain components | Yes |
| Component Type | Options | Type of LangChain component to create | No |

## UseCases

- **Custom Tool Development** : Create specialized tools for specific business needs
- **Data Transformation** : Complex data processing and transformation logic
- **API Integration** : Custom API integrations with advanced error handling
- **Chain Composition** : Build complex LangChain workflows programmatically
- **Memory Management** : Custom memory implementations and session handling
- **Vector Operations** : Advanced vector database operations and similarity searches
- **Content Processing** : Advanced text processing and content analysis
- **Workflow Orchestration** : Dynamic workflow control and conditional logic
- **Protocol Integration** : Custom protocol handlers and data format converters
- **AI Model Chaining** : Complex model orchestration and result aggregation
- **Real-time Analytics** : Live data processing and metric calculation
- **Dynamic Prompt Generation** : Context-aware prompt creation and optimization
- **Error Recovery** : Advanced error handling and workflow recovery mechanisms
- **Performance Optimization** : Custom caching and optimization strategies
- **Security Implementation** : Custom security validation and sanitization
- **Testing and Validation** : Automated testing tools for LangChain workflows

