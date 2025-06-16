# PostHog Node

## Overview

Comprehensive product analytics integration with PostHog for tracking user behavior, events, and product usage. PostHog is an open-source product analytics platform that provides event tracking, user journey analysis, feature flags, and session recording capabilities for data-driven product decisions.

## Credentials

- **Name**: postHogApi
- **Required**: Yes

### Credential Configuration
- **API Key**: Your PostHog project API key
- **Host**: Your PostHog instance URL (e.g., https://app.posthog.com)

## Inputs

- **Main**: Event data, user identities, and tracking information

## Outputs

- **Main**: Confirmation of successful tracking operations and response metadata

## Properties

### Resource: alias

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Alias | string | The alias name to assign to the user | Yes | - |
| Distinct ID | string | The user's unique identifier to create an alias for | Yes | - |
| Additional Fields | collection | Additional properties for the alias operation | No | - |

**Additional Fields:**
- **Timestamp**: Custom timestamp for when the alias was created
- **Context**: Additional context information about the aliasing event

### Resource: event

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Event | string | The name of the event to track | Yes | - |
| Distinct ID | string | The user's unique identifier who performed the event | Yes | - |
| Properties | collection | Event properties and metadata | No | - |
| Additional Fields | collection | Additional event configuration | No | - |

**Additional Fields:**
- **Timestamp**: Custom timestamp for when the event occurred (if not set, current time is used)
- **Send Feature Flags**: Whether to include feature flag information
- **UUID**: Custom UUID for the event (auto-generated if not provided)

**Properties:**
- Custom properties can include any key-value pairs relevant to the event
- Common properties: `$current_url`, `$screen_name`, `$os`, `$browser`, etc.
- Revenue tracking: `revenue`, `currency`
- User properties: `email`, `name`, `plan_type`

### Resource: identity

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Distinct ID | string | The unique identifier for the user identity | Yes | - |
| Properties | collection | User properties and attributes | No | - |
| Additional Fields | collection | Additional identity configuration | No | - |

**Additional Fields:**
- **Timestamp**: Custom timestamp for the identity update (if not set, current time is used)
- **UUID**: Custom UUID for the identify call

**Properties:**
- **Email**: User's email address
- **Name**: User's full name
- **First Name**: User's first name
- **Last Name**: User's last name
- **Phone**: User's phone number
- **Avatar**: URL to user's avatar image
- **Created At**: When the user account was created
- **Plan**: User's subscription plan
- **Company**: User's company name
- **Title**: User's job title
- Custom properties: Any additional user attributes

### Resource: track

#### Operation: page

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Distinct ID | string | The user's unique identifier | Yes | - |
| Name | string | Name of the page being tracked | No | - |
| URL | string | URL of the page | No | - |
| Properties | collection | Page properties and metadata | No | - |
| Additional Fields | collection | Additional page tracking configuration | No | - |

**Properties:**
- **Title**: Page title
- **Path**: URL path
- **Referrer**: Referring page URL
- **Search**: Query parameters
- **User Agent**: Browser user agent
- **Screen Width**: Screen width in pixels
- **Screen Height**: Screen height in pixels

#### Operation: screen

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Distinct ID | string | The user's unique identifier | Yes | - |
| Name | string | Name of the screen being tracked | Yes | - |
| Properties | collection | Screen properties and metadata | No | - |
| Additional Fields | collection | Additional screen tracking configuration | No | - |

**Properties:**
- **Category**: Screen category
- **App Version**: Application version
- **OS Version**: Operating system version
- **Device Model**: Device model name
- **Screen Width**: Screen width in pixels
- **Screen Height**: Screen height in pixels
- **Orientation**: Screen orientation (portrait/landscape)

## UseCases

- **Product Analytics**: Track user behavior and product usage patterns for data-driven decisions
- **User Journey Analysis**: Analyze user flows and identify optimization opportunities in product experience
- **A/B Testing**: Run experiments and feature flags to test product changes and improvements
- **Conversion Optimization**: Track and optimize conversion funnels throughout the user journey
- **Event Tracking**: Monitor custom events and user interactions within applications
- **Cohort Analysis**: Analyze user retention and behavior patterns across different user segments
- **Feature Adoption**: Track adoption rates of new features and product capabilities
- **Performance Monitoring**: Monitor application performance and user experience metrics
- **Customer Insights**: Generate insights about customer behavior and preferences
- **Revenue Analytics**: Track revenue-related events and analyze monetization patterns
- **Onboarding Analysis**: Optimize user onboarding flows and reduce time-to-value
- **Engagement Metrics**: Measure and improve user engagement across different product areas
- **Churn Prevention**: Identify at-risk users and implement retention strategies
- **Personalization**: Use behavioral data to personalize user experiences and recommendations
- **Growth Analytics**: Measure and optimize key growth metrics and user acquisition funnels
- **Mobile App Tracking**: Track mobile app usage, screen views, and user interactions
- **Web Analytics**: Monitor website traffic, page views, and user behavior patterns
- **Cross-Platform Analytics**: Unify analytics across web, mobile, and desktop applications
- **Session Recording**: Understand user behavior through session replay and heatmaps
- **Feature Flag Management**: Control feature rollouts and measure their impact on user behavior

## Integration Tips

- **Event Naming**: Use consistent event naming conventions (e.g., verb_noun format)
- **User Identification**: Implement proper user identification across anonymous and authenticated states
- **Property Standards**: Establish standard property naming and value formats
- **Real-time Processing**: PostHog processes events in real-time for immediate insights
- **Data Retention**: Consider PostHog's data retention policies for your plan
- **Privacy Compliance**: Implement proper consent management and data privacy controls

