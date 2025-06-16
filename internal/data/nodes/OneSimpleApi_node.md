# One Simple API Node

## Overview

A comprehensive toolbox of no-code utilities for web automation workflows. One Simple API provides various utility functions including website operations, social media data extraction, image processing, and data validation services, making it ideal for automating common web-related tasks without complex coding.

## Credentials

- Name: oneSimpleApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: website

#### Operation: pdf

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webpage URL | string | Link to webpage to convert to PDF | Yes |  |
| Download PDF? | boolean | Whether to download the PDF or return a link to it | Yes |  |
| Put Output In Field | string | The name of the output field to put the binary file data in | Yes |  |
| Options | collection | Additional options for PDF generation (force refresh, quality settings) | No |  |

#### Operation: screenshot

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webpage URL | string | Link to webpage to take screenshot of | Yes |  |
| Download Screenshot? | boolean | Whether to download the screenshot or return a link to it | Yes |  |
| Put Output In Field | string | The name of the output field to put the binary file data in | Yes |  |
| Options | collection | Screenshot options (force refresh, dimensions, quality) | No |  |

#### Operation: seo

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webpage URL | string | Webpage to get SEO information for | Yes |  |
| Options | collection | SEO analysis options | No |  |

### Resource: utility

#### Operation: qrCode

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| QR Content | string | The text that should be turned into a QR code (website URL, text, etc.) | Yes |  |
| Download Image? | boolean | Whether to download the QR code or return a link to it | Yes |  |
| Put Output In Field | string | The name of the output field to put the binary file data in | Yes |  |
| Options | collection | QR Code customization options (size, error correction) | No |  |

#### Operation: validateEmail

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email Address | string | Email address to validate | Yes |  |

#### Operation: expandURL

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| URL | string | Shortened URL to expand to full URL | Yes |  |

### Resource: socialProfile

#### Operation: instagramProfile

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Profile Name | string | Instagram profile username to get details of | Yes |  |

#### Operation: spotifyArtistProfile

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Artist Name | string | Spotify artist name to get profile details for | Yes |  |

### Resource: information

#### Operation: exchangeRate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Value | string | Monetary value to convert | Yes |  |
| From Currency | string | Source currency code (e.g., USD, EUR) | Yes |  |
| To Currency | string | Target currency code (e.g., USD, EUR) | Yes |  |

#### Operation: imageMetadata

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Link To Image | string | URL of image to extract metadata from | Yes |  |

## UseCases

- **Website Documentation and Archiving**: Generate PDF documents and screenshots of websites for documentation, archiving, and compliance purposes
- **SEO Analysis and Monitoring**: Analyze website SEO metrics, track performance changes, and generate SEO reports for optimization workflows
- **QR Code Generation for Marketing**: Create QR codes for marketing campaigns, event promotions, and digital marketing materials with custom branding
- **Email Validation and Data Cleaning**: Validate email addresses in bulk for mailing lists, user registration systems, and data quality assurance
- **URL Management and Analytics**: Expand shortened URLs to reveal destinations, track link sources, and analyze referral traffic patterns
- **Social Media Intelligence**: Extract Instagram and Spotify profile data for market research, influencer analysis, and competitive intelligence
- **Financial Data Integration**: Perform real-time currency conversions for e-commerce, financial reporting, and international business workflows
- **Image Processing and Analysis**: Extract metadata from images for content management, photo organization, and digital asset management systems
- **Web Scraping and Content Extraction**: Combine website screenshots and SEO data for comprehensive website analysis and monitoring
- **Automated Report Generation**: Create visual reports combining website screenshots, PDFs, and data analysis for client presentations and dashboards

