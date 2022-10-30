# CardCancelRequestParams

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

### NewCardCancelRequestParams

`func NewCardCancelRequestParams(txnVersion string, dummyRequest string, merchantCcid string, orderId string, amount string, currencyUnit CurrencyUnit, ) *CardCancelRequestParams`

NewCardCancelRequestParams instantiates a new CardCancelRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCancelRequestParamsWithDefaults

`func NewCardCancelRequestParamsWithDefaults() *CardCancelRequestParams`

NewCardCancelRequestParamsWithDefaults instantiates a new CardCancelRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardCancelRequestParams) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardCancelRequestParams) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardCancelRequestParams) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardCancelRequestParams) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardCancelRequestParams) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardCancelRequestParams) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardCancelRequestParams) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardCancelRequestParams) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardCancelRequestParams) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetOrderId

`func (o *CardCancelRequestParams) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardCancelRequestParams) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardCancelRequestParams) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetAmount

`func (o *CardCancelRequestParams) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardCancelRequestParams) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardCancelRequestParams) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrencyUnit

`func (o *CardCancelRequestParams) GetCurrencyUnit() CurrencyUnit`

GetCurrencyUnit returns the CurrencyUnit field if non-nil, zero value otherwise.

### GetCurrencyUnitOk

`func (o *CardCancelRequestParams) GetCurrencyUnitOk() (*CurrencyUnit, bool)`

GetCurrencyUnitOk returns a tuple with the CurrencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyUnit

`func (o *CardCancelRequestParams) SetCurrencyUnit(v CurrencyUnit)`

SetCurrencyUnit sets CurrencyUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


