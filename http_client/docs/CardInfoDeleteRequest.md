# CardInfoDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CardId** | Pointer to **string** |  | [optional] 

## Methods

### NewCardInfoDeleteRequest

`func NewCardInfoDeleteRequest(accountId string, ) *CardInfoDeleteRequest`

NewCardInfoDeleteRequest instantiates a new CardInfoDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoDeleteRequestWithDefaults

`func NewCardInfoDeleteRequestWithDefaults() *CardInfoDeleteRequest`

NewCardInfoDeleteRequestWithDefaults instantiates a new CardInfoDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


