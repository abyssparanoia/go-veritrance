# CardCancelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayNowIdResponse** | Pointer to [**AccountResponsePayNowIdResponse**](AccountResponsePayNowIdResponse.md) |  | [optional] 
**Result** | Pointer to [**CardCancelResponseResult**](CardCancelResponseResult.md) |  | [optional] 

## Methods

### NewCardCancelResponse

`func NewCardCancelResponse() *CardCancelResponse`

NewCardCancelResponse instantiates a new CardCancelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCancelResponseWithDefaults

`func NewCardCancelResponseWithDefaults() *CardCancelResponse`

NewCardCancelResponseWithDefaults instantiates a new CardCancelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayNowIdResponse

`func (o *CardCancelResponse) GetPayNowIdResponse() AccountResponsePayNowIdResponse`

GetPayNowIdResponse returns the PayNowIdResponse field if non-nil, zero value otherwise.

### GetPayNowIdResponseOk

`func (o *CardCancelResponse) GetPayNowIdResponseOk() (*AccountResponsePayNowIdResponse, bool)`

GetPayNowIdResponseOk returns a tuple with the PayNowIdResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdResponse

`func (o *CardCancelResponse) SetPayNowIdResponse(v AccountResponsePayNowIdResponse)`

SetPayNowIdResponse sets PayNowIdResponse field to given value.

### HasPayNowIdResponse

`func (o *CardCancelResponse) HasPayNowIdResponse() bool`

HasPayNowIdResponse returns a boolean if a field has been set.

### GetResult

`func (o *CardCancelResponse) GetResult() CardCancelResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CardCancelResponse) GetResultOk() (*CardCancelResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CardCancelResponse) SetResult(v CardCancelResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CardCancelResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


