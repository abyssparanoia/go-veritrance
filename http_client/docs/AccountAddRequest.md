# AccountAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**AccountId** | **string** |  | 
**CardNumber** | Pointer to **string** |  | [optional] 
**CardExpire** | Pointer to **string** |  | [optional] 
**SecurityCode** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountAddRequest

`func NewAccountAddRequest(txnVersion string, dummyRequest string, merchantCcid string, accountId string, ) *AccountAddRequest`

NewAccountAddRequest instantiates a new AccountAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAddRequestWithDefaults

`func NewAccountAddRequestWithDefaults() *AccountAddRequest`

NewAccountAddRequestWithDefaults instantiates a new AccountAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *AccountAddRequest) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *AccountAddRequest) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *AccountAddRequest) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *AccountAddRequest) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *AccountAddRequest) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *AccountAddRequest) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *AccountAddRequest) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *AccountAddRequest) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *AccountAddRequest) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetAccountId

`func (o *AccountAddRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountAddRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountAddRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCardNumber

`func (o *AccountAddRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *AccountAddRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *AccountAddRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *AccountAddRequest) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpire

`func (o *AccountAddRequest) GetCardExpire() string`

GetCardExpire returns the CardExpire field if non-nil, zero value otherwise.

### GetCardExpireOk

`func (o *AccountAddRequest) GetCardExpireOk() (*string, bool)`

GetCardExpireOk returns a tuple with the CardExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpire

`func (o *AccountAddRequest) SetCardExpire(v string)`

SetCardExpire sets CardExpire field to given value.

### HasCardExpire

`func (o *AccountAddRequest) HasCardExpire() bool`

HasCardExpire returns a boolean if a field has been set.

### GetSecurityCode

`func (o *AccountAddRequest) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *AccountAddRequest) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *AccountAddRequest) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *AccountAddRequest) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

