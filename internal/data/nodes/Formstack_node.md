# Formstack Trigger

## Overview

The Formstack Trigger node starts the workflow when a form submission is received from Formstack, a powerful form builder platform. This trigger node monitors specified Formstack forms and automatically initiates workflows when new submissions are detected through webhooks, enabling real-time response to form submissions.

## Credentials

This node supports two authentication methods:

### Access Token Authentication
- **Credential Name**: `formstackApi`
- **Required Fields**: Access Token

### OAuth2 Authentication  
- **Credential Name**: `formstackOAuth2Api`
- **Required Fields**: OAuth2 configuration

## Inputs

This node is a trigger node and has no inputs.

## Outputs

- **Main Output**: Outputs the form submission data received via webhook from Formstack.

## Properties

### Resource: Form Trigger

#### Operation: Form Submission

| Field Name | Type | Description | Required |
|---|---|---|---|
| Form | Options | Formstack form to monitor | Yes |
| Resolve Data | Boolean | Resolve form field names to human-readable labels | No |
| Options | Collection | Additional configuration options | No |
## UseCases

- **Lead Management**: Automatically register lead information in CRM systems immediately upon form submission
- **Customer Support**: Automatically create support tickets when inquiry forms are submitted
- **Email Marketing**: Automatically add subscribers to email lists when newsletter signup forms are submitted
- **Order Processing**: Automatically initiate order processing workflows when order forms are submitted
- **Event Registration**: Automatically update attendee management systems when event registration forms are submitted
- **Feedback Collection**: Automatically send form submission data to analysis systems for feedback processing
- **Notification Systems**: Send instant notifications to relevant team members when important forms are submitted
- **Data Synchronization**: Synchronize form data across multiple systems in real-time 