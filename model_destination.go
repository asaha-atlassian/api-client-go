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

type Destination struct {
	Links *Links `json:"_links,omitempty"`
	// Unique destination ID.
	Id string `json:"_id,omitempty"`
	// The destination name
	Name string `json:"name,omitempty"`
	// Destination type (\"google-pubsub\", \"kinesis\", \"mparticle\", or \"segment\")
	Kind string `json:"kind,omitempty"`
	// destination-specific configuration.
	Config *interface{} `json:"config,omitempty"`
	// Whether the data export destination is on or not.
	On bool `json:"on,omitempty"`
	Version int32 `json:"version,omitempty"`
}
