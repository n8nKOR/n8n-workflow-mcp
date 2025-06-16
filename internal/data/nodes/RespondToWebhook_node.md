# Respond to Webhook Node

## Overview

The Respond to Webhook node sends HTTP responses back to webhook callers, enabling n8n workflows to function as API endpoints. This node supports multiple response formats including JSON, text, binary files, JWT tokens, and redirects, with customizable headers, status codes, and authentication options. It's essential for creating interactive webhooks that provide feedback to the calling system.

## Credentials

JWT authentication credentials are required only when using JWT response mode:
- **Credential Name**: `jwtAuth`
- **Key Types**: Passphrase (secret-based) or PEM Key (public/private key pair)
- **Algorithms**: HS256, HS384, HS512, RS256, RS384, RS512, ES256, ES384, ES512, PS256, PS384, PS512

## Inputs

- **Main**: Input data to be processed and returned in the webhook response

## Outputs

- **Main**: Original input data (passes through unchanged)
- **Response Output**: Response details including status code and headers (when enabled in v1.3+)

## Properties

### Resource: Webhook Response

#### Operation: Send Response

| Field Name | Type | Description | Required |
|---|---|---|---|
| Response Mode | Options | Type of response to send | Yes |
| Response Body | Conditional | Content to send (varies by mode) | Conditional |
| Response Code | Number | HTTP status code (100-599) | No |

### Response Modes

#### 1. All Incoming Items
| Field Name | Type | Description | Required |
|---|---|---|---|
| No additional fields | - | Returns all input data as JSON array | No |

**Use Case**: Return all workflow data to the webhook caller

#### 2. Binary File  
| Field Name | Type | Description | Required |
|---|---|---|---|
| Select Binary Property | Options | Binary field to return | Yes |

**Use Case**: Serve files, images, or documents through webhooks

#### 3. First Incoming Item
| Field Name | Type | Description | Required |
|---|---|---|---|
| No additional fields | - | Returns first input item as JSON | No |

**Use Case**: Return single record or result from workflow

#### 4. JSON
| Field Name | Type | Description | Required |
|---|---|---|---|
| JSON Body | JSON | Custom JSON response body | Yes |

**Use Case**: Return custom structured data or API responses

#### 5. JWT Token
| Field Name | Type | Description | Required |
|---|---|---|---|
| JWT Payload | JSON | JWT token payload data | Yes |

**Use Case**: Authentication workflows and session management

#### 6. No Data
| Field Name | Type | Description | Required |
|---|---|---|---|
| No additional fields | - | Returns empty response body | No |

**Use Case**: Acknowledgment responses and status confirmations

#### 7. Redirect
| Field Name | Type | Description | Required |
|---|---|---|---|
| Redirect URL | String | URL to redirect to | Yes |

**Use Case**: OAuth flows and form submission redirects

#### 8. Text
| Field Name | Type | Description | Required |
|---|---|---|---|
| Response Body | String | Plain text response content | Yes |

**Use Case**: Simple messages and plain text responses

## Advanced Configuration

### Custom Response Headers

| Field Name | Type | Description | Required |
|---|---|---|---|
| Response Headers | Collection | Custom HTTP headers | No |

#### Header Object Structure
| Property | Type | Description | Required |
|---|---|---|---|
| Name | String | HTTP header name | Yes |
| Value | String | HTTP header value | Yes |

#### Common Headers
```json
[
  {"name": "Content-Type", "value": "application/json"},
  {"name": "Cache-Control", "value": "no-cache"},
  {"name": "Access-Control-Allow-Origin", "value": "*"},
  {"name": "X-Custom-Header", "value": "custom-value"}
]
```

### HTTP Status Codes

#### Success Codes (2xx)
- **200**: OK - Standard success response
- **201**: Created - Resource created successfully
- **202**: Accepted - Request accepted for processing
- **204**: No Content - Success with no response body

#### Redirection Codes (3xx)
- **301**: Moved Permanently - Permanent redirect
- **302**: Found - Temporary redirect
- **304**: Not Modified - Cached version is current

#### Client Error Codes (4xx)
- **400**: Bad Request - Invalid request format
- **401**: Unauthorized - Authentication required
- **403**: Forbidden - Access denied
- **404**: Not Found - Resource not found
- **422**: Unprocessable Entity - Validation errors

#### Server Error Codes (5xx)
- **500**: Internal Server Error - Workflow error
- **502**: Bad Gateway - Upstream service error
- **503**: Service Unavailable - Service temporarily down

## JWT Token Configuration

### Supported Algorithms

#### HMAC Algorithms (Symmetric)
- **HS256**: HMAC using SHA-256 (default)
- **HS384**: HMAC using SHA-384
- **HS512**: HMAC using SHA-512

#### RSA Algorithms (Asymmetric)
- **RS256**: RSA signature with SHA-256
- **RS384**: RSA signature with SHA-384
- **RS512**: RSA signature with SHA-512

#### ECDSA Algorithms (Asymmetric)
- **ES256**: ECDSA using P-256 and SHA-256
- **ES384**: ECDSA using P-384 and SHA-384
- **ES512**: ECDSA using P-521 and SHA-512

#### RSASSA-PSS Algorithms
- **PS256**: RSASSA-PSS using SHA-256
- **PS384**: RSASSA-PSS using SHA-384
- **PS512**: RSASSA-PSS using SHA-512

### JWT Payload Examples

#### Basic User Token
```json
{
  "sub": "user123",
  "name": "John Doe",
  "iat": 1516239022,
  "exp": 1516242622
}
```

#### API Access Token
```json
{
  "iss": "api.example.com",
  "aud": "client-app",
  "scope": "read write",
  "exp": 1516242622
}
```

## Version Differences

### Version 1.0-1.2
- Single output (main data only)
- Basic response modes
- Limited expression support

### Version 1.3+
- Dual output support
- Response Output branch with status details
- Enhanced expression capabilities
- Better error handling

### Version 1.4+
- Optional response output configuration
- Improved binary data handling
- Enhanced JWT token support

## Compatible Nodes

The Respond to Webhook node works with these trigger nodes:

### Primary Compatibility
- **Webhook**: HTTP endpoint triggers
- **Form Trigger**: HTML form submissions
- **Chat Trigger**: Chat interface interactions
- **Wait**: Workflow pause with response capability

### Integration Requirements
- Must be connected to a compatible trigger node
- Trigger node must support response functionality
- Single response per workflow execution

## Response Output Details

When enabled, the Response Output provides metadata:

```json
{
  "statusCode": 200,
  "headers": {
    "Content-Type": "application/json",
    "X-Response-Time": "45ms"
  },
  "body": "Response content",
  "timestamp": "2024-01-15T10:30:00.000Z"
}
```

## Error Handling

### Validation Errors
- **Invalid JSON**: Malformed JSON in response body
- **Missing Binary Data**: Binary mode without binary input
- **Invalid JWT Payload**: Malformed JWT token data
- **Invalid URL**: Invalid redirect URL format

### Runtime Errors
- **JWT Signing Failed**: Invalid credentials or algorithm
- **Large Response**: Response size exceeds limits
- **Encoding Issues**: Character encoding problems

### Best Practices
- **Validate Input**: Ensure response data is properly formatted
- **Handle Errors**: Use Continue on Fail for robust error responses
- **Monitor Size**: Check response size limits
- **Test Tokens**: Validate JWT configuration before production

## Example Configurations

### API Response with Custom Headers
```json
{
  "responseMode": "json",
  "responseBody": {
    "status": "success",
    "data": "{{$json.result}}",
    "timestamp": "{{$now}}"
  },
  "responseCode": 200,
  "responseHeaders": [
    {"name": "Content-Type", "value": "application/json"},
    {"name": "X-API-Version", "value": "1.0"}
  ]
}
```

### File Download Response
```json
{
  "responseMode": "binaryFile",
  "selectBinaryProperty": "document",
  "responseCode": 200,
  "responseHeaders": [
    {"name": "Content-Disposition", "value": "attachment; filename=report.pdf"},
    {"name": "Content-Type", "value": "application/pdf"}
  ]
}
```

### Authentication Token Response
```json
{
  "responseMode": "jwtToken",
  "jwtPayload": {
    "userId": "{{$json.userId}}",
    "role": "{{$json.role}}",
    "exp": "{{$now.plus({hours: 24}).toSeconds()}}"
  },
  "responseCode": 200
}
```

### Error Response
```json
{
  "responseMode": "json",
  "responseBody": {
    "error": "Validation failed",
    "message": "{{$json.errorMessage}}",
    "code": "VALIDATION_ERROR"
  },
  "responseCode": 422
}
```

## UseCases

- **API Endpoint Creation**: Transform n8n workflows into REST API endpoints with proper responses
- **Form Processing**: Provide confirmation messages and redirects for web form submissions
- **Authentication Services**: Generate JWT tokens for user authentication and session management
- **File Serving**: Serve generated reports, images, or documents through webhook endpoints
- **Integration Callbacks**: Respond to webhook calls from third-party services with status updates
- **OAuth Implementations**: Handle OAuth flows with proper redirects and token responses
- **Webhook Acknowledgments**: Send confirmation responses to webhook senders
- **Error Handling**: Provide structured error responses for failed workflow operations
- **Real-time Notifications**: Send immediate responses to real-time event triggers
- **Data Validation**: Return validation results for submitted data
- **Payment Processing**: Respond to payment gateway webhooks with transaction status
- **Chat Bot Responses**: Provide responses in chat applications and messaging platforms
- **Workflow Status**: Return workflow execution status to external monitoring systems
- **Content Delivery**: Serve dynamic content based on webhook parameters
- **Redirect Services**: Implement URL shortening or redirect services through webhooks

