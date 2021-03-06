/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// InstallmentPlanEventSubscriberRecords struct for InstallmentPlanEventSubscriberRecords
type InstallmentPlanEventSubscriberRecords struct {
	Id int64 `json:"Id"` 
	InstallmentPlanEventId int64 `json:"InstallmentPlanEventId"` 
	SerializedInstallmentPlanEventMessage string `json:"SerializedInstallmentPlanEventMessage,omitempty"` 
	IsAcknowledged bool `json:"IsAcknowledged"` 
	SubscriberType string `json:"SubscriberType,omitempty"` 
	InstallmentPlanEvent *InstallmentPlanEvents `json:"InstallmentPlanEvent,omitempty"` 
	InstallmentPlanEventSubscriberRecordSendLogs []InstallmentPlanEventSubscriberRecordSendLogs `json:"InstallmentPlanEventSubscriberRecordSendLogs,omitempty"` 
}
