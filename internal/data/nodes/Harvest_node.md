# Harvest Node

## Overview

Access data on Harvest

## Credentials

- Name: harvestApi, Required: Yes (for Access Token authentication)
- Name: harvestOAuth2Api, Required: Yes (for OAuth2 authentication)

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Client

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | The name of the client | Yes |
| Address | String | The client's address | No |
| Currency | String | The currency for the client | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Client ID | String | The ID of the client to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Client ID | String | The ID of the client to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Client ID | String | The ID of the client to update | Yes |
| Name | String | The updated name of the client | No |
| Address | String | The updated address of the client | No |

### Resource: Company

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Account ID | String | The account ID to retrieve company information | Yes |

### Resource: Contact

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Client ID | String | The client ID for the contact | Yes |
| First Name | String | The first name of the contact | Yes |
| Last Name | String | The last name of the contact | Yes |
| Email | String | The email address of the contact | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | The ID of the contact to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | The ID of the contact to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Contact ID | String | The ID of the contact to update | Yes |
| First Name | String | The updated first name | No |
| Last Name | String | The updated last name | No |
| Email | String | The updated email address | No |

### Resource: Estimate

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Client ID | String | The client ID for the estimate | Yes |
| Subject | String | The subject of the estimate | No |
| Notes | String | Additional notes for the estimate | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Estimate ID | String | The ID of the estimate to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Estimate ID | String | The ID of the estimate to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Estimate ID | String | The ID of the estimate to update | Yes |
| Subject | String | The updated subject | No |
| Notes | String | The updated notes | No |

### Resource: Expense

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Project ID | String | The project ID for the expense | Yes |
| Expense Category ID | String | The expense category ID | Yes |
| Spent Date | String | The date the expense was incurred | Yes |
| Total Cost | Number | The total cost of the expense | Yes |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Expense ID | String | The ID of the expense to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Expense ID | String | The ID of the expense to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Expense ID | String | The ID of the expense to update | Yes |
| Total Cost | Number | The updated total cost | No |
| Notes | String | The updated notes | No |

### Resource: Invoice

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Client ID | String | The client ID for the invoice | Yes |
| Subject | String | The subject of the invoice | No |
| Notes | String | Additional notes for the invoice | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Invoice ID | String | The ID of the invoice to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Invoice ID | String | The ID of the invoice to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Invoice ID | String | The ID of the invoice to update | Yes |
| Subject | String | The updated subject | No |
| Notes | String | The updated notes | No |

### Resource: Project

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Client ID | String | The client ID for the project | Yes |
| Name | String | The name of the project | Yes |
| Code | String | The project code | No |
| Is Active | Boolean | Whether the project is active | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Project ID | String | The ID of the project to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Project ID | String | The ID of the project to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Project ID | String | The ID of the project to update | Yes |
| Name | String | The updated name | No |
| Code | String | The updated project code | No |
| Is Active | Boolean | Whether the project is active | No |

### Resource: Task

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | The name of the task | Yes |
| Billable By Default | Boolean | Whether the task is billable by default | No |
| Default Hourly Rate | Number | The default hourly rate for the task | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | The ID of the task to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | The ID of the task to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | The ID of the task to update | Yes |
| Name | String | The updated name | No |
| Billable By Default | Boolean | Whether the task is billable by default | No |
| Default Hourly Rate | Number | The updated default hourly rate | No |

### Resource: Time Entry

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Project ID | String | The project ID for the time entry | Yes |
| Task ID | String | The task ID for the time entry | Yes |
| Spent Date | String | The date the time was spent | Yes |
| Hours | Number | The number of hours spent | No |
| Started Time | String | The start time of the entry | No |
| Ended Time | String | The end time of the entry | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Time Entry ID | String | The ID of the time entry to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Time Entry ID | String | The ID of the time entry to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Time Entry ID | String | The ID of the time entry to update | Yes |
| Hours | Number | The updated number of hours | No |
| Notes | String | The updated notes | No |

### Resource: User

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| User ID | String | The ID of the user to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| User ID | String | The ID of the user to update | Yes |
| First Name | String | The updated first name | No |
| Last Name | String | The updated last name | No |
| Email | String | The updated email address | No |

## UseCases

- **Time Tracking** : Track time spent on projects and tasks for accurate billing
- **Project Management** : Create and manage projects with budgets and timelines
- **Client Management** : Manage client information and contacts for business relationships
- **Invoicing** : Generate invoices for completed work and track payments
- **Expense Tracking** : Track project-related expenses for accurate cost accounting
- **Team Management** : Manage team members and their time entries for project oversight
- **Reporting** : Generate reports on time, expenses, and project progress for business insights
- **Resource Planning** : Plan and allocate team resources across multiple projects
- **Billing Automation** : Automate billing processes based on tracked time and expenses
- **Performance Analytics** : Analyze team performance and project profitability

