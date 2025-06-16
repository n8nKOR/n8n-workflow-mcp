# Date & Time Node

## Overview

Allows you to manipulate date and time values

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: datetime

#### Operation: getCurrentDate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Include Time | boolean | Whether to include time in the current date | No | false |
| Output Field Name | string | Name of the field to output the current date to | Yes | currentDate |
| Options | collection | Additional options | No |  |

##### Options Collection Fields:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Timezone | string | Timezone to use for the current date | No |  |

#### Operation: addToDate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Date | string | The date to add time to | Yes |  |
| Time Unit | options | Unit of time to add (years, months, weeks, days, hours, minutes, seconds) | Yes | days |
| Duration | number | Amount of time to add | Yes | 1 |
| Output Field Name | string | Name of the field to output the result to | Yes | newDate |

#### Operation: subtractFromDate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Date | string | The date to subtract time from | Yes |  |
| Time Unit | options | Unit of time to subtract (years, months, weeks, days, hours, minutes, seconds) | Yes | days |
| Duration | number | Amount of time to subtract | Yes | 1 |
| Output Field Name | string | Name of the field to output the result to | Yes | newDate |

#### Operation: formatDate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Date | string | The date to format | Yes |  |
| Format | options | Format to use for the date | Yes | MM/dd/yyyy |
| Output Field Name | string | Name of the field to output the formatted date to | Yes | formattedDate |
| Options | collection | Additional options | No |  |

##### Options Collection Fields:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Timezone | boolean | Whether to use timezone in formatting | No | false |
| From Format | string | Format of the input date | No |  |

#### Operation: roundDate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Date | string | The date to round | Yes |  |
| Mode | options | Round mode (roundDown, roundUp) | Yes | roundDown |
| To Nearest | options | Unit to round to (year, month, day, hour, minute) | Yes | day |
| Output Field Name | string | Name of the field to output the result to | Yes | roundedDate |

#### Operation: getTimeBetweenDates

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | string | The start date | Yes |  |
| End Date | string | The end date | Yes |  |
| Units | multiOptions | Units to return the difference in | Yes | days |
| Output Field Name | string | Name of the field to output the result to | Yes | timeDifference |
| Options | collection | Additional options | No |  |

#### Operation: extractDate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Date | string | The date to extract from | Yes |  |
| Part | options | Part of the date to extract (year, month, day, hour, minute, second) | Yes | year |
| Output Field Name | string | Name of the field to output the result to | Yes | extractedPart |

## UseCases

- **Schedule Management**: Calculate meeting times, deadlines, and recurring events with proper timezone handling
- **Age Calculations**: Determine ages, service durations, or time-based eligibility criteria for automated decision making
- **Data Grouping**: Group time-series data by time periods (daily, weekly, monthly) for analytics and reporting
- **Expiration Tracking**: Monitor subscription renewals, license expirations, and time-sensitive processes
- **Business Hours Validation**: Ensure operations occur within business hours across different timezones
- **Report Generation**: Add timestamps and date calculations to automated reports and data exports
- **Event Scheduling**: Calculate optimal timing for automated processes based on business requirements

