# Onfleet Node

## Overview

⚠️ **Field Documentation**: This documentation includes comprehensive field names and descriptions for all operations. The implementation provides detailed field definitions and proper field names as documented.

Comprehensive integration with Onfleet's last-mile delivery management platform. Onfleet provides route optimization, real-time tracking, delivery management, and analytics for logistics operations. This node enables automation of delivery workflows, task management, and fleet coordination.

## Credentials

- **Name**: onfleetApi
- **Required**: Yes

### Credential Configuration
To configure Onfleet API credentials:
1. Log into your Onfleet dashboard
2. Go to Settings → API & Integrations
3. Generate an API key
4. Copy the API key for use in n8n
5. Ensure your API key has appropriate permissions for the operations you need

## Inputs

- **Main**: JSON data for delivery tasks, destinations, teams, and administrative operations

## Outputs

- **Main**: Returns JSON responses with Onfleet data including tasks, deliveries, tracking information, and operation results

## Properties

### Resource: admin

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 64 |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | String | Administrator name | Yes | - |
| Email | String | Administrator email address | Yes | - |
| Phone | String | Administrator phone number | No | - |
| Additional Fields | Collection | Additional administrator configuration | No | - |

**Additional Fields Options:**
- **Role**: Administrator role and permissions
- **Teams**: Teams the administrator manages
- **Metadata**: Custom metadata for the administrator

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Admin ID | String | ID of the administrator to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

### Resource: container

#### Operation: addTask

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Container ID | String | ID of the container (team or worker) | Yes | - |
| Task ID | String | ID of the task to add to container | Yes | - |
| Consider Dependencies | Boolean | Consider task dependencies when adding | No | true |

### Resource: destination

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Address | String | Street address of the destination | Yes | - |
| Latitude | Number | Latitude coordinate | No | - |
| Longitude | Number | Longitude coordinate | No | - |
| Notes | String | Special delivery notes | No | - |
| Additional Fields | Collection | Additional destination configuration | No | - |

**Additional Fields Options:**
- **Apartment**: Apartment or unit number
- **City**: City name
- **State**: State or province
- **Postal Code**: ZIP or postal code
- **Country**: Country code
- **Unparsed**: Unparsed address string

### Resource: hub

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | String | Hub name | Yes | - |
| Address | String | Hub address | Yes | - |
| Teams | Array | Teams assigned to this hub | No | [] |
| Additional Fields | Collection | Additional hub configuration | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Hub ID | String | The ID of the hub object for lookup | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 64 |

### Resource: organization

#### Operation: getDelegatee

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization ID | String | The ID of the organization for delegatee lookup | Yes | - |

### Resource: recipient

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Get By | Options | The variable that is used for looking up a recipient | Yes | id |
| Recipient ID | String | The ID of the recipient object for lookup | Yes | - |
| Name | String | The name of the recipient for lookup | No | - |
| Phone | String | The phone of the recipient for lookup | No | - |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | String | Recipient name | Yes | - |
| Phone | String | Recipient phone number | Yes | - |
| Notes | String | Special notes about recipient | No | - |
| Skip SMS Notifications | Boolean | Skip SMS notifications for this recipient | No | false |
| Additional Fields | Collection | Additional recipient information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Recipient ID | String | The ID of the recipient object for lookup | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Destination | String | Destination ID for the task | Yes | - |
| Recipients | Array | Array of recipient IDs | No | [] |
| Complete After | Number | Earliest completion time (Unix timestamp) | No | - |
| Complete Before | Number | Latest completion time (Unix timestamp) | No | - |
| Pickup Task | Boolean | Whether this is a pickup task | No | false |
| Auto Assign | Object | Auto-assignment configuration | No | - |
| Additional Fields | Collection | Additional task configuration | No | - |

**Additional Fields Options:**
- **Merchant**: Merchant information
- **Executor**: Assigned worker or team
- **Tracking Viewed**: Whether tracking has been viewed
- **Notes**: Task-specific notes
- **Dependencies**: Task dependencies
- **Barcodes**: Associated barcodes

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 64 |
| Filters | Collection | Filters for task queries | No | - |

**Filter Options:**
- **From**: Start time for the range
- **To**: End time for the range
- **Last ID**: Last ID for pagination
- **State**: Task state filter

#### Operation: complete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | String | ID of the task to complete | Yes | - |
| Complete as a Success | Boolean | Whether the task completion is successful | Yes | true |
| Additional Fields | Collection | Completion details | No | - |

**Additional Fields Options:**
- **Completion Details**: Detailed completion information
- **Notes**: Completion notes
- **Photo**: Completion photo
- **Signature**: Digital signature

#### Operation: clone

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | String | ID of the task to clone | Yes | - |
| Override Fields | Collection | Fields to override in the cloned task | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | String | ID of the task to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

### Resource: team

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | String | Team name | Yes | - |
| Workers | Array | Array of worker IDs for the team | Yes | [] |
| Managers | Array | Array of manager IDs for the team | Yes | [] |
| Hub | String | Hub ID associated with the team | No | - |
| Additional Fields | Collection | Additional team configuration | No | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 64 |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team ID | String | ID of the team to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

#### Operation: autoDispatch

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team ID | String | ID of the team for auto-dispatch | Yes | - |
| Additional Fields | Collection | Auto-dispatch configuration | No | - |

#### Operation: getTimeEstimates

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team ID | String | ID of the team for time estimates | Yes | - |
| Filters | Collection | Parameters for time estimation | No | - |

### Resource: webhook

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| URL | String | Webhook endpoint URL | Yes | - |
| Trigger | Options | Event trigger for the webhook | Yes | taskCreated |
| Threshold | Number | Threshold for triggering webhook | No | 0 |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all webhooks | No | false |
| Limit | Number | Max number of results to return | No | 64 |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webhook ID | String | ID of the webhook to delete | Yes | - |

## UseCases

- **Delivery Management**: Manage end-to-end delivery operations including task creation, assignment, and completion tracking
- **Route Optimization**: Optimize delivery routes for maximum efficiency and reduced travel time
- **Real-time Tracking**: Provide real-time tracking capabilities for customers and internal monitoring
- **Fleet Coordination**: Coordinate delivery teams, workers, and vehicles for optimal resource utilization
- **Customer Communication**: Automate customer notifications for delivery updates, ETAs, and completion confirmations
- **Task Automation**: Automate task creation, assignment, and status updates based on business rules
- **Performance Analytics**: Generate reports and analytics on delivery performance, team efficiency, and customer satisfaction
- **Proof of Delivery**: Collect digital signatures, photos, and completion confirmations for delivery verification
- **Integration Workflows**: Connect Onfleet with e-commerce platforms, ERP systems, and customer management tools
- **Webhook Automation**: Set up automated responses to delivery events and status changes
- **Multi-hub Operations**: Manage deliveries across multiple distribution centers and service areas
- **Recipient Management**: Maintain recipient databases with preferences and delivery instructions
- **Team Management**: Organize delivery teams, assign managers, and track worker performance
- **Exception Handling**: Automate handling of delivery exceptions, failures, and customer requests
- **Scalable Operations**: Scale delivery operations efficiently as business volume grows

