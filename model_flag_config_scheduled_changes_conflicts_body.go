/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.0.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FlagConfigScheduledChangesConflictsBody struct {
	// A unix epoch time in milliseconds specifying the date the scheduled changes will be applied
	ExecutionDate int32 `json:"executionDate,omitempty"`
	Instructions *SemanticPatchInstruction `json:"instructions,omitempty"`
}
