/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ConfigKeysAllOf struct for ConfigKeysAllOf
type ConfigKeysAllOf struct {
	Name string `json:"Name,omitempty"` 
	Code string `json:"Code,omitempty"` 
	TouchPointId int64 `json:"TouchPointId,omitempty"` 
	TouchPoint *TouchPoints `json:"TouchPoint,omitempty"` 
	ConfigValues []ConfigValues `json:"ConfigValues,omitempty"` 
}
