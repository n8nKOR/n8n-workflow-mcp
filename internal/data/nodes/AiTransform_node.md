# AiTransform Node

## Overview

The AiTransform node allows you to modify data based on instructions written in plain English. This node uses AI to generate JavaScript code that transforms your input data according to your natural language instructions. It's perfect for data manipulation tasks without requiring manual coding.

## Credentials

This node does not require any credentials.

## Inputs

- **Main**: JSON data to be transformed

## Outputs

- **Main**: Returns transformed JSON data based on the generated code

## Properties

### Resource: Transform

#### Operation: Generate Code
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Instructions | string | Instructions on how to transform the data in plain English | Yes | - |
| Generated JavaScript | string | Read-only generated JavaScript code based on instructions | No | - |
| Generate Code | button | Click to generate JavaScript code from instructions | Yes | - |

## UseCases

- **Name Concatenation** : Combine first and last names into full name fields
- **Address Formatting** : Merge address components into formatted strings
- **Data Normalization** : Standardize field formats across different data sources
- **Field Restructuring** : Move data between nested objects and flat structures
- **Data Type Conversion** : Convert strings to numbers, dates, or other formats
- **Conditional Filtering** : Remove records based on business rules
- **Data Validation** : Filter out invalid or incomplete records
- **Priority Sorting** : Order data by importance, date, or custom criteria
- **Duplicate Removal** : Eliminate duplicate entries based on specific fields
- **Data Segmentation** : Separate data into different categories
- **Calculated Fields** : Add computed values like totals, percentages, or scores
- **Status Determination** : Set status fields based on business logic
- **Category Assignment** : Automatically categorize data based on rules
- **Timestamp Addition** : Add processing timestamps or date calculations
- **Reference Linking** : Create relationships between data records
- **Pricing Calculations** : Apply discounts, taxes, or complex pricing rules
- **Inventory Management** : Update stock levels and availability status
- **Customer Segmentation** : Classify customers based on behavior or demographics
- **Risk Assessment** : Calculate risk scores or compliance ratings
- **Performance Metrics** : Generate KPIs and business intelligence data
- **Data Extraction** : Pull specific fields from complex API responses
- **Format Standardization** : Convert API data to internal formats
- **Error Handling** : Process and categorize API errors
- **Response Aggregation** : Combine multiple API responses into unified datasets
- **Data Validation** : Verify API response completeness and accuracy

