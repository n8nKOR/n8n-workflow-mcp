# Keap Node

## Overview

Consume Keap API

## Credentials

- Name: keapOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: company

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company Name | string |  | Yes |  |
| Additional Fields | collection |  | No |  |
| Addresses | fixedCollection | ISO Alpha-3 Code | No |  |
| Faxes | fixedCollection |  | No |  |
| Phones | fixedCollection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Company name to query on | No |  |

### Resource: contact

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Duplicate Option | options | Performs duplicate checking by one of the following options: Email, EmailAndName. If a match is found using the option provided, the existing contact will be updated. | Yes |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |
| Addresses | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |
| Emails | fixedCollection |  | No |  |
| Faxes | fixedCollection |  | No |  |
| Phones | fixedCollection |  | No |  |
| Social Accounts | fixedCollection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |
| Options | collection | Comma-delimited list of Contact properties to include in the response. (Some fields such as lead_source_id, custom_fields, and job_title aren | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Attribute to order items by | No |  |

### Resource: contactNote

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User Name or ID | options | The infusionsoft user to create the note on behalf of. Choose from the list, or specify an ID using an <a href= | No |  |
| Contact ID | string |  | No |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | string |  | Yes |  |
| Additional Fields | collection | The infusionsoft user to create the note on behalf of. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: contactTag

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |
| Tag Names or IDs | multiOptions | Choose from the list, or specify IDs using an <a href= | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |
| Tag IDs | string |  | Yes | "\"Tag IDs" |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: ecommerceOrder

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string |  | Yes |  |
| Order Date | dateTime |  | Yes |  |
| Order Title | string |  | Yes |  |
| Order Type | options |  | Yes |  |
| Additional Fields | collection | Uses multiple strings separated by comma as promo codes. The corresponding discount will be applied to the order. | No |  |
| Shipping Address | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |
| Order Items | fixedCollection | Overridable price of the product, if not specified, the default will be used | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Order ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Order ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Date to start searching from | No |  |

### Resource: ecommerceProduct

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product Name | string |  | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

### Resource: email

#### Operation: createRecord

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sent To Address | string |  | Yes |  |
| Sent From Address | string |  | Yes |  |
| Additional Fields | collection | Base64 encoded HTML | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Emails sent since the provided date, must be present if untilDate is provided | No |  |

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User Name or ID | options | The infusionsoft user to send the email on behalf of. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Contact IDs | string | Contact IDs to receive the email. Multiple can be added seperated by comma. | No |  |
| Subject | string | The subject line of the email | No |  |
| Additional Fields | collection | Email field of each Contact record to address the email to, such as  | No |  |
| Attachments | fixedCollection | The content of the attachment, encoded in Base64 | No |  |

### Resource: file

#### Operation: upload

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Binary File | boolean | Whether the data to upload should be taken from binary field | No |  |
| Input Binary Field | string |  | Yes |  |
| File Association | options |  | Yes |  |
| Contact ID | string |  | Yes |  |
| File Name | string | The filename of the attached file, including extension | Yes |  |
| File Data | string | The content of the attachment, encoded in Base64 | Yes |  |
| Is Public | boolean |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Filter based on Contact ID, if user has permission to see Contact files | No |  |

## UseCases

- [Contact Management] : Create, update, and manage customer contact information in CRM system
- [Lead Tracking] : Track and manage sales leads through the customer acquisition funnel
- [Sales Pipeline] : Manage sales opportunities and track deal progression through stages
- [Email Marketing] : Send targeted email campaigns and track engagement metrics
- [Customer Segmentation] : Organize contacts into targeted groups based on behavior and demographics
- [E-commerce Integration] : Sync online store orders and customer data with CRM system
- [Support Ticket Management] : Create and track customer support tickets linked to contact records
- [Marketing Automation] : Set up automated email sequences based on customer actions and triggers
- [Sales Reporting] : Generate sales reports and analytics for business performance tracking
- [Contact Tagging] : Apply tags to contacts for organization and automated workflow triggers
- [Product Management] : Manage product catalog and inventory information for sales teams
- [Order Processing] : Process and track customer orders through fulfillment workflow
- [Customer Onboarding] : Automate new customer welcome sequences and setup processes
- [Data Synchronization] : Sync customer data between Keap and other business systems
- [Campaign Analytics] : Track and analyze marketing campaign performance and ROI
- [File Management] : Store and organize customer documents and attachments in CRM
- [Note Taking] : Add and manage customer interaction notes for relationship tracking
- [Company Records] : Manage business-to-business contact and company information
- [Follow-up Automation] : Set up automated follow-up sequences for sales and marketing
- [Customer Lifecycle] : Track customers through entire lifecycle from lead to retention

