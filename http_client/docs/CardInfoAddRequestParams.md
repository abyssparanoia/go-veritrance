# CardInfoAddRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**PayNowIdParam** | [**CardInfoAddRequestParamsAllOfPayNowIdParam**](CardInfoAddRequestParamsAllOfPayNowIdParam.md) |  | 

## Methods

### NewCardInfoAddRequestParams

`func NewCardInfoAddRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam CardInfoAddRequestParamsAllOfPayNowIdParam, ) *CardInfoAddRequestParams`

NewCardInfoAddRequestParams instantiates a new CardInfoAddRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoAddRequestParamsWithDefaults

`func NewCardInfoAddRequestParamsWithDefaults() *CardInfoAddRequestParams`

NewCardInfoAddRequestParamsWithDefaults instantiates a new CardInfoAddRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardInfoAddRequestParams) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardInfoAddRequestParams) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardInfoAddRequestParams) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardInfoAddRequestParams) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardInfoAddRequestParams) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardInfoAddRequestParams) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardInfoAddRequestParams) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardInfoAddRequestParams) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardInfoAddRequestParams) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetPayNowIdParam

`func (o *CardInfoAddRequestParams) GetPayNowIdParam() CardInfoAddRequestParamsAllOfPayNowIdParam`

GetPayNowIdParam returns the PayNowIdParam field if non-nil, zero value otherwise.

### GetPayNowIdParamOk

`func (o *CardInfoAddRequestParams) GetPayNowIdParamOk() (*CardInfoAddRequestParamsAllOfPayNowIdParam, bool)`

GetPayNowIdParamOk returns a tuple with the PayNowIdParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdParam

`func (o *CardInfoAddRequestParams) SetPayNowIdParam(v CardInfoAddRequestParamsAllOfPayNowIdParam)`

SetPayNowIdParam sets PayNowIdParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

