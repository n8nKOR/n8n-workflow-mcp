# Customer Datastore (n8n Training) Node

## Overview

The Customer Datastore node is a simplified training tool designed for n8n educational purposes and skill development. This node provides sample customer data for learning and practice workflows. **Note**: This is a basic training node that uses hardcoded sample data and only supports simple read operations (Get One Person and Get All People).

## Credentials

**None Required** - This is a training node that uses simulated data and does not require external API credentials or authentication.

## Inputs

- **Main**: Input data for customer lookup operations

## Outputs

- **Main**: Sample customer data from the training dataset

## Properties

### Resource: Customer Management

#### Operation: Get One Person

Retrieve a single customer record from the sample dataset.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Customer ID | String | ID of the customer to retrieve | Yes | - |

**Response Structure:**
```json
{
  "id": "cust_001",
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "+1-555-0123",
  "company": "Acme Corp",
  "status": "active"
}
```

#### Operation: Get All People

Retrieve all customer records from the sample dataset.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Return all sample records | No | true |
| Limit | Number | Maximum number of results to return | No | 10 |

**Response Structure:**
```json
{
  "customers": [
    {
      "id": "cust_001",
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "+1-555-0123",
      "company": "Acme Corp",
      "status": "active"
    },
    {
      "id": "cust_002", 
      "name": "Jane Smith",
      "email": "jane.smith@example.com",
      "phone": "+1-555-0124",
      "company": "Tech Corp",
      "status": "active"
    }
  ],
  "total_count": 10
}
```

## Limitations

⚠️ **Important Training Node Limitations:**
- **Read-Only Operations**: Only supports data retrieval operations
- **No CRUD Operations**: Does not support create, update, or delete operations
- **Hardcoded Data**: Uses static sample data, not a real database
- **No Validation**: No data validation or business logic implementation
- **No Bulk Operations**: No batch processing capabilities
- **Training Purpose Only**: Not suitable for production use

## UseCases

- **Training Scenarios**: Practice n8n workflow development with customer data
- **Learning Exercises**: Educational workflow demonstrations using sample customer records
- **Prototype Development**: Create workflow prototypes before implementing real data sources
- **Data Flow Testing**: Test data transformation and processing logic
- **Workflow Logic Development**: Develop and test conditional logic and data routing
- **API Integration Practice**: Learn API integration patterns with safe sample data
- **Error Handling Practice**: Practice error handling without affecting real systems
- **Template Development**: Create reusable workflow templates for customer data operations

