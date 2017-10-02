# Go API client for swagger

Build custom integrations with the LaunchDarkly REST API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://support.launchdarkly.com](https://support.launchdarkly.com)

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditLogApi* | [**GetAuditLogEntries**](docs/AuditLogApi.md#getauditlogentries) | **Get** /auditlog | Fetch a list of all webhooks
*AuditLogApi* | [**GetAuditLogEntry**](docs/AuditLogApi.md#getauditlogentry) | **Get** /auditlog/{resourceId} | Get a webhook by ID
*EnvironmentsApi* | [**DeleteEnvironment**](docs/EnvironmentsApi.md#deleteenvironment) | **Delete** /environments/{projectKey}/{environmentKey} | Delete an environment by ID
*EnvironmentsApi* | [**GetEnvironment**](docs/EnvironmentsApi.md#getenvironment) | **Get** /environments/{projectKey}/{environmentKey} | Get an environment by key.
*EnvironmentsApi* | [**PatchEnvironment**](docs/EnvironmentsApi.md#patchenvironment) | **Patch** /environments/{projectKey}/{environmentKey} | Modify an environment by ID
*EnvironmentsApi* | [**PostEnvironment**](docs/EnvironmentsApi.md#postenvironment) | **Post** /environments/{projectKey} | Create an environment
*FlagsApi* | [**DeleteFeatureFlag**](docs/FlagsApi.md#deletefeatureflag) | **Delete** /flags/{projectKey}/{featureFlagKey} | Delete a feature flag by ID
*FlagsApi* | [**GetFeatureFlag**](docs/FlagsApi.md#getfeatureflag) | **Get** /flags/{projectKey}/{featureFlagKey} | Get a single feature flag by key.
*FlagsApi* | [**GetFeatureFlagStatus**](docs/FlagsApi.md#getfeatureflagstatus) | **Get** /flag-statuses/{projectKey}/{environmentKey} | Get a list of statuses for all feature flags
*FlagsApi* | [**GetFeatureFlagStatuses**](docs/FlagsApi.md#getfeatureflagstatuses) | **Get** /flag-statuses/{projectKey}/{environmentKey}/{featureFlagKey} | Get a list of statuses for all feature flags
*FlagsApi* | [**GetFeatureFlags**](docs/FlagsApi.md#getfeatureflags) | **Get** /flags/{projectKey} | Get a list of all features in the given project.
*FlagsApi* | [**PatchFeatureFlag**](docs/FlagsApi.md#patchfeatureflag) | **Patch** /flags/{projectKey}/{featureFlagKey} | Modify a feature flag by ID
*FlagsApi* | [**PostFeatureFlag**](docs/FlagsApi.md#postfeatureflag) | **Post** /flags/{projectKey} | Create a feature flag
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /projects/{projectKey} | Delete a project by ID
*ProjectsApi* | [**GetProject**](docs/ProjectsApi.md#getproject) | **Get** /projects/{projectKey} | Get a project by key.
*ProjectsApi* | [**GetProjects**](docs/ProjectsApi.md#getprojects) | **Get** /projects | Returns a list of all projects in the account.
*ProjectsApi* | [**PatchProject**](docs/ProjectsApi.md#patchproject) | **Patch** /projects/{projectKey} | Modify a project by ID
*ProjectsApi* | [**PostProject**](docs/ProjectsApi.md#postproject) | **Post** /projects | Create a project
*RootApi* | [**GetRoot**](docs/RootApi.md#getroot) | **Get** / | Get the root resource
*UserSettingsApi* | [**GetUserFlagSetting**](docs/UserSettingsApi.md#getuserflagsetting) | **Get** /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey} | Get a user by key.
*UserSettingsApi* | [**GetUserFlagSettings**](docs/UserSettingsApi.md#getuserflagsettings) | **Get** /users/{projectKey}/{environmentKey}/{userKey}/flags | Lists the current flag settings for a given user.
*UserSettingsApi* | [**PutFlagSetting**](docs/UserSettingsApi.md#putflagsetting) | **Put** /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey} | Specifically enable or disable a feature flag for a user based on their key.
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /users/{projectKey}/{environmentKey}/{userKey} | Delete a user by ID
*UsersApi* | [**GetSearchUsers**](docs/UsersApi.md#getsearchusers) | **Get** /user-search/{projectKey}/{environmentKey} | Search users in LaunchDarkly based on their last active date, or a search query.
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /users/{projectKey}/{environmentKey}/{userKey} | Get a user by key.
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /users/{projectKey}/{environmentKey} | List all users in the environment.
*WebhooksApi* | [**DeleteWebhook**](docs/WebhooksApi.md#deletewebhook) | **Delete** /webhooks/{resourceId} | Delete a webhook by ID
*WebhooksApi* | [**GetWebhook**](docs/WebhooksApi.md#getwebhook) | **Get** /webhooks/{resourceId} | Get a webhook by ID
*WebhooksApi* | [**GetWebhooks**](docs/WebhooksApi.md#getwebhooks) | **Get** /webhooks | Fetch a list of all webhooks
*WebhooksApi* | [**PatchWebhook**](docs/WebhooksApi.md#patchwebhook) | **Patch** /webhooks/{resourceId} | Modify a webhook by ID
*WebhooksApi* | [**PostWebhook**](docs/WebhooksApi.md#postwebhook) | **Post** /webhooks | Create a webhook


## Documentation For Models

 - [AuditLogEntries](docs/AuditLogEntries.md)
 - [AuditLogEntry](docs/AuditLogEntry.md)
 - [AuditLogEntryTarget](docs/AuditLogEntryTarget.md)
 - [Clause](docs/Clause.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentBody](docs/EnvironmentBody.md)
 - [FeatureFlag](docs/FeatureFlag.md)
 - [FeatureFlagBody](docs/FeatureFlagBody.md)
 - [FeatureFlagConfig](docs/FeatureFlagConfig.md)
 - [FeatureFlagConfigFallthrough](docs/FeatureFlagConfigFallthrough.md)
 - [FeatureFlagStatus](docs/FeatureFlagStatus.md)
 - [FeatureFlagStatuses](docs/FeatureFlagStatuses.md)
 - [FeatureFlags](docs/FeatureFlags.md)
 - [Link](docs/Link.md)
 - [Links](docs/Links.md)
 - [Member](docs/Member.md)
 - [PatchDelta](docs/PatchDelta.md)
 - [Project](docs/Project.md)
 - [ProjectBody](docs/ProjectBody.md)
 - [Projects](docs/Projects.md)
 - [Rollout](docs/Rollout.md)
 - [Rule](docs/Rule.md)
 - [Target](docs/Target.md)
 - [User](docs/User.md)
 - [UserFlagSetting](docs/UserFlagSetting.md)
 - [UserFlagSettings](docs/UserFlagSettings.md)
 - [UserSettingsBody](docs/UserSettingsBody.md)
 - [Users](docs/Users.md)
 - [Variation](docs/Variation.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookBody](docs/WebhookBody.md)
 - [Webhooks](docs/Webhooks.md)
 - [WeightedVariation](docs/WeightedVariation.md)


## Documentation For Authorization


## Token

- **Type**: API key 
- **API key parameter name**: Authorization
- **Location**: HTTP header


## Author

support@launchdarkly.com

