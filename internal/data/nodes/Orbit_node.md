# Orbit Node

## Overview

Consume Orbit API to manage community engagement and member relationships. Orbit is a community management platform that helps organizations build stronger relationships with their community members by tracking activities, managing member profiles, and facilitating meaningful interactions across multiple channels and platforms.

## Credentials

- Name: orbitApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: activity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |
| Title | string |  | Yes |  |
| Additional Fields | collection | A user-defined way to group activities of the same nature. Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | When set the post will be filtered by the member ID | No |  |

### Resource: member

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |
| Resolve Identities | boolean | By default, the response just includes the reference of the identity. When set to true the identities will be resolved automatically. | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Resolve Identities | boolean | By default, the response just includes the reference of the identity. When set to true the identities will be resolved automatically. | No |  |
| Options | collection | Name of the field the response will be sorted by | No |  |

#### Operation: lookup

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Source | options | Set to github, twitter, email, discourse or the source of any identities you | Yes |  |
| Search By | options |  | Yes |  |
| ID | string | The username at the source | Yes |  |
| Username | string | The username at the source | Yes |  |
| Email | string | The email address | Yes |  |
| Host | string |  | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |
| Update Fields | collection | Adds tags to member; comma-separated string or array | No |  |

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Identity | fixedCollection | The identity is used to find the member. If no member exists, a new member will be created and linked to the provided identity. | Yes |  |
| Additional Fields | collection | Adds tags to member; comma-separated string or array | No |  |

### Resource: note

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |
| Note | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Resolve Member | boolean |  | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |
| Note ID | string |  | Yes |  |
| Note | string |  | Yes |  |

### Resource: post

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |
| URL | string | Supply any URL and Orbit will do its best job to parse out a title, description, and image | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | When set the post will be filtered by the member ID | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Workspace Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Member ID | string |  | Yes |  |
| Post ID | string |  | Yes |  |

## UseCases

- **Community Member Management**: Track and manage community member profiles, including their activities, contributions, and engagement levels across different platforms
- **Developer Relations Automation**: Automate tracking of developer community interactions, contributions to open source projects, and engagement with documentation and forums
- **Customer Success Workflows**: Monitor customer activities, track support interactions, and identify opportunities for deeper engagement and relationship building
- **Event and Meetup Coordination**: Manage attendee information, track participation in community events, and follow up with participants for continued engagement
- **Content Creator Recognition**: Track content creators and influencers in your community, monitor their contributions, and automate recognition and reward programs
- **Community Health Analytics**: Analyze member engagement patterns, identify top contributors, and generate insights about community growth and health metrics
- **Cross-Platform Activity Aggregation**: Consolidate member activities from GitHub, Twitter, Discord, forums, and other platforms into a unified community view
- **Automated Onboarding Workflows**: Create automated sequences for new community members including welcome messages, resource sharing, and initial engagement tracking
- **Advocacy Program Management**: Identify and nurture community advocates, track their activities, and manage advocacy program participation and rewards
- **Support Ticket Integration**: Link support interactions with community member profiles to provide context and improve customer relationship management

