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
// PisSessionsAllOf struct for PisSessionsAllOf
type PisSessionsAllOf struct {
	SessionId string `json:"SessionId,omitempty"` 
	LastAccessTime *SplititTime `json:"LastAccessTime"` 
	UserId int64 `json:"UserId,omitempty"` 
	UserUniqueId string `json:"UserUniqueId,omitempty"` 
	InstallmentPlanNumber string `json:"InstallmentPlanNumber,omitempty"` 
	VersionedTouchPointId int64 `json:"VersionedTouchPointId,omitempty"` 
	SessionValidPeriodInMinutes int32 `json:"SessionValidPeriodInMinutes"` 
	VersionedTouchPoint *VersionedTouchPoints `json:"VersionedTouchPoint,omitempty"` 
	UserType UserType `json:"UserType,omitempty"` 
	SessionAvailibility SessionAvailibility `json:"SessionAvailibility"` 
}
