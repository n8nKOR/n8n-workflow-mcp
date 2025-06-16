# Mindee Node

## Overview

The Mindee node provides AI-powered document analysis and data extraction capabilities using Mindee's intelligent document processing API. This node can automatically extract structured data from various document types including receipts, invoices, passports, driver's licenses, and custom documents using state-of-the-art OCR and machine learning technologies.

## Credentials

- **Name**: mindeeReceiptApi, Required: Yes

To obtain Mindee API credentials:
1. Sign up for a Mindee account at https://platform.mindee.com
2. Navigate to the API Keys section in your dashboard
3. Generate a new API key for your application
4. Copy the API key for use in n8n credentials

## Inputs

- **Main**: Input containing binary document data or file references

## Outputs

- **Main**: Extracted document data with confidence scores and metadata

## Properties

### Resource: Document Analysis

#### Operation: Predict

| Field Name | Type | Description | Required |
|---|---|---|---|
| API Version | Options | Which Mindee API Version to use (1/3/4) | Yes |
| Resource | Options | Type of document to analyze (Invoice/Receipt) | Yes |
| Operation | Options | Operation to perform (Predict) | Yes |
| Input Binary Field | String | Name of input binary field containing file | Yes |
| RAW Data | Boolean | Return data exactly as received from API | No |

## UseCases

- **Receipt Processing**: Extract merchant info, amounts, and line items from receipts
- **Invoice Analysis**: Parse vendor details, amounts, and payment terms from invoices
- **Document Digitization**: Convert paper documents to structured digital data
- **Expense Management**: Automate expense report creation from receipt images
- **Financial Data Extraction**: Extract key financial information for accounting systems

## Operations

### Analyze Document

Extract structured data from uploaded documents using AI-powered OCR.

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Binary Property | String | Property name containing document data | Yes | data |
| Document Source | Options | Source of document data (binary/URL) | No | binary |
| Document URL | String | URL of document to analyze (if source=URL) | Conditional | - |
| Include Cropped | Boolean | Include cropped field images in response | No | false |
| Include Full Text | Boolean | Include full OCR text in response | No | false |

### Analysis Configuration

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Language | Options | Document language (auto/en/fr/es/de/pt/it) | No | auto |
| Page Range | String | Specific pages to analyze (e.g., "1-3,5") | No | all |
| DPI Enhancement | Boolean | Apply DPI enhancement for low-quality images | No | true |
| Orientation Correction | Boolean | Automatically correct document orientation | No | true |
| Confidence Threshold | Number | Minimum confidence score (0.0-1.0) | No | 0.7 |

### Response Format Options

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Output Format | Options | Response data format (json/structured) | No | structured |
| Include Raw OCR | Boolean | Include raw OCR text in response | No | false |
| Include Geometry | Boolean | Include bounding box coordinates | No | false |
| Normalize Dates | Boolean | Standardize date formats | No | true |
| Normalize Amounts | Boolean | Standardize currency amounts | No | true |

## Document-Specific Fields

### Receipt Analysis Fields

| Field | Type | Description | Confidence |
|---|---|---|---|
| Merchant Name | String | Store or business name | High |
| Merchant Address | Object | Full business address | Medium |
| Transaction Date | Date | Purchase date and time | High |
| Total Amount | Number | Final total paid | High |
| Tax Amount | Number | Total tax charged | Medium |
| Tip Amount | Number | Tip/gratuity amount | Medium |
| Line Items | Array | Individual products/services | High |
| Payment Method | String | Payment type (cash/card/digital) | Low |
| Receipt Number | String | Transaction/receipt ID | Medium |
| Category | String | Business category classification | Medium |

### Invoice Analysis Fields

| Field | Type | Description | Confidence |
|---|---|---|---|
| Vendor Name | String | Supplier business name | High |
| Vendor Address | Object | Supplier address details | High |
| Customer Name | String | Billing customer name | High |
| Customer Address | Object | Billing address | High |
| Invoice Number | String | Unique invoice identifier | High |
| Invoice Date | Date | Invoice issue date | High |
| Due Date | Date | Payment due date | High |
| Subtotal | Number | Pre-tax amount | High |
| Tax Amount | Number | Total tax amount | High |
| Total Amount | Number | Final amount due | High |
| Line Items | Array | Product/service details | High |
| Payment Terms | String | Payment conditions | Medium |
| Purchase Order | String | PO reference number | Low |

### Passport Analysis Fields

| Field | Type | Description | Confidence |
|---|---|---|---|
| Document Type | String | Passport/travel document type | High |
| Country Code | String | Issuing country | High |
| Document Number | String | Passport number | High |
| Surname | String | Last name | High |
| Given Names | String | First and middle names | High |
| Nationality | String | Citizenship | High |
| Date of Birth | Date | Birth date | High |
| Place of Birth | String | Birth location | Medium |
| Gender | String | Gender designation | High |
| Issue Date | Date | Document issue date | High |
| Expiry Date | Date | Document expiration | High |
| Personal Number | String | National ID number | Medium |
| MRZ | Object | Machine readable zone data | High |

### Driver's License Fields

| Field | Type | Description | Confidence |
|---|---|---|---|
| License Number | String | Driver's license ID | High |
| Full Name | String | License holder name | High |
| Address | Object | Residential address | High |
| Date of Birth | Date | Birth date | High |
| Issue Date | Date | License issue date | High |
| Expiry Date | Date | License expiration | High |
| License Class | String | Vehicle class authorization | High |
| Restrictions | Array | Driving restrictions | Medium |
| Endorsements | Array | Special endorsements | Medium |
| Photo | String | Base64 encoded photo | Low |
| Signature | String | Base64 encoded signature | Low |

## Advanced Features

### Confidence Scoring
- **Field-level confidence**: Individual confidence scores for each extracted field
- **Document-level confidence**: Overall document analysis reliability
- **Threshold filtering**: Filter results based on minimum confidence levels
- **Quality assessment**: Document quality and readability scores

### Multi-page Processing
- **Batch processing**: Analyze multiple pages in a single request
- **Page-specific extraction**: Target specific pages for analysis
- **Document splitting**: Separate multi-document files
- **Cross-page validation**: Validate data consistency across pages

### Data Validation
- **Format validation**: Check data format consistency (dates, amounts, etc.)
- **Business rules**: Apply custom validation logic
- **Cross-field validation**: Verify relationships between fields
- **Duplicate detection**: Identify potential duplicate documents

### Error Handling
- **Graceful degradation**: Partial results when some fields fail
- **Retry logic**: Automatic retry for temporary failures
- **Error classification**: Detailed error types and recovery suggestions
- **Quality warnings**: Alerts for low-quality input documents

## API Response Structure

### Standard Response
```json
{
  "api_request": {
    "status": "success",
    "status_code": 200,
    "url": "https://api.mindee.net/v1/products/..."
  },
  "document": {
    "id": "doc_id_12345",
    "filename": "receipt.pdf",
    "total_pages": 1,
    "inference": {
      "pages": [
        {
          "id": 0,
          "orientation": {"value": 0},
          "prediction": {
            "merchant_name": {
              "value": "Walmart",
              "confidence": 0.95
            },
            "total_amount": {
              "value": 45.67,
              "confidence": 0.98
            }
          }
        }
      ],
      "prediction": {
        "merchant_name": {
          "value": "Walmart", 
          "confidence": 0.95
        },
        "total_amount": {
          "value": 45.67,
          "confidence": 0.98
        }
      }
    }
  }
}
```

### Error Response
```json
{
  "api_request": {
    "status": "failure",
    "status_code": 400,
    "error": {
      "message": "Invalid document format",
      "details": "Supported formats: PDF, JPEG, PNG, WEBP, TIFF, HEIC"
    }
  }
}
```

