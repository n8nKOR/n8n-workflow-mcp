# Webex by Cisco

## Overview

The Webex by Cisco node provides comprehensive integration with the Cisco Webex API, enabling automated meeting management and messaging functionality. This integration allows you to streamline video conferencing workflows, automate communication processes, and enhance productivity through programmable access to Webex's collaboration features.

## Credentials

- **Credential Name**: `ciscoWebexOAuth2Api`
- **Required Fields**: Complete OAuth2 configuration with authorization code grant flow
- **Authentication Method**: OAuth2 with Cisco Webex developer application

## Inputs

- **Main Input**: Data to be processed by the Webex operation, varies by resource and operation type

## Outputs

- **Main Output**: Webex API response data containing operation results, meeting details, message information, or error details

## Properties

### Resource: Meeting

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Title | String | Meeting title/subject | Yes |
| Start Time | DateTime | Meeting start time in ISO format | Yes |
| End Time | DateTime | Meeting end time in ISO format | Yes |
| Site URL | Options | Webex site URL selection (dynamically loaded) | No |
| Timezone | String | Meeting timezone (e.g., America/New_York) | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Meeting ID | String | ID of the meeting to retrieve | Yes |

#### Operation: Get All
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all meetings or use pagination | No |
| Limit | Number | Maximum number of meetings to return | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Meeting ID | String | ID of the meeting to delete | Yes |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Meeting ID | String | ID of the meeting to update | Yes |
| Update Fields | Collection | Collection of fields to update | No |

### Resource: Message

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Destination | Options | Message destination type (Room or Person) | Yes |
| Room ID | Options | Target room selection (dynamically loaded) | Yes (when Room selected) |
| To Person ID/Email | String | Recipient person ID or email address | Yes (when Person selected) |
| Text | String | Message content | Yes |
| Markdown | Boolean | Enable markdown formatting in message | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | String | ID of the message to retrieve | Yes |

#### Operation: Get All
| Field Name | Type | Description | Required |
|---|---|---|---|
| Room ID | Options | Room to retrieve messages from | Yes |
| Return All | Boolean | Whether to return all messages or use pagination | No |
| Limit | Number | Maximum number of messages to return | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | String | ID of the message to delete | Yes |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | String | ID of the message to update | Yes |
| Text | String | New message content | Yes |

## UseCases

- **Automated Meeting Scheduling**: Automatically create and schedule Webex meetings based on specific triggers and conditions
- **Event Notification System**: Send automatic alert messages to Webex rooms when important events occur
- **Meeting Lifecycle Management**: Automate the creation, modification, and deletion of meetings for efficient schedule management
- **Team Communication Automation**: Automatically share project status updates and notifications with team rooms
- **Customer Support Integration**: Automatically schedule support meetings and send invitations based on customer inquiries
- **Reporting and Analytics**: Generate automated reports on meeting attendance rates and collaboration activities
- **Workflow Integration**: Connect Webex meetings with CRM, project management, and other business systems
- **Broadcast Messaging**: Send important announcements and updates to multiple rooms simultaneously
- **Meeting Reminder System**: Automatically send meeting reminders and preparation materials to participants
- **Resource Management**: Coordinate meeting rooms and resources with automated scheduling workflows
- **Customer Onboarding**: Automatically schedule and manage customer onboarding meetings and training sessions
- **Emergency Communication**: Quickly create emergency meetings and send urgent messages to relevant teams 