# Dropbox Node

## Overview

The Dropbox node allows you to interact with Dropbox cloud storage through the Dropbox API. This node enables you to upload, download, search, and manage files and folders in your Dropbox account. You can automate file operations, synchronize data, and integrate Dropbox storage with your n8n workflows for comprehensive file management and backup solutions.

## Credentials

- Name: dropboxApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: file

#### Operation: download

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Path | string | The file path of the file to download. Has to contain the full path. | Yes |  |
| Put Output File in Field | string |  | Yes |  |

#### Operation: upload

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Path | string | The file path of the file to upload. Has to contain the full path. The parent folder has to exist. Existing files get overwritten. | Yes |  |
| Binary File | boolean | Whether the data to upload should be taken from binary field | No |  |
| File Content | string | The text content of the file to upload | No |  |
| Input Binary Field | string |  | Yes |  |

### Resource: search

#### Operation: query

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Query | string | The string to search for. May match across multiple fields based on the request arguments. | Yes |  |
| File Status | options | The string to search for. May match across multiple fields based on the request arguments. | No |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |
| Filters | collection | Multiple file extensions can be set separated by comma. Example: jpg,pdf. | No |  |

### Resource: folder

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder | string | The folder to create. The parent folder has to exist. | Yes |  |

#### Operation: list

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder Path | string | The path of which to list the content | No |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether the results will include entries for files and folders that used to exist but were deleted. The default for this field is False. | No |  |

## UseCases

- **File Backup**: Automatically backup important files to Dropbox storage
- **Document Management**: Organize and manage documents in Dropbox folders
- **File Synchronization**: Sync files between different systems via Dropbox
- **Automated Uploads**: Upload generated reports, images, or data files automatically
- **Content Distribution**: Share files and folders with team members or clients
- **Data Archiving**: Archive old files and data for long-term storage
- **Collaborative Workflows**: Enable file sharing and collaboration in workflows
- **Media Processing**: Store and retrieve media files for processing workflows
- **Configuration Management**: Store and retrieve configuration files and templates
- **Log File Management**: Upload and organize log files from various systems

