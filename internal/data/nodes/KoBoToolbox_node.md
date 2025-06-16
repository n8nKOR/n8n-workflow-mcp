# KoBoToolbox Node

## Overview

Work with KoBoToolbox forms and submissions

## Credentials

- Name: koBoToolboxApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: form

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | Yes |  |
| Limit | number | Max number of results to return | No |  |
| Options | collection | Whether to sort by descending order | No |  |
| Filters | collection | A text search query based on form data - e.g.  | No |  |

### Resource: hook

#### Operation: retryOne

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Hook Log ID | string | Hook log ID (starts with hl, e.g. hlSbGKaUKzTVNoWEVMYbLHe) | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | Yes |  |
| Limit | number | Max number of results to return | No |  |

#### Operation: getLogs

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Log Status | options | Only retrieve logs with a specific status | No |  |
| Start Date | dateTime | Minimum date for the hook log to retrieve | No |  |
| End Date | dateTime | Maximum date for the hook log to retrieve | No |  |

### Resource: submission

#### Operation: setValidation

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Validation Status | options | Desired Validation Status | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | Yes |  |
| Limit | number | Max number of results to return | No |  |
| Filter | options |  | No |  |
| See <a href= | notice |  | No |  |
| Filters (JSON) | string |  | No |  |

### Resource: file

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Property Name | string | Name of the binary property to write the file into | Yes |  |
| Download File Content | boolean | Whether to download the file content into a binary property | Yes |  |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Upload Mode | options |  | Yes |  |
| Property Name | string | Name of the binary property containing the file to upload. Supported types: image, audio, video, csv, xml, zip. | Yes |  |
| File URL | string | HTTP(s) link to the file to upload | Yes |  |

## UseCases

- **Survey Data Collection** : Collect survey data and responses from field research
- **Humanitarian Data Gathering** : Gather data for humanitarian and development projects
- **Research Data Management** : Manage research data collection and analysis workflows
- **Field Data Collection** : Collect data in remote and challenging field environments
- **Mobile Data Collection** : Enable mobile data collection through smartphones and tablets
- **Monitoring and Evaluation** : Conduct monitoring and evaluation activities for projects
- **Community Feedback** : Collect feedback and input from communities and stakeholders
- **Emergency Response** : Gather data for emergency response and disaster management
- **Academic Research** : Support academic research data collection and analysis
- **NGO Operations** : Support NGO operations and program management activities

