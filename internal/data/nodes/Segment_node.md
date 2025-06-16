# Segment Node

## Overview

Consume Segment API

## Credentials

- Name: segmentApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: group

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | No |  |
| Group ID | string | A Group ID is the unique identifier which you recognize a group by in your own database | Yes |  |
| Traits | fixedCollection |  | No |  |
| Context | fixedCollection | Whether a user is active | No |  |
| Integration | fixedCollection |  | No |  |

### Resource: identify

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | No |  |
| Traits | fixedCollection |  | No |  |
| Context | fixedCollection | Whether a user is active | No |  |
| Integration | fixedCollection |  | No |  |

### Resource: track

#### Operation: event

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | No |  |
| Event | string | Name of the action that a user has performed | Yes |  |
| Context | fixedCollection | Whether a user is active | No |  |
| Integration | fixedCollection |  | No |  |
| Properties | fixedCollection |  | No |  |

#### Operation: page

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| User ID | string |  | No |  |
| Name | string | Name of the page For example, most sites have a "Signup" page that can be useful to tag, so you can see users as they move through your funnel | No |  |
| Context | fixedCollection | Whether a user is active | No |  |
| Integration | fixedCollection |  | No |  |
| Properties | fixedCollection |  | No |  |

## UseCases

- [Customer Data Management] : Centralize and manage customer data from multiple sources
- [Event Tracking] : Track user events and interactions across web and mobile applications
- [User Analytics] : Analyze user behavior and journey through applications
- [Personalization] : Enable personalized experiences based on customer data
- [Marketing Attribution] : Track marketing campaign effectiveness and attribution
- [A/B Testing] : Support A/B testing with event tracking and analytics
- [Customer Segmentation] : Segment customers based on behavior and characteristics
- [Real-time Analytics] : Process and analyze customer data in real-time
- [Data Integration] : Integrate customer data across multiple tools and platforms
- [Conversion Tracking] : Track conversion events and optimize conversion funnels
- [Customer Journey Mapping] : Map and analyze customer journeys across touchpoints
- [Audience Building] : Build targeted audiences for marketing campaigns
- [Data Warehousing] : Send customer data to data warehouses for analysis
- [GDPR Compliance] : Manage customer data privacy and GDPR compliance
- [Cross-Platform Tracking] : Track users across web, mobile, and other platforms

