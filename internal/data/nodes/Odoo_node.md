# Odoo Node

## Overview

⚠️ **Operation Coverage**: This documentation includes create, get, getAll, update, and delete operations for custom resources. The implementation supports the complete CRUD operations including delete and get operations as documented.

Comprehensive integration with Odoo ERP system for managing business operations, customer relationships, inventory, and custom modules. Odoo is an open-source enterprise resource planning (ERP) platform that provides integrated business applications including CRM, sales, inventory, accounting, and project management.

## Credentials

- **Name**: odooApi
- **Required**: Yes

### Credential Configuration
To configure Odoo API credentials:
1. Access your Odoo instance
2. Go to Settings → Technical → Database Structure → External API
3. Generate an API key or use database credentials
4. Configure the database URL, username, and password/API key
5. Ensure API access is enabled for your user account

## Inputs

- **Main**: JSON data for creating, updating, or filtering Odoo records

## Outputs

- **Main**: Returns JSON responses with Odoo record data, operation results, and metadata

## Properties

### Resource: custom

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resource | String | Name of the custom Odoo model/resource | Yes | - |
| Fields | Collection | Field values for the new record | No | - |

**Available Fields (varies by resource):**
- Field values depend on the selected Odoo model
- Common fields include name, description, date fields
- Relationship fields for linking to other records

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resource | String | Name of the custom Odoo model/resource | Yes | - |
| Record ID | Number | ID of the specific record to retrieve | Yes | - |
| Fields | Array | Specific fields to retrieve | No | [] |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resource | String | Name of the custom Odoo model/resource | Yes | - |
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Filters | Collection | Domain filters for filtering records | No | - |

**Filter Options:**
- **Field**: Field name to filter on
- **Operator**: Comparison operator (=, !=, >, <, like, etc.)
- **Value**: Value to compare against

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resource | String | Name of the custom Odoo model/resource | Yes | - |
| Record ID | Number | ID of the record to update | Yes | - |
| Update Fields | Collection | Fields to update with new values | No | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resource | String | Name of the custom Odoo model/resource | Yes | - |
| Record ID | Number | ID of the record to delete | Yes | - |

### Resource: opportunity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | String | Name of the opportunity | Yes | - |
| Partner ID | Number | Customer/partner associated with opportunity | No | - |
| Expected Revenue | Number | Expected revenue amount | No | 0 |
| Probability | Number | Probability of closing (0-100) | No | 10 |
| Additional Fields | Collection | Additional opportunity fields | No | - |

**Additional Fields Options:**
- **Description**: Detailed description of the opportunity
- **Close Date**: Expected closing date
- **Sales Team**: Assigned sales team
- **Salesperson**: Assigned salesperson
- **Tags**: Opportunity tags and categories
- **Priority**: Opportunity priority level

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | Number | ID of the opportunity to retrieve | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Filters | Collection | Domain filters for opportunities | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | Number | ID of the opportunity to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | Number | ID of the opportunity to delete | Yes | - |

### Resource: contact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | String | Contact name | Yes | - |
| Email | String | Contact email address | No | - |
| Phone | String | Contact phone number | No | - |
| Is Company | Boolean | Whether this contact is a company | No | false |
| Additional Fields | Collection | Additional contact information | No | - |

**Additional Fields Options:**
- **Street Address**: Street address
- **City**: City name
- **Country**: Country selection
- **Website**: Website URL
- **Category**: Contact categories and tags
- **Parent Company**: Parent company if applicable

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | Number | ID of the contact to retrieve | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Filters | Collection | Domain filters for contacts | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | Number | ID of the contact to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | Number | ID of the contact to delete | Yes | - |

### Resource: note

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Memo | String | Note content | Yes | - |
| Subject | String | Note subject/title | No | - |
| Related Model | String | Related model name (e.g., res.partner) | No | - |
| Related ID | Number | ID of the related record | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | Number | ID of the note to retrieve | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Filters | Collection | Domain filters for notes | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | Number | ID of the note to update | Yes | - |
| Memo | String | Updated note content | Yes | - |
| Subject | String | Updated note subject | No | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | Number | ID of the note to delete | Yes | - |

## UseCases

- **ERP Integration**: Integrate Odoo ERP data with other business systems and applications for unified business operations
- **Customer Relationship Management**: Manage customer data, sales processes, and opportunity tracking through Odoo CRM
- **Inventory Management**: Synchronize inventory data between Odoo and external systems for real-time stock management
- **Sales Order Processing**: Automate sales order creation, processing, and fulfillment workflows
- **Financial Data Integration**: Integrate accounting and financial data with reporting systems and business intelligence tools
- **E-commerce Synchronization**: Sync e-commerce data between Odoo and online stores for unified product and order management
- **Project Management**: Manage projects, tasks, and resource allocation through Odoo project management modules
- **Human Resources**: Integrate HR data and employee management processes for payroll and workforce management
- **Manufacturing Operations**: Coordinate manufacturing processes, production planning, and supply chain management
- **Business Intelligence**: Extract Odoo data for business intelligence, analytics, and executive reporting dashboards
- **Custom Module Integration**: Integrate with custom Odoo modules and third-party applications for specialized business needs
- **Data Migration**: Migrate data between Odoo instances or from legacy systems to Odoo
- **Automated Reporting**: Generate automated reports combining Odoo data with external data sources
- **Workflow Automation**: Automate business processes that span multiple Odoo modules and external systems
- **Multi-Company Operations**: Manage data across multiple companies and subsidiaries within Odoo
- **Partner Portal Integration**: Integrate partner and customer portals with external communication systems
- **Quality Management**: Automate quality control processes and compliance tracking through Odoo QMS modules
- **Asset Management**: Track and manage company assets, maintenance schedules, and depreciation
- **Document Management**: Integrate Odoo's document management with external storage and collaboration platforms
- **API Development**: Build custom applications and integrations using Odoo's extensive API capabilities

