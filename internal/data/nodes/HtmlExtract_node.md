# HTML Extract Node

## Overview

Extracts data from HTML

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: HTML Extract

#### Operation: Extract Data

| Field Name | Type | Description | Required |
|---|---|---|---|
| Source Data | Options | If HTML should be read from binary or JSON data | Yes |
| Data Property Name | String | The property containing HTML data | Yes |
| Extraction Values | FixedCollection | CSS selectors and extraction rules | No |
| Key | String | The key under which the extracted value should be saved | Yes |
| CSS Selector | String | The CSS selector to use | Yes |
| Return Value | Options | What kind of data should be returned (text, html, attribute, value) | Yes |
| Attribute | String | The name of the attribute to return (when Return Value is 'attribute') | No |
| Return Array | Boolean | Whether to return values as an array | No |
| Trim Values | Boolean | Whether to trim whitespace from extracted values | No |

## UseCases

- **Web Scraping** : Extract structured data from HTML pages and web content
- **Data Mining** : Parse HTML documents to extract specific information and datasets
- **Content Processing** : Extract text, links, and media from HTML content for further processing
- **Product Information Extraction** : Extract product details, prices, and descriptions from e-commerce sites
- **News and Article Processing** : Extract headlines, authors, and content from news websites
- **Form Data Extraction** : Extract form fields and values from HTML forms

## Source Data Options

### Binary
- **Input Binary Field**: Name of the binary property containing HTML file
- **Use Case**: Processing uploaded HTML files or web scraped content

### JSON
- **JSON Property**: Name of the JSON property containing HTML string or array of strings
- **Use Case**: Processing HTML content from previous nodes

## Return Value Options

- **Text**: Get only the text content of the element
- **HTML**: Get the HTML the element contains  
- **Attribute**: Get an attribute value like "class" from an element
- **Value**: Get value of an input, select or textarea

## CSS Selector Examples

### Basic Selectors
- **Element**: `div`, `p`, `h1`
- **Class**: `.className`, `.price`
- **ID**: `#elementId`, `#header`
- **Attribute**: `[data-value]`, `[href]`

### Advanced Selectors
- **Descendant**: `.parent .child`
- **Direct Child**: `.parent > .child`
- **Adjacent Sibling**: `.element + .next`
- **Multiple**: `.class1, .class2`

### Practical Examples
```css
/* Extract product prices */
.price

/* Get all article titles */
h2.article-title

/* Extract image URLs */
img[src]

/* Get navigation links */
nav a[href]

/* Extract data attributes */
[data-product-id]
```

## Data Processing

### Single vs Array Return
- **Single String**: All matching elements joined as one string
- **Array Mode**: Each matching element as separate array item
- **Use Case**: Choose based on whether you need individual elements or combined text

### Value Trimming
- **Default**: Whitespace is automatically trimmed from extracted values
- **Disable**: Set "Trim Values" to false to preserve original spacing
- **Use Case**: Preserving formatted text vs cleaning up extracted data

## HTML Parsing

### Parser Capabilities
- **Cheerio**: Uses Cheerio library for jQuery-style HTML parsing
- **Malformed HTML**: Handles poorly formatted HTML gracefully
- **Performance**: Efficient parsing for large HTML documents

### Supported HTML Features
- **HTML5**: Full HTML5 element support
- **Nested Elements**: Complex nested structures
- **Attributes**: All standard and custom attributes
- **Text Content**: Plain text extraction from elements

## Error Handling

- **Missing Elements**: Returns undefined for non-existent selectors
- **Invalid HTML**: Attempts to parse malformed HTML
- **Empty Results**: Gracefully handles empty selections
- **Attribute Errors**: Returns undefined for missing attributes

## Usage Examples

### Extract Product Information
```
Source Data: json
Data Property Name: htmlContent

Extraction Values:
1. Key: title
   CSS Selector: .product-title
   Return Value: text

2. Key: price  
   CSS Selector: .price
   Return Value: text

3. Key: imageUrl
   CSS Selector: .product-image
   Return Value: attribute
   Attribute: src

4. Key: description
   CSS Selector: .description
   Return Value: html
```

### Extract Multiple Links
```
Source Data: binary
Data Property Name: webpage

Extraction Values:
1. Key: navigationLinks
   CSS Selector: nav a
   Return Value: attribute
   Attribute: href
   Return Array: true

2. Key: linkTexts
   CSS Selector: nav a
   Return Value: text
   Return Array: true
```

### Extract Article Data
```
Source Data: json
Data Property Name: articleHtml

Extraction Values:
1. Key: headline
   CSS Selector: h1.headline
   Return Value: text

2. Key: author
   CSS Selector: .author-name
   Return Value: text

3. Key: publishDate
   CSS Selector: .publish-date
   Return Value: attribute
   Attribute: datetime

4. Key: content
   CSS Selector: .article-body
   Return Value: html
```

## Migration Notice

**⚠️ Deprecated**: This node is hidden and deprecated. Use the newer **HTML** node instead, which provides:
- Better performance
- Enhanced CSS selector support
- More extraction options

## Related Nodes

- **HTML**: The newer replacement for HTML extraction
- **HTTP Request**: For fetching HTML content from web pages
- **JSON**: For processing extracted data
- **Code**: For custom HTML processing logic

