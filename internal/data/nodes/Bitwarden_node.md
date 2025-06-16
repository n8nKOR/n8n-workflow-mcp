# Bitwarden Node

## Overview

The Bitwarden node allows you to consume the Bitwarden API for password and organization management. Bitwarden is a comprehensive password management solution that also offers enterprise features for teams and organizations. This node enables you to manage collections, groups, members, and events within your Bitwarden organization through n8n workflows.

## Credentials

This node requires Bitwarden Organization API credentials:
- **Client ID**: Your Bitwarden organization client ID
- **Client Secret**: Your Bitwarden organization client secret
- **Environment**: Bitwarden server environment (default: bitwarden.com)

To obtain API credentials:
1. Log into your Bitwarden business account
2. Navigate to Settings → Organizations → API Keys
3. Generate organization API keys
4. Copy the client ID and secret for use in n8n

## Inputs

- **Main**: JSON data for creating/updating organization resources

## Outputs

- **Main**: Returns JSON data from Bitwarden API responses

## Properties

### Resource: Collection

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | String | ID of the collection to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | String | ID of the collection to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all collections | No |
| Limit | Number | Maximum number of collections to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | String | ID of the collection to update | Yes |
| Update Fields | Collection | Collection properties to update | No |

### Resource: Group

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Name of the group | Yes |
| Additional Fields | Collection | Optional group properties | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Group ID | String | ID of the group to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Group ID | String | ID of the group to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all groups | No |
| Limit | Number | Maximum number of groups to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Group ID | String | ID of the group to update | Yes |
| Update Fields | Collection | Group properties to update | No |

### Resource: Member

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | String | Email address of the member | Yes |
| Type | Options | Member type (User, Admin, Owner, etc.) | Yes |
| Additional Fields | Collection | Optional member properties | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Member ID | String | ID of the member to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Member ID | String | ID of the member to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all members | No |
| Limit | Number | Maximum number of members to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Member ID | String | ID of the member to update | Yes |
| Update Fields | Collection | Member properties to update | No |

### Resource: Event

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all events | No |
| Limit | Number | Maximum number of events to return | No |
| Additional Fields | Collection | Event filtering options | No |

## Organization Features

### Collections
Collections organize and control access to vault items:
- **Name**: Collection display name
- **Groups**: Groups with access to the collection
- **External ID**: External system identifier
- **Read-Only Access**: Permission level control

### Groups
Groups manage user permissions and access:
- **Name**: Group display name
- **Access Control**: Permissions for collections
- **Member Assignment**: Users within the group
- **External ID**: External system identifier

### Members
Members are users within the organization:
- **Email**: User's email address
- **Type**: User role (User, Admin, Owner, Manager)
- **Status**: Invitation and activation status
- **Access**: Group and collection assignments
- **Two-Factor**: 2FA requirement settings

### Events
Events provide audit trail functionality:
- **User Actions**: Login, logout, vault access
- **Administrative Actions**: Member management, settings changes
- **Security Events**: Failed logins, policy violations
- **Timestamps**: When events occurred
- **IP Addresses**: Source of actions

## Member Types

- **Owner**: Full organizational control
- **Admin**: User and policy management
- **User**: Standard vault access
- **Manager**: Limited administrative functions
- **Custom**: Role-based permissions

## UseCases

- **Organization Management Automation** : Automate user onboarding and access management for enterprise teams
- **Security Policy Enforcement** : Automatically enforce password policies and security requirements
- **Access Control Management** : Manage user permissions and group assignments across the organization
- **Audit and Compliance** : Track and monitor organizational security events for compliance reporting
- **Collection Management** : Organize and control access to shared vault items and secrets
- **User Lifecycle Management** : Automate user provisioning, deprovisioning, and role changes
- **Security Monitoring** : Monitor organization security events and respond to threats
- **Identity Management Integration** : Integrate with existing identity management systems
- **Password Policy Automation** : Enforce and monitor organizational password policies
- **Enterprise Security Workflows** : Create automated security workflows for large organizations

### Best Practices
- Use specific permissions for automation accounts
- Implement proper error handling
- Monitor organizational events regularly
- Keep API credentials secure and rotated
- Test operations in development environment first

## UseCases

- **Password Management** : Manage passwords and secure credentials across systems
- **Secret Management** : Store and manage API keys, tokens, and sensitive information
- **Team Credential Sharing** : Share credentials securely among team members
- **Automated Password Rotation** : Automate password rotation and security updates
- **Security Compliance** : Maintain security compliance through centralized credential management
- **DevOps Secret Management** : Manage secrets and credentials in DevOps workflows
- **Application Integration** : Integrate secure credential storage with applications
- **Audit and Monitoring** : Monitor credential usage and access for security auditing
- **Emergency Access** : Provide emergency access to critical credentials
- **Multi-Factor Authentication** : Integrate with multi-factor authentication systems

