# Filter Node

## Overview

Remove items matching a condition

## Credentials

None

## Inputs

- Main

## Outputs

- Main (2 outputs: Kept, Discarded)

## Properties

### Resource: Item Filter

#### Operation: Filter Items

| Field Name | Type | Description | Required |
|---|---|---|---|
| Conditions | filter | Filtering conditions to evaluate | Yes |
| Ignore Case | boolean | Whether to ignore letter case when evaluating conditions | No |
| Loose Type Validation | boolean | Enable less strict type validation (v2.1+) | No |

### Condition Types
- **String Conditions**: Contains, equals, starts with, ends with, regex
- **Number Conditions**: Equal, not equal, greater than, less than, between
- **Boolean Conditions**: True, false, is set, is not set
- **Date/Time Conditions**: Before, after, between dates
- **Array Conditions**: Contains, length, empty/not empty

### Logical Operators
- **AND**: All conditions must be true
- **OR**: At least one condition must be true
- **Complex Logic**: Nested groups with mixed AND/OR logic

### Version-Specific Features

#### Version 2.2+
- Enhanced filter component with improved UI
- Better type validation
- Improved performance

#### Version 2.1+
- Loose type validation option
- Better error handling for type mismatches

### Output Behavior

The Filter node has **two outputs**:

1. **Output 1 (Kept)**: Items that **match** the conditions (pass the filter)
2. **Output 2 (Discarded)**: Items that **don't match** the conditions (fail the filter)

### Usage Examples

#### Filter by String Field
```
Field: customer.status
Condition: equals
Value: "active"
```

#### Filter by Number Range
```
Field: order.amount
Condition: greater than
Value: 100
```

#### Complex Filter with Multiple Conditions
```
AND Group:
  - customer.type equals "premium"
  - order.amount greater than 50
OR
  - customer.status equals "vip"
```

### Error Handling

- **Continue on Fail**: When enabled, items that cause errors are sent to the "Discarded" output
- **Type Validation**: Configurable strictness for type checking
- **Case Sensitivity**: Optional for string comparisons

### Usage Notes
- Each input item is evaluated independently against the conditions
- Items maintain their pairedItem relationships through filtering
- Use the two outputs to handle both matching and non-matching items
- Conditions support expressions for dynamic filtering
- Filter component provides visual condition building
- Supports nested logic for complex filtering scenarios

## UseCases

- **Active Route Filtering**: In multi-user email-to-Notion task conversion systems, filter Airtable routing information to process only active routes (Active=true), ensuring only valid user configurations are processed
- **Change Detection Filtering**: In visual regression testing workflows, filter AI analysis results to include only web pages where actual changes were detected, excluding unchanged pages from reports
- **Data Quality Control**: Filter customer records to exclude incomplete or invalid entries before processing in CRM systems
- **Priority-based Processing**: Filter support tickets by priority level to ensure high-priority issues are processed first in automated workflows
- **Content Moderation**: Filter user-generated content to separate approved content from content requiring manual review
- **Performance Optimization**: Filter large datasets to process only records that meet specific criteria, reducing downstream processing overhead
- **Conditional Workflow Routing**: Split workflow execution paths based on data conditions, sending different types of records to appropriate processing chains

