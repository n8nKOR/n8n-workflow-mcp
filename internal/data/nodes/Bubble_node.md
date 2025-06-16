# Bubble Node

## Overview

The Bubble node allows you to consume the Bubble Data API for no-code application integration. Bubble is a visual programming platform that enables users to build web applications without code. This node provides access to your Bubble app's database for creating, reading, updating, and deleting data objects.

## Credentials

This node requires Bubble API credentials:
- **App Name**: Your Bubble application name
- **API Token**: Your Bubble API token

To obtain API credentials:
1. Open your Bubble application editor
2. Navigate to Settings → API
3. Enable the Data API
4. Generate an API token
5. Note your app name from the URL

## Inputs

- **Main**: JSON data for object operations

## Outputs

- **Main**: Returns JSON data from Bubble Data API responses

## Properties

### Resource: Object

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | String | Name of the object type in Bubble app | Yes |
| Properties | Object | Object properties to set during creation | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | String | Name of the object type in Bubble app | Yes |
| Object ID | String | Unique identifier of the object to delete | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | String | Name of the object type in Bubble app | Yes |
| Object ID | String | Unique identifier of the object to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | String | Name of the object type in Bubble app | Yes |
| Return All | Boolean | Whether to return all matching objects | No |
| Limit | Number | Maximum number of objects to return | No |
| Constraints | Array | Filter constraints for the query | No |
| Sort Field | String | Field name to sort results by | No |
| Sort Order | Options | Sort order (ascending, descending) | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | String | Name of the object type in Bubble app | Yes |
| Object ID | String | Unique identifier of the object to update | Yes |
| Properties | Object | Object properties to update | No |

## API Details

### Base URL Structure
```
https://{app-name}.bubbleapps.io/api/1.1/obj/{object-type}
# For self-hosted: https://your-domain.com/api/1.1/obj/{object-type}
```

### Authentication
- Uses Bearer token authentication
- API token required for all requests
- Token configured in Bubble app settings under API tab

### Data API Features
- **CRUD Operations**: Full create, read, update, delete support
- **Filtering**: Advanced constraint-based filtering
- **Sorting**: Sort results by any field in ascending or descending order
- **Pagination**: Efficient pagination for large datasets
- **Relationships**: Support for Bubble's relational data structures

### Constraint Types
- **Equals**: Exact value matching
- **Not Equal**: Exclude specific values
- **Contains**: Text contains substring
- **Empty**: Field is empty or null
- **Not Empty**: Field has a value
- **Greater Than**: Numeric or date comparisons
- **Less Than**: Numeric or date comparisons
- **Geographic**: Location-based constraints

## UseCases

### No-Code Application Integration
- **Data Synchronization**: Sync Bubble app data with external systems and databases
- **Automated Workflows**: Trigger workflows based on Bubble app events and data changes
- **Business Process Automation**: Connect Bubble apps to enterprise systems and processes
- **Multi-App Data Sharing**: Share data between multiple Bubble applications
- **External API Integration**: Bridge Bubble apps with third-party APIs and services

### Customer and User Management
- **User Profile Management**: Manage user accounts and profiles across systems
- **Customer Data Enrichment**: Enhance customer records with external data sources
- **Account Synchronization**: Keep customer data consistent across multiple platforms
- **User Activity Tracking**: Monitor and analyze user interactions and behaviors
- **Automated User Onboarding**: Streamline user registration and setup processes

### E-commerce and Marketplace Operations
- **Product Catalog Management**: Manage product listings and inventory data
- **Order Processing**: Automate order fulfillment and tracking workflows
- **Customer Support**: Integrate support tickets and customer communications
- **Payment Processing**: Connect payment systems with order and customer data
- **Inventory Management**: Track stock levels and trigger reorder processes

### Content and Project Management
- **Content Creation Workflows**: Manage content approval and publishing processes
- **Project Tracking**: Monitor project progress and deliverables
- **Document Management**: Organize and share project documents and assets
- **Team Collaboration**: Facilitate team communication and task management
- **Resource Planning**: Manage team resources and capacity planning

### Analytics and Reporting
- **Data Export**: Extract Bubble app data for external analysis and reporting
- **Business Intelligence**: Feed Bubble data into BI tools and dashboards
- **Performance Monitoring**: Track app performance and user engagement metrics
- **Compliance Reporting**: Generate reports for regulatory compliance requirements
- **Data Backup**: Regular backup of critical business data from Bubble apps

## Error Handling

### Common Error Codes
- **400 Bad Request**: Invalid request parameters or malformed JSON
- **401 Unauthorized**: Invalid or missing API token
- **403 Forbidden**: Insufficient permissions for the operation
- **404 Not Found**: Object type or specific object not found
- **422 Unprocessable Entity**: Validation errors or constraint violations
- **429 Too Many Requests**: API rate limit exceeded
- **500 Internal Server Error**: Bubble API server issues

### Specific Error Scenarios
1. **Object Type Not Found**: Specified object type doesn't exist in Bubble app
2. **Invalid Object ID**: Object ID format is incorrect or object doesn't exist
3. **Permission Denied**: API token lacks permissions for the operation
4. **Validation Errors**: Object properties don't meet Bubble app constraints
5. **Rate Limiting**: API usage exceeds plan limits

### Troubleshooting Tips
1. **Verify API Setup**: Ensure Data API is enabled in Bubble app settings
2. **Check Object Types**: Confirm object type names match exactly (case-sensitive)
3. **Validate Permissions**: Verify API token has appropriate read/write permissions
4. **Test Constraints**: Validate filter constraints syntax and field names
5. **Monitor Rate Limits**: Track API usage to avoid exceeding plan quotas

## Advanced Configuration

### API Token Configuration
- **Token Generation**: Create API tokens in Bubble app Settings → API tab
- **Permission Levels**: Set appropriate read/write permissions for operations
- **Security**: Store tokens securely and rotate regularly
- **Environment Management**: Use different tokens for development and production

### Data Modeling Strategies
- **Object Type Naming**: Use consistent naming conventions for object types
- **Relationship Handling**: Properly handle Bubble's relational data structures
- **Field Mapping**: Map Bubble fields to external system fields accurately
- **Data Validation**: Implement proper validation before creating/updating objects

### Performance Optimization
- **Batch Operations**: Group multiple operations to reduce API calls
- **Efficient Filtering**: Use specific constraints to minimize data transfer
- **Pagination Strategy**: Implement proper pagination for large datasets
- **Caching**: Cache frequently accessed data to reduce API load

### Integration Patterns
- **Real-time Sync**: Implement real-time data synchronization patterns
- **Event-Driven Architecture**: Use Bubble data changes to trigger external workflows
- **Data Transformation**: Transform data between Bubble and external system formats
- **Error Recovery**: Implement robust error handling and retry mechanisms

### Security and Compliance
- **Access Control**: Implement proper access controls and authentication
- **Data Privacy**: Handle personal data in compliance with privacy regulations
- **Audit Logging**: Log all data operations for security and compliance
- **Rate Limiting**: Implement client-side rate limiting to avoid API abuse

#### Operations
- **Create**: Create new data objects
- **Delete**: Remove data objects
- **Get**: Retrieve specific data objects
- **Get Many**: List multiple data objects
- **Update**: Modify existing data objects

## UseCases

- **No-Code Application Integration** : Sync Bubble app data with external systems and databases
- **Customer Management** : Manage user accounts and profiles across multiple platforms
- **E-commerce Operations** : Handle product catalogs, orders, and inventory management
- **Content Management** : Organize and manage content creation and publishing workflows
- **Analytics and Reporting** : Extract Bubble data for external analysis and business intelligence
- **Business Process Automation** : Connect Bubble apps to enterprise systems and processes
- **Multi-App Data Sharing** : Share data between multiple Bubble applications
- **User Activity Tracking** : Monitor and analyze user interactions and behaviors
- **Automated Workflows** : Trigger workflows based on Bubble app events and data changes
- **External API Integration** : Bridge Bubble apps with third-party APIs and services
- **Custom Data Types**: Work with your app's custom data types
- **Field Management**: Handle all field types (text, number, date, etc.)
- **Relationships**: Manage linked data objects
- **Privacy Rules**: Respects Bubble's privacy rule configuration

## Properties

### Resource: Object

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | options | Custom data type name | Yes |
| Fields | collection | Object field values | Yes |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | options | Custom data type name | Yes |
| Object ID | string | Unique identifier of object to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | options | Custom data type name | Yes |
| Object ID | string | Unique identifier of object to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | options | Custom data type name | Yes |
| Limit | number | Maximum number of objects to return | No |
| Cursor | string | Pagination cursor for next page | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type Name | options | Custom data type name | Yes |
| Object ID | string | Unique identifier of object to update | Yes |
| Fields | collection | Updated field values | Yes |

## UseCases

- No-code App Integration: Connect Bubble applications with external systems and workflows
- Database Synchronization: Keep Bubble data in sync with other platforms and services
- Automated Data Processing: Process and transform Bubble app data through n8n workflows
- User Management: Automate user creation and updates in Bubble applications
- E-commerce Integration: Sync product catalogs and order data between Bubble and other platforms

## Notes

### Data Types
- Must define custom data types in Bubble first
- Field names must match exactly as defined in Bubble
- Supports all Bubble field types (text, number, date, list, etc.)

### API Access
- Requires Data API to be enabled in Bubble
- Privacy rules affect what data is accessible
- API tokens provide read/write access based on configuration

### Performance
- Consider pagination for large datasets
- Use specific field selection when possible
- Monitor API usage in Bubble dashboard

