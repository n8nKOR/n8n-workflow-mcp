# Rename Keys Node

## Overview

Update item field names

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Keys

#### Operation: Rename
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Keys | collection | Define keys to be renamed | No | - |
| Current Key Name | string | The current name of the key (supports dot-notation for deep keys) | Yes | - |
| New Key Name | string | The name the key should be renamed to (supports dot-notation) | Yes | - |
| Regex | collection | Regular expression replacements for key names | No | - |
| Regular Expression | string | Regex pattern to match key names | Yes | - |
| Replace With | string | The replacement pattern for matched keys (supports regex captures $1, $2, etc.) | Yes | - |
| Case Insensitive | boolean | Whether to use case insensitive matching | No | false |
| Max Depth | number | Maximum depth to replace keys (-1 for unlimited, 0 for top level only) | No | -1 |

## UseCases

- **Data Transformation**: Rename JSON keys to match target schema
- **API Integration**: Transform field names between different APIs
- **Data Cleaning**: Standardize field naming conventions
- **Legacy System Integration**: Map old field names to new standards
- **Bulk Renaming**: Use regex patterns for bulk key renaming

