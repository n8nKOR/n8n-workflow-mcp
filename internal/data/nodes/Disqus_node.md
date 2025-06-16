# Disqus Node

## Overview

Access data on Disqus

## Credentials

- Name: disqusApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: forum

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Forum Name | string | The short name(aka ID) of the forum to get | Yes |  |
| Additional Fields | collection | You may specify relations to include with your response | No |  |

#### Operation: getPosts

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Forum Name | string | The short name(aka ID) of the forum to get | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | You may specify filters for your response | No |  |

#### Operation: getCategories

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Forum Name | string | The short name(aka ID) of the forum to get Categories | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | You may specify order to sort your response | No |  |

#### Operation: getThreads

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Forum Name | string | The short name(aka ID) of the forum to get Threads | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | You may specify relations to include with your response | No |  |

## UseCases

- **Community Engagement Analytics**: Monitor forum activity, track user engagement, and analyze discussion patterns across different topics
- **Content Moderation Automation**: Extract forum posts and comments for automated content moderation and community management
- **Social Media Monitoring**: Track discussions and conversations happening on Disqus-powered websites and forums
- **Customer Feedback Analysis**: Gather user feedback and comments from product discussions and support forums
- **SEO Content Strategy**: Analyze popular discussion topics and threads to inform content creation and SEO strategy
- **Competitor Research**: Monitor discussions on competitor websites and forums to understand market sentiment and trends
- **Community Management**: Automate responses to forum posts, manage user interactions, and maintain community guidelines
- **Lead Generation**: Identify potential leads from forum discussions and user profiles engaging with relevant topics
- **Brand Monitoring**: Track mentions of your brand or products across various Disqus-powered forums and comment sections
- **Research and Insights**: Collect discussion data for market research, customer insights, and trend analysis

