# Kitemaker Node

## Overview

Consume the Kitemaker GraphQL API

## Credentials

- Name: kitemakerApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: space

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: user

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: workItem

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the work item to create | Yes |  |
| Space Name or ID | options | ID of the space to retrieve the work items from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Status Name or ID | options | ID of the status to set on the item to create. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Additional Fields | collection | Description of the item to create. Markdown supported. | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Work Item ID | string | ID of the work item to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Space Name or ID | options | ID of the space to retrieve the work items from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Work Item ID | string | ID of the work item to update | Yes |  |
| Update Fields | collection | Description of the item to update. Markdown supported. | No |  |

## UseCases

- [Product Management] : Manage product development lifecycles and feature planning
- [Task Management] : Organize and track development tasks and work items
- [Sprint Planning] : Plan development sprints and manage agile workflows
- [Team Collaboration] : Enable development team collaboration and communication
- [Feature Tracking] : Track feature development from concept to deployment
- [Project Roadmaps] : Create and maintain product roadmaps and timelines
- [Issue Management] : Manage bugs, issues, and technical debt systematically
- [Workflow Automation] : Automate project workflows and status updates
- [Progress Monitoring] : Monitor project progress with visual dashboards and metrics
- [Release Management] : Plan and track software releases and deployments
- [Customer Feedback] : Collect and prioritize customer feedback for product improvements
- [Integration Management] : Integrate with development tools and version control systems
- [Resource Planning] : Plan and allocate development resources across projects
- [Goal Tracking] : Set and track team goals and objectives
- [Performance Analytics] : Analyze team performance and project metrics

