# MQTT Node

## Overview

The MQTT node allows you to publish messages to MQTT brokers. MQTT (Message Queuing Telemetry Transport) is a lightweight, publish-subscribe messaging protocol designed for low-bandwidth, high-latency or unreliable networks. This node enables you to send data to MQTT topics as part of your n8n workflows, making it ideal for IoT applications, real-time messaging, and distributed systems communication.

## Credentials

This node requires MQTT broker credentials:
- **Protocol**: Connection protocol (mqtt://, mqtts://, ws://, wss://)
- **Host**: MQTT broker hostname or IP address
- **Port**: MQTT broker port (default: 1883 for mqtt, 8883 for mqtts)
- **Username**: Username for broker authentication (optional)
- **Password**: Password for broker authentication (optional)
- **Client ID**: Unique client identifier (optional, auto-generated if not provided)

To configure MQTT credentials:
1. Set up your MQTT broker (e.g., Mosquitto, HiveMQ, AWS IoT Core)
2. Note the broker's hostname/IP and port
3. Create authentication credentials if required by your broker
4. Configure SSL/TLS settings if using secure connections

## Inputs

- **Main**: JSON data containing the message payload and configuration

## Outputs

- **Main**: Returns confirmation of message publication with metadata

## Properties

### Resource: MQTT Publisher

#### Operation: Publish Message

| Field Name | Type | Description | Required |
|---|---|---|---|
| Topic | string | MQTT topic to publish to | Yes |
| Message | string | Message payload to publish | Yes |
| QoS | options | Quality of Service level (0, 1, 2) | No |
| Retain | boolean | Whether message should be retained by broker | No |
| Message Format | options | Format of the message (String, JSON) | No |
| Payload Encoding | options | Text encoding for message payload | No |
| Duplicate Flag | boolean | Mark message as duplicate | No |
| Message ID | number | Unique message identifier | No |

## Features

- **Multiple QoS Levels**: Support for all MQTT QoS levels (0, 1, 2)
- **Topic Flexibility**: Dynamic topic assignment using expressions
- **Message Formats**: Support for text, JSON, and binary payloads
- **Retained Messages**: Option to retain messages on the broker
- **SSL/TLS Support**: Secure connections to MQTT brokers
- **Authentication**: Username/password and certificate-based auth
- **Connection Persistence**: Efficient connection reuse across executions

## MQTT Concepts

### Topics
- **Topic Structure**: Hierarchical topics using forward slashes (e.g., `home/living-room/temperature`)
- **Wildcards**: Not supported for publishing (only for subscribing)
- **Topic Naming**: Use descriptive, hierarchical naming conventions

### Publishing Best Practices
- **Topic Design**: Use consistent, logical topic hierarchies
- **QoS Selection**: Choose appropriate QoS level for your use case
- **Message Size**: Keep messages reasonably small for better performance
- **Retained Messages**: Use sparingly for state information

### Connection Management
- **Keep Alive**: Automatically managed by the node
- **Clean Session**: Configurable per connection
- **Reconnection**: Automatic reconnection handling
- **Connection Pooling**: Efficient resource usage

## UseCases

### IoT Data Publishing
- **Sensor Data**: Publish sensor readings to MQTT topics
- **Device Status**: Send device health and status updates
- **Command Responses**: Publish acknowledgments and responses
- **Telemetry**: Stream real-time device telemetry data

### Home Automation
- **Smart Devices**: Control and monitor smart home devices
- **Automation Rules**: Trigger automation based on sensor data
- **Status Updates**: Publish device state changes
- **Integration**: Connect with Home Assistant, OpenHAB, etc.

### Industrial Applications
- **Machine Data**: Publish production line data
- **Alerts**: Send critical alerts and notifications
- **Monitoring**: Real-time equipment monitoring
- **SCADA Integration**: Interface with SCADA systems

### Real-time Messaging
- **Chat Applications**: Publish chat messages
- **Notifications**: Send push notifications
- **Live Updates**: Broadcast real-time updates
- **Event Streaming**: Stream application events

### Platform Integration
- **Home Assistant Integration**: Connect smart home sensors and devices to Home Assistant using MQTT topics
- **AWS IoT Core**: Publish device data to AWS IoT Core for cloud processing and analytics
- **Industrial Monitoring**: Send machine status and production data to industrial monitoring systems

## Error Handling

Common issues and solutions:
- **Connection Failed**: Verify broker address, port, and credentials
- **Topic Invalid**: Check topic naming conventions and permissions
- **QoS Not Supported**: Verify broker supports selected QoS level
- **Authentication Failed**: Confirm username/password or certificate setup
- **SSL/TLS Error**: Check certificate configuration and broker SSL settings

## Performance Considerations

### Message Throughput
- **QoS Impact**: Higher QoS levels have lower throughput
- **Message Size**: Smaller messages generally perform better
- **Connection Reuse**: Node reuses connections for efficiency
- **Batch Publishing**: Consider batching for high-volume scenarios

### Network Optimization
- **Keep Alive Interval**: Optimize for your network conditions
- **Connection Timeout**: Configure appropriate timeout values
- **Bandwidth Usage**: Monitor network usage with QoS 1 and 2
- **Broker Selection**: Choose geographically close brokers when possible

## Security

### Authentication Methods
- **Username/Password**: Basic authentication
- **Client Certificates**: X.509 certificate-based authentication
- **API Keys**: Some cloud brokers support API key authentication
- **OAuth**: Advanced authentication for cloud services

### Encryption
- **TLS/SSL**: Encrypt data in transit
- **Certificate Validation**: Verify broker certificates
- **Payload Encryption**: Encrypt sensitive payloads at application level
- **Topic-level Security**: Use broker ACLs for topic permissions


