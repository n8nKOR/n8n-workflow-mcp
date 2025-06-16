# TOTP Node

## Overview

Generate a time-based one-time password

## Credentials

- Name: totpApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: TOTP

#### Operation: Generate Secret

| Field Name | Type | Description | Required |
|---|---|---|---|
| operation | options | The operation to perform | Yes |
| algorithm | options | HMAC hashing algorithm (SHA1, SHA224, SHA256, SHA3-224, SHA3-256, SHA3-384, SHA3-512, SHA384, SHA512) | No |
| digits | number | Number of digits in the generated TOTP code (Default: 6) | No |
| period | number | How many seconds the generated TOTP code is valid for (Default: 30) | No |



## UseCases

- [Two-Factor Authentication]: Generate time-based one-time passwords for secure authentication workflows and MFA implementations
- [Security Token Generation]: Create TOTP codes for accessing secured systems and services that require time-based authentication
- [Automated Login Workflows]: Generate TOTP codes programmatically for automated testing and system integration scenarios

