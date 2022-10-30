# CardCaptureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayNowIdResponse** | Pointer to [**AccountResponsePayNowIdResponse**](AccountResponsePayNowIdResponse.md) |  | [optional] 
**Result** | Pointer to [**CardCaptureResponseResult**](CardCaptureResponseResult.md) |  | [optional] 

## Methods

### NewCardCaptureResponse

`func NewCardCaptureResponse() *CardCaptureResponse`

NewCardCaptureResponse instantiates a new CardCaptureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCaptureResponseWithDefaults

`func NewCardCaptureResponseWithDefaults() *CardCaptureResponse`

NewCardCaptureResponseWithDefaults instantiates a new CardCaptureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayNowIdResponse

`func (o *CardCaptureResponse) GetPayNowIdResponse() AccountResponsePayNowIdResponse`

GetPayNowIdResponse returns the PayNowIdResponse field if non-nil, zero value otherwise.

### GetPayNowIdResponseOk

`func (o *CardCaptureResponse) GetPayNowIdResponseOk() (*AccountResponsePayNowIdResponse, bool)`

GetPayNowIdResponseOk returns a tuple with the PayNowIdResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdResponse

`func (o *CardCaptureResponse) SetPayNowIdResponse(v AccountResponsePayNowIdResponse)`

SetPayNowIdResponse sets PayNowIdResponse field to given value.

### HasPayNowIdResponse

`func (o *CardCaptureResponse) HasPayNowIdResponse() bool`

HasPayNowIdResponse returns a boolean if a field has been set.

### GetResult

`func (o *CardCaptureResponse) GetResult() CardCaptureResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CardCaptureResponse) GetResultOk() (*CardCaptureResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CardCaptureResponse) SetResult(v CardCaptureResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CardCaptureResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


