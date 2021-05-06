/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// UserJobSubscriptions struct for UserJobSubscriptions
type UserJobSubscriptions struct {
	Id int64 `json:"Id"` 
	PisUserUniqueId string `json:"PisUserUniqueId,omitempty"` 
	PisUserId int64 `json:"PisUserId"` 
	SplititJobDefinitionId int64 `json:"SplititJobDefinitionId"` 
	PisUser *PisUsers `json:"PisUser,omitempty"` 
	SplititJobDefinition *SplititJobDefinitions `json:"SplititJobDefinition,omitempty"` 
}
