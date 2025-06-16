# Linear Node

## Overview

Consume Linear API

## Credentials

- Name: linearApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: comment

#### Operation: addComment

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue ID | string |  | Yes |  |
| Comment | string |  | Yes |  |
| Additional Fields | collection | ID of the parent comment if this is a reply | No |  |

### Resource: issue

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Team Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Title | string |  | Yes |  |
| Additional Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Issue ID | string |  | Yes |  |
| Update Fields | collection | Choose from the list, or specify an ID using an <a href= | No |  |

#### Operation: addLink

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Link | string |  | Yes |  |

## UseCases

- [Customer Support Ticketing System] : Automatically create support tickets from Slack messages with ticket emojis, using AI to generate structured titles, summaries, and priority levels while preventing duplicate ticket creation
- [Bug Classification and Routing] : Classify new bugs using GPT-4 analysis and automatically move them to appropriate team workflows based on bug type, severity, and component affected
- [Visual Regression Testing Integration] : Create issues automatically when AI-powered visual regression testing detects changes in website screenshots, including detailed change descriptions and affected components
- [Automated Issue Management] : Query existing issues to check for duplicates, update issue status based on workflow progress, and add contextual comments with analysis results

