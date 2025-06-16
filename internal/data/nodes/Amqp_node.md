# Amqp

## Overview

The AMQP Sender node provides integration with AMQP 1.0 (Advanced Message Queuing Protocol) messaging systems. This node enables sending raw messages to AMQP queues or topics, making it ideal for enterprise messaging, event-driven architectures, and system integration workflows.

## Credentials

- **Credential Name**: `amqp`
- **Required**: Yes
- **Configuration**: 
  - Hostname: AMQP server hostname
  - Port: AMQP server port
  - Username: Authentication username (optional)
  - Password: Authentication password (optional)
  - Transport Type: Connection transport type (optional)

## Inputs

- **Main**: Input data from previous nodes containing messages to send via AMQP

## Outputs

- **Main**: Returns confirmation data about sent messages

## Properties

### Resource: Message Queue

#### Operation: Send Message

| Field Name | Type | Description | Required |
|---|---|---|---|
| Queue / Topic | String | Name of the queue or topic to publish to (e.g. topic://sourcename.something) | Yes |
| Headers | JSON | Header parameters as JSON (flat object), sent as application_properties in AMQP message meta info | No |
| Container ID | String | Container ID to pass to the RHEA backend | No |
| Data as Object | Boolean | Whether to send the data as an object | No |
| Reconnect | Boolean | Whether to automatically reconnect if disconnected | No |
| Reconnect Limit | Number | Maximum number of reconnect attempts | No |
| Send Property | String | The only property to send (if empty, the whole item will be sent) | No |

## UseCases

- **Cross-System Communication** : Send messages between enterprise applications using AMQP protocol
- **Legacy System Integration** : Connect modern workflows with legacy messaging infrastructure
- **Business Process Automation** : Trigger business processes through message queues
- **Event Broadcasting** : Publish events to multiple subscribers in enterprise environments
- **Service Bus Integration** : Connect with Azure Service Bus and other AMQP-compatible platforms
- **Event Publishing** : Publish domain events to message queues for asynchronous processing
- **Saga Pattern Implementation** : Coordinate distributed transactions through message flows
- **CQRS Support** : Send commands and events in CQRS architectural patterns
- **Microservices Orchestration** : Coordinate complex workflows across multiple services
- **Reactive Systems** : Build reactive applications with event-driven messaging
- **Message Translation** : Convert data formats between different systems via message queues
- **Content-Based Routing** : Route messages to different queues based on content
- **Message Aggregation** : Collect and combine messages from multiple sources
- **Batch Processing** : Send batch data to queues for bulk processing systems
- **API Gateway Integration** : Forward API requests to backend services via message queues
- **Stream Processing** : Send data streams to message queues for real-time analytics
- **IoT Data Ingestion** : Collect and queue IoT sensor data for processing
- **Log Aggregation** : Centralize log data through message queues
- **Monitoring Events** : Send system monitoring events for alerting and analytics
- **Audit Trail Management** : Queue audit events for compliance and security tracking
- **Order Processing** : Queue customer orders for fulfillment systems
- **Notification Services** : Send notifications through message queues to delivery services
- **File Processing** : Queue file processing tasks for background workers
- **Data Synchronization** : Synchronize data between different databases via message queues
- **Workflow Coordination** : Coordinate multi-step workflows through message exchanges

