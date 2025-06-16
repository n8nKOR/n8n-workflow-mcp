# WooCommerce Node

## Overview

Consume WooCommerce API

## Credentials

- Name: wooCommerceApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: customer

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string |  | Yes |  |

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
| Filters | collection | Email address to filter customers by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | ID of the customer to update | Yes |  |

### Resource: product

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Product name | Yes |  |
| Additional Fields | collection | If managing stock, this controls if backorders are allowed | No |  |
| Dimensions | fixedCollection | Product dimensions | No |  |
| Images | fixedCollection | Product Image | No |  |
| Metadata | fixedCollection | Meta data | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string |  | No |  |
| Update Fields | collection | If managing stock, this controls if backorders are allowed | No |  |
| Dimensions | fixedCollection | Product dimensions | No |  |
| Images | fixedCollection | Product Image | No |  |
| Metadata | fixedCollection | Meta data | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Limit response to resources published after a given ISO8601 compliant date | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string |  | No |  |

### Resource: order

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Currency the order was created with | No |  |
| Billing | fixedCollection | Billing address | No |  |
| Coupon Lines | fixedCollection | Coupons line data | No |  |
| Fee Lines | fixedCollection | Fee line data | No |  |
| Line Items | fixedCollection | Line item data | No |  |
| Metadata | fixedCollection | Meta data | No |  |
| Shipping | fixedCollection | Shipping address | No |  |
| Shipping Lines | fixedCollection | Shipping line data | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Order ID | string |  | No |  |
| Update Fields | collection | Currency the order was created with | No |  |
| Billing | fixedCollection | Billing address | No |  |
| Coupon Lines | fixedCollection | Coupons line data | No |  |
| Fee Lines | fixedCollection | Fee line data | No |  |
| Line Items | fixedCollection | Line item data | No |  |
| Metadata | fixedCollection | Meta data | No |  |
| Shipping | fixedCollection | Shipping address | No |  |
| Shipping Lines | fixedCollection | Shipping line data | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Order ID | string |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Limit response to resources published after a given ISO8601 compliant date | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Order ID | string |  | No |  |

## UseCases

- **E-commerce Store Management** : Manage WooCommerce store operations including products, orders, and customers
- **Customer Data Synchronization** : Sync customer information between WooCommerce and CRM systems
- **Order Processing Automation** : Automate order fulfillment, shipping notifications, and status updates
- **Inventory Management** : Automatically update product stock levels and inventory across platforms
- **Product Catalog Sync** : Synchronize product information between WooCommerce and external systems
- **Sales Analytics** : Extract order and customer data for business intelligence and reporting
- **Customer Service Integration** : Create support tickets automatically for order issues or returns
- **Marketing Automation** : Trigger marketing campaigns based on customer purchase behavior
- **Multi-Channel Commerce** : Integrate WooCommerce with other sales channels for unified order management
- **Financial Reporting** : Generate automated financial reports and reconcile sales data

