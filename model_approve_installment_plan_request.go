/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ApproveInstallmentPlanRequest struct for ApproveInstallmentPlanRequest
type ApproveInstallmentPlanRequest struct {
	InstallmentPlanNumber string `json:"InstallmentPlanNumber,omitempty"`
	CustomerSignaturePngAsBase64 string `json:"CustomerSignaturePngAsBase64,omitempty"`
	AreTermsAndConditionsApproved bool `json:"AreTermsAndConditionsApproved"`
}