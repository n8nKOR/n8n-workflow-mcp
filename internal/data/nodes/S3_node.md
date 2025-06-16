# S3 Node

## Overview

Sends data to any S3-compatible service. This node is designed for services that use the S3 standard, such as MinIO, Digital Ocean Spaces, Wasabi, Backblaze B2, and other S3-compatible storage providers. It provides comprehensive bucket and file management capabilities including upload, download, delete, and listing operations. For AWS S3 specifically, use the dedicated 'AWS S3' node instead.

## Credentials

This node requires S3-compatible service credentials:
- **Credential Name**: `s3`
- **Required Fields**: 
  - Access Key ID: Your S3 access key
  - Secret Access Key: Your S3 secret key
  - Region: Storage region (e.g., us-east-1)
  - Endpoint: Service endpoint URL (for non-AWS services)

**Supported Services:**
- **MinIO**: Self-hosted object storage
- **Digital Ocean Spaces**: Digital Ocean's object storage
- **Wasabi**: Hot cloud storage service
- **Backblaze B2**: Cloud storage service
- **Linode Object Storage**: Linode's S3-compatible storage
- **Any S3-compatible service**: Custom endpoints supported

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Bucket

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket to create | Yes | - |
| Region | options | AWS region for the bucket | No | us-east-1 |
| ACL | options | Access control list permissions | No | private |

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket to delete | Yes | - |

#### Operation: Get All
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Return all buckets | No | true |

### Resource: File

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket | Yes | - |
| File Key | string | Key/path of the file to delete | Yes | - |

#### Operation: Download
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket | Yes | - |
| File Key | string | Key/path of the file to download | Yes | - |
| Binary Property | string | Property name to store binary data | No | data |

#### Operation: Get All
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket | Yes | - |
| Return All | boolean | Return all files | No | false |
| Limit | number | Maximum number of files to return | No | 100 |
| Prefix | string | Prefix to filter files | No | - |

#### Operation: Upload
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket | Yes | - |
| File Key | string | Key/path for the uploaded file | Yes | - |
| Binary Data | boolean | Whether data is binary | Yes | true |
| Binary Property | string | Property containing binary data | No | data |
| File Content | string | Text content for non-binary files | No | - |
| ACL | options | Access control list permissions | No | private |

### Resource: Folder

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket | Yes | - |
| Folder Key | string | Key/path of the folder to create | Yes | - |

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket | Yes | - |
| Folder Key | string | Key/path of the folder to delete | Yes | - |

#### Operation: Get All
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bucket Name | string | Name of the bucket | Yes | - |
| Folder Key | string | Folder path to list contents | No | - |
| Return All | boolean | Return all folders | No | false |
| Limit | number | Maximum number of folders to return | No | 100 |

## UseCases

- **File Storage**: Store and manage files in S3-compatible storage services
- **Backup Solutions**: Backup files and data to cloud storage providers
- **Static Website Hosting**: Host static websites using S3-compatible services
- **Data Archiving**: Archive old data for long-term storage and compliance
- **Content Distribution**: Distribute files and media content globally
- **Document Management**: Store and organize business documents
- **Media Processing**: Store images, videos, and audio files for processing
- **Data Lake Storage**: Build data lakes with S3-compatible storage
- **Application Assets**: Store application assets and resources
- **Log Storage**: Archive application and system logs



