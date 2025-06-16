# Twilio Node

## Overview

Send SMS and WhatsApp messages or make phone calls

## Credentials

- Name: twilioApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: sms

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| To Whatsapp | boolean | Whether the message should be sent to WhatsApp | No |  |
| Message | string | The message to send | Yes |  |

### Resource: call

#### Operation: make

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Use TwiML | boolean | Whether to use the <a href= | No |  |
| Message | string |  | Yes |  |

## UseCases

- [Appointment Lead Management] : Automatically handle SMS-based appointment inquiries with AI-powered conversation analysis, lead qualification, and integration with calendar booking systems like Cal.com
- [Automated Follow-up Campaigns] : Send scheduled follow-up SMS messages to leads and customers based on predefined criteria, tracking engagement and managing communication frequency
- [Customer Support Notifications] : Send SMS alerts and updates for support tickets, order status changes, and important account notifications to keep customers informed
- [Two-Factor Authentication] : Implement secure SMS-based verification codes for user authentication and account security in web applications and services
- [Emergency Alert System] : Send urgent notifications and alerts to users via SMS for system outages, security incidents, or time-sensitive business communications
- [Marketing Campaign Automation] : Execute SMS marketing campaigns with personalized messages, promotional offers, and automated responses based on customer behavior and preferences

