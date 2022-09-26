# AccountV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CardInfo** | Pointer to [**[]CardInfoV1**](CardInfoV1.md) |  | [optional] 

## Methods

### NewAccountV1

`func NewAccountV1(accountId string, ) *AccountV1`

NewAccountV1 instantiates a new AccountV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountV1WithDefaults

`func NewAccountV1WithDefaults() *AccountV1`

NewAccountV1WithDefaults instantiates a new AccountV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountV1) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountV1) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountV1) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCardInfo

`func (o *AccountV1) GetCardInfo() []CardInfoV1`

GetCardInfo returns the CardInfo field if non-nil, zero value otherwise.

### GetCardInfoOk

`func (o *AccountV1) GetCardInfoOk() (*[]CardInfoV1, bool)`

GetCardInfoOk returns a tuple with the CardInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInfo

`func (o *AccountV1) SetCardInfo(v []CardInfoV1)`

SetCardInfo sets CardInfo field to given value.

### HasCardInfo

`func (o *AccountV1) HasCardInfo() bool`

HasCardInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


