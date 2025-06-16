# Paddle Node

## Overview

Consume Paddle API for comprehensive payment processing and subscription management. Paddle is an all-in-one platform for software businesses that handles payment processing, subscription billing, tax compliance, and customer management. This node enables integration with Paddle's powerful commerce platform for automated billing workflows and revenue operations.

## Credentials

- Name: paddleApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: coupon

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Coupon Type | options | Either product (valid for specified products or subscription plans) or checkout (valid for any checkout) | No |  |
| Product Names or IDs | multiOptions | Comma-separated list of product IDs. Required if coupon_type is product. Choose from the list, or specify IDs using an <a href= | Yes |  |
| Discount Type | options | Either flat or percentage | No |  |
| Discount Amount Currency | number | Discount amount in currency | No |  |
| Discount Amount % | number | Discount amount in percentage | No |  |
| Currency | options | The currency must match the balance currency specified in your account | No |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | Attributes in JSON form | No |  |
| Additional Fields | collection | Number of times a coupon can be used in a checkout. This will be set to 999,999 by default, if not specified. | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string | The specific product/subscription ID | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update By | options | Either flat or percentage | No |  |
| Coupon Code | string | Identify the coupon to update | No |  |
| Group | string | The name of the group of coupons you want to update | No |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | Attributes in JSON form | No |  |
| Additional Fields | collection | Number of times a coupon can be used in a checkout. This will be set to 999,999 by default, if not specified. | No |  |

### Resource: payment

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | Attributes in JSON form | No |  |
| Additional Fields | collection | Payment starting from date | No |  |

#### Operation: reschedule

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payment Name or ID | options | The upcoming subscription payment ID. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Date | dateTime | Date you want to move the payment to | No |  |

### Resource: plan

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Plan ID | string | Filter: The subscription plan ID | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: product

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: order

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Checkout ID | string | The identifier of the buyer’s checkout | Yes |  |

### Resource: user

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | json | Attributes in JSON form | No |  |
| Additional Fields | collection | Filter: The subscription plan ID | No |  |

## UseCases

- **Subscription Revenue Management**: Automate subscription billing cycles, payment collection, and revenue recognition for SaaS and software businesses
- **Promotional Campaign Automation**: Create and manage discount coupons, promotional codes, and special offers with automated distribution and tracking
- **Customer Lifecycle Management**: Track customer payments, subscription changes, and billing history for comprehensive customer relationship management
- **Failed Payment Recovery**: Implement automated workflows to handle failed payments, retry logic, and customer communication for payment issues
- **Revenue Analytics and Reporting**: Generate detailed reports on payment performance, subscription metrics, and revenue trends for business intelligence
- **Tax Compliance Automation**: Leverage Paddle's built-in tax handling to automatically calculate and collect appropriate taxes based on customer location
- **Payment Reconciliation**: Automate the matching of payments with orders and subscriptions for accurate financial record-keeping
- **Churn Prevention Workflows**: Monitor payment patterns and subscription status to identify at-risk customers and trigger retention campaigns
- **International Payment Processing**: Handle global payments with automatic currency conversion and region-specific payment method support
- **Subscription Plan Management**: Dynamically manage subscription tiers, pricing changes, and plan upgrades/downgrades based on customer behavior

