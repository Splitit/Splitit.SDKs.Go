/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// OperationType 
type OperationType string

// List of OperationType
const (
	OPERATIONTYPE_NONE OperationType = "None"
	OPERATIONTYPE_AUTHORIZE OperationType = "Authorize"
	OPERATIONTYPE_CAPTURE OperationType = "Capture"
	OPERATIONTYPE_VOID OperationType = "Void"
	OPERATIONTYPE_REFUND OperationType = "Refund"
	OPERATIONTYPE_CANCEL OperationType = "Cancel"
	OPERATIONTYPE_CHECK_CAPTURE_ASYNC_RESPONSE OperationType = "CheckCaptureAsyncResponse"
	OPERATIONTYPE_VERIFY_CARD OperationType = "VerifyCard"
)
