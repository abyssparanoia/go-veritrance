# \CardApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowV1AuthorizeCardinfo**](CardApi.md#PaynowV1AuthorizeCardinfo) | **Post** /paynow/v2/Authorize/card | 決済の与信を行います
[**PaynowV1ReAuthorizeCardinfo**](CardApi.md#PaynowV1ReAuthorizeCardinfo) | **Post** /paynow/v1/ReAuthorize/card | 決済の再与信を行います
[**PaynowV2CancelCardinfo**](CardApi.md#PaynowV2CancelCardinfo) | **Post** /paynow/v2/Cancel/card | 決済のキャンセルを行います
[**PaynowV2CaptureCardinfo**](CardApi.md#PaynowV2CaptureCardinfo) | **Post** /paynow/v2/Capture/card | 決済の売上確定を行います



## PaynowV1AuthorizeCardinfo

> CardAuthorizeResponse PaynowV1AuthorizeCardinfo(ctx).Body(body).Execute()

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV1AuthorizeCardinfo(context.Background()).Body(body).Execute()
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
 **body** | **string** |  | 

### Return type

[**CardAuthorizeResponse**](CardAuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowV1ReAuthorizeCardinfo

> CardAuthorizeResponse PaynowV1ReAuthorizeCardinfo(ctx).Body(body).Execute()

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV1ReAuthorizeCardinfo(context.Background()).Body(body).Execute()
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
 **body** | **string** |  | 

### Return type

[**CardAuthorizeResponse**](CardAuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowV2CancelCardinfo

> CardCancelResponse PaynowV2CancelCardinfo(ctx).Body(body).Execute()

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2CancelCardinfo(context.Background()).Body(body).Execute()
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
 **body** | **string** |  | 

### Return type

[**CardCancelResponse**](CardCancelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaynowV2CaptureCardinfo

> CardCaptureResponse PaynowV2CaptureCardinfo(ctx).Body(body).Execute()

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2CaptureCardinfo(context.Background()).Body(body).Execute()
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
 **body** | **string** |  | 

### Return type

[**CardCaptureResponse**](CardCaptureResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

