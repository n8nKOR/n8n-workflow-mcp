# HTML Node

## Overview

Work with HTML

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: HTML Processing

#### Operation: Generate HTML Template

| Field Name | Type | Description | Required |
|---|---|---|---|
| HTML Template | string | HTML template to render with dynamic data | Yes |

#### Operation: Extract HTML Content

| Field Name | Type | Description | Required |
|---|---|---|---|
| Source Data | options | Whether to read HTML from binary or JSON data | Yes |
| Data Property Name | string | The property containing HTML data | Yes |
| Extraction Values | fixedCollection | CSS selectors and extraction rules | No |
| Key | string | The key under which the extracted value should be saved | Yes |
| CSS Selector | string | The CSS selector to use | Yes |
| Return Value | options | What kind of data should be returned (text, html, attribute, value) | Yes |
| Attribute | string | The name of the attribute to return (when Return Value is 'attribute') | No |
| Skip Selectors | string | Comma-separated list of selectors to skip in text extraction | No |
| Return Array | boolean | Whether to return values as an array | No |

#### Operation: Convert to HTML Table

| Field Name | Type | Description | Required |
|---|---|---|---|
| Data Property Name | string | JSON property containing data to convert | Yes |
| Table Style | options | Styling options for the HTML table (default, bootstrap, custom) | No |

### Source Data Options

#### Binary Data
- **Input Binary Field**: Name of the binary property containing HTML file
- **Use Case**: Processing uploaded HTML files or web scraped content

#### JSON Data
- **JSON Property**: Name of the JSON property containing HTML string
- **Use Case**: Processing HTML content from previous nodes

### Extraction Values Configuration

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Key | string | The key under which the extracted value should be saved | Yes | - |
| CSS Selector | string | The CSS selector to use | Yes | - |
| Return Value | options | What kind of data should be returned | Yes | text |
| Attribute | string | The name of the attribute to return (when Return Value is 'attribute') | No | - |
| Skip Selectors | string | Comma-separated list of selectors to skip in text extraction | No | - |
| Return Array | boolean | Whether to return values as an array | No | false |

### Return Value Options

- **Text**: Get only the text content of the element
- **HTML**: Get the HTML the element contains
- **Attribute**: Get an attribute value like "class" from an element
- **Value**: Get value of an input, select or textarea

### HTML Template Features

#### Dynamic Content
Use `{{ expression }}` syntax for dynamic content:
```html
<h1>Hello {{ $json.name }}!</h1>
<p>Your order #{{ $json.order_id }} is ready.</p>
```

#### Conditional Content
```html
{{ $json.status === 'active' ? '<span class="active">Active</span>' : '<span class="inactive">Inactive</span>' }}
```

#### Loops and Arrays
```html
<ul>
{{ $json.items.map(item => `<li>${item.name}: $${item.price}</li>`).join('') }}
</ul>
```

#### CSS Styling
```html
<style>
  body { font-family: Arial, sans-serif; }
  .header { background-color: #007bff; color: white; padding: 20px; }
  .content { margin: 20px; }
</style>
```

### CSS Selector Examples

#### Basic Selectors
- **Element**: `div`, `p`, `h1`
- **Class**: `.className`, `.price`
- **ID**: `#elementId`, `#header`
- **Attribute**: `[data-value]`, `[href]`

#### Advanced Selectors
- **Descendant**: `.parent .child`
- **Direct Child**: `.parent > .child`
- **Adjacent Sibling**: `.element + .next`
- **Multiple**: `.class1, .class2`

#### Practical Examples
```css
/* Extract product prices */
.price

/* Get all article titles */
h2.article-title

/* Extract image URLs */
img[src]

/* Get navigation links */
nav a[href]
```

### HTML Table Conversion

#### Table Styling Options
- **Default**: Basic HTML table structure
- **Bootstrap**: Bootstrap-styled table classes
- **Custom**: User-defined CSS classes

#### Generated Table Structure
```html
<table>
  <thead>
    <tr>
      <th>Column1</th>
      <th>Column2</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Value1</td>
      <td>Value2</td>
    </tr>
  </tbody>
</table>
```

### Advanced Features

#### Skip Selectors
When extracting text, you can skip certain elements:
```
Skip Selectors: img, .advertisement, #sidebar
```

This excludes images, advertisements, and sidebar content from text extraction.

#### Array Return Mode
- **Single String**: All matching elements joined as one string
- **Array**: Each matching element as separate array item
- **Use Case**: Processing multiple similar elements

### Error Handling

- **Invalid HTML**: Attempts to parse malformed HTML
- **Missing Selectors**: Gracefully handles non-existent elements
- **Empty Results**: Returns empty values instead of errors
- **Template Errors**: Provides helpful error messages for template syntax issues

### Usage Examples

#### Generate Email Template
```html
<html>
<head>
  <style>
    .email-body { font-family: Arial; padding: 20px; }
    .header { color: #333; }
  </style>
</head>
<body>
  <div class="email-body">
    <h1 class="header">Welcome {{ $json.firstName }}!</h1>
    <p>Thank you for signing up on {{ $json.signupDate }}.</p>
  </div>
</body>
</html>
```

#### Extract Product Information
```
CSS Selector: .product-title
Return Value: text
Key: title

CSS Selector: .price
Return Value: text
Key: price

CSS Selector: .product-image
Return Value: attribute
Attribute: src
Key: imageUrl
```

#### Convert Data to Table
```
Data Property: products
Table Style: bootstrap
```

### Usage Notes
- HTML templates support full HTML5 features
- CSS styles are included in generated HTML
- JavaScript is included but not executed within n8n
- Extraction works with both well-formed and malformed HTML
- Large HTML documents may impact performance
- CSS selectors support jQuery-style syntax
- Template expressions have access to all workflow data
- Binary HTML files are automatically decoded
- Table conversion preserves data types where possible

## UseCases

- **Web Scraping Data Extraction**: Extract specific elements (links, titles, content) from web pages like Paul Graham essay lists using CSS selectors to provide input data for AI summarization systems.
- **Email Template Generation**: Generate HTML email templates with dynamic data for use in personalized email delivery systems.
- **Data Table Conversion**: Convert JSON data to HTML table format for visual representation in reports and dashboards.

