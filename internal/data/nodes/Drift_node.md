# Drift Node

## Overview

Consume Drift API

## Credentials

- Name: driftApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: contact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | The email of the contact | Yes |  |
| Additional Fields | collection | The name of the contact | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | Unique identifier for the contact | Yes |  |
| Update Fields | collection | The email of the contact | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | Unique identifier for the contact | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | Unique identifier for the contact | Yes |  |

## UseCases

- **Live Chat Contact Management**: Automatically create and update contact records from chat interactions and website visitors
- **Conversational Marketing**: Integrate chat conversations with marketing automation workflows for personalized follow-up campaigns
- **Lead Qualification**: Capture and manage leads generated from chat conversations and website interactions
- **Customer Support Integration**: Sync chat contacts with support ticketing systems for unified customer service
- **Sales Pipeline Automation**: Automatically move qualified chat contacts through sales workflows and CRM systems
- **Chat Analytics and Reporting**: Track and analyze chat contact data for conversion optimization and performance insights
- **Multi-channel Contact Sync**: Maintain consistent contact information across chat, email, and other communication channels
- **Website Visitor Tracking**: Create contact records for website visitors engaging with chat widgets and forms
- **Personalized Chat Experiences**: Use contact data to provide contextual and personalized chat interactions
- **Revenue Attribution**: Track chat-generated leads through the entire customer journey for ROI measurement

