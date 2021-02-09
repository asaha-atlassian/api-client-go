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

type MaUbyCategory struct {
	Links *StreamBySdkLinks `json:"_links,omitempty"`
	Metadata []MauMetadata `json:"metadata,omitempty"`
	Series []StreamUsageSeries `json:"series,omitempty"`
}
