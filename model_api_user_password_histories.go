/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ApiUserPasswordHistories struct for ApiUserPasswordHistories
type ApiUserPasswordHistories struct {
	Id int64 `json:"Id"` 
	Password string `json:"Password,omitempty"` 
	PasswordHash string `json:"PasswordHash,omitempty"` 
	ApiUserId int64 `json:"ApiUserId"` 
	ApiUser *ApiUsers `json:"ApiUser,omitempty"` 
}
