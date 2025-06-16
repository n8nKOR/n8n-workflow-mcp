# Stripe Node

## Overview

Consume the Stripe API

## Credentials

- Name: stripeApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: customerCard

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer to be associated with this card | Yes |  |
| Card Token | string | Token representing sensitive card information | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer whose card to remove | Yes |  |
| Card ID | string | ID of the card to remove | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer whose card to retrieve | Yes |  |
| Source ID | string | ID of the source to retrieve | Yes |  |

### Resource: charge

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer to be associated with this charge | Yes |  |
| Amount | number | Amount in cents to be collected for this charge, e.g. enter <code>100</code> for $1.00 | Yes |  |
| Currency Name or ID | options | Three-letter ISO currency code, e.g. <code>USD</code> or <code>EUR</code>. It must be a <a href= | Yes |  |
| Source ID | string | ID of the customer | Yes |  |
| Additional Fields | collection | Arbitrary text to describe the charge to create | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Charge ID | string | ID of the charge to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Charge ID | string | ID of the charge to update | Yes |  |
| Update Fields | collection | Arbitrary text to describe the charge to update | No |  |

### Resource: coupon

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Apply | options | How long the discount will be in effect | Yes |  |
| Discount Type | options | Whether the coupon discount is a percentage or a fixed amount | Yes |  |
| Amount Off | number | Amount in cents to subtract from an invoice total, e.g. enter <code>100</code> for $1.00 | Yes |  |
| Currency Name or ID | options | Three-letter ISO currency code, e.g. <code>USD</code> or <code>EUR</code>. It must be a <a href= | Yes |  |
| Percent Off | number | Percentage to apply with the coupon | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: customer

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Full name or business name of the customer to create | Yes |  |
| Additional Fields | collection | Address of the customer to create | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Customer | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer to update | Yes |  |
| Update Fields | collection | Address of the customer to update | No |  |

### Resource: source

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer to attach the source to | Yes |  |
| Type | options | Type of source (payment instrument) to create | Yes |  |
| Amount | number | Amount in cents to be collected for this charge, e.g. enter <code>100</code> for $1.00 | No |  |
| Currency Name or ID | options | Three-letter ISO currency code, e.g. <code>USD</code> or <code>EUR</code>. It must be a <a href= | No |  |
| Additional Fields | collection | Set of key-value pairs to attach to the source to create | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer whose source to delete | Yes |  |
| Source ID | string | ID of the source to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Source ID | string | ID of the source to retrieve | Yes |  |

### Resource: token

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Type | options | Type of token to create | Yes |  |
| Card Number | string |  | No |  |
| CVC | string | Security code printed on the back of the card | No |  |
| Expiration Month | string | Number of the month when the card will expire | No |  |
| Expiration Year | string | Year when the card will expire | No |  |

## UseCases

- **E-commerce Payment Processing**: Process online payments, handle customer billing, and manage subscription payments for e-commerce platforms
- **Subscription and Recurring Billing**: Automate recurring subscription charges, manage billing cycles, and handle subscription lifecycle events
- **Customer Payment Method Management**: Store and manage customer payment methods, cards, and billing information securely
- **Order and Transaction Processing**: Process single payments, handle refunds, and manage transaction history for order fulfillment systems
- **Marketplace Payment Splitting**: Handle marketplace transactions with automatic payment splitting between sellers and platform fees
- **Subscription Plan Management**: Create and manage different subscription tiers, pricing plans, and promotional coupons
- **Failed Payment Recovery**: Automatically retry failed payments, update expired cards, and manage dunning processes
- **Financial Reporting and Analytics**: Generate revenue reports, track payment metrics, and analyze customer payment behavior
- **Tax and Compliance Management**: Handle tax calculations, compliance requirements, and international payment processing
- **Customer Support Integration**: Manage payment disputes, chargebacks, and customer billing inquiries through automated workflows

