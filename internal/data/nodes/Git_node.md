# Git

## Overview

The Git node provides comprehensive Git version control operations within n8n workflows, enabling automated repository management, code deployment, and version control integration. This node supports all essential Git operations including clone, commit, push, pull, and repository status management.

## Credentials

Git operations support optional authentication for remote repositories:

- **Credential Name**: `gitPassword`
- **Required Fields**: Username and Password/Token for private repositories and push operations

## Inputs

Git operations accept input data for repository paths, commit messages, and file paths from previous workflow nodes.

## Outputs

Git operations return operation results including status information, commit logs, and configuration data to subsequent workflow nodes.

## Properties

### Resource: Git Repository Operations

#### Operation: Add

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Git operation to perform | Yes |
| Repository Path | String | Local path to Git repository | Yes |
| Paths to Add | String | Comma-separated list of file/folder paths to stage | Yes |

#### Operation: Clone

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Git operation to perform | Yes |
| Authentication | Options | Authentication method (None, Authenticate) | No |
| Source Repository | String | Remote repository URL | Yes |
| New Repository Path | String | Local destination path | Yes |

#### Operation: Commit

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Git operation to perform | Yes |
| Repository Path | String | Local path to Git repository | Yes |
| Message | String | Commit message | Yes |
| Paths to Add | String | Specific files to stage and commit | No |

#### Operation: Push

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Git operation to perform | Yes |
| Authentication | Options | Authentication method (None, Authenticate) | No |
| Repository Path | String | Local path to Git repository | Yes |
| Target Repository | String | Remote repository URL if different from origin | No |

#### Operation: Log

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Git operation to perform | Yes |
| Repository Path | String | Local path to Git repository | Yes |
| Return All | Boolean | Return all commits or limit results | No |
| Limit | Number | Maximum number of commits to return | No |
| File | String | Show history for specific file | No |

#### Operation: Status

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Git operation to perform | Yes |
| Repository Path | String | Local path to Git repository | Yes |

## UseCases

- **Automated Deployments**: Trigger deployment workflows based on Git repository changes
- **CI/CD Integration**: Integrate Git operations into continuous integration and deployment pipelines
- **Code Synchronization**: Automatically sync code changes between repositories and environments
- **Version Control Automation**: Automate version tagging and release management processes
- **Repository Management**: Manage multiple repositories from centralized n8n workflows
- **Backup and Archival**: Create automated backup workflows for critical code repositories
- **Development Workflow**: Automate common development tasks like branch creation and merging

