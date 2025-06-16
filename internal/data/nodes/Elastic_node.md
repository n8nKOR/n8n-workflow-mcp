# Elasticsearch Node

## Overview

The Elasticsearch node provides functionality to interact with Elasticsearch search engine for document and index management. This node enables you to leverage Elasticsearch's powerful capabilities for large-scale data search, analysis, and storage within your n8n workflows. Elasticsearch is a distributed, RESTful search and analytics engine capable of handling large volumes of data in real-time.

## Credentials

This node requires Elasticsearch API credentials:
- **Base URL**: The URL of your Elasticsearch instance (e.g., https://localhost:9200)
- **Username**: Username for authentication (optional, if security is enabled)
- **Password**: Password for authentication (optional, if security is enabled)

To configure credentials:
1. Ensure your Elasticsearch instance is running and accessible
2. If security is enabled, provide valid username and password
3. Test the connection to verify access to your Elasticsearch cluster

## Inputs

- **Main**: JSON data for document operations and index management

## Outputs

- **Main**: Returns JSON data from Elasticsearch API responses

## Properties

### Resource: Document

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index where the document will be created | Yes |
| Document ID | String | Unique ID of the document to create | No |
| Body | JSON | Document content in JSON format | Yes |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index containing the document | Yes |
| Document ID | String | ID of the document to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index containing the document | Yes |
| Document ID | String | ID of the document to retrieve | Yes |
| Simple | Boolean | Whether to return in simple format | No |

#### Operation: Get All
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index to search | Yes |
| Return All | Boolean | Whether to return all documents | No |
| Limit | Number | Maximum number of documents to return | No |
| Simple | Boolean | Whether to return in simple format | No |
| Query | JSON | Search query in JSON format | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index containing the document | Yes |
| Document ID | String | ID of the document to update | Yes |
| Body | JSON | Content to update in JSON format | Yes |

### Resource: Index

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index to create | Yes |
| Body | JSON | Index settings in JSON format | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index to retrieve | Yes |

#### Operation: Get All
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all indices | No |
| Limit | Number | Maximum number of indices to return | No |

#### Operation: Open
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index to open | Yes |

#### Operation: Close
| Field Name | Type | Description | Required |
|---|---|---|---|
| Index ID | String | ID of the index to close | Yes |

## UseCases

- **Log Data Analysis**: Store and analyze application logs in Elasticsearch for system monitoring and troubleshooting
- **Search Functionality Implementation**: Index data for advanced search features in websites and applications with relevance scoring
- **Real-time Data Processing**: Store incoming data in Elasticsearch for immediate search availability and real-time analytics
- **Business Intelligence**: Analyze large volumes of business data to derive insights and create dashboards
- **Security Monitoring**: Collect and analyze security event logs for threat detection and incident response
- **Performance Monitoring**: Store system performance metrics and analyze them for optimization opportunities
- **E-commerce Search**: Implement product search with faceting, filtering, and auto-complete functionality
- **Content Management**: Index and search documents, articles, and multimedia content with full-text search capabilities 