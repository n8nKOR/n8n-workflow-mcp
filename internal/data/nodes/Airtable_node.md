# Airtable Node

## Overview

The Airtable node allows you to interact with Airtable, a cloud-based database service that combines the functionality of a database with the simplicity of a spreadsheet. This node enables you to create, read, update, and delete records in your Airtable bases through n8n workflows.

## Credentials

This node requires Airtable API credentials:
- **API Token**: Your Airtable personal access token
- **Base ID** (optional): Can be specified in the node configuration

To obtain API credentials:
1. Log into your Airtable account
2. Visit https://airtable.com/create/tokens
3. Create a new personal access token with appropriate scopes
4. Copy the token for use in n8n

## Inputs

- **Main**: JSON data for creating/updating records

## Outputs

- **Main**: Returns JSON data from Airtable API responses

## Properties

### Resource: Record

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | ID of the Airtable base | Yes |
| Table | String | Name or ID of the table | Yes |
| Columns | Fixed Collection | Field values for the new record | Yes |
| Options | Collection | Additional options (typecast, etc.) | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | ID of the Airtable base | Yes |
| Table | String | Name or ID of the table | Yes |
| Record ID | String | ID of the record to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | ID of the Airtable base | Yes |
| Table | String | Name or ID of the table | Yes |
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Download Attachments | Boolean | Whether to download attachment files | No |
| Additional Options | Collection | Filters, sorting, view options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | ID of the Airtable base | Yes |
| Table | String | Name or ID of the table | Yes |
| Record ID | String | ID of the record to update | Yes |
| Columns | Fixed Collection | Field values to update | Yes |
| Options | Collection | Additional options (typecast, etc.) | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | ID of the Airtable base | Yes |
| Table | String | Name or ID of the table | Yes |
| Record ID | String | ID of the record to delete | Yes |

#### Operation: Append
| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | ID of the Airtable base | Yes |
| Table | String | Name or ID of the table | Yes |
| Columns | Fixed Collection | Field values for the new records | Yes |
| Options | Collection | Additional options (typecast, etc.) | No |

#### Operation: List
| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | ID of the Airtable base | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |
| Additional Options | Collection | Additional filtering options | No |

### Resource: Base

#### Operation: Get Schema
| Field Name | Type | Description | Required |
|---|---|---|---|
| Base ID | String | ID of the Airtable base | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Max number of results to return | No |

## UseCases

- **Multi-user Routing Management** : Store and manage user-specific routing information (email addresses, Notion database IDs, tokens) for Gmail-to-Notion task conversion systems, enabling proper task creation in the correct user databases
- **Route Status Management** : Control email processing by activating/deactivating user routing settings, automatically disabling routes under specific conditions to prevent errors and ensure system stability
- **HR Applicant Tracking System** : Store and manage job applicant information including personal details, contact information, CV links, and application status for automated recruitment workflows
- **AI-Enhanced Data Processing** : Automatically trigger workflows when new records are created or updated, enabling dynamic data extraction and AI-powered analysis of unstructured content
- **UTM Campaign Management** : Store and organize marketing campaign UTM links with associated QR codes, tracking campaign performance and managing marketing attribution data
- **Marketing Analytics Dashboard** : Maintain campaign performance metrics, click-through rates, and conversion data for comprehensive marketing analysis and reporting

