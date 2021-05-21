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

type MultiEnvironmentDependentFlags struct {
	Items []MultiEnvironmentDependentFlag `json:"items,omitempty"`
	Links *DependentFlagsLinks `json:"_links,omitempty"`
	Site *Site `json:"_site,omitempty"`
}