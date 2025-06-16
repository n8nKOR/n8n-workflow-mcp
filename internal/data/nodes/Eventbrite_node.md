# Eventbrite Trigger

## Overview

The Eventbrite Trigger node handles Eventbrite events via webhooks in your n8n workflows. This trigger node listens for various Eventbrite events such as attendee check-ins, order placements, event updates, and more, allowing you to automate workflows based on event management activities.

## Credentials

- **Credential Name**: `eventbriteApi` or `eventbriteOAuth2Api`
- **Required**: Yes
- **Configuration**: Eventbrite API credentials with appropriate permissions

## Inputs

This node does not accept input connections as it is a trigger node.

## Outputs

- **Main**: Outputs event data when an Eventbrite webhook is triggered

## Properties

### Resource: Webhook Events

#### Operation: Eventbrite Event Handling

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | Authentication method (Private Key or OAuth2) | Yes |
| Organization Name or ID | Options | The Eventbrite Organization to work on | Yes |
| Event Name or ID | Options | Limit the triggers to this event | Yes |
| Actions | Multi Options | One or more action to subscribe to | Yes |
| Resolve Data | Boolean | Whether to resolve the data automatically | No |

## Event Actions

- **attendee.checked_in**: Attendee checked into event
- **attendee.checked_out**: Attendee checked out of event
- **attendee.updated**: Attendee information updated
- **event.created**: New event created
- **event.published**: Event published
- **event.unpublished**: Event unpublished
- **event.updated**: Event information updated
- **order.placed**: New order placed
- **order.refunded**: Order refunded
- **order.updated**: Order information updated
- **organizer.updated**: Organizer information updated
- **ticket_class.created**: New ticket class created
- **ticket_class.deleted**: Ticket class deleted
- **ticket_class.updated**: Ticket class updated
- **venue.updated**: Venue information updated

## UseCases

- **Event Management** : Automate event management workflows based on Eventbrite activities
- **Attendee Tracking** : Track attendee check-ins and check-outs automatically
- **Order Processing** : Process new orders and handle refunds automatically
- **Event Analytics** : Collect and analyze event data for reporting
- **Customer Communication** : Send automated communications based on event activities 