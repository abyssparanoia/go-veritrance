# \CardinfoApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowV1AddCardinfo**](CardinfoApi.md#PaynowV1AddCardinfo) | **Post** /paynow/v1/Add/cardinfo | 会員 ID にカード情報を紐付けて登録します
[**PaynowV1DeleteCardinfo**](CardinfoApi.md#PaynowV1DeleteCardinfo) | **Post** /paynow/v1/Delete/cardinfo | 会員 ID に紐付けられたカード情報を削除します
[**PaynowV1GetCardinfo**](CardinfoApi.md#PaynowV1GetCardinfo) | **Post** /paynow/v1/Get/cardinfo | 会員 ID に紐付けられた課金情報を取得します
[**PaynowV1UpdateCardinfo**](CardinfoApi.md#PaynowV1UpdateCardinfo) | **Post** /paynow/v1/Update/cardinfo | 当該会員 ID、およびカード ID に紐付けられたカード情報を更新します



## PaynowV1AddCardinfo

> AccountResponse PaynowV1AddCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

会員 ID にカード情報を紐付けて登録します



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
    resp, r, err := apiClient.CardinfoApi.PaynowV1AddCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowV1AddCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1AddCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowV1AddCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1AddCardinfoRequest struct via the builder pattern


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


## PaynowV1DeleteCardinfo

> PayNowIdResponse PaynowV1DeleteCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

会員 ID に紐付けられたカード情報を削除します



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
    resp, r, err := apiClient.CardinfoApi.PaynowV1DeleteCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowV1DeleteCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1DeleteCardinfo`: PayNowIdResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowV1DeleteCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1DeleteCardinfoRequest struct via the builder pattern


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


## PaynowV1GetCardinfo

> AccountResponse PaynowV1GetCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

会員 ID に紐付けられた課金情報を取得します



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
    resp, r, err := apiClient.CardinfoApi.PaynowV1GetCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowV1GetCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1GetCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowV1GetCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1GetCardinfoRequest struct via the builder pattern


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


## PaynowV1UpdateCardinfo

> AccountResponse PaynowV1UpdateCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

当該会員 ID、およびカード ID に紐付けられたカード情報を更新します



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
    resp, r, err := apiClient.CardinfoApi.PaynowV1UpdateCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowV1UpdateCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1UpdateCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowV1UpdateCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1UpdateCardinfoRequest struct via the builder pattern


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

