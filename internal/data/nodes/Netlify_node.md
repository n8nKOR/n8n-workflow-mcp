# Netlify Node

## Overview

⚠️ **Operation Coverage**: This documentation covers the main operations. The implementation supports additional operations including deploy cancel, deploy delete, site get, and site delete operations beyond the basic getAll and create operations.

Comprehensive integration with Netlify API for managing static site deployments and hosting infrastructure. Netlify is a modern hosting platform that provides continuous deployment, serverless functions, and global CDN for JAMstack applications. This node enables automation of deployment workflows, site management, and hosting operations.

## Credentials

- **Name**: netlifyApi
- **Required**: Yes

### Credential Configuration
To configure Netlify API credentials:
1. Log into your Netlify account
2. Go to User Settings → Applications
3. Click "New access token"
4. Provide a description for the token
5. Copy the generated token for use in n8n

## Inputs

- **Main**: JSON data for site configuration and deployment parameters

## Outputs

- **Main**: Returns JSON responses from Netlify API including deployment status, site information, and operation results

## Properties

### Resource: deploy

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Site ID | String | ID of the site to get deployments for | Yes | - |
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 50 |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deploy ID | String | ID of the deployment to retrieve | Yes | - |

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Site ID | String | ID of the site to deploy to | Yes | - |
| Title | String | Title for the deployment | No | - |
| Additional Fields | Collection | Advanced deployment configuration | No | - |

**Additional Fields Options:**
- **Branch**: Git branch to deploy from
- **Clear Cache**: Clear CDN cache after deployment
- **Production Deploy**: Mark as production deployment
- **Deploy Preview**: Create preview deployment for testing

#### Operation: cancel

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deploy ID | String | ID of the deployment to cancel | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Deploy ID | String | ID of the deployment to delete | Yes | - |

### Resource: site

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | Number | Max number of results to return | No | 50 |
| Filter | String | Filter sites by name or URL | No | - |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Site ID | String | ID of the site to retrieve | Yes | - |

#### Operation: delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Site ID | String | ID of the site to delete | Yes | - |

## UseCases

- **Static Site Deployment**: Deploy and host static websites and single-page applications with automated build processes
- **Continuous Deployment**: Set up continuous deployment pipelines from Git repositories with automatic builds on push
- **JAMstack Development**: Build and deploy JAMstack applications with modern development workflows and performance optimization
- **Form Handling**: Handle form submissions without backend server infrastructure using Netlify Forms
- **CDN Management**: Distribute content globally through integrated CDN network with automatic optimization
- **Domain Management**: Manage custom domains and SSL certificates for websites with automatic HTTPS
- **Branch Previews**: Create preview deployments for feature branches and pull requests for testing and review
- **Build Automation**: Automate build processes and site generation workflows with configurable build settings
- **A/B Testing**: Implement split testing and feature flags for website optimization and user experience testing
- **Serverless Functions**: Deploy serverless functions for backend functionality without managing servers
- **Asset Optimization**: Optimize images and assets for better performance with automatic compression and resizing
- **Analytics Integration**: Integrate web analytics and performance monitoring for data-driven insights
- **Team Collaboration**: Enable team collaboration on web development projects with role-based access control
- **Rollback Management**: Manage deployment rollbacks and version control with instant rollback capabilities
- **Edge Computing**: Deploy edge functions for enhanced performance and functionality at the network edge
- **Environment Management**: Manage different environments (development, staging, production) with environment-specific configurations
- **Webhook Integration**: Set up webhooks for deployment notifications and integration with external services
- **Performance Monitoring**: Monitor site performance and uptime with built-in analytics and alerting
- **Security Management**: Implement security headers, access control, and DDoS protection for hosted sites
- **API Gateway**: Use Netlify as an API gateway for serverless applications and microservices architecture

