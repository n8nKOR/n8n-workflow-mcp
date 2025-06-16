# Emelia Node

## Overview

Consume the Emelia API

## Credentials

- Name: emeliaApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: campaign

#### Operation: addContact

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign Name or ID | options | The ID of the campaign to add the contact to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Contact Email | string | The email of the contact to add to the campaign | Yes |  |
| Additional Fields | collection | Filter by custom fields | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign Name | string | The name of the campaign to create | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign ID | string | The ID of the campaign to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: pause

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign ID | string | The ID of the campaign to pause. The campaign must be in RUNNING mode. | Yes |  |

#### Operation: start

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign ID | string | The ID of the campaign to start. Email provider and contacts must be set. | Yes |  |

#### Operation: duplicate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Campaign Name or ID | options | The ID of the campaign to duplicate. Choose from the list, or specify an ID using an <a href= | Yes |  |
| New Campaign Name | string | The name of the new campaign to create | Yes |  |
| Options | collection | Whether to copy all the contacts from the original campaign | No |  |

### Resource: contactList

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact List Name or ID | options | The ID of the contact list to add the contact to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Contact Email | string | The email of the contact to add to the contact list | Yes |  |
| Additional Fields | collection | Filter by custom fields | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **Cold Email Outreach**: Create and manage email campaigns for cold outreach and lead generation with contact management
- **Email Campaign Automation**: Automate the creation, management, and execution of email marketing campaigns with custom contact lists
- **Sales Prospecting**: Add prospects to targeted email campaigns and track engagement for sales pipeline development
- **Lead Nurturing**: Create multi-touch email sequences for nurturing leads through the sales funnel
- **Contact List Management**: Organize and segment contacts into different lists for targeted email campaigns
- **Campaign Performance Tracking**: Monitor campaign metrics, open rates, and response rates for optimization
- **A/B Testing**: Duplicate campaigns to test different email variations and optimize messaging
- **Automated Follow-up**: Create systematic follow-up campaigns based on prospect responses and engagement
- **CRM Integration**: Sync campaign data and contact information with external CRM systems for unified sales processes
- **Outbound Sales Automation**: Streamline outbound sales processes with automated email sequences and contact management

