/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ApiUsersAllOf struct for ApiUsersAllOf
type ApiUsersAllOf struct {
	CrmId string `json:"CrmId,omitempty"` 
	Merchant *Merchants `json:"Merchant,omitempty"` 
	ApiUserPasswordHistories []ApiUserPasswordHistories `json:"ApiUserPasswordHistories,omitempty"` 
	MerchantId int64 `json:"MerchantId"` 
	MerchantName string `json:"MerchantName,omitempty"` 
	Password *IUserPasswordHistory `json:"Password,omitempty"` 
}
