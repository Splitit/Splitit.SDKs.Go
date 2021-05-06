/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// MerchantStatus 
type MerchantStatus string

// List of MerchantStatus
const (
	MERCHANTSTATUS_SELF_ON_BOARDING MerchantStatus = "SelfOnBoarding"
	MERCHANTSTATUS_SIGNED_UP MerchantStatus = "SignedUp"
	MERCHANTSTATUS_SUBMITTED MerchantStatus = "Submitted"
	MERCHANTSTATUS_TESTED MerchantStatus = "Tested"
	MERCHANTSTATUS_CERTIFIED MerchantStatus = "Certified"
	MERCHANTSTATUS_LIVE MerchantStatus = "Live"
	MERCHANTSTATUS_END MerchantStatus = "End"
	MERCHANTSTATUS_NONE MerchantStatus = "None"
)
