/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ChargebackRequest struct for ChargebackRequest
type ChargebackRequest struct {
	TransactionIdsToMark []int64 `json:"TransactionIdsToMark,omitempty"` 
	TransactionIdsToUnmark []int64 `json:"TransactionIdsToUnmark,omitempty"` 
	InstallmentPlanNumber string `json:"InstallmentPlanNumber,omitempty"` 
	PartialResponseMapping bool `json:"PartialResponseMapping"` 
}
