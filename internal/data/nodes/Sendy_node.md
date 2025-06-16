# Sendy Node

## Overview

Consume Sendy API

## Credentials

- Name: sendyApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: campaign

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| From Name | string | The  | No |  |
| From Email | string | The  | No |  |
| Reply To | string | The  | No |  |
| Title | string | The  | No |  |
| Subject | string | The  | No |  |
| HTML Text | string | The  | No |  |
| Send Campaign | boolean | Whether to send the campaign as well and not just create a draft. Default is false. | No |  |
| Brand ID | string |  | Yes |  |
| Additional Fields | collection | Lists to exclude from your campaign. List IDs should be single or comma-separated. | No |  |

### Resource: subscriber

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | Email address of the subscriber | No |  |
| List ID | string | The list ID you want to subscribe a user to. This encrypted & hashed ID can be found under View all lists section named ID. | No |  |
| Additional Fields | collection | User | No |  |

#### Operation: count

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| List ID | string | The list ID you want to subscribe a user to. This encrypted & hashed ID can be found under View all lists section named ID. | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | Email address of the subscriber | No |  |
| List ID | string | The list ID you want to subscribe a user to. This encrypted & hashed ID can be found under View all lists section named ID. | No |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | Email address of the subscriber | No |  |
| List ID | string | The list ID you want to subscribe a user to. This encrypted & hashed ID can be found under View all lists section named ID. | No |  |

#### Operation: status

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | Email address of the subscriber | No |  |
| List ID | string | The list ID you want to subscribe a user to. This encrypted & hashed ID can be found under View all lists section named ID. | No |  |

## UseCases

- [Email Marketing] : Send marketing emails and newsletters to subscriber lists
- [List Management] : Manage subscriber lists and segmentation for targeted campaigns
- [Campaign Automation] : Automate email campaigns based on subscriber behavior and triggers
- [Subscriber Management] : Add, update, and remove subscribers from mailing lists
- [Newsletter Distribution] : Distribute regular newsletters to engaged subscriber base
- [Transactional Emails] : Send transactional emails like confirmations and notifications
- [A/B Testing] : Test different email content and subject lines for optimization
- [Analytics and Reporting] : Track email performance metrics and engagement statistics
- [Integration Workflows] : Integrate with e-commerce and CRM systems for automated emails
- [Drip Campaigns] : Create automated drip email sequences for lead nurturing
- [Event-Based Emails] : Send emails triggered by specific user actions or events
- [Personalization] : Deliver personalized email content based on subscriber data
- [Cost-Effective Marketing] : Leverage affordable email marketing through Amazon SES integration
- [Compliance Management] : Manage unsubscribes and maintain email compliance standards
- [Multi-Brand Campaigns] : Manage email campaigns for multiple brands or clients

