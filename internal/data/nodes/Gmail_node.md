# Gmail Node

## Overview

The Gmail node allows you to interact with Gmail API to send emails, manage drafts, read messages, and automate email workflows. This node enables you to integrate Gmail functionality into your n8n workflows for email automation and management.

## Credentials

This node requires Gmail API credentials:
- **OAuth2**: Google OAuth2 credentials with Gmail API access
- **Service Account**: Google Service Account for server-to-server authentication

To obtain API credentials:
1. Create a project in Google Cloud Console
2. Enable Gmail API for your project
3. Create OAuth2 credentials or Service Account
4. Configure authorized redirect URIs for OAuth2
5. Download credentials file or copy client ID/secret

## Inputs

- **Main**: JSON data for email operations

## Outputs

- **Main**: Returns JSON data from Gmail API responses

## Properties

### Resource: Message

#### Operation: Send
| Field Name | Type | Description | Required |
|---|---|---|---|
| To | string | Recipient email address | Yes |
| Subject | string | Email subject line | Yes |
| Message | string | Email body content | Yes |
| Options | collection | Additional email options | No |

#### Operation: Send Draft
| Field Name | Type | Description | Required |
|---|---|---|---|
| Draft ID | string | ID of the draft to send | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | string | ID of the message to retrieve | Yes |
| Format | options | Format of the message response | No |
| Options | collection | Additional retrieval options | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | boolean | Whether to return all messages | No |
| Limit | number | Maximum number of messages to return | No |
| Filters | collection | Search filters and criteria | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | string | ID of the message to delete | Yes |

#### Operation: Reply
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | string | ID of the message to reply to | Yes |
| Message | string | Reply message content | Yes |
| Options | collection | Additional reply options | No |

### Resource: Draft

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| To | string | Recipient email address | Yes |
| Subject | string | Email subject line | Yes |
| Message | string | Email body content | Yes |
| Options | collection | Additional draft options | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Draft ID | string | ID of the draft to retrieve | Yes |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Draft ID | string | ID of the draft to delete | Yes |

### Resource: Label

#### Operation: Add
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | string | ID of the message | Yes |
| Label IDs | array | Array of label IDs to add | Yes |

#### Operation: Remove
| Field Name | Type | Description | Required |
|---|---|---|---|
| Message ID | string | ID of the message | Yes |
| Label IDs | array | Array of label IDs to remove | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | boolean | Whether to return all labels | No |
| Limit | number | Maximum number of labels to return | No |

## UseCases

- [AI-Powered Auto-Response System] : Automatically analyze incoming emails to determine if responses are needed and generate AI-powered reply drafts for review and sending
- [Smart PDF Attachment Processing] : Monitor Gmail for PDF attachments, extract content using AI analysis, and automatically organize files based on document type (invoices, resumes, contracts)
- [Multi-User Email-to-Task Conversion] : Convert incoming emails into structured tasks in project management systems like Notion, with intelligent routing based on email addresses and content analysis
- [Automated Email Notifications] : Send confirmation emails, status updates, and alerts as part of automated workflows for job applications, form submissions, and system events
- [Email-Based Content Triggers] : Use incoming emails as triggers for content creation workflows, such as generating blog posts or social media content based on email requests
- [Customer Support Automation] : Automatically categorize support emails, create tickets, and send acknowledgment responses with estimated resolution times
- [Email Marketing Integration] : Sync email interactions with CRM systems, track engagement metrics, and trigger follow-up sequences based on email behavior
- [Document Workflow Automation] : Process email attachments through approval workflows, extract data for database entry, and route documents to appropriate departments
- [Marketing Analytics Reporting] : Send automated Google Analytics reports with AI-generated insights and performance summaries to stakeholders on scheduled intervals
- [Sales Meeting Preparation] : Retrieve recent email correspondence with meeting attendees to provide context and conversation history for automated meeting preparation workflows
- [Email Correspondence Analysis] : Extract and analyze email content to generate meeting briefs, identify key discussion points, and prepare personalized sales materials 