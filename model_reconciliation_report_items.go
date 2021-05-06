/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ReconciliationReportItems struct for ReconciliationReportItems
type ReconciliationReportItems struct {
	Id int64 `json:"Id"` 
	GatewayAmount float64 `json:"GatewayAmount"` 
	GatewayCount int32 `json:"GatewayCount"` 
	Pisamount float64 `json:"Pisamount"` 
	Piscount int32 `json:"Piscount"` 
	IsIdentical bool `json:"IsIdentical"` 
	ReconciliationReportId int64 `json:"ReconciliationReportId"` 
	ReconciliationReport *ReconciliationReports `json:"ReconciliationReport,omitempty"` 
	TransactionType OperationType `json:"TransactionType"` 
	GatewayType ProcessorType `json:"GatewayType"` 
}
