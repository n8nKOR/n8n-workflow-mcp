# urlscan.io Node

## Overview

No description available

## Credentials

- Name: urlScanIoApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: scan

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Scan ID | string | ID of the scan to retrieve | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Query using the <a href= | No |  |

#### Operation: perform

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| URL | string | URL to scan | No |  |
| Additional Fields | collection | <code>User-Agent</code> header to set for this scan. Defaults to <code>n8n</code> | No |  |

## UseCases

- **Security Threat Analysis** : Analyze suspicious URLs and websites for potential security threats
- **Phishing Detection** : Identify and analyze phishing websites and malicious URLs
- **Website Security Auditing** : Perform security audits of websites and web applications
- **Malware Analysis** : Analyze websites for malware distribution and malicious content
- **Brand Protection** : Monitor for fraudulent websites impersonating your brand or organization
- **Incident Response** : Investigate security incidents involving suspicious web activity
- **Threat Intelligence** : Gather threat intelligence data from web-based sources
- **Compliance Monitoring** : Monitor websites for compliance with security and regulatory standards
- **Automated Security Scanning** : Integrate URL scanning into automated security workflows
- **Forensic Investigation** : Conduct forensic analysis of web-based security incidents

