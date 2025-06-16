# ActiveCampaign Node

## Overview

The ActiveCampaign node allows you to create and edit data in ActiveCampaign, a comprehensive email marketing and customer experience automation platform. This node enables you to manage contacts, accounts, deals, lists, tags, and e-commerce data within your n8n workflows.

## Credentials

This node requires ActiveCampaign API credentials:
- **API URL**: Your ActiveCampaign instance URL (e.g., https://youraccountname.api-us1.com)
- **API Key**: Your ActiveCampaign API key

To obtain API credentials:
1. Log into your ActiveCampaign account
2. Navigate to Settings → Developer
3. Copy your API URL and API Key

## Inputs

- **Main**: JSON data for creating/updating records

## Outputs

- **Main**: Returns JSON data from ActiveCampaign API responses

## Properties

### Resource: Account

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Name of the account | Yes |
| Additional Fields | Collection | Optional additional account data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Account ID | Number | ID of the account to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Account ID | Number | ID of the account to update | Yes |
| Update Fields | Collection | Account data to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Account ID | Number | ID of the account to delete | Yes |

### Resource: Account Contact

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Account ID | Number | ID of the account | Yes |
| Contact ID | Number | ID of the contact | Yes |
| Additional Fields | Collection | Optional additional data | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Account Contact ID | Number | ID of the account contact to delete | Yes |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Account Contact ID | Number | ID of the account contact to update | Yes |
| Update Fields | Collection | Data to update | No |

### Resource: Connection

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Service | String | Name of the service | Yes |
| External ID | String | External ID for the connection | Yes |
| Additional Fields | Collection | Optional additional connection data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Connection ID | Number | ID of the connection to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Connection ID | Number | ID of the connection to update | Yes |
| Update Fields | Collection | Connection data to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Connection ID | Number | ID of the connection to delete | Yes |

### Resource: Contact

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | String | Email address of the contact | Yes |
| Update if Exists | Boolean | Whether to update if contact already exists | No |
| Additional Fields | Collection | Optional contact data (first name, last name, phone, custom fields) | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | Number | ID of the contact to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Additional Options | Collection | Filters and sorting options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | Number | ID of the contact to update | Yes |
| Update Fields | Collection | Contact data to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | Number | ID of the contact to delete | Yes |

### Resource: Contact List

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | Number | ID of the contact | Yes |
| List ID | Number | ID of the list | Yes |

#### Operation: Remove
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | Number | ID of the contact | Yes |
| List ID | Number | ID of the list | Yes |

### Resource: Contact Tag

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | Number | ID of the contact | Yes |
| Tag ID | Number | ID of the tag | Yes |

#### Operation: Remove
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | Number | ID of the contact | Yes |
| Tag ID | Number | ID of the tag | Yes |

### Resource: Deal

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Title | String | Title of the deal | Yes |
| Contact ID | Number | ID of the contact | Yes |
| Value | Number | Value of the deal | Yes |
| Currency | String | Currency of the deal | Yes |
| Additional Fields | Collection | Optional deal data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | Number | ID of the deal to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | Number | ID of the deal to update | Yes |
| Update Fields | Collection | Deal data to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | Number | ID of the deal to delete | Yes |

### Resource: E-Commerce Customer

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Connection ID | Number | ID of the connection | Yes |
| External ID | String | External customer ID | Yes |
| Email | String | Customer email address | Yes |
| Additional Fields | Collection | Optional customer data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Customer ID | Number | ID of the e-commerce customer to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Customer ID | Number | ID of the customer to update | Yes |
| Update Fields | Collection | Customer data to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Customer ID | Number | ID of the customer to delete | Yes |

### Resource: E-Commerce Order

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| External ID | String | External order ID | Yes |
| Source | Number | Source of the order | Yes |
| Customer ID | Number | ID of the customer | Yes |
| Currency | String | Order currency | Yes |
| Connection ID | Number | ID of the connection | Yes |
| Total Price | Number | Total price of the order | Yes |
| Additional Fields | Collection | Optional order data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Order ID | Number | ID of the e-commerce order to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Order ID | Number | ID of the order to update | Yes |
| Update Fields | Collection | Order data to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Order ID | Number | ID of the order to delete | Yes |

### Resource: E-Commerce Order Product

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Order ID | Number | ID of the order | Yes |
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

### Resource: List

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

### Resource: Tag

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Name of the tag | Yes |
| Tag Type | String | Type of the tag | Yes |
| Description | String | Description of the tag | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Tag ID | Number | ID of the tag to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Tag ID | Number | ID of the tag to update | Yes |
| Update Fields | Collection | Tag data to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Tag ID | Number | ID of the tag to delete | Yes |

## UseCases

- [Email Marketing] : Create and send targeted email campaigns to segmented audiences
- [Marketing Automation] : Set up automated marketing workflows based on customer behavior
- [Lead Nurturing] : Nurture leads through automated email sequences and content delivery
- [Customer Segmentation] : Segment customers based on behavior, demographics, and preferences
- [Sales Pipeline Management] : Track and manage sales opportunities through CRM integration
- [Contact Management] : Manage customer contacts and their interaction history
- [Behavioral Tracking] : Track customer behavior on websites and applications for personalization
- [Event-Triggered Campaigns] : Trigger marketing campaigns based on specific customer actions
- [A/B Testing] : Test different marketing messages and campaigns for optimization
- [Customer Lifecycle Marketing] : Create marketing campaigns for different stages of customer lifecycle
- [E-commerce Integration] : Integrate with e-commerce platforms for abandoned cart recovery
- [Tag-Based Automation] : Use tags to automate customer journey and engagement workflows
- [Performance Analytics] : Analyze campaign performance and customer engagement metrics
- [Personalization] : Deliver personalized content and offers based on customer data
- [Integration Workflows] : Connect ActiveCampaign with other business tools and platforms

