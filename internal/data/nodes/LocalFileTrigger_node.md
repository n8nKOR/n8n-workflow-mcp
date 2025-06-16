# Local File Trigger Node

## Overview

Triggers a workflow on file system changes using the Chokidar file watching library. This trigger node monitors files and folders for various events like creation, modification, and deletion. Essential for automating file processing workflows, monitoring log files, and creating reactive systems that respond to file system changes. Supports both file and folder monitoring with extensive configuration options.

## Credentials

None

## Inputs

None (Trigger Node)

## Outputs

- Main

## Properties

### Resource: File System Monitoring

#### Operation: Watch Files

| Field Name | Type | Description | Required |
|---|---|---|---|
| Trigger On | options | Choose between monitoring a specific file or folder | Yes |
| File to Watch | string | The specific file path to monitor (when Trigger On = file) | Conditional |
| Folder to Watch | string | The directory path to monitor (when Trigger On = folder) | Conditional |
| Watch for | multiOptions | Types of events to monitor when watching folders | Conditional |

#### Watch for Events (Folder Monitoring)
| Event Type | Value | Description |
|---|---|---|
| File Added | add | Triggers whenever a new file is added to the folder |
| File Changed | change | Triggers whenever an existing file is modified |
| File Deleted | unlink | Triggers whenever a file is deleted from the folder |
| Folder Added | addDir | Triggers whenever a new subfolder is created |
| Folder Deleted | unlinkDir | Triggers whenever a subfolder is deleted |

### Options

| Field Name | Type | Description | Default | Required |
|---|---|---|---|---|
| Await Write Finish | boolean | Wait until files finish writing to avoid reading partial content | false | No |
| Include Linked Files/Folders | boolean | Whether to follow symlinks, aliases, and shortcuts | true | No |
| Ignore | string | Files or paths to ignore using Anymatch syntax (e.g., **/*.txt) | '' | No |
| Ignore Existing Files/Folders | boolean | Whether to ignore existing files/folders on startup | true | No |
| Max Folder Depth | options | How deep into folder structure to watch (0=top only, -1=unlimited) | -1 | No |
| Use Polling | boolean | Use polling instead of native file system events | false | No |

#### Additional Information
- Uses the Chokidar library for reliable cross-platform file watching
- Supports Anymatch syntax for flexible file pattern matching
- Provides detailed execution guidance in the trigger panel
- Can monitor both individual files and entire directory trees
- Handles symlinks, aliases, and shortcuts appropriately
- Offers polling fallback for systems with limited file system event support

## UseCases

- **Automated Document Processing**: Monitor upload folders for new PDF invoices, contracts, or forms and automatically trigger processing workflows for data extraction and routing
- **Log File Analysis**: Watch application log files for changes and trigger real-time analysis, alerting, or log aggregation workflows when new entries are detected
- **Content Management Automation**: Monitor content directories for new images, videos, or documents and automatically trigger optimization, metadata extraction, or publishing workflows
- **Configuration Management**: Watch configuration files and automatically trigger deployment, validation, or notification workflows when changes are detected
- **Data Pipeline Triggers**: Monitor data drop folders for new CSV, JSON, or XML files and automatically trigger ETL processes or data validation workflows
- **Backup and Synchronization**: Watch critical directories for changes and trigger backup operations or file synchronization across multiple systems
- **Development Workflow Automation**: Monitor source code directories for changes and trigger build processes, testing, or deployment workflows
- **Media Processing**: Watch media upload folders and automatically trigger video transcoding, image resizing, or metadata extraction workflows
- **Security Monitoring**: Monitor sensitive directories for unauthorized file changes and trigger security alert or audit workflows
- **Report Generation**: Watch data directories for new files and automatically trigger report generation or dashboard update workflows

**Note**: This node requires local file system access and works with n8n's file monitoring capabilities. Ensure proper file permissions and consider performance implications when monitoring large directory structures.

