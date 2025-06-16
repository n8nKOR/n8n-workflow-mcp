# FTP Node

## Overview

The FTP node enables secure and efficient file transfer operations via FTP (File Transfer Protocol) and SFTP (SSH File Transfer Protocol) protocols. This versatile node supports comprehensive file management operations including upload, download, listing, deletion, and renaming of files and directories on remote servers.

## Credentials

The FTP node requires protocol-specific credentials:

- **FTP Protocol**: Username/Password authentication for FTP connections
- **SFTP Protocol**: Username/Password or SSH Key-based authentication for SFTP connections

## Inputs

- **Main**: Input data for file operations and path information

## Outputs

- **Main**: Returns operation results and file metadata

## Properties

### Resource: File Operations

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Protocol | Options | File transfer protocol (FTP/SFTP) | Yes |
| Path | String | Full path to the file or directory to delete | Yes |
| Options | Collection | Additional deletion options | No |

#### Operation: Download

| Field Name | Type | Description | Required |
|---|---|---|---|
| Protocol | Options | File transfer protocol (FTP/SFTP) | Yes |
| Path | String | Full path to the file to download | Yes |
| Put Output File in Field | String | Name of binary field for downloaded file | Yes |

#### Operation: List

| Field Name | Type | Description | Required |
|---|---|---|---|
| Protocol | Options | File transfer protocol (FTP/SFTP) | Yes |
| Path | String | Directory path to list | Yes |
| Recursive | Boolean | Include subdirectories recursively | No |

#### Operation: Rename

| Field Name | Type | Description | Required |
|---|---|---|---|
| Protocol | Options | File transfer protocol (FTP/SFTP) | Yes |
| Old Path | String | Current path of the file/directory | Yes |
| New Path | String | New path for the file/directory | Yes |
| Options | Collection | Additional rename options | No |

#### Operation: Upload

| Field Name | Type | Description | Required |
|---|---|---|---|
| Protocol | Options | File transfer protocol (FTP/SFTP) | Yes |
| Path | String | Full destination path for the uploaded file | Yes |
| Binary File | Boolean | Use binary data or text content | Yes |
| Input Binary Field | String | Binary field containing file data | Conditional |
| File Content | String | Text content for non-binary uploads | Conditional |
| Options | Collection | Additional upload options | No |

## UseCases

- **File Backup and Synchronization**: Transfer files between local and remote servers for backup purposes
- **Website Deployment**: Upload website files and assets to web servers
- **Data Exchange**: Share files with partners and external systems
- **Automated File Processing**: Download files for processing and upload results
- **Document Management**: Organize and manage documents on remote file servers
- **Log File Collection**: Gather log files from remote servers for analysis 