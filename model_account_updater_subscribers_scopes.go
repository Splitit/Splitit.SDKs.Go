/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// AccountUpdaterSubscribersScopes struct for AccountUpdaterSubscribersScopes
type AccountUpdaterSubscribersScopes struct {
	Id int64 `json:"Id"` 
	TerminalId int64 `json:"TerminalId"` 
	IssuingCountryIso string `json:"IssuingCountryIso,omitempty"` 
	Terminal *Terminals `json:"Terminal,omitempty"` 
	AccountUpdaterProvider AccountUpdaterProvider `json:"AccountUpdaterProvider"` 
	CardBrand CardBrand `json:"CardBrand"` 
}
