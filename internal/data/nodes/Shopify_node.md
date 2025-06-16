# Shopify Node

## Overview

Comprehensive e-commerce integration with Shopify platform for managing online stores, products, orders, and customers. Shopify is a leading cloud-based e-commerce platform that enables businesses to create, customize, and manage online stores with built-in payment processing, inventory management, and shipping capabilities.

## Credentials

This node supports multiple authentication methods:
- **Name**: shopifyOAuth2Api (OAuth2)
- **Name**: shopifyAccessTokenApi (Access Token)  
- **Name**: shopifyApi (API Key)
- **Required**: Yes

### Authentication Methods

#### OAuth2 Authentication (shopifyOAuth2Api)
- **Method**: OAuth 2.0 authorization flow for public apps
- **Client ID**: Your Shopify app client ID
- **Client Secret**: Your Shopify app client secret
- **Scope**: Required permissions (read_products, write_orders, etc.)
- **Authorization URL**: Your Shopify store authorization endpoint
- **Access Token URL**: Your Shopify store token endpoint
- **Use Case**: Public apps requiring user consent and store approval

#### Access Token Authentication (shopifyAccessTokenApi)
- **Method**: Private app access token authentication
- **Shop Domain**: Your Shopify store domain (e.g., your-store.myshopify.com)
- **Access Token**: Private app access token from Shopify admin
- **Use Case**: Private apps with direct API access

#### API Key Authentication (shopifyApi) - Legacy
- **Method**: Basic API key and password authentication (legacy method)
- **Shop Domain**: Your Shopify store domain (e.g., your-store.myshopify.com)
- **API Key**: Your Shopify API key
- **Password**: Your Shopify API password
- **Use Case**: Legacy integrations (not recommended for new implementations)

### API Version
- **Current Version**: 2024-07
- **Supported Versions**: 2023-10, 2024-01, 2024-04, 2024-07

## Inputs

- **Main**: JSON data for creating, updating, or querying Shopify resources

## Outputs

- **Main**: Returns JSON responses with Shopify resource data, operation results, and metadata

## Properties

### Resource: order

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer | Collection | Customer information for the order | No | - |
| Email | String | Customer email address | No | - |
| Financial Status | Options | Payment status of the order | No | pending |
| Fulfillment Status | Options | Fulfillment status of the order | No | unfulfilled |
| Send Receipt | Boolean | Send order confirmation email | No | false |
| Send Fulfillment Receipt | Boolean | Send fulfillment confirmation email | No | false |
| Note | String | Additional order notes | No | - |
| Tags | String | Comma-separated order tags | No | - |
| Tax Lines | Collection | Tax information for the order | No | - |
| Discount Codes | Collection | Discount codes applied to the order | No | - |
| Shipping Address | Collection | Order shipping address | No | - |
| Billing Address | Collection | Order billing address | No | - |
| Additional Fields | Collection | Additional order information | No | - |
| Line Items | FixedCollection | Products and variants included in the order | No | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Order ID | String | Unique identifier of the order to delete | Yes | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Order ID | String | Unique identifier of the order to retrieve | Yes | - |
| Options | Collection | Additional options for data retrieval | No | - |

**Options fields:**
- **Fields**: Comma-separated list of fields to return (e.g., id,name,email)
- **Financial Status**: Filter by financial status
- **Fulfillment Status**: Filter by fulfillment status

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 50 |
| Options | Collection | Query filters and conditions | No | - |

**Options fields:**
- **Attribution App ID**: Show orders attributed to a specific app
- **Created At Min**: Show orders created after this date
- **Created At Max**: Show orders created before this date
- **Financial Status**: Filter by payment status
- **Fulfillment Status**: Filter by fulfillment status
- **IDs**: Comma-separated list of order IDs
- **Processed At Min**: Show orders processed after this date
- **Processed At Max**: Show orders processed before this date
- **Since ID**: Show orders after the specified ID
- **Status**: Order status filter (open, closed, cancelled, any)
- **Updated At Min**: Show orders updated after this date
- **Updated At Max**: Show orders updated before this date

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Order ID | String | Unique identifier of the order to update | Yes | - |
| Update Fields | Collection | Fields to update on the order | No | - |

**Update Fields:**
- **Email**: Customer email address
- **Note**: Order notes
- **Tags**: Comma-separated order tags
- **Buyer Accepts Marketing**: Whether customer accepts marketing
- **Metafields**: Order metafields for custom data

### Resource: product

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Title | String | The name of the product | Yes | - |
| Product Type | String | Product categorization | No | - |
| Vendor | String | Product vendor/manufacturer | No | - |
| Handle | String | URL handle (auto-generated if not provided) | No | - |
| Status | Options | Product status (active, archived, draft) | No | draft |
| Published | Boolean | Whether the product is published | No | false |
| Template Suffix | String | Template suffix for custom templates | No | - |
| SEO Title | String | SEO title for search engines | No | - |
| SEO Description | String | SEO description for search engines | No | - |
| Tags | String | Comma-separated product tags | No | - |
| Images | Collection | Product images | No | - |
| Variants | Collection | Product variants (size, color, etc.) | No | - |
| Options | Collection | Product options (size, color, etc.) | No | - |
| Metafields | Collection | Custom metafields for additional data | No | - |
| Additional Fields | Collection | Additional product information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | String | Unique identifier of the product to update | Yes | - |
| Update Fields | Collection | Fields to update on the product | No | - |

**Update Fields include all fields from create operation**

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | String | Unique identifier of the product to delete | Yes | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | String | Unique identifier of the product to retrieve | Yes | - |
| Additional Fields | Collection | Additional options for data retrieval | No | - |

**Additional Fields:**
- **Fields**: Comma-separated list of fields to return
- **Presentment Prices**: Include presentment prices for markets

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 50 |
| Additional Fields | Collection | Query filters and conditions | No | - |

**Additional Fields:**
- **Collection ID**: Filter by product collection ID
- **Created At Min**: Show products created after this date
- **Created At Max**: Show products created before this date
- **Handle**: Filter by product handle
- **IDs**: Comma-separated list of product IDs
- **Product Type**: Filter by product type
- **Published At Min**: Show products published after this date
- **Published At Max**: Show products published before this date
- **Published Status**: Filter by published status
- **Since ID**: Show products after the specified ID
- **Status**: Product status filter
- **Title**: Filter by product title
- **Updated At Min**: Show products updated after this date
- **Updated At Max**: Show products updated before this date
- **Vendor**: Filter by vendor

### Resource: customer

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | String | Customer email address | Yes | - |
| First Name | String | Customer first name | No | - |
| Last Name | String | Customer last name | No | - |
| Phone | String | Customer phone number | No | - |
| Tags | String | Comma-separated customer tags | No | - |
| Note | String | Notes about the customer | No | - |
| Verified Email | Boolean | Whether email is verified | No | false |
| Password | String | Customer password | No | - |
| Password Confirmation | String | Password confirmation | No | - |
| Send Email Invite | Boolean | Send account invitation email | No | false |
| Addresses | Collection | Customer addresses | No | - |
| Additional Fields | Collection | Additional customer information | No | - |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | String | Unique identifier of the customer to update | Yes | - |
| Update Fields | Collection | Fields to update on the customer | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | String | Unique identifier of the customer to retrieve | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | String | Unique identifier of the customer to delete | Yes | - |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 50 |
| Options | Collection | Query filters and conditions | No | - |

## UseCases

- **E-commerce Management**: Manage online store operations including products, orders, and customers
- **Inventory Management**: Track product inventory levels and automate stock management
- **Order Processing**: Automate order fulfillment and shipping processes
- **Customer Management**: Manage customer data and relationship information
- **Product Catalog**: Maintain product catalogs with descriptions, images, and pricing
- **Sales Analytics**: Analyze sales performance and generate business insights
- **Marketing Automation**: Automate marketing campaigns and customer engagement
- **Payment Processing**: Process payments and handle transaction management
- **Shipping Integration**: Integrate with shipping carriers for automated logistics
- **Multi-Channel Selling**: Manage sales across multiple channels and marketplaces
- **Subscription Commerce**: Handle subscription-based products and recurring billing
- **Dropshipping**: Manage dropshipping operations and supplier relationships
- **POS Integration**: Integrate online store with point-of-sale systems
- **App Ecosystem**: Integrate with Shopify apps and third-party services
- **International Commerce**: Handle international sales with currency and tax management
- **Inventory Synchronization**: Sync inventory between Shopify and external systems
- **Customer Segmentation**: Segment customers for targeted marketing campaigns
- **Abandoned Cart Recovery**: Implement automated abandoned cart recovery workflows
- **Product Recommendations**: Generate personalized product recommendations
- **Loyalty Programs**: Integrate with loyalty program systems and rewards
- **Financial Reporting**: Generate detailed financial reports and tax documentation
- **Warehouse Management**: Integrate with warehouse management systems
- **B2B Commerce**: Handle wholesale and B2B customer relationships
- **Mobile Commerce**: Optimize for mobile shopping experiences
- **SEO Optimization**: Manage product SEO and search engine visibility

