# \CardApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowV1AuthorizeCardinfo**](CardApi.md#PaynowV1AuthorizeCardinfo) | **Post** /paynow/v1/Authorize/card | 決済の与信を行います
[**PaynowV1CancelCardinfo**](CardApi.md#PaynowV1CancelCardinfo) | **Post** /paynow/v1/Cancel/card | 決済のキャンセルを行います
[**PaynowV1CaptureCardinfo**](CardApi.md#PaynowV1CaptureCardinfo) | **Post** /paynow/v1/Capture/card | 決済の売上確定を行います
[**PaynowV1ReAuthorizeCardinfo**](CardApi.md#PaynowV1ReAuthorizeCardinfo) | **Post** /paynow/v1/ReAuthorize/card | 決済の再与信を行います



## PaynowV1AuthorizeCardinfo

> CardAuthorizeResponse PaynowV1AuthorizeCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    resp, r, err := apiClient.CardApi.PaynowV1AuthorizeCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowV1AuthorizeCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1AuthorizeCardinfo`: CardAuthorizeResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowV1AuthorizeCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1AuthorizeCardinfoRequest struct via the builder pattern


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


## PaynowV1CancelCardinfo

> CardCancelResponse PaynowV1CancelCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    resp, r, err := apiClient.CardApi.PaynowV1CancelCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowV1CancelCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1CancelCardinfo`: CardCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowV1CancelCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1CancelCardinfoRequest struct via the builder pattern


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


## PaynowV1CaptureCardinfo

> CardCaptureResponse PaynowV1CaptureCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    resp, r, err := apiClient.CardApi.PaynowV1CaptureCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowV1CaptureCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1CaptureCardinfo`: CardCaptureResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowV1CaptureCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1CaptureCardinfoRequest struct via the builder pattern


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


## PaynowV1ReAuthorizeCardinfo

> CardAuthorizeResponse PaynowV1ReAuthorizeCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    resp, r, err := apiClient.CardApi.PaynowV1ReAuthorizeCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowV1ReAuthorizeCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV1ReAuthorizeCardinfo`: CardAuthorizeResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowV1ReAuthorizeCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV1ReAuthorizeCardinfoRequest struct via the builder pattern


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

