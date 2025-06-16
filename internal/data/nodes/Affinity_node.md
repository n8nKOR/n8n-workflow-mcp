# Affinity Node

## Overview

The Affinity node allows you to consume the Affinity API to manage relationship intelligence data. Affinity is a CRM platform that uses relationship intelligence to help teams leverage their network and close more deals. This node enables you to manage lists, list entries, organizations, and persons within your Affinity workspace through n8n workflows.

## Credentials

This node requires Affinity API credentials:
- **API Key**: Your Affinity API key

To obtain API credentials:
1. Log into your Affinity account
2. Navigate to Settings → Integrations → API
3. Generate an API key
4. Copy the key for use in n8n

## Inputs

- **Main**: JSON data for creating/updating records

## Outputs

- **Main**: Returns JSON data from Affinity API responses

## Properties

### Resource: List

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| List ID | String | ID of the list to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

### Resource: List Entry

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| List ID | String | ID of the list to add entry to | Yes |
| Entity ID | String | ID of the entity (person or organization) | Yes |
| Additional Fields | Collection | Additional entry properties | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| List ID | String | ID of the list | Yes |
| List Entry ID | String | ID of the list entry to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| List ID | String | ID of the list | Yes |
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| List ID | String | ID of the list | Yes |
| List Entry ID | String | ID of the list entry to update | Yes |
| Update Fields | Collection | Entry properties to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| List ID | String | ID of the list | Yes |
| List Entry ID | String | ID of the list entry to delete | Yes |

### Resource: Organization

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Name of the organization | Yes |
| Domain | String | Domain of the organization | No |
| Additional Fields | Collection | Additional organization properties | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Organization ID | String | ID of the organization to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Additional Options | Collection | Filter and search options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Organization ID | String | ID of the organization to update | Yes |
| Update Fields | Collection | Organization properties to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Organization ID | String | ID of the organization to delete | Yes |

### Resource: Person

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| First Name | String | First name of the person | Yes |
| Last Name | String | Last name of the person | Yes |
| Email Addresses | Fixed Collection | Email addresses for the person | No |
| Additional Fields | Collection | Additional person properties | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |
| Additional Options | Collection | Filter and search options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person to update | Yes |
| Update Fields | Collection | Person properties to update | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person to delete | Yes |

## UseCases

- **Relationship Intelligence** : Leverage Affinity's relationship mapping to identify warm introductions and connections
- **Deal Flow Management** : Track and manage investment opportunities and deal pipeline in venture capital
- **Network Analysis** : Analyze professional networks and relationship strength for business development
- **CRM Enhancement** : Enhance traditional CRM data with relationship intelligence and network insights
- **Investment Research** : Research potential investments using relationship data and network connections
- **Partnership Development** : Identify potential partners through mutual connections and relationship mapping
- **Sales Prospecting** : Use relationship intelligence to improve sales prospecting and lead qualification
- **Fundraising Support** : Leverage network connections for fundraising and investor relations activities
- **Business Development** : Identify business opportunities through relationship analysis and network mapping
- **Strategic Planning** : Use relationship data for strategic planning and market entry decisions

