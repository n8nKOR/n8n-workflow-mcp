# HubSpot Node

## Overview

Consume HubSpot API

## Credentials

- Name: hubspotApi, Required: Yes (for API Key authentication)
- Name: hubspotAppToken, Required: Yes (for APP Token authentication)
- Name: hubspotOAuth2Api, Required: Yes (for OAuth2 authentication)

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Contact

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Properties | Collection | Contact properties to set | No |
| Associations | Collection | Associate contact with companies or deals | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact to retrieve | Yes |
| Properties | Multi Options | Properties to retrieve | No |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Properties | Multi Options | Properties to retrieve | No |

#### Operation: Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| Filter Groups | Collection | Groups of filters to apply | No |
| Properties | Multi Options | Properties to retrieve | No |
| Sorts | Collection | Sort criteria | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact to update | Yes |
| Properties | Collection | Properties to update | No |

### Resource: Company

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Properties | Collection | Company properties to set | No |
| Associations | Collection | Associate company with contacts or deals | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Company ID | String | ID of the company to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Company ID | String | ID of the company to retrieve | Yes |
| Properties | Multi Options | Properties to retrieve | No |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Properties | Multi Options | Properties to retrieve | No |

#### Operation: Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| Filter Groups | Collection | Groups of filters to apply | No |
| Properties | Multi Options | Properties to retrieve | No |
| Sorts | Collection | Sort criteria | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Company ID | String | ID of the company to update | Yes |
| Properties | Collection | Properties to update | No |

### Resource: Deal

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Properties | Collection | Deal properties to set | No |
| Associations | Collection | Associate deal with contacts or companies | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | String | ID of the deal to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | String | ID of the deal to retrieve | Yes |
| Properties | Multi Options | Properties to retrieve | No |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Properties | Multi Options | Properties to retrieve | No |

#### Operation: Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| Filter Groups | Collection | Groups of filters to apply | No |
| Properties | Multi Options | Properties to retrieve | No |
| Sorts | Collection | Sort criteria | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | String | ID of the deal to update | Yes |
| Properties | Collection | Properties to update | No |

### Resource: Ticket

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Properties | Collection | Ticket properties to set | No |
| Associations | Collection | Associate ticket with contacts, companies, or deals | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Ticket ID | String | ID of the ticket to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Ticket ID | String | ID of the ticket to retrieve | Yes |
| Properties | Multi Options | Properties to retrieve | No |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Properties | Multi Options | Properties to retrieve | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Ticket ID | String | ID of the ticket to update | Yes |
| Properties | Collection | Properties to update | No |

## UseCases

- **CRM Automation** : Automatically sync customer data between HubSpot and other business systems
- **Lead Management** : Capture and nurture leads through automated workflows and scoring
- **Sales Pipeline Management** : Track deals, manage sales stages, and automate follow-up activities
- **Customer Support** : Manage support tickets and track customer service interactions
- **Marketing Automation** : Create and manage marketing campaigns, contact lists, and engagement tracking
- **Data Integration** : Integrate HubSpot with e-commerce, ERP, and other business applications

## Resources Available

### Contact
Manage HubSpot contacts and customer data

**Operations:**
- **Create**: Create a new contact
- **Delete**: Delete a contact
- **Get**: Retrieve a specific contact
- **Get All**: List all contacts
- **Search**: Search contacts
- **Update**: Update contact information

### Company
Manage companies and organizations

**Operations:**
- **Create**: Create a new company
- **Delete**: Delete a company
- **Get**: Retrieve a specific company
- **Get All**: List all companies
- **Search**: Search companies
- **Update**: Update company information

### Contact List
Manage contact lists and segments

**Operations:**
- **Add**: Add contacts to a list
- **Get All**: List all contact lists
- **Remove**: Remove contacts from a list

### Deal
Manage sales deals and opportunities

**Operations:**
- **Create**: Create a new deal
- **Delete**: Delete a deal
- **Get**: Retrieve a specific deal
- **Get All**: List all deals
- **Get Recently Created**: Get recently created deals
- **Get Recently Modified**: Get recently modified deals
- **Search**: Search deals
- **Update**: Update deal information

### Engagement
Manage communication activities

**Operations:**
- **Create**: Create a new engagement
- **Delete**: Delete an engagement
- **Get**: Retrieve a specific engagement
- **Get All**: List all engagements

**Engagement Types:**
- **Call**: Phone call activities
- **Email**: Email communications
- **Meeting**: Meeting activities
- **Task**: Task assignments

### Ticket
Manage customer service tickets

**Operations:**
- **Create**: Create a new ticket
- **Delete**: Delete a ticket
- **Get**: Retrieve a specific ticket
- **Get All**: List all tickets
- **Update**: Update ticket information

## Dynamic Property Loading

The node provides extensive dynamic loading for HubSpot properties:

### Contact Properties
- **Standard Properties**: First name, last name, email, phone, etc.
- **Custom Properties**: Organization-specific contact fields
- **Lead Status**: Available lead status options
- **Lifecycle Stages**: Customer lifecycle stages
- **Original Sources**: Lead source tracking

### Company Properties
- **Standard Properties**: Name, domain, industry, employee count
- **Custom Properties**: Organization-specific company fields
- **Industries**: Available industry classifications
- **Company Types**: Business type categories
- **Lead Status**: Company lead status options

### Deal Properties
- **Standard Properties**: Deal name, amount, close date
- **Custom Properties**: Organization-specific deal fields
- **Pipelines**: Available sales pipelines
- **Stages**: Pipeline-specific deal stages
- **Deal Types**: Deal categorization

### Ticket Properties
- **Standard Properties**: Subject, content, priority, status
- **Custom Properties**: Organization-specific ticket fields
- **Categories**: Ticket categorization
- **Priorities**: Priority levels
- **Sources**: Ticket source channels

## Search Capabilities

### List Search Methods
- **Search Companies**: Find companies with autocomplete
- **Search Contacts**: Find contacts with autocomplete  
- **Search Deals**: Find deals with autocomplete
- **Search Engagements**: Find engagement activities
- **Search Tickets**: Find tickets with autocomplete
- **Search Owners**: Find HubSpot users and owners

## Associations Management

### Association Types
- **Contact to Company**: Link contacts to companies
- **Contact to Deal**: Associate contacts with deals
- **Company to Deal**: Link companies to deals
- **Deal to Ticket**: Connect deals to support tickets

### Association Operations
- **Create Associations**: Link records together
- **Remove Associations**: Unlink related records
- **Get Associations**: Retrieve associated records

## Related Nodes

- **HTTP Request**: For custom HubSpot API calls
- **JSON**: For processing HubSpot API responses
- **Code**: For custom data transformation logic
- **Schedule Trigger**: For automated HubSpot operations

