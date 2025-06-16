# Markdown Node

## Overview

The Markdown node provides bidirectional conversion between Markdown and HTML formats, enabling seamless content transformation for documentation, web publishing, and content management workflows. The node supports both HTML-to-Markdown and Markdown-to-HTML conversion with extensive customization options and full GitHub Flavored Markdown (GFM) compatibility.

## Credentials

No credentials required - this node operates on content transformation locally.

## Inputs

- **Main**: Input data containing HTML or Markdown content to convert

## Outputs

- **Main**: Returns converted content in the specified output field

## Properties

### Resource: Content Conversion

#### Operation: HTML to Markdown (Default)

| Field Name | Type | Description | Required |
|---|---|---|---|
| HTML | String | HTML content to convert to Markdown | Yes |
| Destination Key | String | Output field name for converted content | Yes |
| Options | Collection | HTML-to-Markdown conversion options | No |

#### Operation: Markdown to HTML

| Field Name | Type | Description | Required |
|---|---|---|---|
| Markdown | String | Markdown content to convert to HTML | Yes |
| Destination Key | String | Output field name for converted content | Yes |
| Options | Collection | Markdown-to-HTML conversion options | No |

## Conversion Options

### HTML to Markdown Options

#### Text Formatting
| Option | Type | Description | Default |
|---|---|---|---|
| Bullet Marker | String | Bullet point marker for lists | `*` |
| Code Block Style | Options | Code block format (fenced/indented) | fenced |
| Emphasis Delimiter | String | Emphasis marker (`*` or `_`) | `*` |
| Strong Delimiter | String | Strong text marker (`**` or `__`) | `**` |

#### Content Processing
| Option | Type | Description | Default |
|---|---|---|---|
| Ignore Elements | Array | HTML elements to ignore during conversion | `[]` |
| Text Replace | Object | Regex patterns for text replacement | `{}` |
| Image Handling | Options | How to handle images (keep/data-uri/remove) | keep |
| Blockquote Style | Options | Blockquote formatting style | `>` |

#### Advanced Options
| Option | Type | Description | Default |
|---|---|---|---|
| Escape Characters | Array | Characters to escape in output | `[]` |
| Max Consecutive Newlines | Number | Limit consecutive blank lines | 3 |
| Use Native Parser | Boolean | Use native HTML parser | false |
| Line Separator | String | Line ending format | `\n` |

### Markdown to HTML Options

#### GitHub Flavored Markdown
| Option | Type | Description | Default |
|---|---|---|---|
| Tables | Boolean | Enable table support | true |
| Task Lists | Boolean | Enable task list checkboxes | true |
| Strikethrough | Boolean | Enable strikethrough text | true |
| Autolink | Boolean | Automatic URL linking | true |

#### Code and Syntax
| Option | Type | Description | Default |
|---|---|---|---|
| Code Blocks | Boolean | Enable fenced code blocks | true |
| GitHub Code Blocks | Boolean | GitHub-style code blocks | true |
| Syntax Highlighting | Boolean | Add syntax highlighting classes | false |

#### Document Structure
| Option | Type | Description | Default |
|---|---|---|---|
| Header ID | Boolean | Generate header IDs | false |
| Complete HTML | Boolean | Output complete HTML document | false |
| Title Tag | String | HTML document title | `""` |
| Metadata | Object | Additional HTML metadata | `{}` |

#### Text Processing
| Option | Type | Description | Default |
|---|---|---|---|
| Smart Quotes | Boolean | Convert straight quotes to curly | false |
| Emoji | Boolean | Enable emoji shortcodes | false |
| Mentions | Boolean | Enable @username mentions | false |
| Email Obfuscation | Boolean | Obfuscate email addresses | false |

## Supported Features

### HTML to Markdown Conversion

#### Supported HTML Elements
- **Text Formatting**: `<b>`, `<strong>`, `<i>`, `<em>`, `<u>`, `<strike>`
- **Headings**: `<h1>` through `<h6>`
- **Lists**: `<ul>`, `<ol>`, `<li>` with proper nesting
- **Links**: `<a>` with href attributes and titles
- **Images**: `<img>` with alt text and title support
- **Code**: `<code>`, `<pre>`, `<samp>` with language detection
- **Blockquotes**: `<blockquote>` with proper nesting
- **Tables**: `<table>`, `<tr>`, `<td>`, `<th>` (basic support)
- **Line Breaks**: `<br>`, `<hr>`

#### Advanced Features
- **Nested Lists**: Multi-level list conversion
- **Code Blocks**: Language detection and preservation
- **HTML Entities**: Proper entity decoding
- **Whitespace**: Intelligent whitespace normalization
- **Data URIs**: Base64 image handling

### Markdown to HTML Conversion

#### Standard Markdown
- **Headers**: `#` through `######` with optional IDs
- **Emphasis**: `*italic*` and `**bold**` text
- **Lists**: Numbered and bulleted lists with nesting
- **Links**: `[text](url)` and `[text][ref]` reference links
- **Images**: `![alt](src)` with title support
- **Code**: Inline `` `code` `` and fenced ```code blocks```
- **Blockquotes**: `>` with multi-line support
- **Horizontal Rules**: `---` or `***`

#### GitHub Flavored Markdown (GFM)
- **Tables**: Pipe-separated tables with alignment
- **Task Lists**: `- [x]` and `- [ ]` checkboxes
- **Strikethrough**: `~~strikethrough~~` text
- **Autolinks**: Automatic URL and email linking
- **Mention Support**: `@username` linking
- **Emoji**: `:emoji:` shortcode support

## Destination Key Configuration

### Simple Field Names
```
"converted_content"  // Creates item.converted_content
"html_output"        // Creates item.html_output
"markdown_result"    // Creates item.markdown_result
```

### Nested Object Paths
```
"content.html"       // Creates item.content.html
"data.processed.md"  // Creates item.data.processed.md
"output.formats.html" // Creates item.output.formats.html
```

## Conversion Libraries

### HTML to Markdown: `node-html-markdown`
- **Features**: Comprehensive HTML parsing and Markdown generation
- **Standards**: CommonMark compliant with extensions
- **Performance**: Optimized for large documents
- **Customization**: Extensive configuration options

### Markdown to HTML: `showdown`
- **Features**: Full Markdown and GFM support
- **Standards**: CommonMark and GitHub Flavored Markdown
- **Extensions**: Plugin architecture for custom features
- **Performance**: Fast parsing and HTML generation

## Example Configurations

### Basic HTML to Markdown
```json
{
  "mode": "htmlToMarkdown",
  "html": "<h1>Title</h1><p>Content with <strong>bold</strong> text.</p>",
  "destinationKey": "markdown_content"
}
```

### Markdown to HTML with GFM
```json
{
  "mode": "markdownToHtml", 
  "markdown": "# Title\n\n- [x] Task complete\n- [ ] Task pending",
  "destinationKey": "html_content",
  "options": {
    "tables": true,
    "taskLists": true,
    "completeHtmlDoc": true
  }
}
```

### Advanced HTML to Markdown with Custom Options
```json
{
  "mode": "htmlToMarkdown",
  "html": "<article>...</article>",
  "destinationKey": "content.markdown",
  "options": {
    "bulletMarker": "-",
    "codeBlockStyle": "fenced",
    "ignoreElements": ["script", "style"],
    "textReplace": {
      "pattern": "\\s+",
      "replacement": " "
    }
  }
}
```

## Error Handling

### Common Issues
- **Invalid HTML**: Malformed HTML may produce unexpected Markdown
- **Unsupported Elements**: Some HTML elements may not convert cleanly
- **Encoding Issues**: Character encoding problems in input content
- **Large Content**: Memory limitations with very large documents

### Best Practices
- **Validate Input**: Ensure input content is well-formed
- **Test Conversions**: Verify output format meets requirements
- **Handle Errors**: Use Continue on Fail for robust workflows
- **Monitor Performance**: Large documents may require processing limits

## UseCases

- **Documentation Publishing**: Convert documentation between formats for different platforms
- **Content Management Systems**: Transform content for web publishing and editing
- **Blog Post Processing**: Convert blog posts between editing and publishing formats
- **API Response Processing**: Transform API responses between HTML and Markdown
- **Email Template Generation**: Convert Markdown templates to HTML for email sending
- **Static Site Generation**: Process Markdown content for static site generators
- **Knowledge Base Management**: Maintain content in Markdown, publish as HTML
- **Code Documentation**: Convert README files between formats for different repositories
- **Report Generation**: Generate HTML reports from Markdown templates
- **Content Migration**: Migrate content between platforms with different format requirements
- **Social Media Publishing**: Convert rich content for different social platforms
- **E-book Publishing**: Transform content between formats for different publishing platforms
- **Wiki Management**: Convert wiki content between editing and display formats
- **Comment System Integration**: Process user comments between Markdown and HTML formats
- **Content Backup**: Create backups in multiple formats for redundancy

