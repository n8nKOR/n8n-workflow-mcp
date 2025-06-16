# FileMaker Node

## Overview

Retrieve data from the FileMaker data API

## Credentials

- Name: fileMaker, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Database Operations

#### Operation: Database Management

| Field Name | Type | Description | Required |
|---|---|---|---|
| Action | Options | Database action to perform | Yes |
| Layout Name or ID | Options | FileMaker layout to use | Yes |
| Record ID | Number | Internal record ID for operations | Conditional |
| Offset | Number | Record number for pagination | No |
| Limit | Number | Maximum number of results | No |
| Get Portals | Boolean | Include portal data in results | No |
| Portals Name or ID | Options | Specific portals to retrieve | No |
| Response Layout Name or ID | Options | Layout for response formatting | No |
## UseCases

- **Database Management**: Perform CRUD operations on FileMaker databases
- **Data Integration**: Integrate FileMaker data with other systems
- **Automated Reports**: Generate reports from FileMaker layouts
- **Script Execution**: Run FileMaker scripts programmatically
- **Data Synchronization**: Keep FileMaker data synchronized with external sources

### Usage Notes
- Uses FileMaker Data API (REST API) for communication
- Requires proper FileMaker Server setup with Data API enabled
- Record IDs are internal FileMaker identifiers, different from field values
- Portal data retrieval can include related table information
- Scripts can be executed with optional parameters
- Offset and limit parameters support pagination for large datasets
- Layout selection determines available fields and portal options

