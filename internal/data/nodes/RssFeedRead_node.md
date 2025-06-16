# RSS Read Node

## Overview

Reads data from an RSS Feed and extracts feed items. This node supports multiple RSS formats including RSS 2.0, RSS 1.0 (RDF), and Atom feeds. It can handle SSL certificate issues and provides comprehensive feed parsing capabilities for automation workflows. The node processes RSS feeds and returns individual feed items as separate workflow items.

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: RSS Feed

#### Operation: Read Feed

| Field Name | Type | Description | Required |
|---|---|---|---|
| URL | string | URL of the RSS feed to read. Supports HTTP and HTTPS protocols | Yes |

### Options

| Field Name | Type | Description | Required |
|---|---|---|---|
| Ignore SSL Issues (Insecure) | boolean | Whether to ignore SSL/TLS certificate issues when connecting to HTTPS feeds | No |

#### Additional Information
- The node automatically sets appropriate Accept headers for RSS/Atom feeds
- Supports RSS 2.0, RSS 1.0 (RDF), and Atom feed formats
- Each feed item becomes a separate output item in the workflow
- Feed metadata is available in the parsed output
- Connection errors are handled gracefully with descriptive error messages

## UseCases

- **AI-Powered Content Monitoring**: Monitor multiple RSS feeds for relevant articles, use AI to classify content relevance, and automatically summarize important articles for team distribution via Slack or email
- **News Aggregation and Analysis**: Collect articles from various news sources, filter by keywords or topics, and create curated content collections for newsletters or content marketing campaigns
- **Industry News Alerting**: Track industry news, competitor updates, or specific topics through RSS feeds and trigger alerts when important developments occur, integrating with notification systems
- **Automated Content Research**: Gather content from RSS feeds as research material for AI-powered content generation, blog writing, or market analysis workflows
- **Blog Content Syndication**: Monitor competitor blogs or industry publications to stay updated on trends and automatically collect content for analysis or inspiration
- **Social Media Content Curation**: Extract RSS feed content to automatically populate social media scheduling tools with relevant industry content
- **E-commerce Price Monitoring**: Monitor product RSS feeds from e-commerce sites to track price changes, new product launches, or inventory updates
- **Podcast Episode Tracking**: Monitor podcast RSS feeds to automatically track new episodes and integrate with content management or notification systems

