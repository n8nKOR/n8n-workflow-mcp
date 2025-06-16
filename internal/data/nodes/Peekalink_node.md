# Peekalink

## Overview

The Peekalink node provides integration with the Peekalink API, a service that generates rich link previews and metadata extraction from URLs. This node enables automated extraction of webpage information including titles, descriptions, images, and other metadata, making it ideal for content curation, social media automation, and link validation workflows.

## Credentials

- **Credential Name**: `peekalinkApi`
- **Required**: Yes
- **Configuration**: 
  - API Key: Your Peekalink API key obtained from the Peekalink dashboard

## Inputs

- **Main**: Input data from previous nodes, which can contain URLs to process or other workflow data

## Outputs

- **Main**: Returns extracted metadata and preview information for the specified URLs

## Properties

### Resource: Link Preview

#### Operation: Preview Generation

| Field Name | Type | Description | Required |
|---|---|---|---|
| Operation | Options | Choose between 'Preview' (return preview for a link) or 'Is Available' (check if preview is available) | Yes |
| URL | String | The URL to generate preview for or check availability | Yes |

## UseCases

- **Peekalink Integration** : Use Peekalink for workflow automation and data processing
- **Data Processing** : Process and transform data using Peekalink capabilities
- **Workflow Enhancement** : Enhance workflows with Peekalink functionality
- **Automation Tasks** : Automate repetitive tasks using Peekalink features

## API Integration

The node communicates with the Peekalink API:
- **Base URL**: `https://api.peekalink.io`
- **Preview Endpoint**: `https://api.peekalink.io` (POST)
- **Availability Endpoint**: `https://api.peekalink.io/is-available/` (POST)
- **Authentication**: API key via `X-API-Key` header
- **Request Format**: JSON with `link` parameter

## Response Data Structure

### Preview Operation Response
When using the "Preview" operation, the response typically includes:

- **Title**: The webpage title or Open Graph title
- **Description**: Meta description or Open Graph description
- **Image**: Preview image URL (Open Graph image, Twitter card image, or first significant image)
- **URL**: Canonical URL of the page
- **Domain**: Domain name of the URL
- **Favicon**: Favicon URL for the website
- **Type**: Content type (article, website, video, etc.)
- **Site Name**: Name of the website or publication
- **Author**: Author information if available
- **Published Time**: Publication date if available
- **Modified Time**: Last modification date if available

### Is Available Operation Response
When using the "Is Available" operation, the response indicates:
- **Available**: Boolean indicating if preview generation is possible
- **Reason**: Explanation if preview is not available
- **Status**: HTTP status information

## Error Handling

The node implements robust error handling:
- **API Errors**: Detailed error messages from Peekalink API
- **Authentication Errors**: Clear indication of API key issues
- **URL Validation**: Handling of invalid or inaccessible URLs
- **Rate Limiting**: Proper handling of API rate limits
- **Continue on Fail**: Option to continue workflow execution even if preview extraction fails

## Use Cases

### Content Management
- **CMS Integration**: Automatically populate link metadata when URLs are added to content
- **Blog Post Enhancement**: Extract preview data for external links in blog posts
- **Content Validation**: Verify that linked content matches expected descriptions
- **SEO Analysis**: Extract and analyze meta information for SEO auditing

### Social Media Automation
- **Rich Link Previews**: Generate preview cards for shared links on social platforms
- **Social Media Posting**: Enhance automated posts with rich link previews
- **Content Curation**: Automatically extract summaries for curated content lists
- **Link Sharing**: Improve link sharing experiences with preview data

### Marketing and Analytics
- **Campaign Validation**: Verify landing page metadata for marketing campaigns
- **Competitor Analysis**: Extract metadata from competitor websites
- **Content Research**: Gather information about trending content and articles
- **Link Quality Assessment**: Evaluate the quality and relevance of external links

### E-commerce and Business
- **Product Link Enhancement**: Extract product information from e-commerce URLs
- **Supplier Validation**: Verify supplier website information and credibility
- **Partnership Research**: Gather information about potential business partners
- **Market Research**: Extract data from industry-related websites

### Development and Integration
- **Bookmark Management**: Enhance bookmark collections with rich metadata
- **Link Aggregation**: Build comprehensive link databases with metadata
- **Content Discovery**: Automate content discovery and categorization
- **Website Monitoring**: Track changes in website metadata over time

### Quality Assurance
- **Link Validation**: Ensure all external links in content are accessible and relevant
- **Content Verification**: Verify that linked content matches descriptions
- **Broken Link Detection**: Identify and flag inaccessible or problematic URLs
- **Metadata Consistency**: Ensure consistent metadata across linked content

## Integration Examples

- **WordPress/CMS**: Auto-populate link previews when editors add external links
- **Social Media Management**: Generate rich previews for scheduled social media posts
- **Email Marketing**: Enhance newsletter links with preview information
- **Content Aggregation**: Build curated content feeds with rich link previews
- **SEO Tools**: Analyze competitor content and metadata strategies
- **Bookmark Services**: Create enhanced bookmark collections with metadata
- **News Aggregation**: Extract article information for news curation platforms
- **Research Tools**: Gather comprehensive information about web resources

## Best Practices

- **Rate Limiting**: Respect API rate limits to avoid service interruptions
- **Caching**: Cache preview data to reduce API calls for frequently accessed URLs
- **Error Handling**: Implement proper fallbacks for failed preview extractions
- **URL Validation**: Validate URLs before sending to the API
- **Batch Processing**: Process multiple URLs efficiently in workflows
- **Data Storage**: Store extracted metadata for future reference and analysis

