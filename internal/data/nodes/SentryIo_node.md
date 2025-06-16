# Sentry.io Node

## Overview

Consume Sentry.io API

## Credentials

- Name: sentryIoOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: event

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the events belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Project Slug Name or ID | options | The slug of the project the events belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Full | boolean | Whether the event payload will include the full event body, including the stack trace | No |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the events belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Project Slug Name or ID | options | The slug of the project the events belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Event ID | string | The ID of the event to retrieve (either the numeric primary-key or the hexadecimal ID as reported by the raven client) | Yes |  |

### Resource: issue

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the issues belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Project Slug Name or ID | options | The slug of the project the issues belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | An optional Sentry structured search query. If not provided, an implied  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue ID | string | ID of issue to get | Yes |  |
| Update Fields | collection | The actor ID (or username) of the user or team that should be assigned to this issue | No |  |

### Resource: organization

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Whether to restrict results to organizations which you have membership | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the team should be created for. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The slug of the organization the team should be created for | Yes |  |
| Agree to Terms | boolean | Whether you agree to the applicable terms of service and privacy policy of Sentry.io | No |  |
| Additional Fields | collection | The unique URL slug for this organization. If this is not provided a slug is automatically generated based on the name. | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Slug Name or ID | options | The slug of the organization to update. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Update Fields | collection | The new name of the organization | No |  |

### Resource: project

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project Slug Name or ID | options | The slug of the project to retrieve. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Slug Name or ID | options | The slug of the team to create a new project for. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string | The name for the new project | Yes |  |
| Additional Fields | collection | Optionally a slug for the new project. If it's not provided a slug is generated from the name. | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the project belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Project Slug Name or ID | options | The slug of the project to update. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Update Fields | collection | The new platform for the updated project | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the project belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Project Slug Name or ID | options | The slug of the project to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |

### Resource: release

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the releases belong to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | This parameter can be used to create a "starts with" filter for the version | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the release belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Version | string | A version identifier for this release. Can be a version number, a commit hash etc. | Yes |  |
| URL | string | A URL that points to the release. This can be the path to an online interface to the sourcecode for instance. | Yes |  |
| Project Names or IDs | multiOptions | A list of project slugs that are involved in this release. Choose from the list, or specify IDs using an <a href= | Yes |  |
| Additional Fields | collection | An optional date that indicates when the release went live. If not provided the current time is assumed. | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the release belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Version | string | A version identifier for this release. Can be a version number, a commit hash etc. | Yes |  |
| Update Fields | collection | An optional list of commit data to be associated with the release | Yes |  |

### Resource: team

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization for which the teams should be listed. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the team belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Team Slug Name or ID | options | The slug of the team to get. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the team belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Name | string | The name of the team | Yes |  |
| Additional Fields | collection | The optional slug for this team. If not provided it will be auto generated from the name. | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the team belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Team Slug Name or ID | options | The slug of the team to update. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Update Fields | collection | The new name of the team | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization Slug Name or ID | options | The slug of the organization the team belongs to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Team Slug Name or ID | options | The slug of the team to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |

## UseCases

- [Error Monitoring] : Monitor and track application errors and exceptions in real-time
- [Performance Monitoring] : Track application performance metrics and identify bottlenecks
- [Issue Management] : Manage and prioritize software issues and bug reports
- [Release Tracking] : Track software releases and monitor their impact on application stability
- [Team Collaboration] : Enable development team collaboration on error resolution
- [Alert Management] : Set up intelligent alerts for critical application issues
- [Debug Information] : Collect detailed debug information for faster issue resolution
- [Performance Optimization] : Identify and optimize application performance issues
- [User Experience Monitoring] : Monitor user experience and track error impact
- [Integration Workflows] : Integrate Sentry with development and deployment workflows
- [Code Quality] : Improve code quality through error analysis and tracking
- [Production Monitoring] : Monitor production applications for stability and reliability
- [Development Support] : Support development teams with error tracking and debugging
- [Incident Response] : Coordinate incident response for application outages and errors
- [Analytics and Reporting] : Generate reports on application health and error trends

