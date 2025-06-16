# If Node

## Overview

Route items to different branches (true/false)

## Credentials

None

## Inputs

- Main

## Outputs

- Main (true)
- Main (false)

## Properties

### Resource: Conditional Router

#### Operation: Route Items

| Field Name | Type | Description | Required |
|---|---|---|---|
| Conditions | filter | Conditions to evaluate for routing items to true/false branches | Yes |
| Ignore Case | boolean | Whether to ignore letter case when evaluating conditions | No |
| Loose Type Validation | boolean | Whether to use loose type validation | No |

### Condition System
The If node uses a powerful filter system to evaluate conditions with:
- **Multiple Conditions**: Combine multiple conditions with AND/OR logic
- **Nested Groups**: Create complex logical groupings
- **Type Validation**: Automatic type checking and validation
- **Case Sensitivity**: Configurable case sensitivity for string comparisons

### Comparison Operators
- **Equal**: Exact match comparison
- **Not Equal**: Non-matching comparison
- **Greater Than**: Numeric greater than
- **Greater Than or Equal**: Numeric greater than or equal
- **Less Than**: Numeric less than
- **Less Than or Equal**: Numeric less than or equal

### String Operators
- **Contains**: String contains substring
- **Starts With**: String starts with prefix
- **Ends With**: String ends with suffix
- **Regex**: Regular expression matching
- **Is Empty**: Check if string is empty
- **Is Not Empty**: Check if string has content

### Array Operators
- **Contains**: Array contains specific value
- **Length**: Array length comparison
- **Is Empty**: Array has no elements
- **Is Not Empty**: Array has elements

### Boolean Operators
- **Is True**: Value is exactly true
- **Is False**: Value is exactly false

### Existence Operators
- **Exists**: Property exists (not null/undefined)
- **Does Not Exist**: Property is null/undefined

### Type Validation

#### Strict Mode (Default)
- **Type Checking**: Enforces strict type matching
- **Error Handling**: Throws errors on type mismatches
- **Data Integrity**: Ensures data consistency

#### Loose Mode
- **Type Coercion**: Attempts automatic type conversion
- **Flexible Matching**: More permissive comparisons
- **Backward Compatibility**: Compatible with older workflows

### Output Routing

#### True Branch (Output 1)
- Items that **match** the conditions
- Continue to the first output connection
- Maintain original item structure and pairing

#### False Branch (Output 2)  
- Items that **do not match** the conditions
- Continue to the second output connection
- Maintain original item structure and pairing

### Data Preservation

#### Paired Items
- **Item Pairing**: Maintains relationship between input and output items
- **Index Tracking**: Preserves original item indices
- **Metadata**: Retains all item metadata

#### Error Handling
- **Continue on Fail**: Option to route failed items to false branch
- **Error Context**: Provides detailed error information
- **Item Index**: Includes item index in error context

### Logical Operators

#### AND Logic
All conditions must be true for item to route to true branch:
```
Condition 1: age > 18
AND
Condition 2: status = "active"
```

#### OR Logic  
Any condition can be true for item to route to true branch:
```
Condition 1: department = "sales"
OR
Condition 2: department = "marketing"
```

#### Mixed Logic
Complex combinations with grouping:
```
(department = "sales" OR department = "marketing")
AND
(status = "active" AND experience > 2)
```

### Expression Support

#### Dynamic Values
All condition values support n8n expressions:
```
Field: $json.age
Operator: Greater Than
Value: {{ $json.minimumAge || 18 }}
```

#### Function Calls
Use JavaScript functions in conditions:
```
Field: $json.email
Operator: Regex
Value: {{ /^[^\s@]+@[^\s@]+\.[^\s@]+$/ }}
```

### Case Sensitivity

#### Case Sensitive (Default: false)
- **Ignore Case**: "Hello" equals "hello"
- **String Matching**: Case-insensitive string operations
- **Flexible Comparisons**: More user-friendly matching

#### Case Sensitive (true)
- **Exact Case**: "Hello" does not equal "hello"  
- **Precise Matching**: Exact string comparisons
- **Strict Validation**: More precise data validation

### Usage Examples

#### Simple Age Filter
```
Conditions:
- Field: $json.age
- Operator: Greater Than or Equal
- Value: 18

Result: Adults to true branch, minors to false branch
```

#### Complex User Filter
```
Conditions:
Group 1 (OR):
- Field: $json.role
- Operator: Equal  
- Value: "admin"

OR

- Field: $json.department
- Operator: Equal
- Value: "management"

AND

Group 2:
- Field: $json.active
- Operator: Is True

Result: Active admins or managers to true branch
```

#### String Pattern Matching
```
Conditions:
- Field: $json.email
- Operator: Regex
- Value: ^[a-zA-Z0-9._%+-]+@company\.com$

Options:
- Ignore Case: true

Result: Company emails to true branch
```

#### Array Filtering
```
Conditions:
- Field: $json.tags
- Operator: Contains
- Value: "priority"

Result: Items with priority tag to true branch
```

### Performance Considerations

#### Condition Optimization
- **Simple Conditions**: Process faster than complex ones
- **Early Termination**: AND logic stops at first false condition
- **OR Logic**: Stops at first true condition
- **Condition Ordering**: Most selective conditions should be placed first

#### Large Datasets
- **Batch Processing**: Processes all items in single execution
- **Memory Usage**: Efficient memory management for large datasets
- **Error Isolation**: Individual item errors don't affect others

### Migration from V1

#### Key Changes
- **Filter System**: New powerful condition builder
- **Type Validation**: Enhanced type checking
- **Error Handling**: Improved error reporting
- **Case Sensitivity**: Better string handling

#### Upgrade Benefits
- **More Flexible**: Complex condition combinations
- **Better UX**: Visual condition builder
- **Improved Performance**: Optimized execution
- **Enhanced Debugging**: Better error messages

### Best Practices

#### Condition Design
- **Keep Simple**: Start with simple conditions, add complexity as needed
- **Test Thoroughly**: Test with representative data
- **Use Grouping**: Group related conditions logically

#### Error Handling
- **Enable Continue on Fail**: For robust workflows
- **Monitor Outputs**: Check both true and false branches
- **Log Errors**: Capture and review processing errors

#### Performance
- **Optimize Conditions**: Put most selective conditions first
- **Avoid Regex**: Use simpler operators when possible
- **Batch Size**: Consider breaking large datasets into smaller batches

### Advanced Features

#### Dynamic Conditions
- **Expression Fields**: Use `{{ }}` syntax for dynamic values
- **Workflow Context**: Access data from previous nodes
- **Environment Variables**: Reference system and custom variables
- **Date Functions**: Utilize Luxon date manipulation functions

#### Error Recovery
- **Continue on Fail**: Routes error items to false branch
- **Error Context**: Provides detailed error information
- **Item Tracking**: Maintains original item index in errors
- **Validation Messages**: Clear error descriptions for debugging

#### Performance Optimization
- **Batch Processing**: Handles multiple items efficiently
- **Memory Management**: Optimized for large datasets
- **Condition Caching**: Reuses compiled conditions when possible
- **Type Coercion**: Configurable type conversion for flexibility

### Usage Notes
- Supports unlimited nested conditions with complex AND/OR logic combinations
- Visual condition builder provides drag-and-drop interface for intuitive setup
- Type validation ensures data consistency with configurable strictness levels
- Case sensitivity can be configured globally for all string operations
- Expression support enables dynamic condition values from workflow context
- Error handling routes failed items to false branch when continue-on-fail is enabled
- Maintains complete item pairing and metadata throughout the routing process
- Performance optimized for large datasets with intelligent early termination
- Regular expression support includes full pattern matching with flags
- Date comparison handles multiple formats with timezone awareness
- Array operations support both content and structural comparisons
- Empty value detection covers all edge cases including NaN and invalid dates

## UseCases

- **Data Validation**: Filter and route data based on quality checks, ensuring only valid records proceed to downstream processing
- **User Segmentation**: Route users to different workflow paths based on demographics, behavior, or subscription status for personalized automation
- **Content Moderation**: Automatically approve or flag content based on compliance rules, keyword detection, and safety criteria
- **Pricing Logic**: Route orders to different processing paths based on order value, customer tier, or promotional eligibility
- **Error Handling**: Separate successful API responses from errors for appropriate handling and retry logic
- **A/B Testing**: Route users to different experimental groups based on randomization or specific criteria for testing workflows
- **Workflow Branching**: Create complex decision trees for business processes like approval workflows, escalation procedures, and routing logic
- **Data Transformation**: Route items to different transformation pipelines based on data format, source system, or processing requirements
- **Alert Filtering**: Route monitoring data to create alerts only when specific thresholds or conditions are met
- **Compliance Checking**: Ensure data meets regulatory requirements before processing, routing non-compliant data for manual review
- **Dynamic Processing**: Choose different processing methods based on data characteristics, file types, or system availability
- **Quality Assurance**: Route test results to different paths based on pass/fail criteria for automated testing workflows
- **Customer Service**: Route support tickets to appropriate teams based on urgency, category, or customer tier
- **Inventory Management**: Route product updates to different systems based on availability, category, or warehouse location
- **Financial Processing**: Route transactions to different approval levels based on amount, account type, or risk assessment
- **Content Publishing**: Route content to different channels or formats based on content type, target audience, or publication schedule
- [File Format Validation] : Verify if email attachments are PDF files and route only PDF files to text extraction and AI analysis paths while handling other formats separately
- [AI Analysis Result Branching] : Check if OpenAI PDF content analysis results meet specific conditions (e.g., "true") and route only matching files to Google Drive upload paths
- [User Authentication] : Validate Telegram user credentials by checking first name, last name, and user ID against authorized user lists for secure bot access
- [Message Type Routing] : Route different types of incoming messages (text, audio, images) to appropriate processing paths in multimodal AI applications

