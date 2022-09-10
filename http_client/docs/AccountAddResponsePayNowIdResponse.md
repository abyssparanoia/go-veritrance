# AccountAddResponsePayNowIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessId** | **string** | 会員管理処理を一意に識別する番号 | 
**Status** | **string** | 会員管理処理要求の実行結果 | 
**Message** | **string** | 会員管理処理要求の実行結果メッセージ | 
**Account** | Pointer to [**AccountV2**](AccountV2.md) |  | [optional] 

## Methods

### NewAccountAddResponsePayNowIdResponse

`func NewAccountAddResponsePayNowIdResponse(processId string, status string, message string, ) *AccountAddResponsePayNowIdResponse`

NewAccountAddResponsePayNowIdResponse instantiates a new AccountAddResponsePayNowIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAddResponsePayNowIdResponseWithDefaults

`func NewAccountAddResponsePayNowIdResponseWithDefaults() *AccountAddResponsePayNowIdResponse`

NewAccountAddResponsePayNowIdResponseWithDefaults instantiates a new AccountAddResponsePayNowIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessId

`func (o *AccountAddResponsePayNowIdResponse) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *AccountAddResponsePayNowIdResponse) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *AccountAddResponsePayNowIdResponse) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetStatus

`func (o *AccountAddResponsePayNowIdResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountAddResponsePayNowIdResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountAddResponsePayNowIdResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *AccountAddResponsePayNowIdResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccountAddResponsePayNowIdResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccountAddResponsePayNowIdResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAccount

`func (o *AccountAddResponsePayNowIdResponse) GetAccount() AccountV2`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountAddResponsePayNowIdResponse) GetAccountOk() (*AccountV2, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountAddResponsePayNowIdResponse) SetAccount(v AccountV2)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccountAddResponsePayNowIdResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


