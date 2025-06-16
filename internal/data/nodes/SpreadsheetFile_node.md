# Spreadsheet File Node

## Overview

Reads and writes data from a spreadsheet file like CSV, XLS, ODS, etc

⚠️ **Version Information**: This node has multiple versions (V1, V2). The current default version is V2, which provides enhanced file format support, improved data processing capabilities, and better error handling. Some features may differ between versions.

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: File Operations

#### Operation: Read From File

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform. Options: "Read From File", "Write to File" | Yes |
| Input Binary Field | String | The name of the input field containing the file data to be processed | Yes |

#### Operation: Write to File

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Action to perform | Yes |
| File Format | Options | Format of the output file. Options: CSV, HTML, ODS, RTF, XLS, XLSX | Yes |
| Put Output File in Field | String | The name of the output binary field to put the file in | Yes |

### Options (Write to File)
| Field Name | Type | Description | Required |
|---|---|---|---|
| Compression | Boolean | Whether compression will be applied (for XLSX, ODS files) | No |
| File Name | String | File name to set in binary data | No |
| Header Row | Boolean | Whether the first row contains header names | No |
| Sheet Name | String | Name of the sheet to create (for ODS, XLS, XLSX files) | No |

### Options (Read From File) 
| Field Name | Type | Description | Required |
|---|---|---|---|
| Delimiter | String | Field delimiter for CSV files | No |
| Encoding | Options | Text encoding for CSV files (ASCII, Latin1, UTF-8, etc.) | No |
| Header Row | Boolean | Whether the first row contains header names | No |
| Include Empty Cells | Boolean | Whether to include empty cells in the output | No |
| Range | String | Range of cells to read (e.g., A1:D10) | No |
| Sheet Name | String | Name of the sheet to read from | No |

## UseCases

- **Data Import and Export** : Import data from Excel files and export processed data to spreadsheets
- **Report Generation** : Generate Excel reports from database queries and API responses
- **Data Analysis** : Analyze spreadsheet data and perform calculations on imported datasets
- **Batch Data Processing** : Process large volumes of data from Excel files in automated workflows
- **Financial Reporting** : Create financial reports and statements in Excel format
- **Inventory Management** : Import and export inventory data from Excel spreadsheets
- **Customer Data Migration** : Migrate customer data between systems using Excel as intermediary
- **Survey Data Processing** : Process survey responses and feedback data from Excel files
- **Template-Based Reporting** : Generate reports using predefined Excel templates
- **Data Validation** : Validate and clean data imported from Excel spreadsheets

