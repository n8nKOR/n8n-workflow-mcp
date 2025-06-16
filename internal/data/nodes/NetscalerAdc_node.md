# Netscaler ADC Node

## Overview

The Netscaler ADC node provides integration with Citrix Netscaler Application Delivery Controller (ADC) through its Nitro REST API. This node enables automation of SSL certificate management and file operations on Netscaler ADC appliances, supporting certificate lifecycle management, file transfers, and system configuration tasks. The node uses specific API endpoints including `/config/systemfile`, `/config/sslcert`, and `/config/sslcertkey` for comprehensive certificate and file management operations.

## Credentials

- **Name**: citrixAdcApi
- **Required**: Yes

### Credential Configuration
To configure Netscaler ADC API credentials:
1. Access your Netscaler ADC management interface
2. Enable NITRO REST API access (typically on port 80/443)
3. Create or use existing user account with appropriate permissions (nsroot or custom user)
4. Configure API access permissions for certificate and file operations
5. Use the management IP/hostname and credentials in n8n
6. Ensure SSL certificate validation is properly configured for HTTPS connections

## Inputs

- **Main**: Certificate data, file content, or configuration parameters

## Outputs

- **Main**: API responses including operation status, certificate details, or file information
  - **JSON Response Structure**: Returns structured JSON with operation results and error handling
  - **File Operations**: Binary data encoded in base64 format for file uploads
  - **Certificate Operations**: Certificate metadata and validation status

## Properties

### Resource: File (Default)

#### Operation: Upload

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Location | String | Target directory path on ADC | Yes | /nsconfig/ssl/ |
| Binary Property | String | Binary data field name containing file content | Yes | data |
| File Name | String | Override filename (optional) | No | - |

**Supported File Locations:**
- `/nsconfig/ssl/` - SSL certificates and keys (default)
- `/var/netscaler/ssl/` - SSL configuration files
- `/nsconfig/` - General configuration files
- `/var/tmp/` - Temporary files

**Technical Parameters:**
- Uses `filecontent` parameter for base64 encoded file data
- `fileencoding` parameter set to "BASE64" for binary transfers
- `filelocation` parameter specifies target directory

#### Operation: Download

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Location | String | Source directory path on ADC | Yes | /nsconfig/ssl/ |
| File Name | String | Name of file to download | Yes | - |
| Binary Property | String | Output binary field name | Yes | data |

#### Operation: Delete

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| File Location | String | Directory path on ADC | Yes | /nsconfig/ssl/ |
| File Name | String | Name of file to delete | Yes | - |

### Resource: Certificate

#### Operation: Create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Certificate File Name | String | Output certificate file name and path | Yes | - |
| Certificate Format | Options | Certificate format (PEM/DER) | Yes | PEM |
| Certificate Type | Options | Certificate type (Root-CA, Intermediate-CA, Server, Client) | No | Server |
| Certificate Request File Name | String | CSR file name and path | Yes | - |
| Private Key File Name | String | Private key file name (for Root-CA) | Conditional | - |
| Private Key Format | Options | Private key format (PEM/DER) | No | PEM |
| CA Certificate File Name | String | CA certificate file name | Conditional | - |
| CA Certificate File Format | Options | CA certificate format (PEM/DER) | No | PEM |
| CA Private Key File Name | String | CA private key file name | Conditional | - |
| CA Private Key File Format | Options | CA private key format (PEM/DER) | No | PEM |
| CA Serial File Number | String | CA serial file number | Conditional | - |
| PEM Passphrase | String | Passphrase for encrypted keys (Required for PEM format) | Conditional | - |
| Subject Alternative Name | String | SAN extension for certificate | No | - |
| Validity Period | Number | Certificate validity in days | No | 365 |

**Certificate Format Requirements:**
- PEM format certificates require passphrase when encrypted
- Password field is required when certificateFormat is PEM and private key is encrypted
- DER format certificates do not require passphrase

#### Operation: Install

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Certificate-Key Pair Name | String | Name for the certificate-key pair | Yes | - |
| Certificate File Name | String | X509 certificate file name and path | Yes | - |
| Private Key File Name | String | Private key file name and path | Yes | - |
| Certificate Format | Options | Certificate format (PEM/DER) | Yes | PEM |
| Password | String | Password for encrypted private keys (Required for PEM with encryption) | Conditional | - |
| Notify When Expires | Boolean | Enable expiration notification alerts | No | false |
| Notification Period | Number | Days before expiration to alert (10-100) | No | 30 |
| Certificate Bundle | Boolean | Parse certificate chain as single file | No | false |

**Certificate Bundle Functionality:**
- When enabled, processes multiple certificates from single PEM file
- Automatically separates certificate chain components
- Supports complex certificate hierarchies and intermediate certificates
- Maintains certificate order and relationship validation

## Technical Details

### API Version Compatibility
- Supports Netscaler ADC 12.0 and later
- Uses NITRO REST API v1
- Requires proper API version headers for compatibility

### API Endpoints
- `/config/systemfile` - File management operations
- `/config/sslcert` - SSL certificate operations  
- `/config/sslcertkey` - Certificate-key pair management
- Base64 encoding required for all file uploads
- JSON response format with structured error handling

### Security Considerations
- Always use HTTPS for production environments
- Implement proper certificate validation
- Use service accounts with minimal required permissions
- Regular rotation of API credentials recommended
- File content transmitted as base64 encoded data for security

### Performance Optimization
- Batch operations when possible
- Monitor API rate limits (default: 100 requests/minute)
- Use appropriate timeout values for large file operations
- Implement proper error handling and retries
- Base64 encoding adds ~33% overhead to file size

### Error Handling
- Structured JSON error responses from API
- Validation errors for certificate format mismatches
- File permission and location validation
- Certificate chain validation and verification

## UseCases

- **SSL Certificate Management**: Automate SSL certificate creation, installation, and renewal processes with proper format validation
- **Certificate Lifecycle Automation**: Create complete certificate chains from root to server certificates with automatic encoding
- **File Transfer Operations**: Upload configuration files, certificates, and keys to ADC appliances using base64 encoding
- **Configuration Backup**: Download configuration files and certificates for backup purposes with proper binary handling
- **Certificate Authority Operations**: Create and manage internal certificate authorities with passphrase management
- **Multi-Environment Deployment**: Automate certificate deployment across development, staging, and production ADCs
- **Compliance Automation**: Ensure certificate compliance by automated renewal and monitoring with format validation
- **Integration with PKI Systems**: Connect with external PKI systems for certificate management using API endpoints
- **Security Policy Enforcement**: Automate security policy deployment through structured file operations
- **Disaster Recovery**: Automate certificate and configuration restoration procedures with proper encoding
- **Certificate Inventory Management**: Track and manage certificate installations across multiple ADCs
- **Automated Certificate Provisioning**: Integrate with ACME protocols for automatic certificate provisioning
- **Custom Certificate Validation**: Implement custom certificate validation workflows with format checking
- **Certificate Monitoring**: Set up automated certificate expiration monitoring and alerts
- **Bulk Certificate Operations**: Process multiple certificates in batch operations with bundle support
- **Cross-Platform Integration**: Integrate ADC certificate management with external security tools
- **Development Environment Setup**: Automate development environment certificate provisioning
- **Load Balancer Certificate Management**: Manage certificates for load-balanced applications
- **API Gateway Security**: Automate API gateway certificate management and updates
- **Zero-Downtime Certificate Renewal**: Implement zero-downtime certificate renewal processes