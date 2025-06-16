# Wait Node

## Overview

Pauses workflow execution for a specified duration or until specific conditions are met. Supports time intervals, webhooks, and form submissions.

## Credentials

- httpBasicAuth (Basic Auth), Required: Yes when incomingAuthentication is "basicAuth"
- httpHeaderAuth (Header Auth), Required: Yes when incomingAuthentication is "headerAuth" 
- jwtAuth (JWT Auth), Required: Yes when incomingAuthentication is "jwtAuth"

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Wait Controller

#### Operation: Time Interval

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resume Mode | options | Determines the waiting mode to use | Yes |
| Amount | number | The time to wait | Yes |
| Unit | options | Time unit (seconds, minutes, hours, days) | Yes |

#### Operation: Specific Time

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resume Mode | options | Determines the waiting mode to use | Yes |
| Date Time | dateTime | The date and time to wait for before continuing | Yes |

#### Operation: Webhook Call

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resume Mode | options | Determines the waiting mode to use | Yes |
| Incoming Authentication | options | Authentication for incoming webhook requests | No |
| HTTP Method | options | The HTTP method of the webhook call | No |
| Response Code | number | HTTP response code to return | No |
| Response Mode | options | When and how to respond | No |
| Response Data | options | What data to return | No |
| Response Binary Property Name | string | Binary property name | No |
| Limit Wait Time | boolean | Whether to limit waiting time | No |
| Limit Type | options | Condition for resuming execution | No |
| Resume Amount | number | Time to wait before auto-resume | No |
| Resume Unit | options | Time unit for auto-resume | No |
| Max Date and Time | dateTime | Latest date/time to wait until | No |
| Webhook Suffix | string | Suffix appended to restart URL | No |

#### Operation: Form Submission

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resume Mode | options | Determines the waiting mode to use | Yes |
| Incoming Authentication | options | Authentication for incoming form requests | No |
| Form Title | string | Title displayed on the form | No |
| Form Description | string | Description shown on the form | No |
| Form Fields | collection | Fields to collect in the form | No |
| Form Respond Mode | options | How to respond after form submission | No |
| Limit Wait Time | boolean | Whether to limit waiting time | No |
| Limit Type | options | Condition for resuming execution | No |
| Resume Amount | number | Time to wait before auto-resume | No |
| Resume Unit | options | Time unit for auto-resume | No |
| Max Date and Time | dateTime | Latest date/time to wait until | No |
| Append Attribution | boolean | Add n8n attribution to forms | No |

### Resume Modes
- **Time Interval**: Wait for a specified amount of time
- **Specific Time**: Wait until a specific date and time
- **Webhook Call**: Wait for an incoming webhook call
- **Form Submission**: Wait for form submission

### Time Units
- **Seconds**: Wait for specified seconds
- **Minutes**: Wait for specified minutes
- **Hours**: Wait for specified hours
- **Days**: Wait for specified days

### Authentication Methods
- **None**: No authentication required
- **Basic Auth**: Username and password authentication
- **Header Auth**: Custom header-based authentication
- **JWT Auth**: JSON Web Token authentication

### Response Modes
- **On Received**: Respond immediately when request is received
- **Last Node**: Respond when workflow execution completes
- **Response Node**: Use a dedicated response node to control response

## UseCases

- **Screenshot Generation Wait**: Wait for external screenshot services like Apify to complete image generation, ensuring stable image processing in visual regression testing workflows. Controls intervals between screenshot generations during batch processing to prevent service overload.

