/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// StateLimitRuleDatasAllOf struct for StateLimitRuleDatasAllOf
type StateLimitRuleDatasAllOf struct {
	TerminalId int64 `json:"TerminalId"` 
	CountrySubdivisionId int64 `json:"CountrySubdivisionId"` 
	CountrySubdivision *CountrySubdivisions `json:"CountrySubdivision,omitempty"` 
	Terminal *Terminals `json:"Terminal,omitempty"` 
}
