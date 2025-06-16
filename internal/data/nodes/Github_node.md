# GitHub Node

## Overview

Consume GitHub API

## Credentials

- Name: githubApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: workflow

#### Operation: dispatchAndWait

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Your execution will pause until a webhook is called. This URL will be generated at runtime and passed to your Github workflow as a resumeUrl input. | notice |  | No |  |

### Resource: file

#### Operation: list

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Path | string | The file path of the file. Has to contain the full path. | Yes |  |
| Path | string | The path of the folder to list | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| As Binary Property | boolean | Whether to set the data of the file as binary property instead of returning the raw API response | No |  |
| Put Output File in Field | string |  | Yes |  |
| Additional Parameters | collection | Additional fields to add | No |  |

### Resource: issue

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | The title of the issue | Yes |  |
| Body | string | The body of the issue | No |  |
| Labels | collection | Label to add to issue | No |  |
| Assignees | collection | User to assign issue too | No |  |

#### Operation: createComment

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Number | number | The number of the issue on which to create the comment on | Yes |  |
| Body | string | The body of the comment | No |  |

#### Operation: edit

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Number | number | The number of the issue edit | Yes |  |
| Edit Fields | collection | User to assign issue to | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Number | number | The issue number to get data for | Yes |  |

#### Operation: lock

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Number | number | The issue number to lock | Yes |  |
| Lock Reason | options | The issue is Off-Topic | No |  |

### Resource: release

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag | string | The tag of the release | Yes |  |
| Additional Fields | collection | The name of the issue | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | The body of the release | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: repository

#### Operation: getIssues

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Return only issues which are assigned to a specific user | No |  |

#### Operation: getPullRequests

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Returns pull requests with any state | No |  |

### Resource: review

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| PR Number | number | The number of the pull request | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| PR Number | number | The number of the pull request to review | Yes |  |
| Event | options | Approve the pull request | No |  |
| Body | string | The body of the review (required for events Request Changes or Comment) | No |  |
| Additional Fields | collection | The SHA of the commit that needs a review, if different from the latest | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Body | string | The body of the review | No |  |

### Resource: user

#### Operation: getRepositories

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: invite

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Organization | string | The GitHub organization that the user is being invited to | Yes |  |
| Email | string | The email address of the invited user | Yes |  |

### Resource: organization

#### Operation: getRepositories

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **DevOps Automation**: Automate GitHub workflows, trigger builds, and manage deployment processes through GitHub Actions integration
- **Issue Management**: Create, update, and track software issues, bugs, and feature requests for development project management
- **Code Review Automation**: Streamline pull request reviews, approvals, and merge processes for collaborative development
- **Release Management**: Automate software release creation, versioning, and distribution through GitHub Releases
- **Repository Management**: Manage file operations, repository access, and organization-level repository administration
- **Team Collaboration**: Automate user invitations, team management, and access control for development teams
- **CI/CD Integration**: Integrate with continuous integration pipelines for automated testing and deployment workflows
- **Project Documentation**: Automate documentation updates, README management, and repository maintenance tasks
- **Open Source Management**: Manage open source contributions, community engagement, and project visibility
- **Security and Compliance**: Monitor repository security, manage access permissions, and ensure compliance with development standards

