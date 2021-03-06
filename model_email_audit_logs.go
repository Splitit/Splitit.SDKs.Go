/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// EmailAuditLogs struct for EmailAuditLogs
type EmailAuditLogs struct {
	Id int64 `json:"Id"` 
	InstallmentPlanId int64 `json:"InstallmentPlanId,omitempty"` 
	UserUniqueId string `json:"UserUniqueId,omitempty"` 
	UserId int64 `json:"UserId,omitempty"` 
	EmailAddress string `json:"EmailAddress,omitempty"` 
	ExternalIdentifier string `json:"ExternalIdentifier,omitempty"` 
	InstallmentPlan *InstallmentPlans `json:"InstallmentPlan,omitempty"` 
	UserType UserType `json:"UserType,omitempty"` 
	EmailType SystemEmailsTypes `json:"EmailType"` 
	EmailHtmlBodyFilename string `json:"EmailHtmlBodyFilename,omitempty"` 
}
