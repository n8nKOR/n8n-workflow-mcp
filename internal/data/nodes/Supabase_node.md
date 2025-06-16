# Supabase Node

## Overview

Add, get, delete and update data in a table

## Credentials

- Name: supabaseApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: row

#### Operation: create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Use Custom Schema | boolean | Whether to use a database schema different from the default "public" schema | No |
| Schema | string | Name of database schema to use for table | No |
| Table Name or ID | options | Choose from the list, or specify an ID using an expression | Yes |
| Data to Send | options | Auto-Map Input Data to Columns or Define Below for Each Column | No |
| Inputs to Ignore | string | List of input properties to avoid sending, separated by commas | No |
| Fields to Send | fixedCollection | Field values to send when defining below | No |

#### Operation: delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Use Custom Schema | boolean | Whether to use a database schema different from the default "public" schema | No |
| Schema | string | Name of database schema to use for table | No |
| Table Name or ID | options | Choose from the list, or specify an ID using an expression | Yes |
| Select Conditions | fixedCollection | Conditions to filter rows for deletion | No |

#### Operation: get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Use Custom Schema | boolean | Whether to use a database schema different from the default "public" schema | No |
| Schema | string | Name of database schema to use for table | No |
| Table Name or ID | options | Choose from the list, or specify an ID using an expression | Yes |
| Select Conditions | fixedCollection | Conditions to filter rows for retrieval | No |

#### Operation: getAll

| Field Name | Type | Description | Required |
|---|---|---|---|
| Use Custom Schema | boolean | Whether to use a database schema different from the default "public" schema | No |
| Schema | string | Name of database schema to use for table | No |
| Table Name or ID | options | Choose from the list, or specify an ID using an expression | Yes |
| Return All | boolean | Whether to return all results or only up to a given limit | No |
| Limit | number | Max number of results to return | No |
| Filters | fixedCollection | Conditions to filter rows | No |

#### Operation: update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Use Custom Schema | boolean | Whether to use a database schema different from the default "public" schema | No |
| Schema | string | Name of database schema to use for table | No |
| Table Name or ID | options | Choose from the list, or specify an ID using an expression | Yes |
| Select Conditions | fixedCollection | Conditions to filter rows for update | No |
| Data to Send | options | Auto-Map Input Data to Columns or Define Below for Each Column | No |
| Inputs to Ignore | string | List of input properties to avoid sending, separated by commas | No |
| Fields to Send | fixedCollection | Field values to send when defining below | No |

## UseCases

- [AI Agent File Management] : Store and manage file metadata for documents uploaded to Supabase Storage, enabling AI agents to track processed files and prevent duplicate processing in RAG systems
- [Vector Database Integration] : Use Supabase as a vector store for AI embeddings, storing document chunks and their vector representations for semantic search and retrieval-augmented generation workflows
- [Telegram Bot Memory Storage] : Store long-term conversation memory and user preferences for Telegram bots, maintaining persistent context across chat sessions with OpenAI assistant integration
- [User Session Management] : Track user interactions, conversation history, and personalized settings for AI chatbots and virtual assistants across multiple platforms
- [Document Processing Pipeline] : Manage document processing workflows by storing file processing status, extracted content, and metadata for automated content analysis systems
- [Real-time Data Synchronization] : Sync data between different applications and services using Supabase's real-time capabilities for live updates and collaborative workflows

