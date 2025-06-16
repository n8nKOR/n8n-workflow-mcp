# Demio Node

## Overview

Consume the Demio API

## Credentials

- Name: demioApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: event

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection |  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event ID | string |  | Yes |  |
| Additional Fields | collection | Whether to return only active dates in series | No |  |

#### Operation: register

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| First Name | string | The registrant | Yes |  |
| Email | string | The registrant | Yes |  |
| Additional Fields | collection | The value for the predefined Company field | No |  |

### Resource: report

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event Name or ID | options | Choose from the list, or specify an ID using an <a href= | No |  |
| Session Name or ID | options | ID of the session. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Filters | collection | Filter results by participation status | No |  |

## UseCases

- **Automated Webinar Registration**: Automatically register participants in webinars based on form submissions, course completions, or purchase events
- **Post-Webinar Follow-up Automation**: Trigger different email sequences based on webinar attendance status (attended, no-show, left early)
- **Lead Generation and Qualification**: Use webinar registration and attendance data to qualify and segment leads for sales follow-up
- **Multi-Session Event Management**: Automate registration and tracking for webinar series and recurring educational events
- **CRM Integration and Sync**: Synchronize webinar participant data with CRM systems for unified lead and customer management
- **Event-Driven Marketing Campaigns**: Create targeted marketing campaigns based on webinar topics, attendance patterns, and engagement levels
- **Customer Education Workflows**: Integrate webinar participation with learning management systems and certification processes
- **Sales Pipeline Automation**: Automatically move high-engaged webinar attendees through sales funnels and schedule follow-up calls
- **Attendance Analytics and Reporting**: Generate comprehensive reports on webinar performance, engagement metrics, and participant behavior
- **Personalized Content Delivery**: Send customized resources and follow-up materials based on specific webinar sessions attended

