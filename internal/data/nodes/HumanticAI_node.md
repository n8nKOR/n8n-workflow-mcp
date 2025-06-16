# Humantic AI Node

## Overview

Consume Humantic AI API

## Credentials

- Name: humanticAiApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: profile

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | The LinkedIn profile URL or email ID for creating a Humantic profile. If you are sending the resume, this should be a unique string. | Yes |  |
| Send Resume | boolean | Whether to send a resume for a resume based analysis | No |  |
| Input Binary Field | string |  | No |  |

#### Operation: get

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | This value is the same as the User ID that was provided when the analysis was created. This could be a LinkedIn URL, email ID, or a unique string in case of resume based analysis. | Yes |  |
| Options | collection | Fetch the Humantic profile of the user for a particular persona type. Multiple persona values can be supported using comma as a delimiter. | No |  |

#### Operation: update

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string | This value is the same as the User ID that was provided when the analysis was created. Currently only supported for profiles created using LinkedIn URL. | Yes |  |
| Send Resume | boolean | Whether to send a resume for a resume of the user | No |  |
| Text | string | Additional text written by the user | No |  |
| Input Binary Field | string |  | No |  |

## UseCases

- **Sales Personality Insights**: Analyze prospect personalities to optimize sales approaches, communication styles, and relationship building
- **Recruitment and Hiring**: Evaluate candidate personalities for role fit, team dynamics, and cultural alignment in hiring processes
- **Customer Relationship Management**: Understand customer personalities to personalize marketing messages and improve customer engagement
- **Team Building and Management**: Analyze team member personalities for better collaboration, role assignments, and conflict resolution
- **Marketing Personalization**: Create targeted marketing campaigns based on personality insights and behavioral predictions
- **Leadership Development**: Assess leadership styles and provide personality-based coaching and development recommendations
- **Negotiation Strategy**: Understand counterpart personalities to optimize negotiation tactics and communication approaches
- **Customer Success**: Tailor customer success strategies based on personality traits and communication preferences
- **Training and Development**: Design personalized training programs based on individual learning styles and personality characteristics
- **Partnership and Collaboration**: Evaluate potential business partners and collaborators for personality compatibility and working styles

