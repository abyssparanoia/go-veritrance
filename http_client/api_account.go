/*
Veritrance API

Veritrance API 

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

type ApiPaynowV2AddAccountRequest struct {
	ctx context.Context
	ApiService *AccountApiService
	accountId *string
	cardNumber *string
	cardExpire *string
	securityCode *string
}

func (r ApiPaynowV2AddAccountRequest) AccountId(accountId string) ApiPaynowV2AddAccountRequest {
	r.accountId = &accountId
	return r
}

func (r ApiPaynowV2AddAccountRequest) CardNumber(cardNumber string) ApiPaynowV2AddAccountRequest {
	r.cardNumber = &cardNumber
	return r
}

func (r ApiPaynowV2AddAccountRequest) CardExpire(cardExpire string) ApiPaynowV2AddAccountRequest {
	r.cardExpire = &cardExpire
	return r
}

func (r ApiPaynowV2AddAccountRequest) SecurityCode(securityCode string) ApiPaynowV2AddAccountRequest {
	r.securityCode = &securityCode
	return r
}

func (r ApiPaynowV2AddAccountRequest) Execute() (*AccountResponse, *http.Response, error) {
	return r.ApiService.PaynowV2AddAccountExecute(r)
}

/*
PaynowV2AddAccount 会員情報を追加します。

・会員 ID を登録します。
・会員 ID の登録と同時に、カード情報および継続課金情報を紐付けることが可能です。
  継続課金情報を設定する際は、カード情報の同時設定が必要です。
  カード情報を指定した場合、有効性確認が行われます。
・会員 ID の登録と同時に、カード情報のみを紐付けることが可能です。


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV2AddAccountRequest
*/
func (a *AccountApiService) PaynowV2AddAccount(ctx context.Context) ApiPaynowV2AddAccountRequest {
	return ApiPaynowV2AddAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountResponse
func (a *AccountApiService) PaynowV2AddAccountExecute(r ApiPaynowV2AddAccountRequest) (*AccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.PaynowV2AddAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v2/Add/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, reportError("accountId is required and must be specified")
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
	localVarFormParams.Add("accountId", parameterToString(*r.accountId, ""))
	if r.cardNumber != nil {
		localVarFormParams.Add("cardNumber", parameterToString(*r.cardNumber, ""))
	}
	if r.cardExpire != nil {
		localVarFormParams.Add("cardExpire", parameterToString(*r.cardExpire, ""))
	}
	if r.securityCode != nil {
		localVarFormParams.Add("securityCode", parameterToString(*r.securityCode, ""))
	}
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

type ApiPaynowV2DeleteAccountRequest struct {
	ctx context.Context
	ApiService *AccountApiService
	accountId *string
}

func (r ApiPaynowV2DeleteAccountRequest) AccountId(accountId string) ApiPaynowV2DeleteAccountRequest {
	r.accountId = &accountId
	return r
}

func (r ApiPaynowV2DeleteAccountRequest) Execute() (*PayNowIdResponse, *http.Response, error) {
	return r.ApiService.PaynowV2DeleteAccountExecute(r)
}

/*
PaynowV2DeleteAccount 会員 ID の会員情報を、指定された「退会年月日」に削除します。

・会員 ID の会員情報を、指定された「退会年月日」に削除します。


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV2DeleteAccountRequest
*/
func (a *AccountApiService) PaynowV2DeleteAccount(ctx context.Context) ApiPaynowV2DeleteAccountRequest {
	return ApiPaynowV2DeleteAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PayNowIdResponse
func (a *AccountApiService) PaynowV2DeleteAccountExecute(r ApiPaynowV2DeleteAccountRequest) (*PayNowIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayNowIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.PaynowV2DeleteAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v2/Delete/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, reportError("accountId is required and must be specified")
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
	localVarFormParams.Add("accountId", parameterToString(*r.accountId, ""))
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

type ApiPaynowV2UpdateAccountRequest struct {
	ctx context.Context
	ApiService *AccountApiService
	accountId *string
}

func (r ApiPaynowV2UpdateAccountRequest) AccountId(accountId string) ApiPaynowV2UpdateAccountRequest {
	r.accountId = &accountId
	return r
}

func (r ApiPaynowV2UpdateAccountRequest) Execute() (*PayNowIdResponse, *http.Response, error) {
	return r.ApiService.PaynowV2UpdateAccountExecute(r)
}

/*
PaynowV2UpdateAccount 会員 ID の「入会年月日」を更新します。

・会員 ID の「入会年月日」を更新します。


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaynowV2UpdateAccountRequest
*/
func (a *AccountApiService) PaynowV2UpdateAccount(ctx context.Context) ApiPaynowV2UpdateAccountRequest {
	return ApiPaynowV2UpdateAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PayNowIdResponse
func (a *AccountApiService) PaynowV2UpdateAccountExecute(r ApiPaynowV2UpdateAccountRequest) (*PayNowIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayNowIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountApiService.PaynowV2UpdateAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paynow/v2/Update/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, reportError("accountId is required and must be specified")
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
	localVarFormParams.Add("accountId", parameterToString(*r.accountId, ""))
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
