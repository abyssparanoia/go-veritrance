# CardCaptureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | [**CardCaptureRequestParams**](CardCaptureRequestParams.md) |  | 
**AuthHash** | **string** |  | 

## Methods

### NewCardCaptureRequest

`func NewCardCaptureRequest(params CardCaptureRequestParams, authHash string, ) *CardCaptureRequest`

NewCardCaptureRequest instantiates a new CardCaptureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCaptureRequestWithDefaults

`func NewCardCaptureRequestWithDefaults() *CardCaptureRequest`

NewCardCaptureRequestWithDefaults instantiates a new CardCaptureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *CardCaptureRequest) GetParams() CardCaptureRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CardCaptureRequest) GetParamsOk() (*CardCaptureRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CardCaptureRequest) SetParams(v CardCaptureRequestParams)`

SetParams sets Params field to given value.


### GetAuthHash

`func (o *CardCaptureRequest) GetAuthHash() string`

GetAuthHash returns the AuthHash field if non-nil, zero value otherwise.

### GetAuthHashOk

`func (o *CardCaptureRequest) GetAuthHashOk() (*string, bool)`

GetAuthHashOk returns a tuple with the AuthHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHash

`func (o *CardCaptureRequest) SetAuthHash(v string)`

SetAuthHash sets AuthHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


