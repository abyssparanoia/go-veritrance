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


// AccountApiService AccountApi service
type AccountApiService service

type ApiPaynowV1AddAccountRequest struct {
	ctx context.Context
	ApiService *AccountApiService
	params *string
	authHash *string
}

func (r ApiPaynowV1AddAccountRequest) Params(params string) ApiPaynowV1AddAccountRequest {
	r.params = &params
	return r
}

func (r ApiPaynowV1AddAccountRequest) AuthHash(authHash string) ApiPaynowV1AddAccountRequest {
	r.authHash = &authHash
	return r
}

func (r ApiPaynowV1AddAccountRequest) Execute() (*AccountResponse, *http.Response, error) {
	return r.ApiService.PaynowV1AddAccountExecute(r)
}

/*
PaynowV1AddAccount 会員情報を追加します。

・会員 ID を登録します。
・会員 ID の登録と同時に、カード情報および継続課金情報を紐付けることが可能です。
  継続課金情報を設定する際は、カード情報の同時設定が必要です。
  カード情報を指定した場合、有効性確認が行われます。
・会員 ID の登録と同時に、カード情報のみを紐付けることが可能です。


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV1AddAccountRequest
*/
func (a *AccountApiService) PaynowV1AddAccount(ctx context.Context) ApiPaynowV1AddAccountRequest {
	return ApiPaynowV1AddAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountResponse
func (a *AccountApiService) PaynowV1AddAccountExecute(r ApiPaynowV1AddAccountRequest) (*AccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.PaynowV1AddAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v1/Add/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.params == nil {
		return localVarReturnValue, nil, reportError("params is required and must be specified")
	}
	if r.authHash == nil {
		return localVarReturnValue, nil, reportError("authHash is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("params", parameterToString(*r.params, ""))
	localVarFormParams.Add("authHash", parameterToString(*r.authHash, ""))
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

type ApiPaynowV1DeleteAccountRequest struct {
	ctx context.Context
	ApiService *AccountApiService
	params *string
	authHash *string
}

func (r ApiPaynowV1DeleteAccountRequest) Params(params string) ApiPaynowV1DeleteAccountRequest {
	r.params = &params
	return r
}

func (r ApiPaynowV1DeleteAccountRequest) AuthHash(authHash string) ApiPaynowV1DeleteAccountRequest {
	r.authHash = &authHash
	return r
}

func (r ApiPaynowV1DeleteAccountRequest) Execute() (*PayNowIdResponse, *http.Response, error) {
	return r.ApiService.PaynowV1DeleteAccountExecute(r)
}

/*
PaynowV1DeleteAccount 会員 ID の会員情報を、指定された「退会年月日」に削除します。

・会員 ID の会員情報を、指定された「退会年月日」に削除します。


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV1DeleteAccountRequest
*/
func (a *AccountApiService) PaynowV1DeleteAccount(ctx context.Context) ApiPaynowV1DeleteAccountRequest {
	return ApiPaynowV1DeleteAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PayNowIdResponse
func (a *AccountApiService) PaynowV1DeleteAccountExecute(r ApiPaynowV1DeleteAccountRequest) (*PayNowIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayNowIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.PaynowV1DeleteAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v1/Delete/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.params == nil {
		return localVarReturnValue, nil, reportError("params is required and must be specified")
	}
	if r.authHash == nil {
		return localVarReturnValue, nil, reportError("authHash is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("params", parameterToString(*r.params, ""))
	localVarFormParams.Add("authHash", parameterToString(*r.authHash, ""))
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

type ApiPaynowV1UpdateAccountRequest struct {
	ctx context.Context
	ApiService *AccountApiService
	params *string
	authHash *string
}

func (r ApiPaynowV1UpdateAccountRequest) Params(params string) ApiPaynowV1UpdateAccountRequest {
	r.params = &params
	return r
}

func (r ApiPaynowV1UpdateAccountRequest) AuthHash(authHash string) ApiPaynowV1UpdateAccountRequest {
	r.authHash = &authHash
	return r
}

func (r ApiPaynowV1UpdateAccountRequest) Execute() (*PayNowIdResponse, *http.Response, error) {
	return r.ApiService.PaynowV1UpdateAccountExecute(r)
}

/*
PaynowV1UpdateAccount 会員 ID の「入会年月日」を更新します。

・会員 ID の「入会年月日」を更新します。


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV1UpdateAccountRequest
*/
func (a *AccountApiService) PaynowV1UpdateAccount(ctx context.Context) ApiPaynowV1UpdateAccountRequest {
	return ApiPaynowV1UpdateAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PayNowIdResponse
func (a *AccountApiService) PaynowV1UpdateAccountExecute(r ApiPaynowV1UpdateAccountRequest) (*PayNowIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayNowIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.PaynowV1UpdateAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v1/Update/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.params == nil {
		return localVarReturnValue, nil, reportError("params is required and must be specified")
	}
	if r.authHash == nil {
		return localVarReturnValue, nil, reportError("authHash is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("params", parameterToString(*r.params, ""))
	localVarFormParams.Add("authHash", parameterToString(*r.authHash, ""))
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
