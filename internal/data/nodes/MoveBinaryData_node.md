# Convert to/from binary data Node

## Overview

Move data between binary and JSON properties

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Data

#### Operation: Binary to JSON
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Set All Data | boolean | Replace all JSON data with binary content | No | true |
| Source Key | string | Binary key to get data from (supports dot-notation) | Yes | data |
| Destination Key | string | JSON key to write data to (when Set All Data is false) | No | data |
| JSON Parse | boolean | Parse binary content as JSON | No | true |
| Encoding | options | Character encoding for text conversion | No | - |
| Strip BOM | boolean | Remove BOM marker from binary data | No | false |
| Keep Source | boolean | Preserve original data after conversion | No | false |

#### Operation: JSON to Binary
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Convert All Data | boolean | Convert all JSON data to binary | No | true |
| Source Key | string | JSON key to get data from (when not converting all) | No | data |
| Destination Key | string | Binary key to write data to (supports dot-notation) | Yes | data |
| Encoding | options | Character encoding for text conversion | No | - |
| Add Byte Order Mark (BOM) | boolean | Add BOM marker for compatible encodings | No | false |
| MIME Type | string | Set MIME type for binary data | No | - |
| Keep Source | boolean | Preserve original data after conversion | No | false |

## UseCases

- **File Processing**: Convert between binary and JSON data formats
- **Data Transformation**: Transform data between different property types
- **File Encoding**: Handle different text encodings and character sets
- **API Integration**: Convert binary data for API consumption
- **Data Migration**: Move data between binary and JSON formats
- **Content Processing**: Process file contents in different formats

