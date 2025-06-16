# Box Node

## Overview

The Box node allows you to consume the Box API for cloud storage management. Box is an enterprise cloud content management platform that enables secure file sharing, collaboration, and storage. This node provides comprehensive file and folder operations for automating document workflows and content management.

## Credentials

This node requires Box OAuth2 API credentials:
- **Client ID**: Your Box application client ID
- **Client Secret**: Your Box application client secret

To obtain API credentials:
1. Create a Box developer account at https://developer.box.com
2. Create a new Box application
3. Configure OAuth2 authentication with appropriate scopes
4. Copy client ID and secret for n8n configuration

## Inputs

- **Main**: JSON data for file/folder operations

## Outputs

- **Main**: Returns JSON data from Box API responses

## Properties

### Resource: File

#### Operation: Copy
| Field Name | Type | Description | Required |
|---|---|---|---|
| File ID | String | ID of the file to copy | Yes |
| Parent ID | String | ID of the destination folder | Yes |
| Name | String | New name for the copied file | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| File ID | String | ID of the file to delete | Yes |

#### Operation: Download
| Field Name | Type | Description | Required |
|---|---|---|---|
| File ID | String | ID of the file to download | Yes |
| Binary Property | String | Name of binary property to store file data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| File ID | String | ID of the file to retrieve | Yes |

#### Operation: Search
| Field Name | Type | Description | Required |
|---|---|---|---|
| Query | String | Search query for files | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Share
| Field Name | Type | Description | Required |
|---|---|---|---|
| File ID | String | ID of the file to share | Yes |
| Access Level | Options | Share access level (open, company, collaborators) | Yes |
| Password | String | Password protection for shared link | No |

#### Operation: Upload
| Field Name | Type | Description | Required |
|---|---|---|---|
| Parent ID | String | ID of the parent folder | Yes |
| Binary Data | Binary | Binary file data to upload | Conditional |
| File Name | String | Name of the file | Yes |

### Resource: Folder

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Name of the new folder | Yes |
| Parent ID | String | ID of the parent folder | Yes |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Folder ID | String | ID of the folder to delete | Yes |
| Recursive | Boolean | Whether to delete recursively | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Folder ID | String | ID of the folder to retrieve | Yes |

#### Operation: Search
| Field Name | Type | Description | Required |
|---|---|---|---|
| Query | String | Search query for folders | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Share
| Field Name | Type | Description | Required |
|---|---|---|---|
| Folder ID | String | ID of the folder to share | Yes |
| Access Level | Options | Share access level (open, company, collaborators) | Yes |
| Password | String | Password protection for shared link | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Folder ID | String | ID of the folder to update | Yes |
| Name | String | New name for the folder | No |
| Description | String | New description for the folder | No |

## UseCases

- **Enterprise Document Storage** : Centralize company documents with proper access controls
- **File Organization Automation** : Automatically sort and organize uploaded files by type or metadata
- **Version Control Workflows** : Manage document versions and track changes across teams
- **Archive Management** : Move old documents to archive folders based on age or usage
- **Document Backup** : Backup critical business documents from other systems to Box
- **Project File Management** : Create and manage project-specific folder structures
- **Team File Sharing** : Automatically share project files with team members based on roles
- **Client Collaboration** : Set up secure client portals for file sharing and collaboration
- **Review and Approval Workflows** : Route documents through approval processes
- **Real-time Synchronization** : Keep files synchronized between Box and other business systems
- **CRM Document Storage** : Store customer-related documents linked to CRM records
- **Invoice and Contract Management** : Organize financial documents with automated filing
- **HR Document Workflows** : Manage employee onboarding documents and personnel files
- **Compliance and Audit** : Maintain regulatory compliance with structured document storage
- **Marketing Asset Management** : Organize and distribute marketing materials across teams
- **Digital Asset Distribution** : Manage and distribute multimedia content and brand assets
- **Training Material Management** : Distribute training content to employees and partners
- **Product Documentation** : Maintain and share product manuals and specifications
- **Client Deliverable Management** : Organize and deliver project deliverables to clients
- **Content Publishing Workflows** : Automate content publication and distribution processes
- **API Data Persistence** : Store API responses and exported data in organized structures
- **Report Storage** : Automatically save business reports and analytics data
- **Database File Storage** : Store database exports and backup files
- **Integration Hub** : Use Box as a central hub for file-based system integrations
- **Multi-system File Synchronization** : Keep files consistent across multiple business platforms