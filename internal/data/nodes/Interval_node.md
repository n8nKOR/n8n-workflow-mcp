# Interval Node

## Overview

Triggers the workflow in a given interval

## Credentials

None

## Inputs

None

## Outputs

- Main

## Properties

### Resource: Schedule Trigger

#### Operation: Interval Timer

| Field Name | Type | Description | Required |
|---|---|---|---|
| Interval | number | Interval value (must be 1 or higher) | Yes |
| Unit | options | Unit of the interval value (seconds, minutes, hours) | Yes |

### Interval Units

- **Seconds**: Trigger every N seconds
- **Minutes**: Trigger every N minutes  
- **Hours**: Trigger every N hours

### Trigger Behavior

#### Automatic Execution
- **Activation Required**: Workflow must be activated to run on schedule
- **Background Execution**: Runs automatically in the background
- **Continuous Operation**: Continues until workflow is deactivated

#### Manual Execution
- **Test Trigger**: Can be manually triggered from canvas for testing
- **Execute Workflow**: Use "execute workflow" button for immediate testing
- **Development**: Useful during workflow development and debugging

### Interval Calculation

#### Time Conversion
- **Seconds**: Direct value (interval × 1)
- **Minutes**: Converted to seconds (interval × 60)
- **Hours**: Converted to seconds (interval × 3600)

#### Internal Processing
- All intervals converted to milliseconds for JavaScript execution
- Uses `setInterval()` for reliable timing
- Maximum interval value: 2,147,483,647 milliseconds (≈ 24.8 days)

### Timing Constraints

#### Minimum Interval
- **1 Second**: Minimum allowed interval
- **Error Handling**: Values ≤ 0 throw validation error
- **Performance**: Very short intervals may impact system performance

#### Maximum Interval
- **24.8 Days**: Maximum interval due to JavaScript limitations
- **Error Handling**: Excessive values throw "interval value too large" error
- **Practical Limits**: Consider system resources for long intervals

### Activation Notice

The node displays an important notice:
> "This workflow will run on the schedule you define here once you activate it. For testing, you can also trigger it manually: by going back to the canvas and clicking 'execute workflow'"

### Output Data

#### Trigger Data
- **Empty Object**: Each trigger produces `[{}]` 
- **Consistent Format**: Same data structure for every execution
- **Timing Information**: Execution time available through workflow context

#### Data Flow
- **Single Item**: Each trigger generates one output item
- **JSON Format**: Output in standard n8n JSON format
- **No Payload**: No meaningful data in trigger output

## UseCases

#### Regular Monitoring
```
Interval: 5
Unit: minutes
Use Case: Check API status every 5 minutes
```

#### Daily Reports
```
Interval: 24
Unit: hours
Use Case: Generate daily summary reports
```

#### Frequent Updates
```
Interval: 30
Unit: seconds
Use Case: Real-time data synchronization
```

#### Weekly Maintenance
```
Interval: 168
Unit: hours (7 days)
Use Case: Weekly cleanup tasks
```

### Integration Patterns

#### Data Collection
- **API Polling**: Regular API calls for new data
- **Database Sync**: Periodic database synchronization
- **File Processing**: Regular file system monitoring

#### Maintenance Tasks
- **Cleanup Operations**: Regular cleanup of temporary data
- **Health Checks**: System health monitoring
- **Backup Tasks**: Automated backup schedules

#### Notification Systems
- **Status Updates**: Regular status notifications
- **Report Generation**: Scheduled report creation
- **Alert Systems**: Periodic alert checks

### Performance Considerations

#### System Impact
- **Resource Usage**: Frequent triggers consume system resources
- **Concurrent Executions**: Multiple workflows can run simultaneously
- **Memory Management**: Each execution creates new workflow instance

#### Optimization Tips
- **Appropriate Intervals**: Don't make intervals unnecessarily short
- **Efficient Workflows**: Optimize connected nodes for performance
- **Error Handling**: Include proper error handling in triggered workflows

### Error Handling

#### Validation Errors
- **Negative Values**: Interval ≤ 0 throws error
- **Excessive Values**: Interval > max limit throws error
- **Type Errors**: Non-numeric intervals cause validation errors

#### Runtime Errors
- **Workflow Errors**: Errors in triggered workflow don't stop trigger
- **Cleanup**: Failed executions are cleaned up automatically
- **Logging**: Execution errors logged for debugging

### Lifecycle Management

#### Activation
- **Manual Activation**: Must be activated through workflow interface
- **Immediate Start**: Trigger starts immediately upon activation
- **Persistent**: Continues running until deactivated

#### Deactivation
- **Manual Stop**: Deactivate workflow to stop trigger
- **Immediate Effect**: Stops scheduling new executions immediately
- **Cleanup**: Clears internal timer resources

### Migration Notice

**⚠️ Hidden Node**: This node is marked as hidden, indicating it may be deprecated. Consider using the newer **Schedule Trigger** node for new workflows, which provides:
- More scheduling options (cron expressions)
- Better timezone support
- Enhanced configuration options

### Comparison with Schedule Trigger

#### Interval Node (Simple)
- **Fixed Intervals**: Simple recurring intervals
- **Three Units**: Seconds, minutes, hours only
- **Easy Setup**: Straightforward configuration

#### Schedule Trigger (Advanced)
- **Cron Expressions**: Complex scheduling patterns
- **Timezone Support**: Multiple timezone options
- **Calendar-based**: Day/month specific scheduling

### Usage Notes
- Node is marked as hidden in current version - consider Schedule Trigger for new workflows
- Requires workflow activation to run automatically
- Manual execution available for testing during development
- Minimum interval is 1 second, maximum approximately 24.8 days
- Each trigger produces an empty object as output data
- System performance may be impacted by very frequent triggers
- Error handling should be implemented in connected workflow nodes
- Deactivating the workflow immediately stops the trigger
- Uses JavaScript setInterval internally for reliable timing

