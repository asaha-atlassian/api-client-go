/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.0.3
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagStatusAcrossEnvironments struct {
	Links *FeatureFlagStatusLinks `json:"_links,omitempty"`
	Key string `json:"key,omitempty"`
	Environments map[string]FeatureFlagStatusForQueriedEnvironment `json:"environments,omitempty"`
}
