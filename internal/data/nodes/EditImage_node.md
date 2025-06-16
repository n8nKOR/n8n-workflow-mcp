# Edit Image Node

## Overview

Edits an image like blur, resize or adding border and text

## Credentials

None

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: Image Editor

#### Operation: Get Information

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |

#### Operation: Blur

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Blur | number | The blur value | Yes |

#### Operation: Border

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Border Width | number | The width of the border | Yes |
| Border Color | color | The color of the border | Yes |

#### Operation: Composite

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Composite Image Property | string | Name of the binary property containing the image to composite | Yes |
| Position X | number | X position of composite image | No |
| Position Y | number | Y position of composite image | No |

#### Operation: Create

| Field Name | Type | Description | Required |
|---|---|---|---|
| Background Color | color | The background color of the image to create | Yes |
| Image Width | number | The width of the image to create | Yes |
| Image Height | number | The height of the image to create | Yes |

#### Operation: Crop

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Position X | number | X coordinate of crop area top-left corner | Yes |
| Position Y | number | Y coordinate of crop area top-left corner | Yes |
| Width | number | Width of crop area | Yes |
| Height | number | Height of crop area | Yes |

#### Operation: Draw

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Primitive | options | The primitive to draw (circle, line, rectangle) | Yes |
| Color | color | The color of the primitive to draw | Yes |
| Start Position X | number | X (horizontal) start position of the primitive | Yes |
| Start Position Y | number | Y (vertical) start position of the primitive | Yes |

#### Operation: Resize

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Width | number | New width of the image | Yes |
| Height | number | New height of the image | Yes |

#### Operation: Rotate

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Rotate | number | How much the image should be rotated (degrees) | Yes |
| Background Color | color | The background color to use for the areas that are empty after rotation | No |

#### Operation: Shear

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Degrees X | number | Shear along X axis (degrees) | Yes |
| Degrees Y | number | Shear along Y axis (degrees) | Yes |

#### Operation: Text

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Text | string | The text to add to the image | Yes |
| Font Size | number | Size of the text | Yes |
| Font Color | color | Color of the text | Yes |
| Position X | number | X coordinate where to position the text | Yes |
| Position Y | number | Y coordinate where to position the text | Yes |
| Font | options | The font to use (Arial, etc.) | No |

#### Operation: Transparent

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Color | color | The color to make transparent | Yes |

#### Operation: Multi Step

| Field Name | Type | Description | Required |
|---|---|---|---|
| Property Name | string | Name of the binary property in which the image data can be found | Yes |
| Operations | fixedCollection | Collection of operations to perform in order | No |

### General Options

| Field Name | Type | Description | Required |
|---|---|---|---|
| File Name | string | File name to set in binary data | No |
| Format | options | Output image format (bmp, gif, jpeg, png, tiff, webp) | No |
| Quality | number | Compression quality for jpeg/png/tiff (0-100) | No |

### Usage Notes
- Input image should be in binary format
- Multiple operations can be chained using the Multi Step operation
- Font selection is available through system fonts
- Transparency is supported with alpha channel colors
- Various image formats are supported for output

## UseCases

- [Image Information Extraction] : Extract image resolution and size information in AI captioning systems to calculate the position and size of caption text overlays
- [AI Processing Image Resizing] : Resize large images to optimal sizes (e.g., 512x512) for efficient AI model processing, improving processing speed and optimizing memory usage
- [Image Caption Overlay] : Overlay AI-generated captions on images with semi-transparent backgrounds using multi-step operations to draw background rectangles and add text sequentially
- [PDF to Image Conversion Processing] : Resize converted PDF images to 75% of original size for faster LLM processing while maintaining readability for multimodal AI analysis
- [Resume Screening Image Preparation] : Process resume PDF images by resizing them to optimal dimensions for Vision Language Model analysis in automated candidate evaluation systems

