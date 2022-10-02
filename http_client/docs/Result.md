# Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceType** | **string** | 要求電文を送信した決済サービスタイプ | 
**Status** | **string** | 処理結果コード | 
**VResultCode** | **string** | 処理の結果を詳細に表すコード 4 桁ずつ 4 つのブロックで構成され、各ブロックでサービス毎の処理結果を表します。  | 
**MerrMsg** | Pointer to **string** | 処理結果を日本語で表示します。 | [optional] 

## Methods

### NewResult

`func NewResult(serviceType string, status string, vResultCode string, ) *Result`

NewResult instantiates a new Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultWithDefaults

`func NewResultWithDefaults() *Result`

NewResultWithDefaults instantiates a new Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceType

`func (o *Result) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Result) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Result) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetStatus

`func (o *Result) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Result) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Result) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVResultCode

`func (o *Result) GetVResultCode() string`

GetVResultCode returns the VResultCode field if non-nil, zero value otherwise.

### GetVResultCodeOk

`func (o *Result) GetVResultCodeOk() (*string, bool)`

GetVResultCodeOk returns a tuple with the VResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVResultCode

`func (o *Result) SetVResultCode(v string)`

SetVResultCode sets VResultCode field to given value.


### GetMerrMsg

`func (o *Result) GetMerrMsg() string`

GetMerrMsg returns the MerrMsg field if non-nil, zero value otherwise.

### GetMerrMsgOk

`func (o *Result) GetMerrMsgOk() (*string, bool)`

GetMerrMsgOk returns a tuple with the MerrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerrMsg

`func (o *Result) SetMerrMsg(v string)`

SetMerrMsg sets MerrMsg field to given value.

### HasMerrMsg

`func (o *Result) HasMerrMsg() bool`

HasMerrMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


