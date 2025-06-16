# GoToWebinar Node

## Overview

Consume the GoToWebinar API

## Credentials

- Name: goToWebinarOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: attendee

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Registrant Key | string | Registrant key of the attendee at the webinar session | Yes |  |

#### Operation: getDetails

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Registrant Key | string | Registrant key of the attendee at the webinar session | Yes |  |
| Details | options | The details to retrieve for the attendee | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: coorganizer

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar that the co-organizer is hosting. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Is External | boolean | Whether the co-organizer has no GoToWebinar account | Yes |  |
| Organizer Key | string | The co-organizer | No |  |
| Given Name | string | The co-organizer | No |  |
| Email | string | The co-organizer | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Co-Organizer Key | string | Key of the co-organizer to delete | No |  |
| Is External | boolean | By default only internal co-organizers (with a GoToWebinar account) can be deleted. If you want to use this call for external co-organizers you have to set this parameter to  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar to retrieve all co-organizers from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: reinvite

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key | string | By default only internal co-organizers (with a GoToWebinar account) can be deleted. If you want to use this call for external co-organizers you have to set this parameter to  | Yes |  |
| Co-Organizer Key | string | Key of the co-organizer to reinvite | No |  |
| Is External | boolean | Whether the co-organizer has no GoToWebinar account | Yes |  |

### Resource: panelist

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the panelist to create | Yes |  |
| Email | string | Email address of the panelist to create | Yes |  |
| Webinar Key Name or ID | options | Key of the webinar that the panelist will present at. Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar to retrieve all panelists from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar to delete the panelist from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Panelist Key | string | Key of the panelist to delete | Yes |  |

#### Operation: reinvite

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar to reinvite the panelist to. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Panelist Key | string | Key of the panelist to reinvite | Yes |  |

### Resource: registrant

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar of the registrant to create. Choose from the list, or specify an ID using an <a href= | Yes |  |
| First Name | string | First name of the registrant to create | No |  |
| Last Name | string | Last name of the registrant to create | No |  |
| Email | string | Email address of the registrant to create | No |  |
| Additional Fields | collection | Full address of the registrant to create | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | The key of the webinar to retrieve registrants from. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar of the registrant to delete. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Registrant Key | string | Key of the registrant to delete | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key Name or ID | options | Key of the webinar of the registrant to retrieve. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Registrant Key | string | Key of the registrant to retrieve | Yes |  |

### Resource: session

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Start of the datetime range for the session | Yes |  |

#### Operation: getDetails

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Details | options | Performance details for a webinar session | No |  |

### Resource: webinar

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subject | string |  | Yes |  |
| Time Range | fixedCollection |  | Yes |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key | string | Key of the webinar to delete | Yes |  |
| Additional Fields | collection |  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key | string | Key of the webinar to retrieve | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Start of the datetime range for the webinar | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Webinar Key | string | Key of the webinar to update | Yes |  |
| Notify Participants | boolean |  | Yes |  |
| Update Fields | collection | Whether the webinar may be watched anytime | Yes |  |

## UseCases

- **Webinar Event Management**: Automate webinar creation, scheduling, and management for online events and training sessions
- **Attendee Registration**: Streamline webinar registration processes with automated attendee management and communication
- **Multi-presenter Coordination**: Manage co-organizers and panelists for complex webinar presentations and events
- **Attendance Tracking**: Monitor webinar attendance, engagement metrics, and generate post-event analytics reports
- **Automated Follow-up**: Create automated follow-up workflows for webinar attendees based on participation and engagement
- **Training and Education**: Deliver online training programs, educational content, and certification courses
- **Marketing Events**: Host marketing webinars, product demonstrations, and customer engagement events
- **Sales Enablement**: Conduct sales presentations, product demos, and lead generation webinars
- **Customer Support**: Provide customer training, support webinars, and product onboarding sessions
- **Event Analytics**: Track webinar performance, attendee behavior, and ROI for webinar marketing campaigns

