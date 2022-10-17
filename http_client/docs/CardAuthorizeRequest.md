# CardAuthorizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | [**CardAuthorizeRequestParams**](CardAuthorizeRequestParams.md) |  | 
**AuthHash** | **string** |  | 

## Methods

### NewCardAuthorizeRequest

`func NewCardAuthorizeRequest(params CardAuthorizeRequestParams, authHash string, ) *CardAuthorizeRequest`

NewCardAuthorizeRequest instantiates a new CardAuthorizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuthorizeRequestWithDefaults

`func NewCardAuthorizeRequestWithDefaults() *CardAuthorizeRequest`

NewCardAuthorizeRequestWithDefaults instantiates a new CardAuthorizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *CardAuthorizeRequest) GetParams() CardAuthorizeRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CardAuthorizeRequest) GetParamsOk() (*CardAuthorizeRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CardAuthorizeRequest) SetParams(v CardAuthorizeRequestParams)`

SetParams sets Params field to given value.


### GetAuthHash

`func (o *CardAuthorizeRequest) GetAuthHash() string`

GetAuthHash returns the AuthHash field if non-nil, zero value otherwise.

### GetAuthHashOk

`func (o *CardAuthorizeRequest) GetAuthHashOk() (*string, bool)`

GetAuthHashOk returns a tuple with the AuthHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHash

`func (o *CardAuthorizeRequest) SetAuthHash(v string)`

SetAuthHash sets AuthHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


