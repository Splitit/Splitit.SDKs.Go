/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// TerminalsAllOf struct for TerminalsAllOf
type TerminalsAllOf struct {
	Name string `json:"Name,omitempty"` 
	Email string `json:"Email,omitempty"` 
	ApiKey string `json:"ApiKey,omitempty"` 
	WizardTimeout int32 `json:"WizardTimeout"` 
	MerchantId int64 `json:"MerchantId"` 
	UtcOffset float64 `json:"UtcOffset"` 
	IsDisabled bool `json:"IsDisabled"` 
	ChargeBeforeShipping bool `json:"ChargeBeforeShipping"` 
	TerminalId string `json:"TerminalId,omitempty"` 
	AgentId int64 `json:"AgentId"` 
	CanCancelInstallmentsPlan bool `json:"CanCancelInstallmentsPlan"` 
	CrmId string `json:"CrmId,omitempty"` 
	CurrencyId int64 `json:"CurrencyId,omitempty"` 
	PendingShipmentReminderDays int32 `json:"PendingShipmentReminderDays"` 
	BusinessUnitId int64 `json:"BusinessUnitId"` 
	KeepNonApprovedPlanLive int64 `json:"KeepNonApprovedPlanLive"` 
	UseTestGateway bool `json:"UseTestGateway"` 
	Agent *Agents `json:"Agent,omitempty"` 
	BusinessUnit *BusinessUnits `json:"BusinessUnit,omitempty"` 
	Currency *Currencies `json:"Currency,omitempty"` 
	Merchant *Merchants `json:"Merchant,omitempty"` 
	AccountUpdaterSubscribersScopes []AccountUpdaterSubscribersScopes `json:"AccountUpdaterSubscribersScopes,omitempty"` 
	InstallmentPlans []InstallmentPlans `json:"InstallmentPlans,omitempty"` 
	ReconciliationReports []ReconciliationReports `json:"ReconciliationReports,omitempty"` 
	StateLimitRuleDatas []StateLimitRuleDatas `json:"StateLimitRuleDatas,omitempty"` 
	TerminalCountries []TerminalCountries `json:"TerminalCountries,omitempty"` 
	TerminalGatewayDatas []TerminalGatewayDatas `json:"TerminalGatewayDatas,omitempty"` 
	TerminalKvps []TerminalKvps `json:"TerminalKvps,omitempty"` 
	ContinueExistingPlanWithOriginalGateway bool `json:"ContinueExistingPlanWithOriginalGateway"` 
	ActiveTerminalData *TerminalGatewayDatas `json:"ActiveTerminalData,omitempty"` 
	ChbDefaultAction ChbDefaultAction `json:"ChbDefaultAction"` 
	FailureUnderFrozenInstallmentsPlan FailureUnderFrozenInstallmentsPlan `json:"FailureUnderFrozenInstallmentsPlan"` 
	RefundUnderCancelation RefundUnderCancelation `json:"RefundUnderCancelation"` 
	IntegrationType IntegrationType `json:"IntegrationType"` 
	TestMode TestModes `json:"TestMode"` 
}
