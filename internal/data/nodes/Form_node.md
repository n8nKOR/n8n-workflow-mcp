# n8n Form Trigger Node

## Overview

Generate webforms in n8n and pass their responses to the workflow

## Credentials

- Name: httpBasicAuth, Required: Optional (when authentication is enabled)

## Inputs

None

## Outputs

- Main

## Properties

### Resource: Web Form

#### Operation: Generate Form

| Field Name | Type | Description | Required |
|---|---|---|---|
| Authentication | options | Authentication method for form access (none, basic) | No |
| Form Title | string | Title displayed at the top of the form | No |
| Form Description | string | Description text shown below the title | No |
| Form Fields | fixedCollection | Collection of form fields to include | Yes |
| Response Mode | options | How to respond after form submission (onReceived, lastNode, responseNode) | No |
| Button Label | string | The label of the submit button in the form | No |
| Webhook Path | string | Custom path for the form URL | No |
| Ignore Bots | boolean | Whether to ignore requests from bots and crawlers | No |
| Use Workflow Timezone | boolean | Whether to use workflow timezone in 'submittedAt' field | No |
| Custom Form Styling | string | Override default styling with custom CSS | No |

#### Form Field Configuration

| Field Name | Type | Description | Required |
|---|---|---|---|
| Field Label | string | Label displayed for the field | Yes |
| Field Type | options | Type of input field (text, textarea, number, email, date, time, datetime, select, multi-select, radio, checkbox, file) | Yes |
| Field Name | string | Internal name for the field (used in data) | Yes |
| Required | boolean | Whether the field is mandatory | No |
| Placeholder | string | Placeholder text for the field | No |

### Custom CSS Styling (v2.2+)

Override the default form appearance using CSS:

```css
/* Example custom styling */
.form-container {
    background-color: #f5f5f5;
    border-radius: 8px;
    padding: 20px;
}

.form-title {
    color: #333;
    font-size: 24px;
}

.form-field {
    margin-bottom: 15px;
}
```

### Output Data Structure

When the form is submitted, the following data is provided:

| Field | Type | Description |
|-------|------|-------------|
| submittedAt | string | Timestamp of form submission |
| formData | object | Object containing all form field values |
| [fieldName] | various | Individual field values based on field configuration |

### Webhook URLs

The node generates two webhook URLs:

1. **Setup URL (GET)**: Displays the form interface
2. **Submit URL (POST)**: Receives form submissions

### Usage Examples

#### Contact Form
```
Fields:
- Name (text, required)
- Email (email, required) 
- Message (textarea, required)
- Company (text, optional)
```

#### Survey Form
```
Fields:
- Rating (select: 1-5)
- Feedback (textarea)
- Recommend (radio: Yes/No)
- Categories (multi-select)
```

## UseCases

- [AI-Powered Blog Content Creation] : Collect keywords, chapter count, and word limits from users to automatically generate WordPress blog posts with AI-generated content and images
- [HR Job Application Processing] : Create job application forms that automatically collect candidate information, CVs, and experience details for AI-powered resume analysis and evaluation
- [Contact Form Automation] : Build contact forms that automatically route inquiries to appropriate departments and trigger follow-up workflows
- [Survey and Feedback Collection] : Create dynamic surveys with various field types to collect user feedback and automatically process responses

### Usage Notes
- Form URLs are automatically generated and can be customized
- Supports multi-step forms when combined with Form nodes in workflow
- Bot detection helps prevent spam submissions
- Custom CSS allows complete visual customization
- Form data validation is handled automatically based on field types
- Supports file uploads with proper handling of binary data
- Can be integrated with authentication systems
- Response customization allows branded success pages

