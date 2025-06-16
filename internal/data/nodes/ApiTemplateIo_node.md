# ApiTemplateIo Node

## Overview

The APITemplate.io node allows you to consume the APITemplate.io API to generate images and PDFs from templates. APITemplate.io is a service that enables automated creation of visual content such as social media images, certificates, invoices, and other documents using predefined templates with dynamic data insertion.

## Credentials

This node requires APITemplate.io API credentials:
- **API Key**: Your APITemplate.io API key

To obtain API credentials:
1. Sign up for an account at APITemplate.io
2. Navigate to your account settings
3. Generate an API key
4. Copy the API key for use in n8n

## Inputs

- **Main**: JSON data to be used for template rendering

## Outputs

- **Main**: Returns generated content URL or binary data

## Properties

### Resource: Account

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| - | - | Retrieves account information | No |

### Resource: Image

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Template ID | Options | ID of the image template to use | Yes |
| JSON Parameters | Boolean | Whether to define parameters as JSON | No |
| Download | Boolean | Whether to download the file as binary data | No |
| Put Output File in Field | String | Name of the binary property field (when download is enabled) | No |
| Overrides (JSON) | JSON | Template parameters as JSON object | No |
| Properties | Collection | Individual template properties | No |

### Resource: PDF

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Template ID | Options | ID of the PDF template to use | Yes |
| JSON Parameters | Boolean | Whether to define parameters as JSON | No |
| Download | Boolean | Whether to download the file as binary data | No |
| Put Output File in Field | String | Name of the binary property field (when download is enabled) | No |
| Overrides (JSON) | JSON | Template parameters as JSON object | No |
| Properties | Collection | Individual template properties | No |

## UseCases

- **Automated Report Generation** : Generate PDF reports and documents from dynamic data sources
- **Invoice and Receipt Creation** : Create professional invoices and receipts from transaction data
- **Certificate Generation** : Generate certificates and credentials for courses and achievements
- **Marketing Material Creation** : Create personalized marketing materials and promotional documents
- **Contract and Agreement Generation** : Generate contracts and legal documents from template data
- **Label and Badge Creation** : Create labels, badges, and identification materials
- **Presentation Automation** : Generate presentations and slides from structured data
- **Document Personalization** : Create personalized documents for customers and clients
- **Bulk Document Processing** : Generate large volumes of documents from batch data
- **Template-Based Publishing** : Automate publishing workflows using document templates

