# ChatTrigger

## Overview

The ChatTrigger node provides a webhook-based chat interface that enables real-time conversational interactions within n8n workflows. This powerful trigger node supports both hosted and embedded chat modes, allowing users to create interactive AI assistants, customer support systems, and conversational applications. The node handles file uploads, session management, and multiple authentication methods while providing extensive customization options for chat appearance and behavior.

## Credentials

The ChatTrigger node supports multiple authentication methods:

### Supported Authentication Types

| Authentication Type | Description | Use Case |
|---|---|---|
| None | No authentication required | Public chat interfaces |
| HTTP Basic Auth | Username/password authentication | Simple protected access |
| n8n User Auth | n8n platform user authentication | Internal team access |

### Authentication Configuration

#### HTTP Basic Auth
- **Username**: Required for basic authentication
- **Password**: Required for basic authentication
- **Realm**: Optional authentication realm description

#### n8n User Auth
- Leverages existing n8n user accounts
- Automatic session management
- Role-based access control

## Inputs

### Optional Inputs
- **Memory**: Connection to memory system for conversation context preservation

## Outputs

The ChatTrigger node provides comprehensive output for each chat interaction:

### Chat Message Data
- **Message Content**: User's text input
- **Session ID**: Unique identifier for conversation session
- **User Information**: Authentication details (when applicable)
- **Timestamp**: Message timestamp
- **File Attachments**: Uploaded files (when file upload is enabled)

### Session Management
- **Session State**: Current conversation state
- **Memory Context**: Historical conversation data
- **User Preferences**: Stored user settings and preferences

## Properties

### Resource: Chat Trigger

#### Operation: Chat Interface

| Field Name | Type | Description | Required |
|---|---|---|---|
| Mode | Options | Chat interface mode (hosted/embedded) | Yes |
| Authentication | Options | Authentication method | Yes |
| Allow File Uploads | Boolean | Enable file upload functionality | No |
| Options | Collection | Additional configuration options | No |

## UseCases

- **Customer Support Chat** : Real-time customer service with file support and session persistence
- **AI Assistant Integration** : Embed conversational AI into websites and applications
- **Technical Support** : File upload for debugging with conversation history
- **Educational Tutoring** : Interactive learning with document sharing capabilities
- **Sales Chat** : Lead qualification with conversation tracking and context
- **Help Desk Systems** : Internal support with authentication and session management
- **Product Demos** : Interactive product demonstrations with guided conversations
- **Feedback Collection** : Structured feedback gathering with conversational interface
- **Document Analysis** : Upload and discuss documents with AI assistance
- **Code Review** : Share code files for AI-powered review and suggestions
- **Content Creation** : Collaborative content development with file sharing
- **Training and Onboarding** : Interactive training systems with session tracking
- **Survey and Research** : Dynamic questionnaires with conversation flow
- **Troubleshooting Guides** : Interactive diagnostic tools with context awareness
- **Language Learning** : Conversational practice with pronunciation file uploads
- **Creative Collaboration** : Brainstorming sessions with media file sharing