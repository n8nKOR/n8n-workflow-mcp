# NPM Node

## Overview

The NPM node allows you to interact with the NPM registry API to retrieve package information, search for packages, manage distribution tags, and analyze package metadata. This node enables automation of package management tasks, dependency analysis, and registry monitoring workflows within n8n.

## Credentials

This node supports optional NPM API credentials:
- **API Key**: NPM authentication token (required for write operations like updating distribution tags)
- **Registry URL**: Custom NPM registry URL (defaults to official NPM registry)

To obtain NPM API credentials:
1. Log into your NPM account at npmjs.com
2. Navigate to Access Tokens in your account settings
3. Generate a new token with appropriate permissions
4. Use the token for authenticated operations

## Inputs

- **Main**: Input data for package operations

## Outputs

- **Main**: Returns NPM registry data including package metadata, search results, or distribution tag information

## Properties

### Resource: Package

#### Operation: Get Metadata
| Field Name | Type | Description | Required |
|---|---|---|---|
| Package Name | String | Name of the NPM package | Yes |
| Version | String | Specific version to retrieve (defaults to latest) | No |

#### Operation: Get Versions
| Field Name | Type | Description | Required |
|---|---|---|---|
| Package Name | String | Name of the NPM package | Yes |

#### Operation: Search
| Field Name | Type | Description | Required |
|---|---|---|---|
| Query | String | Search query string | Yes |
| Size | Number | Number of results to return (max 250) | No |
| From | Number | Offset for pagination | No |
| Quality | Number | Quality weight for search ranking (0-1) | No |
| Popularity | Number | Popularity weight for search ranking (0-1) | No |
| Maintenance | Number | Maintenance weight for search ranking (0-1) | No |

### Resource: Distribution Tag

#### Operation: Get All
| Field Name | Type | Description | Required |
|---|---|---|---|
| Package Name | String | Name of the NPM package | Yes |

#### Operation: Update
| Field Name | Type | Description | Required |
|---|---|---|---|
| Package Name | String | Name of the NPM package | Yes |
| Tag Name | String | Distribution tag name (e.g., latest, beta, alpha) | Yes |
| Version | String | Package version to tag | Yes |

## UseCases

- **Dependency Monitoring**: Monitor versions and updates of project dependencies
- **Package Analysis**: Analyze package metadata, dependencies, and maintenance status
- **Registry Search**: Search for packages matching specific criteria or functionality
- **Version Management**: Automate distribution tag management for package releases
- **Security Auditing**: Check package versions against known security vulnerabilities
- **Release Automation**: Automate package publishing and tagging workflows
- **Dependency Updates**: Identify outdated dependencies and available updates
- **License Compliance**: Audit package licenses for compliance requirements
- **Download Analytics**: Track package download statistics and popularity trends
- **Quality Assessment**: Evaluate package quality scores and maintenance status
- **CI/CD Integration**: Integrate package management into continuous integration workflows
- **Private Registry Management**: Manage packages in private NPM registries
- **Package Discovery**: Discover new packages and alternatives for specific use cases
- **Ecosystem Analysis**: Analyze the NPM ecosystem and package relationships
- **Automated Documentation**: Generate dependency documentation from package metadata

