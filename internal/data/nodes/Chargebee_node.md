# Chargebee Node

## Overview

Retrieve data from Chargebee API, a subscription billing and revenue management platform. Manage customers, invoices, and subscriptions with comprehensive billing operations.

## Credentials

- Name: chargebeeApi, Required: Yes
- Account Name: Your Chargebee account subdomain
- API Key: Chargebee API key for authentication
- Authentication Method: Basic auth using API key

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: customer

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Properties | collection | Properties to set on the new user | No |  |

### Resource: invoice

#### Operation: list

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Max Results | number | Max. amount of results to return(< 100). | No |  |
| Filters | fixedCollection | Filter for invoices | No |  |

#### Operation: pdfUrl

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice ID | string | The ID of the invoice to get | Yes |  |

### Resource: subscription

#### Operation: cancel

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subscription ID | string | The ID of the subscription to cancel | Yes |  |
| Schedule End of Term | boolean | Whether it will not cancel it directly in will instead schedule the cancelation for the end of the term | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subscription ID | string | The ID of the subscription to delete | Yes |  |

## UseCases

- **Subscription Management** : Manage customer subscriptions, billing cycles, and plan changes
- **Revenue Management** : Track and manage recurring revenue and subscription metrics
- **Customer Billing** : Automate customer billing processes and invoice generation
- **Payment Processing** : Handle subscription payments and payment method updates
- **Subscription Analytics** : Analyze subscription performance, churn, and growth metrics
- **Customer Lifecycle Management** : Manage customer onboarding, upgrades, and cancellations
- **Dunning Management** : Handle failed payments and subscription recovery processes
- **Invoice Automation** : Automate invoice creation, delivery, and payment collection
- **Pricing Strategy** : Implement and manage different pricing plans and billing models
- **Financial Reporting** : Generate financial reports for subscription-based businesses
- **Customer Portal Integration** : Integrate customer self-service portals with billing data
- **Subscription Upgrades** : Automate subscription upgrades and plan modifications
- **Trial Management** : Manage free trials and trial-to-paid conversions
- **Tax Management** : Handle tax calculations and compliance for subscription billing
- **Multi-Currency Support** : Manage international subscriptions with multiple currencies
- **Webhook Integration** : Integrate Chargebee webhooks for real-time event processing
- **CRM Integration** : Synchronize subscription data with CRM systems
- **Usage-Based Billing** : Implement metered billing and usage-based pricing models
- **Subscription Forecasting** : Forecast subscription revenue and business growth
- **Compliance Management** : Ensure billing compliance with regulatory requirements

