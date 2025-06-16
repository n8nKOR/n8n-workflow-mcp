# ReadBinaryFiles

## Overview

The Read Binary Files node is a **deprecated** node that reads binary files from disk based on file patterns. This node has been hidden from the node palette and its functionality has been superseded by more modern file handling nodes. It was designed to read multiple files matching a pattern and convert them into binary data for workflow processing.

**⚠️ DEPRECATION NOTICE**: This node is deprecated and hidden from the node palette. For new implementations, consider using the **Read Binary File** node or other modern file handling nodes.

## Credentials

The Read Binary Files node does not require credentials as it operates on local file system access:

- **File System Access**: Requires appropriate file system permissions to read files
- **Local File Access**: Operates on files accessible to the n8n instance

## Inputs

- **Main**: Input data from previous nodes (used for workflow context)

## Outputs

- **Main**: Returns binary data for each file found matching the pattern

## Properties

### Resource: File Operations

#### Operation: Binary File Reading

| Field Name | Type | Description | Required |
|---|---|---|---|
| File Selector | String | Pattern for files to read (e.g., "*.jpg", "/path/to/files/*") | Yes |
| Property Name | String | Name of the binary property to store the file data | Yes |

## UseCases

- **Batch File Processing** : Process multiple files matching a specific pattern
- **Image Processing** : Read image files for batch processing workflows
- **Document Processing** : Load multiple documents for analysis or conversion
- **File Migration** : Transfer files from one system to another
- **Archive Processing** : Extract and process files from directories

### Document Processing

1. **PDF Processing**: Read PDF files for text extraction or conversion
2. **Image Processing**: Load images for analysis, transformation, or upload
3. **Archive Processing**: Read archive files for extraction or analysis
4. **Document Conversion**: Load documents for format conversion
5. **Batch Processing**: Process multiple files simultaneously

### Data Migration

1. **File System Migration**: Transfer files between systems
2. **Cloud Upload**: Prepare files for cloud storage upload
3. **Database Storage**: Store files as binary data in databases
4. **Content Management**: Import files into content management systems
5. **Backup Operations**: Read files for backup and archiving

### Media Processing

1. **Image Galleries**: Load images for gallery processing
2. **Video Processing**: Read video files for processing or streaming
3. **Audio Processing**: Load audio files for analysis or conversion
4. **Thumbnail Generation**: Read images to generate thumbnails
5. **Media Optimization**: Load media files for optimization

### Integration Workflows

1. **API Uploads**: Read files to upload via APIs
2. **Email Attachments**: Attach files to emails
3. **Form Processing**: Handle file uploads from forms
4. **Report Generation**: Include files in generated reports
5. **Content Delivery**: Prepare files for content delivery networks

### Automation Tasks

1. **File Monitoring**: Process files from monitored directories
2. **Scheduled Processing**: Batch process files on schedule
3. **Workflow Triggers**: Use file reading as workflow triggers
4. **Data Validation**: Read files to validate content
5. **Quality Assurance**: Process files for quality checks

### Development and Testing

1. **Test Data**: Load test files for application testing
2. **Development Assets**: Read development resources
3. **Configuration Files**: Load configuration files
4. **Template Processing**: Read template files for processing
5. **Resource Management**: Manage application resources

