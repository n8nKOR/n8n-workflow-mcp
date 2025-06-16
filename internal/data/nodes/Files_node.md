# Read/Write Files from Disk Node

## Overview

Read or write files from the computer that runs n8n. This node provides direct interaction with the local file system for file processing operations. Essential for local file management, data processing, and workflow automation that requires file system access. Note that this node only works with files on the same computer running n8n - for handling files between different computers, use other nodes like FTP, HTTP Request, or AWS.

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: File Operations

#### Operation: Read File(s) From Disk

| Field Name | Type | Description | Required |
|---|---|---|---|
| File Path(s) | string/collection | Path(s) to the file(s) to read. Supports absolute paths or relative paths from n8n execution directory | Yes |
| Property Name | string | Name of the binary property to store the file data | No |

##### Read Options
| Field Name | Type | Description | Default | Required |
|---|---|---|---|---|
| Encoding | options | File encoding method (utf8, ascii, base64, etc.) | utf8 | No |
| Max File Size | number | Maximum file size limit in bytes | - | No |
| Include File Info | boolean | Whether to include file metadata information | false | No |

#### Operation: Write File to Disk

| Field Name | Type | Description | Required |
|---|---|---|---|
| File Path | string | Path where the file should be saved | Yes |
| Binary Data | string | Name of the binary property containing the data to write | Yes |

##### Write Options
| Field Name | Type | Description | Default | Required |
|---|---|---|---|---|
| Overwrite | boolean | Whether to overwrite existing files | false | No |
| Create Directory | boolean | Whether to automatically create directories if they don't exist | false | No |
| File Mode | string | File permissions in octal format (e.g., 0644) | - | No |

#### Additional Information
- Supports both absolute and relative file paths
- Handles multiple file formats and encodings
- Provides comprehensive error handling for file operations
- Maintains file metadata and permissions when specified
- Supports batch file reading operations
- Automatically handles directory creation when needed

## UseCases

- **Log File Processing**: Read system log files for analysis, monitoring, and automated alerting based on log patterns and error detection
- **Data Backup and Archiving**: Create local backup files from processed data, ensuring data persistence and recovery capabilities
- **Configuration File Management**: Read and modify application configuration files for dynamic system configuration and deployment automation
- **Report Generation**: Save processed data and generated reports to local files for distribution, archiving, or further processing
- **Temporary File Handling**: Create and manage temporary files during complex workflow processing for intermediate data storage
- **File Format Conversion**: Read files in one format and save them in another format after processing and transformation
- **Batch File Processing**: Process multiple files in sequence, reading, transforming, and saving results for bulk operations
- **System Integration**: Interface with local applications and systems that require file-based data exchange and communication
- **Development and Testing**: Create test files, read sample data, and manage development resources during workflow testing
- **Data Export and Import**: Export workflow results to files for external systems or import data from local files for processing

**Important Notes:**
- This node only works with files on the same computer running n8n
- For file operations between different computers, use FTP, HTTP Request, AWS, or similar nodes
- File paths can be absolute or relative to the n8n execution directory
- Ensure proper file permissions and security considerations when working with sensitive files 