/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// Schedule struct for Schedule
type Schedule struct {
	NumberOfInstallments int32 `json:"NumberOfInstallments"` 
	Deposit bool `json:"Deposit"` 
	Elements []ScheduleElements `json:"Elements,omitempty"` 
}
