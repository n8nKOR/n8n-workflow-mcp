# Webhook Node

## Overview

Creates HTTP endpoints that trigger workflow execution when called. Supports all standard HTTP methods and various data formats.

## Credentials

- httpBasicAuth (Basic Auth), Required: Yes when authentication is "basicAuth"
- httpHeaderAuth (Header Auth), Required: Yes when authentication is "headerAuth"  
- jwtAuth (JWT Auth), Required: Yes when authentication is "jwtAuth"

## Inputs

None

## Outputs

- Dynamic based on configuration

## Properties

### Resource: HTTP Endpoint

#### Operation: Listen for Requests

| Field Name | Type | Description | Required |
|---|---|---|---|
| Multiple Methods | boolean | Whether to allow multiple HTTP methods | No |
| HTTP Method | options/multiOptions | HTTP method(s) to listen to (GET, POST, PUT, DELETE, etc.) | Yes |
| Path | string | The path to listen to (supports dynamic values with ':') | Yes |
| Authentication | options | Authentication method (none, basicAuth, headerAuth, jwtAuth) | No |
| Response Mode | options | When and how to respond (onReceived, lastNode, responseNode) | No |
| Response Code | number | HTTP Response code to return (100-599) | No |
| Response Data | options | What data should be returned | No |
| Response Binary Property Name | string | Name of the binary property to return | No |
| Binary Data | boolean | Whether webhook receives binary data | No |
| Binary Property Name | string | Name of output field for binary data | No |
| Ignore Bots | boolean | Whether to ignore bot requests | No |
| IP Whitelist | string | Comma-separated list of allowed IP addresses | No |
| No Response Body | boolean | Whether to send any body in response | No |
| Raw Body | boolean | Whether to return the raw body | No |
| Response Content Type | string | Custom content-type header | No |
| Response Headers | fixedCollection | Add headers to webhook response | No |
| Response Property Name | string | Property name to return data from | No |

### HTTP Methods
- **GET**: Retrieve data from the endpoint
- **POST**: Send data to the endpoint
- **PUT**: Update data at the endpoint
- **DELETE**: Delete data at the endpoint
- **PATCH**: Partially update data
- **HEAD**: Get headers only
- **OPTIONS**: Get allowed methods

### Authentication Methods
- **None**: No authentication required
- **Basic Auth**: Username and password authentication
- **Header Auth**: Custom header-based authentication
- **JWT Auth**: JSON Web Token authentication

### Response Modes
- **On Received**: Respond immediately when request is received
- **Last Node**: Respond when workflow execution completes
- **Response Node**: Use a dedicated response node to control response

### Response Data Options
- **All Entries**: Return all workflow data
- **First Entry Only**: Return only the first data item
- **Last Entry Only**: Return only the last data item
- **No Data**: Return empty response

## UseCases

- [AI Slack Bot with Google Gemini] : Used as an endpoint to receive POST requests from Slack and trigger AI chatbot workflows. Receives Slack message data and forwards it to AI agents
- [Telegram Bot Integration] : Receive webhook events from Telegram Bot API to process incoming messages, including text, voice, and image content for AI-powered bot responses
- [Third-party Service Integration] : Create endpoints for external services to trigger workflows, enabling real-time data processing and automated responses to external events

