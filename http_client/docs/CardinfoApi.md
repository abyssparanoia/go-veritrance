# \CardinfoApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowV2AddCardinfo**](CardinfoApi.md#PaynowV2AddCardinfo) | **Post** /paynow/v2/Add/cardinfo | 会員 ID にカード情報を紐付けて登録します
[**PaynowV2DeleteCardinfo**](CardinfoApi.md#PaynowV2DeleteCardinfo) | **Post** /paynow/v2/Delete/cardinfo | 会員 ID に紐付けられたカード情報を削除します
[**PaynowV2GetCardinfo**](CardinfoApi.md#PaynowV2GetCardinfo) | **Post** /paynow/v2/Get/cardinfo | 会員 ID に紐付けられた課金情報を取得します
[**PaynowV2UpdateCardinfo**](CardinfoApi.md#PaynowV2UpdateCardinfo) | **Post** /paynow/v2/Update/cardinfo | 当該会員 ID、およびカード ID に紐付けられたカード情報を更新します



## PaynowV2AddCardinfo

> AccountResponse PaynowV2AddCardinfo(ctx).AccountId(accountId).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).Execute()

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
    accountId := "accountId_example" // string | 
    cardNumber := "cardNumber_example" // string | 
    cardExpire := "cardExpire_example" // string | 
    securityCode := "securityCode_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardinfoApi.PaynowV2AddCardinfo(context.Background()).AccountId(accountId).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowV2AddCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2AddCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowV2AddCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2AddCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **cardNumber** | **string** |  | 
 **cardExpire** | **string** |  | 
 **securityCode** | **string** |  | 

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


## PaynowV2DeleteCardinfo

> PayNowIdResponse PaynowV2DeleteCardinfo(ctx).AccountId(accountId).CardId(cardId).Execute()

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
    accountId := "accountId_example" // string | 
    cardId := "cardId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardinfoApi.PaynowV2DeleteCardinfo(context.Background()).AccountId(accountId).CardId(cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowV2DeleteCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2DeleteCardinfo`: PayNowIdResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowV2DeleteCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2DeleteCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **cardId** | **string** |  | 

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


## PaynowV2GetCardinfo

> AccountResponse PaynowV2GetCardinfo(ctx).AccountId(accountId).Execute()

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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardinfoApi.PaynowV2GetCardinfo(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowV2GetCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2GetCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowV2GetCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2GetCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 

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


## PaynowV2UpdateCardinfo

> AccountResponse PaynowV2UpdateCardinfo(ctx).AccountId(accountId).CardId(cardId).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).Execute()

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
    accountId := "accountId_example" // string | 
    cardId := "cardId_example" // string | 
    cardNumber := "cardNumber_example" // string | 
    cardExpire := "cardExpire_example" // string | 
    securityCode := "securityCode_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardinfoApi.PaynowV2UpdateCardinfo(context.Background()).AccountId(accountId).CardId(cardId).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardinfoApi.PaynowV2UpdateCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2UpdateCardinfo`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CardinfoApi.PaynowV2UpdateCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2UpdateCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **cardId** | **string** |  | 
 **cardNumber** | **string** |  | 
 **cardExpire** | **string** |  | 
 **securityCode** | **string** |  | 

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

