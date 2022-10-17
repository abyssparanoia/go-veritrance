# CardAuthorizeRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**PayNowIdParam** | [**CardAuthorizeRequestParamsAllOfPayNowIdParam**](CardAuthorizeRequestParamsAllOfPayNowIdParam.md) |  | 

## Methods

### NewCardAuthorizeRequestParams

`func NewCardAuthorizeRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam CardAuthorizeRequestParamsAllOfPayNowIdParam, ) *CardAuthorizeRequestParams`

NewCardAuthorizeRequestParams instantiates a new CardAuthorizeRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuthorizeRequestParamsWithDefaults

`func NewCardAuthorizeRequestParamsWithDefaults() *CardAuthorizeRequestParams`

NewCardAuthorizeRequestParamsWithDefaults instantiates a new CardAuthorizeRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardAuthorizeRequestParams) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardAuthorizeRequestParams) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardAuthorizeRequestParams) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardAuthorizeRequestParams) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardAuthorizeRequestParams) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardAuthorizeRequestParams) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardAuthorizeRequestParams) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardAuthorizeRequestParams) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardAuthorizeRequestParams) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetPayNowIdParam

`func (o *CardAuthorizeRequestParams) GetPayNowIdParam() CardAuthorizeRequestParamsAllOfPayNowIdParam`

GetPayNowIdParam returns the PayNowIdParam field if non-nil, zero value otherwise.

### GetPayNowIdParamOk

`func (o *CardAuthorizeRequestParams) GetPayNowIdParamOk() (*CardAuthorizeRequestParamsAllOfPayNowIdParam, bool)`

GetPayNowIdParamOk returns a tuple with the PayNowIdParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdParam

`func (o *CardAuthorizeRequestParams) SetPayNowIdParam(v CardAuthorizeRequestParamsAllOfPayNowIdParam)`

SetPayNowIdParam sets PayNowIdParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


