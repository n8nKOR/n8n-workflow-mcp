# MongoDB Node

## Overview

Find, insert, update, and aggregate documents in MongoDB databases. This node provides comprehensive operations for interacting with MongoDB collections, supporting complex queries, data manipulation, and aggregation pipelines.

## Credentials

- Name: mongoDb, Required: Yes

The MongoDB credentials support two configuration types:
- **Connection String**: Direct MongoDB connection string
- **Parametric Configuration**: Individual server details (host, port, database, authentication)

## Inputs

- Main: Input data for insert/update operations

## Outputs

- Main: Query results, operation status, or error information

## Properties

### Resource: Database Operations

#### Operation: Document Management

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Database operation to perform | Yes |
| Collection | String | MongoDB collection name | Yes |
| Query | JSON | Query filter or aggregation pipeline | Conditional |
| Update Key | String | Field for document identification | Conditional |
| Fields | String | Fields to include/update | No |
| Upsert | Boolean | Create if not exists | No |
| Options | Collection | Additional operation options | No |
## Operations

### Find Documents

Search for documents in a MongoDB collection with advanced filtering and options.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Query (JSON Format) | JSON | MongoDB query filter | Yes | {} |

**Query Options:**
- **Limit**: Maximum number of documents to return (0 = unlimited)
- **Skip**: Number of documents to skip for pagination
- **Sort (JSON Format)**: Sort order specification `{ "field": 1 }` (1=asc, -1=desc)
- **Projection (JSON Format)**: Field selection `{ "_id": 0, "field": 1 }`

**Example Query:**
```json
{ "age": { "$gte": 18 }, "status": "active" }
```

### Insert Documents

Insert new documents into a MongoDB collection.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Fields | String | Comma-separated list of fields to include | No | (all fields) |

**Options:**
- **Date Fields**: Fields to parse as MongoDB Date type
- **Use Dot Notation**: Enable dot notation for nested field access

### Update Documents

Update existing documents in a MongoDB collection.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Key | String | Field used to identify documents (typically "_id") | Yes | id |
| Fields | String | Comma-separated list of fields to update | No | (all fields) |
| Upsert | Boolean | Create document if it doesn't exist | No | false |

**Options:**
- **Date Fields**: Fields to parse as MongoDB Date type
- **Use Dot Notation**: Enable dot notation for nested field updates

### Find and Update

Find documents and update specific fields using MongoDB's `$set` operator.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Key | String | Field used to identify documents | Yes | id |
| Fields | String | Fields to update | No | (all fields) |
| Upsert | Boolean | Create document if none found | No | false |

### Find and Replace

Find documents and replace them entirely with new data.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Update Key | String | Field used to identify documents | Yes | id |
| Fields | String | Fields for the replacement document | No | (all fields) |
| Upsert | Boolean | Create document if none found | No | false |

### Delete Documents

Remove documents matching the specified criteria.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Delete Query (JSON Format) | JSON | Filter criteria for deletion | Yes | {} |

**Returns:** `{ "deletedCount": number }`

### Aggregate Documents

Execute MongoDB aggregation pipelines for complex data processing.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Query | JSON | Aggregation pipeline array | Yes | [] |

**Example Pipeline:**
```json
[
  { "$match": { "status": "active" } },
  { "$group": { "_id": "$category", "count": { "$sum": 1 } } },
  { "$sort": { "count": -1 } }
]
```

## Advanced Features

### ObjectId Handling
- Automatically converts string `_id` values to MongoDB ObjectId
- Supports ObjectId queries and updates
- Returns ObjectIds as strings in results

### Date Field Processing
- Parse specified fields as MongoDB Date objects
- Support for ISO date strings and timestamps
- Configurable through `dateFields` option

### Dot Notation Support
- Access nested document fields using dot notation
- Enable via `useDotNotation` option
- Example: `"user.profile.name"`

### Error Handling
- Continue on Fail support for batch operations
- Detailed error messages with operation context
- Per-item error handling for multi-document operations

### Connection Management
- Automatic connection pooling
- Connection string or parametric configuration
- Database validation and error reporting

## Query Examples

### Basic Find
```json
{ "status": "published", "category": "technology" }
```

### Complex Query with Operators
```json
{
  "price": { "$gte": 100, "$lte": 500 },
  "tags": { "$in": ["featured", "popular"] },
  "createdAt": { "$gte": "2023-01-01T00:00:00Z" }
}
```

### Projection Example
```json
{ "_id": 0, "title": 1, "price": 1, "category": 1 }
```

### Sort Example
```json
{ "createdAt": -1, "price": 1 }
```

### Aggregation Pipeline
```json
[
  { "$match": { "status": "active" } },
  { "$lookup": {
      "from": "categories",
      "localField": "categoryId",
      "foreignField": "_id",
      "as": "category"
  }},
  { "$unwind": "$category" },
  { "$project": {
      "title": 1,
      "categoryName": "$category.name",
      "price": 1
  }}
]
```

## UseCases

### Data Operations
- **Document Retrieval**: Query collections with complex filters and sorting
- **Bulk Data Import**: Insert large datasets efficiently
- **Data Synchronization**: Update documents based on external data sources
- **Content Management**: Manage articles, products, or user-generated content

### Analytics and Reporting
- **Aggregation Pipelines**: Generate complex reports and analytics
- **Data Transformation**: Process and reshape data for analysis
- **Statistical Analysis**: Calculate metrics, averages, and summaries
- **Data Migration**: Transform data during database migrations

### Application Integration
- **API Backends**: Power REST APIs with MongoDB operations
- **Microservices**: Integrate MongoDB with service architectures
- **Event Processing**: Handle events and update document states
- **Real-time Applications**: Support live data updates and queries

### Business Logic
- **Inventory Management**: Track product quantities and availability
- **User Management**: Handle user profiles and authentication data
- **Order Processing**: Manage e-commerce orders and transactions
- **Audit Trails**: Maintain logs of data changes and user actions

### Advanced Use Cases
- **Search Implementation**: Build text search with MongoDB queries
- **Data Validation**: Enforce business rules through query logic
- **Caching Layer**: Use MongoDB as a caching solution
- **Session Management**: Store and manage user session data

