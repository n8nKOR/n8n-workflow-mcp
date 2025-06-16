# JotForm Trigger

## Overview

The JotForm Trigger node allows you to handle JotForm events via webhooks. This trigger node automatically starts workflows when form submissions occur in your JotForm account. It enables real-time automation based on form responses and provides options to resolve form data with human-readable field names.

## Credentials

This node requires JotForm API credentials:
- **Credential Name**: `jotFormApi`
- **Required Fields**: 
  - API Key: Your JotForm API key

To obtain API credentials:
1. Log into your JotForm account
2. Navigate to Settings > API
3. Generate an API key for integration use

## Inputs

This node is a trigger node and has no inputs.

## Outputs

- **Main**: Returns JSON data containing the form submission information from JotForm

## Properties

### Resource: Form Trigger

#### Operation: Form Submission

| Field Name | Type | Description | Required |
|---|---|---|---|
| Form | Options | JotForm form to monitor | Yes |
| Resolve Data | Boolean | Resolve form field names to human-readable labels | No |
| Options | Collection | Additional configuration options | No |
## UseCases

- **Lead Management**: Automatically add new leads to CRM when contact forms are submitted
- **Customer Support**: Create support tickets when help request forms are filled
- **Event Registration**: Process event registrations and send confirmation emails
- **Survey Analysis**: Collect and analyze survey responses in real-time
- **Newsletter Signup**: Add subscribers to email marketing lists automatically
- **Order Processing**: Handle product orders and initiate fulfillment workflows
- **Feedback Collection**: Process customer feedback and route to appropriate teams
- **Application Processing**: Handle job applications, membership requests, etc.
- **Data Validation**: Validate and clean form data before storing in databases
- **Notification Systems**: Send alerts to teams when important forms are submitted 