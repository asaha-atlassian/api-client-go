/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 4.0.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type UserSegmentBody struct {
	// A human-friendly name for the user segment.
	Name string `json:"name"`
	// A unique key that will be used to reference the user segment in feature flags.
	Key string `json:"key"`
	// A description for the user segment.
	Description string `json:"description,omitempty"`
	// Controls whether this segment can support unlimited numbers of users. Requires the beta API and additional setup. Include/exclude lists in this payload are not used in unbounded segments.
	Unbounded bool `json:"unbounded,omitempty"`
	// Tags for the user segment.
	Tags []string `json:"tags,omitempty"`
}
