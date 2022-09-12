# AccountV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CardInfo** | Pointer to [**[]CardInfoV2**](CardInfoV2.md) |  | [optional] 

## Methods

### NewAccountV2

`func NewAccountV2(accountId string, ) *AccountV2`

NewAccountV2 instantiates a new AccountV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountV2WithDefaults

`func NewAccountV2WithDefaults() *AccountV2`

NewAccountV2WithDefaults instantiates a new AccountV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountV2) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountV2) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountV2) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCardInfo

`func (o *AccountV2) GetCardInfo() []CardInfoV2`

GetCardInfo returns the CardInfo field if non-nil, zero value otherwise.

### GetCardInfoOk

`func (o *AccountV2) GetCardInfoOk() (*[]CardInfoV2, bool)`

GetCardInfoOk returns a tuple with the CardInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInfo

`func (o *AccountV2) SetCardInfo(v []CardInfoV2)`

SetCardInfo sets CardInfo field to given value.

### HasCardInfo

`func (o *AccountV2) HasCardInfo() bool`

HasCardInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


