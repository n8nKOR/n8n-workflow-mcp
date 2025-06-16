# ChainSummarization

## Overview

Text summarization using AI language models with multiple processing methods

## Credentials

- Name: Language Model Connection, Required: Yes

## Inputs

- Main
- Model
- Document (optional)

## Outputs

- Main

## Properties

### Resource: Text Summarization

#### Operation: Document Summarization

| Field Name | Type | Description | Required |
|---|---|---|---|
| Summarization Method | Options | Choose between 'Map Reduce' or 'Refine' summarization methods | Yes |
| Prompt Template | String | Custom prompt template for summarization | No |
| Input Text Field | String | Field name containing the text to summarize | No |
| Chunk Size | Number | Maximum size of text chunks for processing | No |
| Chunk Overlap | Number | Number of characters to overlap between chunks | No |

## UseCases

- **Document Summarization** : Generate concise summaries of long documents and articles
- **Content Analysis** : Extract key insights from large volumes of text data
- **Research Assistance** : Summarize research papers and academic documents
- **Business Intelligence** : Create executive summaries from business reports
- **Knowledge Management** : Process and summarize organizational knowledge bases
- **Email Processing** : Summarize email threads and conversations
- **News Aggregation** : Generate news summaries from articles and feeds
- **Meeting Notes** : Create concise summaries from meeting transcripts
- **Legal Document Review** : Create briefs from legal documents and contracts
- **Customer Feedback Analysis** : Summarize customer reviews and feedback
- **Technical Documentation** : Create overview summaries of technical manuals
- **Financial Report Summary** : Condense financial reports for executive review
- **Social Media Monitoring** : Summarize social media conversations and trends
- **Content Migration** : Summarize legacy content during system migrations
- **Compliance Documentation** : Summarize regulatory documents and requirements
- **Market Research** : Create summaries from market research reports