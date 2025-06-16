# Crypto Node

## Overview

Provides comprehensive cryptographic operations for data security and integrity. Supports generating random strings, creating hashes (MD5, SHA variants), HMAC authentication, and digital signatures. Handles both text and binary data with multiple encoding formats (hex, base64). Essential for data validation, secure token generation, and implementing authentication mechanisms in workflows.

## Credentials

None (User-provided keys required for signing and HMAC operations)

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Crypto Operations

#### Operation: Generate

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the property to which to write the random string | Yes |
| Encoding Type | options | Encoding that will be used to generate string (ASCII, BASE64, HEX, UUID) | Yes |
| Length | number | Length of the generated string (for ASCII, BASE64, HEX) | No |

#### Operation: Hash

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type | options | The hash algorithm to use (MD5, SHA256, SHA3-256, SHA3-384, SHA3-512, SHA384, SHA512) | Yes |
| Binary Data | boolean | Whether the data to be hashed should be taken from binary field | Yes |
| Binary Property Name | string | Name of the binary property which contains the input data (when Binary Data = true) | Yes |
| Value | string | The value that should be hashed (when Binary Data = false) | Yes |
| Property Name | string | Name of the property to which to write the hash | Yes |
| Encoding | options | Output encoding format (BASE64, HEX) | Yes |

#### Operation: HMAC

| Field Name | Type | Description | Required |
|---|---|---|---|
| Type | options | The hash algorithm to use for HMAC (MD5, SHA256, SHA3-256, SHA3-384, SHA3-512, SHA384, SHA512) | Yes |
| Binary Data | boolean | Whether the data to be processed should be taken from binary field | Yes |
| Binary Property Name | string | Name of the binary property which contains the input data (when Binary Data = true) | Yes |
| Value | string | The value of which the HMAC should be created (when Binary Data = false) | Yes |
| Property Name | string | Name of the property to which to write the HMAC | Yes |
| Secret | string | Secret key used for HMAC generation | Yes |
| Encoding | options | Output encoding format (BASE64, HEX) | Yes |

#### Operation: Sign

| Field Name | Type | Description | Required |
|---|---|---|---|
| Value | string | The value that should be signed | Yes |
| Property Name | string | Name of the property to which to write the signed value | Yes |
| Algorithm | options | Cryptographic algorithm to use for signing | Yes |
| Encoding | options | Output encoding format for the signature (BASE64, HEX) | Yes |
| Private Key | string | Private key to use when signing the string | Yes |

## UseCases

- **Secure Token Generation** : Generate random strings for session tokens, API keys, and secure identifiers
- **Password Hashing** : Create secure password hashes using SHA256 or other strong algorithms
- **API Request Signing** : Sign API requests using HMAC for authentication and integrity verification
- **Digital Document Signing** : Create digital signatures for documents using RSA or other signing algorithms
- **File Integrity Verification** : Generate checksums for files to verify data integrity and detect corruption
- **Data Validation** : Create hashes for data validation and tamper detection
- **Authentication Tokens** : Generate secure authentication tokens for user sessions
- **Webhook Verification** : Verify webhook payloads using HMAC signatures
- **License Key Generation** : Create secure license keys and activation codes
- **Secure Random Generation** : Generate cryptographically secure random values for various purposes
- **Message Authentication** : Authenticate messages using HMAC with shared secrets
- **Certificate Signing** : Sign certificates and other cryptographic materials
- **Data Fingerprinting** : Create unique fingerprints for data deduplication and identification
- **Secure Communication** : Implement secure communication protocols with cryptographic operations
- **Audit Trail Creation** : Generate cryptographic proofs for audit trails and compliance
- **Key Derivation** : Derive encryption keys from passwords or other key material
- **Digital Timestamping** : Create cryptographic timestamps for document authenticity
- **Blockchain Operations** : Perform cryptographic operations for blockchain and cryptocurrency applications
- **Secure Storage** : Hash sensitive data before storing in databases
- **Identity Verification** : Create cryptographic proofs for identity verification systems

