# AgileCrm Node

## Overview

The AgileCrm node allows you to consume the Agile CRM API to manage contacts, companies, and deals within your Agile CRM account. Agile CRM is a comprehensive customer relationship management platform that combines sales, marketing, and service automation. This node enables you to integrate Agile CRM data and operations into your n8n workflows.

## Credentials

This node requires AgileCrm API credentials:
- **Domain**: Your Agile CRM domain (e.g., yourcompany.agilecrm.com)
- **Email**: Your Agile CRM account email
- **API Key**: Your Agile CRM API key

To obtain API credentials:
1. Log into your Agile CRM account
2. Navigate to Admin Settings → API & Analytics → API Key
3. Generate or copy your existing API key

## Inputs

- **Main**: JSON data for creating/updating records

## Outputs

- **Main**: Returns JSON data from Agile CRM API responses

## Properties

### Resource: Contact

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| JSON Parameters | Boolean | Whether to define contact data as JSON | No |
| Additional Fields (JSON) | String | Contact data in JSON format | No |
| First Name | String | Contact's first name | No |
| Last Name | String | Contact's last name | No |
| Email | String | Contact's email address | No |
| Phone | String | Contact's phone number | No |
| Tags | String | Comma-separated list of tags | No |
| Star Value | Number | Star rating for the contact (1-5) | No |
| Lead Score | Number | Lead score for the contact | No |
| Additional Fields | Collection | Additional contact properties | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | Unique identifier for the contact | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return (default: 20) | No |
| Filter | Options | How to filter results (None, Build Manually, JSON) | No |
| Must Match | Options | Whether filters should match any or all conditions | No |
| Simplify | Boolean | Whether to return simplified response | No |
| Filters | Fixed Collection | Conditions for filtering contacts | No |
| Sort | Collection | Sorting options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact to update | Yes |
| JSON Parameters | Boolean | Whether to define update data as JSON | No |
| Additional Fields (JSON) | String | Update data in JSON format | No |
| Update Fields | Collection | Contact properties to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | ID of the contact to delete | Yes |

### Resource: Company

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| JSON Parameters | Boolean | Whether to define company data as JSON | No |
| Additional Fields (JSON) | String | Company data in JSON format | No |
| Name | String | Company name | No |
| URL | String | Company website URL | No |
| Phone | String | Company phone number | No |
| Tags | String | Comma-separated list of tags | No |
| Star Value | Number | Star rating for the company (1-5) | No |
| Additional Fields | Collection | Additional company properties | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Company ID | String | Unique identifier for the company | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return (default: 20) | No |
| Filter | Options | How to filter results (None, Build Manually, JSON) | No |
| Must Match | Options | Whether filters should match any or all conditions | No |
| Simplify | Boolean | Whether to return simplified response | No |
| Filters | Fixed Collection | Conditions for filtering companies | No |
| Sort | Collection | Sorting options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Company ID | String | ID of the company to update | Yes |
| JSON Parameters | Boolean | Whether to define update data as JSON | No |
| Additional Fields (JSON) | String | Update data in JSON format | No |
| Update Fields | Collection | Company properties to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Company ID | String | ID of the company to delete | Yes |

### Resource: Deal

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| JSON Parameters | Boolean | Whether to define deal data as JSON | No |
| Additional Fields (JSON) | String | Deal data in JSON format | No |
| Name | String | Deal name | No |
| Expected Value | Number | Expected value of the deal | No |
| Milestone | String | Deal milestone/stage | No |
| Probability | Number | Probability of closing (0-100) | No |
| Close Date | DateTime | Expected close date | No |
| Contact IDs | String | Comma-separated list of contact IDs | No |
| Additional Fields | Collection | Additional deal properties | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | String | Unique identifier for the deal | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return (default: 20) | No |
| Filter | Options | How to filter results (None, Build Manually, JSON) | No |
| Must Match | Options | Whether filters should match any or all conditions | No |
| Simplify | Boolean | Whether to return simplified response | No |
| Filters | Fixed Collection | Conditions for filtering deals | No |
| Sort | Collection | Sorting options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | String | ID of the deal to update | Yes |
| JSON Parameters | Boolean | Whether to define update data as JSON | No |
| Additional Fields (JSON) | String | Update data in JSON format | No |
| Update Fields | Collection | Deal properties to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Deal ID | String | ID of the deal to delete | Yes |

## UseCases

- **Lead Management** : Automatically create and update contact records from web forms and landing pages
- **Sales Pipeline Automation** : Track and manage deals through various sales stages and milestones
- **Customer Data Synchronization** : Sync customer information between Agile CRM and other business systems
- **Sales Reporting** : Generate automated sales reports and dashboards from CRM data
- **Company Profile Management** : Maintain comprehensive company profiles with contact relationships
- **Deal Tracking** : Monitor deal progress, probabilities, and expected close dates
- **Contact Enrichment** : Enhance contact records with additional data from external sources
- **Marketing Automation** : Trigger marketing campaigns based on contact behavior and lead scores
- **Customer Service Integration** : Connect customer service activities with CRM contact records
- **Sales Performance Analytics** : Analyze sales team performance and deal conversion rates

