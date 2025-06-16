# Raindrop Node

## Overview

Consume the Raindrop API

## Credentials

- Name: raindropOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: bookmark

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Collection Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| Link | string | Link of the bookmark to be created | Yes |  |
| Additional Fields | collection | Whether this bookmark is marked as favorite | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bookmark ID | string | The ID of the bookmark to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bookmark ID | string | The ID of the bookmark to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Collection Name or ID | options | The ID of the collection from which to retrieve all bookmarks. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bookmark ID | string | The ID of the bookmark to update | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: collection

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the collection to create | Yes |  |
| Additional Fields | collection | URL of an image to use as cover for the collection | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Collection ID | string | The ID of the collection to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Collection ID | string | The ID of the collection to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Type | options | Root-level collections | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Collection ID | string | The ID of the collection to update | Yes |  |
| Update Fields | collection | Name of the binary property containing the data for the image to upload as a cover | No |  |

### Resource: tag

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tags | string | One or more tags to delete. Enter comma-separated values to delete multiple tags. | Yes |  |
| Additional Fields | collection | It\ | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: user

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Self | boolean | Whether to return details on the logged-in user | Yes |  |
| User ID | string | The ID of the user to retrieve | Yes |  |

## UseCases

- [Bookmark Management] : Organize and manage web bookmarks across devices and browsers
- [Content Curation] : Curate and organize web content for research and reference
- [Knowledge Management] : Build personal knowledge bases with organized web resources
- [Research Organization] : Organize research materials and reference sources
- [Team Collaboration] : Share bookmarks and collections with team members
- [Content Discovery] : Discover and save interesting web content for later reading
- [Reading List Management] : Maintain reading lists and save articles for later
- [Reference Library] : Build reference libraries for projects and topics
- [Link Sharing] : Share useful links and resources with colleagues and friends
- [Cross-Device Sync] : Synchronize bookmarks across multiple devices and platforms
- [Tag-Based Organization] : Use tags to categorize and find bookmarks efficiently
- [Search and Discovery] : Search through saved bookmarks and discover related content
- [Automation Workflows] : Automate bookmark saving and organization workflows
- [Archive Management] : Archive and preserve important web content and resources
- [Integration Management] : Integrate bookmark management with other productivity tools

