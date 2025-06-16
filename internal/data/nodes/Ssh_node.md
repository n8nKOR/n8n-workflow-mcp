# SSH Node

## Overview

Execute commands via SSH

## Credentials

- Name: sshPassword, Required: Yes (when using password authentication)
- Name: sshPrivateKey, Required: Yes (when using private key authentication)

## Inputs

- Main

## Outputs

- Main

## Properties

### Authentication

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Password or Private Key | Yes |

### Resource: command

#### Operation: execute

| Field Name | Type | Description | Required |
|---|---|---|---|
| Command | string | The command to be executed on a remote device | No |
| Working Directory | string | The working directory for command execution | Yes |

### Resource: file

#### Operation: upload

| Field Name | Type | Description | Required |
|---|---|---|---|
| Input Binary Field | string | The name of the input binary field containing the file to be uploaded | Yes |
| Target Directory | string | The directory to upload the file to. The name of the file does not need to be specified, it's taken from the binary data file name | Yes |
| Options | collection | Additional options for file upload | No |

#### Operation: download

| Field Name | Type | Description | Required |
|---|---|---|---|
| Path | string | The file path of the file to download. Has to contain the full path including file name | Yes |
| File Property | string | Object property name which holds binary data | Yes |
| Options | collection | Additional options for file download | No |

## UseCases

- **Server Administration** : Execute remote server administration commands and maintenance tasks
- **Deployment Automation** : Automate application deployments and server configurations
- **System Monitoring** : Monitor system health and performance through remote commands
- **File Management** : Manage files and directories on remote servers
- **Database Operations** : Execute database maintenance and backup operations remotely
- **Log Analysis** : Analyze server logs and system information remotely
- **Security Auditing** : Perform security audits and compliance checks on remote systems
- **Backup Operations** : Automate backup processes and data synchronization
- **Configuration Management** : Manage server configurations and environment settings
- **DevOps Automation** : Integrate SSH operations into DevOps pipelines and workflows

