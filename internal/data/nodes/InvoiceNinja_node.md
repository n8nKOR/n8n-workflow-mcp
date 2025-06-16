# Invoice Ninja Node

## Overview

Consume Invoice Ninja API

## Credentials

- Name: invoiceNinjaApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: client

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection |  | No |  |
| Billing Address | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |
| Contacts | fixedCollection |  | No |  |
| Shipping Address | fixedCollection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Client ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Client ID | string |  | Yes |  |
| Options | collection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection |  | No |  |

### Resource: invoice

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |
| Invoice Items | fixedCollection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice ID | string |  | Yes |  |

#### Operation: email

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice ID | string |  | Yes |  |
| Options | collection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection |  | No |  |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |
| Time Logs | fixedCollection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string |  | Yes |  |
| Options | collection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection |  | No |  |

### Resource: payment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Invoice Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| Amount | number |  | No |  |
| Additional Fields | collection |  | No |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payment ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Payment ID | string |  | Yes |  |
| Options | collection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection |  | No |  |

### Resource: expense

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Expense ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Expense ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: quote

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |
| Invoice Items | fixedCollection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Quote ID | string |  | Yes |  |

#### Operation: email

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Quote ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Quote ID | string |  | Yes |  |
| Options | collection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection |  | No |  |

### Resource: bank_transaction

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bank Transaction ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bank Transaction ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: matchPayment

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Bank Transaction ID | string |  | Yes |  |
| Payment Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |

## UseCases

- **Invoice Management**: Automate invoice creation, sending, and tracking for business billing and accounts receivable processes
- **Payment Processing**: Streamline payment collection, recording, and reconciliation for improved cash flow management
- **Client Relationship Management**: Maintain comprehensive client databases with billing history and contact information
- **Project Time Tracking**: Track billable hours, project tasks, and time logs for accurate client billing and project management
- **Financial Reporting**: Generate financial reports, expense tracking, and business analytics for accounting and tax purposes
- **Quote and Estimate Management**: Create and manage business quotes, estimates, and proposal workflows for sales processes
- **Expense Management**: Track business expenses, categorize costs, and integrate with tax reporting and reimbursement processes
- **Banking Integration**: Sync bank transactions, automate payment matching, and streamline financial reconciliation
- **Recurring Billing**: Set up automated recurring invoices for subscription services, retainers, and ongoing client relationships
- **Multi-business Operations**: Manage multiple business entities, departments, or service lines with separate billing and tracking

