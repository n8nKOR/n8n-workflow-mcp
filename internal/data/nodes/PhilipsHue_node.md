# Philips Hue Node

## Overview

Consume Philips Hue API

## Credentials

- Name: philipsHueOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: light

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Light ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Light ID | string |  | Yes |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Light Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| On | boolean | On/Off state of the light | Yes |  |
| Additional Fields | collection | The light is not performing an alert effect | No |  |

## UseCases

- [Smart Home Automation] : Automate lighting based on time, occupancy, or other smart home events
- [Mood Lighting] : Create dynamic lighting scenes for different moods and activities
- [Energy Management] : Control lighting to optimize energy consumption and reduce costs
- [Security Lighting] : Automate security lighting patterns when away from home or during night hours
- [Wake-up Lighting] : Create gradual wake-up lighting routines that simulate natural sunrise
- [Entertainment Integration] : Sync lighting with music, movies, or gaming for immersive experiences
- [Productivity Enhancement] : Adjust lighting throughout the day to support circadian rhythms and productivity
- [Remote Control] : Control home lighting remotely from anywhere via mobile apps or voice commands
- [Scene Management] : Create and manage custom lighting scenes for different rooms and occasions
- [Schedule Automation] : Set up automatic lighting schedules for daily routines and energy savings
- [Notification Lighting] : Use lights as visual notifications for alerts, calls, or smart home events
- [Health and Wellness] : Use lighting therapy for seasonal affective disorder or sleep improvement
- [Party and Event Lighting] : Create dynamic lighting effects for parties and special events
- [Home Office Optimization] : Optimize lighting for video calls and work productivity
- [Integration with IoT] : Connect with other IoT devices for comprehensive smart home experiences

