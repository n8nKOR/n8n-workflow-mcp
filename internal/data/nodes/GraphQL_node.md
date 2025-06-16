# GraphQL Node

## Overview

Makes a GraphQL request and returns the received data. GraphQL is a query language and runtime for APIs that provides a complete and understandable description of the data in your API. This node enables you to execute GraphQL queries, mutations, and subscriptions against any GraphQL endpoint, making it easy to integrate with modern APIs that use GraphQL instead of REST.

## Credentials

- Name: httpBasicAuth, Required: Yes (for Basic Auth)
- Name: httpCustomAuth, Required: Yes (for Custom Auth)
- Name: httpDigestAuth, Required: Yes (for Digest Auth)
- Name: httpHeaderAuth, Required: Yes (for Header Auth)
- Name: httpQueryAuth, Required: Yes (for Query Auth)
- Name: oAuth1Api, Required: Yes (for OAuth1)
- Name: oAuth2Api, Required: Yes (for OAuth2)

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: GraphQL Request

#### Operation: Execute Query

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | Options | The way to authenticate (default: none) | No |
| HTTP Request Method | Options | The underlying HTTP request method to use (default: POST) | No |
| Endpoint | String | The GraphQL endpoint | Yes |
| Ignore SSL Issues (Insecure) | Boolean | Whether to download the response even if SSL certificate validation is not possible (default: false) | No |
| Request Format | Options | The format for the query payload (GraphQL Raw or JSON) | Yes |
| Query | String | The GraphQL query to execute | Yes |
| Variables | JSON | GraphQL variables as JSON object | No |
| Operation Name | String | Name of the operation to execute (for multiple operations) | No |
| Headers | Collection | Custom headers to send with the request | No |

## UseCases

- **API Integration**: Connect to GraphQL APIs for efficient data retrieval and manipulation
- **Content Management**: Query and mutate CMS data via GraphQL endpoints
- **E-commerce**: Manage products, orders, and customer data through GraphQL APIs
- **Social Media**: Query social media platforms with GraphQL APIs (GitHub, Facebook, etc.)
- **Real-time Data**: Execute GraphQL subscriptions for real-time data updates
- **Database Operations**: Perform complex database queries with precise field selection
- **Microservices**: Communicate with GraphQL microservices and federated schemas
- **Mobile Backend**: Optimize mobile app data fetching with GraphQL queries
- **Analytics**: Query analytics platforms that provide GraphQL interfaces
- **Headless CMS**: Integrate with headless CMS systems via GraphQL



