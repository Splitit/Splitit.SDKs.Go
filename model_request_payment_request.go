/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// RequestPaymentRequest struct for RequestPaymentRequest
type RequestPaymentRequest struct {
	InstallmentPlanNumber string `json:"InstallmentPlanNumber,omitempty"` 
	PaymentApprovalPhone string `json:"PaymentApprovalPhone,omitempty"` 
	PaymentApprovalEmail string `json:"PaymentApprovalEmail,omitempty"` 
}
