/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// RolesAllOf struct for RolesAllOf
type RolesAllOf struct {
	Name string `json:"Name,omitempty"` 
	RoleStrength int32 `json:"RoleStrength"` 
	BusinessParty BusinessParty `json:"BusinessParty"` 
	BusinessActivities []BusinessActivity `json:"BusinessActivities,omitempty"` 
	BusinessActivitiesSerialized string `json:"BusinessActivitiesSerialized,omitempty"` 
}
