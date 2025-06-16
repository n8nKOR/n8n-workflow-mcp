# Asana Node

## Overview

The Asana node allows you to consume the Asana REST API to manage projects, tasks, subtasks, and team collaboration. Asana is a web and mobile work management platform designed to help teams track their work. This node enables you to automate task management, project tracking, and team coordination within your n8n workflows.

## Credentials

This node supports two authentication methods:

### Access Token
- **API Token**: Your personal Asana access token

To obtain an access token:
1. Log into your Asana account
2. Go to Profile Settings → Apps → Developer Apps
3. Create a personal access token
4. Copy the token for use in n8n

### OAuth2
- Uses OAuth2 flow for authentication
- More secure for production environments
- Requires app registration in Asana

## Inputs

- **Main**: JSON data for creating/updating Asana resources

## Outputs

- **Main**: Returns JSON data from Asana API responses

## Properties

### Resource: Project

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Team | Options | Team to create the project in | Yes |
| Name | String | Name of the project | Yes |
| Additional Fields | Collection | Optional project properties | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Project ID | String | ID of the project to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Project ID | String | ID of the project to retrieve | Yes |
| Options | Collection | Additional query options | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |
| Filters | Collection | Project filtering options | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Project ID | String | ID of the project to update | Yes |
| Update Fields | Collection | Project properties to update | No |

### Resource: Task

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Name of the task | Yes |
| Additional Fields | Collection | Optional task properties | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the task to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the task to retrieve | Yes |
| Options | Collection | Additional query options | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |
| Filters | Collection | Task filtering options | No |

#### Operation: Search
| Field Name | Type | Description | Required |
|---|---|---|---|
| Workspace | Options | Workspace to search in | Yes |
| Additional Options | Collection | Search parameters | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the task to update | Yes |
| Update Fields | Collection | Task properties to update | No |

### Resource: Subtask

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Parent Task ID | String | ID of the parent task | Yes |
| Name | String | Name of the subtask | Yes |
| Additional Fields | Collection | Optional subtask properties | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the parent task | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

### Resource: Task Comment

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the task to comment on | Yes |
| Is Text Html | Boolean | Whether the comment is HTML formatted | No |
| Text | String | Comment text content | Yes |

#### Operation: Remove
| Field Name | Type | Description | Required |
|---|---|---|---|
| Comment ID | String | ID of the comment to remove | Yes |

### Resource: Task Project

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the task | Yes |
| Project | Options | Project to add the task to | Yes |
| Additional Fields | Collection | Optional properties | No |

#### Operation: Remove
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the task | Yes |
| Project | Options | Project to remove the task from | Yes |

### Resource: Task Tag

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the task | Yes |
| Tag | Options | Tag to add to the task | Yes |

#### Operation: Remove
| Field Name | Type | Description | Required |
|---|---|---|---|
| Task ID | String | ID of the task | Yes |
| Tag | Options | Tag to remove from the task | Yes |

### Resource: User

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| User ID | String | ID of the user to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Workspace | Options | Workspace to get users from | Yes |
| Return All | Boolean | Whether to return all results | No |
| Limit | Number | Maximum number of results to return | No |

## UseCases

- **Project Management** : Manage projects, tasks, and team collaboration in Asana
- **Task Automation** : Automate task creation, assignment, and status updates
- **Team Collaboration** : Facilitate team collaboration and communication through Asana
- **Workflow Integration** : Integrate Asana with other business tools and systems
- **Progress Tracking** : Track project progress and milestone completion
- **Resource Management** : Manage team resources and workload distribution
- **Client Project Management** : Manage client projects and deliverables
- **Agile Development** : Support agile development processes and sprint management
- **Time Tracking Integration** : Integrate time tracking with project management
- **Reporting and Analytics** : Generate project reports and performance analytics

