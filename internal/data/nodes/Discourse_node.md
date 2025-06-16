# Discourse Node

## Overview

Consume Discourse API

## Credentials

- Name: discourseApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: category

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the category | Yes |  |
| Color | color | Color of the category | Yes |  |
| Text Color | color | Text color of the category | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Category ID | string | ID of the category | Yes |  |
| Name | string | New name of the category | Yes |  |
| Update Fields | collection | Color of the category | No |  |

### Resource: group

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Group ID | string | ID of the group to update | Yes |  |
| Name | string | New name of the group | Yes |  |

### Resource: post

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the post | No |  |
| Content | string | Content of the post | Yes |  |
| Additional Fields | collection | ID of the category. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | ID of the post | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | ID of the post | Yes |  |
| Content | string | Content of the post. HTML is supported. | Yes |  |
| Update Fields | collection |  | No |  |

### Resource: search

#### Operation: query

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Term | string | Term to search for | Yes |  |
| Simplify | boolean | Whether to return a simplified version of the response instead of the raw data | No |  |

### Resource: user

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the user to create | Yes |  |
| Email | string | Email of the user to create | Yes |  |
| Username | string | The username of the user to create | Yes |  |
| Password | string | The password of the user to create | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| By | options | What to search by | Yes |  |
| Username | string | The username of the user to return | Yes |  |
| SSO External ID | string | Discourse SSO external ID | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Flag | options | User flags to search for | No |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to sort ascending | No |  |

### Resource: userGroup

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Usernames | string | Usernames to add to group. Multiples can be defined separated by comma. | Yes |  |
| Group ID | string | ID of the group | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Usernames | string | Usernames to remove from group. Multiples can be defined separated by comma. | Yes |  |
| Group ID | string | ID of the group to remove | Yes |  |

## UseCases

- **Community Management Automation**: Automatically create posts, manage categories, and moderate community discussions
- **Customer Support Integration**: Create and manage support topics, track user questions, and provide automated responses
- **Content Publishing**: Automatically publish content, announcements, and updates to community forums
- **User Onboarding**: Automate user registration, group assignments, and welcome message workflows
- **Knowledge Base Management**: Create and organize knowledge base articles, FAQs, and documentation within community forums
- **Discussion Analytics**: Track forum activity, analyze user engagement, and generate community health reports
- **Moderation Workflows**: Implement automated moderation rules, user group management, and content oversight
- **Cross-platform Integration**: Sync discussions with external systems like CRM, help desk, or project management tools
- **Community Engagement**: Trigger community activities based on user behavior, create discussion topics, and manage user participation
- **Content Search and Discovery**: Implement advanced search functionality and content recommendation systems for forum users

