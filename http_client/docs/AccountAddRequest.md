# AccountAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | [**AccountAddRequestParams**](AccountAddRequestParams.md) |  | 
**AuthHash** | **string** |  | 

## Methods

### NewAccountAddRequest

`func NewAccountAddRequest(params AccountAddRequestParams, authHash string, ) *AccountAddRequest`

NewAccountAddRequest instantiates a new AccountAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAddRequestWithDefaults

`func NewAccountAddRequestWithDefaults() *AccountAddRequest`

NewAccountAddRequestWithDefaults instantiates a new AccountAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *AccountAddRequest) GetParams() AccountAddRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AccountAddRequest) GetParamsOk() (*AccountAddRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AccountAddRequest) SetParams(v AccountAddRequestParams)`

SetParams sets Params field to given value.


### GetAuthHash

`func (o *AccountAddRequest) GetAuthHash() string`

GetAuthHash returns the AuthHash field if non-nil, zero value otherwise.

### GetAuthHashOk

`func (o *AccountAddRequest) GetAuthHashOk() (*string, bool)`

GetAuthHashOk returns a tuple with the AuthHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHash

`func (o *AccountAddRequest) SetAuthHash(v string)`

SetAuthHash sets AuthHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


