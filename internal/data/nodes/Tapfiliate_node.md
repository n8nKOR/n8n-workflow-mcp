# Tapfiliate Node

## Overview

Consume Tapfiliate API

## Credentials

- Name: tapfiliateApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: affiliate

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | The affiliate's email | Yes |  |
| First Name | string | The affiliate's firstname | Yes |  |
| Last Name | string | The affiliate's lastname | Yes |  |
| Additional Fields | collection | The country's ISO_3166-1 code. <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Affiliate ID | string | The ID of the affiliate | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Affiliate ID | string | The ID of the affiliate | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Retrieves affiliates for a certain affiliate group | No |  |

### Resource: affiliateMetadata

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Affiliate ID | string | The ID of the affiliate | Yes |  |
| Metadata | fixedCollection | Meta data | No |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Affiliate ID | string | The ID of the affiliate | Yes |  |
| Key | string | Name of the metadata key to remove | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Affiliate ID | string | The ID of the affiliate | Yes |  |
| Key | string | Name of the metadata key to update | No |  |
| Value | string | Value to set for the metadata key | No |  |

### Resource: programAffiliate

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Program Name or ID | options | The ID of the Program to add the affiliate to. This ID can be found as part of the URL when viewing the program on the platform. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Affiliate ID | string | The ID of the affiliate | Yes |  |
| Additional Fields | collection | An optional approval status | No |  |

#### Operation: approve

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Program Name or ID | options | The ID of the Program to add the affiliate to. This ID can be found as part of the URL when viewing the program on the platform. Choose from the list, or specify an ID using an <a href= | No |  |
| Affiliate ID | string | The ID of the affiliate | No |  |

#### Operation: disapprove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Program Name or ID | options | The ID of the Program to add the affiliate to. This ID can be found as part of the URL when viewing the program on the platform. Choose from the list, or specify an ID using an <a href= | No |  |
| Affiliate ID | string | The ID of the affiliate | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Program Name or ID | options | The ID of the Program to add the affiliate to. This ID can be found as part of the URL when viewing the program on the platform. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Affiliate ID | string | The ID of the affiliate | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Program Name or ID | options | The ID of the Program to add the affiliate to. This ID can be found as part of the URL when viewing the program on the platform. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Retrieves affiliates for a certain affiliate group | No |  |

## UseCases

- **Affiliate Program Management** : Manage affiliate programs and partner relationships
- **Commission Tracking** : Track affiliate commissions and payment calculations
- **Partner Onboarding** : Automate affiliate partner onboarding and registration
- **Performance Analytics** : Analyze affiliate performance and conversion metrics
- **Payment Processing** : Automate affiliate payment processing and disbursements
- **Conversion Tracking** : Track conversions and attribute sales to affiliates
- **Campaign Management** : Manage affiliate marketing campaigns and promotions
- **Referral Programs** : Implement and manage customer referral programs
- **Multi-Tier Commissions** : Handle complex multi-tier commission structures
- **Fraud Prevention** : Monitor and prevent affiliate fraud and abuse

