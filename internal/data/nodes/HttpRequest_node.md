# HTTP Request

## Overview

The HTTP Request node is a versatile and powerful node that enables communication with external APIs and web services by sending HTTP requests and processing responses. It supports all standard HTTP methods (GET, POST, PUT, DELETE, PATCH, etc.) and provides comprehensive configuration options for headers, query parameters, request bodies, and authentication methods. This node is essential for integrating n8n workflows with REST APIs, webhooks, and web services.

## Credentials

The HTTP Request node supports multiple authentication methods to accommodate various API security requirements:

| Authentication Type | Description | Use Case |
|---|---|---|
| None | No authentication required | Public APIs, internal services |
| Basic Auth | Username/password-based authentication | Traditional web services |
| Bearer Token | Token-based authentication | Modern APIs, JWT tokens |
| Digest Auth | MD5 hash-based authentication | Legacy systems requiring digest auth |
| Header Auth | Custom header-based authentication | APIs with custom auth headers |
| Query Auth | Query parameter-based authentication | APIs requiring auth in URL parameters |
| OAuth1 | OAuth 1.0 authentication | Twitter API, legacy OAuth services |
| OAuth2 | OAuth 2.0 authentication | Google APIs, modern OAuth services |
| SSL Certificates | Client certificate authentication | Enterprise APIs requiring mutual TLS |

## Inputs

The node accepts input data from previous nodes in the workflow, which can be used to:
- Dynamically construct request URLs
- Populate request headers and parameters
- Generate request body content
- Control authentication credentials

## Outputs

The HTTP Request node outputs the response data in JSON format, including:
- **Response Body**: The main response content (JSON, text, binary)
- **Status Code**: HTTP status code (200, 404, 500, etc.)
- **Headers**: Response headers from the server
- **Status Message**: HTTP status message

## Properties

### Resource: HTTP Request

#### Operation: HTTP Request

| Field Name | Type | Description | Required |
|---|---|---|---|
| Method | Options | HTTP method (GET, POST, PUT, DELETE, PATCH, etc.) | Yes |
| URL | String | The URL to make the request to | Yes |
| Authentication | Options | Authentication method to use | Yes |
| Send Query | Boolean | Whether to send query parameters | No |
| Send Headers | Boolean | Whether to send custom headers | No |
| Send Body | Boolean | Whether to send request body | No |

#### Query Parameters

| Field Name | Type | Description | Required |
|---|---|---|---|
| Query Parameters | Collection | List of query parameters to send | No |

#### Headers

| Field Name | Type | Description | Required |
|---|---|---|---|
| Headers | Collection | List of custom headers to send | No |

#### Body

| Field Name | Type | Description | Required |
|---|---|---|---|
| Body Content Type | Options | Type of body content (JSON, Form Data, Binary, etc.) | No |
| Body | String/JSON | Request body content | No |

#### Options

| Field Name | Type | Description | Required |
|---|---|---|---|
| Response | Options | Response format options | No |
| Pagination | Collection | Pagination configuration | No |
| SSL Certificates | Boolean | Whether to provide SSL certificates | No |
| Proxy | String | HTTP proxy configuration | No |
| Timeout | Number | Request timeout in milliseconds | No |
## UseCases

- **HttpRequest Integration** : Use HttpRequest for workflow automation and data processing
- **Data Processing** : Process and transform data using HttpRequest capabilities
- **Workflow Enhancement** : Enhance workflows with HttpRequest functionality
- **Automation Tasks** : Automate repetitive tasks using HttpRequest features

## Pagination Support

The HTTP Request node includes built-in pagination support for APIs that return large datasets:
- **Parameter-based Pagination**: Updates parameters for subsequent requests
- **URL-based Pagination**: Follows next page URLs from responses
- **Cursor-based Pagination**: Uses cursors or tokens for pagination
- **Automatic Stopping**: Configurable conditions to stop pagination

## Error Handling

The node implements robust error handling:
- **HTTP Status Codes**: Proper handling of 4xx and 5xx errors
- **Network Errors**: Timeout and connection error management
- **SSL/TLS Errors**: Certificate validation and security errors
- **Authentication Errors**: Clear reporting of auth failures
- **Rate Limiting**: Automatic retry with exponential backoff

## Performance Features

- **Connection Pooling**: Efficient connection reuse for multiple requests
- **Compression**: Automatic gzip/deflate compression support
- **Streaming**: Support for large file uploads and downloads
- **Batch Processing**: Efficient handling of multiple requests
- **Caching**: Optional response caching for improved performance

## Use Cases

- **API Integration**: Connect with REST APIs to fetch or send data
- **Webhook Delivery**: Send data to external systems via webhooks
- **Data Synchronization**: Synchronize data between different systems and platforms
- **File Operations**: Upload files to cloud storage or download resources
- **Monitoring and Health Checks**: Monitor website availability and API status
- **Authentication Flows**: Handle OAuth flows and token refresh operations
- **Microservices Communication**: Enable communication between microservices
- **Third-party Service Integration**: Connect with SaaS platforms and external services
- **Data Collection**: Gather data from multiple APIs for analysis and reporting
- **Notification Systems**: Send alerts and notifications to external systems
- **Content Management**: Interact with CMS APIs for content operations
- **E-commerce Integration**: Connect with payment gateways and shopping platforms
- **Social Media Automation**: Interact with social media APIs for posting and monitoring
- **IoT Device Communication**: Send commands and receive data from IoT devices 