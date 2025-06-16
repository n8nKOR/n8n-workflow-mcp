# iCalendar Node

## Overview

The iCalendar node generates RFC 5545 compliant calendar files (.ics) that can be imported into any standard calendar application. This node creates event files with comprehensive metadata including attendees, recurrence patterns, location data, and timezone information.

## Credentials

No credentials required - this node generates calendar files locally without external API dependencies.

## Inputs

- **Main**: Input data containing event information and attendee details

## Outputs

- **Main**: Returns binary data containing the generated iCalendar (.ics) file

## Properties

### Resource: Calendar Event

#### Operation: Create iCalendar File

| Field Name | Type | Description | Required |
|---|---|---|---|
| Event Title | String | Title/summary of the calendar event | No |
| Start | DateTime | Event start date and time | Yes |
| End | DateTime | Event end date and time | Yes |
| All Day | Boolean | Mark event as all-day (ignores time components) | No |
| Put Output File in Field | String | Binary field name for the generated .ics file | Yes |
| Description | String | Detailed event description or agenda | No |
| Location | String | Event venue or meeting location | No |
| Calendar Name | String | Calendar name for Apple iCal and Microsoft Outlook | No |
| File Name | String | Name of the generated .ics file | No |
| Attendees | Collection | List of event attendees | No |
| Organizer | Collection | Event organizer information | No |

## UseCases

- **Meeting Scheduling**: Generate calendar invites for meetings and appointments
- **Event Management**: Create calendar events for conferences, workshops, and seminars
- **Automated Reminders**: Generate calendar entries for deadlines and important dates
- **Workflow Integration**: Connect calendar events with other business processes
- **Mass Event Creation**: Bulk generate calendar events from data sources
- **Cross-Platform Calendar**: Create calendar files compatible with all major calendar applications 