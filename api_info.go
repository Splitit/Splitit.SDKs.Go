/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// InfoApiService InfoApi service
type InfoApiService service

/*
InfoGetLearnMoreDetails Method for InfoGetLearnMoreDetails
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param request
@return LearnMoreDetailsResponse
*/
func (a *InfoApiService) InfoGetLearnMoreDetails(ctx _context.Context, request LearnMoreDetailsRequest) (LearnMoreDetailsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LearnMoreDetailsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/Merchant/GetLearnMoreDetails"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &request
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v LearnMoreDetailsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarReturnValue.ResponseHeader.Succeeded != true {
		if len(*localVarReturnValue.ResponseHeader.Errors) > 0 {
			newErr := GenericOpenAPIError{
				body:  localVarBody,
				error: (*localVarReturnValue.ResponseHeader.Errors)[0].Message,
				model: localVarReturnValue.ResponseHeader,
			}
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: "Error ocurred while contacting Splitit API.",
			model: localVarReturnValue.ResponseHeader,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}