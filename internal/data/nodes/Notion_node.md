# Notion Node (V2)

## Overview

Consume the Notion API to interact with Notion workspaces, databases, pages, and blocks for content automation and data synchronization. This node uses Notion API v2 with support for multiple versions (2, 2.1, 2.2) and provides advanced resource locator functionality for flexible resource referencing.

## Credentials

This node requires Notion API credentials:
- **Name**: notionApi
- **Required**: Yes
- **Type**: Internal Integration Token for Notion integrations with specific permissions

### Credential Configuration
To obtain API credentials:
1. Go to [Notion Integrations](https://www.notion.so/my-integrations)
2. Click **New Integration**
3. Give your integration a name and select associated workspace
4. Configure capabilities (Read content, Update content, Insert content)
5. Copy the **Internal Integration Token**
6. **Important**: Share pages/databases with your integration in Notion
7. Add your connection to integrate with Notion workspace

## Inputs

- **Main**: JSON data for creating/updating Notion content

## Outputs

- **Main**: Returns JSON data from Notion API responses

## Properties

### Resource: Block

#### Operation: Append
| Field Name | Type | Description | Required |
|---|---|---|---|
| Block ID | resourceLocator | Block to append to (supports URL/ID modes with validation) | Yes |
| Block Content | Collection | Blocks to append | Yes |

**Block ID Resource Locator:**
- **By URL**: Extract block ID from Notion URL automatically
- **By ID**: Directly specify the block ID (format: 1234-5678-9abc-def0)
- **From List**: Choose from available blocks
- **Validation**: Automatic URL parsing and ID format validation

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Block ID | resourceLocator | Parent block to get children from (supports URL/ID modes) | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |

### Resource: Database

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | resourceLocator | Database to retrieve (supports URL/ID modes with validation) | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Search
| Field Name | Type | Description | Required |
|---|---|---|---|
| Search Text | String | Text to search for | Yes |
| Sort | Collection | Sorting options | No |
| Filter | Collection | Filtering options | No |

### Resource: Database Page

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | resourceLocator | Database to create page in (supports URL/ID modes) | Yes |
| Title | String | Title of the page | Yes |
| Properties | Collection | Database properties to set | No |
| Page Content | Collection | Blocks to add to page | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Page ID | resourceLocator | Page to retrieve (supports URL/ID modes with validation) | Yes |
| Property Names | Array | Specific properties to retrieve | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | resourceLocator | Database to query (supports URL/ID modes) | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |
| Filter | Collection | Filtering criteria | No |
| Sort | Collection | Sorting options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Page ID | resourceLocator | Page to update (supports URL/ID modes with validation) | Yes |
| Properties | Collection | Properties to update | No |

### Resource: Page

#### Operation: Archive
| Field Name | Type | Description | Required |
|---|---|---|---|
| Page ID | resourceLocator | Page to archive (supports URL/ID modes with validation) | Yes |

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Parent | resourceLocator | Parent page or database (supports URL/ID modes) | Yes |
| Title | String | Title of the page | Yes |
| Page Content | Collection | Blocks to add to page | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Page ID | resourceLocator | Page to retrieve (supports URL/ID modes with validation) | Yes |

#### Operation: Search
| Field Name | Type | Description | Required |
|---|---|---|---|
| Search Text | String | Text to search for | Yes |
| Filter | Collection | Filtering options | No |
| Sort | Collection | Sorting options | No |

### Resource: User

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| User ID | resourceLocator | User to retrieve (supports URL/ID modes) | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |

## Resource Locator System

The `resourceLocator` type is an advanced field type that provides multiple ways to reference Notion resources:

### Input Methods
- **Select from List**: Choose from a dropdown of available resources
- **By ID**: Directly specify the resource ID with format validation
- **By URL**: Use a Notion URL to automatically extract the ID

### URL Processing
- **Automatic Extraction**: Extracts resource IDs from Notion URLs
- **Format Validation**: Validates Notion URL structure and format
- **Error Handling**: Provides clear error messages for invalid URLs
- **Link Copying**: Instructions for copying links from Notion interface

### ID Validation
- **Format Checking**: Validates Notion ID format (UUID with dashes)
- **Regex Patterns**: Uses validation regex patterns for URL and ID formats
- **Placeholder Guidance**: Field placeholders show expected formats

## Version Support

### Node Versions
- **Version 2**: Base functionality with core Notion API features
- **Version 2.1**: Enhanced field validation and improved error handling
- **Version 2.2**: Latest version with full resource locator support (default)

### Version-Specific Features
- **Field Validation**: Enhanced validation rules in newer versions
- **URL Extraction**: Improved URL parsing in version 2.1+
- **Resource Locator**: Advanced locator functionality in version 2.2
- **Error Messages**: More detailed error messages in recent versions

### Default Version
- **Current Default**: Version 2.2
- **Recommendation**: Use latest version for new workflows
- **Compatibility**: Older versions maintained for compatibility

## Authentication Integration

### Connection Notice
- **Setup Requirement**: Integration must be connected to Notion workspace
- **Permission Configuration**: Integration needs appropriate workspace permissions
- **Sharing Requirement**: Pages/databases must be shared with integration
- **Connection Validation**: Node validates connection before operations

### Security Features
- **Token Validation**: Validates Internal Integration Token format
- **Permission Checking**: Verifies integration has required permissions
- **Workspace Access**: Ensures integration has workspace access
- **Error Handling**: Clear error messages for authentication issues

## Technical Details

### API Integration
- **Notion API Version**: Uses Notion API v1 endpoints
- **Rate Limiting**: Handles Notion API rate limits automatically
- **Error Handling**: Comprehensive error handling for API responses
- **Response Processing**: Structured JSON response processing

### Field Processing
- **Dynamic Validation**: Context-aware field validation
- **Resource Dependencies**: Fields update based on resource selections
- **Format Conversion**: Automatic format conversion between URL and ID
- **Placeholder Updates**: Dynamic placeholder text based on field context

## UseCases

- **Notion Knowledge Base AI Assistant**: Implement a question-answer AI assistant using Notion databases as knowledge bases. Search related information by keywords or tags, maintain conversation history, and provide contextually appropriate responses
- **AI Assistant Workflow Generator**: Automatically generate custom n8n workflows for Notion databases by analyzing database schemas and creating tailored AI assistant configurations with appropriate search queries and response formats
- **Audio Transcription and Summarization**: Transcribe audio files using AI, generate summaries with GPT-4, and automatically store structured content in Notion databases for knowledge management and content organization
- **Vector Database Integration**: Convert Notion pages into vector embeddings and store them in vector databases like Pinecone for semantic search and retrieval-augmented generation (RAG) applications
- **Research Paper Analysis**: Automatically analyze academic papers from sources like Hugging Face, extract key insights using AI, and organize findings in structured Notion databases for research management
- **Dynamic Content Management**: Create and update Notion pages programmatically based on external data sources, enabling automated content creation and database population for content management systems

