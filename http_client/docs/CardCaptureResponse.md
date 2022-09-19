# CardCaptureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceType** | Pointer to **string** | 要求電文を送信した決済サービスタイプ | [optional] 
**Status** | Pointer to **string** | 処理結果コード | [optional] 
**VResultCode** | Pointer to **string** | 処理の結果を詳細に表すコード 4 桁ずつ 4 つのブロックで構成され、各ブロックでサービス毎の処理結果を表します。  | [optional] 
**MerrMsg** | Pointer to **string** | 処理結果を日本語で表示します。 | [optional] 
**MarchTxn** | Pointer to **string** | 決済サーバーにて決済処理電文（内部処理も含む）毎に付与する ID １つの取引 ID に対して、複数の ID が付与されます。  | [optional] 
**OrderId** | Pointer to **string** | 決済要求時に店舗様にて任意に採番し送信された取引 ID | [optional] 
**CustTxn** | Pointer to **string** | 決済サーバーがオーダー（取引 ID）と紐付ける為に採番する ID | [optional] 
**TxnVersion** | Pointer to **string** | 電文のバージョン | [optional] 
**CardTransactiontype** | Pointer to **string** |  | [optional] 
**GatewayRequestDate** | Pointer to **time.Time** |  | [optional] 
**GatewayResponseDate** | Pointer to **time.Time** |  | [optional] 
**CenterRequestDate** | Pointer to **time.Time** |  | [optional] 
**CenterResponseDate** | Pointer to **time.Time** |  | [optional] 
**Pending** | Pointer to **int32** |  | [optional] 
**Loopback** | Pointer to **int32** |  | [optional] 
**ConnectedCenterId** | Pointer to **string** |  | [optional] 
**CenterRequestNumber** | Pointer to **string** |  | [optional] 
**CenterReferenceNumber** | Pointer to **string** |  | [optional] 
**ResReturnReferenceNumber** | Pointer to **string** |  | [optional] 
**ResAuthCode** | Pointer to **string** |  | [optional] 
**ResActionCode** | Pointer to **string** |  | [optional] 
**ResCenterErrorCode** | Pointer to **string** |  | [optional] 
**AcquirerCode** | Pointer to **string** |  | [optional] 

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

### GetServiceType

`func (o *CardCaptureResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CardCaptureResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CardCaptureResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *CardCaptureResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *CardCaptureResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardCaptureResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardCaptureResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardCaptureResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVResultCode

`func (o *CardCaptureResponse) GetVResultCode() string`

GetVResultCode returns the VResultCode field if non-nil, zero value otherwise.

### GetVResultCodeOk

`func (o *CardCaptureResponse) GetVResultCodeOk() (*string, bool)`

GetVResultCodeOk returns a tuple with the VResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVResultCode

`func (o *CardCaptureResponse) SetVResultCode(v string)`

SetVResultCode sets VResultCode field to given value.

### HasVResultCode

`func (o *CardCaptureResponse) HasVResultCode() bool`

HasVResultCode returns a boolean if a field has been set.

### GetMerrMsg

`func (o *CardCaptureResponse) GetMerrMsg() string`

GetMerrMsg returns the MerrMsg field if non-nil, zero value otherwise.

### GetMerrMsgOk

`func (o *CardCaptureResponse) GetMerrMsgOk() (*string, bool)`

GetMerrMsgOk returns a tuple with the MerrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerrMsg

`func (o *CardCaptureResponse) SetMerrMsg(v string)`

SetMerrMsg sets MerrMsg field to given value.

### HasMerrMsg

`func (o *CardCaptureResponse) HasMerrMsg() bool`

HasMerrMsg returns a boolean if a field has been set.

### GetMarchTxn

`func (o *CardCaptureResponse) GetMarchTxn() string`

GetMarchTxn returns the MarchTxn field if non-nil, zero value otherwise.

### GetMarchTxnOk

`func (o *CardCaptureResponse) GetMarchTxnOk() (*string, bool)`

GetMarchTxnOk returns a tuple with the MarchTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarchTxn

`func (o *CardCaptureResponse) SetMarchTxn(v string)`

SetMarchTxn sets MarchTxn field to given value.

### HasMarchTxn

`func (o *CardCaptureResponse) HasMarchTxn() bool`

HasMarchTxn returns a boolean if a field has been set.

### GetOrderId

`func (o *CardCaptureResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardCaptureResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardCaptureResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CardCaptureResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCustTxn

`func (o *CardCaptureResponse) GetCustTxn() string`

GetCustTxn returns the CustTxn field if non-nil, zero value otherwise.

### GetCustTxnOk

`func (o *CardCaptureResponse) GetCustTxnOk() (*string, bool)`

GetCustTxnOk returns a tuple with the CustTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustTxn

`func (o *CardCaptureResponse) SetCustTxn(v string)`

SetCustTxn sets CustTxn field to given value.

### HasCustTxn

`func (o *CardCaptureResponse) HasCustTxn() bool`

HasCustTxn returns a boolean if a field has been set.

### GetTxnVersion

`func (o *CardCaptureResponse) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardCaptureResponse) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardCaptureResponse) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.

### HasTxnVersion

`func (o *CardCaptureResponse) HasTxnVersion() bool`

HasTxnVersion returns a boolean if a field has been set.

### GetCardTransactiontype

`func (o *CardCaptureResponse) GetCardTransactiontype() string`

GetCardTransactiontype returns the CardTransactiontype field if non-nil, zero value otherwise.

### GetCardTransactiontypeOk

`func (o *CardCaptureResponse) GetCardTransactiontypeOk() (*string, bool)`

GetCardTransactiontypeOk returns a tuple with the CardTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTransactiontype

`func (o *CardCaptureResponse) SetCardTransactiontype(v string)`

SetCardTransactiontype sets CardTransactiontype field to given value.

### HasCardTransactiontype

`func (o *CardCaptureResponse) HasCardTransactiontype() bool`

HasCardTransactiontype returns a boolean if a field has been set.

### GetGatewayRequestDate

`func (o *CardCaptureResponse) GetGatewayRequestDate() time.Time`

GetGatewayRequestDate returns the GatewayRequestDate field if non-nil, zero value otherwise.

### GetGatewayRequestDateOk

`func (o *CardCaptureResponse) GetGatewayRequestDateOk() (*time.Time, bool)`

GetGatewayRequestDateOk returns a tuple with the GatewayRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRequestDate

`func (o *CardCaptureResponse) SetGatewayRequestDate(v time.Time)`

SetGatewayRequestDate sets GatewayRequestDate field to given value.

### HasGatewayRequestDate

`func (o *CardCaptureResponse) HasGatewayRequestDate() bool`

HasGatewayRequestDate returns a boolean if a field has been set.

### GetGatewayResponseDate

`func (o *CardCaptureResponse) GetGatewayResponseDate() time.Time`

GetGatewayResponseDate returns the GatewayResponseDate field if non-nil, zero value otherwise.

### GetGatewayResponseDateOk

`func (o *CardCaptureResponse) GetGatewayResponseDateOk() (*time.Time, bool)`

GetGatewayResponseDateOk returns a tuple with the GatewayResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseDate

`func (o *CardCaptureResponse) SetGatewayResponseDate(v time.Time)`

SetGatewayResponseDate sets GatewayResponseDate field to given value.

### HasGatewayResponseDate

`func (o *CardCaptureResponse) HasGatewayResponseDate() bool`

HasGatewayResponseDate returns a boolean if a field has been set.

### GetCenterRequestDate

`func (o *CardCaptureResponse) GetCenterRequestDate() time.Time`

GetCenterRequestDate returns the CenterRequestDate field if non-nil, zero value otherwise.

### GetCenterRequestDateOk

`func (o *CardCaptureResponse) GetCenterRequestDateOk() (*time.Time, bool)`

GetCenterRequestDateOk returns a tuple with the CenterRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestDate

`func (o *CardCaptureResponse) SetCenterRequestDate(v time.Time)`

SetCenterRequestDate sets CenterRequestDate field to given value.

### HasCenterRequestDate

`func (o *CardCaptureResponse) HasCenterRequestDate() bool`

HasCenterRequestDate returns a boolean if a field has been set.

### GetCenterResponseDate

`func (o *CardCaptureResponse) GetCenterResponseDate() time.Time`

GetCenterResponseDate returns the CenterResponseDate field if non-nil, zero value otherwise.

### GetCenterResponseDateOk

`func (o *CardCaptureResponse) GetCenterResponseDateOk() (*time.Time, bool)`

GetCenterResponseDateOk returns a tuple with the CenterResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterResponseDate

`func (o *CardCaptureResponse) SetCenterResponseDate(v time.Time)`

SetCenterResponseDate sets CenterResponseDate field to given value.

### HasCenterResponseDate

`func (o *CardCaptureResponse) HasCenterResponseDate() bool`

HasCenterResponseDate returns a boolean if a field has been set.

### GetPending

`func (o *CardCaptureResponse) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CardCaptureResponse) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CardCaptureResponse) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *CardCaptureResponse) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetLoopback

`func (o *CardCaptureResponse) GetLoopback() int32`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *CardCaptureResponse) GetLoopbackOk() (*int32, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *CardCaptureResponse) SetLoopback(v int32)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *CardCaptureResponse) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetConnectedCenterId

`func (o *CardCaptureResponse) GetConnectedCenterId() string`

GetConnectedCenterId returns the ConnectedCenterId field if non-nil, zero value otherwise.

### GetConnectedCenterIdOk

`func (o *CardCaptureResponse) GetConnectedCenterIdOk() (*string, bool)`

GetConnectedCenterIdOk returns a tuple with the ConnectedCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedCenterId

`func (o *CardCaptureResponse) SetConnectedCenterId(v string)`

SetConnectedCenterId sets ConnectedCenterId field to given value.

### HasConnectedCenterId

`func (o *CardCaptureResponse) HasConnectedCenterId() bool`

HasConnectedCenterId returns a boolean if a field has been set.

### GetCenterRequestNumber

`func (o *CardCaptureResponse) GetCenterRequestNumber() string`

GetCenterRequestNumber returns the CenterRequestNumber field if non-nil, zero value otherwise.

### GetCenterRequestNumberOk

`func (o *CardCaptureResponse) GetCenterRequestNumberOk() (*string, bool)`

GetCenterRequestNumberOk returns a tuple with the CenterRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestNumber

`func (o *CardCaptureResponse) SetCenterRequestNumber(v string)`

SetCenterRequestNumber sets CenterRequestNumber field to given value.

### HasCenterRequestNumber

`func (o *CardCaptureResponse) HasCenterRequestNumber() bool`

HasCenterRequestNumber returns a boolean if a field has been set.

### GetCenterReferenceNumber

`func (o *CardCaptureResponse) GetCenterReferenceNumber() string`

GetCenterReferenceNumber returns the CenterReferenceNumber field if non-nil, zero value otherwise.

### GetCenterReferenceNumberOk

`func (o *CardCaptureResponse) GetCenterReferenceNumberOk() (*string, bool)`

GetCenterReferenceNumberOk returns a tuple with the CenterReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterReferenceNumber

`func (o *CardCaptureResponse) SetCenterReferenceNumber(v string)`

SetCenterReferenceNumber sets CenterReferenceNumber field to given value.

### HasCenterReferenceNumber

`func (o *CardCaptureResponse) HasCenterReferenceNumber() bool`

HasCenterReferenceNumber returns a boolean if a field has been set.

### GetResReturnReferenceNumber

`func (o *CardCaptureResponse) GetResReturnReferenceNumber() string`

GetResReturnReferenceNumber returns the ResReturnReferenceNumber field if non-nil, zero value otherwise.

### GetResReturnReferenceNumberOk

`func (o *CardCaptureResponse) GetResReturnReferenceNumberOk() (*string, bool)`

GetResReturnReferenceNumberOk returns a tuple with the ResReturnReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResReturnReferenceNumber

`func (o *CardCaptureResponse) SetResReturnReferenceNumber(v string)`

SetResReturnReferenceNumber sets ResReturnReferenceNumber field to given value.

### HasResReturnReferenceNumber

`func (o *CardCaptureResponse) HasResReturnReferenceNumber() bool`

HasResReturnReferenceNumber returns a boolean if a field has been set.

### GetResAuthCode

`func (o *CardCaptureResponse) GetResAuthCode() string`

GetResAuthCode returns the ResAuthCode field if non-nil, zero value otherwise.

### GetResAuthCodeOk

`func (o *CardCaptureResponse) GetResAuthCodeOk() (*string, bool)`

GetResAuthCodeOk returns a tuple with the ResAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResAuthCode

`func (o *CardCaptureResponse) SetResAuthCode(v string)`

SetResAuthCode sets ResAuthCode field to given value.

### HasResAuthCode

`func (o *CardCaptureResponse) HasResAuthCode() bool`

HasResAuthCode returns a boolean if a field has been set.

### GetResActionCode

`func (o *CardCaptureResponse) GetResActionCode() string`

GetResActionCode returns the ResActionCode field if non-nil, zero value otherwise.

### GetResActionCodeOk

`func (o *CardCaptureResponse) GetResActionCodeOk() (*string, bool)`

GetResActionCodeOk returns a tuple with the ResActionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResActionCode

`func (o *CardCaptureResponse) SetResActionCode(v string)`

SetResActionCode sets ResActionCode field to given value.

### HasResActionCode

`func (o *CardCaptureResponse) HasResActionCode() bool`

HasResActionCode returns a boolean if a field has been set.

### GetResCenterErrorCode

`func (o *CardCaptureResponse) GetResCenterErrorCode() string`

GetResCenterErrorCode returns the ResCenterErrorCode field if non-nil, zero value otherwise.

### GetResCenterErrorCodeOk

`func (o *CardCaptureResponse) GetResCenterErrorCodeOk() (*string, bool)`

GetResCenterErrorCodeOk returns a tuple with the ResCenterErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResCenterErrorCode

`func (o *CardCaptureResponse) SetResCenterErrorCode(v string)`

SetResCenterErrorCode sets ResCenterErrorCode field to given value.

### HasResCenterErrorCode

`func (o *CardCaptureResponse) HasResCenterErrorCode() bool`

HasResCenterErrorCode returns a boolean if a field has been set.

### GetAcquirerCode

`func (o *CardCaptureResponse) GetAcquirerCode() string`

GetAcquirerCode returns the AcquirerCode field if non-nil, zero value otherwise.

### GetAcquirerCodeOk

`func (o *CardCaptureResponse) GetAcquirerCodeOk() (*string, bool)`

GetAcquirerCodeOk returns a tuple with the AcquirerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerCode

`func (o *CardCaptureResponse) SetAcquirerCode(v string)`

SetAcquirerCode sets AcquirerCode field to given value.

### HasAcquirerCode

`func (o *CardCaptureResponse) HasAcquirerCode() bool`

HasAcquirerCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


