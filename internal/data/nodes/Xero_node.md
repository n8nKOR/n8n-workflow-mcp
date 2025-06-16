# Xero Node

## Overview

Consume Xero API

## Credentials

- Name: xeroOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: contact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string | Full name of contact/organisation | Yes |  |
| Additional Fields | collection | A user defined account number | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Contact ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether contacts with a status of ARCHIVED will be included in the response | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Contact ID | string |  | Yes |  |
| Update Fields | collection | A user defined account number | No |  |

### Resource: invoice

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Type | options | Accounts Payable or supplier invoice | Yes |  |
| Contact ID | string |  | Yes |  |
| Line Items | fixedCollection | Line item data | Yes |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Invoice ID | string |  | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Invoice ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether you | No |  |

## UseCases

- **Accounting Integration** : Integrate Xero accounting data with other business systems and workflows
- **Invoice Management** : Create, update, and manage invoices for automated billing processes
- **Contact Management** : Manage customer and supplier contact information synchronization
- **Financial Reporting** : Generate financial reports by extracting Xero accounting data
- **Automated Bookkeeping** : Automate bookkeeping tasks and data entry into Xero
- **Multi-Organization Management** : Manage accounting data across multiple Xero organizations
- **Payment Processing** : Track and manage payment status for invoices and bills
- **Customer Onboarding** : Automatically create customer contacts when new accounts are created
- **Expense Tracking** : Monitor and categorize business expenses through Xero integration
- **Tax Compliance** : Ensure tax compliance by maintaining accurate financial records
- **Cash Flow Management** : Monitor cash flow by tracking receivables and payables
- **Vendor Management** : Manage supplier relationships and payment schedules
- **Budget Planning** : Use Xero data for budget planning and financial forecasting
- **Sales Analytics** : Analyze sales data and customer payment patterns
- **Audit Trail Maintenance** : Maintain comprehensive audit trails for financial transactions
- **Bank Reconciliation** : Assist with bank reconciliation processes using transaction data
- **Subscription Billing** : Manage recurring invoices and subscription-based billing
- **Project Accounting** : Track project-related expenses and profitability
- **Client Portal Integration** : Integrate client portals with Xero accounting data
- **Financial Dashboard Creation** : Create real-time financial dashboards using Xero data

