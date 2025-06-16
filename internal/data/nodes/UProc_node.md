# uProc Node

## Overview

Consume uProc API

## Credentials

- Name: uprocApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Audio

#### Operation: Process Audio

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific audio processing tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Communication

#### Operation: Communication Tools

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific communication tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Company

#### Operation: Company Information

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific company information tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Finance

#### Operation: Financial Operations

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific financial tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Geographic

#### Operation: Geographic Data

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific geographic tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Image

#### Operation: Image Processing

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific image processing tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Internet

#### Operation: Internet Operations

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific internet tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Personal

#### Operation: Personal Data

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific personal data tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Product

#### Operation: Product Information

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific product information tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Security

#### Operation: Security Operations

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific security tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

### Resource: Text

#### Operation: Text Processing

| Field Name | Type | Description | Required |
|---|---|---|---|
| tool | options | Specific text processing tool | Yes |
| parameters | object | Tool-specific parameters | Varies |
| dataWebhook | string | URL to send tool response | No |

## UseCases

- API Integration: Connect to uProc services for various data processing tasks
- Multi-tool Processing: Access multiple processing tools through a single interface
- Automated Workflows: Build workflows using uProc's diverse tool ecosystem
- Data Transformation: Process and transform data using specialized tools

