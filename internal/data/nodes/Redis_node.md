# Redis Node

## Overview

Get, send and update data in Redis

## Credentials

- Name: redis, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Operations

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform. Options: "Delete", "Get", "Increment", "Info", "Keys", "Pop", "Publish", "Push", "Set" | Yes |
| Key | String | Name of the key to delete from Redis | Yes |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Name | String | Name of the property to write received data to (supports dot-notation) | Yes |
| Key | String | Name of the key to get from Redis | Yes |
| Key Type | Options | Type of key data. Options: "Automatic", "Hash", "List", "Sets", "String" | Yes |

#### Operation: Increment
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Key | String | Name of the key to increment | Yes |

#### Operation: Info
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |

#### Operation: Keys
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Key Pattern | String | Pattern to match keys (supports Redis wildcards) | Yes |

#### Operation: Pop
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| List | String | Name of the Redis list to pop from | Yes |
| Tail | Boolean | Whether to pop from tail (right) instead of head (left) | No |

#### Operation: Publish
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Channel | String | Redis channel to publish message to | Yes |
| Message | String | Message to publish | Yes |

#### Operation: Push
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| List | String | Name of the Redis list to push to | Yes |
| Value | String | Value to push to the list | Yes |
| Tail | Boolean | Whether to push to tail (right) instead of head (left) | No |

#### Operation: Set
| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| Key | String | Name of the key to set in Redis | Yes |
| Value | String | Value to set for the key | Yes |
| Expire | Boolean | Whether to set an expiration time | No |
| TTL | Number | Time to live in seconds | No |

## UseCases

- **Caching and Performance Optimization**: Store frequently accessed data in Redis for fast retrieval, reducing database load and improving application response times
- **Session Management**: Store user session data, authentication tokens, and temporary user states for web applications and APIs
- **Real-time Analytics**: Track and store real-time metrics, counters, and analytics data with automatic expiration and increment operations
- **Message Queue Implementation**: Use Redis lists and pub/sub functionality to implement job queues, task processing, and inter-service communication
- **Rate Limiting and Throttling**: Implement API rate limiting, request throttling, and usage tracking using Redis counters with expiration
- **Temporary Data Storage**: Store temporary workflow data, intermediate processing results, and short-lived application state
- **Distributed Locking**: Implement distributed locks and synchronization mechanisms for coordinating operations across multiple processes
- **Configuration Management**: Store and retrieve application configuration, feature flags, and dynamic settings with fast access
- **Chat and Messaging Systems**: Build real-time chat applications using Redis pub/sub for message broadcasting and delivery
- **Workflow State Management**: Maintain workflow execution state, progress tracking, and temporary data during complex automation processes

