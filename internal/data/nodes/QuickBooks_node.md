# QuickBooks Online Node

## Overview

Consume the QuickBooks Online API

## Credentials

- Name: quickBooksOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: bill

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Vendor | options | Vendor for the bill | Yes |  |
| Line Items | fixedCollection | Line items for the bill | Yes |  |
| Additional Fields | collection | Additional fields for the bill | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bill ID | string | Unique identifier for the bill | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bill ID | string | Unique identifier for the bill | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 50 |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bill ID | string | Unique identifier for the bill | Yes |  |
| Update Fields | collection | Fields to update | No |  |

### Resource: customer

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Customer name | Yes |  |
| Additional Fields | collection | Additional customer fields | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | Unique identifier for the customer | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 50 |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | string | Unique identifier for the customer | Yes |  |
| Update Fields | collection | Fields to update | No |  |

### Resource: invoice

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer | options | Customer for the invoice | Yes |  |
| Line Items | fixedCollection | Line items for the invoice | Yes |  |
| Additional Fields | collection | Additional fields for the invoice | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice ID | string | Unique identifier for the invoice | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice ID | string | Unique identifier for the invoice | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 50 |

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice ID | string | Unique identifier for the invoice | Yes |  |
| Email | string | Email address to send the invoice to | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice ID | string | Unique identifier for the invoice | Yes |  |
| Update Fields | collection | Fields to update | No |  |

### Resource: item

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Item name | Yes |  |
| Type | options | Type of item | Yes |  |
| Additional Fields | collection | Additional item fields | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Item ID | string | Unique identifier for the item | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 50 |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Item ID | string | Unique identifier for the item | Yes |  |
| Update Fields | collection | Fields to update | No |  |

### Resource: payment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer | options | Customer for the payment | Yes |  |
| Total Amount | number | Total amount of the payment | Yes |  |
| Additional Fields | collection | Additional payment fields | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payment ID | string | Unique identifier for the payment | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payment ID | string | Unique identifier for the payment | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 50 |

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payment ID | string | Unique identifier for the payment | Yes |  |
| Email | string | Email address to send the payment receipt to | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payment ID | string | Unique identifier for the payment | Yes |  |
| Update Fields | collection | Fields to update | No |  |

### Resource: vendor

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Vendor name | Yes |  |
| Additional Fields | collection | Additional vendor fields | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Vendor ID | string | Unique identifier for the vendor | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 50 |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Vendor ID | string | Unique identifier for the vendor | Yes |  |
| Update Fields | collection | Fields to update | No |  |

## UseCases

- **Accounting Automation** : Automate accounting processes and financial data management
- **Invoice Management** : Create, send, and track invoices and billing information
- **Expense Tracking** : Track business expenses and categorize financial transactions
- **Financial Reporting** : Generate financial reports and statements for business analysis
- **Customer Management** : Manage customer information and payment histories
- **Vendor Management** : Manage vendor relationships and payment processing
- **Tax Preparation** : Prepare tax documents and compliance reporting
- **Cash Flow Management** : Monitor cash flow and financial performance
- **Integration with E-commerce** : Integrate e-commerce sales data with accounting systems
- **Multi-Company Accounting** : Manage accounting for multiple business entities

