# Write Binary File Node

## Overview

Writes a binary file to disk

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: File

#### Operation: Write

| Field Name | Type | Description | Required |
|---|---|---|---|
| fileName | string | Path to which the file should be written | Yes |
| dataPropertyName | string | Name of the binary property which contains the data | Yes |
| append | boolean | Whether to append to an existing file | No |

## UseCases

- File Storage: Save binary data from previous nodes to disk
- Data Export: Export processed images, documents, or other binary files
- Backup Creation: Create backup copies of binary data
- File System Integration: Store files in specific locations for other applications

