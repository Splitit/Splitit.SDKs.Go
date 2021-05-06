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
// AccountUpdaterPendingCards struct for AccountUpdaterPendingCards
type AccountUpdaterPendingCards struct {
	Id int64 `json:"Id"` 
	CardId int64 `json:"CardId"` 
	ExportTime *SplititTime `json:"ExportTime,omitempty"` 
	ExportBatchId string `json:"ExportBatchId,omitempty"` 
	ExportRecordId string `json:"ExportRecordId,omitempty"` 
	ImportTime *SplititTime `json:"ImportTime,omitempty"` 
	Card *Cards `json:"Card,omitempty"` 
	AccountUpdater AccountUpdaterProvider `json:"AccountUpdater"` 
}
