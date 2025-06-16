# Oura Node

## Overview

Consume Oura API

## Credentials

- Name: ouraApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: profile

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| (No additional fields) | - | Get the user's personal information | - | - |

### Resource: summary

#### Operation: getActivity

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 5 |
| Filters | collection | Add Filter | No |  |

##### Filters Collection Fields:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | dateTime | Start date for the summary retrieval. If omitted, it defaults to a week ago | No |  |
| End Date | dateTime | End date for the summary retrieval. If omitted, it defaults to the current day | No |  |

#### Operation: getReadiness

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 5 |
| Filters | collection | Add Filter | No |  |

##### Filters Collection Fields:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | dateTime | Start date for the summary retrieval. If omitted, it defaults to a week ago | No |  |
| End Date | dateTime | End date for the summary retrieval. If omitted, it defaults to the current day | No |  |

#### Operation: getSleep

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return | No | 5 |
| Filters | collection | Add Filter | No |  |

##### Filters Collection Fields:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | dateTime | Start date for the summary retrieval. If omitted, it defaults to a week ago | No |  |
| End Date | dateTime | End date for the summary retrieval. If omitted, it defaults to the current day | No |  |

## UseCases

- [Health Monitoring] : Track and analyze personal health metrics and vital signs
- [Sleep Analysis] : Monitor sleep patterns and quality for better sleep hygiene
- [Fitness Tracking] : Track daily activity levels and exercise performance metrics
- [Recovery Monitoring] : Monitor recovery metrics after workouts and stress periods
- [Health Reporting] : Generate regular health reports and insights for personal tracking
- [Wellness Integration] : Integrate health data with other wellness and fitness applications
- [Research Data Collection] : Collect health data for research and medical studies
- [Personal Analytics] : Analyze long-term health trends and patterns for lifestyle optimization
- [Automated Health Alerts] : Set up alerts for unusual health metrics or patterns
- [Healthcare Integration] : Share health data with healthcare providers and medical systems
- [Coaching Insights] : Provide data to fitness coaches and health professionals
- [Lifestyle Optimization] : Use health data to optimize daily routines and habits
- [Team Health Monitoring] : Monitor health metrics for sports teams or group wellness programs
- [Insurance Integration] : Share health data with insurance providers for wellness programs
- [Goal Tracking] : Track progress towards health and fitness goals with objective metrics

