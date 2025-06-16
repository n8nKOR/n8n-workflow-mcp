# Automizy Node

## Overview

The Automizy node allows you to consume the Automizy API for email marketing automation. Automizy was a marketing automation platform that helped businesses create and manage email campaigns with automation workflows. **Note: This service may no longer exist and will be removed from n8n in a future release.**

## Credentials

This node requires Automizy API credentials:
- **API Key**: Your Automizy API key

To obtain API credentials:
1. Log into your Automizy account
2. Navigate to Settings → API
3. Generate or copy your API key
4. Use the API key in n8n

## Inputs

- **Main**: JSON data for creating/updating contacts and managing lists

## Outputs

- **Main**: Returns JSON data from Automizy API responses

## Properties

### Resource: Contact

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| List ID | Options | List to add the contact to | Yes |
| Email | String | Email address of the contact | Yes |
| Additional Fields | Collection | Optional contact properties | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| List ID | Options | List to get contacts from | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |
| Additional Options | Collection | Filtering and search options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact to update | Yes |
| Update Fields | Collection | Contact properties to update | No |

### Resource: List

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all lists | No |
| Limit | Number | Maximum number of results to return | No |

## UseCases

- **Email Marketing Automation** : Create and manage automated email marketing campaigns and sequences
- **Contact List Management** : Organize and segment contact lists for targeted marketing campaigns
- **Lead Nurturing** : Automate lead nurturing workflows with personalized email sequences
- **Customer Onboarding** : Create automated onboarding email series for new customers
- **Event-based Messaging** : Trigger emails based on customer behavior and engagement
- **Newsletter Management** : Manage newsletter subscriptions and automated delivery
- **Customer Retention** : Create retention campaigns to re-engage inactive customers
- **Sales Follow-up** : Automate sales follow-up sequences and lead qualification
- **E-commerce Marketing** : Send cart abandonment, product recommendations, and promotional emails
- **Webinar Promotion** : Automate webinar registration confirmations and follow-up sequences

