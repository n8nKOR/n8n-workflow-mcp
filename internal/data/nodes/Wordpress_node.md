# Wordpress Node

## Overview

Consume Wordpress API

## Credentials

- Name: wordpressApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: post

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | The title for the post | Yes |  |
| Additional Fields | collection | The ID for the author of the object. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | Unique identifier for the object | Yes |  |
| Update Fields | collection | The ID for the author of the object. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | Unique identifier for the object | Yes |  |
| Options | collection | The password for the post if it is password protected | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Limit response to posts published after a given ISO8601 compliant date | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | Unique identifier for the object | Yes |  |
| Options | collection | Whether to bypass trash and force deletion | No |  |

### Resource: page

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | The title for the page | Yes |  |
| Additional Fields | collection | The ID for the author of the object. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Page ID | string | Unique identifier for the object | Yes |  |
| Update Fields | collection | The ID for the author of the object. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Page ID | string | Unique identifier for the object | Yes |  |
| Options | collection | The password for the page if it is password protected | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Limit response to pages published after a given ISO8601 compliant date | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Page ID | string | Unique identifier for the object | Yes |  |
| Options | collection | Whether to bypass trash and force deletion | No |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | string | Login name for the user | Yes |  |
| Name | string | Display name for the user | Yes |  |
| First Name | string | First name for the user | Yes |  |
| Last Name | string | Last name for the user | Yes |  |
| Email | string | The email address for the user | Yes |  |
| Password | string | Password for the user (never included) | Yes |  |
| Additional Fields | collection | URL of the user | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | Unique identifier for the user | Yes |  |
| Update Fields | collection | Login name for the user | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | Unique identifier for the user | Yes |  |
| Options | collection | Scope under which the request is made; determines fields present in response | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Scope under which the request is made; determines fields present in response | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Reassign | string | Reassign the deleted user | Yes |  |

## UseCases

- [AI-Powered Blog Content Creation] : Automatically create WordPress blog posts with AI-generated content including titles, subtitles, chapters, and featured images based on user-provided keywords
- [Content Management Automation] : Streamline content publishing workflows by automatically creating draft posts that can be reviewed and published by content teams
- [Multi-Channel Content Distribution] : Synchronize content across multiple platforms by creating WordPress posts alongside social media content for consistent brand messaging
- [SEO-Optimized Content Generation] : Generate WordPress posts with AI-optimized content structure, meta descriptions, and keyword integration for better search engine visibility
- [Automated Content Scheduling] : Create and schedule WordPress posts based on content calendars, trending topics, or automated content generation workflows

