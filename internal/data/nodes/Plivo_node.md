# Plivo Node

## Overview

Comprehensive communication platform integration with Plivo for sending SMS/MMS messages and making voice calls. Plivo provides cloud-based APIs for global voice and SMS communication, enabling businesses to integrate communication capabilities into their applications and workflows.

**Service Limitations:**
- **MMS Geographic Restriction**: MMS service is available only in the US and Canada

## Credentials

- **Name**: plivoApi
- **Required**: Yes

### Credential Configuration
- **Auth ID**: Your Plivo account Auth ID
- **Auth Token**: Your Plivo account Auth Token

## Inputs

- **Main**: Message content, phone numbers, and communication parameters

## Outputs

- **Main**: Delivery status, message IDs, call information, and API responses

## Properties

### Resource: sms

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| From | string | Plivo phone number to send the SMS from (must be a verified Plivo number) | Yes | - |
| To | string | Destination phone number (E.164 format recommended) | Yes | - |
| Message | string | Text message content to send (up to 1600 characters) | Yes | - |
| Type | options | Message type (sms, unicode) | No | sms |
| URL | string | Callback URL for delivery receipts | No | - |
| Method | options | HTTP method for callback (GET, POST) | No | POST |
| Log | boolean | Whether to log the message | No | true |
| Trackable | boolean | Enable click tracking for URLs | No | false |
| Additional Fields | collection | Additional SMS parameters | No | - |

**Additional Fields:**
- **Powerpack UUID**: Use Powerpack for sending (overrides From number)
- **Media URLs**: URLs of media files (converts SMS to MMS)
- **DLT Entity ID**: DLT entity ID for India SMS
- **DLT Template ID**: DLT template ID for India SMS

### Resource: mms

#### Operation: send

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| From | string | Plivo phone number to send the MMS from (must be MMS-enabled) | Yes | - |
| To | string | Destination phone number (E.164 format recommended) | Yes | - |
| Message | string | Text message content (optional when media is included) | No | - |
| Media URLs | string | Comma-separated list of media file URLs (images, videos, audio) | No | - |
| Type | options | Message type (mms) | No | mms |
| URL | string | Callback URL for delivery receipts | No | - |
| Method | options | HTTP method for callback (GET, POST) | No | POST |
| Log | boolean | Whether to log the message | No | true |
| Additional Fields | collection | Additional MMS parameters | No | - |

**Service Limitations:**
- **Geographic Restriction**: MMS service is available only in the US and Canada
- **Media Requirements**: Supported formats include JPEG, PNG, GIF (images), MP4, 3GP (videos), MP3, AAC (audio)
- **File Size Limits**: Maximum 5MB per media file, 600KB recommended for better delivery
- **Carrier Support**: Not all carriers support MMS; delivery may vary

### Resource: call

#### Operation: make

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| From | string | Caller ID for the outbound call (must be a verified Plivo number) | Yes | - |
| To | string | Destination phone number to call (E.164 format recommended) | Yes | - |
| Answer URL | string | URL that returns XML to handle the call when answered | Yes | - |
| Answer Method | options | HTTP method for Answer URL request (GET, POST) | Yes | POST |
| Hangup URL | string | URL called when the call ends | No | - |
| Hangup Method | options | HTTP method for Hangup URL | No | POST |
| Fallback URL | string | Fallback URL if Answer URL fails | No | - |
| Fallback Method | options | HTTP method for Fallback URL | No | POST |
| Caller Name | string | Caller name displayed to called party | No | - |
| Send Digits | string | DTMF digits to send when call connects | No | - |
| Send On Preanswer | boolean | Send digits on preanswer | No | false |
| Time Limit | number | Maximum call duration in seconds | No | 14400 |
| Timeout | number | Ring timeout in seconds | No | 120 |
| Additional Fields | collection | Additional call parameters | No | - |

**Default Values:**
- **Answer Method**: Default is "POST"

**Additional Fields:**
- **Machine Detection**: Enable answering machine detection
- **Machine Detection Time**: Time to wait for machine detection
- **Machine Detection URL**: URL called when machine is detected
- **Ring URL**: URL called when call starts ringing
- **Ring Method**: HTTP method for Ring URL

### Resource: account

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subaccount | boolean | Whether to retrieve subaccount information | No | false |

**Account Information Retrieved:**
- **Account Balance**: Current account balance and currency
- **Account Type**: Main account or subaccount details
- **Auto Recharge**: Auto-recharge configuration status
- **Cash Credits**: Available cash credits
- **Timezone**: Account timezone configuration
- **Resource URI**: API resource endpoint for the account

### Resource: application

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Application Name | string | Name for the application | Yes | - |
| Answer URL | string | URL to fetch XML when call is answered | Yes | - |
| Answer Method | options | HTTP method for Answer URL (GET, POST) | No | POST |
| Additional Fields | collection | Additional application configuration | No | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Application ID | string | ID of the application to delete | Yes | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Application ID | string | ID of the application to retrieve | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all applications or limit results | No | false |
| Limit | number | Maximum number of applications to return | No | 20 |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Application ID | string | ID of the application to update | Yes | - |
| Update Fields | collection | Fields to update for the application | No | - |

### Resource: conference

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all conferences or limit results | No | false |
| Limit | number | Maximum number of conferences to return | No | 20 |

### Resource: endpoint

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | string | Username for the endpoint | Yes | - |
| Password | string | Password for the endpoint | Yes | - |
| Alias | string | Alias for the endpoint | Yes | - |
| Additional Fields | collection | Additional endpoint configuration | No | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Endpoint ID | string | ID of the endpoint to delete | Yes | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Endpoint ID | string | ID of the endpoint to retrieve | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all endpoints or limit results | No | false |
| Limit | number | Maximum number of endpoints to return | No | 20 |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Endpoint ID | string | ID of the endpoint to update | Yes | - |
| Update Fields | collection | Fields to update for the endpoint | No | - |

### Resource: number

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Number | string | Phone number to retrieve information for | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all numbers or limit results | No | false |
| Limit | number | Maximum number of numbers to return | No | 20 |
| Type | options | Type of number (local, tollfree, mobile) | No | - |
| Number Startswith | string | Filter numbers that start with specific digits | No | - |
| Additional Fields | collection | Additional filtering options | No | - |

### Resource: pricing

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Country ISO | string | Country ISO code for pricing information | Yes | - |

**Pricing Information Retrieved:**
- **SMS Pricing**: Per-message pricing for SMS to the country
- **Voice Pricing**: Per-minute pricing for voice calls
- **Outbound Call Rates**: Rates for different phone number types
- **Country Information**: Country name and ISO code
- **Currency**: Currency for pricing information

## Service Coverage

### SMS Coverage
- **Global**: 190+ countries and territories
- **Popular Markets**: US, Canada, UK, India, Australia, Germany, France, etc.
- **Pricing**: Varies by destination country and message volume

### Voice Coverage
- **Global**: 190+ countries for outbound calls
- **DID Numbers**: Available in 60+ countries
- **Quality**: Premium voice routes with high connectivity

### Geographic Restrictions
- **MMS**: Limited to US and Canada only
- **Toll-Free**: Available in multiple countries with local regulations
- **Short Codes**: Country-specific availability and approval processes

## UseCases

- **SMS Marketing Campaigns**: Send promotional SMS messages and marketing campaigns to customers
- **Transactional SMS**: Send order confirmations, shipping notifications, and account alerts via SMS
- **Two-Factor Authentication**: Implement SMS-based two-factor authentication for security
- **Customer Notifications**: Send appointment reminders, service updates, and important notifications
- **Emergency Alerts**: Send urgent notifications and emergency alerts to users and staff
- **Customer Support**: Provide SMS-based customer support and service communications
- **Event Notifications**: Send event reminders and updates to attendees and participants
- **Verification Codes**: Send verification codes for account registration and password resets
- **Survey and Feedback**: Collect customer feedback and survey responses via SMS
- **International Messaging**: Send SMS messages to international customers and global audiences
- **Voice Notifications**: Make automated voice calls for critical alerts and announcements
- **Call Center Integration**: Integrate Plivo voice capabilities with call center operations
- **IVR Systems**: Build interactive voice response systems for customer service
- **Conference Calls**: Host conference calls and group voice communications
- **Call Recording**: Record calls for quality assurance and compliance purposes
- **Click-to-Call**: Enable website visitors to initiate calls with customer service
- **Number Masking**: Protect user privacy by masking phone numbers in communications
- **Bulk Messaging**: Send large volumes of SMS messages for campaigns and notifications
- **API Integration**: Integrate SMS and voice capabilities into existing applications
- **Multi-Channel Communication**: Combine SMS, MMS, and voice for comprehensive communication strategies

## Best Practices

### SMS Best Practices
- **Opt-in Compliance**: Ensure proper opt-in consent for all recipients
- **Message Timing**: Send messages during appropriate hours for the recipient's timezone
- **Character Limits**: Keep messages under 160 characters to avoid multiple SMS charges
- **Clear CTAs**: Include clear call-to-action instructions in messages

### Voice Best Practices
- **Answer URL Preparation**: Ensure Answer URL returns valid Plivo XML
- **Error Handling**: Implement fallback URLs for reliability
- **Call Duration**: Set appropriate time limits to control costs
- **Caller ID**: Use recognizable caller IDs to improve answer rates

### Security Considerations
- **Webhook Validation**: Validate webhook signatures for security
- **Rate Limiting**: Implement rate limiting to prevent abuse
- **Number Verification**: Verify phone numbers before sending messages
- **Sensitive Data**: Avoid sending sensitive information via SMS

