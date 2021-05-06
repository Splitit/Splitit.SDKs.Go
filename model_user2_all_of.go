/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// User2AllOf struct for User2AllOf
type User2AllOf struct {
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
}
