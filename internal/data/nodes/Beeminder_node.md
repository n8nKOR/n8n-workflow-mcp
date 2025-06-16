# Beeminder Node

## Overview

The Beeminder node allows you to consume the Beeminder API for goal tracking and data logging. Beeminder is a goal-tracking service that uses behavioral economics to help users achieve their goals by putting money on the line. This node enables you to create, update, and manage datapoints for your Beeminder goals within your n8n workflows.

## Credentials

This node requires Beeminder API credentials:
- **Username**: Your Beeminder username
- **API Token**: Your Beeminder API token

To obtain API credentials:
1. Log into your Beeminder account
2. Navigate to Settings → Account → API
3. Generate or copy your API token
4. Use your username and API token in n8n

## Inputs

- **Main**: JSON data for creating/updating datapoints

## Outputs

- **Main**: Returns JSON data from Beeminder API responses

## Properties

### Resource: Datapoint

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Goal Name | Options | Goal to add datapoint to | Yes |
| Value | Number | Datapoint value to send | Yes |
| Additional Fields | Collection | Optional datapoint properties | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Goal Name | Options | Goal containing the datapoint | Yes |
| Datapoint ID | String | ID of the datapoint to delete | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Goal Name | Options | Goal to retrieve datapoints from | Yes |
| Return All | Boolean | Whether to return all datapoints | No |
| Limit | Number | Maximum number of datapoints to return (max: 300) | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Goal Name | Options | Goal containing the datapoint | Yes |
| Datapoint ID | String | ID of the datapoint to update | Yes |
| Update Fields | Collection | Datapoint properties to update | No |

## Datapoint Properties

### Core Properties
- **Value**: The numeric value to track (required for creation)
- **Comment**: Optional comment or note about the datapoint
- **Timestamp**: When the datapoint occurred (defaults to current time)
- **Request ID**: Unique identifier to prevent duplicate entries

### Additional Fields for Creation
| Field Name | Type | Description |
|---|---|---|
| Comment | String | Optional comment about this datapoint |
| Timestamp | DateTime | When this data occurred (defaults to now) |
| Request ID | String | Unique identifier to prevent duplicates |

### Update Fields
| Field Name | Type | Description |
|---|---|---|
| Value | Number | New value for the datapoint |
| Comment | String | New comment for the datapoint |
| Timestamp | DateTime | New timestamp for the datapoint |



## UseCases

- **Fitness Tracking Automation** : Automatically log exercise data from fitness trackers to Beeminder goals
- **Productivity Monitoring** : Track work hours and task completion from project management tools
- **Habit Formation** : Sync daily habit completion from habit tracking apps to Beeminder
- **Weight Management** : Automatically log weight measurements from smart scales
- **Reading Goals** : Track pages read or books finished from reading apps
- **Time Tracking Integration** : Log focused work time from time tracking applications
- **Financial Goal Tracking** : Monitor spending or savings goals with bank account data
- **Sleep Quality Monitoring** : Track sleep hours from sleep tracking devices
- **Writing Progress** : Count words written or articles published from writing tools
- **Learning Goals** : Track study time or course completion from educational platforms

