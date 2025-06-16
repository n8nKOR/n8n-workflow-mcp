# Zoom Node

## Overview

Consume Zoom API

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: meeting

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Topic | string | Topic of the meeting | No |  |
| Additional Fields | collection | Meeting agenda | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | Meeting ID | Yes |  |
| Additional Fields | collection | To view meeting details of a particular occurrence of the recurring meeting | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | This includes all valid past meetings, live meetings and upcoming scheduled meetings | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | Meeting ID | Yes |  |
| Additional Fields | collection | Meeting occurrence ID | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| ID | string | Meeting ID | Yes |  |
| Update Fields | collection | Meeting agenda | No |  |

### Resource: meetingRegistrant

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Meeting ID | string |  | Yes |  |
| Email | string | Valid Email-ID | Yes |  |
| First Name | string |  | Yes |  |
| Additional Fields | collection | Valid address of registrant | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Meeting ID | string |  | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Registrant Status | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Meeting ID | string |  | Yes |  |
| Action | options | Registrant Status | Yes |  |
| Additional Fields | collection |  | No |  |

### Resource: webinar

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | User ID or email ID | Yes |  |
| Additional Fields | collection | Webinar agenda | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar ID | string |  | Yes |  |
| Additional Fields | collection | To view webinar details of a particular occurrence of the recurring webinar | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | User ID or email-ID | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Webinar occurrence ID | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar ID | string | User ID or email address of user | Yes |  |
| Additional Fields | collection | Webinar agenda | No |  |

### Resource: webinarId

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar ID | string |  | Yes |  |

## UseCases

- [Meeting Automation] : Automate meeting creation and scheduling for recurring events
- [Participant Management] : Manage meeting participants and attendance tracking
- [Recording Management] : Automatically download and organize meeting recordings
- [Meeting Analytics] : Analyze meeting participation and engagement metrics
- [Integration Workflows] : Integrate Zoom with CRM and calendar systems
- [Event Management] : Manage large-scale webinars and virtual events
- [Attendance Reporting] : Generate attendance reports for training and compliance
- [Meeting Notifications] : Send automated meeting reminders and follow-ups
- [Room Management] : Manage Zoom rooms and conference room scheduling
- [User Administration] : Automate user provisioning and account management
- [Cloud Recording Processing] : Process and distribute cloud recordings automatically
- [Meeting Templates] : Create standardized meeting templates for different purposes
- [Survey and Polling] : Collect feedback through post-meeting surveys and polls
- [Calendar Integration] : Sync meeting schedules with calendar applications
- [Quality Monitoring] : Monitor meeting quality and performance metrics

