/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// UserPermissionLevel 
type UserPermissionLevel string

// List of UserPermissionLevel
const (
	USERPERMISSIONLEVEL_ADMINISTRATOR UserPermissionLevel = "Administrator"
	USERPERMISSIONLEVEL_MANAGER UserPermissionLevel = "Manager"
	USERPERMISSIONLEVEL_CASHIER UserPermissionLevel = "Cashier"
	USERPERMISSIONLEVEL_SPLITIT_ADMIN UserPermissionLevel = "SplititAdmin"
	USERPERMISSIONLEVEL_CHOOSE UserPermissionLevel = "Choose"
)
