# Transform Family Nodes

## Overview

⚠️ **CRITICAL IMPORTANT**: The "Transform" nodes in n8n are **NOT a single node** with multiple operations. Each Transform functionality is a **completely separate, independent node** with its own unique name and purpose. Do not look for a "Transform" node in n8n - instead, look for specific nodes like "Aggregate", "Limit", "Sort", etc.

**Node Architecture**: Multiple independent nodes, not a single unified node
**Naming Convention**: Each node has a specific name (e.g., "Aggregate", "Limit", "Sort")
**Workflow Usage**: Add each node individually to your workflow as separate components
**Important**: There is NO single "Transform" node - this document covers multiple separate nodes

**Available Independent Transform Nodes:**
- **Aggregate** - Combines and summarizes data (Node name: "Aggregate")
- **Limit** - Restricts the number of items (Node name: "Limit")  
- **Sort** - Orders workflow items (Node name: "Sort")
- **Remove Duplicates** - Filters out duplicate items (Node name: "Remove Duplicates")
- **Split Out** - Splits arrays into individual items (Node name: "Split Out")
- **Summarize** - Groups data and performs aggregation (Node name: "Summarize")

## ⚠️ User Guidance - How to Find These Nodes

### In n8n Node Palette:
1. **Search for SPECIFIC node names**, not "Transform"
2. Each node appears as a separate entry (e.g., search "Aggregate", not "Transform")
3. Add each node individually to your workflow
4. Configure each node's specific properties independently

### Common Mistakes to Avoid:
- ❌ Don't search for "Transform" in node palette
- ❌ Don't expect a single node with Transform operations
- ❌ Don't look for Transform resource with multiple operations
- ✅ Search for specific node names (Aggregate, Limit, Sort, etc.)
- ✅ Treat each as completely independent nodes
- ✅ Configure each node separately in your workflow

## How to Use Transform Nodes

### Adding to Workflow
1. Search for the specific node name (e.g., "Aggregate", not "Transform")
2. Each node appears as a separate item in the node palette
3. Add each node individually to your workflow
4. Configure each node's specific properties

### Node Independence
- Each Transform node is completely independent
- No shared configuration or resources
- Each has its own specific properties and behavior
- Connect them individually in your workflow chain

## Individual Node Documentation

### Aggregate Node
**Node Name**: `Aggregate`

Combines and summarizes data from multiple items into consolidated results.

**Key Properties:**
- **Aggregate**: Fields to aggregate data
- **Include All Fields**: Include all input fields in output

### Limit Node
**Node Name**: `Limit`

Restricts the number of items that pass through the workflow.

**Key Properties:**
- **Max Items**: Maximum number of items to pass through

### Sort Node
**Node Name**: `Sort`

Orders workflow items based on specified fields and criteria.

**Key Properties:**
- **Type**: Field types for sorting
- **Sort**: Fields and order for sorting

### Remove Duplicates Node
**Node Name**: `Remove Duplicates`

Filters out duplicate items based on specified comparison fields.

**Key Properties:**
- **Compare**: Fields to compare for duplicates
- **Include Binary Data**: Include binary data in duplicate check

### Split Out Node
**Node Name**: `Split Out`

Splits arrays or objects into individual workflow items.

**Key Properties:**
- **Field to Split Out**: Field containing array/object to split
- **Include**: Which data to include in output

### Summarize Node
**Node Name**: `Summarize`

Groups data and performs aggregation operations on grouped items.

**Key Properties:**
- **Aggregate**: Fields and operations for summarization
- **Group By**: Fields to group data by

## General Information

### Credentials
Transform nodes do not require any credentials as they operate on data within the n8n workflow environment.

### Inputs
Transform nodes accept standard workflow data from previous nodes in the workflow.

### Outputs
Transform nodes provide transformed data output to subsequent nodes in the workflow.

## UseCases

- **Data Aggregation**: Combine multiple API responses into consolidated reports
- **List Management**: Remove duplicates from customer lists and contact databases
- **Data Sampling**: Limit large datasets for testing and development workflows
- **Report Generation**: Summarize sales data, user analytics, and performance metrics
- **Data Cleaning**: Standardize and clean imported data from various sources
- **Array Processing**: Convert nested arrays into individual workflow items
- **Sort Operations**: Order data by date, priority, or custom business logic
- **Batch Processing**: Group items by category, date, or other criteria for bulk operations
- **Data Transformation**: Convert data formats and structures using AI-powered transformations
- **Quality Control**: Filter and validate data quality before downstream processing
- **Performance Optimization**: Optimize workflow performance by limiting and organizing data flow
- **ETL Operations**: Extract, transform, and load data between different systems
- **Analytics Preparation**: Prepare data for analytics and business intelligence tools
- **Workflow Testing**: Create sample datasets for workflow development and testing
- **Data Migration**: Transform data structures during system migrations
- **API Response Processing**: Process and reshape API responses for specific use cases
- **Database Operations**: Prepare data for database insertions and updates
- **Export Formatting**: Format data for CSV, JSON, or other export formats
- **Real-time Processing**: Transform streaming data in real-time workflows
- **Compliance Processing**: Ensure data meets compliance and regulatory requirements