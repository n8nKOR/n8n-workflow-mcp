# Baserow Node

## Overview

The Baserow node allows you to consume the Baserow API to manage data in your Baserow databases. Baserow is an open-source no-code database platform that provides a user-friendly alternative to traditional databases and spreadsheets. This node enables you to create, read, update, and delete rows in your Baserow tables through n8n workflows.

## Credentials

This node requires Baserow API credentials:
- **Username/Email**: Your Baserow account email
- **Password**: Your Baserow account password
- **Host**: Baserow instance URL (e.g., https://api.baserow.io for cloud, or your self-hosted URL)

To obtain API credentials:
1. Sign up for a Baserow account (cloud or self-hosted)
2. Use your account email and password
3. For cloud: use https://api.baserow.io as host
4. For self-hosted: use your custom Baserow instance URL

## Inputs

- **Main**: JSON data for creating/updating table rows

## Outputs

- **Main**: Returns JSON data from Baserow API responses

## Properties

### Resource: Row

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | Options | Database containing the target table | Yes |
| Table ID | Options | Table to create the row in | Yes |
| Data to Send | Options | Choose between defining fields or using JSON | Yes |
| Fields | Collection | Individual field values (when using Define Below) | No |
| JSON Data | JSON | Complete row data as JSON object | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | Options | Database containing the target table | Yes |
| Table ID | Options | Table containing the row | Yes |
| Row ID | Number | ID of the row to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | Options | Database containing the target table | Yes |
| Table ID | Options | Table containing the row | Yes |
| Row ID | Number | ID of the row to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | Options | Database containing the target table | Yes |
| Table ID | Options | Table to retrieve rows from | Yes |
| Return All | Boolean | Whether to return all rows | No |
| Limit | Number | Maximum number of rows to return | No |
| Additional Options | Collection | Filtering, sorting, and search options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | Options | Database containing the target table | Yes |
| Table ID | Options | Table containing the row | Yes |
| Row ID | Number | ID of the row to update | Yes |
| Data to Send | Options | Choose between defining fields or using JSON | Yes |
| Fields | Collection | Individual field values to update | No |
| JSON Data | JSON | Complete row data as JSON object | No |

## Field Types

Baserow supports various field types that you can work with:

### Text Fields
- **Text**: Single-line text field
- **Long Text**: Multi-line text field
- **Rich Text**: Formatted text with styling

### Number Fields
- **Number**: Numeric values with decimal support
- **Rating**: Star rating field (1-5 scale)

### Selection Fields
- **Single Select**: Choose one option from predefined list
- **Multiple Select**: Choose multiple options from predefined list

### Date and Time
- **Date**: Date values
- **Last Modified**: Automatically updated timestamp
- **Created On**: Automatically set creation timestamp

### Relationship Fields
- **Link to Table**: References to rows in other tables
- **Lookup**: Display values from linked rows
- **Count**: Count of linked rows

### Special Fields
- **Boolean**: True/false checkbox
- **URL**: Website links
- **Email**: Email addresses
- **Phone Number**: Phone number field
- **File**: File attachments
- **Formula**: Calculated values based on other fields

## Additional Options for Get Many

### Filtering
| Field Name | Type | Description |
|---|---|---|
| Filter Type | Options | AND or OR logic for multiple filters |
| Filters | Collection | Field-based filtering conditions |

### Sorting
| Field Name | Type | Description |
|---|---|---|
| Sort | Collection | Field and direction for sorting results |

### Search
| Field Name | Type | Description |
|---|---|---|
| Search | String | Search term to find rows |

## UseCases

- **Database Management Automation** : Automate data entry and management in Baserow databases
- **Data Synchronization** : Sync data between Baserow and external systems
- **Form Processing** : Process form submissions and store data in structured tables
- **Project Management** : Manage projects, tasks, and team collaboration data
- **Customer Data Management** : Maintain customer information and interaction records
- **Inventory Tracking** : Track product inventory and stock management
- **Event Management** : Manage event registrations and attendee information
- **Content Management** : Organize and manage content databases
- **Reporting and Analytics** : Generate reports from Baserow data for business insights
- **Workflow Automation** : Create automated workflows based on database changes
For self-hosted Baserow instances:
- Ensure proper SSL configuration
- Configure CORS settings if needed
- Monitor server resources and performance
- Keep Baserow updated to latest version
- Implement proper backup strategies

## UseCases

- **Database Management** : Manage structured data in Baserow's no-code database platform
- **Project Management** : Track projects, tasks, and team collaboration using Baserow tables
- **Customer Relationship Management** : Maintain customer databases and interaction histories
- **Inventory Management** : Track inventory levels, products, and stock movements
- **Content Management** : Organize and manage content for websites and applications
- **Event Planning** : Coordinate events, attendees, and logistics using structured data
- **Survey Data Collection** : Collect and analyze survey responses and feedback data
- **Team Collaboration** : Enable team collaboration through shared databases and workflows
- **Data Integration** : Integrate Baserow data with other business systems and applications
- **Reporting and Analytics** : Generate reports and analytics from Baserow database content

