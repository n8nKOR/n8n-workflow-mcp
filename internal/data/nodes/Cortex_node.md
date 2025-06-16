# Cortex Node

## Overview

Apply the Cortex analyzer/responder on the given entity

## Credentials

- Name: cortexApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: analyzer

#### Operation: execute

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Analyzer Type Name or ID | options | Choose the analyzer. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Observable Type Name or ID | options | Choose the observable type. Choose from the list, or specify an ID using an <a href= | Yes |  |
| Observable Value | string | Enter the observable value | Yes |  |
| Put Output File in Field | string |  | Yes |  |
| TLP | options | The TLP of the analyzed observable | No |  |
| Additional Fields | collection | Whether to force bypassing the cache | No |  |

## UseCases

- **Threat Intelligence Analysis**: Analyze suspicious files, URLs, IP addresses, and domains for potential security threats
- **Automated Malware Detection**: Submit files to various security analyzers for comprehensive malware analysis and reputation checking
- **IOC Enrichment**: Enrich indicators of compromise with additional context and threat intelligence data
- **Incident Response Automation**: Execute automated response actions on security entities like cases, alerts, and artifacts
- **Security Orchestration**: Integrate threat analysis into broader security workflows and SOAR platforms
- **Phishing Investigation**: Analyze suspicious URLs and email attachments for phishing campaigns
- **Network Security Monitoring**: Automate analysis of suspicious network artifacts and indicators

