# Webflow Node

## Overview

Consume the Webflow API

## Credentials

- Name: webflowOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Item

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| site | options | Choose the Webflow site to work with | Yes |
| collection | options | Select the collection within the site | Yes |
| fields | collection | Dynamic fields based on the selected collection structure | Yes |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| site | options | Choose the Webflow site to work with | Yes |
| collection | options | Select the collection within the site | Yes |
| itemId | string | ID of the item to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| site | options | Choose the Webflow site to work with | Yes |
| collection | options | Select the collection within the site | Yes |
| itemId | string | ID of the item to get | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| site | options | Choose the Webflow site to work with | Yes |
| collection | options | Select the collection within the site | Yes |
| filters | collection | Filter items when getting multiple items | No |
| limit | number | Maximum number of items to return | No |
| offset | number | Number of items to skip | No |
| sort | options | Sort order for retrieved items | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| resource | options | The resource to operate on | Yes |
| operation | options | Operation to perform | Yes |
| site | options | Choose the Webflow site to work with | Yes |
| collection | options | Select the collection within the site | Yes |
| itemId | string | ID of the item to update | Yes |
| fields | collection | Dynamic fields based on the selected collection structure | Yes |

## UseCases

- [Content Management]: Manage CMS content in Webflow collections, including blog posts, product catalogs, and portfolio items
- [E-commerce Integration]: Sync product data between external systems and Webflow CMS collections for online stores
- [Dynamic Website Updates]: Automatically update website content based on external data sources or user interactions

