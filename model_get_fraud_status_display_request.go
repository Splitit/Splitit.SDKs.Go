/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// GetFraudStatusDisplayRequest struct for GetFraudStatusDisplayRequest
type GetFraudStatusDisplayRequest struct {
	ProviderReferenceId string `json:"ProviderReferenceId,omitempty"`
	MerchantId int64 `json:"MerchantId,omitempty"`
	InstallmentPlanNumber string `json:"InstallmentPlanNumber,omitempty"`
}