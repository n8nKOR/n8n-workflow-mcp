# Jira Software Node

## Overview

Consume Jira Software API

## Credentials

- Name: jiraSoftwareCloudApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: issue

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project | resourceLocator |  | Yes | "{ mode: \"list\"" |
| Issue Type | resourceLocator |  | Yes | "{ mode: \"list\"" |
| Summary | string |  | Yes |  |
| Additional Fields | collection | Choose from the list, or specify IDs using an <a href= | No | "{ mode: \"list\"" |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string |  | Yes |  |
| Update Fields | collection | Value of the field to set | No | "{ mode: \"list\"" |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string |  | Yes |  |
| Delete Subtasks | boolean |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string |  | Yes |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |
| Additional Fields | collection | A list of fields to return for the issue. This parameter accepts a comma-separated list. Use it to retrieve a subset of fields. Allowed values: <code>*all</code> Returns all fields. <code>*navigable</code> Returns navigable fields. Any issue field, prefixed with a minus to exclude. | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Returns a list of recent updates to an issue, sorted by date, starting from the most recent | No |  |

#### Operation: changelog

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string |  | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: notify

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string |  | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | collection | The HTML body of the email notification for the issue | No |  |
| Notification Recipients | fixedCollection | The recipients of the email notification for the issue | No |  |
| Notification Recipients | json | The recipients of the email notification for the issue | No |  |
| Notification Recipients Restrictions | fixedCollection | Restricts the notifications to users with the specified permissions | No |  |
| Notification Recipients Restrictions | json | Restricts the notifications to users with the specified permissions | No |  |

#### Operation: transitions

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string |  | Yes |  |
| Additional Fields | collection | Use expand to include additional information about transitions in the response. This parameter accepts transitions.fields, which returns information about the fields in the transition screen for each transition. Fields hidden from the screen are not returned. Use this information to populate the fields and update fields in Transition issue. | No |  |

### Resource: issueAttachment

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string |  | Yes |  |
| Input Binary Field | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attachment ID | string | The ID of the attachment | Yes |  |
| Download | boolean |  | Yes |  |
| Put Output File in Field | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string |  | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Download | boolean |  | Yes |  |
| Put Output File in Field | string |  | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Attachment ID | string | The ID of the attachment | Yes |  |

### Resource: issueComment

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string | issueComment Key | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Comment | string | Comment | No |  |
| Document Format (JSON) | json | The Atlassian Document Format (ADF). Online builder can be found <a href= | No |  |
| Options | collection | Use expand to include additional information about comments in the response. This parameter accepts Rendered Body, which returns the comment body rendered in HTML. | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string | The ID or key of the issue | Yes |  |
| Comment ID | string | The ID of the comment | Yes |  |
| Options | collection | Use expand to include additional information about comments in the response. This parameter accepts Rendered Body, which returns the comment body rendered in HTML. | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string | The ID or key of the issue | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Use expand to include additional information about comments in the response. This parameter accepts Rendered Body, which returns the comment body rendered in HTML. | No |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string | The ID or key of the issue | Yes |  |
| Comment ID | string | The ID of the comment | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Key | string | The Issue Comment key | Yes |  |
| Comment ID | string | The ID of the comment | Yes |  |
| JSON Parameters | boolean |  | No |  |
| Comment | string | Comment | No |  |
| Document Format (JSON) | json | The Atlassian Document Format (ADF). Online builder can be found <a href= | No |  |
| Options | collection | Use expand to include additional information about comments in the response. This parameter accepts Rendered Body, which returns the comment body rendered in HTML. | No |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | string |  | Yes |  |
| Email Address | string |  | Yes |  |
| Display Name | string |  | Yes |  |
| Additional Fields | collection | Password for the user. If a password is not set, a random password is generated. | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Account ID | string | Account ID of the user to delete | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Account ID | string | Account ID of the user to retrieve | No |  |
| Additional Fields | collection | Include more information about the user | No |  |

## UseCases

- **Project Management Automation**: Automate issue creation, tracking, and management for agile development and project coordination
- **Bug Tracking and Resolution**: Streamline bug reporting, assignment, and resolution workflows for software development teams
- **Agile Development**: Support scrum and kanban methodologies with automated sprint planning, backlog management, and progress tracking
- **Team Collaboration**: Facilitate team communication through issue comments, notifications, and collaborative workflow management
- **Customer Support Integration**: Convert customer support requests into development issues and track resolution progress
- **DevOps Integration**: Connect development issues with CI/CD pipelines, deployments, and release management processes
- **Release Management**: Coordinate software releases, version planning, and delivery tracking across development teams
- **User Management**: Automate user provisioning, access control, and team member management for development projects
- **Reporting and Analytics**: Generate project reports, velocity metrics, and team performance analytics for management visibility
- **Cross-platform Integration**: Sync Jira issues with external tools, repositories, and project management systems

