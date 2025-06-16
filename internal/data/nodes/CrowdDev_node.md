# crowd.dev Node

## Overview

crowd.dev is an open-source suite of community and data tools built to unlock community-led growth for your organization. The crowd.dev node allows you to manage community members, activities, organizations, notes, tasks, and automations within your crowd.dev platform.

## Credentials

- Name: crowdDevApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Activity

#### Operation: Create or Update with Member
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | string | Platform-specific username | Yes | - |
| Type | string | Type of activity | Yes | - |
| Timestamp | dateTime | Date and time when the activity took place | Yes | - |
| Platform | string | Platform on which the activity took place | Yes | - |
| Source ID | string | The ID of the activity in the platform | Yes | - |
| Display Name | string | UI friendly name of the member | No | - |
| Emails | collection | Email addresses of the member | No | - |
| Joined At | dateTime | Date of joining the community | No | - |
| Title | string | Title of the activity | No | - |
| Body | string | Body of the activity | No | - |
| Channel | string | Channel of the activity | No | - |
| Source Parent ID | string | The ID of the parent activity in the platform | No | - |

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Member | string | The ID of the member that performed the activity | Yes | - |
| Type | string | Type of activity | Yes | - |
| Timestamp | dateTime | Date and time when the activity took place | Yes | - |
| Platform | string | Platform on which the activity took place | Yes | - |
| Source ID | string | The ID of the activity in the platform | Yes | - |
| Title | string | Title of the activity | No | - |
| Body | string | Body of the activity | No | - |
| Channel | string | Channel of the activity | No | - |
| Source Parent ID | string | The ID of the parent activity in the platform | No | - |

### Resource: Member

#### Operation: Create or Update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Platform | string | Platform for which to check member existence | Yes | - |
| Username | string | Username of the member in platform | Yes | - |
| Display Name | string | UI friendly name of the member | No | - |
| Emails | collection | Email addresses of the member | No | - |
| Joined At | dateTime | Date of joining the community | No | - |
| Organizations | collection | Organizations associated with the member | No | - |
| Tags | collection | Tags associated with the member | No | - |

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the member | Yes | - |

#### Operation: Find
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the member | Yes | - |

#### Operation: Update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the member | Yes | - |
| Platform | string | Platform for which to check member existence | Yes | - |
| Username | string | Username of the member in platform | Yes | - |
| Display Name | string | UI friendly name of the member | No | - |
| Emails | collection | Email addresses of the member | No | - |

### Resource: Note

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Body | string | The body of the note | No | - |

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the note | Yes | - |

#### Operation: Find
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the note | Yes | - |

#### Operation: Update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the note | Yes | - |
| Body | string | The body of the note | No | - |

### Resource: Organization

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The name of the organization | Yes | - |
| URL | string | The URL of the organization | No | - |
| Description | string | A short description of the organization | No | - |
| Logo | string | A URL for logo of the organization | No | - |
| Employees | number | The number of employees of the organization | No | - |
| Members | collection | Members associated with the organization | No | - |

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the organization | Yes | - |

#### Operation: Find
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the organization | Yes | - |

#### Operation: Update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the organization | Yes | - |
| Name | string | The name of the organization | Yes | - |
| URL | string | The URL of the organization | No | - |
| Description | string | A short description of the organization | No | - |
| Logo | string | A URL for logo of the organization | No | - |
| Employees | number | The number of employees of the organization | No | - |

### Resource: Task

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The name of the task | No | - |
| Body | string | The body of the task | No | - |
| Status | string | The status of the task | No | - |
| Members | collection | Members associated with the task | No | - |
| Activities | collection | Activities associated with the task | No | - |
| Assignees | string | Users assigned with the task | No | - |

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the task | Yes | - |

#### Operation: Find
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the task | Yes | - |

#### Operation: Update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the task | Yes | - |
| Name | string | The name of the task | No | - |
| Body | string | The body of the task | No | - |
| Status | string | The status of the task | No | - |

### Resource: Automation

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Trigger | options | What will trigger an automation (New Activity, New Member) | Yes | - |
| URL | string | URL to POST webhook data to | Yes | - |

#### Operation: Destroy
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the automation | Yes | - |

#### Operation: Find
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the automation | Yes | - |

#### Operation: List
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| (No additional fields required) | - | - | - | - |

#### Operation: Update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | The ID of the automation | Yes | - |
| Trigger | options | What will trigger an automation (New Activity, New Member) | Yes | - |
| URL | string | URL to POST webhook data to | Yes | - |

## UseCases

- **Community Management**: Manage community members and activities
- **Activity Tracking**: Track and analyze community engagement
- **Task Management**: Create and manage community tasks
- **Organization Management**: Manage community organizations
- **Automation**: Set up automated workflows for community events
- **Note Taking**: Create and manage community notes
