# AbstractPayNowIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessId** | **string** | 会員管理処理を一意に識別する番号 | 
**Status** | **string** | 会員管理処理要求の実行結果 | 
**Message** | **string** | 会員管理処理要求の実行結果メッセージ | 

## Methods

### NewAbstractPayNowIdResponse

`func NewAbstractPayNowIdResponse(processId string, status string, message string, ) *AbstractPayNowIdResponse`

NewAbstractPayNowIdResponse instantiates a new AbstractPayNowIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractPayNowIdResponseWithDefaults

`func NewAbstractPayNowIdResponseWithDefaults() *AbstractPayNowIdResponse`

NewAbstractPayNowIdResponseWithDefaults instantiates a new AbstractPayNowIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessId

`func (o *AbstractPayNowIdResponse) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *AbstractPayNowIdResponse) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *AbstractPayNowIdResponse) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetStatus

`func (o *AbstractPayNowIdResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AbstractPayNowIdResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AbstractPayNowIdResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *AbstractPayNowIdResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AbstractPayNowIdResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AbstractPayNowIdResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


