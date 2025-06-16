# Monica CRM Node

## Overview

Consume the Monica CRM API

## Credentials

- Name: monicaCrmApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: activity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity Type Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Contacts | string | Comma-separated list of IDs of the contacts to associate the activity with | Yes |  |
| Happened At | dateTime | Date when the activity happened | Yes |  |
| Summary | string | Brief description of the activity - max 255 characters | Yes |  |
| Additional Fields | collection | Description of the activity - max 100,000 characters | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | string | ID of the activity to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | string | ID of the activity to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | string | ID of the activity to update | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: call

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to associate the call with | Yes |  |
| Called At | dateTime | Date when the call happened | Yes |  |
| Description | string | Description of the call - max 100,000 characters | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Call ID | string | ID of the call to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Call ID | string | ID of the call to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Call ID | string | ID of the call to update | Yes |  |
| Update Fields | collection | Date when the call happened | No |  |

### Resource: contact

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| First Name | string |  | Yes |  |
| Gender Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Additional Fields | collection | Whether the contact has passed away | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Search term to filter results by | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to update | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

### Resource: contactField

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to associate the contact field with | Yes |  |
| Contact Field Type Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Content | string | Content of the contact field - max 255 characters | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact Field ID | string | ID of the contactField to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact Field ID | string | ID of the contact field to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact whose fields to retrieve | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to associate the contact field with | Yes |  |
| Contact Field ID | string | ID of the contact field to update | Yes |  |
| Contact Field Type Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Content | string | Content of the contact field - max 255 characters | Yes |  |

### Resource: contactTag

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to add a tag to | Yes |  |
| Tag Names or IDs | multiOptions | Tags to add to the contact. Choose from the list, or specify IDs using an <a href= | Yes |  |

#### Operation: remove

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to remove the tag from | Yes |  |
| Tag Names or IDs | multiOptions | Tags to remove from the contact. Choose from the list, or specify IDs using an <a href= | Yes |  |

### Resource: conversation

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to associate the conversation with | Yes |  |
| Contact Field Type Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Happened At | dateTime | Date when the conversation happened | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation ID | string | ID of the conversation to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation ID | string | ID of the conversation to retrieve | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation ID | string | ID of the conversation to update | Yes |  |
| Contact Field Type Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Happened At | dateTime | Date when the conversation happened | Yes |  |

### Resource: conversationMessage

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Conversation ID | string | ID of the contact whose conversation | Yes |  |
| Content | string | Content of the message | Yes |  |
| Written At | dateTime | Date when the message was written | Yes |  |
| Written By | options | Author of the message | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message to update | Yes |  |
| Conversation ID | string | ID of the conversation whose message to update | Yes |  |
| Update Fields | collection | ID of the contact to associate the conversationMessage with | No |  |

### Resource: journalEntry

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | string | Title of the journal entry - max 250 characters | Yes |  |
| Content | string | Content of the journal entry - max 100,000 characters | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Journal Entry ID | string | ID of the journal entry to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Journal Entry ID | string | ID of the journal entry to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Journal Entry ID | string | ID of the journal entry to update | Yes |  |
| Update Fields | collection | Content of the journal entry - max 100,000 characters | No |  |

### Resource: note

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to associate the note with | Yes |  |
| Body | string | Body of the note - max 100,000 characters | Yes |  |
| Additional Fields | collection | Whether the note has been favorited | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | string | ID of the note to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | string | ID of the note to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Note ID | string | ID of the note to update | Yes |  |
| Update Fields | collection | Body of the note - max 100,000 characters | No |  |

### Resource: reminder

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to associate the reminder with | No |  |
| Frequency Type | options | Type of frequency of the reminder | Yes |  |
| Recurring Interval | number | Interval for the reminder | No |  |
| Initial Date | dateTime | Date of the reminder | Yes |  |
| Title | string | Title of the reminder - max 100,000 characters | Yes |  |
| Additional Fields | collection | Description about the reminder - Max 100,000 characters | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Reminder ID | string | ID of the reminder to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Reminder ID | string | ID of the reminder to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Reminder ID | string | ID of the reminder to update | Yes |  |
| Update Fields | collection | ID of the contact to associate the reminder with | No |  |

### Resource: tag

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the tag - max 250 characters | Yes |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag ID | string | ID of the tag to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag ID | string | ID of the tag to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tag ID | string | ID of the tag to update | Yes |  |
| Name | string | Name of the tag - max 250 characters | Yes |  |

### Resource: task

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to associate the task with | Yes |  |
| Title | string | Title of the task entry - max 250 characters | Yes |  |
| Additional Fields | collection | Description of the task - max 100,000 characters | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of the task to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of the task to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Task ID | string | ID of the task to update | Yes |  |
| Update Fields | collection | ID of the contact to associate the task with | No |  |

## UseCases

- [Personal Relationship Management] : Manage personal relationships and keep track of friends and family
- [Contact Organization] : Organize personal contacts with detailed information and notes
- [Relationship Tracking] : Track important dates, events, and milestones in relationships
- [Gift Planning] : Plan and track gifts for birthdays, holidays, and special occasions
- [Social Network Management] : Manage your social network and relationship connections
- [Meeting Notes] : Keep notes about conversations and meetings with personal contacts
- [Event Management] : Track and manage personal events and social gatherings
- [Family Tree Management] : Build and maintain family trees and genealogy information
- [Professional Networking] : Manage professional contacts and networking relationships
- [Follow-up Reminders] : Set reminders to follow up with friends and colleagues
- [Personal Analytics] : Analyze your relationship patterns and social connections
- [Contact History] : Maintain a complete history of interactions with each contact
- [Privacy Management] : Keep personal relationship data private and secure
- [Mobile Access] : Access personal CRM data on mobile devices for on-the-go updates
- [Integration Workflows] : Connect Monica with calendar and communication tools

