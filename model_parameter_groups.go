/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ParameterGroups struct for ParameterGroups
type ParameterGroups struct {
	Id int64 `json:"Id"` 
	Tokens []Tokens `json:"Tokens,omitempty"` 
	Merchants []Merchants `json:"Merchants,omitempty"` 
	Parameters []Parameters `json:"Parameters,omitempty"` 
	SplititJobDefinitions []SplititJobDefinitions `json:"SplititJobDefinitions,omitempty"` 
}
