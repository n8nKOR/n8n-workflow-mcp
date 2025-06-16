# Schedule Trigger Node

## Overview

The Schedule Trigger node executes workflows on a defined schedule using various timing intervals. This trigger node supports scheduling options from simple intervals to cron expressions for automated workflow execution.

**Core Features**:
- Basic interval scheduling (seconds, minutes, hours, days, weeks, months)
- Custom cron expression support
- Timezone-aware scheduling

## Credentials

None - This node requires no external authentication credentials.

## Inputs

None - This is a trigger node that starts workflows based on time schedules.

## Outputs

- Main: Outputs timestamp data when the scheduled trigger executes

## Properties

### Resource: Schedule

#### Operation: Trigger

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Trigger Interval | options | The type of interval for triggering (Seconds, Minutes, Hours, Days, Weeks, Months, Custom Cron) | Yes | Minutes |
| Seconds | number | Number of seconds between triggers (minimum 30) | No | 30 |
| Minutes | number | Number of minutes between triggers | No | 5 |
| Hours | number | Number of hours between triggers | No | 1 |
| Hour | options | Specific hour to trigger (0-23) | No | 0 |
| Minute | number | Specific minute to trigger (0-59) | No | 0 |
| Days | number | Number of days between triggers | No | 1 |
| Weekdays | multiOptions | Specific weekdays for weekly triggers | No | Monday |
| Weeks | number | Number of weeks between triggers | No | 1 |
| Day of Month | number | Specific day of month for monthly triggers (1-31) | No | 1 |
| Months | number | Number of months between triggers | No | 1 |
| Expression | string | Custom cron expression for advanced scheduling | No | - |

### Basic Trigger Intervals

#### Seconds
Triggers every N seconds (minimum 30 seconds).

#### Minutes  
Triggers every N minutes - most common option.

#### Hours
Triggers every N hours with optional minute specification.

#### Days
Triggers every N days with optional hour and minute.

#### Weeks
Triggers every N weeks on specified weekdays.

#### Months
Triggers every N months on specified day.

#### Custom Cron
Allows custom cron expressions: `[Minute] [Hour] [Day of Month] [Month] [Day of Week]`

### Output Data

When triggered, outputs timestamp information:

| Field | Type | Description |
|-------|------|-------------|
| timestamp | string | ISO 8601 timestamp of execution |
| year | number | Current year |
| month | number | Current month (1-12) |
| day | number | Current day of month |
| hour | number | Current hour (0-23) |
| minute | number | Current minute (0-59) |
| second | number | Current second (0-59) |

## UseCases

- **Automated Reporting**: Generate and distribute reports on a daily, weekly, or monthly schedule
- **Data Synchronization**: Regularly sync data between different systems and databases
- **Maintenance Tasks**: Schedule automated cleanup, backup, and optimization processes
- **Business Process Automation**: Trigger recurring business workflows like invoice processing or inventory updates
- **Monitoring and Alerts**: Perform regular health checks and send alerts when issues are detected
- **Content Publishing**: Schedule content publication and social media posts at optimal times
- **Batch Processing**: Process large datasets during off-peak hours for optimal performance
- **API Rate Limit Management**: Schedule API calls to respect rate limits and avoid throttling
- **Email Campaigns**: Send scheduled marketing emails and newsletters to customers
- **System Monitoring**: Regular monitoring of system performance and resource usage