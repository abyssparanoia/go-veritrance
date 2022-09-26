# \AccountApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowV1AddAccount**](AccountApi.md#PaynowV1AddAccount) | **Post** /paynow/v1/Add/account | 会員情報を追加します。
[**PaynowV1DeleteAccount**](AccountApi.md#PaynowV1DeleteAccount) | **Post** /paynow/v1/Delete/account | 会員 ID の会員情報を、指定された「退会年月日」に削除します。
[**PaynowV1UpdateAccount**](AccountApi.md#PaynowV1UpdateAccount) | **Post** /paynow/v1/Update/account | 会員 ID の「入会年月日」を更新します。



## PaynowV1AddAccount

> AccountResponse PaynowV1AddAccount(ctx).Params(params).AuthHash(authHash).Execute()

会員情報を追加します。



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    params := "params_example" // string | 
    authHash := "authHash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.PaynowV1AddAccount(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.PaynowV1AddAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1AddAccount`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.PaynowV1AddAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1AddAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string** |  | 
 **authHash** | **string** |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowV1DeleteAccount

> PayNowIdResponse PaynowV1DeleteAccount(ctx).Params(params).AuthHash(authHash).Execute()

会員 ID の会員情報を、指定された「退会年月日」に削除します。



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    params := "params_example" // string | 
    authHash := "authHash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.PaynowV1DeleteAccount(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.PaynowV1DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1DeleteAccount`: PayNowIdResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.PaynowV1DeleteAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string** |  | 
 **authHash** | **string** |  | 

### Return type

[**PayNowIdResponse**](PayNowIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowV1UpdateAccount

> PayNowIdResponse PaynowV1UpdateAccount(ctx).Params(params).AuthHash(authHash).Execute()

会員 ID の「入会年月日」を更新します。



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    params := "params_example" // string | 
    authHash := "authHash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.PaynowV1UpdateAccount(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.PaynowV1UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1UpdateAccount`: PayNowIdResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.PaynowV1UpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1UpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string** |  | 
 **authHash** | **string** |  | 

### Return type

[**PayNowIdResponse**](PayNowIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

