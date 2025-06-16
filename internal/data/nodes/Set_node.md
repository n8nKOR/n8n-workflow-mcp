# Set Node

## Overview

Add or edit fields on an input item and optionally remove other fields

⚠️ **Version Information**: This node has multiple versions (V1, V2). The current default version is V2 (3.4), which provides enhanced field editing capabilities and improved data handling. Some features may differ between versions.

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Field Editor

#### Operation: Set Fields

| Field Name | Type | Description | Required |
|---|---|---|---|
| Mode | options | How to edit fields (Manual Mapping, JSON) | Yes |
| Fields to Set | collection | Define fields to add or modify manually | No |
| Include Other Input Fields | boolean | Whether to pass all input fields to output along with set fields | No |
| Input Fields to Include | options | How to select fields to include (All, Selected, All Except) | No |
| JSON Output | json | Customize item output with JSON specification (JSON mode only) | No |
| Duplicate Item | boolean | Whether to duplicate the item (node setting) | No |
| Duplicate Item Count | number | How many times to duplicate the item for testing | No |
| Fields to Include | string | Comma-separated list of field names to include | No |
| Fields to Exclude | string | Comma-separated list of field names to exclude | No |

### Mode Options
- **Manual Mapping**: Define fields individually with names and values
- **JSON**: Define output structure using JSON format

### Field Include Options
- **All**: Include all input fields in output
- **Selected**: Include only specified fields
- **All Except**: Include all fields except specified ones

### Manual Mapping Configuration
When using Manual Mapping mode, you can:
- Add new fields with custom names and values
- Modify existing field values
- Use expressions for dynamic field values
- Control which input fields to preserve in output

### JSON Mode Configuration
When using JSON mode, you can:
- Define complete output structure as JSON
- Use expressions within JSON structure
- Create complex nested objects
- Transform data structure completely

### Expression Support
All field values support n8n expressions:
```
Field Name: fullName
Field Value: {{ $json.firstName }} {{ $json.lastName }}

JSON Output: {
  "user": {
    "name": "{{ $json.name }}",
    "email": "{{ $json.email }}",
    "timestamp": "{{ $now.toISO() }}"
  }
}
```

## UseCases

- [Data Transformation and Cleanup] : Transform raw data from APIs or databases into structured formats, clean field names, and standardize data types for downstream processing
- [Email Content Simplification] : Extract and simplify email data by selecting only essential fields like date, subject, text, sender, and recipient for AI processing workflows
- [Workflow Configuration Variables] : Set configuration variables like API URLs, collection names, batch sizes, and embedding dimensions for use across multiple nodes in complex workflows
- [Slack Message Formatting] : Format Slack message data by extracting message ID, channel, user information, and content while cleaning up formatting for ticket creation systems
- [API Response Processing] : Process API responses by extracting relevant fields, renaming properties, and preparing data for integration with other services or databases
- [Meeting Data Extraction] : Extract and structure meeting information from calendar events, including attendee details, meeting URLs, and descriptions for automated meeting preparation workflows

