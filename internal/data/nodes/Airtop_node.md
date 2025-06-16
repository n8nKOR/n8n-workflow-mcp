# Airtop Node

## Overview

The Airtop node allows you to scrape and control any website using Airtop's browser automation platform. Airtop provides cloud-based browser sessions that you can control programmatically to extract data, interact with web pages, and automate web workflows.

## Credentials

This node requires Airtop API credentials:
- **API Key**: Your Airtop API key

To obtain API credentials:
1. Sign up for an Airtop account at https://airtop.ai
2. Navigate to your dashboard
3. Generate an API key from the settings or API section

## Inputs

- **Main**: JSON data for session and interaction control

## Outputs

- **Main**: Returns JSON data from Airtop API responses, including extracted data and session information

## Properties

### Resource: Session

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Configuration | Collection | Session configuration options | No |
| Profile ID | String | Browser profile ID to use | No |
| Proxy | Collection | Proxy configuration | No |
| Screen Size | Collection | Browser window dimensions | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Return All | Boolean | Whether to return all sessions | No |
| Limit | Number | Max number of sessions to return | No |

#### Operation: Delete
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session to terminate | Yes |

### Resource: Window

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session to create window in | Yes |
| URL | String | URL to navigate to | No |
| Window Options | Collection | Window configuration options | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Return All | Boolean | Whether to return all windows | No |
| Limit | Number | Max number of windows to return | No |

#### Operation: Close
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window to close | Yes |

### Resource: Interaction

#### Operation: Click
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| Selector | String | CSS selector or XPath for element to click | Yes |
| Options | Collection | Click options (wait, timeout, etc.) | No |

#### Operation: Type
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| Selector | String | CSS selector or XPath for input element | Yes |
| Text | String | Text to type | Yes |
| Options | Collection | Typing options (clear, delay, etc.) | No |

#### Operation: Navigate
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| URL | String | URL to navigate to | Yes |
| Options | Collection | Navigation options (wait, timeout, etc.) | No |

#### Operation: Wait
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| Wait Type | Options | Type of wait (element, timeout, page load) | Yes |
| Wait Options | Collection | Wait configuration options | Yes |

### Resource: Extraction

#### Operation: Get Page Content
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| Content Type | Options | Type of content to extract (HTML, text, etc.) | Yes |
| Options | Collection | Extraction options | No |

#### Operation: Get Element
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| Selector | String | CSS selector or XPath for element | Yes |
| Extract | Options | What to extract (text, attribute, HTML) | Yes |
| Options | Collection | Extraction options | No |

#### Operation: Get Elements
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| Selector | String | CSS selector or XPath for elements | Yes |
| Extract | Options | What to extract from each element | Yes |
| Options | Collection | Extraction options | No |

#### Operation: Take Screenshot
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| Options | Collection | Screenshot options (format, quality, area) | No |

### Resource: File

#### Operation: Upload
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| File Input Selector | String | CSS selector for file input element | Yes |
| File Data | String | Binary data or file path | Yes |
| Options | Collection | Upload options | No |

#### Operation: Download
| Field Name | Type | Description | Required |
|---|---|---|---|
| Session ID | String | ID of the session | Yes |
| Window ID | String | ID of the window | Yes |
| URL | String | URL of file to download | Yes |
| Options | Collection | Download options | No |

## UseCases

- **Web Scraping** : Extract data from websites that require browser interaction
- **Automated Testing** : Perform automated testing of web applications and user interfaces
- **Data Collection** : Collect data from dynamic websites with JavaScript-heavy content
- **Form Automation** : Automate form filling and submission processes
- **Social Media Automation** : Automate social media interactions and content posting
- **E-commerce Monitoring** : Monitor product prices, availability, and competitor analysis
- **Lead Generation** : Automate lead collection from various online sources
- **Content Management** : Automate content updates and management across multiple platforms
- **Quality Assurance** : Perform automated quality assurance testing for web applications
- **Market Research** : Gather market intelligence and competitive analysis data
- **Website Monitoring** : Monitor websites for changes, updates, and performance issues
- **SEO Analysis** : Perform automated SEO analysis and ranking monitoring

