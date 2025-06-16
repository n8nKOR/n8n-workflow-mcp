# Nextcloud Node

## Overview

⚠️ **Operation Coverage**: This documentation includes the major file, folder, and user operations. The implementation supports comprehensive operations including copy, delete, move, and share operations for both files and folders, plus user delete and get operations.

Comprehensive integration with Nextcloud for file management, user administration, and collaboration workflows. Nextcloud is an open-source cloud storage and collaboration platform that provides file synchronization, sharing, and team collaboration features. This node enables automation of file operations, user management, and document workflows.

## Credentials

- **Name**: nextCloudApi
- **Required**: Yes

### Credential Configuration
To configure Nextcloud API credentials:
1. Access your Nextcloud instance
2. Go to Settings → Security
3. Create an app password or use OAuth2 authentication
4. For OAuth2: Register your application in Nextcloud admin settings
5. Use the appropriate authentication method (Basic Auth or OAuth2)

### Authentication Methods
- **Basic Authentication**: Username and app password
- **OAuth2**: Client ID, client secret, and OAuth2 flow

## Inputs

- **Main**: File data, user information, or folder parameters

## Outputs

- **Main**: Returns JSON responses including file information, operation status, and metadata

## Properties

### Resource: file

#### Operation: download

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Path | String | The file path of the file to download. Has to contain the full path starting with / | Yes | - |
| Put Output File in Field | String | Field name to store the downloaded file | Yes | data |

#### Operation: upload

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Path | String | The absolute file path of the file to upload. Must contain the full path. Parent folder must exist. | Yes | - |
| Binary File | Boolean | Whether the input is binary file data | Yes | - |
| File Content | String | The text content of the file to upload | No | - |
| Input Binary Field | String | Field name containing binary file data | Yes | data |

#### Operation: copy

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Source Path | String | Path of the file to copy | Yes | - |
| Destination Path | String | Destination path for the copied file | Yes | - |
| Overwrite | Boolean | Overwrite existing file at destination | No | false |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Path | String | Path of the file to delete | Yes | - |

#### Operation: move

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Source Path | String | Current path of the file | Yes | - |
| Destination Path | String | New path for the file | Yes | - |
| Overwrite | Boolean | Overwrite existing file at destination | No | false |

#### Operation: share

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Path | String | Path of the file to share | Yes | - |
| Share Type | Options | Type of share (user/group/public link) | Yes | user |
| Share With | String | Username or group name to share with | No | - |
| Permissions | Options | Share permissions (read/write/admin) | No | read |
| Password | String | Password for public links | No | - |
| Expiration Date | String | Share expiration date (YYYY-MM-DD) | No | - |

### Resource: folder

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder | String | The folder to create. Parent folder must exist. Path should start with / | Yes | - |

#### Operation: list

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder Path | String | The path of which to list the content. Path should start with / | No | / |
| Include Properties | Boolean | Include detailed file properties | No | false |

#### Operation: copy

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Source Path | String | Path of the folder to copy | Yes | - |
| Destination Path | String | Destination path for the copied folder | Yes | - |
| Overwrite | Boolean | Overwrite existing folder at destination | No | false |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder Path | String | Path of the folder to delete | Yes | - |
| Recursive | Boolean | Delete folder and all contents | No | true |

#### Operation: move

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Source Path | String | Current path of the folder | Yes | - |
| Destination Path | String | New path for the folder | Yes | - |
| Overwrite | Boolean | Overwrite existing folder at destination | No | false |

#### Operation: share

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder Path | String | Path of the folder to share | Yes | - |
| Share Type | Options | Type of share (user/group/public link) | Yes | user |
| Share With | String | Username or group name to share with | No | - |
| Permissions | Options | Share permissions (read/write/admin) | No | read |
| Password | String | Password for public links | No | - |
| Expiration Date | String | Share expiration date (YYYY-MM-DD) | No | - |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | String | Username for the new user | Yes | - |
| Email | String | Email address of the user | Yes | - |
| Password | String | Password for the new user | No | - |
| Display Name | String | Display name of the user | No | - |
| Groups | Array | Groups to add the user to | No | [] |
| Quota | String | Storage quota for the user | No | - |
| Additional Fields | Collection | Additional user configuration | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | String | Username of the user to retrieve | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 50 |
| Search | String | Search term for filtering users | No | - |
| Options | Collection | Additional query options | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | String | Username of the user to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

**Update Fields Options:**
- **Display Name**: New display name
- **Email**: New email address
- **Password**: New password
- **Quota**: New storage quota
- **Groups**: Update group memberships

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | String | Username of the user to delete | Yes | - |

## UseCases

- **File Synchronization**: Sync files between Nextcloud and external systems for data backup and cross-platform sharing
- **Document Management**: Manage document storage, versioning, sharing, and collaboration workflows for teams
- **User Management**: Automate user account creation, updates, and management for team onboarding and administration
- **Content Collaboration**: Enable team collaboration on shared files and documents with permission management
- **Backup and Archive**: Automate backup processes for important files and data with scheduled uploads
- **File Processing**: Process and transform files stored in Nextcloud through automated workflows
- **Integration Workflows**: Connect Nextcloud with other business applications and services for seamless data flow
- **Automated File Organization**: Organize files automatically based on content, date, metadata, or business rules
- **Notification Systems**: Send notifications when files are uploaded, modified, shared, or when storage quotas are reached
- **Data Migration**: Migrate files and data between Nextcloud instances or other cloud storage platforms
- **Access Control Automation**: Automate file and folder sharing permissions based on user roles and business logic
- **Compliance and Audit**: Track file access, modifications, and sharing for compliance and audit requirements
- **Bulk Operations**: Perform bulk file operations like copying, moving, or sharing multiple files simultaneously
- **Workflow Triggers**: Trigger business processes when specific files are uploaded or modified
- **Storage Management**: Monitor and manage storage usage, quotas, and cleanup of old or unused files
- **External Integration**: Integrate Nextcloud with CRM, ERP, or other business systems for centralized file management
- **Team Onboarding**: Automate creation of user accounts, folder structures, and initial file sharing for new team members
- **Project Management**: Create and manage project folders, share resources, and coordinate team collaboration
- **Document Approval**: Implement document approval workflows with automated sharing and notification systems
- **Content Distribution**: Distribute content to specific users or groups based on roles and access requirements

