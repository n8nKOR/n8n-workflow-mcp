# Merge Node

## Overview

⚠️ **Implementation Note**: This documentation describes the Merge node concept. The actual implementation logic is located in `/v3/MergeV3.node.ts` (version 3.2) which provides the core merging functionality.

Merges data of multiple streams once data from both is available

## Credentials

None

## Inputs

- Dynamic (configured based on mode selection)

## Outputs

- Main

## Properties

### Resource: Data Merger

#### Operation: Merge Data

| Field Name | Type | Description | Required |
|---|---|---|---|
| Mode | options | How input data should be merged (Append, Combine, SQL Query, Choose Branch) | Yes |
| Combine By | options | How input data should be merged (when Mode is Combine) | No |

### Merge Modes
- **Append**: Output items of each input, one after the other
- **Combine**: Merge matching items together
- **SQL Query**: Write a query to do the merge
- **Choose Branch**: Output data from a specific branch, without modifying it

### Combine Options (when Mode is Combine)
- **Matching Fields**: Combine items with the same field values
- **Position**: Combine items based on their order
- **All Possible Combinations**: Every pairing of every two items (cross join)

### Mode-Specific Behavior

#### Append Mode
- **Simple Concatenation**: Combines all input data in sequence
- **Preserves Order**: Maintains original item ordering from each input
- **No Data Modification**: Items pass through unchanged

#### Combine by Matching Fields
- **Field Selection**: Choose which fields to match on
- **Join Types**: Configure how non-matching items are handled
- **Conflict Resolution**: Handle duplicate field names

#### Combine by Position
- **Index-Based Merging**: Matches items by their position in inputs
- **Padding Options**: Handle inputs of different lengths
- **Order Preservation**: Maintains positional relationships

#### All Possible Combinations
- **Cross Join**: Creates cartesian product of all inputs
- **Output Multiplier**: Result count = input1_count × input2_count × ...
- **Memory Considerations**: Can generate large result sets

#### SQL Query Mode
- **Custom Logic**: Write SQL-like queries for complex merging
- **Field Mapping**: Access input data through virtual tables
- **Advanced Filtering**: Combine with WHERE clauses and functions

#### Choose Branch Mode
- **Branch Selection**: Pick specific input to output
- **Conditional Logic**: Dynamic branch selection based on data
- **Pass-through**: Selected data flows unchanged

### Dynamic Input Configuration
- **Required Inputs**: Varies by mode (1 for most, 2 for chooseBranch)
- **Input Labels**: Dynamically configured based on mode
- **Connection Types**: All inputs accept Main connection type

### Advanced Features
- **Version**: Uses Merge v3.2 (latest)
- **Expression Support**: Dynamic configuration through expressions
- **Memory Efficient**: Optimized for large data sets
- **Error Handling**: Robust handling of mismatched data structures

## UseCases

- **PDF Analysis Result Merging**: Merge PDF text extraction results with OpenAI analysis results to combine file information and AI analysis results for integrated processing.
- **Web Scraping Data Integration**: Merge title information extracted from web pages with AI summary results to generate complete content information.
- **Multi-Source Data Combination**: Match related information from multiple APIs or databases using specific key values to create unified datasets.
- **Workflow Branch Recombination**: Recombine data processing paths that were split by conditional branching to generate integrated data for subsequent processing.
- **Image and AI Analysis Result Combination**: Combine image information (size, resolution) with AI-generated caption data based on position to create complete datasets needed for caption overlay.
- **Visual Regression Test Data Integration**: Merge AI comparison analysis results of reference images and current images to generate integrated data for change detection reports.

