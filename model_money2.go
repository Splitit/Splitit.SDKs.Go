/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// Money2 struct for Money2
type Money2 struct {
	Amount float64 `json:"Amount"` 
	Currency *Currencies `json:"Currency,omitempty"` 
}