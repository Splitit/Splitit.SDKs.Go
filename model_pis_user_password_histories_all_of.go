/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PisUserPasswordHistoriesAllOf struct for PisUserPasswordHistoriesAllOf
type PisUserPasswordHistoriesAllOf struct {
	PisUserId int64 `json:"PisUserId"` 
	Password string `json:"Password,omitempty"` 
	PasswordHash string `json:"PasswordHash,omitempty"` 
	PisUser *PisUsers `json:"PisUser,omitempty"` 
}