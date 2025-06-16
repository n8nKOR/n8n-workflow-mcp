# Mautic Node

## Overview

Consume Mautic API

## Credentials

This node supports multiple Mautic authentication methods:
- **mauticApi**: API token authentication
- **mauticOAuth2Api**: OAuth2 authentication for enhanced security

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: campaignContact

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign ID | string | ID of the campaign to add contact to | Yes |  |
| Contact ID | string | ID of the contact to add to campaign | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign ID | string | ID of the campaign to remove contact from | Yes |  |
| Contact ID | string | ID of the contact to remove from campaign | Yes |  |

### Resource: company

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company Name | string | The name of the company to create | No |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |
| Additional Fields | collection | Additional company information | No |  |

#### Additional Company Fields
| Field Name | Type | Description |
|---|---|---|
| Annual Revenue | number | Company's annual revenue |
| Description | string | Company description |
| Industry | string | Company industry |
| Employee Count | number | Number of employees |
| Website | string | Company website URL |
| Phone | string | Company phone number |
| Address | collection | Company address information |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | The ID of the company to update | Yes |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |
| Update Fields | collection | Fields to update with new values | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | The ID of the company to return | Yes |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |
| Additional Fields | collection | Filtering and sorting options | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | The ID of the company to delete | Yes |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |

### Resource: companyContact

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | ID of the company | Yes |  |
| Contact ID | string | ID of the contact to add to company | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Company ID | string | ID of the company | Yes |  |
| Contact ID | string | ID of the contact to remove from company | Yes |  |

### Resource: contact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| JSON Parameters | boolean | Use JSON parameters for complex field configuration | No | false |
| Email | string | Email address of the contact | No |  |
| First Name | string | Contact's first name | No |  |
| Last Name | string | Contact's last name | No |  |
| Primary Company Name or ID | options | Primary company association. Choose from the list, or specify an ID using expression | No |  |
| Position | string | Contact's job position | No |  |
| Title | string | Contact's professional title | No |  |
| Body | json | Contact parameters in JSON format (when JSON Parameters enabled) | No |  |
| Additional Fields | collection | Extended contact information and custom fields | No |  |

#### Contact Additional Fields
| Field Name | Type | Description |
|---|---|---|
| Phone | string | Contact phone number |
| Mobile | string | Contact mobile number |
| Address | collection | Contact address information |
| Social Media | collection | Social media profiles |
| Custom Fields | collection | Custom field values |
| Tags | array | Contact tags |
| Owner | string | Contact owner ID |
| Stage | string | Contact stage ID |
| Points | number | Contact points/score |
| Attribution | string | Attribution information |
| UTM Source | string | UTM source tracking |
| UTM Medium | string | UTM medium tracking |
| UTM Campaign | string | UTM campaign tracking |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to update | Yes |  |
| JSON Parameters | boolean | Use JSON parameters for complex field updates | No | false |
| Update Fields | collection | Contact parameters to update | No |  |

#### Operation: editDoNotContactList

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact | Yes |  |
| Action | options | Action to perform (add/remove) | Yes |  |
| Channel | options | Communication channel (email/sms) | Yes |  |
| Additional Fields | collection | Do Not Contact details | No |  |

#### Operation: editContactPoint

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact | Yes |  |
| Action | options | Action to perform (add/subtract) | Yes |  |
| Points | number | Number of points to add or subtract | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Search and filtering options | No |  |

#### Contact Search Options
| Field Name | Type | Description |
|---|---|---|
| Search | string | Search term for contact lookup |
| Order By | string | Field to order results by |
| Order Direction | options | Sort direction (ASC/DESC) |
| Published Only | boolean | Return only published contacts |
| Minimal | boolean | Return minimal contact data |
| Where | string | SQL-like WHERE clause for filtering |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to delete | Yes |  |

#### Operation: sendEmail

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign Email Name or ID | options | Email campaign to send. Choose from the list, or specify an ID using expression | Yes |  |
| Contact ID | string | ID of the contact to send email to | Yes |  |

### Resource: contactSegment

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Segment ID | string | ID of the segment | Yes |  |
| Contact ID | string | ID of the contact to add to segment | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Segment ID | string | ID of the segment | Yes |  |
| Contact ID | string | ID of the contact to remove from segment | Yes |  |

### Resource: segmentEmail

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Segment Email Name or ID | options | Segment email campaign to send. Choose from the list, or specify an ID using expression | Yes |  |

## UseCases

- **Marketing Automation** : Automate marketing campaigns and lead nurturing workflows
- **Lead Management** : Manage leads and prospects through the sales funnel
- **Email Marketing** : Create and manage email marketing campaigns and sequences
- **Contact Segmentation** : Segment contacts based on behavior and demographics
- **Campaign Analytics** : Track and analyze marketing campaign performance
- **Lead Scoring** : Implement lead scoring systems to prioritize prospects
- **Customer Journey Mapping** : Map and optimize customer journey touchpoints
- **Multi-Channel Marketing** : Coordinate marketing across multiple channels and platforms
- **Personalization** : Deliver personalized marketing content and experiences
- **CRM Integration** : Integrate marketing data with customer relationship management systems

