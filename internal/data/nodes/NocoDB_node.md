# NocoDB Node

## Overview

Read, update, write and delete data from NocoDB. NocoDB is an open-source no-code database platform that turns any database into a smart spreadsheet interface. It provides a user-friendly way to manage data with features like forms, views, and API access, making it ideal for rapid application development and data management workflows.

## Credentials

This node supports multiple authentication methods:
- **Name**: nocoDb (User Token) 
- **Name**: nocoDbApiToken (API Token)
- **Required**: Yes

### Authentication Options

#### API Token Authentication (nocoDbApiToken) - Recommended
- **Access Method**: Uses API token for programmatic access
- **Token Type**: Project-specific or workspace-specific API tokens
- **Security**: Recommended for production environments and automated workflows
- **Credential Configuration**: API token with base URL
- **Use Case**: Server-to-server integration, automated workflows, production systems

#### User Token Authentication (nocoDb)
- **Access Method**: Uses user session-based authentication
- **Token Type**: User authentication tokens from login sessions
- **Security**: Suitable for development and testing environments
- **Credential Configuration**: User credentials with base URL
- **Use Case**: Development environments, testing, user-interactive scenarios

### Credential Configuration
To configure NocoDB API credentials:
1. Access your NocoDB instance
2. Navigate to Account Settings → Tokens
3. Generate a new API token with appropriate permissions
4. Copy the token and base URL for use in n8n
5. Ensure the token has access to the workspaces and bases you plan to use

## Inputs

- **Main**: JSON data for creating/updating NocoDB records

## Outputs

- **Main**: Returns JSON responses from NocoDB API including record data and operation results

## Properties

### Resource: row

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Authentication | Options | Authentication method (API Token/User Token) | Yes | API Token |
| API Version | Options | NocoDB API version (Before v0.90.0/v0.90.0 Onwards/v0.200.0 Onwards) | Yes | v0.200.0 Onwards |
| Workspace Name or ID | Options | Choose from the list, or specify an ID using an expression | No | none |
| Base Name or ID | Options | Choose from the list, or specify an ID using an expression | Yes | - |
| Table Name or ID | Options | The table to operate on | Yes | - |
| Data to Send | Collection | The data to send | No | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Authentication | Options | Authentication method (API Token/User Token) | Yes | API Token |
| API Version | Options | NocoDB API version (Before v0.90.0/v0.90.0 Onwards/v0.200.0 Onwards) | Yes | v0.200.0 Onwards |
| Workspace Name or ID | Options | Choose from the list, or specify an ID using an expression | No | none |
| Base Name or ID | Options | Choose from the list, or specify an ID using an expression | Yes | - |
| Table Name or ID | Options | The table to operate on | Yes | - |
| Primary Key Type | Options | Type of primary key field (varies by API version) | Yes | id |
| Row ID Value | String | The value of the ID field | Yes | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Authentication | Options | Authentication method (API Token/User Token) | Yes | API Token |
| API Version | Options | NocoDB API version (Before v0.90.0/v0.90.0 Onwards/v0.200.0 Onwards) | Yes | v0.200.0 Onwards |
| Workspace Name or ID | Options | Choose from the list, or specify an ID using an expression | No | none |
| Base Name or ID | Options | Choose from the list, or specify an ID using an expression | Yes | - |
| Table Name or ID | Options | The table to operate on | Yes | - |
| Row ID Value | String | The value of the ID field | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Authentication | Options | Authentication method (API Token/User Token) | Yes | API Token |
| API Version | Options | NocoDB API version (Before v0.90.0/v0.90.0 Onwards/v0.200.0 Onwards) | Yes | v0.200.0 Onwards |
| Workspace Name or ID | Options | Choose from the list, or specify an ID using an expression | No | none |
| Base Name or ID | Options | Choose from the list, or specify an ID using an expression | Yes | - |
| Table Name or ID | Options | The table to operate on | Yes | - |
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 50 |
| Download Attachments | Boolean | Whether to download attachments | No | false |
| Options | Collection | Additional options | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Authentication | Options | Authentication method (API Token/User Token) | Yes | API Token |
| API Version | Options | NocoDB API version (Before v0.90.0/v0.90.0 Onwards/v0.200.0 Onwards) | Yes | v0.200.0 Onwards |
| Workspace Name or ID | Options | Choose from the list, or specify an ID using an expression | No | none |
| Base Name or ID | Options | Choose from the list, or specify an ID using an expression | Yes | - |
| Table Name or ID | Options | The table to operate on | Yes | - |
| Row ID Value | String | The value of the ID field | Yes | - |
| Data to Send | Collection | The data to send | No | - |

## Authentication Methods

### API Token (Recommended)
- **Description**: Uses NocoDB API token configured in credentials
- **Credential Type**: nocoDbApiToken
- **Use Case**: Production environments and automated workflows
- **Security**: More secure, token-based authentication with fine-grained permissions

### User Token  
- **Description**: Uses user session-based authentication
- **Credential Type**: nocoDb
- **Use Case**: Development and testing environments
- **Security**: Session-based authentication suitable for user-interactive scenarios

## API Version Differences

### Before v0.90.0
- **Legacy API**: Basic functionality with limited features
- **Primary Key**: Simple ID-based primary keys
- **Field Display**: Basic field configuration without advanced options
- **Compatibility**: For older NocoDB installations

### v0.90.0 Onwards
- **Enhanced API**: Improved functionality and better error handling
- **Primary Key Options**: Extended primary key type support
- **Field Requirements**: Version-specific field visibility and requirements
- **Workspace Support**: Enhanced workspace and base management

### v0.200.0 Onwards (Recommended)
- **Modern API**: Latest features with full functionality support
- **Advanced Primary Keys**: Complex primary key configurations with custom options
- **Dynamic Fields**: Version-dependent field display logic
- **Load Options**: Dynamic field loading with dependency management (getWorkspaces, getBases, getTables)
- **Optimized Performance**: Best performance and reliability

## Version-Specific Field Logic

### Primary Key Configuration
- **Before v0.90.0**: Basic "id" field only
- **v0.90.0 Onwards**: Extended primary key types with conditional display
- **v0.200.0 Onwards**: Full custom primary key support with complex configurations

### Field Visibility
- **Dynamic Display**: Fields shown/hidden based on selected API version
- **Option Dependencies**: Field options loaded based on version-specific capabilities
- **Conditional Requirements**: Some fields required only for specific versions and operations

### Workspace and Base Management
- **Version-Dependent**: Workspace features available in newer versions
- **Load Options Methods**: Dynamic loading of workspaces, bases, and tables
- **Dependency Chain**: loadOptionsDependsOn functionality for related field updates

## Technical Details

### API Endpoint Structure
- **Version-Specific URLs**: Different API endpoints for each version
- **Authentication Headers**: Varies by authentication method and version
- **Response Format**: Consistent JSON response structure across versions

### Dynamic Field Loading
- **getWorkspaces**: Load available workspaces dynamically
- **getBases**: Load bases based on selected workspace
- **getTables**: Load tables based on selected base
- **Dependency Management**: Fields update automatically based on selections

### Error Handling
- **Version-Specific Errors**: Different error handling for each API version
- **Authentication Errors**: Specific handling for token and user authentication failures
- **Validation**: Field validation based on version requirements

## UseCases

- **Rapid Prototyping and Development**: Quickly create database-backed applications and prototypes using NocoDB's no-code interface for faster development cycles
- **Data Collection and Form Management**: Use NocoDB forms to collect data from users and automatically store responses in structured database tables
- **Small Business Data Management**: Manage customer information, inventory, orders, and business operations through an intuitive spreadsheet-like interface
- **Content Management Systems**: Create and manage content databases for websites, blogs, and applications with easy-to-use views and filters
- **Project and Task Management**: Build custom project management solutions with tables for tasks, assignments, deadlines, and progress tracking
- **Data Migration and Integration**: Transfer data between systems, synchronize databases, and integrate NocoDB with existing workflows and applications
- **Collaborative Database Management**: Enable team collaboration on shared databases with role-based access control and real-time updates
- **API-First Database Solutions**: Leverage NocoDB's REST API to create custom applications and integrations with external systems
- **Reporting and Analytics**: Generate reports and dashboards from NocoDB data using its built-in views and external analytics tools
- **Legacy System Modernization**: Replace outdated spreadsheets and legacy systems with modern, API-enabled database solutions

