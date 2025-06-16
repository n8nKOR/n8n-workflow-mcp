# Zoho CRM

## Overview

The Zoho CRM node provides comprehensive integration with Zoho CRM, enabling full customer relationship management automation through n8n workflows. This node supports complete CRUD operations across all major CRM resources including Leads, Contacts, Accounts, Deals, Products, and more.

## Credentials

Zoho CRM requires OAuth2 authentication:

- **Credential Name**: `zohoOAuth2Api`
- **Required**: Yes
- **Authentication Type**: OAuth2 with Client ID, Client Secret, and appropriate CRM scopes

## Inputs

Zoho CRM operations accept input data from previous workflow nodes for CRM record processing.

## Outputs

Zoho CRM operations return CRM data, operation results, and record information to subsequent workflow nodes.

## Properties

### Resource: CRM Records

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | CRM resource type (Account, Contact, Deal, etc.) | Yes |
| Custom Fields | Collection | Resource-specific fields and values | Varies |
| Additional Fields | Collection | Extended properties for the record | No |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | CRM resource type (Account, Contact, Deal, etc.) | Yes |
| Record ID | String | Unique record identifier | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | CRM resource type (Account, Contact, Deal, etc.) | Yes |
| Return All | Boolean | Return all records or limit results | No |
| Limit | Number | Maximum records to return | No |
| Fields | Multi-Options | Specific fields to retrieve | No |
| Sort Field | Options | Field to sort by | No |
| Sort Order | Options | Ascending or descending order | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | CRM resource type (Account, Contact, Deal, etc.) | Yes |
| Record ID | String | Record identifier to update | Yes |
| Update Fields | Collection | Fields to modify with new values | Yes |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | CRM resource type (Account, Contact, Deal, etc.) | Yes |
| Record ID | String | Record identifier to delete | Yes |

#### Operation: Upsert

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | CRM resource type (Account, Contact, Deal, etc.) | Yes |
| Duplicate Check Fields | Multi-Options | Fields to check for existing records | No |
| Custom Fields | Collection | Record data for creation or update | Yes |

## UseCases

- **Lead Management**: Automate lead capture, qualification, and nurturing processes
- **Sales Pipeline Automation**: Manage deals through sales stages with automated actions
- **Customer Onboarding**: Streamline new customer setup and onboarding workflows
- **Quote-to-Cash**: Automate the entire sales process from quote to payment
- **Contact Management**: Maintain comprehensive customer contact databases
- **Marketing Automation**: Integrate marketing campaigns with CRM data
- **Support Case Management**: Create and manage customer support cases
- **Account Management**: Track and manage customer account relationships
- **Product Catalog Management**: Maintain product information and pricing
- **Vendor Management**: Track vendor relationships and purchase orders
- **Invoice Processing**: Automate invoice creation and management
- **Sales Reporting**: Generate sales reports and analytics
- **Data Synchronization**: Sync CRM data with other business systems
- **Workflow Automation**: Create custom business process automation