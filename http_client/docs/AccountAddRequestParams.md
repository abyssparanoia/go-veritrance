# AccountAddRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**PayNowIdParam** | [**AccountAddRequestParamsAllOfPayNowIdParam**](AccountAddRequestParamsAllOfPayNowIdParam.md) |  | 

## Methods

### NewAccountAddRequestParams

`func NewAccountAddRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam AccountAddRequestParamsAllOfPayNowIdParam, ) *AccountAddRequestParams`

NewAccountAddRequestParams instantiates a new AccountAddRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAddRequestParamsWithDefaults

`func NewAccountAddRequestParamsWithDefaults() *AccountAddRequestParams`

NewAccountAddRequestParamsWithDefaults instantiates a new AccountAddRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *AccountAddRequestParams) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *AccountAddRequestParams) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *AccountAddRequestParams) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *AccountAddRequestParams) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *AccountAddRequestParams) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *AccountAddRequestParams) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *AccountAddRequestParams) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *AccountAddRequestParams) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *AccountAddRequestParams) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetPayNowIdParam

`func (o *AccountAddRequestParams) GetPayNowIdParam() AccountAddRequestParamsAllOfPayNowIdParam`

GetPayNowIdParam returns the PayNowIdParam field if non-nil, zero value otherwise.

### GetPayNowIdParamOk

`func (o *AccountAddRequestParams) GetPayNowIdParamOk() (*AccountAddRequestParamsAllOfPayNowIdParam, bool)`

GetPayNowIdParamOk returns a tuple with the PayNowIdParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdParam

`func (o *AccountAddRequestParams) SetPayNowIdParam(v AccountAddRequestParamsAllOfPayNowIdParam)`

SetPayNowIdParam sets PayNowIdParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


