# Switch Node

## Overview

Route items depending on defined expression or rules

⚠️ **Version Information**: This node has multiple versions (V1, V2, V3). The current default version is V3.2, which provides enhanced routing capabilities, improved condition handling, and advanced rule-based routing options. Some features may differ between versions.

## Credentials

None

## Inputs

- Main

## Outputs

- Dynamic outputs based on mode (Expression or Rules)

## Properties

### Resource: Switch

#### Operation: Expression Mode

| Field Name | Type | Description | Required |
|---|---|---|---|
| Mode | Options | How data should be routed (Expression) | Yes |
| Number of Outputs | Number | How many outputs to create | Yes |
| Output Index | Number | The output index to send the input item to | Yes |

#### Operation: Rules Mode

| Field Name | Type | Description | Required |
|---|---|---|---|
| Mode | Options | How data should be routed (Rules) | Yes |
| Routing Rules | Fixed Collection | Define rules to route items to different outputs | Yes |
| Conditions | Filter | Build matching conditions for each rule | Yes |
| Rename Output | Boolean | Whether to rename the output label | No |
| Output Name | String | Custom label for the output when rule matches | No |
| Fallback Output | Options | Where to send items when no rules match | No |
| Ignore Case | Boolean | Whether to ignore letter case when evaluating conditions | No |
| Rename Fallback Output | String | Custom name for fallback output | No |
| Send data to all matching outputs | Boolean | Whether to send data to all outputs meeting conditions | No |
| Less strict type validation | Boolean | Use less strict validation for data types | No |

## UseCases

- **Data Classification** : Route items based on data type, category, or content
- **Error Handling** : Route items to different paths based on error conditions
- **Priority Routing** : Send high-priority items to expedited processing paths
- **Status-based Routing** : Route items based on their processing status
- **User Type Routing** : Route requests based on user roles or permissions
- **Approval Workflows** : Route items through different approval chains
- **Customer Segmentation** : Route customers based on segment criteria
- **Product Categorization** : Route products to appropriate processing workflows
- **Order Processing** : Route orders based on type, value, or complexity
- **Lead Qualification** : Route leads based on qualification scores
- **Content Filtering** : Filter and route content based on specific criteria
- **Format Handling** : Route data to different processors based on format
- **Size-based Processing** : Route large vs small datasets to appropriate handlers
- **Quality Control** : Route items based on quality assessment results
- **Duplicate Detection** : Route duplicates to separate processing paths
- **API Response Handling** : Route responses based on status codes or content
- **Multi-system Integration** : Route data to different systems based on criteria
- **Protocol Selection** : Choose communication protocols based on requirements
- **Service Selection** : Route requests to different service instances
- **Version Handling** : Route requests to different API versions

