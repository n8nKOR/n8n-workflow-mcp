# Cockpit Node

## Overview

The Cockpit node allows you to interact with Cockpit CMS through its API. Cockpit is a headless content management system that provides APIs for managing collections, forms, and assets. This node enables you to create, read, update, and delete content in Cockpit collections and submit data to Cockpit forms.

## Credentials

- Name: cockpitApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: collection

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Comma-separated list of fields to get | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Entry ID | string |  | Yes |  |

### Resource: form

#### Operation: submit

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| JSON Data Fields | boolean | Whether form fields should be set via the value-key pair UI or JSON | No |  |
| Form Data | json | Form data to send as JSON | No |  |
| Form Data | fixedCollection | Name of the field | No |  |

## UseCases

- **Content Management**: Manage content collections in a headless CMS architecture
- **Form Submissions**: Handle form submissions from websites or applications
- **Data Synchronization**: Sync content between Cockpit and other systems
- **Website Updates**: Automatically update website content through Cockpit collections
- **User-Generated Content**: Process and store user submissions via Cockpit forms
- **Multi-Channel Publishing**: Distribute content across multiple platforms from Cockpit
- **Content Validation**: Validate and process content before publishing
- **Dynamic Content**: Create and update dynamic content for web applications

