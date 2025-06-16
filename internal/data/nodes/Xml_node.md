# XML Node

## Overview

Convert data from and to XML

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Data

#### Operation: JSON to XML

| Field Name | Type | Description | Required |
|---|---|---|---|
| mode | options | From and to what format the data should be converted | Yes |
| dataPropertyName | string | Name of the property to contain the converted XML data | Yes |
| allowSurrogateChars | boolean | Allow Unicode surrogate characters (Default: false) | No |
| attrkey | string | Prefix for accessing attributes (Default: $) | No |
| cdata | boolean | Wrap text nodes in CDATA instead of escaping (Default: false) | No |
| charkey | string | Prefix for accessing character content (Default: _) | No |
| headless | boolean | Omit the XML header (Default: false) | No |
| rootName | string | Root element name (Default: root) | No |

#### Operation: XML to JSON

| Field Name | Type | Description | Required |
|---|---|---|---|
| mode | options | From and to what format the data should be converted | Yes |
| dataPropertyName | string | Name of the property containing XML data to convert | Yes |
| attrkey | string | Prefix for accessing attributes (Default: $) | No |
| charkey | string | Prefix for accessing character content (Default: _) | No |
| explicitArray | boolean | Always put child nodes in array (Default: false) | No |
| explicitRoot | boolean | Include root node in resulting object (Default: true) | No |
| ignoreAttrs | boolean | Ignore XML attributes, create text nodes only (Default: false) | No |
| mergeAttrs | boolean | Merge attributes and child elements as parent properties (Default: true) | No |
| normalize | boolean | Trim whitespaces inside text nodes (Default: false) | No |
| normalizeTags | boolean | Normalize tag names to lowercase (Default: false) | No |
| trim | boolean | Trim whitespace at beginning and end of text nodes (Default: false) | No |

## UseCases

- [API Integration]: Convert XML responses from legacy APIs to JSON format for modern application processing and data manipulation
- [Data Migration]: Transform data between XML and JSON formats when migrating between systems or databases with different data format requirements
- [Configuration Management]: Convert configuration files between XML and JSON formats for different application environments and deployment scenarios

