# CardInfoUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**PayNowIdParam** | [**CardInfoUpdateRequestParamsAllOfPayNowIdParam**](CardInfoUpdateRequestParamsAllOfPayNowIdParam.md) |  | 

## Methods

### NewCardInfoUpdateRequestParams

`func NewCardInfoUpdateRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam CardInfoUpdateRequestParamsAllOfPayNowIdParam, ) *CardInfoUpdateRequestParams`

NewCardInfoUpdateRequestParams instantiates a new CardInfoUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoUpdateRequestParamsWithDefaults

`func NewCardInfoUpdateRequestParamsWithDefaults() *CardInfoUpdateRequestParams`

NewCardInfoUpdateRequestParamsWithDefaults instantiates a new CardInfoUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardInfoUpdateRequestParams) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardInfoUpdateRequestParams) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardInfoUpdateRequestParams) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardInfoUpdateRequestParams) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardInfoUpdateRequestParams) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardInfoUpdateRequestParams) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardInfoUpdateRequestParams) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardInfoUpdateRequestParams) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardInfoUpdateRequestParams) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetPayNowIdParam

`func (o *CardInfoUpdateRequestParams) GetPayNowIdParam() CardInfoUpdateRequestParamsAllOfPayNowIdParam`

GetPayNowIdParam returns the PayNowIdParam field if non-nil, zero value otherwise.

### GetPayNowIdParamOk

`func (o *CardInfoUpdateRequestParams) GetPayNowIdParamOk() (*CardInfoUpdateRequestParamsAllOfPayNowIdParam, bool)`

GetPayNowIdParamOk returns a tuple with the PayNowIdParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdParam

`func (o *CardInfoUpdateRequestParams) SetPayNowIdParam(v CardInfoUpdateRequestParamsAllOfPayNowIdParam)`

SetPayNowIdParam sets PayNowIdParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


