/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// UserWithActionItemAllOf struct for UserWithActionItemAllOf
type UserWithActionItemAllOf struct {
	FirstName string `json:"FirstName,omitempty"` 
	LastName string `json:"LastName,omitempty"` 
	JobTitle string `json:"JobTitle,omitempty"` 
	FaxNumber string `json:"FaxNumber,omitempty"` 
	MobileNumber string `json:"MobileNumber,omitempty"` 
	WorkPhoneNumber string `json:"WorkPhoneNumber,omitempty"` 
	Notes string `json:"Notes,omitempty"` 
	Type UserPermissionLevel `json:"Type"` 
}
