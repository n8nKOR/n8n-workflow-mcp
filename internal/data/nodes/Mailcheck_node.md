# Mailcheck

## Overview

The Mailcheck node allows you to validate email addresses using the Mailcheck.co API service. This node provides email verification capabilities to check the validity, deliverability, and quality of email addresses within your n8n workflows. Mailcheck helps reduce bounce rates and improve email campaign effectiveness by identifying invalid, risky, or undeliverable email addresses.

## Credentials

This node requires Mailcheck API credentials:
- **Credential Name**: `mailcheckApi`
- **Required Fields**: 
  - API Key: Your Mailcheck.co API key

To obtain API credentials:
1. Sign up for a Mailcheck.co account at https://mailcheck.co
2. Navigate to your API settings or dashboard
3. Generate an API key
4. Copy the API key for use in n8n

## Inputs

- **Main**: Email addresses to validate (can be single item or array)

## Outputs

- **Main**: Returns JSON data containing email validation results with status, deliverability score, and detailed analysis

## Properties

### Resource: Email

#### Operation: Check

| Field Name | Type | Description | Required |
|---|---|---|---|
| Email | String | Email address to validate and check | Yes |

**Email Validation Process:**
- Syntax validation
- Domain verification
- MX record checks
- Mailbox existence validation
- Risk assessment scoring
- Spam trap detection

**Response Fields:**
- `result`: Overall validation result (valid/invalid/unknown)
- `score`: Deliverability score (0-100)
- `reason`: Detailed reason for the validation result
- `disposable`: Whether the email is from a disposable email provider
- `mx`: MX record information
- `smtp`: SMTP server verification results

## UseCases

- **Email List Cleaning**: Validate email lists before marketing campaigns to reduce bounce rates and improve deliverability
- **User Registration Validation**: Verify email addresses during user signup to ensure valid contact information
- **Lead Quality Assessment**: Score and validate leads based on email quality to prioritize sales efforts
- **Newsletter Subscription Management**: Clean subscriber lists by removing invalid or risky email addresses
- **E-commerce Customer Validation**: Verify customer email addresses during checkout to prevent order delivery issues
- **CRM Data Quality**: Maintain clean contact databases by validating existing email records
- **Form Submission Filtering**: Validate email addresses from web forms to prevent spam and fake submissions
- **Marketing Automation**: Integrate email validation into automated marketing workflows for better campaign performance
- **Compliance and Deliverability**: Ensure email marketing compliance by maintaining clean, validated recipient lists
- **Data Migration Cleaning**: Validate email addresses when migrating contact data between systems
