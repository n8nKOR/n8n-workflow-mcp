# Pipedrive Node

## Overview

Comprehensive CRM integration with Pipedrive for managing sales pipelines, contacts, deals, and activities. Pipedrive is a sales-focused CRM platform that helps businesses track leads through the sales process with visual pipeline management, activity tracking, and performance analytics.

## Credentials

This node supports multiple authentication methods:
- **Name**: pipedriveOAuth2Api (OAuth2)
- **Name**: pipedriveApi (API Token)
- **Required**: Yes

### Authentication Methods

#### OAuth2 Authentication (pipedriveOAuth2Api)
- **Client ID**: Your Pipedrive app client ID
- **Client Secret**: Your Pipedrive app client secret
- **Authorization URL**: Pipedrive OAuth authorization endpoint
- **Access Token URL**: Pipedrive OAuth token endpoint
- **Scope**: Required permissions (read, write, admin)

#### API Token Authentication (pipedriveApi)
- **API Token**: Your Pipedrive API token
- **Company Domain**: Your Pipedrive company domain

### Authentication Parameter
- **Authentication**: Choose between "API Token" and "OAuth2" methods

## Inputs

- **Main**: JSON data for creating, updating, or querying Pipedrive resources

## Outputs

- **Main**: Returns JSON responses with Pipedrive resource data, operation results, and metadata

## Properties

### Resource: activity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subject | string | The subject of the activity to create | Yes |  |
| Done | options | Whether the activity is done or not | No |  |
| Type | string | Type of the activity like  | Yes |  |
| Additional Fields | collection | ID of the deal this activity will be associated with | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | number | ID of the activity to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | number | ID of the activity to get | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | number | ID of the activity to update | Yes |  |
| Update Fields | collection | Whether the user is set to busy during the activity | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Whether the Activity is done or not. 0 = Not done, 1 = Done. If omitted returns both Done and Not done activities. | No |  |

### Resource: deal

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | The title of the deal to create | Yes |  |
| Associate With | options | Type of entity to link to this deal | Yes |  |
| Organization ID | number | ID of the organization this deal will be associated with | Yes |  |
| Person ID | number | ID of the person this deal will be associated with | No |  |
| Additional Fields | collection | Currency of the deal. Accepts a 3-character currency code. Like EUR, USD, ... | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | number | ID of the deal to delete | Yes |  |

#### Operation: duplicate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | number | ID of the deal to duplicate | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | number | ID of the deal to get | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal ID | number | ID of the deal to update | Yes |  |
| Update Fields | collection | Currency of the deal. Accepts a 3-character currency code. Like EUR, USD, ... | No |  |

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Term | string | The search term to look for. Minimum 2 characters (or 1 if using exact_match). | Yes |  |
| Exact Match | boolean | Whether only full exact matches against the given term are returned. It is not case sensitive. | No |  |
| Additional Fields | collection | Supports including optional fields in the results which are not provided by default. Example: deal.cc_email. | No | "[\"custom_fields\"" |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection | Predefined filter to apply to the deals to retrieve. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: dealProduct

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal Name or ID | options | The ID of the deal to add a product to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Product Name or ID | options | The ID of the product to add to a deal. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Item Price | number | Price at which to add or update this product in a deal | Yes |  |
| Quantity | number | How many items of this product to add/update in a deal | Yes |  |
| Additional Fields | collection | Text to describe this product-deal attachment | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal Name or ID | options | The ID of the deal whose product to update. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Product Attachment Name or ID | options | ID of the deal-product (the ID of the product attached to the deal). Choose from the list, or specify an ID using an <a href= | Yes |  |
| Update Fields | collection | Text to describe this product-deal attachment | No |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal Name or ID | options | The ID of the deal whose product to remove. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Product Attachment Name or ID | options | ID of the deal-product (the ID of the product attached to the deal). Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deal Name or ID | options | The ID of the deal whose products to retrieve. Choose from the list, or specify an ID using an <a href= | Yes |  |

### Resource: file

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Input Binary Field | string |  | Yes |  |
| Additional Fields | collection | ID of the activite this file will be associated with | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File ID | number | ID of the file to delete | Yes |  |

#### Operation: download

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File ID | number | ID of the file to download | Yes |  |
| Put Output File in Field | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File ID | number | ID of the file to get | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File ID | number | ID of the file to update | Yes |  |
| Update Fields | collection | The updated visible name of the file | No |  |

### Resource: lead

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | The title of the lead | Yes |  |
| Person ID | number | ID of the person this lead will be associated with | No |  |
| Organization ID | number | ID of the organization this lead will be associated with | No |  |
| Additional Fields | collection | Additional lead information | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | number | ID of the lead to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | number | ID of the lead to get | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | number | ID of the lead to update | Yes |  |
| Update Fields | collection | Fields to update for the lead | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection | Filters to apply to lead retrieval | No |  |

### Resource: note

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Content | string | The content of the note | Yes |  |
| Deal ID | number | ID of the deal this note will be associated with | No |  |
| Person ID | number | ID of the person this note will be associated with | No |  |
| Organization ID | number | ID of the organization this note will be associated with | No |  |
| Additional Fields | collection | Additional note information | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | number | ID of the note to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | number | ID of the note to get | Yes |  |

#### Operation: getMany

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all notes or limit results | No | false |
| Limit | number | Maximum number of notes to return | No | 100 |
| Filters | collection | Filters to apply to note retrieval | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection | Filters to apply to note retrieval | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | number | ID of the note to update | Yes |  |
| Update Fields | collection | Fields to update for the note | No |  |

### Resource: organization

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The name of the organization | Yes |  |
| Additional Fields | collection | Additional organization information | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | number | ID of the organization to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | number | ID of the organization to get | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection | Filters to apply to organization retrieval | No |  |

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Term | string | Search term for organizations | Yes |  |
| Additional Fields | collection | Additional search parameters | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | number | ID of the organization to update | Yes |  |
| Update Fields | collection | Fields to update for the organization | No |  |

### Resource: person

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The name of the person | Yes |  |
| Additional Fields | collection | Additional person information | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Person ID | number | ID of the person to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Person ID | number | ID of the person to get | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection | Filters to apply to person retrieval | No |  |

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Term | string | Search term for persons | Yes |  |
| Additional Fields | collection | Additional search parameters | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Person ID | number | ID of the person to update | Yes |  |
| Update Fields | collection | Fields to update for the person | No |  |

### Resource: product

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The name of the product | Yes |  |
| Code | string | Product code/SKU | No |  |
| Unit | string | Unit of measurement | No |  |
| Tax | number | Tax percentage | No |  |
| Active Flag | boolean | Whether the product is active | No | true |
| Additional Fields | collection | Additional product information | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | number | ID of the product to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | number | ID of the product to get | Yes |  |

#### Operation: getMany

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all products or limit results | No | false |
| Limit | number | Maximum number of products to return | No | 100 |
| Filters | collection | Filters to apply to product retrieval | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | number | ID of the product to update | Yes |  |
| Update Fields | collection | Fields to update for the product | No |  |

## Resource Coverage

### Comprehensive CRM Resources
- **Activity**: Task and activity management with scheduling and tracking
- **Deal**: Sales opportunity management through pipeline stages
- **Deal Product**: Product association and pricing for deals
- **File**: Document and attachment management
- **Lead**: Lead capture and qualification processes
- **Note**: Communication and note tracking across entities
- **Organization**: Company and organization management
- **Person**: Contact and person relationship management
- **Product**: Product catalog and pricing management

### Standard CRUD Operations
Most resources support standard operations:
- **Create**: Add new records with validation
- **Get**: Retrieve individual records by ID
- **Get All**: List multiple records with filtering
- **Update**: Modify existing records
- **Delete**: Remove records (where applicable)
- **Search**: Find records using search terms (select resources)

## Authentication Integration

### Dual Authentication Support
- **API Token**: Direct token-based authentication for automated workflows
- **OAuth2**: Secure OAuth2 flow for user-based authentication
- **Dynamic Selection**: Runtime choice between authentication methods
- **Credential Flexibility**: Support for different authentication scenarios

### Security Features
- **Token Management**: Secure token storage and rotation
- **Scope Control**: Granular permission control through OAuth2 scopes
- **Session Handling**: Proper session management for OAuth2 flows
- **Error Handling**: Clear authentication error messages and guidance

## UseCases

- **Sales Pipeline Automation**: Automate deal progression through pipeline stages with activity tracking and notifications
- **Lead Management**: Capture leads from multiple sources and qualify them through automated workflows
- **Customer Relationship Management**: Maintain comprehensive customer profiles with notes, activities, and communication history
- **Sales Reporting**: Generate sales reports and analytics by extracting deal, activity, and performance data
- **Quote and Proposal Generation**: Create quotes and proposals by combining deal information with product data
- **Follow-up Automation**: Automate follow-up activities and reminders based on deal stages and customer interactions
- **Data Synchronization**: Sync Pipedrive data with other business systems like marketing automation and accounting platforms
- **Team Performance Tracking**: Monitor sales team performance through activity tracking and deal progression metrics
- **Customer Onboarding**: Automate customer onboarding processes using organization and person data
- **Integration Workflows**: Connect Pipedrive with email marketing, support systems, and other business tools

