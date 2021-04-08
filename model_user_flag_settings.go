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

type UserFlagSettings struct {
	Links *Links `json:"_links,omitempty"`
	Items map[string]UserFlagSetting `json:"items,omitempty"`
}
