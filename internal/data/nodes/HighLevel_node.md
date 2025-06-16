# HighLevel Node

## Overview

Consume HighLevel API v2

## Credentials

- Name: highLevelOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: CRM Management

#### Operation: Resource Management

| Field Name | Type | Description | Required |
|---|---|---|---|
| Resource | Options | Resource type to work with | Yes |
| Operation | Options | Operation to perform | Yes |
| Contact ID | String | Contact identifier | Conditional |
| Opportunity ID | String | Opportunity identifier | Conditional |
| Task ID | String | Task identifier | Conditional |
| Calendar ID | String | Calendar identifier | Conditional |
| Options | Collection | Additional configuration options | No |
## UseCases

- **CRM Management**: Manage contacts, opportunities, and customer relationships
- **Sales Pipeline**: Track deals and opportunities through sales stages
- **Task Management**: Create and manage follow-up tasks and activities
- **Calendar Integration**: Schedule appointments and manage calendar events
- **Marketing Automation**: Integrate with HighLevel marketing campaigns

### Usage Notes
- Requires active HighLevel subscription and API access
- All operations are scoped to the authenticated location
- Custom fields are dynamically loaded per location
- Pipeline and stage assignments must match location configuration
- Contact tags support multiple values for categorization
- Calendar events support timezone management
- Task priorities can be customized per location
- OAuth2 provides secure, scoped access to HighLevel data
- Pagination automatically handles large datasets
- Search functionality supports partial matching for better UX

