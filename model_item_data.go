/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ItemData struct for ItemData
type ItemData struct {
	Name string `json:"Name,omitempty"` 
	Sku string `json:"Sku,omitempty"` 
	Price *MoneyWithCurrencyCode `json:"Price,omitempty"` 
	Quantity float64 `json:"Quantity"` 
	Description string `json:"Description,omitempty"` 
}
