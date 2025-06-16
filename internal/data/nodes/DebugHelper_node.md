# DebugHelper Node

## Overview

Causes problems intentionally and generates useful data for debugging

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: debug

#### Operation: doNothing

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Category | options | What action to perform | Yes | doNothing |

#### Operation: throwError

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Category | options | What action to perform | Yes | throwError |
| Error Type | options | Type of error to throw (NodeApiError, NodeOperationError, Error) | Yes | NodeApiError |
| Error Message | string | The message to send as part of the error | Yes | Node has thrown an error |

#### Operation: oom

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Category | options | What action to perform | Yes | oom |
| Memory Size to Generate | number | The approximate amount of memory to generate | Yes | 10 |

#### Operation: randomData

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Category | options | What action to perform | Yes | randomData |
| Data Type | options | Type of random data to generate | Yes | user |
| Number of Items | number | Number of items to generate | No | 1 |
| Seed | number | Seed for random generation (for reproducible results) | No |  |

##### Data Type Options:
- Address: Generate random addresses
- Coordinates: Generate latitude/longitude coordinates  
- Credit Card: Generate credit card information
- Email: Generate email addresses
- IPv4: Generate IPv4 addresses
- IPv6: Generate IPv6 addresses
- MAC: Generate MAC addresses
- NanoIds: Generate NanoIDs
- URL: Generate URLs
- User Data: Generate user information
- UUID: Generate UUIDs
- Version: Generate semantic version numbers

##### Additional Fields for NanoId:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| NanoId Alphabet | string | The alphabet to use for generating the nanoIds | No | 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ |
| NanoId Length | number | The length of the nanoIds to generate | No | 21 |

## UseCases

- **Error Handling Testing**: Test workflow error handling by intentionally triggering various error types to ensure robust error recovery
- **Load Testing**: Generate memory pressure to test workflow performance under resource constraints and identify optimization opportunities
- **Test Data Creation**: Generate realistic test data for development environments, automated testing, and demo scenarios
- **Workflow Debugging**: Use as a controlled failure point to test workflow branching and error recovery mechanisms
- **Performance Benchmarking**: Create consistent test scenarios to measure workflow performance and identify bottlenecks
- **Development Prototyping**: Generate sample data during development to test data transformations and processing logic
- **Quality Assurance**: Ensure workflows handle edge cases and error conditions properly before production deployment

