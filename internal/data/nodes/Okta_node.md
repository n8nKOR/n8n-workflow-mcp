# Okta Node

## Overview

Integrate with Okta's identity and access management platform for user management and authentication operations. Okta is a leading identity provider that offers secure authentication, authorization, and user lifecycle management services. This node enables automation of user provisioning and identity operations within n8n workflows.

⚠️ **Resource Limitation**: This node implementation supports only User resource operations. Group and Application resources mentioned in other documentation do not exist in the actual implementation. For comprehensive Okta management including groups and applications, consider using Okta's REST API directly through HTTP Request nodes.

## Credentials

This node requires Okta API credentials:
- **Credential Name**: `oktaApi`
- **Required Fields**: 
  - URL: Your Okta domain URL (e.g., https://your-domain.okta.com)
  - API Token: Your Okta API token

To obtain API credentials:
1. Log into your Okta Admin Console
2. Navigate to Security > API > Tokens
3. Create a new API token
4. Copy the token value for use in n8n
5. Use your full Okta domain URL (including https://)

## Inputs

- **Main**: User data for creation, updates, or query parameters

## Outputs

- **Main**: Returns JSON responses containing user information, operation results, and metadata

## Properties

### Resource: User

#### Operation: Create
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | String | User email address | Yes | - |
| First Name | String | User first name | Yes | - |
| Last Name | String | User last name | Yes | - |
| Login | String | User login name (if different from email) | No | - |
| Password | String | User password for credential-based authentication | No | - |
| Activate | Boolean | Activate user account after creation | No | true |
| Additional Fields | Collection | Additional user profile fields | No | - |

**Additional Fields Options:**
- **Mobile Phone**: User's mobile phone number
- **Department**: User's department
- **Title**: Job title
- **Manager**: Manager's user ID
- **Employee Number**: Employee identification number
- **Cost Center**: Cost center assignment
- **Organization**: Organization name
- **Division**: Division within organization

#### Operation: Delete
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | String | ID of the user to delete | Yes | - |
| Send Admin Email | Boolean | Send email notification to admins | No | false |

#### Operation: Get
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | String | ID of the user to retrieve | Yes | - |

#### Operation: Get All
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Return all users | No | false |
| Limit | Number | Maximum number of users to return | No | 100 |
| Filter | String | Filter expression for users (e.g., status eq "ACTIVE") | No | - |
| Search | String | Search query for users | No | - |

**Common Filter Examples:**
- `status eq "ACTIVE"` - Active users only
- `lastLogin gt "2023-01-01T00:00:00.000Z"` - Users who logged in after date
- `profile.department eq "Engineering"` - Users in specific department

#### Operation: Update
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | String | ID of the user to update | Yes | - |
| Update Fields | Collection | Fields to update | No | - |

**Update Fields Options:**
- **First Name**: Updated first name
- **Last Name**: Updated last name
- **Email**: Updated email address
- **Mobile Phone**: Updated mobile phone
- **Department**: Updated department
- **Title**: Updated job title
- **Status**: User status (ACTIVE, SUSPENDED, DEPROVISIONED)
- **Password**: New password (if applicable)

#### Operation: Activate
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | String | ID of the user to activate | Yes | - |
| Send Email | Boolean | Send activation email to user | No | true |

#### Operation: Deactivate
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | String | ID of the user to deactivate | Yes | - |
| Send Admin Email | Boolean | Send notification to administrators | No | false |

#### Operation: Suspend
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | String | ID of the user to suspend | Yes | - |

#### Operation: Unsuspend
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | String | ID of the user to unsuspend | Yes | - |

## UseCases

- **Identity Management**: Manage user identities and access across your organization with centralized user lifecycle management
- **User Provisioning**: Automate user account creation, updates, and deactivation for efficient onboarding and offboarding
- **Employee Onboarding**: Streamline employee onboarding with automated account setup and profile configuration
- **Employee Offboarding**: Automate user deactivation and access removal during employee departures
- **Compliance Reporting**: Generate reports on user access, status changes, and account activities for compliance requirements
- **Security Automation**: Automate security policies, user status monitoring, and access reviews
- **Directory Synchronization**: Sync user data between Okta and other systems like HR platforms and business applications
- **Audit and Monitoring**: Track user activities, status changes, and access patterns for security monitoring
- **Bulk User Operations**: Perform bulk user creation, updates, or status changes for large-scale user management
- **Integration Workflows**: Connect Okta user management with business processes and external systems
- **User Status Management**: Automate user activation, deactivation, suspension based on business rules
- **Profile Data Management**: Keep user profiles synchronized across multiple systems and applications
- **Access Control Automation**: Implement automated access control based on user attributes and organizational changes
- **Identity Governance**: Support identity governance initiatives with automated user lifecycle management
- **Self-Service Automation**: Enable automated responses to self-service requests and profile updates
