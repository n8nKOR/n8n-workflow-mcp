# PostBin Node

## Overview

Simple webhook testing tool for creating temporary HTTP endpoints to capture and inspect incoming requests. PostBin provides basic bin management and request handling for development and testing purposes.

**Node Structure**: This node has two separate resources - "Bin" and "Request" - each with specific operations.

## Credentials

None required - PostBin is a free service that doesn't require authentication.

## Inputs

- **Main**: Input data for bin and request operations

## Outputs

- **Main**: Bin information, request data, or operation status

## Properties

### Resource: bin

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| - | - | Creates a new PostBin endpoint | - | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bin ID | string | ID of the PostBin to retrieve | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bin ID | string | ID of the PostBin to delete | Yes | - |

### Resource: request

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bin ID | string | ID of the PostBin to get requests from | Yes | - |

#### Operation: removeFirst

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bin ID | string | ID of the PostBin to remove first request from | Yes | - |

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bin ID | string | ID of the PostBin to send request to | Yes | - |
| Data | string | Data to send in the request | No | - |

## UseCases

- **Webhook Testing**: Test and debug webhook payloads from external services
- **API Development**: Inspect outgoing API requests during development
- **Integration Testing**: Test connections between services
- **Development Debugging**: Capture HTTP requests for analysis
- **Service Validation**: Verify third-party service integrations
- **Request Analysis**: Inspect headers and body content
- **Prototype Testing**: Quick testing during development phases

## Important Notes

- This is a basic testing tool with simple bin and request management
- Bins are temporary and may be automatically cleaned up
- Do not send sensitive data to PostBin endpoints
- Use for development and testing purposes only
- Request data is stored temporarily and may be publicly accessible

## Example Workflows

### Basic Webhook Testing
1. Create a new bin
2. Configure webhook service to use bin URL
3. Trigger webhook events
4. Retrieve and analyze received requests

### Request Debugging
1. Create bin for request capture
2. Send API requests to bin URL
3. Review captured request data
4. Debug and fix issues

