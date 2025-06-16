# Salesforce Node

## Overview

Comprehensive integration with Salesforce CRM platform for managing customer relationships, sales processes, and business operations. Salesforce is a leading cloud-based CRM that provides sales automation, marketing automation, customer service, and business analytics capabilities.

## Credentials

This node supports multiple authentication methods:
- **Name**: salesforceOAuth2Api (OAuth2)
- **Name**: salesforceJwtApi (JWT)
- **Required**: Yes

### Authentication Methods

#### OAuth2 Authentication (salesforceOAuth2Api)
- **Method**: OAuth 2.0 authorization flow
- **Client ID**: Your Salesforce connected app client ID
- **Client Secret**: Your Salesforce connected app client secret
- **Authorization URL**: Salesforce authorization endpoint
- **Access Token URL**: Salesforce token endpoint
- **Use Case**: Interactive authentication with user consent

#### JWT Authentication (salesforceJwtApi)
- **Method**: JSON Web Token (JWT) bearer flow
- **Client ID**: Your Salesforce connected app client ID
- **Private Key**: Private key for JWT signing (PEM format)
- **Username**: Salesforce username for JWT flow
- **Subdomain**: Your Salesforce subdomain (e.g., your-company.my.salesforce.com)
- **Use Case**: Server-to-server authentication without user interaction

### Credential Configuration
1. Create a Connected App in Salesforce Setup
2. Configure OAuth settings and permissions
3. For JWT: Upload certificate and enable JWT flow
4. Copy credentials to n8n configuration

## Inputs

- **Main**: JSON data for creating, updating, or querying Salesforce records

## Outputs

- **Main**: Returns JSON responses with Salesforce record data, operation results, and metadata

## Properties

### Resource: account

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | String | Account name | Yes | - |
| Type | Options | Account type (Customer, Partner, etc.) | No | - |
| Industry | String | Industry category | No | - |
| Phone | String | Main phone number | No | - |
| Website | String | Company website | No | - |
| Additional Fields | Collection | Additional account information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Account ID | String | ID of account to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Account ID | String | ID of account to retrieve | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Account ID | String | ID of account to delete | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Options | Collection | Query conditions and filters | No | - |

### Resource: attachment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Parent ID | String | ID of parent record to attach to | Yes | - |
| Name | String | Name of the attachment | Yes | - |
| Body | String | Base64 encoded file content | Yes | - |
| Content Type | String | MIME type of the attachment | No | - |
| Description | String | Description of the attachment | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attachment ID | String | ID of attachment to retrieve | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Parent ID | String | ID of parent record | No | - |
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attachment ID | String | ID of attachment to delete | Yes | - |

### Resource: case

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subject | String | Case subject | Yes | - |
| Status | Options | Case status | No | New |
| Priority | Options | Case priority | No | Medium |
| Origin | Options | Case origin (Web, Phone, Email, etc.) | No | - |
| Account ID | String | Associated account ID | No | - |
| Contact ID | String | Associated contact ID | No | - |
| Additional Fields | Collection | Additional case information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Case ID | String | ID of case to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Case ID | String | ID of case to retrieve | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Case ID | String | ID of case to delete | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Options | Collection | Query conditions and filters | No | - |

### Resource: lead

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Match Against | Options | The field to check to see if the lead already exists | Yes | - |
| Value to Match | String | Value to match against the specified field | Yes | - |
| First Name | String | Lead first name | No | - |
| Last Name | String | Lead last name | Yes | - |
| Company | String | Lead company | Yes | - |
| Email | String | Lead email address | No | - |
| Additional Fields | Collection | Additional lead information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | String | ID of Lead that needs to be updated | Yes | - |
| Update Fields | Collection | Fields to update including revenue, status, etc. | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | String | ID of Lead that needs to be fetched | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | String | ID of Lead that needs to be deleted | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Options | Collection | Query conditions and filters | No | - |

#### Operation: addToCampaign

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | String | ID of lead to add to campaign | Yes | - |
| Campaign Name or ID | Options | ID of the campaign to add the lead to | Yes | - |
| Options | Collection | Campaign member options | No | - |

#### Operation: addNote

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | String | ID of lead to add note to | Yes | - |
| Title | String | Title of the note | Yes | - |
| Body | String | Body of the note. Limited to 32 KB. | No | - |
| Options | Collection | Additional note options | No | - |

### Resource: contact

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Match Against | Options | The field to check to see if the contact already exists | Yes | - |
| Value to Match | String | Value to match against the specified field | Yes | - |
| First Name | String | Contact first name | No | - |
| Last Name | String | Contact last name | Yes | - |
| Email | String | Contact email address | No | - |
| Account ID | String | Associated account ID | No | - |
| Additional Fields | Collection | Additional contact information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | String | ID of contact that needs to be updated | Yes | - |
| Update Fields | Collection | Fields to update including account association, phone, etc. | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | String | ID of contact that needs to be fetched | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | String | ID of contact that needs to be deleted | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Options | Collection | Query conditions and filters | No | - |

#### Operation: addToCampaign

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | String | ID of contact to add to campaign | Yes | - |
| Campaign Name or ID | Options | ID of the campaign to add the contact to | Yes | - |
| Options | Collection | Campaign member options | No | - |

#### Operation: addNote

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | String | ID of contact to add note to | Yes | - |
| Title | String | Title of the note | Yes | - |
| Body | String | Body of the note. Limited to 32 KB. | No | - |
| Options | Collection | Additional note options | No | - |

### Resource: customObject

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Custom Object Name or ID | Options | Name of the custom object | Yes | - |
| Match Against | Options | The field to check to see if the object already exists | Yes | - |
| Value to Match | String | Value to match against the specified field | Yes | - |
| Fields | Collection | Custom object field values | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Custom Object Name or ID | Options | Name of the custom object | Yes | - |
| Record ID | String | Record ID to be updated | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Custom Object Name or ID | Options | Name of the custom object | Yes | - |
| Record ID | String | Record ID to be retrieved | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Custom Object Name or ID | Options | Name of the custom object | Yes | - |
| Record ID | String | Record ID to be deleted | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Custom Object Name or ID | Options | Name of the custom object | Yes | - |
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Options | Collection | Query conditions and filters | No | - |

### Resource: document

#### Operation: upload

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | String | Name of the file | Yes | - |
| Input Binary Field | String | Field containing binary file data | Yes | data |
| Additional Fields | Collection | Document metadata and associations | No | - |

### Resource: flow

#### Operation: invoke

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Flow Name or ID | String | Name or ID of the flow to invoke | Yes | - |
| Input Variables | Collection | Input variables for the flow | No | - |

### Resource: opportunity

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Match Against | Options | The field to check to see if the opportunity already exists | Yes | - |
| Value to Match | String | Value to match against the specified field | Yes | - |
| Name | String | Opportunity name | Yes | - |
| Stage Name | String | Sales stage of the opportunity | Yes | - |
| Close Date | String | Expected close date (YYYY-MM-DD) | Yes | - |
| Account ID | String | Associated account ID | No | - |
| Additional Fields | Collection | Additional opportunity information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | String | ID of opportunity that needs to be updated | Yes | - |
| Update Fields | Collection | Fields to update including amount, stage, etc. | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | String | ID of opportunity that needs to be fetched | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | String | ID of opportunity that needs to be deleted | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Options | Collection | Query conditions and filters | No | - |

#### Operation: addNote

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Opportunity ID | String | ID of opportunity to add note to | Yes | - |
| Title | String | Title of the note | Yes | - |
| Body | String | Body of the note. Limited to 32 KB. | No | - |
| Options | Collection | Additional note options | No | - |

### Resource: search

#### Operation: query

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Query | String | SOSL search query | Yes | - |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subject | String | Task subject | Yes | - |
| Status | Options | Task status | No | Not Started |
| Priority | Options | Task priority | No | Normal |
| Due Date | String | Task due date (YYYY-MM-DD) | No | - |
| What ID | String | Related record ID | No | - |
| Who ID | String | Related contact/lead ID | No | - |
| Additional Fields | Collection | Additional task information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | String | ID of task to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | String | ID of task to retrieve | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | String | ID of task to delete | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Options | Collection | Query conditions and filters | No | - |

### Resource: user

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | String | ID of user to retrieve | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 100 |
| Options | Collection | Query conditions and filters | No | - |

## UseCases

- **CRM Integration**: Integrate Salesforce CRM data with other business systems and applications for unified customer management
- **Lead Management**: Automate lead capture, qualification, and routing processes for sales teams
- **Opportunity Tracking**: Manage sales opportunities, pipeline forecasting, and revenue recognition
- **Customer Service**: Automate case management, ticket routing, and customer support workflows
- **Marketing Automation**: Sync marketing campaign data and lead scoring with Salesforce
- **Data Synchronization**: Keep customer data synchronized between Salesforce and other business applications
- **Reporting and Analytics**: Extract Salesforce data for business intelligence and custom reporting
- **Workflow Automation**: Automate business processes that span Salesforce and external systems
- **Contact Management**: Centralize contact information and communication history across all customer touchpoints
- **Sales Process Automation**: Automate sales activities, follow-ups, and opportunity progression
- **Document Management**: Attach and manage documents related to accounts, contacts, and opportunities
- **Campaign Management**: Track marketing campaign effectiveness and ROI through Salesforce integration
- **Territory Management**: Automate territory assignments and sales team organization
- **Compliance Tracking**: Maintain audit trails and compliance records for regulatory requirements
- **Custom Object Integration**: Work with custom Salesforce objects for industry-specific requirements
- **API Integration**: Connect third-party applications with Salesforce through automated API calls
- **Data Migration**: Migrate data between Salesforce orgs or from legacy systems
- **Real-time Updates**: Provide real-time data updates between Salesforce and external applications
- **Quote and Proposal Generation**: Automate quote generation and proposal processes
- **Customer Journey Tracking**: Track customer interactions across multiple touchpoints and channels

