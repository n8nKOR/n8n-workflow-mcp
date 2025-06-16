# Figma Trigger (Beta)

## Overview

The Figma Trigger node starts the workflow when Figma events occur in the design platform. This trigger node monitors specified Figma teams and automatically initiates workflows when various events are detected through webhooks, such as file updates, comments, deletions, version updates, and library publications.

## Credentials

- **Credential Name**: `figmaApi`
- **Required Fields**: 
  - Personal Access Token

## Inputs

This node is a trigger node and has no inputs.

## Outputs

- **Main Output**: Outputs the Figma event data received via webhook.

## Properties

### Resource: Figma Trigger

#### Operation: Team Events

| Field Name | Type | Description | Required |
|---|---|---|---|
| Team ID | String | Figma team ID to monitor for events | Yes |
| Events | Multi Options | Types of events to monitor | Yes |
| Options | Collection | Additional configuration options | No |

#### Options Collection

| Field Name | Type | Description | Required |
|---|---|---|---|
| Response Code | Number | HTTP response code for webhook | No |
| Response Mode | Options | How to respond to webhook requests | No |

## UseCases

- **Design Review Automation**: Automatically send notifications to reviewers when files are updated
- **Version Management**: Automatically document and share changes with the team when new versions are created
- **Comment Notifications**: Send instant notifications to relevant team members when comments are added to design files
- **Backup Automation**: Automatically save files to backup systems when they are updated
- **Project Management Integration**: Automatically sync Figma activities with project management tools
- **Design System Distribution**: Automatically notify development teams when design system libraries are published
- **Collaboration Enhancement**: Streamline design-to-development handoff processes
- **Change Tracking**: Monitor and log all design changes for audit and review purposes 