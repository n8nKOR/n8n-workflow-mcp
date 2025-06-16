# Read PDF Node

## Overview

Reads a PDF file and extracts its text content for further processing in workflows. This node can handle both encrypted and unencrypted PDF files, making it essential for document processing automation. It takes binary PDF data as input and outputs the extracted text content, enabling text analysis, data extraction, and content processing workflows.

## Credentials

None

## Inputs

- Main (with binary PDF data)

## Outputs

- Main (with extracted text content)

## Properties

### Resource: PDF Reader

#### Operation: Extract Text

| Field Name | Type | Description | Required |
|---|---|---|---|
| Input Binary Field | string | Name of the binary property from which to read the PDF file. Default is 'data' | Yes |
| Encrypted | boolean | Whether this is an encrypted/password-protected PDF file | Yes |
| Password | string | Password to decrypt the PDF file with. Only shown when Encrypted is true | No |

#### Additional Information
- The node preserves the original binary data in the output alongside the extracted text
- Supports password-protected PDF files with proper decryption
- Extracted text is returned in the JSON output for further processing
- Error handling allows workflows to continue on failure with error information
- The node is marked as hidden in the UI but fully functional

## UseCases

- **Email Attachment Processing**: Extract text content from PDF attachments received via Gmail for AI-powered analysis and classification before routing to appropriate team members
- **Document Content Analysis**: Convert PDF documents to text format for automated content analysis, keyword extraction, and information retrieval workflows
- **Invoice and Receipt Processing**: Extract text from financial documents like invoices, receipts, and statements for automated data entry and accounting system integration
- **Contract and Legal Document Analysis**: Process legal documents and contracts to extract key terms, dates, and clauses for compliance monitoring and contract management
- **Resume and CV Screening**: Extract text content from job applicant PDF resumes for AI-powered analysis and evaluation against job requirements and qualifications
- **Research Paper Processing**: Convert academic papers and research documents to text for content analysis, citation extraction, and knowledge management systems
- **Form Data Extraction**: Process filled PDF forms to extract submitted data for database entry and workflow automation
- **Compliance Document Review**: Extract text from regulatory documents and compliance reports for automated review and flagging of important information

