# Compression Node

## Overview

Handles file compression and decompression operations for ZIP and GZIP formats with batch processing and automatic optimization. This node uses the fflate library for efficient compression operations and supports multiple binary fields for batch processing. Essential for file archiving, data transfer optimization, and storage space management in automation workflows.

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: File Compression

#### Operation: Compress

| Field Name | Type | Description | Required |
|---|---|---|---|
| Input Binary Field(s) | string | The name of the input binary field(s) containing the file(s) to be compressed. Use comma-separated list for multiple files | Yes |
| Output Format | options | Format of the output compressed file (zip, gzip) | Yes |
| File Name | string | Name of the output compressed file (for ZIP format) | Conditional |
| Put Output File in Field | string | The name of the output binary field to put the compressed file in | No |

#### Operation: Decompress

| Field Name | Type | Description | Required |
|---|---|---|---|
| Input Binary Field(s) | string | The name of the input binary field(s) containing the file(s) to decompress. Use comma-separated list for multiple files | Yes |
| Output Prefix | string | Prefix to add to the decompressed files | No |

## Supported Compression Formats

### Input Formats (Decompression)
- **Zip:** `.zip` files - Extracts all files within the archive
- **Gzip:** `.gz`, `.gzip` files - Decompresses single files

### Output Formats (Compression)
- **Zip:** Creates `.zip` archives containing multiple files
- **Gzip:** Creates `.gz` files for individual file compression

## File Handling Behavior

### Compression Optimization
The node automatically optimizes compression based on file types. Files with extensions that are already compressed use compression level 0 (store only), while other files use compression level 6:

**Already Compressed File Types:** 7z, aifc, bz2, doc, docx, gif, gz, heic, heif, jpg, jpeg, mov, mp3, mp4, pdf, png, ppt, pptx, rar, webm, webp, xls, xlsx, zip

### Decompression Handling
- **Zip Files:** Extracts all files and assigns them to separate binary fields with the specified output prefix
- **Gzip Files:** Decompresses individual files and attempts to determine the original file extension and MIME type
- **macOS Metadata:** Automatically filters out `__MACOSX` metadata files during zip extraction

### Binary Field Management
- **Input:** Accepts single or multiple binary fields (comma-separated list)
- **Output:** Creates appropriately named binary fields based on operation and configuration
- **File Naming:** Preserves original filenames and extensions where possible
- **MIME Type Detection:** Uses mime-types library for accurate content type identification

#### Additional Information
- Uses fflate library for efficient compression/decompression operations
- Supports batch processing of multiple files simultaneously
- Automatically handles different compression levels based on file types
- Provides proper error handling for corrupted or invalid archives
- Maintains file metadata and structure during compression/decompression

## UseCases

- **Database Backup Compression**: Compress SQLite database files and other backup data into ZIP archives for efficient storage and transfer, reducing storage costs and transfer times
- **Log File Archiving**: Compress application log files using GZIP for long-term storage, maintaining accessibility while significantly reducing storage requirements
- **Document Package Creation**: Create ZIP packages containing multiple documents, images, and files for client delivery or internal distribution workflows
- **Data Export Optimization**: Compress large CSV, JSON, or XML data exports before sending via email or uploading to cloud storage, improving transfer efficiency
- **Media File Organization**: Package related media files (images, videos, documents) into organized ZIP archives for project delivery or client presentation
- **Configuration Backup**: Compress configuration files and settings into portable archives for system backup and deployment automation
- **Report Distribution**: Package generated reports, charts, and supporting documents into ZIP files for automated distribution to stakeholders
- **Archive Processing**: Decompress received ZIP or GZIP files to extract content for further processing in automation workflows
- **Batch File Processing**: Compress multiple processed files into single archives for organized storage and simplified file management
- **Email Attachment Optimization**: Compress large attachments before sending emails to comply with size limits and improve delivery reliability

