# Cal.com Trigger

## Overview

The Cal.com Trigger node enables you to handle Cal.com events via webhooks in your n8n workflows. This trigger node listens for various Cal.com events such as booking creation, cancellation, rescheduling, and meeting completion, allowing you to automate workflows based on calendar activities.

## Credentials

- **Credential Name**: `calApi`
- **Required**: Yes
- **Configuration**: Cal.com API credentials with appropriate permissions

## Inputs

This node does not accept input connections as it is a trigger node.

## Outputs

- **Main**: Outputs event data when a Cal.com webhook is triggered

## Properties

### Resource: Webhook Events

#### Operation: Cal.com Event Handling

| Field Name | Type | Description | Required |
|---|---|---|---|
| Events | Multi Options | Types of Cal.com events to listen for | Yes |
| API Version | Options | Cal.com API version to use (Before v2.0 or v2.0 Onwards) | Yes |
| App ID | String | The ID of the App to monitor | No |
| EventType Name or ID | Options | The EventType to monitor | No |
| Payload Template | String | Template to customize the webhook payload | No |

## Event Types

- **Booking Cancelled**: Receive notifications when a Cal event is canceled
- **Booking Created**: Receive notifications when a new Cal event is created
- **Booking Rescheduled**: Receive notifications when a Cal event is rescheduled
- **Meeting Ended**: Receive notifications when a Cal event or meeting has ended

## UseCases

- **Booking Notifications** : Send notifications when new bookings are created
- **Calendar Automation** : Automate workflows based on calendar events
- **Meeting Follow-up** : Trigger follow-up actions after meetings end
- **Booking Management** : Handle booking cancellations and rescheduling
- **Event Analytics** : Track and analyze booking patterns and metrics

**Important Note**: This integration only provides trigger functionality. There is no action node for creating or managing bookings through n8n - only for receiving webhook notifications about Cal.com events. 