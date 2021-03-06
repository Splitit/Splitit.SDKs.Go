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
// InstallmentPlanEvents struct for InstallmentPlanEvents
type InstallmentPlanEvents struct {
	Id int64 `json:"Id"` 
	InstallmentPlanId int64 `json:"InstallmentPlanId"` 
	ActivityTime *SplititTime `json:"ActivityTime"` 
	InstallmentPlan *InstallmentPlans `json:"InstallmentPlan,omitempty"` 
	InstallmentPlanEventSubscriberRecords []InstallmentPlanEventSubscriberRecords `json:"InstallmentPlanEventSubscriberRecords,omitempty"` 
	Event InstallmentPlanEventType `json:"Event"` 
}
