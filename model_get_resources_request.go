/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// GetResourcesRequest struct for GetResourcesRequest
type GetResourcesRequest struct {
	SystemTextCategories *[]SystemTextCategory `json:"SystemTextCategories,omitempty"`
	RequestContext *GetResourcesRequestContext `json:"RequestContext,omitempty"`
}
