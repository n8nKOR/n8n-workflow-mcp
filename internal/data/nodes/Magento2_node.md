# Magento 2 Node

## Overview

The Magento 2 node provides comprehensive integration with Magento's REST API, enabling automation of e-commerce operations including customer management, product catalog management, order processing, and invoice generation. This node supports all major Magento resources with advanced filtering, custom attributes, and bulk operations for enterprise-grade e-commerce automation.

## Credentials

- Name: magento2Api, Required: Yes

### Credential Configuration
To configure Magento 2 API credentials:
1. Log into your Magento 2 admin panel
2. Navigate to System > Extensions > Integrations
3. Create a new integration with appropriate API permissions
4. Copy the Access Token for use in n8n credentials
5. Ensure API user has permissions for required resources (customers, products, orders, invoices)

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Customer

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | String | Customer email address | Yes |
| First Name | String | Customer first name | Yes |
| Last Name | String | Customer last name | Yes |
| Website ID | Options | Website ID for the customer | No |
| Store ID | Options | Store ID for the customer | No |
| Group ID | Options | Customer group ID | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Customer ID | Number | ID of the customer to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Customer ID | Number | ID of the customer to retrieve | Yes |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or use pagination | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Customer ID | Number | ID of the customer to update | Yes |
| Update Fields | Collection | Fields to update | No |

### Resource: Invoice

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Order ID | Number | ID of the order to create invoice for | Yes |
| Items | Collection | Invoice items | No |
| Notify | Boolean | Whether to notify customer | No |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Invoice ID | Number | ID of the invoice to retrieve | Yes |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or use pagination | No |
| Limit | Number | Maximum number of results to return | No |

### Resource: Order

#### Operation: Cancel

| Field Name | Type | Description | Required |
|---|---|---|---|
| Order ID | Number | ID of the order to cancel | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Order ID | Number | ID of the order to retrieve | Yes |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or use pagination | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Ship

| Field Name | Type | Description | Required |
|---|---|---|---|
| Order ID | Number | ID of the order to ship | Yes |
| Items | Collection | Items to ship | No |
| Notify | Boolean | Whether to notify customer | No |

### Resource: Product

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| SKU | String | Product SKU | Yes |
| Name | String | Product name | Yes |
| Price | Number | Product price | Yes |
| Type ID | Options | Product type | Yes |
| Attribute Set ID | Options | Attribute set ID | Yes |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Product SKU | String | SKU of the product to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Product SKU | String | SKU of the product to retrieve | Yes |

#### Operation: Get All

| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all results or use pagination | No |
| Limit | Number | Maximum number of results to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Product SKU | String | SKU of the product to update | Yes |
| Update Fields | Collection | Fields to update | No |

## UseCases

- **Customer Management**: Automate customer account creation, updates, and maintenance workflows
- **Order Processing**: Stream order data to external systems, fulfillment services, and accounting platforms
- **Inventory Synchronization**: Sync product information between Magento and warehouse management systems
- **Customer Service Integration**: Connect customer data with support ticket systems and CRM platforms
- **Marketing Automation**: Trigger email campaigns based on customer behavior and order history
- **Reporting and Analytics**: Extract sales data for business intelligence and reporting systems
- **Multi-Channel Integration**: Synchronize inventory and orders across multiple sales channels
- **B2B Integration**: Connect with enterprise systems for customer account management and bulk ordering
- **Payment Processing**: Integrate with payment gateways and financial systems for invoice management
- **Shipping Integration**: Automate shipping label creation and tracking information updates
- **Compliance Reporting**: Generate compliance reports for tax authorities and regulatory bodies
- **Data Migration**: Migrate customer and product data between Magento instances or other platforms
- **Quality Assurance**: Automated testing of product data integrity and customer account validation
- **Pricing Management**: Sync pricing information with external pricing engines and ERP systems
- **Returns Processing**: Automate return merchandise authorization (RMA) workflows
- **Subscription Management**: Handle subscription-based products and recurring billing integration
- **Affiliate Management**: Track affiliate sales and commission calculations
- **Mobile App Integration**: Provide data to mobile applications for customer and product information
- **Social Commerce**: Sync product catalogs with social media platforms for social selling
- **Drop-shipping Integration**: Automate vendor notifications and order routing for drop-shipped products

### Usage Notes
- All operations support Magento's custom attribute system for extensibility
- Dynamic option loading provides real-time store configuration data
- Filtering system supports both simple and complex search criteria
- Bulk operations are optimized for performance with large datasets
- Multi-store support enables management of complex store hierarchies
- Extension attributes provide integration with third-party Magento modules
- Authentication uses Magento's standard REST API token system
- Error handling provides detailed debugging information for failed operations
- Custom attributes support all Magento field types including multi-select and date fields