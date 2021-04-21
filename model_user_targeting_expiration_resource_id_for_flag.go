/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.1.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type UserTargetingExpirationResourceIdForFlag struct {
	Kind string `json:"kind,omitempty"`
	ProjectKey string `json:"projectKey,omitempty"`
	EnvironmentKey string `json:"environmentKey,omitempty"`
	FlagKey string `json:"flagKey,omitempty"`
	Key string `json:"key,omitempty"`
}
