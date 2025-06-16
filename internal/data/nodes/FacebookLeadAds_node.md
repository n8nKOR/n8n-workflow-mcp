# Facebook Lead Ads Trigger

## Overview

The Facebook Lead Ads Trigger node handles Facebook Lead Ads events via webhooks in your n8n workflows. This trigger node listens for new lead submissions from Facebook Lead Ads campaigns, allowing you to automatically process and respond to leads as they are generated.

## Credentials

- **Credential Name**: `facebookLeadAdsOAuth2Api`
- **Required**: Yes
- **Configuration**: Facebook OAuth2 credentials with Lead Ads permissions

## Inputs

This node does not accept input connections as it is a trigger node.

## Outputs

- **Main**: Outputs lead data when a Facebook Lead Ads webhook is triggered

## Properties

### Resource: Webhook Events

#### Operation: Facebook Lead Ads Event Handling

| Field Name | Type | Description | Required |
|---|---|---|---|
| Event | Options | Type of Facebook Lead Ads event to listen for | Yes |
| Page | Resource Locator | The Facebook page linked to the form for retrieving new leads | Yes |
| Form | Resource Locator | The form to monitor for fetching lead details upon submission | Yes |
| Simplify Output | Boolean | Whether to return a simplified version of the webhook event | No |

## Event Types

- **New Lead**: Triggered when a new lead is submitted through Facebook Lead Ads

## UseCases

- **Lead Processing** : Automatically process new leads from Facebook campaigns
- **CRM Integration** : Add new leads directly to your CRM system
- **Lead Notifications** : Send instant notifications when new leads are generated
- **Lead Qualification** : Automatically qualify and score incoming leads
- **Follow-up Automation** : Trigger immediate follow-up sequences for new leads

## Important Notes

- **API Limitation**: Due to Facebook API restrictions, you can use only one Facebook Lead Ads trigger per Facebook App. Multiple triggers require separate Facebook Apps.
- **Automatic App Installation**: When setting up the webhook, the app is automatically installed on the specified Facebook page.
- **Webhook Management**: The webhook subscription is automatically created during setup and removed when the trigger is deleted.
- **Verification Token**: The node uses its own ID as the verification token for webhook security.
- **HMAC Verification**: Webhook payloads are verified using HMAC-SHA256 for security.

## Webhook Details

The node creates two webhook endpoints:
- **Setup Endpoint** (GET): Used for Facebook webhook verification during setup
- **Default Endpoint** (POST): Receives actual lead data when new leads are generated

## Output Data Structure

When a new lead is generated, the node outputs data containing:
- Lead ID and creation time
- Form ID and associated page information
- Lead field data (name, email, phone, custom fields, etc.)
- Form submission metadata

With `Simplify Output` enabled, the output focuses on essential lead information. When disabled, the complete webhook payload is provided.

## Use Cases

- **Automated Lead Processing**: Instantly register new leads in CRM systems upon form submission
- **Immediate Response System**: Send automated welcome emails or SMS messages to new leads
- **Lead Scoring**: Automatically analyze and assign priority scores to incoming leads
- **Sales Team Notifications**: Alert sales teams immediately when high-quality leads are generated
- **Marketing Automation**: Trigger personalized marketing campaigns based on lead characteristics
- **Data Synchronization**: Real-time synchronization of lead data across multiple marketing tools
- **Lead Qualification**: Automatically route leads to appropriate teams based on form responses
- **Follow-up Scheduling**: Create calendar events or tasks for lead follow-up activities 