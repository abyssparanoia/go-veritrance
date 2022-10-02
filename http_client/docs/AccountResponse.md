# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayNowIdResponse** | Pointer to [**AccountResponsePayNowIdResponse**](AccountResponsePayNowIdResponse.md) |  | [optional] 
**ServiceType** | Pointer to **string** | 要求電文を送信した決済サービスタイプ | [optional] 
**Status** | Pointer to **string** | 処理結果コード | [optional] 
**VResultCode** | Pointer to **string** | 処理の結果を詳細に表すコード 4 桁ずつ 4 つのブロックで構成され、各ブロックでサービス毎の処理結果を表します。  | [optional] 
**MerrMsg** | Pointer to **string** | 処理結果を日本語で表示します。 | [optional] 

## Methods

### NewAccountResponse

`func NewAccountResponse() *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayNowIdResponse

`func (o *AccountResponse) GetPayNowIdResponse() AccountResponsePayNowIdResponse`

GetPayNowIdResponse returns the PayNowIdResponse field if non-nil, zero value otherwise.

### GetPayNowIdResponseOk

`func (o *AccountResponse) GetPayNowIdResponseOk() (*AccountResponsePayNowIdResponse, bool)`

GetPayNowIdResponseOk returns a tuple with the PayNowIdResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdResponse

`func (o *AccountResponse) SetPayNowIdResponse(v AccountResponsePayNowIdResponse)`

SetPayNowIdResponse sets PayNowIdResponse field to given value.

### HasPayNowIdResponse

`func (o *AccountResponse) HasPayNowIdResponse() bool`

HasPayNowIdResponse returns a boolean if a field has been set.

### GetServiceType

`func (o *AccountResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AccountResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AccountResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *AccountResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *AccountResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVResultCode

`func (o *AccountResponse) GetVResultCode() string`

GetVResultCode returns the VResultCode field if non-nil, zero value otherwise.

### GetVResultCodeOk

`func (o *AccountResponse) GetVResultCodeOk() (*string, bool)`

GetVResultCodeOk returns a tuple with the VResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVResultCode

`func (o *AccountResponse) SetVResultCode(v string)`

SetVResultCode sets VResultCode field to given value.

### HasVResultCode

`func (o *AccountResponse) HasVResultCode() bool`

HasVResultCode returns a boolean if a field has been set.

### GetMerrMsg

`func (o *AccountResponse) GetMerrMsg() string`

GetMerrMsg returns the MerrMsg field if non-nil, zero value otherwise.

### GetMerrMsgOk

`func (o *AccountResponse) GetMerrMsgOk() (*string, bool)`

GetMerrMsgOk returns a tuple with the MerrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerrMsg

`func (o *AccountResponse) SetMerrMsg(v string)`

SetMerrMsg sets MerrMsg field to given value.

### HasMerrMsg

`func (o *AccountResponse) HasMerrMsg() bool`

HasMerrMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


