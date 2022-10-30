# CardCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | [**CardCancelRequestParams**](CardCancelRequestParams.md) |  | 
**AuthHash** | **string** |  | 

## Methods

### NewCardCancelRequest

`func NewCardCancelRequest(params CardCancelRequestParams, authHash string, ) *CardCancelRequest`

NewCardCancelRequest instantiates a new CardCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCancelRequestWithDefaults

`func NewCardCancelRequestWithDefaults() *CardCancelRequest`

NewCardCancelRequestWithDefaults instantiates a new CardCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *CardCancelRequest) GetParams() CardCancelRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CardCancelRequest) GetParamsOk() (*CardCancelRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CardCancelRequest) SetParams(v CardCancelRequestParams)`

SetParams sets Params field to given value.


### GetAuthHash

`func (o *CardCancelRequest) GetAuthHash() string`

GetAuthHash returns the AuthHash field if non-nil, zero value otherwise.

### GetAuthHashOk

`func (o *CardCancelRequest) GetAuthHashOk() (*string, bool)`

GetAuthHashOk returns a tuple with the AuthHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHash

`func (o *CardCancelRequest) SetAuthHash(v string)`

SetAuthHash sets AuthHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


