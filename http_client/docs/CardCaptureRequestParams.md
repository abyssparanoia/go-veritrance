# CardCaptureRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**OrderId** | **string** |  | 
**Amount** | **string** |  | 
**CurrencyUnit** | [**CurrencyUnit**](CurrencyUnit.md) |  | 

## Methods

### NewCardCaptureRequestParams

`func NewCardCaptureRequestParams(txnVersion string, dummyRequest string, merchantCcid string, orderId string, amount string, currencyUnit CurrencyUnit, ) *CardCaptureRequestParams`

NewCardCaptureRequestParams instantiates a new CardCaptureRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCaptureRequestParamsWithDefaults

`func NewCardCaptureRequestParamsWithDefaults() *CardCaptureRequestParams`

NewCardCaptureRequestParamsWithDefaults instantiates a new CardCaptureRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardCaptureRequestParams) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardCaptureRequestParams) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardCaptureRequestParams) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardCaptureRequestParams) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardCaptureRequestParams) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardCaptureRequestParams) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardCaptureRequestParams) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardCaptureRequestParams) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardCaptureRequestParams) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetOrderId

`func (o *CardCaptureRequestParams) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardCaptureRequestParams) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardCaptureRequestParams) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetAmount

`func (o *CardCaptureRequestParams) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardCaptureRequestParams) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardCaptureRequestParams) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrencyUnit

`func (o *CardCaptureRequestParams) GetCurrencyUnit() CurrencyUnit`

GetCurrencyUnit returns the CurrencyUnit field if non-nil, zero value otherwise.

### GetCurrencyUnitOk

`func (o *CardCaptureRequestParams) GetCurrencyUnitOk() (*CurrencyUnit, bool)`

GetCurrencyUnitOk returns a tuple with the CurrencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyUnit

`func (o *CardCaptureRequestParams) SetCurrencyUnit(v CurrencyUnit)`

SetCurrencyUnit sets CurrencyUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


