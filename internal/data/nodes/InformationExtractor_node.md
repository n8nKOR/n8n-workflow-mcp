# InformationExtractor

## Overview

The InformationExtractor node provides structured information extraction from unstructured text using AI language models within n8n's LangChain ecosystem. This node transforms raw text into organized, structured data by leveraging advanced natural language processing and schema-based extraction techniques.

## Credentials

The InformationExtractor node requires AI language model connection:

- **AI Language Model**: Connection to any supported language model node

## Inputs

- **Input Data**: Standard n8n workflow data containing text to process
- **AI Language Model**: Required connection to language model for processing

## Outputs

- **Extracted Data**: Structured JSON object containing extracted information

## Properties

### Resource: Information Extraction

#### Operation: Extract Information

| Field Name | Type | Description | Required |
|---|---|---|---|
| Text | String | The unstructured text to extract information from | Yes |
| Schema Definition Method | Options | Method to define extraction schema (Attributes, JSON Example, Manual Schema) | Yes |
| Attributes | Collection | List of attributes to extract (when using Attributes method) | No |
| JSON Example | JSON | Example JSON structure (when using JSON Example method) | No |
| JSON Schema | JSON | Manual JSON Schema definition (when using Manual Schema method) | No |
| System Prompt | String | Custom system prompt for extraction | No |
| Batch Processing | Boolean | Enable batch processing for multiple texts | No |
| Batch Size | Number | Number of items to process in each batch | No |
| Enable Output Fixing | Boolean | Enable automatic output correction | No |

## UseCases

- **Document Processing**: Extract structured data from contracts, invoices, and reports
- **Customer Data Extraction**: Parse customer information from emails and forms
- **Lead Generation**: Extract contact information from web content and documents
- **Content Analysis**: Structure data from articles, blogs, and research papers
- **E-commerce Data**: Extract product information from descriptions and reviews
- **Legal Document Processing**: Parse legal documents for key terms and clauses
- **Healthcare Records**: Extract patient information from medical documents
- **Financial Data Processing**: Parse financial statements and transaction records
- **HR Document Processing**: Extract candidate information from resumes and applications
- **Survey Response Analysis**: Structure open-ended survey responses 