/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// AddressData struct for AddressData
type AddressData struct {
	AddressLine string `json:"AddressLine,omitempty"`
	AddressLine2 string `json:"AddressLine2,omitempty"`
	City string `json:"City,omitempty"`
	Country string `json:"Country,omitempty"`
	State string `json:"State,omitempty"`
	Zip string `json:"Zip,omitempty"`
	FullAddressLine string `json:"FullAddressLine,omitempty"`
}