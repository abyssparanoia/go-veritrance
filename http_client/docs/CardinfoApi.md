# \CardinfoApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowidV1AddCardinfo**](CardinfoApi.md#PaynowidV1AddCardinfo) | **Post** /paynowid/v1/Add/cardinfo | 会員 ID にカード情報を紐付けて登録します
[**PaynowidV1DeleteCardinfo**](CardinfoApi.md#PaynowidV1DeleteCardinfo) | **Post** /paynowid/v1/Delete/cardinfo | 会員 ID に紐付けられたカード情報を削除します
[**PaynowidV1GetCardinfo**](CardinfoApi.md#PaynowidV1GetCardinfo) | **Post** /paynowid/v1/Get/cardinfo | 会員 ID に紐付けられた課金情報を取得します
[**PaynowidV1UpdateCardinfo**](CardinfoApi.md#PaynowidV1UpdateCardinfo) | **Post** /paynowid/v1/Update/cardinfo | 当該会員 ID、およびカード ID に紐付けられたカード情報を更新します



## PaynowidV1AddCardinfo

> AccountResponse PaynowidV1AddCardinfo(ctx).Body(body).Execute()

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardinfoApi.PaynowidV1AddCardinfo(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowidV1AddCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowidV1AddCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowidV1AddCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowidV1AddCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowidV1DeleteCardinfo

> PayNowIdResponse PaynowidV1DeleteCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    resp, r, err := apiClient.CardinfoApi.PaynowidV1DeleteCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowidV1DeleteCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowidV1DeleteCardinfo`: PayNowIdResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowidV1DeleteCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowidV1DeleteCardinfoRequest struct via the builder pattern


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


## PaynowidV1GetCardinfo

> AccountResponse PaynowidV1GetCardinfo(ctx).Body(body).Execute()

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardinfoApi.PaynowidV1GetCardinfo(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowidV1GetCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowidV1GetCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowidV1GetCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowidV1GetCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowidV1UpdateCardinfo

> AccountResponse PaynowidV1UpdateCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    resp, r, err := apiClient.CardinfoApi.PaynowidV1UpdateCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowidV1UpdateCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowidV1UpdateCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowidV1UpdateCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowidV1UpdateCardinfoRequest struct via the builder pattern


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

