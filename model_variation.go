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

type Variation struct {
	Id string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Value *interface{} `json:"value"`
}
