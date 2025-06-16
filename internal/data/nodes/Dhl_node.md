# DHL Node

## Overview

Consume DHL API

## Credentials

- Name: dhlApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: shipment

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Tracking Number | string | The tracking number to get details for | Yes |  |
| Options | collection | Additional options for tracking | No |  |

##### Options Collection Fields:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Recipient's Postal Code | string | DHL will return more detailed information when you provide the Recipient's Postal Code - it acts as a verification step | No |  |

##### API Features:
- Real-time shipment tracking
- Detailed delivery status information
- Location tracking throughout delivery process
- Estimated delivery time
- Proof of delivery when available
- Enhanced details with recipient postal code verification

## Features

- **Real-time Tracking**: Get current shipment status and location information
- **Delivery Notifications**: Monitor shipment progress throughout the delivery process  
- **Enhanced Verification**: Use recipient postal code for additional tracking details
- **Global Coverage**: Track international and domestic DHL shipments worldwide
- **Status Updates**: Receive detailed status information including estimated delivery times

## Tracking Information

### Shipment Status Types
- **Pre-Transit**: Shipment information received, package not yet picked up
- **In Transit**: Package is moving through the DHL network
- **Out for Delivery**: Package is on the delivery vehicle for final delivery
- **Delivered**: Package has been successfully delivered
- **Exception**: Delivery attempt failed or package requires special handling

### Available Data
- **Current Location**: Real-time package location information
- **Transit History**: Complete journey history with timestamps
- **Estimated Delivery**: Predicted delivery date and time window
- **Delivery Proof**: Signature and delivery confirmation when available
- **Service Details**: Shipping service type and features used

## Authentication

DHL API requires authentication credentials:
- **API Key**: Your DHL API key for tracking services
- **Account Access**: Valid DHL business account with API access

To obtain API credentials:
1. Register for a DHL business account
2. Apply for API access through DHL Developer Portal
3. Generate API keys for tracking services
4. Configure rate limits and usage permissions

## UseCases

- **E-commerce Order Tracking**: Automatically track customer orders and provide real-time delivery updates through your platform
- **Supply Chain Monitoring**: Monitor shipments across your supply chain to identify delays and optimize logistics operations
- **Customer Service Automation**: Automatically update customers about delivery status and handle delivery-related inquiries
- **Inventory Management**: Track incoming shipments to coordinate receiving operations and inventory updates
- **Delivery Exception Handling**: Automatically detect and respond to delivery exceptions or failed delivery attempts
- **Performance Analytics**: Analyze shipping performance metrics and delivery times for business optimization
- **Multi-Carrier Integration**: Integrate DHL tracking with other shipping providers for unified shipment visibility

