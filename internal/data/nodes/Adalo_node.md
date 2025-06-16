# Adalo Node

## Overview

The Adalo node allows you to consume the Adalo API to interact with collections in your Adalo applications. Adalo is a no-code platform for building native mobile and web apps. This node enables you to create, read, update, and delete records in your Adalo app collections through n8n workflows.

## Credentials

This node requires Adalo API credentials:
- **App ID**: Your Adalo application ID
- **API Key**: Your Adalo API key

To obtain API credentials:
1. Log into your Adalo account
2. Open your application
3. Navigate to the API Documentation section
4. Find your App ID and generate an API key

## Inputs

- **Main**: JSON data for creating/updating collection records

## Outputs

- **Main**: Returns JSON data from Adalo API responses

## Properties

### Resource: Collection

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | String | ID of the collection to create a record in | Yes |
| Data to Send | Options | How to specify the data (Auto-Map Input Data or Define Below) | Yes |
| Fields to Send | Fixed Collection | Field values to send when "Define Below" is selected | No |
| Inputs to Ignore | String | Comma-separated list of input properties to ignore when auto-mapping | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | String | ID of the collection to retrieve a record from | Yes |
| Row ID | String | ID of the record to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | String | ID of the collection to retrieve records from | Yes |
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return (1-100, default: 100) | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | String | ID of the collection containing the record to update | Yes |
| Row ID | String | ID of the record to update | Yes |
| Data to Send | Options | How to specify the data (Auto-Map Input Data or Define Below) | Yes |
| Fields to Send | Fixed Collection | Field values to send when "Define Below" is selected | No |
| Inputs to Ignore | String | Comma-separated list of input properties to ignore when auto-mapping | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | String | ID of the collection containing the record to delete | Yes |
| Row ID | String | ID of the record to delete | Yes |

## UseCases

- **Product Catalog Management** : Sync product data between Adalo app and external inventory systems
- **Order Processing** : Create order records when customers make purchases
- **Customer Registration** : Auto-create customer profiles from form submissions
- **Inventory Updates** : Real-time stock level updates from warehouse management systems
- **Review Management** : Collect and store customer reviews in Adalo collections
- **Blog Post Creation** : Automatically publish blog content from CMS to Adalo app
- **User-Generated Content** : Store and moderate user submissions like photos and comments
- **Event Management** : Create and update event listings from external calendar systems
- **Lead Management** : Import leads from marketing campaigns into Adalo CRM collections
- **Employee Directory** : Maintain up-to-date staff information across systems
- **Task Assignment** : Create and update task records for project management
- **Multi-Platform Sync** : Keep data consistent between Adalo and other business tools
- **Backup Operations** : Regular data exports for backup and compliance purposes
- **Data Migration** : Transfer historical data when upgrading or changing systems
- **Real-time Updates** : Push notifications and data updates to mobile users
- **Webhook Integration** : Process incoming data from external APIs and services

