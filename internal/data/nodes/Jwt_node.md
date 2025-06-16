# JWT Node

## Overview

JWT (JSON Web Token) operations for signing, verifying, and decoding tokens

## Credentials

- Name: jwtAuth, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: JWT Operations

#### Operation: Sign

| Field Name | Type | Description | Required |
|---|---|---|---|
| Use JSON to Build Payload | boolean | Whether to use JSON to build the claims | No |
| Audience | string | Identifies the recipients that the JWT is intended for | No |
| Expires In | number | The lifetime of the token in seconds | No |
| Issuer | string | Identifies the principal that issued the JWT | No |
| JWT ID | string | Unique identifier for the JWT | No |
| Not Before | number | Time before which JWT must not be accepted | No |
| Subject | string | Identifies the principal that is the subject of the JWT | No |
| Payload Claims (JSON) | json | Claims to add to the token in JSON format (when using JSON format) | No |
| Override Algorithm | options | Algorithm to use for signing (HS256, HS384, HS512, RS256, RS384, RS512, PS256, PS384, PS512, ES256, ES384, ES512) | No |

#### Operation: Verify

| Field Name | Type | Description | Required |
|---|---|---|---|
| Token | string | The JWT token to verify | Yes |
| Return Additional Info | boolean | Return complete token info vs just payload | No |
| Ignore Expiration | boolean | Whether to ignore token expiration | No |
| Ignore Not Before Claim | boolean | Whether to ignore the not before claim | No |
| Clock Tolerance | number | Seconds to tolerate for clock differences | No |
| Key ID | string | The kid (key ID) claim for signature validation | No |
| Override Algorithm | options | Algorithm to use for verifying | No |

#### Operation: Decode

| Field Name | Type | Description | Required |
|---|---|---|---|
| Token | string | The JWT token to decode | Yes |
| Return Additional Info | boolean | Return complete token info vs just payload | No |

### JWT Claims

#### Standard Claims
- **iss** (Issuer): Principal that issued the JWT
- **sub** (Subject): Principal that is the subject of the JWT
- **aud** (Audience): Recipients that the JWT is intended for
- **exp** (Expiration): Time after which the JWT expires
- **nbf** (Not Before): Time before which the JWT must not be accepted
- **iat** (Issued At): Time at which the JWT was issued
- **jti** (JWT ID): Unique identifier for the JWT

#### Custom Claims
Any additional claims can be added through JSON format:
```json
{
  "user_id": "12345",
  "role": "admin",
  "permissions": ["read", "write"],
  "department": "engineering"
}
```

### Supported Algorithms

#### HMAC (Symmetric)
- **HS256**: HMAC using SHA-256 hash algorithm
- **HS384**: HMAC using SHA-384 hash algorithm  
- **HS512**: HMAC using SHA-512 hash algorithm

#### RSA (Asymmetric)
- **RS256**: RSASSA-PKCS1-v1_5 using SHA-256
- **RS384**: RSASSA-PKCS1-v1_5 using SHA-384
- **RS512**: RSASSA-PKCS1-v1_5 using SHA-512

#### RSA-PSS (Asymmetric)
- **PS256**: RSASSA-PSS using SHA-256 and MGF1 with SHA-256
- **PS384**: RSASSA-PSS using SHA-384 and MGF1 with SHA-384
- **PS512**: RSASSA-PSS using SHA-512 and MGF1 with SHA-512

#### ECDSA (Asymmetric)
- **ES256**: ECDSA using P-256 and SHA-256
- **ES384**: ECDSA using P-384 and SHA-384
- **ES512**: ECDSA using P-521 and SHA-512

### Credential Configuration

#### Key Types
- **Passphrase**: Use secret string for HMAC algorithms
- **PEM Key**: Use public/private key pair for RSA/ECDSA algorithms

#### Key Format
- **Secret**: String for symmetric algorithms
- **Private Key**: PEM formatted private key for signing
- **Public Key**: PEM formatted public key for verification

### Signing Process

#### Steps
1. **Create Header**: Algorithm and key ID information
2. **Build Payload**: Add claims and custom data
3. **Generate Signature**: Sign header and payload
4. **Encode Token**: Create final JWT string

#### Example Process
```
Header: {"alg": "HS256", "typ": "JWT"}
Payload: {"sub": "user123", "exp": 1672531200}
Signature: HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)
Token: <base64header>.<base64payload>.<base64signature>
```

### Verification Process

#### Steps
1. **Parse Token**: Extract header, payload, and signature
2. **Verify Algorithm**: Check algorithm matches expectation
3. **Validate Signature**: Verify signature using appropriate key
4. **Check Claims**: Validate expiration, not before, etc.
5. **Return Result**: Success or failure with details

#### Clock Tolerance
Handle small time differences between servers:
```
Clock Tolerance: 30 (seconds)
Allows 30 second variance for exp/nbf claims
```

### Error Handling

#### Common Errors
- **Invalid Token Format**: Malformed JWT structure
- **Invalid Signature**: Signature verification failed
- **Token Expired**: Current time after expiration
- **Token Not Active**: Current time before not-before time
- **Invalid Algorithm**: Algorithm mismatch or not allowed
- **Missing Claims**: Required claims not present

#### Error Messages
```
"The JWT token was not provided"
"Token expired"
"Token not active"
"Invalid signature"
"Invalid token format"
```

### Security Considerations

#### Algorithm Selection
- **HMAC**: Shared secret, simpler but requires secret sharing
- **RSA/ECDSA**: Public/private key pairs, more secure for distributed systems
- **Key Size**: Use appropriate key sizes (256+ bits for HMAC, 2048+ bits for RSA)

#### Best Practices
- **Short Expiration**: Use reasonable expiration times
- **Secure Storage**: Protect private keys and secrets
- **Algorithm Validation**: Verify algorithm matches expectations
- **Claim Validation**: Validate all relevant claims

## UseCases

#### Authentication Token
```
Operation: Sign
Claims:
- Subject: user123
- Issuer: auth-service
- Audience: app-service
- Expires In: 3600
Custom Claims:
- role: admin
- permissions: ["read", "write"]
```

#### API Access Token
```
Operation: Sign
Claims:
- Issuer: api-gateway
- Subject: service-account
- Audience: backend-api
- Expires In: 7200
Custom Claims:
- scope: ["users", "orders"]
```

#### Token Verification
```
Operation: Verify
Token: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9...
Options:
- Clock Tolerance: 60
- Ignore Expiration: false
```

#### Token Inspection
```
Operation: Decode
Token: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9...
Options:
- Return Additional Info: true
```

### Integration Patterns

#### Authentication Flow
1. **User Login**: Validate credentials
2. **Token Generation**: Sign JWT with user claims
3. **Token Response**: Return token to client
4. **Request Validation**: Verify token on subsequent requests

#### API Gateway Pattern
1. **Token Validation**: Verify incoming JWT tokens
2. **Claims Extraction**: Extract user/permission information
3. **Authorization**: Check permissions for requested resources
4. **Request Forwarding**: Forward validated requests to backend

### Response Formats

#### Successful Sign
```json
{
  "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1c2VyMTIzIiwiZXhwIjoxNjcyNTMxMjAwfQ.signature"
}
```

#### Successful Verify (Simple)
```json
{
  "sub": "user123",
  "exp": 1672531200,
  "iat": 1672527600
}
```

#### Successful Verify (Complete)
```json
{
  "header": {"alg": "HS256", "typ": "JWT"},
  "payload": {"sub": "user123", "exp": 1672531200},
  "signature": "signature_string"
}
```

#### Successful Decode
```json
{
  "header": {"alg": "HS256", "typ": "JWT"},
  "payload": {"sub": "user123", "exp": 1672531200}
}
```

### Performance Considerations

#### Algorithm Performance
- **HMAC**: Fastest for small to medium tokens
- **RSA**: Slower but more secure for distributed systems
- **ECDSA**: Good balance of security and performance

#### Key Management
- **Cache Keys**: Cache public keys for verification
- **Key Rotation**: Plan for regular key rotation
- **Key Storage**: Use secure key storage solutions

### Migration and Compatibility

#### Token Migration
- **Algorithm Upgrades**: Plan for algorithm migrations
- **Key Rotation**: Handle key rotation gracefully
- **Backward Compatibility**: Support multiple algorithms during transitions

#### Standards Compliance
- **RFC 7519**: JWT standard compliance
- **RFC 7515**: JWS (JSON Web Signature) compliance
- **RFC 7517**: JWK (JSON Web Key) support where applicable

### Usage Notes
- Supports all major JWT algorithms (HMAC, RSA, ECDSA)
- Flexible payload construction with collection or JSON format
- Complete token information available with additional info option
- Clock tolerance handles time synchronization issues
- Secure credential management with multiple key type support
- Industry-standard JWT library (jsonwebtoken) for reliability
- Comprehensive error handling and validation
- Tool-compatible for use in workflow automations
