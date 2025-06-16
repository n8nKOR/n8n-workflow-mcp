# Metabase Node

## Overview

⚠️ **Documentation Status**: This documentation provides basic coverage of the Metabase node. The actual implementation includes comprehensive Alert, Database, Metric, and Question resources with modular structure and separate description files. The implementation provides much richer API functionality than currently documented here.

The Metabase node provides comprehensive integration with Metabase's REST API for business intelligence and data analytics automation. Metabase is an open-source business intelligence tool that allows you to create dashboards, charts, and automated reports from your data. This node enables you to automate database operations, manage alerts, work with metrics, and interact with questions (queries) programmatically.

## Credentials

- **Name**: metabaseApi
- **Required**: Yes

### Credential Configuration
To configure Metabase API credentials:
1. Access your Metabase instance as an administrator
2. Go to Admin → Settings → API
3. Create an API key or use session-based authentication
4. Configure the base URL of your Metabase instance
5. Set up appropriate permissions for the operations you plan to use

## Inputs

- **Main**: JSON data for creating/updating Metabase resources

## Outputs

- **Main**: Returns JSON responses from Metabase API including data, metadata, and operation results

## Properties

### Resource: Alert

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Card ID | number | ID of the card/question to monitor | Yes |
| Alert Name | string | Name for the alert | Yes |
| Alert Type | options | Type of alert (goal/progress) | Yes |
| Channels | array | Notification channels (email, Slack, etc.) | Yes |
| Alert Condition | object | Conditions for triggering the alert | Yes |
| Recipients | array | List of recipients for alert notifications | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Alert ID | number | ID of the alert to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | boolean | Whether to return all alerts or limit results | No |
| Limit | number | Maximum number of alerts to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Alert ID | number | ID of the alert to update | Yes |
| Alert Name | string | Updated alert name | No |
| Alert Condition | object | Updated alert conditions | No |
| Recipients | array | Updated recipient list | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Alert ID | number | ID of the alert to delete | Yes |

### Resource: Database

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | number | ID of the database to retrieve | Yes |
| Include Tables | boolean | Include table information | No |
| Include Fields | boolean | Include field information | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | boolean | Whether to return all databases | No |
| Limit | number | Maximum number of databases to return | No |
| Include Metadata | boolean | Include database metadata | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | number | ID of the database to update | Yes |
| Database Name | string | Updated database name | No |
| Connection Details | object | Updated connection parameters | No |

#### Operation: Sync Schema
| Field Name | Type | Description | Required |
|---|---|---|---|
| Database ID | number | ID of the database to sync | Yes |
| Full Sync | boolean | Perform full schema synchronization | No |

### Resource: Metric

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Table ID | number | ID of the table for the metric | Yes |
| Metric Name | string | Name of the metric | Yes |
| Definition | object | Metric definition/aggregation | Yes |
| Description | string | Description of the metric | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Metric ID | number | ID of the metric to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Table ID | number | Filter metrics by table | No |
| Return All | boolean | Whether to return all metrics | No |
| Limit | number | Maximum number of metrics to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Metric ID | number | ID of the metric to update | Yes |
| Metric Name | string | Updated metric name | No |
| Definition | object | Updated metric definition | No |
| Description | string | Updated description | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Metric ID | number | ID of the metric to delete | Yes |

### Resource: Question

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Question Name | string | Name of the question/query | Yes |
| Database ID | number | ID of the database to query | Yes |
| Query Type | options | Type of query (native/structured) | Yes |
| Query | object/string | Query definition or SQL | Yes |
| Collection ID | number | ID of collection to save in | No |
| Description | string | Description of the question | No |
| Visualization Settings | object | Chart/visualization configuration | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Question ID | number | ID of the question to retrieve | Yes |
| Include Query | boolean | Include query definition | No |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Collection ID | number | Filter by collection | No |
| Database ID | number | Filter by database | No |
| Return All | boolean | Whether to return all questions | No |
| Limit | number | Maximum number of questions to return | No |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Question ID | number | ID of the question to update | Yes |
| Question Name | string | Updated question name | No |
| Query | object/string | Updated query definition | No |
| Visualization Settings | object | Updated visualization settings | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Question ID | number | ID of the question to delete | Yes |

#### Operation: Execute
| Field Name | Type | Description | Required |
|---|---|---|---|
| Question ID | number | ID of the question to execute | Yes |
| Parameters | object | Query parameters | No |
| Format | options | Output format (json/csv/xlsx) | No |

## UseCases

- **Automated Business Intelligence**: Create automated BI workflows that generate reports, update dashboards, and distribute insights based on data changes
- **Alert Management**: Set up intelligent alerts that monitor KPIs and notify stakeholders when thresholds are reached or anomalies are detected
- **Data Pipeline Integration**: Integrate Metabase with data pipelines to automatically create questions and dashboards for new data sources
- **Report Automation**: Automate the creation and distribution of regular reports, eliminating manual report generation tasks
- **Database Monitoring**: Monitor database performance, schema changes, and data quality through automated Metabase interactions
- **Dynamic Dashboard Creation**: Programmatically create dashboards based on changing business requirements or new data sources
- **Metric Lifecycle Management**: Automate the creation, updating, and retirement of business metrics as requirements evolve
- **Cross-Platform Analytics Integration**: Connect Metabase with other analytics tools and data sources for comprehensive business intelligence
- **Compliance Reporting**: Generate compliance reports automatically by executing predefined questions and formatting results
- **Performance Analytics**: Monitor application performance by automatically querying metrics and generating performance dashboards
- **Custom Alert Systems**: Build sophisticated alert systems that combine multiple data sources and business logic
- **Data Governance**: Implement data governance workflows that track metric definitions, question lineage, and data usage
