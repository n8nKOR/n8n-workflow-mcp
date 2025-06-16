# Simulate Node

## Overview

Simulate a node

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Simulation

#### Operation: Simulate
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Icon | options | Icon selector for the simulated node | No | - |
| Subtitle | string | Subtitle text for the simulated node | No | - |
| Output | options | How to handle output | Yes | - |
| Number of Items | number | Number of input items to return (when Output is "Specify how many") | No | - |
| JSON Output | string | Custom JSON output specification (when Output is "Specify output as JSON") | No | - |
| Execution Duration | number | Simulation execution time in milliseconds | No | - |

**Output Options:**
- **Returns all input items**: Pass through all input items
- **Specify how many of input items to return**: Return limited number of items
- **Specify output as JSON**: Return custom JSON output

## UseCases

- **Workflow Testing**: Test workflow logic without real operations
- **Development**: Create workflow prototypes with simulated nodes
- **Training**: Practice workflow development with safe simulation
- **Debugging**: Isolate workflow sections for troubleshooting
- **Performance Testing**: Test workflow execution with controlled delays

