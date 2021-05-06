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
// FeesDocuments struct for FeesDocuments
type FeesDocuments struct {
	Id int64 `json:"Id"` 
	FeeAmountInEffectiveCurrencyAmount float64 `json:"FeeAmountInEffectiveCurrencyAmount"` 
	FeeAmountInTransactionCurrencyAmount float64 `json:"FeeAmountInTransactionCurrencyAmount"` 
	FeeAmountInUniformCurrencyAmount float64 `json:"FeeAmountInUniformCurrencyAmount"` 
	EffectiveCurrencyId int64 `json:"EffectiveCurrencyId"` 
	TransactionCurrencyId int64 `json:"TransactionCurrencyId,omitempty"` 
	UniformCurrencyId int64 `json:"UniformCurrencyId"` 
	ReferenceFeesDocumentId int64 `json:"ReferenceFeesDocumentId,omitempty"` 
	InstallmentId int64 `json:"InstallmentId,omitempty"` 
	PlanId int64 `json:"PlanId,omitempty"` 
	BusinessUnitId int64 `json:"BusinessUnitId,omitempty"` 
	EffectiveDateUtc *SplititTime `json:"EffectiveDateUtc"` 
	FeeRuleDataId int64 `json:"FeeRuleDataId,omitempty"` 
	FeeRuleData *FeesRuleDatas `json:"FeeRuleData,omitempty"` 
	BusinessUnit *BusinessUnits `json:"BusinessUnit,omitempty"` 
	EffectiveCurrency *Currencies `json:"EffectiveCurrency,omitempty"` 
	Installment *Installments `json:"Installment,omitempty"` 
	Plan *InstallmentPlans `json:"Plan,omitempty"` 
	ReferenceFeesDocument *FeesDocuments `json:"ReferenceFeesDocument,omitempty"` 
	TransactionCurrency *Currencies `json:"TransactionCurrency,omitempty"` 
	UniformCurrency *Currencies `json:"UniformCurrency,omitempty"` 
	InverseReferenceFeesDocument []FeesDocuments `json:"InverseReferenceFeesDocument,omitempty"` 
	FeeEntity AccountingParty `json:"FeeEntity"` 
	FeeType FeesTypes `json:"FeeType"` 
	RangeType RangeType `json:"RangeType,omitempty"` 
	FeeAmountInEffectiveCurrency *Money2 `json:"FeeAmountInEffectiveCurrency,omitempty"` 
	FeeAmountInTransactionCurrency *Money2 `json:"FeeAmountInTransactionCurrency,omitempty"` 
	FeeAmountInUniformCurrency *Money2 `json:"FeeAmountInUniformCurrency,omitempty"` 
}
