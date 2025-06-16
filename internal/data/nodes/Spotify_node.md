# Spotify Node

## Overview

Access public song data via the Spotify API

## Credentials

- Name: spotifyOAuth2Api, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: player

#### Operation: startMusic

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Resource ID | string | Enter a playlist, artist, or album URI or ID | Yes |  |

#### Operation: addSongToQueue

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Track ID | string | Enter a track URI or ID | Yes |  |

#### Operation: volume

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Volume | number | The volume percentage to set | Yes |  |

### Resource: album

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Album ID | string | The album | Yes |  |
| Search Keyword | string | The keyword term to search for | Yes |  |

#### Operation: getNewReleases

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Filters | collection | Country to filter new releases by | No |  |

### Resource: artist

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Artist ID | string | The artist | Yes |  |
| Search Keyword | string | The keyword term to search for | Yes |  |

#### Operation: getTopTracks

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Country | string | Top tracks in which country? Enter the postal abbreviation | Yes |  |

### Resource: playlist

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Name of the playlist to create | Yes |  |
| Additional Fields | collection | Description for the playlist to create | No |  |

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | The new track | No |  |

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Search Keyword | string | The keyword term to search for | Yes |  |

### Resource: track

#### Operation: search

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Track ID | string | The track | Yes |  |
| Search Keyword | string | The keyword term to search for | Yes |  |

## UseCases

- [AI-Powered Playlist Classification] : Automatically analyze user's saved tracks using audio features (tempo, energy, danceability) and create themed playlists based on musical characteristics and listening patterns
- [Monthly Music Archiving] : Archive and organize monthly saved tracks into categorized playlists, providing automated music library management and discovery of listening trends over time
- [Intelligent Music Recommendation] : Use AI to analyze user preferences and Spotify's audio features to generate personalized playlist recommendations and discover new music based on listening history
- [Telegram Music Integration] : Enable users to search and add Spotify tracks to playlists through Telegram bot interactions, combining social messaging with music discovery and playlist management
- [Podcast Content Creation] : Convert personal notes and content into AI-generated audio podcasts, making them available through Spotify-compatible RSS feeds for automated content distribution
- [Music Mood Analysis] : Analyze playlist compositions and track features to determine mood patterns and create automated mood-based music recommendations for different activities and times of day

