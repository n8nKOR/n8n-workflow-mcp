# Freshservice Node

## Overview

Consume the Freshservice API

## Credentials

- Name: freshserviceApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: agent

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string |  | Yes |  |
| First Name | string |  | Yes |  |
| Roles | fixedCollection | Role to assign to the agent | Yes |  |
| Additional Fields | collection | IDs of the departments to which the agent belongs. Choose from the list or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent ID | string | ID of the agent to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent ID | string | ID of the agent to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | ID of the department to which the agent belongs. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent ID | string | ID of the agent to update | Yes |  |
| Update Fields | collection | IDs of the departments to which the agent belongs. Choose from the list or specify an ID using an <a href= | No |  |

### Resource: agentGroup

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Additional Fields | collection | ID of the user to whom an escalation email is sent if a ticket in this group is unassigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent Group ID | string | ID of the agent group to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent Group ID | string | ID of the agent group to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent Group ID | string | ID of the agent group to update | Yes |  |
| Update Fields | collection | ID of the agent to whom an escalation email is sent if a ticket in this group is unassigned. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: agentRole

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent Role ID | string | ID of the agent role to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: announcement

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string |  | Yes |  |
| Body | string | HTML supported | Yes |  |
| Visibility | options |  | Yes |  |
| Visible From | dateTime | Timestamp at which announcement becomes active | Yes |  |
| Additional Fields | collection | Comma-separated additional email addresses to which the announcement needs to be sent | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Announcement ID | string | ID of the announcement to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Announcement ID | string | ID of the announcement to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Announcement ID | string | ID of the announcement to update | Yes |  |
| Update Fields | collection | Comma-separated additional email addresses to which the announcement needs to be sent | No |  |

### Resource: asset

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asset Name | string |  | Yes |  |
| Asset Type Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Asset Fields | fixedCollection | The ID of the field to add custom field to. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asset Display ID | string | Display ID of the asset to delete. Do not confuse with asset ID. | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asset Display ID | string | Display ID of the asset to retrieve. Do not confuse with asset ID. | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | ID of the agent by whom the asset is managed. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asset Display ID | string | Display ID of the asset to update. Do not confuse with asset ID. | Yes |  |
| Asset Fields | fixedCollection | The ID of the field to add custom field to. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: assetType

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asset Type ID | string | ID of the asset type to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asset Type ID | string | ID of the asset type to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asset Type ID | string | ID of the asset type to update | Yes |  |
| Update Fields | collection |  | No |  |

### Resource: change

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Requester Name or ID | options | ID of the requester of the change. Choose from the list or specify an ID. You can also specify the ID using an <a href= | Yes |  |
| Subject | string |  | Yes |  |
| Planned Start Date | dateTime |  | Yes |  |
| Planned End Date | dateTime |  | Yes |  |
| Additional Fields | collection | ID of the agent to whom the change is assigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Change ID | string | ID of the change to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Change ID | string | ID of the change to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Change ID | string | ID of the change to update | Yes |  |
| Update Fields | collection | ID of the agent to whom the change is assigned. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: department

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Additional Fields | collection | Comma-separated email domains associated with the department | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Department ID | string | ID of the department to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Department ID | string | ID of the department to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Name of the department | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Department ID | string | ID of the department to update | Yes |  |
| Update Fields | collection | Comma-separated email domains associated with the department | No |  |

### Resource: location

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the location | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Location ID | string | ID of the location to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Location ID | string | ID of the location to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Location ID | string | ID of the location to update | Yes |  |
| Update Fields | collection |  | No |  |

### Resource: problem

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subject | string |  | Yes |  |
| Requester Name or ID | options | ID of the initiator of the problem. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Due By | dateTime | Date when the problem is due to be solved | No |  |
| Additional Fields | collection | ID of the agent to whom the problem is assigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Problem ID | string | ID of the problem to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Problem ID | string | ID of the problem to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Problem ID | string | ID of the problem to update | Yes |  |
| Update Fields | collection | ID of the agent to whom the problem is assigned. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: product

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asset Type Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string |  | Yes |  |
| Additional Fields | collection | HTML supported | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string | ID of the product to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string | ID of the product to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string | ID of the product to update | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: release

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subject | string |  | Yes |  |
| Release Type | options |  | No |  |
| Priority | options |  | No |  |
| Status | options |  | No |  |
| Planned Start Date | dateTime |  | Yes |  |
| Planned End Date | dateTime |  | Yes |  |
| Additional Fields | collection | ID of the department initiating the release. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Release ID | string | ID of the release to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Release ID | string | ID of the release to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Release ID | string | ID of the release to update | Yes |  |
| Update Fields | collection | ID of the department initiating the release. Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: requester

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| First Name | string |  | Yes |  |
| Primary Email | string |  | No |  |
| Additional Fields | collection | Comma-separated IDs of the departments associated with the requester. Choose from the list or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Requester ID | string | ID of the requester to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Requester ID | string | ID of the requester to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Requester ID | string | ID of the requester to update | Yes |  |
| Update Fields | collection | Comma-separated IDs of the departments associated with the requester. Choose from the list or specify an ID using an <a href= | No |  |

### Resource: requesterGroup

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string |  | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Requester Group ID | string | ID of the requester group to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Requester Group ID | string | ID of the requester group to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Requester Group ID | string | ID of the requester group to update | Yes |  |
| Update Fields | collection | Description of the requester group | No |  |

### Resource: software

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Application Type | options |  | Yes |  |
| Name | string |  | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Software ID | string | ID of the software application to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Software ID | string | ID of the software application to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Software ID | string | ID of the software application to update | Yes |  |
| Update Fields | collection | Type of the software | No |  |

### Resource: ticket

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | Email address of the ticket author | Yes |  |
| Subject | string |  | No |  |
| Description | string | HTML supported | No |  |
| Priority | options |  | No |  |
| Status | options |  | No |  |
| Additional Fields | collection | Comma-separated email addresses to add in the CC field of the ticket email | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string | ID of the ticket to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string | ID of the ticket to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | ID of the agent to whom the tickets have been assigned. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Ticket ID | string | ID of the ticket to update | Yes |  |
| Update Fields | collection | ID of the department to which this ticket has been assigned. Choose from the list, or specify an ID using an <a href= | No |  |

## UseCases

- **IT Service Desk Automation**: Automate ticket creation, assignment, and escalation for IT support requests and incident management
- **Asset Management**: Track and manage IT assets, software applications, and hardware inventory across the organization
- **Change Management**: Automate change request workflows, approvals, and implementation tracking for IT infrastructure changes
- **Problem Management**: Create and track problem records, root cause analysis, and resolution workflows for recurring issues
- **Release Management**: Coordinate software releases, deployment schedules, and rollback procedures through automated workflows
- **Agent and Department Management**: Automate user provisioning, role assignments, and departmental organization for support teams
- **ITIL Process Automation**: Implement ITIL best practices through automated incident, problem, change, and release management processes
- **Service Catalog Integration**: Automate service request fulfillment and approval workflows for business services
- **Asset Lifecycle Management**: Track asset deployment, maintenance schedules, and retirement processes
- **Multi-location Operations**: Manage IT services across multiple locations with centralized asset and user management

