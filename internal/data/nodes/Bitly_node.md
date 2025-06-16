# Bitly Node

## Overview

The Bitly node allows you to consume the Bitly API for URL shortening and link management. Bitly is a popular URL shortening service that provides analytics, custom domains, and link management features. This node enables you to create, update, and manage short links within your n8n workflows for marketing campaigns, social media, and analytics tracking.

## Credentials

This node supports two authentication methods:

### Access Token
- **Access Token**: Your Bitly access token

To obtain an access token:
1. Log into your Bitly account
2. Navigate to Settings → Developer Settings
3. Generate a generic access token
4. Copy the token for use in n8n

### OAuth2
- Uses OAuth2 flow for authentication
- More secure for production environments
- Requires app registration in Bitly

## Inputs

- **Main**: JSON data for creating/updating links

## Outputs

- **Main**: Returns JSON data from Bitly API responses

## Properties

### Resource: Link

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Long URL | String | URL to be shortened | Yes |
| Additional Fields | Collection | Optional link properties | No |
| Deep Links | Collection | Mobile app deep link configuration | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Link ID | String | ID of the link to update | Yes |
| Update Fields | Collection | Link properties to update | No |
| Deep Links | Collection | Mobile app deep link configuration | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Link ID | String | ID of the link to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all links | No |
| Limit | Number | Maximum number of links to return | No |
| Additional Fields | Collection | Filtering and search options | No |

## Link Properties

### Basic Properties
| Field Name | Type | Description |
|---|---|---|
| Title | String | Custom title for the link |
| Domain | String | Custom domain for shortening (if available) |
| Group | Options | Group to organize the link under |
| Tags | Array | Tags for categorizing the link |

### Deep Link Configuration
| Field Name | Type | Description |
|---|---|---|
| App URI Path | String | Path within the mobile app |
| Install Type | Options | How to handle app installation |
| Install URL | String | URL for app installation |
| App ID | String | Mobile app identifier |

### Update Properties
| Field Name | Type | Description |
|---|---|---|
| Long URL | String | New destination URL |
| Title | String | New custom title |
| Archived | Boolean | Whether to archive the link |
| Group | Options | New group assignment |
| Tags | Array | New tag assignments |

## UseCases

- **Social Media Marketing** : Shorten links for Twitter, LinkedIn, and other social media posts
- **Email Campaign Tracking** : Track email marketing performance with shortened, trackable links
- **QR Code Generation** : Create short URLs for QR code campaigns and print materials
- **Link Analytics** : Collect detailed click analytics and user engagement data
- **A/B Testing** : Create multiple short links to test different marketing approaches
- **Brand Management** : Use custom domains for consistent brand presence in links
- **Event Promotion** : Create trackable links for event registration and promotion
- **Content Distribution** : Manage and track content sharing across multiple channels
- **Mobile App Marketing** : Use deep links to drive app installations and engagement
- **Affiliate Marketing** : Track affiliate link performance and conversions

