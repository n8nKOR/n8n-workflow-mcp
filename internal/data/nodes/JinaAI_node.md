# Jina AI Node

## Overview

Interact with Jina AI API

## Credentials

- Name: jinaAiApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Reader

#### Operation: Read

| Field Name | Type | Description | Required |
|---|---|---|---|
| URL | string | The URL to fetch content from | Yes |
| Simplify | boolean | Return simplified response instead of raw data | No |
| Output Format | options | Desired output format (JSON, HTML, Markdown, Text, Screenshot) | No |
| Target CSS Selector | string | CSS selector to focus on specific elements | No |
| Exclude CSS Selector | string | CSS selector for elements to exclude | No |
| Enable Image Captioning | boolean | Generate captions for images within content | No |
| Wait for CSS Selector | string | Wait for element to appear before extracting | No |

#### Operation: Search

| Field Name | Type | Description | Required |
|---|---|---|---|
| Search Query | string | The search query to execute | Yes |
| Simplify | boolean | Return simplified response instead of raw data | No |
| Output Format | options | Desired output format (JSON, HTML, Markdown, Text, Screenshot) | No |
| Site Filter | string | Restrict search to specific websites | No |
| Page Number | number | Page number of search results to retrieve | No |

### Resource: Research

#### Operation: Deep Research

| Field Name | Type | Description | Required |
|---|---|---|---|
| Research Query | string | The topic or question for AI to research | Yes |
| Simplify | boolean | Return simplified response instead of raw data | No |
| Max Returned Sources | number | Maximum number of URLs in final answer | No |
| Prioritize Sources | string | Domains given higher priority for content retrieval | No |
| Exclude Sources | string | Domains to be strictly excluded from retrieval | No |
| Site Filter | string | Restrict research to specific websites | No |

### Output Formats

#### Available Formats
- **JSON**: Structured JSON data (default)
- **HTML**: Clean HTML markup
- **Markdown**: Markdown formatted text
- **Text**: Plain text content
- **Screenshot**: Visual screenshot of content

#### Format Selection
All operations support multiple output formats through the `Output Format` option:
```
X-Return-Format: html|markdown|text|screenshot|json
```

### CSS Selectors

#### Target Selector
Focus content extraction on specific page elements:
```
Target CSS Selector: #main-content .article
Target CSS Selector: .post-content, .article-body
```

#### Exclude Selector  
Remove unwanted elements from extraction:
```
Exclude CSS Selector: header, footer, .ads
Exclude CSS Selector: .sidebar, .comments, nav
```

#### Wait for Selector
Wait for dynamic content to load:
```
Wait for CSS Selector: #results-loaded
Wait for CSS Selector: .content-ready
```

### Reader URL Processing

#### Supported URLs
- **Web Pages**: Standard HTTP/HTTPS URLs
- **Dynamic Content**: JavaScript-rendered pages
- **Single Page Apps**: React, Vue, Angular applications
- **Media Files**: Images, PDFs (with appropriate handling)

#### Content Extraction
- **Clean Text**: Removes navigation, ads, and clutter
- **Structured Data**: Maintains hierarchy and formatting
- **Image Handling**: Optional image captioning for accessibility
- **Link Preservation**: Maintains important hyperlinks

### Search Functionality

#### Search Capabilities
- **Web Search**: Comprehensive web search across multiple sources
- **Site Filtering**: Restrict searches to specific domains
- **Pagination**: Access multiple pages of search results
- **Format Options**: Return results in various formats

#### Search Query Examples
```
Search Query: "machine learning trends 2024"
Search Query: site:github.com machine learning
Search Query: "artificial intelligence" -ads
```

#### Site Filtering
```
Site Filter: jina.ai, github.com
Site Filter: wikipedia.org, arxiv.org
```

### Deep Research Features

#### Research Capabilities
- **Comprehensive Analysis**: Multi-source research synthesis
- **Structured Reports**: Organized research findings
- **Source Attribution**: Proper citation of sources
- **Quality Control**: Filter high-quality sources

#### Research Query Examples
```
Research Query: "Impact of renewable energy on climate change"
Research Query: "Latest developments in quantum computing"
Research Query: "Best practices for API security in 2024"
```

#### Source Management
```
Prioritize Sources: nature.com, sciencedirect.com
Exclude Sources: socialmedia.com, unreliable.org
Site Filter: edu, gov, org
```

### Authentication

#### API Key Requirements
- **Jina AI API Key**: Required for all operations
- **Rate Limits**: Respect API rate limiting
- **Quotas**: Monitor usage quotas and limits

#### Credential Configuration
```
Credential Type: jinaAiApi
Authentication: API Key based
Header: Authorization: Bearer <api_key>
```

### Response Handling

#### Simplified Responses
When `Simplify` is enabled (default):

**Reader Operations**: Returns clean content directly
```json
{
  "content": "Clean extracted text...",
  "title": "Page title",
  "url": "source URL"
}
```

**Research Operations**: Returns structured research
```json
{
  "content": "Research findings...",
  "annotations": [...],
  "usage": {...}
}
```

#### Raw Responses
When `Simplify` is disabled: Returns full API response with metadata

### Error Handling

#### Common Errors
- **Invalid URL**: URL format or accessibility issues
- **Rate Limiting**: API quota exceeded
- **Authentication**: Invalid or missing API key
- **Content Issues**: Page load failures or selector problems

#### Error Recovery
- **Retry Logic**: Automatic retry for transient failures
- **Fallback Options**: Alternative content extraction methods
- **Validation**: URL and parameter validation

### Performance Considerations

#### Optimization Tips
- **URL Validation**: Validate URLs before processing
- **Selector Testing**: Test CSS selectors for accuracy
- **Format Selection**: Choose appropriate output format
- **Caching**: Consider caching for repeated requests

#### Rate Limiting
- **API Limits**: Respect Jina AI rate limits
- **Batch Processing**: Process multiple URLs efficiently
- **Error Handling**: Handle rate limit responses gracefully

## UseCases

#### Content Extraction
```
Operation: Read
URL: https://example.com/article
Output Format: markdown
Target Selector: article .content
Exclude Selector: .ads, .sidebar
```

#### Web Research
```
Operation: Search  
Search Query: "sustainable technology trends"
Site Filter: nature.com, science.org
Output Format: text
```

#### Academic Research
```
Operation: Deep Research
Research Query: "Climate change mitigation strategies"
Prioritize Sources: edu, gov, org
Max Returned Sources: 10
```

#### Data Mining
```
Operation: Read
URL: https://data-source.com
Wait for Selector: .data-loaded
Output Format: json
```

### Integration Patterns

#### Content Processing Pipeline
1. **URL Collection**: Gather URLs to process
2. **Content Extraction**: Use Reader to extract clean content
3. **Data Processing**: Process extracted content
4. **Storage/Analysis**: Store or analyze results

#### Research Workflow
1. **Topic Definition**: Define research questions
2. **Deep Research**: Use Research operation
3. **Result Analysis**: Process research findings
4. **Report Generation**: Generate final reports

### Best Practices

#### URL Processing
- **Validate URLs**: Check URL accessibility
- **Test Selectors**: Verify CSS selectors work correctly
- **Handle Errors**: Implement proper error handling
- **Monitor Performance**: Track processing times

#### Research Quality
- **Clear Queries**: Use specific, well-defined research questions
- **Source Quality**: Prioritize authoritative sources
- **Result Validation**: Verify research findings
- **Citation Tracking**: Maintain source attribution

### Usage Notes
- Supports multiple output formats for flexible content processing
- CSS selectors enable precise content targeting and filtering
- Image captioning provides accessibility for visual content
- Deep research combines multiple sources for comprehensive analysis
- Rate limiting applies to all API operations
- Authentication required for all operations
- Simplify option provides clean, processed responses vs raw API data
- Wait selectors handle dynamic JavaScript-rendered content
- Site filtering enables domain-specific searches and research
