# Ghost Node

## Overview

Consume Ghost API

## Credentials

- Name: ghostAdminApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: post

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Post | Yes |  |
| Content Format | options | The format of the post | No |  |
| Content | string | The content of the post to create | No |  |
| Content (JSON) | json | Mobiledoc is the raw JSON format that Ghost uses to store post contents. <a href= | No |  |
| Content (JSON) | json | Lexical is the JSON format returned by the Ghost Default editor | No |  |
| Additional Fields | collection | Choose from the list, or specify IDs using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | The ID of the post to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| By | options | Get the post either by slug or ID | Yes |  |
| Identifier | string | The ID or slug of the post to get | Yes |  |
| Options | collection | Limit the fields returned in the response object. E.g. for posts fields=title,url. | No |  |
| Options | collection | Limit the fields returned in the response object. E.g. for posts fields=title,url. | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Tells the API to return additional data related to the resource you have requested | No |  |
| Options | collection | Tells the API to return additional data related to the resource you have requested | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | The ID of the post to update | No |  |
| Content Format | options | The format of the post | No |  |
| Update Fields | collection | Choose from the list, or specify IDs using an <a href= | No |  |

## UseCases

- **Content Publishing Automation**: Automate blog post creation, scheduling, and publication workflows for content marketing
- **Editorial Workflow Management**: Streamline content creation processes with automated draft management and publishing schedules
- **Cross-platform Content Distribution**: Automatically publish content to Ghost and sync with other platforms and social media
- **SEO and Metadata Management**: Automate SEO optimization, meta descriptions, and featured image assignments for published content
- **Content Migration**: Migrate content from other platforms to Ghost with automated post creation and formatting
- **Newsletter Integration**: Combine Ghost publishing with email marketing campaigns for subscriber engagement
- **Multi-author Collaboration**: Manage collaborative content creation with automated author assignments and publication workflows
- **Content Analytics Integration**: Connect Ghost content with analytics platforms for performance tracking and insights
- **Backup and Archival**: Automate content backup processes and maintain content archives for business continuity
- **Dynamic Content Generation**: Create automated content based on data sources, APIs, or user-generated content

