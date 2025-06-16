# Wise Node

## Overview

Consume the Wise API

## Credentials

- Name: wiseApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: account

#### Operation: getBalances

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Profile Name or ID | options | ID of the user profile to retrieve the balance of. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getStatement

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Profile Name or ID | options | ID of the user profile whose account to retrieve the statement of. Choose from the list, or specify an ID using an <a href= | No |  |
| Borderless Account Name or ID | options | ID of the borderless account to retrieve the statement of. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Currency | string | Code of the currency of the borderless account to retrieve the statement of | No |  |
| Format | options | File format to retrieve the statement in | No |  |
| Put Output File in Field | string |  | Yes |  |
| File Name | string | Name of the file that will be downloaded | Yes |  |
| Additional Fields | collection | Line style to retrieve the statement in | No |  |

### Resource: exchangeRate

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Source Currency | string | Code of the source currency to retrieve the exchange rate for | No |  |
| Target Currency | string | Code of the target currency to retrieve the exchange rate for | No |  |
| Additional Fields | collection | Range of time to retrieve the exchange rate for | No |  |

### Resource: profile

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Profile Name or ID | options | ID of the user profile to retrieve. Choose from the list, or specify an ID using an <a href= | Yes |  |

### Resource: quote

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Profile Name or ID | options | ID of the user profile to create the quote under. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Target Account Name or ID | options | ID of the account that will receive the funds. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Amount Type | options | Whether the amount is to be sent or received | No |  |
| Amount | number | Amount of funds for the quote to create | No |  |
| Source Currency | string | Code of the currency to send for the quote to create | No |  |
| Target Currency | string | Code of the currency to receive for the quote to create | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Quote ID | string | ID of the quote to retrieve | Yes |  |

### Resource: recipient

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: transfer

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Profile Name or ID | options | ID of the user profile to retrieve the balance of. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Quote ID | string | ID of the quote based on which to create the transfer | Yes |  |
| Target Account Name or ID | options | ID of the account that will receive the funds. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Additional Fields | collection | Reference text to show in the recipient | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Transfer ID | string | ID of the transfer to delete | Yes |  |

#### Operation: execute

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Profile Name or ID | options | ID of the user profile to execute the transfer under. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Transfer ID | string | ID of the transfer to execute | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Transfer ID | string | ID of the transfer to retrieve | Yes |  |
| Download Receipt | boolean | Whether to download the transfer receipt as a PDF file. Only for executed transfers, having status  | Yes |  |
| Put Output File in Field | string |  | Yes |  |
| File Name | string | Name of the file that will be downloaded | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Profile Name or ID | options | ID of the user profile to retrieve. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Range of time for filtering the transfers | No |  |

## UseCases

- **International Payment Automation** : Automate international money transfers and currency exchanges
- **Multi-Currency Account Management** : Manage multiple currency accounts and balances
- **Cross-Border E-commerce** : Process international payments for e-commerce transactions
- **Freelancer Payment Processing** : Pay international freelancers and contractors efficiently
- **Business Expense Management** : Manage international business expenses and reimbursements
- **Currency Exchange Monitoring** : Monitor exchange rates and execute transfers at optimal rates
- **Supplier Payment Automation** : Automate payments to international suppliers and vendors
- **Investment Portfolio Management** : Manage international investment transfers and currency exposure
- **Travel Expense Processing** : Process travel expenses and international transaction reconciliation
- **Financial Reporting Integration** : Integrate Wise transaction data with accounting and reporting systems

