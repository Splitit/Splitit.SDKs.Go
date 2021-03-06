/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// InstallmentPaymentGatewayTransactionLogs struct for InstallmentPaymentGatewayTransactionLogs
type InstallmentPaymentGatewayTransactionLogs struct {
	InstallmentId int64 `json:"InstallmentId"` 
	PaymentGatewayTransactionLogId int64 `json:"PaymentGatewayTransactionLogId"` 
	Installment *Installments `json:"Installment,omitempty"` 
	PaymentGatewayTransactionLog *PaymentGatewayTransactionLogs `json:"PaymentGatewayTransactionLog,omitempty"` 
}
