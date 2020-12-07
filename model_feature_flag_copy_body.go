/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.10.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagCopyBody struct {
	Source *FeatureFlagCopyObject `json:"source,omitempty"`
	Target *FeatureFlagCopyObject `json:"target,omitempty"`
	// comment will be included in audit log item for change.
	Comment string `json:"comment,omitempty"`
	// Define the parts of the flag configuration that will be copied.
	IncludedActions []CopyActions `json:"includedActions,omitempty"`
	// Define the parts of the flag configuration that will not be copied.
	ExcludedActions []CopyActions `json:"excludedActions,omitempty"`
}
