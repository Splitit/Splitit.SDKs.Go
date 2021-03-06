/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// TouchPointColorValues struct for TouchPointColorValues
type TouchPointColorValues struct {
	Id int64 `json:"Id"` 
	TouchPointId int64 `json:"TouchPointId"` 
	MerchantId int64 `json:"MerchantId,omitempty"` 
	Merchant *Merchants `json:"Merchant,omitempty"` 
	Colors []ConfigValues `json:"Colors,omitempty"` 
	TouchPoint *TouchPoints `json:"TouchPoint,omitempty"` 
}
