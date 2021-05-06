/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PayDocuments struct for PayDocuments
type PayDocuments struct {
	Id int64 `json:"Id"` 
	TotalAmount float64 `json:"TotalAmount"` 
	ScpProviderId int64 `json:"ScpProviderId,omitempty"` 
	BusinessUnitId int64 `json:"BusinessUnitId,omitempty"` 
	Discriminator string `json:"Discriminator,omitempty"` 
	CurrencyId int64 `json:"CurrencyId,omitempty"` 
	BusinessUnit *BusinessUnits `json:"BusinessUnit,omitempty"` 
	Currency *Currencies `json:"Currency,omitempty"` 
	ScpProvider *ScpProviders `json:"ScpProvider,omitempty"` 
	FundingPayDocumentDetails []FundingPayDocumentDetails `json:"FundingPayDocumentDetails,omitempty"` 
}
