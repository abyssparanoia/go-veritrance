# CardInfoGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | [**CardInfoGetRequestParams**](CardInfoGetRequestParams.md) |  | 
**AuthHash** | **string** |  | 

## Methods

### NewCardInfoGetRequest

`func NewCardInfoGetRequest(params CardInfoGetRequestParams, authHash string, ) *CardInfoGetRequest`

NewCardInfoGetRequest instantiates a new CardInfoGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoGetRequestWithDefaults

`func NewCardInfoGetRequestWithDefaults() *CardInfoGetRequest`

NewCardInfoGetRequestWithDefaults instantiates a new CardInfoGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *CardInfoGetRequest) GetParams() CardInfoGetRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CardInfoGetRequest) GetParamsOk() (*CardInfoGetRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CardInfoGetRequest) SetParams(v CardInfoGetRequestParams)`

SetParams sets Params field to given value.


### GetAuthHash

`func (o *CardInfoGetRequest) GetAuthHash() string`

GetAuthHash returns the AuthHash field if non-nil, zero value otherwise.

### GetAuthHashOk

`func (o *CardInfoGetRequest) GetAuthHashOk() (*string, bool)`

GetAuthHashOk returns a tuple with the AuthHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHash

`func (o *CardInfoGetRequest) SetAuthHash(v string)`

SetAuthHash sets AuthHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


