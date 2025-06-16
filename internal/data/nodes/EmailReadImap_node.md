# Email Trigger (IMAP) Node

## Overview

Triggers the workflow when a new email is received

## Credentials

- Name: imap, Required: Yes

## Inputs

None

## Outputs

- Main

## Properties

### Resource: Email

#### Operation: Monitor
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Mailbox Name | string | Name of the mailbox to monitor for new emails | Yes | INBOX |
| Action | options | What to do after the email has been received | Yes | read |
| Download Attachments | boolean | Whether attachments of emails should be downloaded (simple format only) | No | false |
| Format | options | The format to return the message in | Yes | simple |
| Property Prefix Name | string | Prefix for name of the binary property to write attachments | No | attachment_ |
| Custom Email Rules | string | Custom email fetching rules using IMAP search syntax | No | ["UNSEEN"] |
| Force Reconnect Every Minutes | number | Sets an interval (in minutes) to force a reconnection | No | 60 |

**Action Options:**
- **Mark as Read**: Marks the email as read after processing
- **Nothing**: Leaves the email as is (may be processed multiple times)

**Format Options:**
- **RAW**: Returns the full email message data with body content as base64url encoded string
- **Resolved**: Returns the full email with all data resolved and attachments saved as binary data  
- **Simple**: Returns the full email (do not use for inline attachments)

## UseCases

- **Email Processing**: Automatically process incoming emails
- **Document Extraction**: Extract attachments from emails
- **Email Monitoring**: Monitor specific mailboxes for new messages
- **Support Ticket Creation**: Create tickets from support emails
- **Email Archiving**: Archive emails based on rules

