# Contentful

## Overview

The Contentful node allows you to consume the Contentful API to retrieve content from your Contentful space. This node supports both Delivery API and Preview API, enabling you to access published content or preview unpublished content for various content management workflows.

## Credentials

- **Credential Name**: `contentfulApi`
- **Required**: Yes
- **Configuration**: Contentful API credentials including space ID and access tokens

## Inputs

- **Main**: Accepts input data for Contentful operations

## Outputs

- **Main**: Returns content data from Contentful API

## Properties

### Resource: Content Management

#### Operation: Contentful API Access

| Field Name | Type | Description | Required |
|---|---|---|---|
| Source | Options | API source (Delivery API or Preview API) | Yes |
| Resource | Options | Type of Contentful resource to work with | Yes |
| Operation | Options | Operation to perform on the resource | Yes |
| Environment ID | String | Contentful environment identifier | Yes |
| Content Type ID | String | Content type identifier | No |
| Entry ID | String | Entry identifier | No |
| Additional Fields | Collection | Additional configuration options | No |

## Resources

- **Asset**: Manage Contentful assets (images, files, etc.)
- **Content Type**: Work with content type definitions
- **Entry**: Retrieve and manage content entries
- **Locale**: Handle localization settings
- **Space**: Access space information

## UseCases

- **Content Retrieval** : Fetch published content from Contentful for websites and applications
- **Content Preview** : Access unpublished content for preview and testing purposes
- **Content Management** : Integrate Contentful content into automated workflows
- **Multi-language Content** : Handle localized content across different languages
- **Asset Management** : Process and distribute media assets from Contentful

## Space Operations

### Get
Retrieve information about a Contentful space.

**Fields**: None (uses space ID from credentials)

## Content Type Operations

### Get
Retrieve a specific content type definition.

**Fields**:
- **Environment ID**
  - Type: String
  - Default: master
  - Description: The ID for the Contentful environment (e.g. master, staging, etc.). Depending on your plan, you might not have environments. In that case use "master".

- **Content Type ID**
  - Type: String
  - Required: Yes
  - Description: The ID of the content type to retrieve

- **Additional Fields**
  - Type: Collection
  - **RAW Data**
    - Type: Boolean
    - Default: false
    - Description: Whether the data should be returned RAW instead of parsed

## Entry Operations

### Get
Retrieve a specific entry by ID.

**Fields**:
- **Environment ID**
  - Type: String
  - Default: master
  - Description: The ID for the Contentful environment (e.g. master, staging, etc.). Depending on your plan, you might not have environments. In that case use "master".

- **Entry ID**
  - Type: String
  - Required: Yes
  - Description: The ID of the entry to retrieve

- **Additional Fields**
  - Type: Collection
  - **RAW Data**
    - Type: Boolean
    - Default: false
    - Description: Whether the data should be returned RAW instead of parsed

### Get Many
Retrieve multiple entries with optional filtering.

**Fields**:
- **Environment ID**
  - Type: String
  - Default: master
  - Description: The ID for the Contentful environment (e.g. master, staging, etc.). Depending on your plan, you might not have environments. In that case use "master".

- **Return All**
  - Type: Boolean
  - Default: false
  - Description: Whether to return all results or only up to a given limit

- **Limit**
  - Type: Number
  - Default: 100
  - Min: 1, Max: 500
  - Description: Max number of results to return (only when Return All is false)

- **Additional Fields**
  - Type: Collection
  - **Content Type ID**
    - Type: String
    - Description: To search for entries with a specific content type
  - **Equal**
    - Type: String
    - Placeholder: fields.title=n8n
    - Description: Search for all data that matches the condition: {attribute}={value}. Attribute can use dot notation.
  - **Not Equal**
    - Type: String
    - Placeholder: fields.title[ne]=n8n
    - Description: Search for all data that matches the condition: {attribute}[ne]={value}. Attribute can use dot notation.
  - **Include**
    - Type: String
    - Placeholder: fields.tags[in]=accessories,flowers
    - Description: Search for all data that matches the condition: {attribute}[in]={value}. Attribute can use dot notation.
  - **Exclude**
    - Type: String
    - Placeholder: fields.tags[nin]=accessories,flowers
    - Description: Search for all data that matches the condition: {attribute}[nin]={value}. Attribute can use dot notation.
  - **Exist**
    - Type: String
    - Placeholder: fields.tags[exists]=true
    - Description: Search for all data that matches the condition: {attribute}[exists]={value}. Attribute can use dot notation.
  - **Fields**
    - Type: String
    - Placeholder: fields.title
    - Description: The select operator allows you to choose what fields to return from an entity. You can choose multiple values by combining comma-separated operators.
  - **Order**
    - Type: String
    - Placeholder: sys.createdAt
    - Description: You can order items in the response by specifying the order search parameter. You can use sys properties (such as sys.createdAt) or field values (such as fields.myCustomDateField) for ordering.
  - **Query**
    - Type: String
    - Description: Full-text search is case insensitive and might return more results than expected. A query will only take values with more than 1 character.
  - **RAW Data**
    - Type: Boolean
    - Default: false
    - Description: Whether the data should be returned RAW instead of parsed

## Asset Operations

### Get
Retrieve a specific asset by ID.

**Fields**:
- **Environment ID**
  - Type: String
  - Default: master
  - Description: The ID for the Contentful environment (e.g. master, staging, etc.). Depending on your plan, you might not have environments. In that case use "master".

- **Asset ID**
  - Type: String
  - Required: Yes
  - Description: The ID of the asset to retrieve

### Get Many
Retrieve multiple assets with optional filtering.

**Fields**:
- **Environment ID**
  - Type: String
  - Default: master
  - Description: The ID for the Contentful environment (e.g. master, staging, etc.). Depending on your plan, you might not have environments. In that case use "master".

- **Return All**
  - Type: Boolean
  - Default: false
  - Description: Whether to return all results or only up to a given limit

- **Limit**
  - Type: Number
  - Default: 100
  - Min: 1, Max: 500
  - Description: Max number of results to return (only when Return All is false)

- **Additional Fields**
  - Type: Collection
  - **Equal**
    - Type: String
    - Placeholder: fields.title=n8n
    - Description: Search for all data that matches the condition: {attribute}={value}. Attribute can use dot notation.
  - **Not Equal**
    - Type: String
    - Placeholder: fields.title[ne]=n8n
    - Description: Search for all data that matches the condition: {attribute}[ne]={value}. Attribute can use dot notation.
  - **Include**
    - Type: String
    - Placeholder: fields.tags[in]=accessories,flowers
    - Description: Search for all data that matches the condition: {attribute}[in]={value}. Attribute can use dot notation.
  - **Exclude**
    - Type: String
    - Placeholder: fields.tags[nin]=accessories,flowers
    - Description: Search for all data that matches the condition: {attribute}[nin]={value}. Attribute can use dot notation.
  - **Exist**
    - Type: String
    - Placeholder: fields.tags[exists]=true
    - Description: Search for all data that matches the condition: {attribute}[exists]={value}. Attribute can use dot notation.
  - **Fields**
    - Type: String
    - Placeholder: fields.title
    - Description: The select operator allows you to choose what fields to return from an entity. You can choose multiple values by combining comma-separated operators.
  - **Order**
    - Type: String
    - Placeholder: sys.createdAt
    - Description: You can order items in the response by specifying the order search parameter. You can use sys properties (such as sys.createdAt) or field values (such as fields.myCustomDateField) for ordering.
  - **Query**
    - Type: String
    - Description: Full-text search is case insensitive and might return more results than expected. A query will only take values with more than 1 character.
  - **RAW Data**
    - Type: Boolean
    - Default: false
    - Description: Whether the data should be returned RAW instead of parsed

## Locale Operations

### Get Many
Retrieve multiple locales from the Contentful space.

**Fields**:
- **Environment ID**
  - Type: String
  - Default: master
  - Description: The ID for the Contentful environment (e.g. master, staging, etc.). Depending on your plan, you might not have environments. In that case use "master".

- **Return All**
  - Type: Boolean
  - Default: false
  - Description: Whether to return all results or only up to a given limit

- **Limit**
  - Type: Number
  - Default: 100
  - Min: 1, Max: 500
  - Description: Max number of results to return (only when Return All is false)

