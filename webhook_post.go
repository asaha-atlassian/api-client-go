/* 
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@launchdarkly.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type WebhookPost struct {

	Url string `json:"url"`

	Secret string `json:"secret,omitempty"`

	Sign bool `json:"sign"`

	On bool `json:"on"`
}