# Instagram Node

## Overview

The Instagram node allows you to interact with Instagram's API to manage content, upload media, and automate social media workflows. This node enables you to create posts, upload images and videos, and manage your Instagram presence through n8n workflows.

## Credentials

This node requires Instagram API credentials:
- **Access Token**: Instagram Basic Display API or Instagram Graph API access token
- **App ID**: Your Instagram app ID
- **App Secret**: Your Instagram app secret

To obtain API credentials:
1. Create a Facebook Developer account
2. Create a new app and add Instagram Basic Display or Instagram Graph API
3. Generate access tokens with appropriate permissions
4. Configure webhook URLs if needed

## Inputs

- **Main**: JSON data for creating posts or media uploads

## Outputs

- **Main**: Returns JSON data from Instagram API responses

## Properties

### Resource: Media

#### Operation: Create Post
| Field Name | Type | Description | Required |
|---|---|---|---|
| Media Type | options | Type of media (image, video, carousel) | Yes |
| Media URL | string | URL of the media file to upload | Yes |
| Caption | string | Caption text for the post | No |
| Location | string | Location tag for the post | No |
| Tags | array | User tags for the post | No |

#### Operation: Upload Media
| Field Name | Type | Description | Required |
|---|---|---|---|
| Media File | binary | Media file to upload | Yes |
| Media Type | options | Type of media (image, video) | Yes |
| Caption | string | Caption for the media | No |

#### Operation: Get Media
| Field Name | Type | Description | Required |
|---|---|---|---|
| Media ID | string | ID of the media to retrieve | Yes |
| Fields | array | Fields to include in response | No |

### Resource: User

#### Operation: Get Profile
| Field Name | Type | Description | Required |
|---|---|---|---|
| User ID | string | Instagram user ID | No |
| Fields | array | Profile fields to retrieve | No |

#### Operation: Get Media List
| Field Name | Type | Description | Required |
|---|---|---|---|
| User ID | string | Instagram user ID | No |
| Limit | number | Number of media items to retrieve | No |

## UseCases

- [Automated Content Publishing] : Schedule and automatically publish AI-generated content with images to Instagram based on trending topics and social media calendars
- [Social Media Campaign Management] : Coordinate multi-platform social media campaigns by automatically posting content to Instagram along with other social networks
- [Content Repurposing] : Automatically convert blog posts, articles, or other content formats into Instagram-optimized posts with appropriate images and captions
- [Trend-Based Content Creation] : Monitor trending topics and automatically generate relevant Instagram content using AI-powered content creation and image generation
- [E-commerce Product Promotion] : Automatically create Instagram posts for new products, sales, or promotions by pulling data from e-commerce platforms and generating marketing content 