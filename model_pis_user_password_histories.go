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
// PisUserPasswordHistories struct for PisUserPasswordHistories
type PisUserPasswordHistories struct {
	Id int64 `json:"Id"` 
	Password string `json:"Password,omitempty"` 
	PasswordHash string `json:"PasswordHash,omitempty"` 
	PasswordExpDate *SplititTime `json:"PasswordExpDate"` 
	PisUserId int64 `json:"PisUserId"` 
	PisUser *PisUsers `json:"PisUser,omitempty"` 
}
