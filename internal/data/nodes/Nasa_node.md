# NASA Node

## Overview

⚠️ **Resource Coverage**: This documentation covers the major resources available in the NASA node. The actual implementation supports 13+ resources including all DONKI space weather events, asteroid data, and Earth observation capabilities.

Retrieve comprehensive data from the NASA API including asteroid information, space weather events, Earth imagery, and astronomical pictures. This node provides access to NASA's extensive collection of scientific data and imagery for space research, education, and analysis.

## Credentials

- **Name**: nasaApi
- **Required**: Yes

### Credential Configuration
To configure NASA API credentials:
1. Visit https://api.nasa.gov/
2. Click "Get Started" to generate an API key
3. Fill out the registration form with your information
4. Copy the API key provided
5. Use the API key in your n8n credential configuration

## Inputs

- **Main**: Input data for query parameters and filters

## Outputs

- **Main**: Returns JSON data from NASA API including images, scientific data, and metadata

## Properties

### Resource: asteroidNeoLookup

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Asteroid ID | String | The ID of the asteroid to be returned | Yes | - |
| Additional Fields | Collection | Whether to include all the close approach data in the asteroid lookup | No | - |

### Resource: asteroidNeoBrowse

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | Boolean | Whether to return all results or limit | No | false |
| Limit | Number | Maximum number of asteroids to return | No | 20 |
| Additional Fields | Collection | Additional query parameters | No | - |

### Resource: asteroidNeoFeed

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for asteroid feed (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for asteroid feed (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: astronomyPictureOfTheDay

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Put Output File in Field | String | Field name to store the image data | Yes | data |
| Additional Fields | Collection | Date selection and other options | No | - |

**Additional Fields Options:**
- **Date**: Specific date for APOD (YYYY-MM-DD format)
- **Start Date**: Start date for date range
- **End Date**: End date for date range
- **Count**: Number of random images to return
- **Thumbs**: Include thumbnail URLs

### Resource: donkiCoronalMassEjection

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for CME events (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for CME events (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: donkiHighSpeedStream

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for HSS events (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for HSS events (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: donkiInterplanetaryShock

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for shock events (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for shock events (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | The location of the geomagnetic storm and other options | No | - |

### Resource: donkiMagnetopauseCrossing

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for magnetopause crossing events (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for magnetopause crossing events (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: donkiNotifications

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for notifications (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for notifications (YYYY-MM-DD) | No | - |
| Type | Options | Type of notification (all/report/forecast) | No | all |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: donkiRadiationBeltEnhancement

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for RBE events (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for RBE events (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: donkiSolarEnergeticParticle

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for SEP events (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for SEP events (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: donkiSolarFlare

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for solar flare events (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for solar flare events (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: donkiWsaEnlilSimulation

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Start Date | String | Start date for WSA-ENLIL simulations (YYYY-MM-DD) | Yes | - |
| End Date | String | End date for WSA-ENLIL simulations (YYYY-MM-DD) | No | - |
| Additional Fields | Collection | Additional filtering options | No | - |

### Resource: earthAssets

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Latitude | Number | Latitude coordinate | Yes | - |
| Longitude | Number | Longitude coordinate | Yes | - |
| Date | String | Date for Earth assets (YYYY-MM-DD) | No | - |
| Dimension | Number | Width and height of image in degrees | No | 0.15 |
| Additional Fields | Collection | Additional query parameters | No | - |

### Resource: earthImagery

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Latitude | Number | Latitude coordinate | Yes | - |
| Longitude | Number | Longitude coordinate | Yes | - |
| Date | String | Date for imagery (YYYY-MM-DD) | No | - |
| Dimension | Number | Width and height of image in degrees | No | 0.15 |
| Put Output File in Field | String | Field name to store the image data | Yes | data |
| Additional Fields | Collection | Additional imagery options | No | - |

## UseCases

- **Space Research**: Retrieve NASA data for space research and scientific analysis of asteroids, solar events, and planetary phenomena
- **Astronomy Education**: Access astronomical images and data for educational purposes and STEM learning programs
- **Asteroid Tracking**: Monitor near-Earth objects, asteroid approach data, and potential impact assessments
- **Earth Observation**: Access Earth imagery for environmental monitoring, climate studies, and geographic analysis
- **Space Weather Monitoring**: Track coronal mass ejections, solar flares, interplanetary shocks, and space weather events that affect Earth
- **Educational Content**: Create educational content using NASA's astronomy picture of the day and scientific imagery
- **Scientific Applications**: Support scientific research with NASA's comprehensive datasets including DONKI space weather data
- **Environmental Studies**: Use Earth imagery and assets for climate research, agriculture monitoring, and natural disaster assessment
- **Public Outreach**: Share NASA content for public science communication and astronomy awareness programs
- **Data Visualization**: Create visualizations and presentations using NASA imagery and scientific data
- **Research Publications**: Access NASA data for academic research, peer-reviewed publications, and scientific analysis
- **App Development**: Integrate NASA data into mobile and web applications for educational and scientific purposes
- **Space Exploration Tracking**: Monitor space missions, solar system exploration, and planetary science discoveries
- **Planetary Science**: Access planetary data for geological studies, atmospheric research, and comparative planetology
- **Satellite Data Analysis**: Analyze Earth observation satellite imagery for various research and commercial applications
- **Space Mission Planning**: Use NASA data for mission planning, orbital mechanics analysis, and space situational awareness
- **Citizen Science**: Enable citizen science projects with accessible NASA data and encourage public participation in space research
- **Media Production**: Access high-quality space imagery for documentaries, educational media, and entertainment content
- **STEM Education**: Support STEM education with real NASA data, images, and scientific information
- **Interactive Exhibits**: Create interactive displays, planetarium shows, and museum exhibits using NASA content
- **Space Weather Services**: Provide space weather information for satellite operators, airlines, and power grid operators
- **Asteroid Defense Research**: Support planetary defense research and asteroid impact mitigation studies

