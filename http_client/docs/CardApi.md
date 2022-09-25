# \CardApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowV2AuthorizeCardinfo**](CardApi.md#PaynowV2AuthorizeCardinfo) | **Post** /paynow/v2/Authorize/card | 決済の与信を行います
[**PaynowV2CancelCardinfo**](CardApi.md#PaynowV2CancelCardinfo) | **Post** /paynow/v2/Cancel/card | 決済のキャンセルを行います
[**PaynowV2CaptureCardinfo**](CardApi.md#PaynowV2CaptureCardinfo) | **Post** /paynow/v2/Capture/card | 決済の売上確定を行います
[**PaynowV2ReAuthorizeCardinfo**](CardApi.md#PaynowV2ReAuthorizeCardinfo) | **Post** /paynow/v2/ReAuthorize/card | 決済の再与信を行います



## PaynowV2AuthorizeCardinfo

> CardAuthorizeResponse PaynowV2AuthorizeCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    params := TODO // CardAuthorizeRequest |  (optional)
    authHash := "authHash_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2AuthorizeCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowV2AuthorizeCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2AuthorizeCardinfo`: CardAuthorizeResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowV2AuthorizeCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2AuthorizeCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | [**CardAuthorizeRequest**](CardAuthorizeRequest.md) |  | 
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


## PaynowV2CancelCardinfo

> CardCancelResponse PaynowV2CancelCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    params := TODO // CardCancelRequest |  (optional)
    authHash := "authHash_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2CancelCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowV2CancelCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2CancelCardinfo`: CardCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowV2CancelCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2CancelCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | [**CardCancelRequest**](CardCancelRequest.md) |  | 
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


## PaynowV2CaptureCardinfo

> CardCaptureResponse PaynowV2CaptureCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    params := TODO // CardCaptureRequest |  (optional)
    authHash := "authHash_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2CaptureCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowV2CaptureCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2CaptureCardinfo`: CardCaptureResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowV2CaptureCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2CaptureCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | [**CardCaptureRequest**](CardCaptureRequest.md) |  | 
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


## PaynowV2ReAuthorizeCardinfo

> CardAuthorizeResponse PaynowV2ReAuthorizeCardinfo(ctx).Params(params).AuthHash(authHash).Execute()

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
    params := TODO // CardReAuthorizeRequest |  (optional)
    authHash := "authHash_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2ReAuthorizeCardinfo(context.Background()).Params(params).AuthHash(authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardApi.PaynowV2ReAuthorizeCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2ReAuthorizeCardinfo`: CardAuthorizeResponse
    fmt.Fprintf(os.Stdout, "Response from `CardApi.PaynowV2ReAuthorizeCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2ReAuthorizeCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | [**CardReAuthorizeRequest**](CardReAuthorizeRequest.md) |  | 
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

