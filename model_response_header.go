/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ResponseHeader struct for ResponseHeader
type ResponseHeader struct {
	Succeeded bool `json:"Succeeded"` 
	Errors []Error `json:"Errors,omitempty"` 
	TraceId string `json:"TraceId,omitempty"` 
}
