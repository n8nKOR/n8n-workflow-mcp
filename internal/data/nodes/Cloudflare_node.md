# Cloudflare Node

## Overview

The Cloudflare node allows you to consume the Cloudflare API to manage zone-level authenticated origin pulls. Cloudflare is a global content delivery network and security service that provides DNS, DDoS protection, and performance optimization. This node specifically handles the management of TLS client certificates for zone-level authenticated origin pulls, which allow you to cryptographically verify that requests to your origin server have come from Cloudflare.

## Credentials

This node requires Cloudflare API credentials:
- **API Token**: Your Cloudflare API token with appropriate permissions
- **Zone:Zone:Read and Zone:Zone SSL and Certificates:Edit permissions are required**

To obtain API credentials:
1. Log into your Cloudflare account
2. Navigate to My Profile → API Tokens
3. Create a new token with Zone:Zone:Read and Zone:Zone SSL and Certificates:Edit permissions
4. Copy the token for use in n8n

## Inputs

- **Main**: JSON data for certificate upload and management operations

## Outputs

- **Main**: Returns JSON data from Cloudflare API responses

## Properties

### Resource: zoneCertificate

Zone-level authenticated origin pulls allow you to cryptographically verify that requests to your origin server have come from Cloudflare using TLS client certificate authentication.

#### Operation: upload

Upload a certificate for zone-level authenticated origin pulls.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Zone Name or ID | options | The zone to upload the certificate to | Yes |  |
| Certificate Content | string | The zone's leaf certificate | Yes |  |
| Private Key | string | The certificate's private key | Yes |  |

#### Operation: get

Get details of a specific certificate.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Zone Name or ID | options | The zone containing the certificate | Yes |  |
| Certificate ID | string | The ID of the certificate to retrieve | Yes |  |

#### Operation: getMany

Get multiple certificates for a zone.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Zone Name or ID | options | The zone to retrieve certificates from | Yes |  |
| Return All | boolean | Whether to return all results or only up to a given limit | No | false |
| Limit | number | Max number of results to return (1-50) | No | 25 |
| Filters | collection | Filters to apply to the results | No |  |

##### Filters for getMany

| Field Name | Type | Description | Options |
|---|---|---|---|
| Status | options | Status of the zone's custom SSL | active, expired, deleted, pending |

#### Operation: delete

Delete a certificate from zone-level authenticated origin pulls.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Zone Name or ID | options | The zone containing the certificate | Yes |  |
| Certificate ID | string | The ID of the certificate to delete | Yes |  |

## Features

- **Certificate Management**: Upload, retrieve, and delete TLS client certificates for authenticated origin pulls
- **Zone-Level Security**: Implement zone-level authenticated origin pulls to verify requests from Cloudflare
- **Certificate Lifecycle**: Manage the complete lifecycle of origin pull certificates
- **Status Filtering**: Filter certificates by status (active, expired, deleted, pending)
- **Multi-Zone Support**: Work with certificates across multiple Cloudflare zones

## Security Features

### Authenticated Origin Pulls
Zone-level authenticated origin pulls provide several security benefits:
- **Request Verification**: Cryptographically verify that requests originate from Cloudflare
- **TLS Client Authentication**: Use client certificates to authenticate Cloudflare requests
- **Origin Security**: Protect your origin server from direct attacks by requiring valid certificates
- **Traffic Validation**: Ensure all traffic to your origin comes through Cloudflare's network

### Certificate Management
- **Secure Upload**: Safely upload certificates and private keys to Cloudflare
- **Certificate Rotation**: Update and rotate certificates for ongoing security
- **Status Monitoring**: Track certificate status and expiration
- **Multi-Certificate Support**: Manage multiple certificates per zone

## UseCases

- **Origin Server Security**: Implement authenticated origin pulls to ensure all traffic to your origin server comes from Cloudflare, preventing direct attacks
- **Certificate Automation**: Automate the upload and rotation of TLS client certificates for authenticated origin pulls
- **Security Compliance**: Maintain security compliance by regularly updating and monitoring certificate status across multiple zones
- **Multi-Environment Management**: Manage certificates across development, staging, and production environments with different Cloudflare zones
- **Certificate Lifecycle Management**: Automate certificate expiration monitoring and renewal processes
- **Security Incident Response**: Quickly revoke and replace compromised certificates during security incidents

