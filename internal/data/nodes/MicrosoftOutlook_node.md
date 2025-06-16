# Microsoft Outlook Node

## Overview

Consume Microsoft Outlook API to manage emails, calendars, contacts, events, and attachments through Microsoft Graph API integration with comprehensive support for shared mailboxes and advanced features.

## Credentials

- Name: microsoftOutlookOAuth2Api, Required: Yes

### Required Scopes
- openid
- offline_access
- Contacts.Read
- Contacts.ReadWrite
- Calendars.Read
- Calendars.Read.Shared
- Calendars.ReadWrite
- Mail.ReadWrite
- Mail.ReadWrite.Shared
- Mail.Send
- Mail.Send.Shared
- MailboxSettings.Read

### Authentication Setup
1. Register application in Azure AD
2. Configure OAuth2 credentials with appropriate scopes
3. Supports shared mailbox access via User Principal Name configuration

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: message

#### Operation: create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| To | collection | Recipients email addresses | Yes |  |
| Subject | string | Email subject line | Yes |  |
| Body Content Type | options | HTML or Text format | Yes | HTML |
| Body | string | Email message body | No |  |
| Additional Fields | collection | CC, BCC, attachments, importance, etc. | No |  |

#### Operation: delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message to delete | Yes |  |

#### Operation: get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message to retrieve | Yes |  |
| Additional Fields | collection | Fields to include in response | No |  |

#### Operation: getAll
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Search and filter criteria | No |  |

#### Operation: move
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message to move | Yes |  |
| Folder | string | Destination folder ID | Yes |  |

#### Operation: reply
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message to reply to | Yes |  |
| Reply Type | options | Reply or Reply All | Yes |  |
| Body Content Type | options | HTML or Text format | Yes | HTML |
| Body | string | Reply message body | Yes |  |
| Additional Fields | collection | Additional reply options | No |  |

#### Operation: send
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| To | collection | Recipients email addresses | Yes |  |
| Subject | string | Email subject line | Yes |  |
| Body Content Type | options | HTML or Text format | Yes | HTML |
| Body | string | Email message body | No |  |
| Additional Fields | collection | CC, BCC, attachments, importance, etc. | No |  |

#### Operation: update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message to update | Yes |  |
| Update Fields | collection | Fields to update | Yes |  |

### Resource: calendar

#### Operation: create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Calendar name | Yes |  |
| Additional Fields | collection | Color, description, etc. | No |  |

#### Operation: delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Calendar ID | string | ID of the calendar to delete | Yes |  |

#### Operation: get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Calendar ID | string | ID of the calendar to retrieve | Yes |  |

#### Operation: getAll
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Calendar ID | string | ID of the calendar to update | Yes |  |
| Update Fields | collection | Fields to update | Yes |  |

### Resource: contact

#### Operation: create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Display Name | string | Contact display name | Yes |  |
| Additional Fields | collection | Email, phone, address, etc. | No |  |

#### Operation: delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to delete | Yes |  |

#### Operation: get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to retrieve | Yes |  |

#### Operation: getAll
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Contact ID | string | ID of the contact to update | Yes |  |
| Update Fields | collection | Fields to update | Yes |  |

### Resource: draft

#### Operation: create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subject | string | Email subject line | Yes |  |
| Body Content Type | options | HTML or Text format | Yes | HTML |
| Body | string | Email message body | No |  |
| Additional Fields | collection | To, CC, BCC, attachments, etc. | No |  |

#### Operation: delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the draft to delete | Yes |  |

#### Operation: get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the draft to retrieve | Yes |  |

#### Operation: send
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the draft to send | Yes |  |

#### Operation: update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the draft to update | Yes |  |
| Update Fields | collection | Fields to update | Yes |  |

### Resource: event

#### Operation: create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Calendar ID | string | ID of the calendar | Yes |  |
| Subject | string | Event subject | Yes |  |
| Start | dateTime | Event start time | Yes |  |
| End | dateTime | Event end time | Yes |  |
| Additional Fields | collection | Location, attendees, body, etc. | No |  |

#### Operation: delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Calendar ID | string | ID of the calendar | Yes |  |
| Event ID | string | ID of the event to delete | Yes |  |

#### Operation: get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Calendar ID | string | ID of the calendar | Yes |  |
| Event ID | string | ID of the event to retrieve | Yes |  |

#### Operation: getAll
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Calendar ID | string | ID of the calendar | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Date range, search filters | No |  |

#### Operation: update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Calendar ID | string | ID of the calendar | Yes |  |
| Event ID | string | ID of the event to update | Yes |  |
| Update Fields | collection | Fields to update | Yes |  |

### Resource: folder

#### Operation: create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Display Name | string | Folder display name | Yes |  |
| Parent Folder | string | Parent folder ID | No |  |

#### Operation: delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder ID | string | ID of the folder to delete | Yes |  |

#### Operation: get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder ID | string | ID of the folder to retrieve | Yes |  |

#### Operation: getAll
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Folder ID | string | ID of the folder to update | Yes |  |
| Update Fields | collection | Fields to update | Yes |  |

### Resource: messageAttachment

#### Operation: add
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message | Yes |  |
| Binary Property | string | Binary data property name | Yes |  |
| Additional Fields | collection | Attachment name, type, etc. | No |  |

#### Operation: download
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message | Yes |  |
| Attachment ID | string | ID of the attachment | Yes |  |
| Binary Property | string | Property name for binary data | Yes |  |

#### Operation: get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message | Yes |  |
| Attachment ID | string | ID of the attachment | Yes |  |

#### Operation: getAll
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Message ID | string | ID of the message | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **Email Marketing Automation**: Send personalized email campaigns, track responses, and manage subscriber communications through Microsoft Outlook integration
- **Customer Support Ticketing**: Automatically create support tickets from incoming emails, categorize by importance, and route to appropriate teams
- **Meeting and Event Management**: Schedule meetings, send calendar invites, manage attendee responses, and sync events with CRM systems
- **Document Workflow Automation**: Send documents via email with attachments, track delivery status, and manage approval workflows
- **Lead Generation and Follow-up**: Automatically send follow-up emails to leads, schedule meetings, and track engagement through email responses
- **Internal Communication Automation**: Distribute company announcements, policy updates, and team communications through automated email workflows
- **Contact and CRM Synchronization**: Sync contacts between Outlook and external CRM systems, maintaining up-to-date customer information
- **Calendar Integration for Project Management**: Create project milestones as calendar events, schedule team meetings, and send deadline reminders
- **Email Response Processing**: Process incoming emails for specific content, extract data, and trigger automated business processes
- **Shared Mailbox Management**: Manage shared mailboxes for team communications, customer service queues, and departmental email handling