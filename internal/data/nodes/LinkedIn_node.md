# LinkedIn Node

## Overview

Consume LinkedIn API

## Credentials

This node supports multiple LinkedIn API credentials:

### Authentication Options
- **Standard LinkedIn API**: `linkedInOAuth2Api` - For personal LinkedIn operations
- **Community Management API**: `linkedInCommunityManagementOAuth2Api` - For organization management operations

### Authentication Methods
| Method | Description | Use Case |
|---|---|---|
| Standard | Personal LinkedIn API access | Individual user posts and activities |
| Community Management | Organization management API | Company page management and posting |

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: post

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Authentication | options | Choose authentication method (Standard or Community Management) | Yes | Standard |
| Post As | options | Determines the posting context (Person or Organization) | Yes | Person |
| Person Name or ID | options | Person as which the post should be posted. Choose from the list, or specify an ID using expression | Yes (when Post As = Person) |  |
| Organization URN | string | URN of Organization as which the post should be posted | Yes (when Post As = Organization) |  |
| Text | string | The primary content of the post | No |  |
| Media Category | options | Media type for the post | No | NONE |
| Input Binary Field | string | Binary field containing media data | Yes (when Media Category = IMAGE) |  |
| Additional Fields | collection | Additional post configuration options | No |  |

#### Media Category Options
| Option | Description |
|---|---|
| NONE | The post does not contain any media, and will only consist of text |
| ARTICLE | Post contains an article link with metadata |
| IMAGE | Post contains an image attachment |

#### Additional Fields Options
| Field Name | Type | Description |
|---|---|---|
| Description | string | Short description for your image or article |
| Original URL | string | Original URL for article posts |
| Title | string | Title for article posts |
| Visibility | options | Post visibility settings (Public, Connections, etc.) |

## UseCases

- **Professional Content Publishing**: Automate LinkedIn post creation and publishing for personal branding and thought leadership
- **Social Media Marketing**: Schedule and manage LinkedIn content for business marketing and brand awareness campaigns
- **Employee Advocacy**: Enable team members to share company content on their personal LinkedIn profiles for organic reach
- **Lead Generation**: Share valuable content to attract prospects and generate leads through LinkedIn engagement
- **Corporate Communications**: Publish company updates, announcements, and news through organizational LinkedIn pages
- **Content Marketing Automation**: Integrate LinkedIn posting with content marketing workflows and editorial calendars
- **Personal Branding**: Build professional presence through consistent content sharing and industry thought leadership
- **Event Promotion**: Promote webinars, conferences, and business events through LinkedIn posts with rich media
- **Recruitment Marketing**: Share job openings, company culture content, and talent acquisition updates on LinkedIn
- **Cross-Platform Social Strategy**: Coordinate LinkedIn content with other social media platforms for unified messaging

