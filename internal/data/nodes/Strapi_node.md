# Strapi Node

## Overview

Consume Strapi API

## Credentials

- Name: strapiApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: entry

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Content Type | string | Name of the content type | Yes |  |
| Columns | string | Comma-separated list of the properties which should used as columns for the new rows | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Content Type | string | Name of the content type | Yes |  |
| Entry ID | string | The ID of the entry to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Content Type | string | Name of the content type | Yes |  |
| Entry ID | string | The ID of the entry to get | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Content Type | string | Name of the content type | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Only select entries matching the publication state provided | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Content Type | string | Name of the content type | Yes |  |
| Update Key | string | Name of the property which decides which rows in the database should be updated. Normally that would be  | Yes |  |
| Columns | string | Comma-separated list of the properties which should used as columns for the new rows | No |  |

## UseCases

- **Headless CMS Integration** : Integrate Strapi CMS content with frontend applications and websites
- **Content Management Automation** : Automate content creation, updates, and publishing workflows
- **Multi-Channel Publishing** : Distribute content across multiple platforms and channels simultaneously
- **Content Synchronization** : Sync content between Strapi and other content management systems
- **API-First Content Delivery** : Deliver content through APIs to mobile apps and web applications
- **Content Workflow Automation** : Automate editorial workflows and content approval processes
- **E-commerce Content Management** : Manage product catalogs and e-commerce content through Strapi
- **Blog and News Automation** : Automate blog post creation and news content distribution
- **Multilingual Content Management** : Manage and distribute multilingual content across platforms
- **Content Analytics Integration** : Connect content performance data with analytics and reporting systems

