/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// FraudCheckResult 
type FraudCheckResult string

// List of FraudCheckResult
const (
	FRAUDCHECKRESULT_FAILED FraudCheckResult = "Failed"
	FRAUDCHECKRESULT_APPROVED FraudCheckResult = "Approved"
	FRAUDCHECKRESULT_REJECTED FraudCheckResult = "Rejected"
	FRAUDCHECKRESULT_NOT_REVIEWED FraudCheckResult = "NotReviewed"
)
