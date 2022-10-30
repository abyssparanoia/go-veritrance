# CardCancelResponseResult

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
**GatewayRequestDate** | Pointer to **string** |  | [optional] 
**GatewayResponseDate** | Pointer to **string** |  | [optional] 
**CenterRequestDate** | Pointer to **string** |  | [optional] 
**CenterResponseDate** | Pointer to **string** |  | [optional] 
**Pending** | Pointer to **string** |  | [optional] 
**Loopback** | Pointer to **string** |  | [optional] 
**ConnectedCenterId** | Pointer to **string** |  | [optional] 
**CenterRequestNumber** | Pointer to **string** |  | [optional] 
**CenterReferenceNumber** | Pointer to **string** |  | [optional] 
**ResReturnReferenceNumber** | Pointer to **string** |  | [optional] 
**ResAuthCode** | Pointer to **string** |  | [optional] 
**ResActionCode** | Pointer to **string** |  | [optional] 
**ResCenterErrorCode** | Pointer to **string** |  | [optional] 
**AcquirerCode** | Pointer to **string** |  | [optional] 

## Methods

### NewCardCancelResponseResult

`func NewCardCancelResponseResult() *CardCancelResponseResult`

NewCardCancelResponseResult instantiates a new CardCancelResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCancelResponseResultWithDefaults

`func NewCardCancelResponseResultWithDefaults() *CardCancelResponseResult`

NewCardCancelResponseResultWithDefaults instantiates a new CardCancelResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceType

`func (o *CardCancelResponseResult) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CardCancelResponseResult) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CardCancelResponseResult) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *CardCancelResponseResult) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *CardCancelResponseResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardCancelResponseResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardCancelResponseResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardCancelResponseResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVResultCode

`func (o *CardCancelResponseResult) GetVResultCode() string`

GetVResultCode returns the VResultCode field if non-nil, zero value otherwise.

### GetVResultCodeOk

`func (o *CardCancelResponseResult) GetVResultCodeOk() (*string, bool)`

GetVResultCodeOk returns a tuple with the VResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVResultCode

`func (o *CardCancelResponseResult) SetVResultCode(v string)`

SetVResultCode sets VResultCode field to given value.

### HasVResultCode

`func (o *CardCancelResponseResult) HasVResultCode() bool`

HasVResultCode returns a boolean if a field has been set.

### GetMerrMsg

`func (o *CardCancelResponseResult) GetMerrMsg() string`

GetMerrMsg returns the MerrMsg field if non-nil, zero value otherwise.

### GetMerrMsgOk

`func (o *CardCancelResponseResult) GetMerrMsgOk() (*string, bool)`

GetMerrMsgOk returns a tuple with the MerrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerrMsg

`func (o *CardCancelResponseResult) SetMerrMsg(v string)`

SetMerrMsg sets MerrMsg field to given value.

### HasMerrMsg

`func (o *CardCancelResponseResult) HasMerrMsg() bool`

HasMerrMsg returns a boolean if a field has been set.

### GetMarchTxn

`func (o *CardCancelResponseResult) GetMarchTxn() string`

GetMarchTxn returns the MarchTxn field if non-nil, zero value otherwise.

### GetMarchTxnOk

`func (o *CardCancelResponseResult) GetMarchTxnOk() (*string, bool)`

GetMarchTxnOk returns a tuple with the MarchTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarchTxn

`func (o *CardCancelResponseResult) SetMarchTxn(v string)`

SetMarchTxn sets MarchTxn field to given value.

### HasMarchTxn

`func (o *CardCancelResponseResult) HasMarchTxn() bool`

HasMarchTxn returns a boolean if a field has been set.

### GetOrderId

`func (o *CardCancelResponseResult) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardCancelResponseResult) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardCancelResponseResult) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CardCancelResponseResult) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCustTxn

`func (o *CardCancelResponseResult) GetCustTxn() string`

GetCustTxn returns the CustTxn field if non-nil, zero value otherwise.

### GetCustTxnOk

`func (o *CardCancelResponseResult) GetCustTxnOk() (*string, bool)`

GetCustTxnOk returns a tuple with the CustTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustTxn

`func (o *CardCancelResponseResult) SetCustTxn(v string)`

SetCustTxn sets CustTxn field to given value.

### HasCustTxn

`func (o *CardCancelResponseResult) HasCustTxn() bool`

HasCustTxn returns a boolean if a field has been set.

### GetTxnVersion

`func (o *CardCancelResponseResult) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardCancelResponseResult) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardCancelResponseResult) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.

### HasTxnVersion

`func (o *CardCancelResponseResult) HasTxnVersion() bool`

HasTxnVersion returns a boolean if a field has been set.

### GetCardTransactiontype

`func (o *CardCancelResponseResult) GetCardTransactiontype() string`

GetCardTransactiontype returns the CardTransactiontype field if non-nil, zero value otherwise.

### GetCardTransactiontypeOk

`func (o *CardCancelResponseResult) GetCardTransactiontypeOk() (*string, bool)`

GetCardTransactiontypeOk returns a tuple with the CardTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTransactiontype

`func (o *CardCancelResponseResult) SetCardTransactiontype(v string)`

SetCardTransactiontype sets CardTransactiontype field to given value.

### HasCardTransactiontype

`func (o *CardCancelResponseResult) HasCardTransactiontype() bool`

HasCardTransactiontype returns a boolean if a field has been set.

### GetGatewayRequestDate

`func (o *CardCancelResponseResult) GetGatewayRequestDate() string`

GetGatewayRequestDate returns the GatewayRequestDate field if non-nil, zero value otherwise.

### GetGatewayRequestDateOk

`func (o *CardCancelResponseResult) GetGatewayRequestDateOk() (*string, bool)`

GetGatewayRequestDateOk returns a tuple with the GatewayRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRequestDate

`func (o *CardCancelResponseResult) SetGatewayRequestDate(v string)`

SetGatewayRequestDate sets GatewayRequestDate field to given value.

### HasGatewayRequestDate

`func (o *CardCancelResponseResult) HasGatewayRequestDate() bool`

HasGatewayRequestDate returns a boolean if a field has been set.

### GetGatewayResponseDate

`func (o *CardCancelResponseResult) GetGatewayResponseDate() string`

GetGatewayResponseDate returns the GatewayResponseDate field if non-nil, zero value otherwise.

### GetGatewayResponseDateOk

`func (o *CardCancelResponseResult) GetGatewayResponseDateOk() (*string, bool)`

GetGatewayResponseDateOk returns a tuple with the GatewayResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseDate

`func (o *CardCancelResponseResult) SetGatewayResponseDate(v string)`

SetGatewayResponseDate sets GatewayResponseDate field to given value.

### HasGatewayResponseDate

`func (o *CardCancelResponseResult) HasGatewayResponseDate() bool`

HasGatewayResponseDate returns a boolean if a field has been set.

### GetCenterRequestDate

`func (o *CardCancelResponseResult) GetCenterRequestDate() string`

GetCenterRequestDate returns the CenterRequestDate field if non-nil, zero value otherwise.

### GetCenterRequestDateOk

`func (o *CardCancelResponseResult) GetCenterRequestDateOk() (*string, bool)`

GetCenterRequestDateOk returns a tuple with the CenterRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestDate

`func (o *CardCancelResponseResult) SetCenterRequestDate(v string)`

SetCenterRequestDate sets CenterRequestDate field to given value.

### HasCenterRequestDate

`func (o *CardCancelResponseResult) HasCenterRequestDate() bool`

HasCenterRequestDate returns a boolean if a field has been set.

### GetCenterResponseDate

`func (o *CardCancelResponseResult) GetCenterResponseDate() string`

GetCenterResponseDate returns the CenterResponseDate field if non-nil, zero value otherwise.

### GetCenterResponseDateOk

`func (o *CardCancelResponseResult) GetCenterResponseDateOk() (*string, bool)`

GetCenterResponseDateOk returns a tuple with the CenterResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterResponseDate

`func (o *CardCancelResponseResult) SetCenterResponseDate(v string)`

SetCenterResponseDate sets CenterResponseDate field to given value.

### HasCenterResponseDate

`func (o *CardCancelResponseResult) HasCenterResponseDate() bool`

HasCenterResponseDate returns a boolean if a field has been set.

### GetPending

`func (o *CardCancelResponseResult) GetPending() string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CardCancelResponseResult) GetPendingOk() (*string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CardCancelResponseResult) SetPending(v string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *CardCancelResponseResult) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetLoopback

`func (o *CardCancelResponseResult) GetLoopback() string`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *CardCancelResponseResult) GetLoopbackOk() (*string, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *CardCancelResponseResult) SetLoopback(v string)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *CardCancelResponseResult) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetConnectedCenterId

`func (o *CardCancelResponseResult) GetConnectedCenterId() string`

GetConnectedCenterId returns the ConnectedCenterId field if non-nil, zero value otherwise.

### GetConnectedCenterIdOk

`func (o *CardCancelResponseResult) GetConnectedCenterIdOk() (*string, bool)`

GetConnectedCenterIdOk returns a tuple with the ConnectedCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedCenterId

`func (o *CardCancelResponseResult) SetConnectedCenterId(v string)`

SetConnectedCenterId sets ConnectedCenterId field to given value.

### HasConnectedCenterId

`func (o *CardCancelResponseResult) HasConnectedCenterId() bool`

HasConnectedCenterId returns a boolean if a field has been set.

### GetCenterRequestNumber

`func (o *CardCancelResponseResult) GetCenterRequestNumber() string`

GetCenterRequestNumber returns the CenterRequestNumber field if non-nil, zero value otherwise.

### GetCenterRequestNumberOk

`func (o *CardCancelResponseResult) GetCenterRequestNumberOk() (*string, bool)`

GetCenterRequestNumberOk returns a tuple with the CenterRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestNumber

`func (o *CardCancelResponseResult) SetCenterRequestNumber(v string)`

SetCenterRequestNumber sets CenterRequestNumber field to given value.

### HasCenterRequestNumber

`func (o *CardCancelResponseResult) HasCenterRequestNumber() bool`

HasCenterRequestNumber returns a boolean if a field has been set.

### GetCenterReferenceNumber

`func (o *CardCancelResponseResult) GetCenterReferenceNumber() string`

GetCenterReferenceNumber returns the CenterReferenceNumber field if non-nil, zero value otherwise.

### GetCenterReferenceNumberOk

`func (o *CardCancelResponseResult) GetCenterReferenceNumberOk() (*string, bool)`

GetCenterReferenceNumberOk returns a tuple with the CenterReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterReferenceNumber

`func (o *CardCancelResponseResult) SetCenterReferenceNumber(v string)`

SetCenterReferenceNumber sets CenterReferenceNumber field to given value.

### HasCenterReferenceNumber

`func (o *CardCancelResponseResult) HasCenterReferenceNumber() bool`

HasCenterReferenceNumber returns a boolean if a field has been set.

### GetResReturnReferenceNumber

`func (o *CardCancelResponseResult) GetResReturnReferenceNumber() string`

GetResReturnReferenceNumber returns the ResReturnReferenceNumber field if non-nil, zero value otherwise.

### GetResReturnReferenceNumberOk

`func (o *CardCancelResponseResult) GetResReturnReferenceNumberOk() (*string, bool)`

GetResReturnReferenceNumberOk returns a tuple with the ResReturnReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResReturnReferenceNumber

`func (o *CardCancelResponseResult) SetResReturnReferenceNumber(v string)`

SetResReturnReferenceNumber sets ResReturnReferenceNumber field to given value.

### HasResReturnReferenceNumber

`func (o *CardCancelResponseResult) HasResReturnReferenceNumber() bool`

HasResReturnReferenceNumber returns a boolean if a field has been set.

### GetResAuthCode

`func (o *CardCancelResponseResult) GetResAuthCode() string`

GetResAuthCode returns the ResAuthCode field if non-nil, zero value otherwise.

### GetResAuthCodeOk

`func (o *CardCancelResponseResult) GetResAuthCodeOk() (*string, bool)`

GetResAuthCodeOk returns a tuple with the ResAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResAuthCode

`func (o *CardCancelResponseResult) SetResAuthCode(v string)`

SetResAuthCode sets ResAuthCode field to given value.

### HasResAuthCode

`func (o *CardCancelResponseResult) HasResAuthCode() bool`

HasResAuthCode returns a boolean if a field has been set.

### GetResActionCode

`func (o *CardCancelResponseResult) GetResActionCode() string`

GetResActionCode returns the ResActionCode field if non-nil, zero value otherwise.

### GetResActionCodeOk

`func (o *CardCancelResponseResult) GetResActionCodeOk() (*string, bool)`

GetResActionCodeOk returns a tuple with the ResActionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResActionCode

`func (o *CardCancelResponseResult) SetResActionCode(v string)`

SetResActionCode sets ResActionCode field to given value.

### HasResActionCode

`func (o *CardCancelResponseResult) HasResActionCode() bool`

HasResActionCode returns a boolean if a field has been set.

### GetResCenterErrorCode

`func (o *CardCancelResponseResult) GetResCenterErrorCode() string`

GetResCenterErrorCode returns the ResCenterErrorCode field if non-nil, zero value otherwise.

### GetResCenterErrorCodeOk

`func (o *CardCancelResponseResult) GetResCenterErrorCodeOk() (*string, bool)`

GetResCenterErrorCodeOk returns a tuple with the ResCenterErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResCenterErrorCode

`func (o *CardCancelResponseResult) SetResCenterErrorCode(v string)`

SetResCenterErrorCode sets ResCenterErrorCode field to given value.

### HasResCenterErrorCode

`func (o *CardCancelResponseResult) HasResCenterErrorCode() bool`

HasResCenterErrorCode returns a boolean if a field has been set.

### GetAcquirerCode

`func (o *CardCancelResponseResult) GetAcquirerCode() string`

GetAcquirerCode returns the AcquirerCode field if non-nil, zero value otherwise.

### GetAcquirerCodeOk

`func (o *CardCancelResponseResult) GetAcquirerCodeOk() (*string, bool)`

GetAcquirerCodeOk returns a tuple with the AcquirerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerCode

`func (o *CardCancelResponseResult) SetAcquirerCode(v string)`

SetAcquirerCode sets AcquirerCode field to given value.

### HasAcquirerCode

`func (o *CardCancelResponseResult) HasAcquirerCode() bool`

HasAcquirerCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


