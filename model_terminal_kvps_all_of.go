/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// TerminalKvpsAllOf struct for TerminalKvpsAllOf
type TerminalKvpsAllOf struct {
	Key string `json:"Key,omitempty"` 
	Value string `json:"Value,omitempty"` 
	TerminalId int64 `json:"TerminalId"` 
	Terminal *Terminals `json:"Terminal,omitempty"` 
}
