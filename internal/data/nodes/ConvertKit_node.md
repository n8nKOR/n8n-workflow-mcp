# ConvertKit Node

## Overview

Consume ConvertKit API

## Credentials

- Name: convertKitApi, Required: Yes

## Inputs

- Main

## Outputs

- Main

## Properties

### Resource: customField

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: form

#### Operation: addSubscriber

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | The subscriber | Yes |  |
| Additional Fields | collection | Object of key/value pairs for custom fields (the custom field must exist before you can use it here) | No |  |

#### Operation: getSubscriptions

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Receive only active subscribers or cancelled subscribers | No |  |

## UseCases

- **Lead Generation Automation**: Automatically add new subscribers to ConvertKit forms from website sign-ups, webinar registrations, or other lead magnets
- **Email Course Enrollment**: Enroll subscribers in automated email sequences based on their interests, behaviors, or purchase history
- **Audience Segmentation**: Create and manage tags to segment subscribers by demographics, interests, or engagement levels for targeted campaigns
- **Subscriber Lifecycle Management**: Automate the process of adding, updating, and managing subscriber data across different marketing funnels
- **Content Personalization**: Use custom fields to collect and store subscriber preferences for personalized email content delivery
- **E-commerce Integration**: Sync customer data from e-commerce platforms to create targeted email campaigns based on purchase behavior
- **Event-Driven Marketing**: Trigger specific email sequences based on user actions like course completions, product purchases, or website interactions
- **Newsletter Management**: Automate the process of adding subscribers to newsletters and managing their subscription preferences
- **Abandoned Cart Recovery**: Create automated sequences to re-engage subscribers who abandoned their shopping carts
- **Customer Onboarding**: Set up automated welcome sequences and onboarding emails for new subscribers or customers

### Resource: sequence

#### Operation: addSubscriber

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Email | string | The subscriber | Yes |  |
| Additional Fields | collection | Object of key/value pairs for custom fields (the custom field must exist before you can use it here) | No |  |

#### Operation: getSubscriptions

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Receive only active subscribers or cancelled subscribers | No |  |

## UseCases

- **Lead Generation Automation**: Automatically add new subscribers to ConvertKit forms from website sign-ups, webinar registrations, or other lead magnets
- **Email Course Enrollment**: Enroll subscribers in automated email sequences based on their interests, behaviors, or purchase history
- **Audience Segmentation**: Create and manage tags to segment subscribers by demographics, interests, or engagement levels for targeted campaigns
- **Subscriber Lifecycle Management**: Automate the process of adding, updating, and managing subscriber data across different marketing funnels
- **Content Personalization**: Use custom fields to collect and store subscriber preferences for personalized email content delivery
- **E-commerce Integration**: Sync customer data from e-commerce platforms to create targeted email campaigns based on purchase behavior
- **Event-Driven Marketing**: Trigger specific email sequences based on user actions like course completions, product purchases, or website interactions
- **Newsletter Management**: Automate the process of adding subscribers to newsletters and managing their subscription preferences
- **Abandoned Cart Recovery**: Create automated sequences to re-engage subscribers who abandoned their shopping carts
- **Customer Onboarding**: Set up automated welcome sequences and onboarding emails for new subscribers or customers

### Resource: tag

#### Operation: create

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Name | string | Tag name, multiple can be added separated by comma | Yes |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |

### Resource: tagSubscriber

#### Operation: add

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Additional Fields | collection | Object of key/value pairs for custom fields (the custom field must exist before you can use it here) | No |  |

#### Operation: getAll

| Field Name | Type | Description | Required | Default |
|---|---|---|---|---|
| Return All | boolean | Whether to return all results or only up to a given limit | No |  |
| Limit | number | Max number of results to return | No |  |
| Additional Fields | collection | Receive only active subscribers or cancelled subscribers | No |  |

## UseCases

- **Lead Generation Automation**: Automatically add new subscribers to ConvertKit forms from website sign-ups, webinar registrations, or other lead magnets
- **Email Course Enrollment**: Enroll subscribers in automated email sequences based on their interests, behaviors, or purchase history
- **Audience Segmentation**: Create and manage tags to segment subscribers by demographics, interests, or engagement levels for targeted campaigns
- **Subscriber Lifecycle Management**: Automate the process of adding, updating, and managing subscriber data across different marketing funnels
- **Content Personalization**: Use custom fields to collect and store subscriber preferences for personalized email content delivery
- **E-commerce Integration**: Sync customer data from e-commerce platforms to create targeted email campaigns based on purchase behavior
- **Event-Driven Marketing**: Trigger specific email sequences based on user actions like course completions, product purchases, or website interactions
- **Newsletter Management**: Automate the process of adding subscribers to newsletters and managing their subscription preferences
- **Abandoned Cart Recovery**: Create automated sequences to re-engage subscribers who abandoned their shopping carts
- **Customer Onboarding**: Set up automated welcome sequences and onboarding emails for new subscribers or customers

