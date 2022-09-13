# \AccountApi

All URIs are relative to *https://api.veritrans.co.jp:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaynowV2AddAccount**](AccountApi.md#PaynowV2AddAccount) | **Post** /paynow/v2/Add/account | 会員情報を追加します。
[**PaynowV2DeleteAccount**](AccountApi.md#PaynowV2DeleteAccount) | **Post** /paynow/v2/Delete/account | 会員 ID の会員情報を、指定された「退会年月日」に削除します。
[**PaynowV2UpdateAccount**](AccountApi.md#PaynowV2UpdateAccount) | **Post** /paynow/v2/Update/account | 会員 ID の「入会年月日」を更新します。



## PaynowV2AddAccount

> AccountResponse PaynowV2AddAccount(ctx).AccountId(accountId).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).Execute()

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
    accountId := "accountId_example" // string | 
    cardNumber := "cardNumber_example" // string |  (optional)
    cardExpire := "cardExpire_example" // string |  (optional)
    securityCode := "securityCode_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.PaynowV2AddAccount(context.Background()).AccountId(accountId).CardNumber(cardNumber).CardExpire(cardExpire).SecurityCode(securityCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.PaynowV2AddAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2AddAccount`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.PaynowV2AddAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2AddAccountRequest struct via the builder pattern


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


## PaynowV2DeleteAccount

> PayNowIdResponse PaynowV2DeleteAccount(ctx).AccountId(accountId).Execute()

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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.PaynowV2DeleteAccount(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.PaynowV2DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2DeleteAccount`: PayNowIdResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.PaynowV2DeleteAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 

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


## PaynowV2UpdateAccount

> PayNowIdResponse PaynowV2UpdateAccount(ctx).AccountId(accountId).Execute()

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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.PaynowV2UpdateAccount(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.PaynowV2UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaynowV2UpdateAccount`: PayNowIdResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.PaynowV2UpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaynowV2UpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 

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

