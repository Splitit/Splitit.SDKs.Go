/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// CancelInstallmentPlanRequest struct for CancelInstallmentPlanRequest
type CancelInstallmentPlanRequest struct {
	InstallmentPlanNumber string `json:"InstallmentPlanNumber,omitempty"` 
	RefundUnderCancelation RefundUnderCancelation `json:"RefundUnderCancelation"` 
	CancelationReason InstallmentPlanCancelationReason `json:"CancelationReason,omitempty"` 
	IsExecutedUnattended bool `json:"IsExecutedUnattended"` 
	PartialResponseMapping bool `json:"PartialResponseMapping"` 
}
