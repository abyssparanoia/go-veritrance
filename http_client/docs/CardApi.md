# \CardApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowV2AuthorizeCardinfo**](CardApi.md#PaynowV2AuthorizeCardinfo) | **Post** /paynow/v2/Authorize/card | 決済の与信を行います
[**PaynowV2CancelCardinfo**](CardApi.md#PaynowV2CancelCardinfo) | **Post** /paynow/v2/Cancel/card | 決済のキャンセルを行います
[**PaynowV2CaptureCardinfo**](CardApi.md#PaynowV2CaptureCardinfo) | **Post** /paynow/v2/Capture/card | 決済の売上確定を行います
[**PaynowV2ReAuthorizeCardinfo**](CardApi.md#PaynowV2ReAuthorizeCardinfo) | **Post** /paynow/v2/ReAuthorize/card | 決済の再与信を行います



## PaynowV2AuthorizeCardinfo

> CardAuthorizeResponse PaynowV2AuthorizeCardinfo(ctx).OrderId(orderId).OriginalOrderId(originalOrderId).Amount(amount).Token(token).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).CardOptionType(cardOptionType).Jpo(jpo).WithCapture(withCapture).Execute()

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
    orderId := "orderId_example" // string | 
    originalOrderId := "originalOrderId_example" // string | 
    amount := "amount_example" // string |  (optional)
    token := "token_example" // string | トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値 (optional)
    cardNumber := "cardNumber_example" // string | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 (optional)
    cardExpire := "cardExpire_example" // string | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 (optional)
    securityCode := "securityCode_example" // string | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 (optional)
    cardOptionType := "cardOptionType_example" // string | カードオプションタイプ （MPI 有り/無し） (optional)
    jpo := "jpo_example" // string | 支払種別 \\\"10\\\"： 一括払い \\\"21\\\"： ボーナス一括 \\\"61Cxx\\\"： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\\\"10\\\"（一括払い）が適用されます。  (optional)
    withCapture := true // bool | 売上フラグ \\\"true\\\"： 与信・売上 \\\"false\\\"： 与信のみ  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2AuthorizeCardinfo(context.Background()).OrderId(orderId).OriginalOrderId(originalOrderId).Amount(amount).Token(token).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).CardOptionType(cardOptionType).Jpo(jpo).WithCapture(withCapture).Execute()
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
 **orderId** | **string** |  | 
 **originalOrderId** | **string** |  | 
 **amount** | **string** |  | 
 **token** | **string** | トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値 | 
 **cardNumber** | **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | 
 **cardExpire** | **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | 
 **securityCode** | **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | 
 **cardOptionType** | **string** | カードオプションタイプ （MPI 有り/無し） | 
 **jpo** | **string** | 支払種別 \\\&quot;10\\\&quot;： 一括払い \\\&quot;21\\\&quot;： ボーナス一括 \\\&quot;61Cxx\\\&quot;： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\\\&quot;10\\\&quot;（一括払い）が適用されます。  | 
 **withCapture** | **bool** | 売上フラグ \\\&quot;true\\\&quot;： 与信・売上 \\\&quot;false\\\&quot;： 与信のみ  | 

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

> CardCancelResponse PaynowV2CancelCardinfo(ctx).OrderId(orderId).Amount(amount).Execute()

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
    orderId := "orderId_example" // string | 
    amount := "amount_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2CancelCardinfo(context.Background()).OrderId(orderId).Amount(amount).Execute()
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
 **orderId** | **string** |  | 
 **amount** | **string** |  | 

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

> CardCaptureResponse PaynowV2CaptureCardinfo(ctx).OrderId(orderId).Amount(amount).Execute()

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
    orderId := "orderId_example" // string | 
    amount := "amount_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2CaptureCardinfo(context.Background()).OrderId(orderId).Amount(amount).Execute()
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
 **orderId** | **string** |  | 
 **amount** | **string** |  | 

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

> CardAuthorizeResponse PaynowV2ReAuthorizeCardinfo(ctx).OrderId(orderId).Amount(amount).Token(token).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).CardOptionType(cardOptionType).Jpo(jpo).WithCapture(withCapture).Execute()

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
    orderId := "orderId_example" // string | 
    amount := "amount_example" // string | 
    token := "token_example" // string | トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値 (optional)
    cardNumber := "cardNumber_example" // string | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 (optional)
    cardExpire := "cardExpire_example" // string | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 (optional)
    securityCode := "securityCode_example" // string | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 (optional)
    cardOptionType := "cardOptionType_example" // string | カードオプションタイプ （MPI 有り/無し） (optional)
    jpo := "jpo_example" // string | 支払種別 \\\"10\\\"： 一括払い \\\"21\\\"： ボーナス一括 \\\"61Cxx\\\"： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\\\"10\\\"（一括払い）が適用されます。  (optional)
    withCapture := true // bool | 売上フラグ \\\"true\\\"： 与信・売上 \\\"false\\\"： 与信のみ  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardApi.PaynowV2ReAuthorizeCardinfo(context.Background()).OrderId(orderId).Amount(amount).Token(token).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).CardOptionType(cardOptionType).Jpo(jpo).WithCapture(withCapture).Execute()
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
 **orderId** | **string** |  | 
 **amount** | **string** |  | 
 **token** | **string** | トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値 | 
 **cardNumber** | **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | 
 **cardExpire** | **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | 
 **securityCode** | **string** | （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。 | 
 **cardOptionType** | **string** | カードオプションタイプ （MPI 有り/無し） | 
 **jpo** | **string** | 支払種別 \\\&quot;10\\\&quot;： 一括払い \\\&quot;21\\\&quot;： ボーナス一括 \\\&quot;61Cxx\\\&quot;： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\\\&quot;10\\\&quot;（一括払い）が適用されます。  | 
 **withCapture** | **bool** | 売上フラグ \\\&quot;true\\\&quot;： 与信・売上 \\\&quot;false\\\&quot;： 与信のみ  | 

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

