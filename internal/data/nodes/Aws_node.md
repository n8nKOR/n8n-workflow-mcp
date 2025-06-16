# AWS

## Overview

The AWS node collection provides comprehensive integration with Amazon Web Services, enabling automation across 15+ AWS services. This robust implementation includes compute, storage, messaging, AI/ML, security, and database services, making it one of the most complete cloud service integrations in n8n. Each AWS service is implemented as a dedicated node with specialized operations and features.

**Supported Services**: Lambda, SNS, S3, DynamoDB, SES, SQS, Cognito, IAM, Rekognition, Textract, Comprehend, Transcribe, Certificate Manager, ELB, and more.

## Credentials

All AWS nodes use unified authentication with the following credential configuration:

- **Credential Name**: `aws`
- **Required Fields**: 
  - **Access Key ID**: Your AWS access key ID
  - **Secret Access Key**: Your AWS secret access key
  - **Region**: AWS region for API calls (e.g., us-east-1, eu-west-1)

## Inputs

- **Main Input**: Data to be processed by the AWS service, varies by service and operation

## Outputs

- **Main Output**: Service-specific response data from AWS APIs, including operation results and metadata

## Properties

### AWS Lambda

#### Resource: Lambda Functions

##### Operation: Invoke

| Field Name | Type | Description | Required |
|---|---|---|---|
| Function Name or ID | Options/String | Select Lambda function from list or enter ARN | Yes |
| Qualifier | String | Function version or alias (default: $LATEST) | Yes |
| Invocation Type | Options | RequestResponse (sync) or Event (async) | No |
| JSON Input | String | JSON payload to pass to Lambda function | No |

### AWS SNS (Simple Notification Service)

#### Resource: Topics

##### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Name | String | Topic name to create | Yes |
| Display Name | String | Display name for SMS subscriptions | No |
| Fifo Topic | Boolean | Create FIFO topic for ordered messages | No |

##### Operation: Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Topic | Resource Locator | Select topic to delete (List/URL/ARN) | Yes |

##### Operation: Publish

| Field Name | Type | Description | Required |
|---|---|---|---|
| Topic | Resource Locator | Select topic to publish message to | Yes |
| Subject | String | Email endpoint subject line | Yes |
| Message | String | Message content to send | Yes |

### AWS S3 (Simple Storage Service)

#### Resource: Files

##### Operations: Upload, Download, Delete, Copy, Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| Bucket Name | Resource Locator | S3 bucket selection | Yes |
| File Key | String | Object key/path in bucket | Yes |
| Binary Data | Boolean | Handle binary file data | No |

### AWS DynamoDB

#### Resource: Items

##### Operations: Upsert, Get, Get All, Delete

| Field Name | Type | Description | Required |
|---|---|---|---|
| Table | Resource Locator | DynamoDB table selection | Yes |
| Expression Attribute Names | String | Attribute name mappings | No |
| Condition Expression | String | Conditional operation expression | No |

### AWS SES (Simple Email Service)

#### Resource: Email

##### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| To | String | Recipient email addresses | Yes |
| Subject | String | Email subject line | Yes |
| Body | String | Email body content | Yes |
| Is Body HTML | Boolean | HTML formatted email body | No |

### AWS SQS (Simple Queue Service)

#### Resource: Messages

##### Operation: Send

| Field Name | Type | Description | Required |
|---|---|---|---|
| Queue | Resource Locator | SQS queue selection | Yes |
| Message Body | String | Message content | Yes |
| Message Attributes | Collection | Custom message attributes | No |

### AWS Rekognition

#### Resource: Image Analysis

##### Operations: Detect Labels, Faces, Text, Celebrities, Moderation

| Field Name | Type | Description | Required |
|---|---|---|---|
| Binary Data | Boolean | Use binary data from previous node | Yes |
| Max Labels | Number | Maximum labels to return | No |
| Min Confidence | Number | Minimum confidence threshold | No |

### AWS Textract

#### Resource: Document Analysis

##### Operations: Analyze Receipt, Analyze ID

| Field Name | Type | Description | Required |
|---|---|---|---|
| Binary Data | Boolean | Document binary data input | Yes |
| Simplify Output | Boolean | Return simplified, structured data | No |

## UseCases

- **Serverless Computing**: Execute Lambda functions for scalable, event-driven processing
- **Notification Systems**: Build comprehensive alert systems using SNS for email, SMS, and webhooks
- **File Storage & Processing**: Store, retrieve, and process files using S3 with automated workflows
- **Database Operations**: Perform NoSQL operations on DynamoDB for scalable data storage
- **Email Automation**: Send transactional and marketing emails through SES integration
- **Message Queuing**: Implement reliable message queuing with SQS for decoupled architectures
- **AI/ML Integration**: Process images and documents using Rekognition and Textract for content analysis
- **User Management**: Manage authentication and authorization with Cognito integration
- **Infrastructure Automation**: Automate AWS resource management and deployment processes
- **Data Analytics**: Build data pipelines combining multiple AWS services for comprehensive analytics
- **Security & Compliance**: Implement security workflows using IAM and Certificate Manager
- **Real-time Processing**: Create event-driven architectures using SNS triggers and Lambda
- **Content Moderation**: Automatically moderate user-generated content using AI services
- **Document Processing**: Extract and process information from receipts, invoices, and IDs
- **Voice Processing**: Convert speech to text using Transcribe for audio content analysis 