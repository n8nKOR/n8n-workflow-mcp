# Hunter Node

## Overview

Consume Hunter API

## Credentials

- Name: hunterApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Hunter API

#### Operation: Domain Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | String | Domain name to search for email addresses | Yes |
| Company | String | Company name to search | No |
| Limit | Number | Maximum number of results to return | No |
| Offset | Number | Number of results to skip | No |

#### Operation: Email Finder

| Field Name | Type | Description | Required |
|---|---|---|---|
| Domain | String | Domain name | Yes |
| First Name | String | First name of the person | Yes |
| Last Name | String | Last name of the person | Yes |

#### Operation: Email Verifier

| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | String | Email address to verify | Yes |

#### Operation: Account Information

| Field Name | Type | Description | Required |
|---|---|---|---|
| (No additional fields) | - | Returns account information | - |

## UseCases

- [Email Discovery] : Find professional email addresses for business contacts and leads
- [Lead Generation] : Discover contact information for potential customers and clients
- [Sales Prospecting] : Build prospect lists with verified email addresses for outreach
- [Email Verification] : Verify the deliverability and validity of email addresses
- [Contact Enrichment] : Enhance existing contact databases with additional email information
- [Recruitment] : Find contact information for potential job candidates and talent acquisition
- [Business Development] : Locate decision makers and key contacts at target companies
- [Marketing Campaigns] : Build targeted email lists for marketing and promotional campaigns
- [Account-Based Marketing] : Find contacts at specific target accounts for personalized outreach
- [Journalist Outreach] : Discover media contacts and journalists for PR and media relations
- [Partnership Development] : Find contacts at potential partner organizations and companies
- [Customer Research] : Research and verify customer contact information for support
- [Data Cleansing] : Clean and verify existing email databases for better deliverability
- [Competitive Intelligence] : Research contact information at competitor organizations
- [Event Networking] : Find contact information for networking and event follow-up activities
