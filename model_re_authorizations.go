/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
import (
	
)
// ReAuthorizations struct for ReAuthorizations
type ReAuthorizations struct {
	Id int64 `json:"Id"` 
	InstallmentPlanId int64 `json:"InstallmentPlanId"` 
	ProcessorId int64 `json:"ProcessorId"` 
	Amount float64 `json:"Amount"` 
	ProcessDateTime *SplititTime `json:"ProcessDateTime,omitempty"` 
	IsTest bool `json:"IsTest"` 
	CardStateId int64 `json:"CardStateId,omitempty"` 
	CardState *CardStateLogs `json:"CardState,omitempty"` 
	InstallmentPlan *InstallmentPlans `json:"InstallmentPlan,omitempty"` 
	Processor *Processors `json:"Processor,omitempty"` 
	ReAuthorizationPaymentGatewayTransactionLogs []ReAuthorizationPaymentGatewayTransactionLogs `json:"ReAuthorizationPaymentGatewayTransactionLogs,omitempty"` 
	Status InstallmentStatus `json:"Status"` 
}
