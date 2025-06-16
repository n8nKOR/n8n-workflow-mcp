# BambooHr Node

## Overview

The BambooHR node allows you to consume the BambooHR API for human resources management. BambooHR is a comprehensive HR software platform that helps organizations manage employee information, time tracking, performance reviews, and HR reporting. This node enables you to automate HR processes and integrate employee data with your workflows.

## Credentials

This node requires BambooHR API credentials:
- **Subdomain**: Your BambooHR company subdomain
- **API Key**: Your BambooHR API key

To obtain API credentials:
1. Log into your BambooHR account
2. Navigate to Settings → API Keys
3. Generate a new API key
4. Note your company subdomain from the URL (e.g., company.bamboohr.com)
5. Use both subdomain and API key in n8n

## Inputs

- **Main**: JSON data for creating/updating employee records and files

## Outputs

- **Main**: Returns JSON data from BambooHR API responses

## Properties

### Resource: Employee

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| First Name | String | Employee's first name | Yes |
| Last Name | String | Employee's last name | Yes |
| Additional Fields | Collection | Optional employee properties | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Employee ID | String | ID of the employee to retrieve | Yes |
| Fields | Multiselect | Specific fields to retrieve | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all employees | No |
| Limit | Number | Maximum number of results to return | No |
| Fields | Multiselect | Specific fields to retrieve | No |
| Filters | Collection | Employee filtering options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Employee ID | String | ID of the employee to update | Yes |
| Update Fields | Collection | Employee properties to update | No |

### Resource: Employee Document

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Employee ID | String | ID of the employee | Yes |
| File ID | String | ID of the document to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Employee ID | String | ID of the employee | Yes |

#### Operation: Upload
| Field Name | Type | Description | Required |
|---|---|---|---|
| Employee ID | String | ID of the employee | Yes |
| Category | Options | Document category | Yes |
| File Name | String | Name of the file | Yes |
| Binary Data | String | Binary property containing file data | Yes |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Employee ID | String | ID of the employee | Yes |
| File ID | String | ID of the document to update | Yes |
| Update Fields | Collection | Document properties to update | No |

### Resource: File

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| File ID | String | ID of the file to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all files | No |
| Limit | Number | Maximum number of results to return | No |

### Resource: Company Report

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Report Format | Options | Format of the report (CSV, PDF, XLS, XML, JSON) | Yes |
| Fields | Multiselect | Fields to include in the report | No |
| Filters | Collection | Report filtering options | No |

## UseCases

- **Employee Onboarding** : Automate new employee creation and document collection during onboarding
- **HR Data Synchronization** : Sync employee information between BambooHR and other business systems
- **Document Management** : Automate employee document uploads and organization
- **Reporting and Analytics** : Generate automated HR reports for management and compliance
- **Employee Status Tracking** : Monitor and update employee status changes and lifecycle events
- **Performance Management** : Manage performance review documents and tracking
- **Compliance Reporting** : Generate compliance reports for regulatory requirements
- **Employee Directory Updates** : Maintain up-to-date employee directories across systems
- **Benefits Administration** : Manage employee benefits documentation and enrollment
- **Payroll Integration** : Sync employee data with payroll systems for accurate processing
