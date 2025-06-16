# Dropcontact Node

## Overview

Find B2B emails and enrich contacts

## Credentials

- Name: dropcontactApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: contact

#### Operation: fetchRequest

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Request ID | string |  | Yes |  |

#### Operation: enrich

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string |  | No |  |
| Simplify Output (Faster) | boolean | When off, waits for the contact data before completing. Waiting time can be adjusted with Extend Wait Time option. When on, returns a request_id that can be used later in the Fetch Request operation. | No |  |
| Additional Fields | collection |  | No |  |
| Options | collection | When not simplifying the response, data will be fetched in two steps. This parameter controls how long to wait (in seconds) before trying the second step. | No |  |

## UseCases

- **B2B Lead Generation**: Automatically find professional email addresses for prospects using their name and company website
- **Sales Prospecting**: Enrich CRM data with verified email addresses and professional contact information for outreach campaigns
- **Contact Database Enrichment**: Enhance existing contact lists with missing professional email addresses and company details
- **Cold Email Campaigns**: Build targeted email lists with verified B2B contacts for cold outreach and lead generation
- **French Market Intelligence**: Leverage specialized French company data including SIREN codes, NAF codes, and company leadership information
- **Marketing Automation**: Automate the process of finding and verifying professional contacts for marketing campaigns
- **Recruitment and Talent Acquisition**: Find contact information for potential candidates and hiring managers at target companies
- **Partnership Development**: Identify and contact key decision-makers at potential partner companies
- **Account-Based Marketing**: Enrich target account data with professional contacts for personalized marketing campaigns
- **Data Cleaning and Validation**: Improve data quality by enriching incomplete contact records with verified professional information

