/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// DekDataAllOf struct for DekDataAllOf
type DekDataAllOf struct {
	CiphertextBlob string `json:"CiphertextBlob,omitempty"` 
	IsActive bool `json:"IsActive"` 
	DekCards []Cards `json:"DekCards,omitempty"` 
}
