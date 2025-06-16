# Splunk Node

## Overview

The Splunk node provides comprehensive integration with Splunk Enterprise API, enabling automated data analysis, search operations, alert management, and reporting workflows. This powerful node supports multiple Splunk resources including search jobs, alerts, reports, and user management, making it an essential tool for security operations, monitoring, analytics, and business intelligence workflows.

⚠️ **Version Information**: This node has multiple versions (V1, V2). The current default version is V2, which provides enhanced API integration, improved search functionality, and expanded resource management capabilities. Some features may differ between versions.

## Credentials

The Splunk node requires Splunk API credentials:

- **Credential Name**: `splunkApi`
- **Required**: Yes

## Inputs

- **Main**: Accepts data for search queries, alert configurations, and other operations

## Outputs

- **Main**: Returns search results, alert data, reports, or operation responses

## Properties

### Resource: Alert

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Alert Name | String | Unique identifier for the alert | Yes |
| Search Query | String | SPL query to trigger the alert | Yes |
| Alert Type | Options | Real-time or scheduled alert | Yes |
| Trigger Conditions | Collection | Conditions that trigger the alert | Yes |
| Actions | Collection | Actions to perform when triggered | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Alert Name | String | Name of the alert to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Alert Name | String | Name of the alert to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Limit | Number | Maximum number of alerts to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Alert Name | String | Name of the alert to update | Yes |
| Search Query | String | Updated SPL query | No |
| Alert Type | Options | Updated alert type | No |

### Resource: Report

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Report Name | String | Unique name for the report | Yes |
| Search Query | String | SPL query for data extraction | Yes |
| Schedule | Collection | Report generation schedule | No |
| Output Format | Options | CSV, PDF, XML output formats | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Report Name | String | Name of the report to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Report Name | String | Name of the report to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Limit | Number | Maximum number of reports to return | No |

### Resource: Search

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Query | String | SPL search query to execute | Yes |
| Execution Mode | Options | Blocking, Normal, or One Shot | No |
| Earliest Time | DateTime | Start time for search range | No |
| Latest Time | DateTime | End time for search range | No |
| Max Time | Number | Maximum execution time (seconds) | No |
| Search Mode | Options | Normal or Real Time | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Search ID | String | ID of the search job to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Search ID | String | ID of the search job to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Limit | Number | Maximum number of search jobs to return | No |

#### Operation: Get Result

| Field Name | Type | Description | Required |
|---|---|---|---|
| Search ID | String | ID of the search job | Yes |
| Output Mode | Options | Result output format | No |

### Resource: User

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Username | String | Unique username identifier | Yes |
| Password | String | User password | Yes |
| Full Name | String | User's full name | No |
| Email | String | User's email address | No |
| Roles | Collection | Assigned user roles | No |

#### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Username | String | Username to delete | Yes |

#### Operation: Get

| Field Name | Type | Description | Required |
|---|---|---|---|
| Username | String | Username to retrieve | Yes |

#### Operation: Get Many

| Field Name | Type | Description | Required |
|---|---|---|---|
| Limit | Number | Maximum number of users to return | No |

#### Operation: Update

| Field Name | Type | Description | Required |
|---|---|---|---|
| Username | String | Username to update | Yes |
| Password | String | New password | No |
| Full Name | String | Updated full name | No |
| Email | String | Updated email address | No |
| Roles | Collection | Updated user roles | No |

## UseCases

- **Threat Detection**: Monitor security events and identify threats
- **Incident Response**: Investigate security incidents with detailed queries
- **Compliance Monitoring**: Track compliance violations and generate reports
- **User Behavior Analytics**: Analyze user activities for anomalies
- **Network Security**: Monitor network traffic and detect intrusions
- **Application Performance**: Monitor application logs and performance metrics
- **Infrastructure Monitoring**: Track server health and resource utilization
- **Error Analysis**: Identify and analyze application errors and failures
- **Capacity Planning**: Analyze usage trends for resource planning
- **Service Level Monitoring**: Track SLA compliance and service quality
- **Business Metrics**: Track KPIs and business performance indicators
- **Customer Analytics**: Analyze customer behavior and engagement
- **Sales Reporting**: Generate sales reports and trend analysis
- **Marketing Analytics**: Measure campaign effectiveness and ROI
- **Operational Efficiency**: Monitor business process performance
- **CI/CD Monitoring**: Track deployment success and failure rates
- **Code Quality Analysis**: Monitor code commits and quality metrics
- **Performance Testing**: Analyze application performance under load
- **Error Tracking**: Monitor application errors and exceptions
- **Release Management**: Track release cycles and deployment metrics
- **Regulatory Compliance**: Monitor compliance with industry regulations
- **Audit Trail**: Maintain detailed audit logs for security and compliance
- **Data Governance**: Track data access and usage patterns
- **Risk Management**: Identify and assess operational risks
- **Policy Enforcement**: Monitor adherence to organizational policies
- **Anomaly Detection**: Identify unusual patterns in data
- **Predictive Analytics**: Forecast trends and potential issues
- **Correlation Analysis**: Find relationships between different events
- **Time Series Analysis**: Analyze data patterns over time
- **Statistical Modeling**: Build statistical models for data analysis
