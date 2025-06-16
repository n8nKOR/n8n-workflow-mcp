# Storyblok Node

## Overview

Consume Storyblok API

## Credentials

- Name: storyblokContentApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: story

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Identifier | string | The ID or slug of the story to get | Yes |  |
| Space Name or ID | options | The name of the space. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Story ID | string | Numeric ID of the story | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Filter by slug | No |  |
| Space Name or ID | options | The name of the space. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Filter by slug | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Space Name or ID | options | The name of the space. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Story ID | string | Numeric ID of the story | Yes |  |

#### Operation: publish

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Space Name or ID | options | The name of the space. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Story ID | string | Numeric ID of the story | Yes |  |
| Options | collection | Numeric ID of release | No |  |

#### Operation: unpublish

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Space Name or ID | options | The name of the space. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Story ID | string | Numeric ID of the story | Yes |  |

## UseCases

- [Content Management] : Manage website and application content through headless CMS architecture
- [Multi-Channel Publishing] : Publish content across websites, mobile apps, and digital platforms
- [Editorial Workflows] : Create editorial workflows for content creation, review, and approval
- [Website Development] : Build dynamic websites with component-based content management
- [E-commerce Content] : Manage product catalogs and e-commerce content for online stores
- [Blog and News Sites] : Create and manage blog posts and news articles with rich media
- [Landing Page Creation] : Build and optimize landing pages for marketing campaigns
- [Localization] : Manage multilingual content and localization workflows
- [Content Automation] : Automate content publishing and synchronization across platforms
- [Digital Marketing] : Create and manage marketing content and campaign materials
- [API-First Development] : Build applications using Storyblok's API-first approach
- [Component Libraries] : Create reusable content components and design systems
- [Content Personalization] : Deliver personalized content experiences based on user data
- [SEO Optimization] : Optimize content for search engines with structured data and metadata
- [Content Analytics] : Track and analyze content performance across different channels

