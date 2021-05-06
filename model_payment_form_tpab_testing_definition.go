/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PaymentFormTpabTestingDefinition struct for PaymentFormTpabTestingDefinition
type PaymentFormTpabTestingDefinition struct {
	Id int64 `json:"Id"` 
	IsActive bool `json:"IsActive"` 
	AbTestName string `json:"AbTestName,omitempty"` 
	PFVersionAId int64 `json:"PFVersionAId,omitempty"` 
	PFVersionBId int64 `json:"PFVersionBId,omitempty"` 
	AbTestDescription string `json:"AbTestDescription,omitempty"` 
	PFVersionAPercentage int32 `json:"PFVersionAPercentage"` 
	PFVersionBPercentage int32 `json:"PFVersionBPercentage"` 
	PFVersionA *VersionedTouchPoints `json:"PFVersionA,omitempty"` 
	PFVersionB *VersionedTouchPoints `json:"PFVersionB,omitempty"` 
}
