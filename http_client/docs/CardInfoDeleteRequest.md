# CardInfoDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**AccountId** | **string** |  | 
**CardId** | Pointer to **string** |  | [optional] 

## Methods

### NewCardInfoDeleteRequest

`func NewCardInfoDeleteRequest(txnVersion string, dummyRequest string, merchantCcid string, accountId string, ) *CardInfoDeleteRequest`

NewCardInfoDeleteRequest instantiates a new CardInfoDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoDeleteRequestWithDefaults

`func NewCardInfoDeleteRequestWithDefaults() *CardInfoDeleteRequest`

NewCardInfoDeleteRequestWithDefaults instantiates a new CardInfoDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardInfoDeleteRequest) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardInfoDeleteRequest) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardInfoDeleteRequest) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardInfoDeleteRequest) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardInfoDeleteRequest) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardInfoDeleteRequest) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardInfoDeleteRequest) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardInfoDeleteRequest) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardInfoDeleteRequest) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetAccountId

`func (o *CardInfoDeleteRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CardInfoDeleteRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CardInfoDeleteRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCardId

`func (o *CardInfoDeleteRequest) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *CardInfoDeleteRequest) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *CardInfoDeleteRequest) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *CardInfoDeleteRequest) HasCardId() bool`

HasCardId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


