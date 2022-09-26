# \CardApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowidV1AuthorizeCardinfo**](CardApi.md#PaynowidV1AuthorizeCardinfo) | **Post** /paynowid/v1/Authorize/card | 決済の与信を行います
[**PaynowidV1CancelCardinfo**](CardApi.md#PaynowidV1CancelCardinfo) | **Post** /paynowid/v1/Cancel/card | 決済のキャンセルを行います
[**PaynowidV1CaptureCardinfo**](CardApi.md#PaynowidV1CaptureCardinfo) | **Post** /paynowid/v1/Capture/card | 決済の売上確定を行います
[**PaynowidV1ReAuthorizeCardinfo**](CardApi.md#PaynowidV1ReAuthorizeCardinfo) | **Post** /paynowid/v1/ReAuthorize/card | 決済の再与信を行います



## PaynowidV1AuthorizeCardinfo

> CardAuthorizeResponse PaynowidV1AuthorizeCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

決済の与信を行います

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
    resp, r, err := apiClient.CardApi.PaynowidV1AuthorizeCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowidV1AuthorizeCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowidV1AuthorizeCardinfo`: CardAuthorizeResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowidV1AuthorizeCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowidV1AuthorizeCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string** |  | 
 **authHash** | **string** |  | 

### Return type

[**CardAuthorizeResponse**](CardAuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowidV1CancelCardinfo

> CardCancelResponse PaynowidV1CancelCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

決済のキャンセルを行います

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
    resp, r, err := apiClient.CardApi.PaynowidV1CancelCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowidV1CancelCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowidV1CancelCardinfo`: CardCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowidV1CancelCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowidV1CancelCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string** |  | 
 **authHash** | **string** |  | 

### Return type

[**CardCancelResponse**](CardCancelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowidV1CaptureCardinfo

> CardCaptureResponse PaynowidV1CaptureCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

決済の売上確定を行います

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
    resp, r, err := apiClient.CardApi.PaynowidV1CaptureCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowidV1CaptureCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowidV1CaptureCardinfo`: CardCaptureResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowidV1CaptureCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowidV1CaptureCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string** |  | 
 **authHash** | **string** |  | 

### Return type

[**CardCaptureResponse**](CardCaptureResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowidV1ReAuthorizeCardinfo

> CardAuthorizeResponse PaynowidV1ReAuthorizeCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

決済の再与信を行います

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
    resp, r, err := apiClient.CardApi.PaynowidV1ReAuthorizeCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowidV1ReAuthorizeCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowidV1ReAuthorizeCardinfo`: CardAuthorizeResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowidV1ReAuthorizeCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowidV1ReAuthorizeCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string** |  | 
 **authHash** | **string** |  | 

### Return type

[**CardAuthorizeResponse**](CardAuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

