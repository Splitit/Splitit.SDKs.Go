/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// VersionedTouchPoints struct for VersionedTouchPoints
type VersionedTouchPoints struct {
	Id int64 `json:"Id"` 
	Version string `json:"Version,omitempty"` 
	TouchPointId int64 `json:"TouchPointId"` 
	SubVersion string `json:"SubVersion,omitempty"` 
	TouchPoint *TouchPoints `json:"TouchPoint,omitempty"` 
	PisSessions []PisSessions `json:"PisSessions,omitempty"` 
	InstallmentPlanAuditLogs []InstallmentPlanAuditLogs `json:"InstallmentPlanAuditLogs,omitempty"` 
}
