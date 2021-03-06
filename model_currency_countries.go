/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// CurrencyCountries struct for CurrencyCountries
type CurrencyCountries struct {
	CountryId int64 `json:"CountryId"` 
	CurrencyId int64 `json:"CurrencyId"` 
	Country *Countries `json:"Country,omitempty"` 
	Currency *Currencies `json:"Currency,omitempty"` 
}
