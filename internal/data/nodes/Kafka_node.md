# Kafka Node

## Overview

Sends messages to a Kafka topic

## Credentials

- Name: kafka, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Kafka Producer

#### Operation: Send Message

| Field Name | Type | Description | Required |
|---|---|---|---|
| Topic | string | Name of the queue or topic to publish to | Yes |
| Send Input Data | boolean | Whether to send the data the node receives as JSON to Kafka | No |
| Message | string | The message to be sent (when not sending input data) | No |
| JSON Parameters | boolean | Whether to use JSON format for parameters | No |
| Use Schema Registry | boolean | Whether to use Confluent Schema Registry | No |
| Schema Registry URL | string | URL of the schema registry (when using schema registry) | Yes |
| Event Name | string | Namespace and Name of Schema (namespace.name) | Yes |
| Use Key | boolean | Whether to use a message key | No |
| Key | string | The message key (when using key) | Yes |
| Headers | fixedCollection | Headers to include with the message | No |
| Headers (JSON) | json | Header parameters as JSON (flat object) | No |
| Acks | boolean | Whether producer must wait for acknowledgement from all replicas | No |
| Compression | boolean | Whether to send data in compressed format using GZIP codec | No |
| Timeout | number | The time to await a response in milliseconds | No |

### Message Sending Modes

#### Send Input Data
When enabled (default), the node sends the complete input data as JSON:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 30
}
```

#### Custom Message
When disabled, send a custom message instead of input data:
```
Message: "Custom notification message"
```

### Schema Registry Integration

#### Confluent Schema Registry
- **Schema Management**: Centralized schema management
- **Schema Evolution**: Handle schema changes over time
- **Serialization**: Automatic serialization using registered schemas

#### Event Name Format
```
Namespace.SchemaName
Example: user.profile.updated
```

### Message Key Usage

#### Partitioning
- **Key-based Partitioning**: Messages with same key go to same partition
- **Load Distribution**: Distribute messages across partitions
- **Order Guarantee**: Messages with same key maintain order

### Headers

Headers provide metadata about the message:
```json
{
  "source": "user-service",
  "version": "1.0",
  "timestamp": "2023-12-01T10:00:00Z"
}
```

### Compression

#### GZIP Compression
- **Reduced Bandwidth**: Smaller message sizes
- **Network Efficiency**: Better network utilization
- **CPU Trade-off**: Compression overhead vs network savings

### Acknowledgment Modes

#### Acks = 0 (Fire and Forget)
- **No Acknowledgment**: Producer doesn't wait for confirmation
- **High Throughput**: Fastest message sending
- **Risk**: Potential message loss

#### Acks = 1 (Leader Acknowledgment)
- **Leader Confirmation**: Wait for leader replica acknowledgment
- **Balanced**: Good balance of performance and reliability

## UseCases

#### Event Streaming
```
Topic: user-events
Message: {"userId": "12345", "action": "login"}
Key: "user_12345"
Headers: {"source": "auth-service"}
```

#### Log Aggregation
```
Topic: application-logs
Message: {"level": "ERROR", "service": "api", "message": "Database connection failed"}
Key: "api-service"
```

### Best Practices

- **Schema Evolution**: Design schemas for backward compatibility
- **Message Size**: Keep messages reasonably sized
- **Key Strategy**: Choose partition keys carefully
- **Error Handling**: Implement retry logic and monitoring

### Usage Notes
- Supports both input data forwarding and custom message sending
- Schema Registry integration for structured data serialization
- Flexible header configuration with collection or JSON format
- Message key support for partition control and ordering
- Compression option for network efficiency
- Configurable acknowledgment and timeout settings

