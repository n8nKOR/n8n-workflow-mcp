# Yourls Node

## Overview

Consume Yourls API

## Credentials

- Name: yourlsApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: url

#### Operation: shorten

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| URL | string | The URL to shorten | Yes |  |
| Additional Fields | collection | Title for custom short URLs | No |  |

#### Operation: expand

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Short URL | string | The short URL to expand | Yes |  |

#### Operation: stats

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Short URL | string | The short URL for which to get stats | Yes |  |

## UseCases

- [URL Shortening] : Create short URLs for sharing long links in social media and marketing
- [Link Tracking] : Track click analytics and engagement metrics for shortened URLs
- [Campaign Management] : Create trackable links for marketing campaigns and promotions
- [Social Media Marketing] : Share short URLs on social platforms with limited character counts
- [Email Marketing] : Include short URLs in email campaigns for better click tracking
- [Analytics and Reporting] : Generate reports on link performance and user engagement
- [Brand Customization] : Create branded short URLs using custom domains
- [Link Management] : Organize and manage collections of shortened URLs
- [A/B Testing] : Create multiple short URLs to test different marketing messages
- [QR Code Generation] : Generate QR codes for shortened URLs for offline marketing
- [API Integration] : Integrate URL shortening into applications and workflows
- [Content Distribution] : Share content across multiple channels with trackable links
- [Performance Monitoring] : Monitor link performance and optimize marketing strategies
- [Custom Redirects] : Create custom redirect rules for enhanced user experience
- [Security and Privacy] : Manage URL access and implement security measures for links

