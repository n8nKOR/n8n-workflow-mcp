# GitLab Node

## Overview

Retrieve data from GitLab API

## Credentials

- Name: gitlabApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: issue

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | The title of the issue | Yes |  |
| Body | string | The body of the issue | No |  |
| Due Date | dateTime | Due Date for issue | No |  |
| Labels | collection | Label to add to issue | No |  |
| Assignees | collection | User ID to assign issue to | No |  |

#### Operation: createComment

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Number | number | The number of the issue on which to create the comment on | Yes |  |
| Body | string | The body of the comment | No |  |

#### Operation: edit

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Number | number | The number of the issue edit | Yes |  |
| Edit Fields | collection | The title of the issue | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Number | number | The number of the issue get data of | Yes |  |

#### Operation: lock

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue Number | number | The number of the issue to lock | Yes |  |
| Lock Reason | options | The issue is Off-Topic | No |  |

### Resource: release

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag | string | The tag of the release | Yes |  |
| Additional Fields | collection | The name of the release | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project ID | string | The ID or URL-encoded path of the project | Yes |  |
| Additional Fields | collection | The field to use as order | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Project ID | string | The ID or URL-encoded path of the project | Yes |  |
| Tag Name | string | The Git tag the release is associated with | Yes |  |
| Additional Fields | collection | The release name | No |  |

### Resource: repository

#### Operation: getIssues

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection | Return only issues which are assigned to a specific user | No |  |

### Resource: file

#### Operation: list

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Path | string | The file path of the file. Has to contain the full path or leave it empty for root folder. | No |  |
| Path | string | The path of the folder to list | No |  |
| Page | number | Page of results to display | No |  |
| Additional Parameters | collection | Additional fields to add | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| As Binary Property | boolean | Whether to set the data of the file as binary property instead of returning the raw API response | No |  |
| Put Output File in Field | string |  | Yes |  |
| Additional Parameters | collection | Additional fields to add | No |  |

## UseCases

- **DevOps Pipeline Automation**: Automate issue creation, release management, and deployment workflows for software development projects
- **Issue and Bug Tracking**: Create, update, and manage development issues, bug reports, and feature requests through automated workflows
- **Release Management**: Automate software release creation, tagging, and deployment processes for continuous delivery
- **Code Repository Management**: Automate file operations, repository management, and version control workflows
- **Project Management Integration**: Sync GitLab issues with project management tools for unified development tracking
- **CI/CD Integration**: Trigger builds, deployments, and notifications based on issue status changes and release events
- **Team Collaboration**: Automate issue assignments, comments, and notifications for development team coordination
- **Documentation Management**: Manage project documentation, README files, and code comments through automated updates
- **Quality Assurance**: Integrate issue tracking with testing workflows and quality assurance processes
- **Customer Support Integration**: Connect customer support tickets with development issues for streamlined bug resolution

