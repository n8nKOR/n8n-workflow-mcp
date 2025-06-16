# Cron Node

## Overview

Automatically triggers workflow execution based on scheduled time intervals using cron expressions. This trigger node supports multiple trigger times per node, allowing complex scheduling patterns for automation workflows. Essential for automating recurring tasks, data synchronization, batch processing operations, and time-based workflow automation. Integrates with n8n's internal scheduler for reliable execution and supports both scheduled and manual triggering.

## Credentials

None

## Inputs

None (Trigger Node)

## Outputs

- Main

## Properties

### Resource: Scheduled Trigger

#### Operation: Schedule Execution

| Field Name | Type | Description | Required |
|---|---|---|---|
| Trigger Times | fixedCollection | Collection of cron expressions that define when the workflow should be triggered. Multiple trigger times can be added | Yes |

#### Trigger Time Configuration
Each trigger time consists of the following fields:

| Field Name | Type | Description | Range | Required |
|---|---|---|---|---|
| Minute | number/string | Minute when to trigger | 0-59, or * for every minute | Yes |
| Hour | number/string | Hour when to trigger | 0-23, or * for every hour | Yes |
| Day of Month | number/string | Day of the month when to trigger | 1-31, or * for every day | Yes |
| Month | number/string | Month when to trigger | 1-12, or * for every month | Yes |
| Day of Week | number/string | Day of the week when to trigger | 0-7 (0 and 7 are Sunday), or * for every day | Yes |

#### Additional Information
- This workflow will run on the schedule you define once you activate it
- For testing, you can trigger it manually by going back to the canvas and clicking 'execute workflow'
- Multiple trigger times can be added to run the workflow at different schedules
- Uses standard cron expression format for scheduling
- The node is hidden in the UI but fully functional
- Supports complex scheduling patterns with multiple triggers
- Activation message confirms when the cron trigger is successfully set up

## UseCases

- **Daily Report Generation**: Schedule daily business reports to be generated and sent every morning at 9 AM on weekdays
  ```
  Minute: 0, Hour: 9, Day of Month: *, Month: *, Day of Week: 1-5
  ```

- **Weekly Data Backup**: Perform weekly database backups every Sunday at 2 AM to ensure data safety
  ```
  Minute: 0, Hour: 2, Day of Month: *, Month: *, Day of Week: 0
  ```

- **Monthly Analytics Processing**: Generate monthly analytics reports on the first day of each month at 1 AM
  ```
  Minute: 0, Hour: 1, Day of Month: 1, Month: *, Day of Week: *
  ```

- **Hourly Health Checks**: Monitor system health and send alerts every hour on the hour
  ```
  Minute: 0, Hour: *, Day of Month: *, Month: *, Day of Week: *
  ```

- **Quarterly Financial Reports**: Generate quarterly financial summaries on the first day of each quarter
  ```
  Minute: 0, Hour: 8, Day of Month: 1, Month: 1,4,7,10, Day of Week: *
  ```

- **Daily Data Synchronization**: Sync data between systems every day at midnight
  ```
  Minute: 0, Hour: 0, Day of Month: *, Month: *, Day of Week: *
  ```

- **Weekly Newsletter Distribution**: Send weekly newsletters every Friday at 10 AM
  ```
  Minute: 0, Hour: 10, Day of Month: *, Month: *, Day of Week: 5
  ```

- **End-of-Month Processing**: Trigger month-end financial processing on the last day of each month
  ```
  Minute: 0, Hour: 23, Day of Month: 28-31, Month: *, Day of Week: *
  ```

