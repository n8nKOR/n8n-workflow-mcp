# Unleashed Software Node

## Overview

Consume Unleashed Software API

## Credentials

- Name: unleashedSoftwareApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: salesOrder

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Only returns orders for a specified Customer GUID. The CustomerId can be specified as a list of comma-separated GUIDs. | No |  |

### Resource: stockOnHand

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Product ID | string |  | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Returns the stock on hand for a specific date | No |  |

## UseCases

- **Inventory Management** : Manage stock levels, product information, and inventory tracking
- **Sales Order Processing** : Process sales orders and manage customer transactions
- **Purchase Order Automation** : Automate purchase orders and supplier management
- **Warehouse Operations** : Coordinate warehouse operations and stock movements
- **Customer Management** : Manage customer information and sales relationships
- **Supplier Integration** : Integrate with suppliers for automated procurement processes
- **Product Catalog Management** : Maintain product catalogs and pricing information
- **Stock Level Monitoring** : Monitor stock levels and trigger reorder alerts
- **Sales Reporting** : Generate sales reports and performance analytics
- **Multi-Location Inventory** : Manage inventory across multiple warehouses and locations

