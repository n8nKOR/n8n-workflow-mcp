# Lemlist Node

## Overview

Consume the Lemlist API

## Credentials

- Name: lemlistApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Activity

#### Operation: Get All
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Return all activities or limit results | No | false |
| Limit | number | Maximum number of activities to return | No | 50 |
| Filters | collection | Filter activities by various criteria | No | - |

### Resource: Campaign

#### Operation: Get All
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Return all campaigns or limit results | No | false |
| Limit | number | Maximum number of campaigns to return | No | 50 |
| Filters | collection | Filter campaigns by various criteria | No | - |

#### Operation: Get Statistics
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign ID | options | Campaign to get statistics for | Yes | - |
| Start Date | dateTime | Statistics start date | Yes | - |
| End Date | dateTime | Statistics end date | Yes | - |
| Timezone | string | Timezone for date calculations | Yes | - |

### Resource: Lead

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign | options | Campaign to add lead to | Yes | - |
| Email | string | Lead email address | Yes | - |
| First Name | string | Lead first name | No | - |
| Last Name | string | Lead last name | No | - |
| Company Name | string | Lead company name | No | - |

#### Operation: Get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | string | ID of the lead to retrieve | Yes | - |

#### Operation: Get All
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Return all leads or limit results | No | false |
| Limit | number | Maximum number of leads to return | No | 50 |
| Campaign | options | Filter by campaign | No | - |

#### Operation: Update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | string | ID of the lead to update | Yes | - |
| First Name | string | Lead first name | No | - |
| Last Name | string | Lead last name | No | - |
| Company Name | string | Lead company name | No | - |

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Lead ID | string | ID of the lead to delete | Yes | - |

### Resource: Enrichment

#### Operation: Enrich Email
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | Email address to enrich | Yes | - |

#### Operation: Enrich Company
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company Name | string | Company name to enrich | Yes | - |
| Domain | string | Company domain | No | - |

### Resource: Team

#### Operation: Get Info
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team ID | string | ID of the team | No | - |

### Resource: Unsubscribe

#### Operation: Add
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | Email address to unsubscribe | Yes | - |

## UseCases

- **Lead Generation**: Automate lead creation and management in campaigns
- **Email Campaign Management**: Monitor campaign performance and statistics
- **Data Enrichment**: Enhance contact and company information
- **Team Collaboration**: Manage team access and permissions
- **Unsubscribe Management**: Handle opt-out requests automatically

