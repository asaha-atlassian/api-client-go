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

type FlagConfigScheduledChangesPatchBody struct {
	// Used to describe the scheduled changes.
	Comment string `json:"comment,omitempty"`
	Instructions *SemanticPatchInstruction `json:"instructions,omitempty"`
}
