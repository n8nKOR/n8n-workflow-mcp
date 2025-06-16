# Autopilot Node

## Overview

The Autopilot node allows you to consume the Autopilot API for marketing automation. Autopilot is a visual marketing automation platform that helps businesses engage customers through personalized journeys. This node enables you to manage contacts, lists, contact journeys, and automated marketing workflows within your n8n automation.

## Credentials

This node requires Autopilot API credentials:
- **API Key**: Your Autopilot API key

To obtain API credentials:
1. Log into your Autopilot account
2. Navigate to Settings → Account → API
3. Generate or copy your API key
4. Use the API key in n8n

## Inputs

- **Main**: JSON data for creating/updating contacts and managing lists

## Outputs

- **Main**: Returns JSON data from Autopilot API responses

## Properties

### Resource: Contact

#### Operation: Upsert
| Field Name | Type | Description | Required |
|---|---|---|---|
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
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

### Resource: Contact Journey

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact | Yes |
| Trigger ID | Options | Journey trigger to add contact to | Yes |

### Resource: Contact List

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact | Yes |
| List ID | Options | List to add contact to | Yes |

#### Operation: Remove
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact | Yes |
| List ID | Options | List to remove contact from | Yes |

#### Operation: Check if Exists
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact | Yes |
| List ID | Options | List to check for contact membership | Yes |

### Resource: List

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all lists | No |
| Limit | Number | Maximum number of results to return | No |

## UseCases

- **Marketing Automation** : Create automated marketing journeys and personalized customer experiences
- **Lead Nurturing** : Develop automated lead nurturing sequences based on customer behavior
- **Customer Onboarding** : Build automated onboarding workflows for new customers and users
- **Email Campaign Management** : Manage and automate email marketing campaigns with personalized content
- **Contact Segmentation** : Segment contacts into targeted lists for specific marketing campaigns
- **Behavioral Triggering** : Trigger automated workflows based on customer actions and engagement
- **E-commerce Automation** : Create automated workflows for cart abandonment, upselling, and cross-selling
- **Event-based Marketing** : Automate marketing responses to specific customer events and milestones
- **Re-engagement Campaigns** : Automatically re-engage inactive customers with targeted messaging
- **Customer Journey Mapping** : Design and automate complex customer journey experiences

