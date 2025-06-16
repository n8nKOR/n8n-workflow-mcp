# Split In Batches Node

## Overview

Split data into batches and iterate over each batch in a controlled loop. This node is essential for processing large datasets in manageable chunks, preventing memory overflow and API rate limiting issues. It maintains state between executions to track progress through the dataset and provides two outputs: one for completed processing and another for continuing the loop. Note that n8n nodes automatically run once for each input item, so this node is typically used for specific batch processing scenarios.

⚠️ **Version Information**: This node has multiple versions (V1, V2, V3). The current default version is V3, which provides enhanced batch processing capabilities, improved memory management, and better error handling. Some features may differ between versions.

## Credentials

None

## Inputs

- Main

## Outputs

- done (first output) - Returns all processed items when the loop is complete
- loop (second output) - Returns current batch items for processing in the loop

## Properties

### Resource: Loop Control

#### Operation: Split in Batches

| Field Name | Type | Description | Required |
|---|---|---|---|
| Batch Size | number | The number of items to return with each call. Minimum value is 1 | Yes |

### Options

| Field Name | Type | Description | Default | Required |
|---|---|---|---|---|
| Reset | boolean | Whether the node starts again from the beginning of input items. Treats incoming data as a new set rather than continuing with previous items | false | No |

#### Additional Information
- The node maintains internal state to track current position in the dataset
- Uses context storage to preserve items between executions
- Automatically calculates the maximum number of runs needed based on batch size
- Provides proper paired item information for workflow tracing
- Handles source data overwriting for complex workflow scenarios
- Returns empty arrays when no items are left to process

#### Loop Mechanics
1. **First Execution**: Initializes context, splits input into batches, returns first batch via 'loop' output
2. **Subsequent Executions**: Returns next batch via 'loop' output until all items are processed
3. **Final Execution**: Returns all processed items via 'done' output when loop is complete

## UseCases

- **Large Dataset Vector Processing**: Process thousands of images or documents in batches for vector embedding generation, preventing memory overflow and API rate limit issues while maintaining processing efficiency
- **Bulk API Operations**: Split large datasets into manageable batches for API calls, database insertions, or file processing operations while maintaining system stability and respecting rate limits
- **Machine Learning Data Preparation**: Process large machine learning datasets in batches for feature extraction, data transformation, or model training preparation without overwhelming system resources
- **Email Campaign Processing**: Send bulk emails in batches to avoid spam filters and respect email service provider rate limits while maintaining delivery reliability and tracking
- **Database Migration**: Migrate large amounts of data between systems in controlled batches, allowing for error handling, progress monitoring, and system stability during transfers
- **Social Media Content Processing**: Process large volumes of social media posts, comments, or images in batches for sentiment analysis, content moderation, or data extraction workflows
- **File Processing Workflows**: Handle large numbers of files in batches for conversion, analysis, or transformation operations without overwhelming file system or processing resources
- **ETL Pipeline Processing**: Extract, transform, and load large datasets in controlled batches to maintain data integrity and system performance during complex data operations

