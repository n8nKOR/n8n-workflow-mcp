# LingvaNex Node

## Overview

Consume LingvaNex API

## Credentials

- Name: lingvaNexApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Translation

#### Operation: Translate
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Text | string | The input text to translate | Yes | - |
| Translate To | options | Target language for translation | Yes | - |
| From | options | Source language code (auto-detect if not specified) | No | - |
| Platform | string | API platform identifier | No | api |
| Translate Mode | string | Input text format ("html" for HTML structure preservation) | No | - |

**Note**: The "Translate To" and "From" fields use dynamic options loaded from the LingvaNex API, supporting all languages listed in their language catalog.

## UseCases

- **Content Localization**: Translate website content for international audiences
- **Document Translation**: Translate business documents and reports
- **Customer Support**: Translate customer messages in multiple languages
- **Email Translation**: Translate email content for global communication
- **Social Media**: Translate social media posts for different markets

