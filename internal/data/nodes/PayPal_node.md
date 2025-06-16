# PayPal Node

## Overview

Consume PayPal API

## Credentials

- Name: payPalApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: payout

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Sender Batch ID | string | A sender-specified ID number. Tracks the payout in an accounting system. | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Items | fixedCollection | The unencrypted phone number | Yes |  |
| Items | json | An array of individual payout items | No |  |
| Additional Fields | collection | The subject line for the email that PayPal sends when payment for a payout item completes. The subject line is the same for all recipients. Max length: 255 characters. | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payout Batch ID | string | The ID of the payout for which to show details | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: payoutItem

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payout Item ID | string | The ID of the payout item for which to show details | Yes |  |

#### Operation: cancel

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payout Item ID | string | The ID of the payout item to cancel | Yes |  |

## UseCases

- [Payment Processing] : Process online payments and transactions for e-commerce businesses
- [Subscription Management] : Manage recurring subscriptions and billing cycles
- [Payout Management] : Send payments to vendors, affiliates, and service providers
- [E-commerce Integration] : Integrate PayPal payments into online stores and marketplaces
- [Invoice Management] : Create and send invoices with PayPal payment options
- [Refund Processing] : Process refunds and handle payment disputes automatically
- [Financial Reporting] : Generate financial reports and transaction analytics
- [Marketplace Payments] : Handle payments for marketplace platforms and multi-vendor systems
- [Donation Processing] : Process donations for non-profit organizations and fundraising
- [Mobile Payments] : Enable mobile payment processing for apps and mobile commerce
- [International Payments] : Process international payments with currency conversion
- [Payment Security] : Implement secure payment processing with fraud protection
- [API Integration] : Integrate PayPal services with custom applications and workflows
- [Transaction Monitoring] : Monitor payment transactions and identify suspicious activity
- [Customer Support] : Handle payment-related customer support and dispute resolution

