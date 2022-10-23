/*
Veritrance API

Veritrance API Requestの形式はcomponentsを参照。文字列にしたものを各Requestのparamsに用いる 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package veritrance_paynow

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// CardApiService CardApi service
type CardApiService service

type ApiPaynowV1AuthorizeCardinfoRequest struct {
	ctx context.Context
	ApiService *CardApiService
	body *string
}

func (r ApiPaynowV1AuthorizeCardinfoRequest) Body(body string) ApiPaynowV1AuthorizeCardinfoRequest {
	r.body = &body
	return r
}

func (r ApiPaynowV1AuthorizeCardinfoRequest) Execute() (*CardAuthorizeResponse, *http.Response, error) {
	return r.ApiService.PaynowV1AuthorizeCardinfoExecute(r)
}

/*
PaynowV1AuthorizeCardinfo 決済の与信を行います

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV1AuthorizeCardinfoRequest
*/
func (a *CardApiService) PaynowV1AuthorizeCardinfo(ctx context.Context) ApiPaynowV1AuthorizeCardinfoRequest {
	return ApiPaynowV1AuthorizeCardinfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CardAuthorizeResponse
func (a *CardApiService) PaynowV1AuthorizeCardinfoExecute(r ApiPaynowV1AuthorizeCardinfoRequest) (*CardAuthorizeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CardAuthorizeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardApiService.PaynowV1AuthorizeCardinfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v2/Authorize/card"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	fmt.Printf("client.decode", string(localVarBody))
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPaynowV1CancelCardinfoRequest struct {
	ctx context.Context
	ApiService *CardApiService
	body *string
}

func (r ApiPaynowV1CancelCardinfoRequest) Body(body string) ApiPaynowV1CancelCardinfoRequest {
	r.body = &body
	return r
}

func (r ApiPaynowV1CancelCardinfoRequest) Execute() (*CardCancelResponse, *http.Response, error) {
	return r.ApiService.PaynowV1CancelCardinfoExecute(r)
}

/*
PaynowV1CancelCardinfo 決済のキャンセルを行います

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV1CancelCardinfoRequest
*/
func (a *CardApiService) PaynowV1CancelCardinfo(ctx context.Context) ApiPaynowV1CancelCardinfoRequest {
	return ApiPaynowV1CancelCardinfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CardCancelResponse
func (a *CardApiService) PaynowV1CancelCardinfoExecute(r ApiPaynowV1CancelCardinfoRequest) (*CardCancelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CardCancelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardApiService.PaynowV1CancelCardinfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v1/Cancel/card"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPaynowV1CaptureCardinfoRequest struct {
	ctx context.Context
	ApiService *CardApiService
	body *string
}

func (r ApiPaynowV1CaptureCardinfoRequest) Body(body string) ApiPaynowV1CaptureCardinfoRequest {
	r.body = &body
	return r
}

func (r ApiPaynowV1CaptureCardinfoRequest) Execute() (*CardCaptureResponse, *http.Response, error) {
	return r.ApiService.PaynowV1CaptureCardinfoExecute(r)
}

/*
PaynowV1CaptureCardinfo 決済の売上確定を行います

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV1CaptureCardinfoRequest
*/
func (a *CardApiService) PaynowV1CaptureCardinfo(ctx context.Context) ApiPaynowV1CaptureCardinfoRequest {
	return ApiPaynowV1CaptureCardinfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CardCaptureResponse
func (a *CardApiService) PaynowV1CaptureCardinfoExecute(r ApiPaynowV1CaptureCardinfoRequest) (*CardCaptureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CardCaptureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardApiService.PaynowV1CaptureCardinfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v1/Capture/card"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPaynowV1ReAuthorizeCardinfoRequest struct {
	ctx context.Context
	ApiService *CardApiService
	body *string
}

func (r ApiPaynowV1ReAuthorizeCardinfoRequest) Body(body string) ApiPaynowV1ReAuthorizeCardinfoRequest {
	r.body = &body
	return r
}

func (r ApiPaynowV1ReAuthorizeCardinfoRequest) Execute() (*CardAuthorizeResponse, *http.Response, error) {
	return r.ApiService.PaynowV1ReAuthorizeCardinfoExecute(r)
}

/*
PaynowV1ReAuthorizeCardinfo 決済の再与信を行います

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV1ReAuthorizeCardinfoRequest
*/
func (a *CardApiService) PaynowV1ReAuthorizeCardinfo(ctx context.Context) ApiPaynowV1ReAuthorizeCardinfoRequest {
	return ApiPaynowV1ReAuthorizeCardinfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CardAuthorizeResponse
func (a *CardApiService) PaynowV1ReAuthorizeCardinfoExecute(r ApiPaynowV1ReAuthorizeCardinfoRequest) (*CardAuthorizeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CardAuthorizeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardApiService.PaynowV1ReAuthorizeCardinfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v1/ReAuthorize/card"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
