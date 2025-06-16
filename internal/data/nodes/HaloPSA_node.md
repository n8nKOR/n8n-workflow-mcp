# HaloPSA Node

## Overview

Consume HaloPSA API

## Credentials

- Name: haloPSAApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: client

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Enter client name | Yes |  |
| Additional Fields | collection | Whether the client is VIP or not | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether to include active customers in the response | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Client ID | string |  | Yes |  |
| Update Fields | collection | Whether the client is VIP or not | No |  |

### Resource: ticket

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket Type Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Summary | string |  | Yes |  |
| Details | string |  | Yes |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether to include active customers in the response | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string |  | No |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: site

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Enter site name | Yes |  |
| Select Client by ID | boolean | Whether client can be selected by ID | No |  |
| Client ID | string |  | Yes |  |
| Client Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether to include active sites in the response | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Site ID | string |  | No |  |
| Update Fields | collection | Enter site name | No |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Enter user name | Yes |  |
| Site Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Additional Fields | collection | Your new password must be at least 8 characters long and contain at least one letter, one number or symbol, one upper case character and one lower case character | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Whether to include active customers in the response | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | No |  |
| Update Fields | collection | Enter user name | No |  |

## UseCases

- **Professional Services Automation**: Streamline client management, project tracking, and service delivery for professional services businesses
- **IT Service Management**: Manage IT support tickets, client requests, and technical service delivery workflows
- **Client Relationship Management**: Automate client onboarding, account management, and customer communication processes
- **Multi-Site Operations**: Coordinate service delivery across multiple client sites and locations
- **User Account Management**: Automate user provisioning, access management, and account lifecycle processes
- **Ticket Workflow Automation**: Streamline support ticket creation, assignment, and resolution tracking
- **Service Desk Operations**: Manage help desk operations, incident tracking, and problem resolution workflows
- **Project Management Integration**: Connect PSA data with project management tools for unified service delivery
- **Billing and Time Tracking**: Integrate service delivery with billing systems and time tracking applications
- **Performance Analytics**: Track service delivery metrics, client satisfaction, and operational efficiency

