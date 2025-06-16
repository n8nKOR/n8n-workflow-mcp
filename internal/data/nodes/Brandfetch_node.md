# Brandfetch Node

## Overview

The Brandfetch node allows you to consume the Brandfetch API to retrieve company brand assets and information. Brandfetch provides access to logos, colors, fonts, and other brand elements for millions of companies worldwide. This node enables automated brand asset collection and brand analysis workflows.

## Credentials

This node requires Brandfetch API credentials:
- **API Key**: Your Brandfetch API key

To obtain API credentials:
1. Sign up for a Brandfetch account
2. Navigate to API settings
3. Generate an API key
4. Copy the key for use in n8n

## Inputs

- **Main**: JSON data with domain information

## Outputs

- **Main**: Returns brand data or binary files

## Properties

### Resource: Brand Assets

#### Operation: Logo
| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | String | Company domain name (e.g., google.com) | Yes |
| Download | Boolean | Whether to download logo files as binary data | Yes |
| Image Type | MultiOptions | Types of images to retrieve (Logo, Icon) | Conditional |
| Image Format | MultiOptions | Formats to download (PNG, SVG) | Conditional |

#### Operation: Color
| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | String | Company domain name (e.g., google.com) | Yes |

#### Operation: Company
| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | String | Company domain name (e.g., google.com) | Yes |

#### Operation: Font
| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | String | Company domain name (e.g., google.com) | Yes |

#### Operation: Industry
| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | String | Company domain name (e.g., google.com) | Yes |

## API Details

### Base URL
```
https://api.brandfetch.io/v2/
```

### Authentication
- Uses API Key authentication via Bearer token
- API key required for all requests
- Rate limits apply based on subscription plan

### Response Data Structure
- **Logos**: Array of logo objects with different types and formats
- **Colors**: Array of color objects with hex values and usage context
- **Fonts**: Array of font objects with family names and weights
- **Company**: Object containing company information and metadata
- **Industry**: Object containing industry classification data

### Logo Download Options
- **Types**: Logo (primary brand logo), Icon (favicon/symbol)
- **Formats**: PNG (raster), SVG (vector)
- **Binary Data**: Downloaded files stored as n8n binary data properties

## UseCases

- **Brand Asset Management** : Automatically download company logos for marketing materials and presentations
- **Lead Enrichment** : Enhance prospect data with company logos and brand information
- **CRM Enhancement** : Add company logos to CRM contact records and company profiles
- **Content Creation** : Include company logos in case studies, articles, and marketing materials
- **E-commerce Integration** : Display brand logos for product listings and vendor management
- **Design Automation** : Populate design templates with correct brand assets and colors
- **Marketing Collateral** : Create branded materials with accurate company assets
- **Sales Presentations** : Automatically populate client logos in sales pitch decks
- **Website Integration** : Display partner and client logos with consistent branding
- **Analytics Dashboards** : Include company logos in business reports and visualizations

| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | string | Company domain name | Yes |

## Features

- **Logo Retrieval**: Get company logos in multiple formats
- **Brand Colors**: Extract brand color palettes
- **Typography**: Retrieve brand font information
- **Company Data**: Get company information and industry
- **Binary Downloads**: Download logo files directly
- **Multiple Formats**: Support for PNG, SVG formats

## UseCases

- **Brand Analysis**: Analyze competitor brand elements
- **Design Systems**: Build brand asset libraries
- **Marketing Automation**: Standardize brand usage
- **Content Creation**: Access brand assets for materials
- **Brand Monitoring**: Track brand asset changes

## Notes

### Domain Requirements
- Use clean domain names (e.g., "google.com" not "https://www.google.com")
- International domains supported
- Subdomains may have different brand data

### Logo Downloads
- Binary downloads create file attachments
- Multiple formats available per logo
- Icon vs logo types serve different purposes
- File sizes vary by format and resolution

### API Limits
- Rate limiting based on plan
- Monitor usage in Brandfetch dashboard
- Consider caching for frequently accessed brands

