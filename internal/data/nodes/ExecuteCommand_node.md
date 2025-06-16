# Execute Command

## Overview

The Execute Command node executes system commands on the host machine where n8n is running. This node allows you to run shell commands, scripts, and system utilities directly from your workflows, enabling integration with system-level operations and external tools.

## Credentials

This node does not require credentials but runs with the permissions of the n8n process.

## Inputs

- **Main**: Accepts input data that can be used in command expressions

## Outputs

- **Main**: Returns command execution results including exit code, stdout, and stderr

## Properties

### Resource: System Operations

#### Operation: Command Execution

| Field Name | Type | Description | Required |
|---|---|---|---|
| Execute Once | Boolean | Whether to execute only once instead of once for each input item | No |
| Command | String | The command to execute on the host system | Yes |

## Security Considerations

- Commands run with n8n process permissions
- Be cautious with user input in commands to prevent injection attacks
- Consider using allowlists for permitted commands in production
- Monitor command execution for security and performance

## UseCases

- **System Integration** : Execute system commands and scripts from workflows
- **File Operations** : Perform file system operations like copying, moving, or processing files
- **External Tool Integration** : Run external tools and utilities as part of workflows
- **System Monitoring** : Execute monitoring commands and collect system information
- **Deployment Automation** : Automate deployment scripts and system configuration tasks

