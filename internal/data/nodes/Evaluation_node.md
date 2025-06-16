# Evaluation Node

## Overview

The Evaluation node enables workflow performance assessment and metrics collection for n8n workflows. This node allows you to evaluate workflow execution results, configure outputs, and collect metrics to measure workflow quality and performance. It's particularly useful for A/B testing, performance monitoring, and quality assurance workflows.

## Credentials

### Google Sheets Integration (for Set Outputs operation)
- **Service Account**: `googleApi`
- **OAuth2**: `googleSheetsOAuth2Api`

Authentication credentials are only required when using the Set Outputs operation with Google Sheets integration for storing evaluation results.

## Inputs

- **Main**: Input data to be evaluated

## Outputs

- **Dynamic Outputs**: Output ports are dynamically determined based on the configured operation

## Properties

### Resource: Evaluation

#### Operation: Set Outputs

| Field Name | Type | Description | Required |
|---|---|---|---|
| Output Configuration | Fixed Collection | Define output ports and routing | Yes |
| Options | Collection | Additional configuration options | No |

#### Output Configuration Fixed Collection

| Field Name | Type | Description | Required |
|---|---|---|---|
| Output Name | String | Name of the output port | Yes |
| Condition | String | Condition for routing to this output | No |

#### Options Collection

| Field Name | Type | Description | Required |
|---|---|---|---|
| Google Sheets Integration | Boolean | Enable Google Sheets output | No |
| Spreadsheet ID | String | Google Sheets spreadsheet identifier | No |
| Sheet Name | String | Name of the sheet to write to | No |
## UseCases

- **A/B Testing**: Compare two workflow versions to determine optimal performance and user experience outcomes
- **Quality Management**: Evaluate data processing workflow accuracy and completeness for quality assurance
- **Performance Monitoring**: Track workflow execution time, throughput, and resource utilization metrics
- **Business KPI Tracking**: Monitor revenue, conversion rates, customer satisfaction, and other business metrics
- **ML Model Evaluation**: Calculate machine learning model accuracy, precision, recall, and F1-score metrics
- **Workflow Optimization**: Compare performance across different configurations to identify optimal settings
- **Compliance Monitoring**: Ensure workflows meet regulatory and quality standards through systematic evaluation
- **Error Rate Analysis**: Track and analyze failure rates and error patterns for process improvement
- **Resource Utilization**: Monitor system resource usage and identify optimization opportunities
- **User Experience Testing**: Evaluate user interaction workflows and optimize for better engagement
- **Cost Analysis**: Track and evaluate workflow operational costs and resource efficiency
- **Security Assessment**: Evaluate security metrics and compliance with security policies

### Usage Notes
- Dynamic output configuration enables flexible workflow routing based on evaluation results
- Google Sheets integration provides external storage and analysis capabilities for evaluation data
- Supports both real-time and batch evaluation modes for different use case requirements
- Metric collection can be combined with alerting systems for automated monitoring
- Evaluation mode detection enables conditional workflow behavior during testing phases
- Authentication is only required when using external integrations like Google Sheets
- Multiple metrics can be collected simultaneously for comprehensive evaluation
- Output ports are created dynamically based on configuration, enabling flexible workflow design 