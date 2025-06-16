# SecurityScorecard Node

## Overview

Consume SecurityScorecard API

## Credentials

- Name: securityScorecardApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: company

#### Operation: getScorePlan

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Score | number | Score target | Yes |  |

#### Operation: getFactor

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection | Filter issues by comma-separated severity list | No |  |

### Resource: industry

#### Operation: getFactorHistorical

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Options | collection | History start date | No |  |

### Resource: invite

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string |  | Yes |  |
| First Name | string |  | Yes |  |
| Last Name | string |  | Yes |  |
| Message | string | Message for the invitee | Yes |  |
| Additional Fields | collection | Minimum days to resolve a scorecard issue | No |  |

### Resource: portfolio

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: portfolioCompany

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Filters | collection | Company score grade filter | No |  |

### Resource: report

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: generate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Report | options |  | Yes |  |
| Scorecard Identifier | string | Primary identifier of a company or scorecard, i.e. domain. | Yes |  |
| Portfolio ID | string |  | Yes |  |
| Branding | options |  | No |  |
| Date | dateTime |  | Yes |  |
| Options | collection |  | No |  |
| Options | collection |  | No |  |

#### Operation: download

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Report URL | string | URL to a generated report | Yes |  |
| Put Output File in Field | string |  | Yes |  |

## UseCases

- **Vendor Risk Assessment** : Assess security risks of third-party vendors and suppliers
- **Continuous Security Monitoring** : Monitor security posture of organizations and partners
- **Compliance Reporting** : Generate security compliance reports for regulatory requirements
- **Supply Chain Security** : Evaluate security risks across the entire supply chain
- **Security Benchmarking** : Compare security performance against industry standards
- **Risk Management** : Integrate security scores into enterprise risk management processes
- **Due Diligence** : Perform security due diligence for mergers and acquisitions
- **Security Analytics** : Analyze security trends and patterns across organizations
- **Threat Intelligence** : Gather threat intelligence data for security decision making
- **Security Automation** : Automate security assessment and monitoring workflows

