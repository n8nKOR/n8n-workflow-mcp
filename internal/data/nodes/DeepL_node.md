# DeepL Node

## Overview

Translate data using DeepL

## Credentials

- Name: deepLApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: language

#### Operation: translate

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Text | string | The text to translate | Yes |  |
| Translate To | options | Target language for translation | Yes |  |
| Additional Fields | collection | Additional translation options | No |  |

##### Additional Fields Collection:
| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Source Language | options | Source language of the text (auto-detected if not specified) | No |  |
| Formality | options | Formality level for translation | No | default |
| Preserve Formatting | boolean | Whether to preserve formatting in translation | No | false |

##### Supported Languages:
- Target languages: Loaded dynamically from DeepL API
- Source languages: Auto-detected or manually specified
- Includes major world languages like English, German, French, Spanish, Italian, Japanese, Chinese, etc.

##### Translation Features:
- High-quality neural machine translation
- Automatic language detection
- Formality control for supported languages
- Formatting preservation options
- Support for multiple text formats

## UseCases

- **Content Localization** : Translate website content and marketing materials for international markets
- **Document Translation** : Translate business documents, contracts, and technical documentation
- **Customer Support** : Provide multilingual customer support by translating inquiries and responses
- **E-commerce Internationalization** : Translate product descriptions and store content for global sales
- **Social Media Management** : Translate social media content for international audience engagement
- **Email Marketing** : Translate email campaigns and newsletters for multilingual audiences
- **Educational Content** : Translate educational materials and course content for global learners
- **News and Media** : Translate news articles and media content for international distribution
- **API Response Translation** : Translate API responses and data for multilingual applications
- **Real-time Communication** : Enable real-time translation in chat applications and messaging systems

