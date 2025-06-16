# Gumroad

## Overview

The Gumroad Trigger node provides webhook-based integration with Gumroad, a popular e-commerce platform for digital products. This trigger node automatically responds to various Gumroad events such as sales, refunds, cancellations, and disputes, enabling real-time workflow automation for digital product businesses.

## Credentials

- **Credential Name**: `gumroadApi`
- **Required**: Yes
- **Configuration**: 
  - API Key: Your Gumroad API key obtained from the Gumroad dashboard

## Inputs

This is a trigger node and does not accept input connections.

## Outputs

- **Main**: Returns event data when triggered by Gumroad events (sales, refunds, cancellations, disputes)

## Properties

### Resource: Webhook Events

#### Operation: Event Handling

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | Select the event type to trigger on (Sale, Refund, Cancellation, Dispute, Dispute Won) | Yes |

## UseCases

- **Sales Notifications** : Send automatic notifications when products are sold
- **Revenue Tracking** : Track sales data and update financial systems
- **Customer Management** : Add customers to CRM systems after purchases
- **Refund Processing** : Handle refund notifications and update accounting systems
- **Dispute Management** : Monitor and respond to payment disputes automatically 