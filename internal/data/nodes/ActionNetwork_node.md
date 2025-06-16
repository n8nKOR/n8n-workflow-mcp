# ActionNetwork Node

## Overview

The ActionNetwork node allows you to consume the Action Network API to manage people, events, petitions, signatures, attendances, and tags within the Action Network platform. This node enables you to integrate Action Network's organizing tools into your n8n workflows.

## Credentials

This node requires ActionNetwork API credentials:
- **API Key**: Your Action Network API key

To obtain API credentials:
1. Log into your Action Network account
2. Navigate to your account settings
3. Generate an API key for integration use

## Inputs

- **Main**: JSON data for creating/updating records

## Outputs

- **Main**: Returns JSON data from Action Network API responses

## Properties

### Resource: Attendance

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person to create an attendance for | Yes |
| Event ID | String | ID of the event to create an attendance for | Yes |

#### Operation: Get  
| Field Name | Type | Description | Required |
|---|---|---|---|
| Event ID | String | ID of the event whose attendance to retrieve | Yes |
| Attendance ID | String | ID of the attendance to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Event ID | String | ID of the event to get attendances for | Yes |
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return (default: 50) | No |

### Resource: Event

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Origin System | String | System where event originated | Yes |
| Title | String | Title of the event | Yes |
| Additional Fields | Collection | Optional additional event data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Event ID | String | ID of the event to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Event ID | String | ID of the event to update | Yes |
| Update Fields | Collection | Event data to update | No |

### Resource: Person

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Email Addresses | Collection | Person's email addresses | No |
| Given Name | String | Person's first name | No |
| Family Name | String | Person's last name | No |
| Additional Fields | Collection | Optional additional person data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person to update | Yes |
| Update Fields | Collection | Person data to update | No |

### Resource: Person Tag

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person to tag | Yes |
| Tag ID | String | ID of the tag to add | Yes |

#### Operation: Remove
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person to untag | Yes |
| Tag ID | String | ID of the tag to remove | Yes |

### Resource: Petition

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Origin System | String | System where petition originated | Yes |
| Title | String | Title of the petition | Yes |
| Additional Fields | Collection | Optional additional petition data | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Petition ID | String | ID of the petition to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Petition ID | String | ID of the petition to update | Yes |
| Update Fields | Collection | Petition data to update | No |

### Resource: Signature

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Person ID | String | ID of the person signing the petition | Yes |
| Petition ID | String | ID of the petition to sign | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Petition ID | String | ID of the petition | Yes |
| Signature ID | String | ID of the signature to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Petition ID | String | ID of the petition to get signatures for | Yes |
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Petition ID | String | ID of the petition | Yes |
| Signature ID | String | ID of the signature to update | Yes |
| Update Fields | Collection | Signature data to update | No |

### Resource: Tag

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Name of the tag | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Tag ID | String | ID of the tag to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No |
| Limit | Number | Max number of results to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Tag ID | String | ID of the tag to update | Yes |
| Update Fields | Collection | Tag data to update | No |

## UseCases

- **Political Campaign Management** : Manage voter outreach, volunteer coordination, and campaign events
- **Grassroots Organizing** : Coordinate grassroots advocacy campaigns and community organizing efforts
- **Nonprofit Fundraising** : Automate donation processing and donor management for nonprofit organizations
- **Event Management** : Organize and manage political rallies, town halls, and community events
- **Volunteer Coordination** : Recruit, manage, and coordinate volunteers for campaigns and causes
- **Advocacy Campaigns** : Run issue-based advocacy campaigns with petition and action management
- **Voter Registration** : Facilitate voter registration drives and civic engagement initiatives
- **Email Marketing for Causes** : Send targeted email campaigns to supporters and constituents
- **Donation Processing** : Process and track political donations and fundraising activities
- **Community Engagement** : Build and engage communities around political and social causes

