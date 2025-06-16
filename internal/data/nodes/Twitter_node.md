# X (Formerly Twitter) Node

## Overview

Post, like, and search tweets, send messages, search users, and add users to lists

## Credentials

- Name: twitterOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Direct Message

#### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| user | string | User ID or username to send message to | Yes |
| text | string | Content of the direct message | Yes |

### Resource: List

#### Operation: Add Member

| Field Name | Type | Description | Required |
|---|---|---|---|
| listId | string | ID of the list | Yes |
| userId | string | User ID to add to the list | Yes |

### Resource: Tweet

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| text | string | Content of the tweet | Yes |
| replyToTweetId | string | Tweet ID to reply to | No |
| media | collection | Media attachments | No |
| poll | collection | Poll options | No |
| location | object | Geographic location | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| tweetId | string | ID of the tweet to delete | Yes |

#### Operation: Like

| Field Name | Type | Description | Required |
|---|---|---|---|
| tweetId | string | ID of the tweet to like | Yes |

#### Operation: Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| searchText | string | Text to search for | Yes |
| returnAll | boolean | Whether to return all results | No |
| sortOrder | options | Order of results (recency, relevancy) | No |
| startTime | dateTime | Beginning of search time range | No |
| endTime | dateTime | End of search time range | No |
| tweetFields | multiOptions | Additional tweet metadata to retrieve | No |

#### Operation: Unlike

| Field Name | Type | Description | Required |
|---|---|---|---|
| tweetId | string | ID of the tweet to unlike | Yes |

### Resource: User

#### Operation: Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| username | string | Username to search for | No |
| userId | string | User ID to look up | No |
| me | boolean | Get information about the authenticated user | No |
| userFields | multiOptions | User metadata to retrieve | No |

## UseCases

- Social Media Management: Automate posting tweets and managing social media presence
- Content Monitoring: Search and track mentions, hashtags, or keywords
- User Engagement: Like tweets, reply to mentions, and send direct messages
- Data Collection: Gather tweet data for analysis and reporting
- List Management: Organize users into lists for targeted content delivery

