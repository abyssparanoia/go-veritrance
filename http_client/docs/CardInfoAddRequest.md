# CardInfoAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxnVersion** | **string** |  | 
**DummyRequest** | **string** |  | 
**MerchantCcid** | **string** |  | 
**AccountId** | Pointer to **string** |  | [optional] 
**Token** | **string** |  | 
**CardNumber** | Pointer to **string** |  | [optional] 
**CardExpire** | Pointer to **string** |  | [optional] 
**SecurityCode** | Pointer to **string** |  | [optional] 

## Methods

### NewCardInfoAddRequest

`func NewCardInfoAddRequest(txnVersion string, dummyRequest string, merchantCcid string, token string, ) *CardInfoAddRequest`

NewCardInfoAddRequest instantiates a new CardInfoAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoAddRequestWithDefaults

`func NewCardInfoAddRequestWithDefaults() *CardInfoAddRequest`

NewCardInfoAddRequestWithDefaults instantiates a new CardInfoAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxnVersion

`func (o *CardInfoAddRequest) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardInfoAddRequest) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardInfoAddRequest) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetDummyRequest

`func (o *CardInfoAddRequest) GetDummyRequest() string`

GetDummyRequest returns the DummyRequest field if non-nil, zero value otherwise.

### GetDummyRequestOk

`func (o *CardInfoAddRequest) GetDummyRequestOk() (*string, bool)`

GetDummyRequestOk returns a tuple with the DummyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyRequest

`func (o *CardInfoAddRequest) SetDummyRequest(v string)`

SetDummyRequest sets DummyRequest field to given value.


### GetMerchantCcid

`func (o *CardInfoAddRequest) GetMerchantCcid() string`

GetMerchantCcid returns the MerchantCcid field if non-nil, zero value otherwise.

### GetMerchantCcidOk

`func (o *CardInfoAddRequest) GetMerchantCcidOk() (*string, bool)`

GetMerchantCcidOk returns a tuple with the MerchantCcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCcid

`func (o *CardInfoAddRequest) SetMerchantCcid(v string)`

SetMerchantCcid sets MerchantCcid field to given value.


### GetAccountId

`func (o *CardInfoAddRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CardInfoAddRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CardInfoAddRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CardInfoAddRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetToken

`func (o *CardInfoAddRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardInfoAddRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardInfoAddRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetCardNumber

`func (o *CardInfoAddRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardInfoAddRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardInfoAddRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CardInfoAddRequest) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpire

`func (o *CardInfoAddRequest) GetCardExpire() string`

GetCardExpire returns the CardExpire field if non-nil, zero value otherwise.

### GetCardExpireOk

`func (o *CardInfoAddRequest) GetCardExpireOk() (*string, bool)`

GetCardExpireOk returns a tuple with the CardExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpire

`func (o *CardInfoAddRequest) SetCardExpire(v string)`

SetCardExpire sets CardExpire field to given value.

### HasCardExpire

`func (o *CardInfoAddRequest) HasCardExpire() bool`

HasCardExpire returns a boolean if a field has been set.

### GetSecurityCode

`func (o *CardInfoAddRequest) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *CardInfoAddRequest) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *CardInfoAddRequest) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *CardInfoAddRequest) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


