# ChainRetrievalQa

## Overview

Question-answering with document retrieval using Retrieval-Augmented Generation (RAG) workflows

## Credentials

- Name: Language Model Credentials, Required: Yes (inherited from connected LLM node)
- Name: Retriever Credentials, Required: Yes (inherited from connected retriever)

## Inputs

- AI Language Model
- AI Retriever
- Main (optional)

## Outputs

- Main

## Properties

### Resource: Question Answering

#### Operation: Retrieval QA

| Field Name | Type | Description | Required |
|---|---|---|---|
| Query | String | Question or query to answer | Yes |
| Prompt Type | Options | Type of prompt handling | No |
| Text | String | User message content | Conditional |
| Options | Collection | Additional configuration options | No |

## UseCases

- **Knowledge Base Q&A** : Answer questions about company documentation, policies, and procedures
- **Customer Support** : Provide automated responses based on support documentation
- **Research Assistance** : Help researchers find answers in large document collections
- **Educational Tools** : Create interactive learning systems with course materials
- **Legal Document Analysis** : Answer questions about legal documents and contracts
- **Technical Documentation** : Provide answers from API docs, user manuals, and guides
- **Medical Information Systems** : Answer medical queries based on clinical documentation
- **Financial Analysis** : Answer questions about financial reports and market data
- **Product Information** : Provide product details from catalogs and specifications
- **Compliance Assistance** : Answer regulatory and compliance questions
- **Academic Research** : Support academic research with paper and thesis analysis
- **Content Discovery** : Help users find relevant information in large content libraries
- **Troubleshooting Systems** : Provide technical support based on troubleshooting guides
- **Policy Consultation** : Answer questions about organizational policies and procedures
- **Historical Research** : Query historical documents and archives
- **Scientific Literature Review** : Answer questions from scientific papers and journals