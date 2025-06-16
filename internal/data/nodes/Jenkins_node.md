# Jenkins Node

## Overview

Consume Jenkins API

## Credentials

- Name: jenkinsApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: job

#### Operation: triggerParams

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Make sure the job is setup to support triggering with parameters. <a href= | notice |  | No |  |
| Parameters | fixedCollection | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| XML | string | XML of Jenkins config | Yes |  |
| To get the XML of an existing job, add ‘config.xml’ to the end of the job URL | notice |  | No |  |

### Resource: instance

#### Operation: quietDown

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Reason | string | Freeform reason for quiet down mode | No |  |

### Resource: build

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Job Name or ID | options | Name of the job. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **CI/CD Pipeline Automation**: Automate continuous integration and deployment pipelines with job triggering and build management
- **Build Process Management**: Trigger builds, monitor build status, and manage build artifacts for software development projects
- **Automated Testing**: Integrate automated testing workflows with parameterized job execution and result tracking
- **Deployment Orchestration**: Coordinate deployment processes across multiple environments with automated job triggering
- **Infrastructure as Code**: Manage infrastructure provisioning and configuration through Jenkins job automation
- **Development Workflow Integration**: Connect Jenkins with version control systems, issue tracking, and project management tools
- **Release Management**: Automate software release processes, versioning, and deployment coordination
- **Environment Management**: Control development, staging, and production environment deployments through automated jobs
- **Monitoring and Alerting**: Monitor build processes, send notifications, and manage Jenkins instance health
- **Multi-team Coordination**: Manage complex build pipelines across multiple development teams and projects

