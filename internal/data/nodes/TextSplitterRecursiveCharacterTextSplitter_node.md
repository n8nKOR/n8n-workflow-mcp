# TextSplitterRecursiveCharacterTextSplitter

## Overview

The TextSplitterRecursiveCharacterTextSplitter node provides advanced recursive text splitting capabilities within n8n's LangChain ecosystem, designed as the recommended solution for most text processing use cases. This intelligent text splitter employs a hierarchical approach, attempting to split text using a sequence of separators from most semantic (paragraphs) to least semantic (characters), ensuring optimal chunk boundaries while preserving document structure and context. It features specialized language-aware splitting for code and markup, configurable chunk sizes with overlap, and sophisticated boundary optimization, making it ideal for preparing diverse content types for vector storage, language model processing, and document analysis.

**Key Features:**
- Recursive hierarchical splitting with multiple separator attempts
- Language-aware splitting for 14+ programming languages and markup formats
- Intelligent separator sequence optimization (paragraphs → sentences → words → characters)
- Configurable chunk size and overlap with boundary optimization
- Preserves semantic structure and context across content types
- Optimized for vector storage and retrieval systems
- Memory-efficient processing with intelligent boundary detection
- Comprehensive support for code, documentation, and natural language

## Credentials

This node does not require any credentials as it operates as a text splitter provider within the LangChain ecosystem.

## Inputs

The TextSplitterRecursiveCharacterTextSplitter node does not accept direct inputs as it serves as a text splitter provider:

- **No Direct Inputs**: Text splitter nodes provide splitting capabilities to other components
- **Configuration-Based**: All configuration is done through node properties

## Outputs

The node outputs a text splitter connection:

- **AI Text Splitter**: LangChain recursive text splitter for document processing
- **Connection Type**: `NodeConnectionTypes.AiTextSplitter`
- **Usage**: Connect to document loaders, vector stores, or other components requiring intelligent text splitting

## Properties

### Resource: Text Splitter Configuration

#### Operation: Recursive Character Text Splitting

| Field Name | Type | Description | Required |
|---|---|---|---|
| Chunk Size | Number | Maximum size of each text chunk (default: 1000) | Yes |
| Chunk Overlap | Number | Number of characters to overlap between chunks (default: 0) | Yes |
| Split Code | Options | Language-specific code splitting mode (default: markdown) | No |

### Supported Languages for Split Code

| Language | Value | Description |
|---|---|---|
| C++ | cpp | Optimized for C++ code structure |
| Go | go | Optimized for Go code structure |
| Java | java | Optimized for Java code structure |
| JavaScript | js | Optimized for JavaScript code structure |
| PHP | php | Optimized for PHP code structure |
| Protocol Buffers | proto | Optimized for Protocol Buffer definitions |
| Python | python | Optimized for Python code structure |
| reStructuredText | rst | Optimized for reStructuredText documents |
| Ruby | ruby | Optimized for Ruby code structure |
| Rust | rust | Optimized for Rust code structure |
| Scala | scala | Optimized for Scala code structure |
| Swift | swift | Optimized for Swift code structure |
| Markdown | markdown | Optimized for Markdown documents |
| LaTeX | latex | Optimized for LaTeX documents |
| HTML | html | Optimized for HTML documents |

## UseCases

- **Document Processing**: Split large documents into manageable chunks for vector storage and retrieval
- **Code Analysis**: Process source code files with language-aware splitting for better context preservation
- **Content Preparation**: Prepare diverse content types for language model processing and analysis
- **Vector Database Optimization**: Create optimally sized chunks for embedding models and similarity search
- **Multi-Language Support**: Handle mixed content types with intelligent boundary detection