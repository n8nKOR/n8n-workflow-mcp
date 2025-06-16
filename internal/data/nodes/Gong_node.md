# Gong Node

## Overview

Interact with Gong API

## Credentials

- Name: gongApi, Required: Yes (for Access Token authentication)
- Name: gongOAuth2Api, Required: Yes (for OAuth2 authentication)

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Call

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | Authentication method (Access Token or OAuth2) | Yes |
| Resource | Options | Resource type to interact with | Yes |
| Call to Get | Resource Locator | Call to retrieve (From List, By ID, or By URL) | Yes |
| Options | Collection | Additional options for call data | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | Authentication method (Access Token or OAuth2) | Yes |
| Resource | Options | Resource type to interact with | Yes |
| Return All | Boolean | Whether to return all results or limit | No |
| Limit | Number | Maximum number of results to return | No |
| Options | Collection | Additional filtering and data options | No |

### Resource: User

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | Authentication method (Access Token or OAuth2) | Yes |
| Resource | Options | Resource type to interact with | Yes |
| User to Get | Resource Locator | User to retrieve (From List or By ID) | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | Authentication method (Access Token or OAuth2) | Yes |
| Resource | Options | Resource type to interact with | Yes |
| Return All | Boolean | Whether to return all results or limit | No |
| Limit | Number | Maximum number of results to return | No |


## UseCases

- **Sales Analytics** : Extract call data to analyze sales rep performance and conversation quality
- **Customer Intelligence** : Extract insights from customer calls and conversations for better understanding
- **CRM Integration** : Sync call data and recordings to CRM systems for comprehensive customer records
- **Sales Coaching** : Use call transcripts and recordings for targeted sales coaching and training
- **Compliance Monitoring** : Ensure sales calls meet regulatory and company standards
- **Revenue Attribution** : Track which calls and conversations lead to closed revenue

## Resources Available

### Call
Operations related to sales calls and conversations

**Operations:**
- **Get**: Retrieve specific call information
- **Get All**: List multiple calls with filtering options
- **Get Transcript**: Retrieve call transcript data
- **Get Recording**: Access call recording files

### User
Operations related to Gong users

**Operations:**
- **Get**: Retrieve specific user information
- **Get All**: List multiple users with filtering options

## List Search Methods

The node provides enhanced search capabilities:

### Get Calls
- **Searchable**: Filter calls by title or ID
- **Paginated**: Supports cursor-based pagination
- **Sorted**: Results sorted alphabetically by title

### Get Users
- **Searchable**: Filter users by name, email, or ID
- **Display Format**: Shows "First Last (email@domain.com)"
- **Paginated**: Supports cursor-based pagination
- **Sorted**: Results sorted alphabetically by name

## Base URL Configuration

The node automatically configures the base URL using:
```
baseURL: {{ $credentials.baseUrl.replace(new RegExp("/$"), "") }}
```

This removes trailing slashes from the configured base URL.

## API Integration

- **API Version**: Uses Gong API v2
- **Response Format**: JSON
- **Pagination**: Cursor-based pagination support
- **Filtering**: Built-in search and filter capabilities

## Usage Examples

### Retrieve Call Data
```
Resource: Call
Operation: Get
Call ID: <call_id>
```

### List Users
```
Resource: User  
Operation: Get All
```

### Search Calls
```
Resource: Call
Operation: Get All
Filter: <search_term>
```

## Error Handling

- **Authentication**: Validates credentials before API calls
- **API Limits**: Handles rate limiting and pagination
- **Network Errors**: Proper error handling for network issues

## Usage Notes
- Requires active Gong subscription and API access
- Base URL must be configured in credentials
- OAuth2 provides more secure authentication than access tokens
- Pagination tokens are automatically handled for large result sets
- Search filtering works on both names/titles and IDs
- Call transcripts and recordings may require additional permissions
- API responses include cursor information for pagination
- User search includes email addresses for better identification

## Related Nodes

- **HTTP Request**: For custom Gong API calls
- **JSON**: For processing Gong API responses
- **Code**: For custom call data analysis
- **Schedule Trigger**: For automated Gong data extraction

