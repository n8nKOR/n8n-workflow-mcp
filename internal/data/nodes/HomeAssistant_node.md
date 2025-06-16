# Home Assistant Node

## Overview

Consume Home Assistant API

## Credentials

- Name: homeAssistantApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: cameraProxy

#### Operation: getScreenshot

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Camera Entity Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Put Output File in Field | string |  | Yes |  |

### Resource: event

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event Type | string | The Entity ID for which an event will be created | Yes |  |
| Event Attributes | fixedCollection | Name of the attribute | No |  |

### Resource: history

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | The end of the period | No |  |

### Resource: log

#### Operation: getLogbookEntries

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | The end of the period | No |  |

### Resource: service

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: call

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Domain Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Service Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Service Attributes | fixedCollection | Name of the field | No |  |

### Resource: state

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Entity Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: upsert

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Entity Name or ID | options | The entity ID for which a state will be created. Choose from the list, or specify an ID using an <a href= | Yes |  |
| State | string |  | Yes |  |
| State Attributes | fixedCollection | Name of the attribute | No |  |

### Resource: template

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Template | string | Render a Home Assistant template. <a href= | Yes |  |

## UseCases

- **Smart Home Automation**: Automate smart home devices, lighting, climate control, and security systems through Home Assistant integration
- **Device State Management**: Monitor and control IoT device states, sensors, and smart home equipment across different protocols
- **Home Security Monitoring**: Integrate security cameras, door sensors, and alarm systems for comprehensive home security automation
- **Energy Management**: Monitor energy consumption, solar panels, and smart appliances for efficient home energy management
- **Climate Control**: Automate heating, cooling, and ventilation systems based on occupancy, weather, and energy efficiency
- **Notification Systems**: Create custom home automation notifications for security events, system status, and device alerts
- **Voice and Mobile Integration**: Connect home automation with voice assistants and mobile apps for seamless control
- **Data Analytics**: Collect and analyze home automation data for optimization, energy savings, and usage patterns
- **Event-driven Automation**: Create complex automation scenarios based on time, occupancy, weather, and device interactions
- **Remote Home Management**: Monitor and control home systems remotely for vacation homes, rental properties, or while traveling

