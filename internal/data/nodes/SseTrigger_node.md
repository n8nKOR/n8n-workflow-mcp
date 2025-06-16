# SSE Trigger Node

## Overview

Triggers the workflow when Server-Sent Events occur

## Credentials

None

## Inputs

None

## Outputs

- Main

## Properties

### Resource: SSE

#### Operation: Listen
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| URL | string | The URL to receive the Server-Sent Events from | Yes | - |
| Headers | collection | Custom headers to send with the SSE request | No | - |
| Ignore SSL Issues | boolean | Whether to ignore SSL certificate issues | No | false |

## UseCases

- **Real-time Data**: Receive real-time data streams from web services
- **Live Updates**: Get live updates from applications and APIs
- **Event Streaming**: Process continuous event streams
- **Chat Applications**: Receive real-time chat messages
- **Stock Prices**: Monitor real-time stock price updates
- **IoT Data**: Receive continuous IoT sensor data streams

