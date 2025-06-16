# Intercom Node

## Overview

Consume Intercom API

## Credentials

- Name: intercomApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: user

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The Intercom defined ID representing the Lead | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Company ID representing the user | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Select By | options | The Intercom defined ID representing the Lead | No |  |
| Value | string | View by value | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update By | options | The Intercom defined ID representing the user | No |  |
| Value | string | Value of the property to identify the user to update | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Identifier Type | options | A unique string identifier for the user. It is required on creation if an email is not supplied. | No |  |
| Value | string | Unique string identifier value | Yes |  |

### Resource: lead

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Delete By | options | The Intercom defined ID representing the Lead | No |  |
| Value | string | Delete by value | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Select By | options | Email representing the Lead | No |  |
| Value | string | View by value | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | The email address of the lead | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update By | options | Automatically generated identifier for the Lead | No |  |
| Value | string | Value of the property to identify the lead to update | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | The email of the user | Yes |  |

### Resource: company

#### Operation: users

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List By | options | The Intercom defined ID representing the company | No |  |
| Value | string | View by value | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Segment representing the Lead | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Select By | options | The company_id you have given to the company | No |  |
| Value | string | View by value | Yes |  |

## UseCases

- **Customer Communication Management**: Automate customer interactions, live chat responses, and customer support conversations
- **Lead Generation and Qualification**: Capture, qualify, and nurture leads through automated conversation workflows
- **User and Customer Data Management**: Maintain comprehensive customer profiles with interaction history and behavioral data
- **Sales Pipeline Integration**: Connect customer conversations with sales processes and CRM systems for unified lead management
- **Customer Support Automation**: Streamline customer support with automated ticketing, conversation routing, and response management
- **Onboarding and User Engagement**: Create automated onboarding sequences and user engagement campaigns
- **Company and Account Management**: Organize customers by company, track account relationships, and manage B2B customer interactions
- **Marketing Automation**: Trigger targeted marketing campaigns based on customer behavior and conversation data
- **Customer Success Workflows**: Monitor customer health, track usage patterns, and automate retention campaigns
- **Multi-channel Customer Experience**: Unify customer interactions across chat, email, and in-app messaging for seamless support

