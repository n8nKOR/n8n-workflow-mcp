# Strava Node

## Overview

Consume Strava API

## Credentials

- Name: stravaOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: activity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | The name of the activity | Yes |  |
| Type | string | Type of activity. For example - Run, Ride etc. | Yes |  |
| Sport Type | options | Type of sport | No |  |
| Start Date | dateTime | ISO 8601 formatted date time | Yes |  |
| Elapsed Time (Seconds) | number | In seconds | Yes |  |
| Additional Fields | collection | Whether to mark as commute | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | string | ID or email of activity | Yes |  |
| Update Fields | collection | Whether to mark as commute | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Activity ID | string | ID or email of activity | Yes |  |

#### Operation: getStreams

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Keys | multiOptions | Desired stream types to return | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

## UseCases

- **Fitness Data Integration** : Integrate Strava fitness data with health and wellness applications
- **Activity Tracking** : Track and analyze athletic activities and performance metrics
- **Training Analytics** : Analyze training data for performance improvement and coaching
- **Social Fitness** : Share and compare fitness activities with friends and communities
- **Goal Tracking** : Track fitness goals and progress over time
- **Health Monitoring** : Monitor health and fitness trends through activity data
- **Coaching Applications** : Support coaching and training programs with activity data
- **Fitness Challenges** : Create and manage fitness challenges and competitions
- **Performance Reporting** : Generate performance reports and analytics for athletes
- **Wearable Device Integration** : Integrate with wearable devices and fitness trackers

