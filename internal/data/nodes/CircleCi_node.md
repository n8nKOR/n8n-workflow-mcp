# CircleCI Node

## Overview

Consume CircleCI API v2. This node provides access to CircleCI's pipeline management functionality, allowing you to trigger, retrieve, and monitor CI/CD pipelines programmatically.

## Credentials

- **Name**: circleCiApi
- **Required**: Yes
- **Type**: Personal API Token
- **Description**: CircleCI Personal API Token for authentication

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: pipeline

The pipeline resource allows you to interact with CircleCI pipelines for your projects.

#### Operation: get

Retrieve a specific pipeline by its number.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Provider | options | Source control system (GitHub, Bitbucket) | Yes | - |
| Project Slug | string | Project slug in the form org-name/repo-name (e.g., "n8n-io/n8n") | Yes | - |
| Pipeline Number | number | The number of the pipeline to retrieve | Yes | 1 |

**API Endpoint**: `GET /project/{vcs}/{project-slug}/pipeline/{pipeline-number}`

#### Operation: getAll

Retrieve multiple pipelines for a project with optional filtering.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Provider | options | Source control system (GitHub, Bitbucket) | Yes | - |
| Project Slug | string | Project slug in the form org-name/repo-name (e.g., "n8n-io/n8n") | Yes | - |
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return (1-500) | No | 100 |
| Filters | collection | Optional filters to apply | No | - |
| └─ Branch | string | Filter by VCS branch name | No | - |

**API Endpoint**: `GET /project/{vcs}/{project-slug}/pipeline`

#### Operation: trigger

Trigger a new pipeline for a project.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Provider | options | Source control system (GitHub, Bitbucket) | Yes | - |
| Project Slug | string | Project slug in the form org-name/repo-name (e.g., "n8n-io/n8n") | Yes | - |
| Additional Fields | collection | Optional parameters for pipeline triggering | No | - |
| └─ Branch | string | The branch where the pipeline should run. The HEAD commit on this branch will be used. Mutually exclusive with tag. | No | - |
| └─ Tag | string | The tag for the pipeline. The commit that this tag points to will be used. Mutually exclusive with branch. | No | - |

**API Endpoint**: `POST /project/{vcs}/{project-slug}/pipeline`

## UseCases

- **CI/CD Pipeline Management**: Monitor and trigger build pipelines for continuous integration and deployment workflows
- **Automated Testing**: Trigger pipeline runs for automated testing when specific conditions are met in your workflow
- **Release Management**: Coordinate software releases by triggering pipelines based on branch updates or tag creation
- **Build Status Monitoring**: Track pipeline status and notify teams of build failures or successes
- **Multi-Environment Deployments**: Trigger pipelines for different environments (staging, production) based on workflow conditions
- **Quality Gates**: Implement quality gates by monitoring pipeline results before proceeding with deployments

## Implementation Notes

This node currently implements only the **Pipeline** resource from the CircleCI API v2. The CircleCI API offers many additional resources that are not yet implemented in this node:

### Available but Not Implemented Resources:
- **Workflows**: Get workflow information, approve, cancel, and rerun workflows
- **Jobs**: Cancel jobs, get job details, retrieve artifacts and test metadata  
- **Projects**: Create and manage projects, handle checkout keys and environment variables
- **Contexts**: Create, list, get, and delete contexts and their environment variables
- **Insights**: Get project and organization-level metrics and analytics
- **Webhooks**: Create, list, update, and delete webhooks
- **Schedules**: Create, list, and manage pipeline schedules
- **Users**: Get current user information and collaborations

### Technical Details:
- Uses CircleCI API v2 (`https://circleci.com/api/v2`)
- Authentication via Personal API Token in header: `Circle-Token: {token}`
- Project slugs must be URL-encoded (slashes converted to `%2F`)
- Supports pagination with `page-token` parameter for `getAll` operations
- VCS providers supported: GitHub and Bitbucket

