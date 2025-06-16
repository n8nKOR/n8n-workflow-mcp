# Calendly Trigger

## Overview

The Calendly Trigger node starts workflows when Calendly events occur. This trigger node listens for Calendly webhook events such as event creation and cancellation, allowing you to automate workflows based on scheduling activities.

## Credentials

- **Credential Name**: `calendlyApi` or `calendlyOAuth2Api`
- **Required**: Yes
- **Configuration**: Calendly API credentials with appropriate permissions

## Inputs

This node does not accept input connections as it is a trigger node.

## Outputs

- **Main**: Outputs event data when a Calendly webhook is triggered

## Properties

### Resource: Webhook Events

#### Operation: Calendly Event Handling

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | Authentication method (OAuth2 or API Key) | Yes |
| Scope | Options | Scope of webhook (Organization or User) | Yes |
| Events | Multi Options | Types of Calendly events to listen for | Yes |

## Event Types

- **Event Created**: Receive notifications when a new Calendly event is created
- **Event Canceled**: Receive notifications when a Calendly event is canceled

## Authentication Methods

- **OAuth2 (recommended)**: Modern OAuth2 authentication
- **API Key**: Personal Access Token authentication (being deprecated)

## UseCases

- **Event Notifications** : Send notifications when Calendly events are created or canceled
- **Calendar Integration** : Integrate Calendly events with other calendar systems
- **Workflow Automation** : Trigger automated workflows based on scheduling activities
- **Customer Communication** : Send automated messages when events are scheduled or canceled
- **Analytics Tracking** : Track scheduling patterns and event metrics 