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
// PisUsers struct for PisUsers
type PisUsers struct {
	Id int64 `json:"Id"` 
	Culture string `json:"Culture,omitempty"` 
	UserFullName string `json:"UserFullName,omitempty"` 
	PhoneNumber string `json:"PhoneNumber,omitempty"` 
	Role *Roles `json:"Role,omitempty"` 
	UserName string `json:"UserName,omitempty"` 
	UniqueId string `json:"UniqueId,omitempty"` 
	LoginAttempt int32 `json:"LoginAttempt"` 
	MaxInvalidLoginAttempts int32 `json:"MaxInvalidLoginAttempts"` 
	Status UserStatus `json:"Status"` 
	IsLocked bool `json:"IsLocked"` 
	Email string `json:"Email,omitempty"` 
	CultureName string `json:"CultureName,omitempty"` 
	RoleId int64 `json:"RoleId"` 
	LoginUserName string `json:"LoginUserName,omitempty"` 
	LastTimeLogin *SplititTime `json:"LastTimeLogin"` 
	IsDataPrivateRestricted bool `json:"IsDataPrivateRestricted"` 
	FirstName string `json:"FirstName,omitempty"` 
	LastName string `json:"LastName,omitempty"` 
	JobTitle string `json:"JobTitle,omitempty"` 
	FaxNumber string `json:"FaxNumber,omitempty"` 
	MobileNumber string `json:"MobileNumber,omitempty"` 
	WorkPhoneNumber string `json:"WorkPhoneNumber,omitempty"` 
	Notes string `json:"Notes,omitempty"` 
	Type UserPermissionLevel `json:"Type"` 
	CrmId string `json:"CrmId,omitempty"` 
	Agents []Agents `json:"Agents,omitempty"` 
	PisUserBusinessUnits []PisUserBusinessUnits `json:"PisUserBusinessUnits,omitempty"` 
	PisUserPasswordHistories []PisUserPasswordHistories `json:"PisUserPasswordHistories,omitempty"` 
	UserJobSubscriptions []UserJobSubscriptions `json:"UserJobSubscriptions,omitempty"` 
	Password *IUserPasswordHistory `json:"Password,omitempty"` 
}
