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
// PlanData struct for PlanData
type PlanData struct {
	NumberOfInstallments int32 `json:"NumberOfInstallments,omitempty"`
	Amount *MoneyWithCurrencyCode `json:"Amount,omitempty"`
	FirstInstallmentAmount *MoneyWithCurrencyCode `json:"FirstInstallmentAmount,omitempty"`
	RefOrderNumber string `json:"RefOrderNumber,omitempty"`
	TestMode *TestModes `json:"TestMode,omitempty"`
	PurchaseMethod *PurchaseMethod `json:"PurchaseMethod,omitempty"`
	Strategy *PlanStrategy `json:"Strategy,omitempty"`
	ExtendedParams map[string]string `json:"ExtendedParams,omitempty"`
	FirstChargeDate *SplititTime `json:"FirstChargeDate,omitempty"`
	AutoCapture bool `json:"AutoCapture,omitempty"`
	IsFunded bool `json:"IsFunded,omitempty"`
	Attempt3DSecure bool `json:"Attempt3DSecure,omitempty"`
}