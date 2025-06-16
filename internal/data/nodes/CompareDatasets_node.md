# Compare Datasets

## Overview

The Compare Datasets node compares two input datasets to identify differences, similarities, and unique items. This node is useful for data synchronization, change detection, and data quality analysis by comparing datasets from different sources or time periods.

## Credentials

This node does not require credentials.

## Inputs

- **Input A**: First dataset to compare
- **Input B**: Second dataset to compare

## Outputs

- **In A only**: Items that exist only in Input A
- **Same**: Items that are identical in both inputs
- **Different**: Items that exist in both inputs but have differences
- **In B only**: Items that exist only in Input B

## Properties

### Resource: Data Comparison

#### Operation: Dataset Comparison

| Field Name | Type | Description | Required |
|---|---|---|---|
| Fields to Match | Fixed Collection | Fields used to pair items between datasets | Yes |
| When There Are Differences | Options | How to handle different versions of the same item | Yes |
| Fuzzy Compare | Boolean | Whether to tolerate small type differences when comparing | No |
| Prefer | Options | Which input to prefer when using mix mode | No |
| For Everything Except | String | Fields to exclude when using mix mode | No |

## Comparison Options

- **Use Input A Version**: Keep the version from Input A when differences exist
- **Use Input B Version**: Keep the version from Input B when differences exist
- **Use a Mix of Versions**: Use different inputs for different fields
- **Include Both Versions**: Include all data from both inputs

## UseCases

- **Data Synchronization** : Compare datasets to identify what needs to be synchronized
- **Change Detection** : Detect changes between different versions of data
- **Data Quality Analysis** : Identify inconsistencies between data sources
- **Merge Operations** : Prepare data for merging by identifying differences
- **Audit Trails** : Track changes in data over time by comparing snapshots
