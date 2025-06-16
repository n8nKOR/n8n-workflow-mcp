# Phantombuster Node

## Overview

Consume Phantombuster API

## Credentials

- Name: phantombusterApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: agent

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent ID | string |  | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: getOutput

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Resolve Data | boolean | By default the outpout is presented as string. If this option gets activated, it will resolve the data automatically. | No |  |
| Additional Fields | collection | If set, the output will be retrieved from the container after the specified previous container ID | No |  |

#### Operation: launch

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Agent Name or ID | options | Choose from the list, or specify an ID using an <a href= | Yes |  |
| Resolve Data | boolean | By default the launch just include the container ID. If this option gets activated, it will resolve the data automatically. | No |  |
| JSON Parameters | boolean |  | No |  |
| Additional Fields | collection | Agent argument. Can either be a JSON string or a plain object. The argument can be retrieved with buster.argument in the agent's script. | No |  |

## UseCases

- [Web Scraping] : Extract data from websites and web applications automatically
- [Social Media Automation] : Automate social media tasks and engagement activities
- [Lead Generation] : Generate leads by scraping business directories and social networks
- [Market Research] : Collect market data and competitor intelligence from web sources
- [Data Collection] : Gather large datasets from various online sources
- [LinkedIn Automation] : Automate LinkedIn prospecting and networking activities
- [Email Discovery] : Find and verify email addresses from web sources
- [Content Aggregation] : Aggregate content from multiple websites and platforms
- [Monitoring and Alerts] : Monitor websites for changes and receive automated alerts
- [API Alternative] : Access data from platforms without official APIs
- [Growth Hacking] : Implement growth hacking strategies through automation
- [Sales Prospecting] : Automate sales prospecting and lead qualification processes
- [Competitive Analysis] : Analyze competitor activities and strategies automatically
- [Data Enrichment] : Enrich existing datasets with additional information from web sources
- [Social Listening] : Monitor social media mentions and brand conversations

