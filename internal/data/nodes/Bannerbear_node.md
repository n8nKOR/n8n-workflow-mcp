# Bannerbear Node

## Overview

The Bannerbear node allows you to consume the Bannerbear API for automated image and video generation. Bannerbear is a service that generates images and videos from templates using API calls. This is perfect for creating social media content, e-commerce banners, personalized images, and automated visual content at scale.

## Credentials

This node requires Bannerbear API credentials:
- **API Key**: Your Bannerbear project API key

To obtain API credentials:
1. Sign up for a Bannerbear account
2. Create a new project or select an existing one
3. Go to Settings → API Keys
4. Copy your project API key
5. Use the API key in n8n

## Inputs

- **Main**: JSON data for template parameters and modifications

## Outputs

- **Main**: Returns JSON data with generated image/video URLs and metadata

## Properties

### Resource: Image

#### Operation: Create
| Field Name | Type | Description | Required |
|---|---|---|---|
| Template ID | Options | Template to use for image generation | Yes |
| Modifications | Collection | Template modifications and parameters | No |
| Additional Fields | Collection | Optional settings | No |

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Image ID | String | ID of the generated image to retrieve | Yes |

### Resource: Template

#### Operation: Get
| Field Name | Type | Description | Required |
|---|---|---|---|
| Template ID | Options | ID of the template to retrieve | Yes |

#### Operation: Get Many
| Field Name | Type | Description | Required |
|---|---|---|---|
| - | - | Retrieves all available templates | No |

## Template Modifications

### Text Modifications
| Field Name | Type | Description |
|---|---|---|
| Text | String | Text content to replace in template |
| Color | String | Text color (hex code, e.g., #FF0000) |
| Background | String | Background color for text element |

### Image Modifications
| Field Name | Type | Description |
|---|---|---|
| Image URL | String | URL of image to insert |
| Fit | Options | How image should fit (crop, fit, fill) |
| Alignment | Options | Image alignment within element |

### Advanced Modifications
| Field Name | Type | Description |
|---|---|---|
| Font Family | String | Font family for text elements |
| Font Weight | String | Font weight (normal, bold, etc.) |
| Font Size | Number | Font size in pixels |
| Line Height | Number | Line height multiplier |
| Letter Spacing | Number | Letter spacing in pixels |

## Additional Options

### Image Creation Options
| Field Name | Type | Description |
|---|---|---|
| Webhook URL | String | URL to receive completion notification |
| Metadata | String | Custom metadata to attach to image |
| Wait for Image | Boolean | Wait for image processing to complete |
| Max Tries | Number | Maximum attempts when waiting for completion |

## UseCases

- **Automated Image Generation** : Generate marketing images and graphics automatically from templates
- **Social Media Content** : Create social media posts and promotional graphics at scale
- **E-commerce Product Images** : Generate product images and promotional banners for online stores
- **Marketing Campaign Assets** : Create marketing materials and campaign assets dynamically
- **Personalized Graphics** : Generate personalized images for customers and users
- **Video Thumbnail Creation** : Create custom video thumbnails and preview images
- **Certificate Generation** : Generate certificates and awards with custom data
- **Report Visualization** : Create visual reports and infographics from data
- **Brand Asset Creation** : Generate branded marketing materials and promotional content
- **Dynamic Content Creation** : Create personalized content based on user data and preferences
- **Dynamic Content Creation** : Create dynamic visual content for websites and applications

