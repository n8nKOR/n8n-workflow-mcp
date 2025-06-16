# ProfitWell

## Overview

The ProfitWell node allows you to access subscription analytics and financial metrics using the ProfitWell API. ProfitWell is a revenue recognition and subscription analytics platform designed for SaaS businesses to track key financial metrics, monitor subscription performance, and gain insights into customer behavior and revenue trends. This node enables you to integrate ProfitWell's comprehensive analytics into your n8n workflows for automated reporting and business intelligence.

## Credentials

This node requires ProfitWell API credentials:
- **Credential Name**: `profitWellApi`
- **Required Fields**: 
  - Access Token: Your ProfitWell API access token

To obtain API credentials:
1. Log into your ProfitWell account at https://www.profitwell.com
2. Navigate to "Settings" → "Integrations" → "API"
3. Generate an API access token
4. Copy the access token for use in n8n

## Inputs

- **Main**: Query parameters for metrics retrieval and settings requests

## Outputs

- **Main**: Returns JSON data containing financial metrics, company settings, or plan information from ProfitWell

## Properties

### Resource: Company

#### Operation: Get Settings

| Field Name | Type | Description | Required |
|---|---|---|---|
| (No additional fields) | - | Retrieves company's ProfitWell account settings | - |

### Resource: Metric

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type | Options | Metric breakdown type (daily/monthly) | Yes |
| Month | String | Target month in YYYY-MM format (current or previous month for daily) | Yes* |
| Simplify | Boolean | Return simplified response format | No |
| Options | Collection | Additional filtering options | No |

*Required only for daily metrics

##### Additional Options:

| Field Name | Type | Description | Default |
|---|---|---|---|
| Plan ID | String | Filter metrics by specific plan ID | All plans |

## UseCases

- **Financial Reporting** : Generate automated monthly and quarterly financial reports with key SaaS metrics
- **Revenue Monitoring** : Track MRR, ARR, and revenue growth trends for executive dashboards
- **Churn Analysis** : Monitor customer churn rates and identify patterns for retention strategies
- **Cohort Analysis** : Analyze customer behavior and lifetime value by acquisition cohorts
- **Performance Alerts** : Set up automated alerts for significant changes in key metrics
- **Business Intelligence** : Integrate subscription metrics into data warehouses and BI tools
- **Investor Reporting** : Prepare standardized investor reports with SaaS KPIs and growth metrics
- **Plan Optimization** : Analyze individual subscription plan performance for pricing strategies
- **Customer Success** : Monitor customer health scores and expansion opportunities
- **Forecasting Models** : Use historical data for revenue forecasting and growth projections
- **Marketing ROI** : Correlate marketing spend with customer acquisition and LTV metrics
- **Operational Dashboards** : Create real-time dashboards for operations and finance teams
