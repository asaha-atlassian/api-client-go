/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.5.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type StreamSdkVersion struct {
	Links *StreamBySdkLinks `json:"_links,omitempty"`
	SdkVersions []StreamSdkVersionData `json:"sdkVersions,omitempty"`
}