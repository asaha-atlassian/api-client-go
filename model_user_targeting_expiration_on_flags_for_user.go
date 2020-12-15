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

type UserTargetingExpirationOnFlagsForUser struct {
	Links *Links `json:"_links,omitempty"`
	Items []UserTargetingExpirationForFlag `json:"items,omitempty"`
}
