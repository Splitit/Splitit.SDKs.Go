/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ResponseStatus struct for ResponseStatus
type ResponseStatus struct {
	ErrorCode string `json:"ErrorCode,omitempty"`
	Message string `json:"Message,omitempty"`
	StackTrace string `json:"StackTrace,omitempty"`
	Errors *[]ResponseError `json:"Errors,omitempty"`
}
