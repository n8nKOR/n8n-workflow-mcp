# AcuityScheduling

## Overview

The Acuity Scheduling Trigger node provides webhook-based integration with Acuity Scheduling, a popular online appointment scheduling service. This trigger node automatically responds to various appointment-related events, enabling real-time workflow automation based on scheduling activities.

**Important Note**: This integration only provides trigger functionality. There is no action node for creating or modifying appointments through n8n.

## Credentials

- **Credential Name**: `acuitySchedulingApi` or `acuitySchedulingOAuth2Api`
- **Required**: Yes
- **Authentication Options**:
  - **API Key**: Direct API key authentication
  - **OAuth2**: OAuth2 authentication flow

## Inputs

This is a trigger node and does not accept input connections.

## Outputs

- **Main**: Returns appointment or order data when triggered by Acuity Scheduling events

## Properties

### Resource: Webhook Events

#### Operation: Event Handling

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | Choose between 'API Key' or 'OAuth2' authentication methods | Yes |
| Event | Options | Select the event type to trigger on (appointment.scheduled, appointment.canceled, etc.) | Yes |
| Resolve Data | Boolean | Whether to resolve full data automatically (if false, only returns ID) | No |

## UseCases

- **Appointment Notifications** : Send automatic notifications when appointments are scheduled, canceled, or rescheduled
- **Calendar Integration** : Sync Acuity appointments with external calendar systems
- **Customer Communication** : Trigger follow-up emails or SMS messages based on appointment events
- **Business Analytics** : Track appointment metrics and generate reports on scheduling activity
- **Payment Processing** : Process payments automatically when orders are completed
- **CRM Integration**: Automatically update CRM systems when appointment details change
- **Customer Follow-up**: Trigger automated follow-up sequences based on appointment status
- **Analytics and Reporting**: Collect appointment data for business intelligence and reporting
- **Multi-platform Synchronization**: Sync appointment changes across multiple calendar systems 