# OpenWeatherMap Node

## Overview

Gets current and future weather information

## Credentials

- Name: openWeatherMapApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Weather

#### Operation: Current Weather
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Location | string | City name, coordinates, or location identifier | Yes | - |
| Units | options | Temperature units (standard, metric, imperial) | No | standard |
| Language | options | Language for weather descriptions | No | en |

#### Operation: 5 Day Forecast
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Location | string | City name, coordinates, or location identifier | Yes | - |
| Units | options | Temperature units (standard, metric, imperial) | No | standard |
| Language | options | Language for weather descriptions | No | en |
| Count | number | Number of forecast periods to return | No | 40 |

#### Operation: One Call
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Latitude | number | Latitude coordinate | Yes | - |
| Longitude | number | Longitude coordinate | Yes | - |
| Units | options | Temperature units (standard, metric, imperial) | No | standard |
| Language | options | Language for weather descriptions | No | en |
| Exclude | collection | Parts of weather data to exclude | No | - |

**Location Formats:**
- **City name**: "London" or "London,UK"
- **Coordinates**: "lat,lon" format
- **City ID**: OpenWeatherMap city ID
- **ZIP code**: "ZIP,country" format

## UseCases

- **Weather Monitoring**: Monitor current weather conditions
- **Weather Forecasting**: Get weather forecasts for planning
- **Agricultural Applications**: Weather data for farming decisions
- **Travel Planning**: Check weather conditions for travel
- **Event Planning**: Weather-based event planning
- **IoT Applications**: Weather data for smart home systems
