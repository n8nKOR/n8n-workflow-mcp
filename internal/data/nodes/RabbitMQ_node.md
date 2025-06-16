# RabbitMQ

## Overview

The RabbitMQ node provides integration with RabbitMQ, a robust message broker that implements the Advanced Message Queuing Protocol (AMQP). This node enables sending messages to RabbitMQ queues and exchanges, making it ideal for building distributed systems, microservices communication, and event-driven architectures.

## Credentials

- **Credential Name**: `rabbitmq`
- **Required**: Yes
- **Configuration**: 
  - Hostname: RabbitMQ server hostname
  - Port: RabbitMQ server port (default: 5672)
  - Username: Authentication username
  - Password: Authentication password
  - Virtual Host: RabbitMQ virtual host (optional)

## Inputs

- **Main**: Input data from previous nodes containing messages to send to RabbitMQ

## Outputs

- **Main**: Returns confirmation data about sent messages or deleted queue items

## Properties

### Resource: Message Queue

#### Operation: Message Operations

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Choose between 'Send a Message to RabbitMQ' or 'Delete From Queue' | Yes |
| Mode | Options | Select 'Queue' or 'Exchange' for message destination | No |
| Queue / Topic | String | Name of the queue to publish to (for queue mode) | No |
| Exchange | String | Name of the exchange to publish to (for exchange mode) | No |
| Type | Options | Exchange type: Direct, Topic, Headers, or Fanout | No |
| Routing Key | String | Routing key for message routing | No |
| Send Input Data | Boolean | Whether to send the input data as message content | No |
| Message | String | Custom message content to send | No |
| Headers | JSON | Message headers as JSON object | No |
| Options | Collection | Additional message options and configurations | No |

## UseCases

- **Microservices Communication** : Enable reliable communication between microservices
- **Event-Driven Architecture** : Publish and consume events across distributed systems
- **Task Queue Management** : Distribute tasks across multiple worker processes
- **System Integration** : Connect different systems through message queuing
- **Asynchronous Processing** : Decouple application components for better scalability

### Message Queue Patterns

1. **Work Queue**: Distribute tasks among multiple workers for load balancing
2. **Publish/Subscribe**: Broadcast messages to multiple subscribers
3. **Request/Reply**: Implement RPC-style communication between services
4. **Topic-based Routing**: Route messages based on content topics
5. **Priority Queues**: Handle high-priority messages first

### Enterprise Integration

1. **Microservices Communication**: Enable reliable communication between microservices
2. **Event-Driven Architecture**: Build event-driven systems with guaranteed delivery
3. **System Decoupling**: Decouple system components for better scalability
4. **Async Processing**: Handle time-consuming tasks asynchronously
5. **Data Pipeline**: Build reliable data processing pipelines

### Real-time Applications

1. **Live Notifications**: Send real-time notifications to users
2. **Chat Systems**: Build scalable chat and messaging applications
3. **IoT Data Processing**: Handle IoT device data streams
4. **Financial Trading**: Process trading orders and market data
5. **Monitoring Systems**: Send alerts and monitoring data

### Workflow Orchestration

1. **Task Scheduling**: Schedule and distribute background tasks
2. **Batch Processing**: Coordinate batch job processing
3. **ETL Operations**: Orchestrate data extraction and transformation
4. **Approval Workflows**: Implement multi-step approval processes
5. **File Processing**: Handle file upload and processing workflows

### High Availability Scenarios

1. **Load Balancing**: Distribute workload across multiple consumers
2. **Failover Support**: Implement failover with message persistence
3. **Dead Letter Queues**: Handle failed message processing
4. **Message Durability**: Ensure messages survive system failures
5. **Cluster Support**: Scale across multiple RabbitMQ nodes

### Business Applications

1. **E-commerce**: Process orders, payments, and inventory updates
2. **Customer Service**: Route support tickets and updates
3. **Marketing Automation**: Send targeted marketing messages
4. **Supply Chain**: Coordinate logistics and inventory management
5. **Content Management**: Handle content publishing and updates

