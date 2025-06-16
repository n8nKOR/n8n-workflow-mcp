# Clearbit Node

## Overview

Consume Clearbit API to enrich data about people and companies. Provides social information, company details, and person enrichment capabilities for lead generation and data enhancement.

## Credentials

- Name: clearbitApi, Required: Yes
- API Key: Clearbit API key for authentication
- Authentication Method: API key-based authentication

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: company

#### Operation: enrich

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Domain | string | The domain to look up | Yes |  |
| Additional Fields | collection | The name of the company | No |  |

#### Operation: autocomplete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name is the partial name of the company | Yes |  |

### Resource: person

#### Operation: enrich

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | The email address to look up | Yes |  |
| Additional Fields | collection | The name of the person's employer | No |  |

## UseCases

- **Lead Enrichment** : Enrich lead data with company and contact information
- **Customer Data Enhancement** : Enhance customer profiles with additional business intelligence
- **Sales Prospecting** : Research prospects and companies for sales outreach
- **Marketing Personalization** : Personalize marketing campaigns with enriched data
- **CRM Data Enrichment** : Automatically enrich CRM records with company and contact data
- **Account-Based Marketing** : Support account-based marketing with detailed company insights
- **Lead Scoring** : Improve lead scoring with enriched company and contact information
- **Market Research** : Conduct market research and competitive analysis
- **Data Quality Improvement** : Improve data quality by filling in missing information
- **Sales Intelligence** : Provide sales teams with actionable intelligence about prospects

