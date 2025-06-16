# QuickChart Node

## Overview

Create a chart via QuickChart

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: chart

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Chart Type | options | The type of chart to create | Yes | bar |
| Add Labels | options | How to add labels to the chart | Yes | manually |
| Labels | fixedCollection | Labels to use in the chart | Yes |  |
| Labels Array | string | The array of labels to be used in the chart | Yes |  |
| Data | json | Data to use for the dataset | Yes |  |
| Put Output In Field | string | The name of the output field to put the binary file data in | Yes | data |
| Chart Options | collection | Additional options for the chart | No |  |

##### Chart Options Collection Fields:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Background Color | color | Background color of the chart | No |  |
| Device Pixel Ratio | number | Pixel ratio of the chart | No | 2 |
| Format | options | File format of the resulting chart (PNG, PDF, SVG, WebP) | No | png |
| Height | number | Height of the chart | No | 300 |
| Width | number | Width of the chart | No | 500 |
| Horizontal | boolean | Whether the chart should use its Y axis horizontal | No | false |

##### Dataset Options Collection Fields:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Background Color | color | Background color for the dataset | No |  |
| Border Color | color | Border color for the dataset | No |  |
| Border Width | number | Border width for the dataset | No |  |
| Fill | boolean | Whether to fill the area under the line | No | false |
| Point Style | options | Style of the data points | No | circle |
| Item Style | options | Style of chart items | No |  |

## UseCases

- [Data Visualization] : Create charts and graphs to visualize business data and analytics
- [Report Generation] : Generate visual reports with charts for stakeholders and clients
- [Dashboard Creation] : Build dynamic dashboards with real-time chart visualizations
- [Email Reports] : Include charts in automated email reports and newsletters
- [Social Media Graphics] : Create chart images for sharing insights on social media platforms
- [Presentation Materials] : Generate charts for presentations and business proposals
- [Website Analytics] : Display website performance metrics and analytics in chart format
- [Sales Reporting] : Visualize sales data and performance metrics for teams and management
- [Marketing Analytics] : Create charts showing marketing campaign performance and ROI
- [Financial Reporting] : Generate financial charts for budget tracking and analysis
- [Performance Monitoring] : Visualize system performance metrics and KPIs
- [Survey Results] : Create charts to display survey responses and research data
- [Progress Tracking] : Visualize project progress and milestone achievements
- [Comparative Analysis] : Create comparison charts for competitor analysis and benchmarking
- [Data Export] : Export data visualizations for use in documents and presentations

