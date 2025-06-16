# Reddit Node

## Overview

Consume the Reddit API

## Credentials

- Name: redditOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: postComment

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | ID of the post to write the comment to. Found in the post URL: <code>/r/[subreddit_name]/comments/[post_id]/[post_title]</code> | Yes |  |
| Comment Text | string | Text of the comment. Markdown supported. | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subreddit | string | The name of subreddit where the post is | Yes |  |
| Post ID | string | ID of the post to get all comments from. Found in the post URL: <code>/r/[subreddit_name]/comments/[post_id]/[post_title]</code> | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comment ID | string | ID of the comment to remove. Found in the comment URL:<code>/r/[subreddit_name]/comments/[post_id]/[post_title]/[comment_id]</code> | Yes |  |

#### Operation: reply

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Comment ID | string | ID of the comment to reply to. To be found in the comment URL: <code>www.reddit.com/r/[subreddit_name]/comments/[post_id]/[post_title]/[comment_id]</code> | Yes |  |
| Reply Text | string | Text of the reply. Markdown supported. | Yes |  |

### Resource: profile

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Details | options | Details of my account to retrieve | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: subreddit

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Content | options | Subreddit content to retrieve | Yes |  |
| Subreddit | string | The name of subreddit to retrieve the content from | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | The keyword for the subreddit search | No |  |

### Resource: post

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subreddit | string | Subreddit to create the post in | Yes |  |
| Kind | options | The kind of the post to create | No |  |
| Title | string | Title of the post, up to 300 characters long | Yes |  |
| URL | string | URL of the post | Yes |  |
| Text | string | Text of the post. Markdown supported. | Yes |  |
| Resubmit | boolean | Whether the URL will be posted even if it was already posted to the subreddit before. Otherwise, the re-posting will trigger an error. | No |  |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Post ID | string | ID of the post to delete. Found in the post URL: <code>/r/[subreddit_name]/comments/[post_id]/[post_title]</code> | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subreddit | string | The name of subreddit to retrieve the post from | Yes |  |
| Post ID | string | ID of the post to retrieve. Found in the post URL: <code>/r/[subreddit_name]/comments/[post_id]/[post_title]</code> | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Subreddit | string | The name of subreddit to retrieve the posts from | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Category of the posts to retrieve | No |  |

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Location | options | Location where to search for posts | No |  |
| Subreddit | string | The name of subreddit to search in | Yes |  |
| Keyword | string | The keyword for the search | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | The category to sort results by | No |  |

### Resource: user

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Username | string | Reddit ID of the user to retrieve | Yes |  |
| Details | options | Details of the user to retrieve | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- [Community Management] : Monitor and manage subreddit communities and user interactions
- [Content Aggregation] : Collect and analyze posts from relevant subreddits for insights
- [Automated Posting] : Schedule and publish content to specific subreddits automatically
- [Comment Monitoring] : Track comments on posts for engagement and feedback analysis
- [Social Media Analytics] : Analyze Reddit discussions for market research and sentiment
- [User Profile Analysis] : Study user behavior and posting patterns for research purposes
- [Content Moderation] : Automate moderation tasks like removing inappropriate content
- [Trend Detection] : Identify trending topics and discussions across subreddits
- [Engagement Automation] : Automatically reply to comments and engage with community
- [Brand Monitoring] : Monitor mentions of brands or products across Reddit communities
- [Research Data Collection] : Gather data from Reddit for academic or business research
- [Notification Systems] : Set up alerts for specific keywords or trending topics
- [Cross-Platform Integration] : Sync Reddit content with other social media platforms
- [Content Curation] : Curate interesting content from Reddit for newsletters or blogs
- [Community Insights] : Generate reports on community activity and engagement metrics

