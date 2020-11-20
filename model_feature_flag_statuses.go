/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.9.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagStatuses struct {
	Links *Links `json:"_links,omitempty"`
	Items []FeatureFlagStatus `json:"items,omitempty"`
}
