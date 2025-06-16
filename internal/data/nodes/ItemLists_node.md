# Item Lists Node

## Overview

Helper for working with lists of items and transforming arrays

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Item Lists

#### Operation: Split Out Items

| Field Name | Type | Description | Required |
|---|---|---|---|
| Field to Split Out | String | The field containing array or object to split | Yes |
| Include Other Fields | Boolean | Whether to include other fields from original item | No |
| Destination Field Name | String | Name for the field containing split values | No |

#### Operation: Concatenate Items

| Field Name | Type | Description | Required |
|---|---|---|---|
| Fields to Aggregate | Multi Options | Fields to combine into arrays | Yes |
| Put Output in Field | String | Field name for the aggregated result | No |
| Keep Only Unique Values | Boolean | Remove duplicate values from arrays | No |

#### Operation: Limit

| Field Name | Type | Description | Required |
|---|---|---|---|
| Max Items | Number | Maximum number of items to keep | Yes |

#### Operation: Remove Duplicates

| Field Name | Type | Description | Required |
|---|---|---|---|
| Compare | Multi Options | Fields to compare for duplicates | Yes |
| Remove Other Fields | Boolean | Remove fields not used for comparison | No |

#### Operation: Sort

| Field Name | Type | Description | Required |
|---|---|---|---|
| Sort | Collection | Sort criteria and direction | Yes |
| Field Name | String | Field to sort by | Yes |
| Order | Options | Sort order (ascending/descending) | Yes |

#### Operation: Summarize

| Field Name | Type | Description | Required |
|---|---|---|---|
| Group By | Multi Options | Fields to group by | No |
| Values | Collection | Fields to aggregate and operations | Yes |
| Aggregation | Options | Aggregation function (sum, average, count, etc.) | Yes |
| Field | String | Field to aggregate | Yes |

## UseCases

- **Array Processing** : Convert array elements to individual items for parallel processing
- **Data Consolidation** : Combine multiple items into single collections for reporting
- **Performance Optimization** : Limit items to prevent processing too many records
- **Data Deduplication** : Remove duplicate records to ensure data quality
- **Data Ordering** : Sort items by specific criteria for prioritization
- **Data Aggregation** : Calculate sums, averages, and counts for analytics

### Split Out Items Operation

#### Use Cases
- **Array Processing**: Convert array elements to individual items
- **Object Flattening**: Extract object properties as separate items
- **Data Normalization**: Transform nested data into flat structure
- **Batch Processing**: Prepare data for parallel processing

#### Configuration Options
- **Field Selection**: Choose which fields contain arrays/objects
- **Key Preservation**: Maintain original item keys
- **Index Information**: Include array index in output

### Concatenate Items Operation

#### Use Cases
- **Data Consolidation**: Combine multiple items into single collection
- **Report Generation**: Gather data for summary reports
- **Batch Packaging**: Package items for bulk operations
- **List Creation**: Create lists from individual data points

#### Configuration Options
- **Field Mapping**: Specify how fields should be combined
- **Structure Definition**: Define output data structure
- **Duplicate Handling**: Handle overlapping field names

### Limit Operation

#### Use Cases
- **Performance Optimization**: Prevent processing too many items
- **Sampling**: Take a subset of data for testing
- **Rate Limiting**: Control downstream processing load
- **Top-N Selection**: Get first N items from ordered data

#### Configuration Options
- **Max Items**: Maximum number of items to keep
- **Skip Items**: Number of items to skip from beginning
- **Random Selection**: Randomly select items instead of first N

### Remove Duplicates Operation

#### Use Cases
- **Data Deduplication**: Remove duplicate records
- **Unique Values**: Ensure uniqueness in datasets
- **Clean Processing**: Prevent duplicate processing
- **Data Quality**: Improve data consistency

#### Configuration Options
- **Compare Fields**: Specify which fields to compare for duplicates
- **Comparison Mode**: Exact match vs. fuzzy matching
- **Keep Strategy**: Which duplicate to keep (first, last, etc.)

### Sort Operation

#### Use Cases
- **Data Ordering**: Sort items by specific criteria
- **Prioritization**: Order by importance or urgency
- **Alphabetical Listing**: Sort text data alphabetically
- **Numerical Ordering**: Sort by numeric values

#### Configuration Options
- **Sort Fields**: Specify fields to sort by
- **Sort Direction**: Ascending or descending order
- **Data Types**: Handle different data types appropriately
- **Multiple Criteria**: Multi-level sorting

### Summarize Operation

#### Use Cases
- **Data Aggregation**: Calculate sums, averages, counts
- **Pivot Tables**: Create cross-tabulated summaries
- **Reporting**: Generate summary statistics
- **Analytics**: Perform data analysis operations

#### Configuration Options
- **Group By**: Fields to group data by
- **Aggregations**: Mathematical operations (sum, avg, count, min, max)
- **Filters**: Pre-aggregation filtering
- **Output Format**: Structure of summary results

### Data Types Support

#### Supported Types
- **Arrays**: JavaScript arrays and nested arrays
- **Objects**: JSON objects and nested objects
- **Strings**: Text data with various operations
- **Numbers**: Numeric data with mathematical operations
- **Booleans**: True/false values
- **Dates**: Date and time values

#### Type Handling
- **Automatic Detection**: Automatically detect data types
- **Type Conversion**: Convert between compatible types
- **Type Validation**: Validate data types for operations
- **Error Handling**: Handle type mismatches gracefully

### Performance Considerations

#### Large Datasets
- **Memory Usage**: Efficient processing of large item collections
- **Streaming**: Process items in batches when possible
- **Optimization**: Optimized algorithms for common operations
- **Resource Management**: Manage memory and CPU usage

#### Best Practices
- **Batch Size**: Consider appropriate batch sizes
- **Operation Order**: Optimize operation sequence
- **Memory Limits**: Be aware of system memory constraints
- **Error Recovery**: Implement proper error handling

### Advanced Features

#### Expression Support
- **Dynamic Configuration**: Use expressions for dynamic behavior
- **Conditional Logic**: Apply operations conditionally
- **Field References**: Reference other fields in operations
- **Custom Functions**: Use JavaScript functions when needed

#### Error Handling
- **Graceful Degradation**: Continue processing on non-critical errors
- **Error Reporting**: Detailed error messages and context
- **Data Validation**: Validate input data before processing
- **Recovery Options**: Options for handling failed operations

### Usage Examples

#### Split Array to Items
```
Operation: Split Out Items
Input: [{"tags": ["red", "blue", "green"]}]
Output: [{"tag": "red"}, {"tag": "blue"}, {"tag": "green"}]
```

#### Combine Items to List
```
Operation: Concatenate Items
Input: [{"name": "John"}, {"name": "Jane"}]
Output: [{"names": ["John", "Jane"]}]
```

#### Limit Items
```
Operation: Limit
Max Items: 5
Input: 10 items
Output: First 5 items
```

#### Remove Duplicates
```
Operation: Remove Duplicates
Compare Fields: ["email"]
Input: Multiple items with same email
Output: Unique items by email
```

#### Sort Items
```
Operation: Sort
Sort Field: "score"
Direction: Descending
Input: Items with various scores
Output: Items sorted by score (high to low)
```

#### Summarize Data
```
Operation: Summarize
Group By: ["department"]
Aggregate: Count of items, Sum of salary
Input: Employee records
Output: Department summaries
```

### Integration Patterns

#### Data Pipeline
- **Pre-processing**: Prepare data for main processing
- **Transformation**: Convert data formats and structures
- **Post-processing**: Clean up and format results
- **Quality Control**: Ensure data quality and consistency

#### Workflow Optimization
- **Performance**: Optimize data flow for better performance
- **Scalability**: Handle varying data volumes
- **Reliability**: Ensure consistent processing results
- **Maintainability**: Create clear and understandable workflows

### Migration Notice

**⚠️ Hidden Node**: This node is marked as hidden, indicating it may be deprecated or have limited visibility. Functionality may be available in other nodes:
- **Split In Batches**: For splitting operations
- **Set**: For data transformation
- **Function**: For custom operations

### Usage Notes
- Node is marked as hidden in current version
- Supports complex array and object transformations
- Optimized for performance with large datasets
- Expression support for dynamic operations
- Error handling ensures robust data processing
- Multiple operations available for different use cases
- Maintains data integrity throughout transformations
- Compatible with all n8n data types and formats

## UseCases

### Data Processing and Transformation

1. **Array Flattening**: Convert nested arrays into individual items for parallel processing
2. **Data Aggregation**: Combine multiple data points into summary reports
3. **List Management**: Clean, sort, and organize data collections
4. **Batch Processing**: Prepare data batches for downstream operations
5. **Data Deduplication**: Remove duplicate entries from datasets

### Workflow Optimization

1. **Performance Control**: Limit processing load by restricting item counts
2. **Data Sampling**: Create test datasets by taking samples from larger collections
3. **Priority Sorting**: Reorder items by importance or business rules
4. **Quality Control**: Remove duplicates and inconsistencies from data flows
5. **Load Balancing**: Distribute work evenly across workflow branches

### Reporting and Analytics

1. **Summary Statistics**: Generate counts, sums, and averages from data collections
2. **Pivot Tables**: Create cross-tabulated reports from raw data
3. **Top-N Analysis**: Find top performers or most frequent items
4. **Data Grouping**: Organize data by categories for analysis
5. **Trend Analysis**: Sort and limit data for time-series analysis

### Integration Patterns

1. **API Response Processing**: Transform API responses into usable data structures
2. **Database Result Handling**: Process query results for application use
3. **File Processing**: Handle large file imports by limiting and organizing data
4. **Message Queue Management**: Control message flow and prevent overload
5. **ETL Operations**: Extract, transform, and load data between systems

### E-commerce and Business

1. **Product Catalog Management**: Sort and organize product listings
2. **Customer Data Processing**: Deduplicate and organize customer records
3. **Order Processing**: Summarize orders by various dimensions
4. **Inventory Management**: Track and organize stock information
5. **Sales Analytics**: Analyze sales data with grouping and summarization

### Content Management

1. **Content Aggregation**: Combine content from multiple sources
2. **Feed Processing**: Organize and limit RSS or social media feeds
3. **Search Results**: Sort and filter search results by relevance
4. **Tag Management**: Organize and deduplicate content tags
5. **Media Processing**: Handle collections of images, videos, or documents

