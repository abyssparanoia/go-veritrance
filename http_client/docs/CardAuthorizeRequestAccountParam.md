# CardAuthorizeRequestAccountParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CardParam** | [**CardAuthorizeRequestCardParam**](CardAuthorizeRequestCardParam.md) |  | 

## Methods

### NewCardAuthorizeRequestAccountParam

`func NewCardAuthorizeRequestAccountParam(accountId string, cardParam CardAuthorizeRequestCardParam, ) *CardAuthorizeRequestAccountParam`

NewCardAuthorizeRequestAccountParam instantiates a new CardAuthorizeRequestAccountParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuthorizeRequestAccountParamWithDefaults

`func NewCardAuthorizeRequestAccountParamWithDefaults() *CardAuthorizeRequestAccountParam`

NewCardAuthorizeRequestAccountParamWithDefaults instantiates a new CardAuthorizeRequestAccountParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CardAuthorizeRequestAccountParam) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CardAuthorizeRequestAccountParam) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CardAuthorizeRequestAccountParam) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCardParam

`func (o *CardAuthorizeRequestAccountParam) GetCardParam() CardAuthorizeRequestCardParam`

GetCardParam returns the CardParam field if non-nil, zero value otherwise.

### GetCardParamOk

`func (o *CardAuthorizeRequestAccountParam) GetCardParamOk() (*CardAuthorizeRequestCardParam, bool)`

GetCardParamOk returns a tuple with the CardParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardParam

`func (o *CardAuthorizeRequestAccountParam) SetCardParam(v CardAuthorizeRequestCardParam)`

SetCardParam sets CardParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


