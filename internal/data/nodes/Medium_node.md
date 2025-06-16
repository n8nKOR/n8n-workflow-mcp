# Medium Node

## Overview

Consume Medium API

## Credentials

This node supports multiple Medium authentication methods:
- **mediumApi**: API token authentication
- **mediumOAuth2Api**: OAuth2 authentication for enhanced security

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: post

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Publication | boolean | Whether you are posting for a publication | No | false |
| Publication Name or ID | options | Publication IDs. Choose from the list, or specify an ID using expression | No |  |
| Title | string | Title of the post. Max Length : 100 characters. | Yes |  |
| Content Format | options | The format of the content to be posted | Yes |  |
| Content | string | The body of the post, in a valid semantic HTML fragment, or Markdown | Yes |  |
| Additional Fields | collection | Additional post configuration options | No |  |

#### Additional Post Fields
| Field Name | Type | Description | Default |
|---|---|---|---|
| Canonical URL | string | The original home of this content, if it was originally published elsewhere |  |
| License | options | License for the post content | all-rights-reserved |
| Notify Followers | boolean | Whether to notify followers about the new post | false |
| Publish Status | options | Post publication status (public, draft, unlisted) | public |
| Tags | array | Tags to associate with the post |  |

### Resource: publication

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 50 |

## UseCases

- **Content Publishing Automation**: Automatically publish blog posts to Medium from content management systems, enabling seamless content distribution workflows
- **Cross-Platform Content Distribution**: Simultaneously publish articles to Medium and other platforms like WordPress, LinkedIn, or personal blogs for maximum reach
- **Content Curation and Republishing**: Curate content from various sources, format it appropriately, and republish to Medium with proper attribution and formatting
- **Newsletter to Blog Conversion**: Convert newsletter content, email campaigns, or internal documentation into Medium articles for public sharing
- **AI-Generated Content Publishing**: Use AI tools to generate articles, summaries, or tutorials and automatically publish them to Medium with proper formatting
- **Documentation Publishing**: Transform technical documentation, API guides, or product tutorials into readable Medium articles for developer communities
- **Social Media Integration**: Create workflows that convert social media threads, LinkedIn posts, or Twitter content into structured Medium articles
- **Personal Branding Automation**: Automate personal branding efforts by scheduling and publishing thought leadership content to Medium based on triggers
- **Content Syndication**: Syndicate high-performing content from company blogs or websites to Medium to increase visibility and engagement
- **Event Content Publishing**: Automatically publish event summaries, conference notes, or webinar transcripts to Medium for knowledge sharing

