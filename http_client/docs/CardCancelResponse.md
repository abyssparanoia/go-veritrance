# CardCancelResponse

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

### GetServiceType

`func (o *CardCancelResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CardCancelResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CardCancelResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *CardCancelResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *CardCancelResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardCancelResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardCancelResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardCancelResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVResultCode

`func (o *CardCancelResponse) GetVResultCode() string`

GetVResultCode returns the VResultCode field if non-nil, zero value otherwise.

### GetVResultCodeOk

`func (o *CardCancelResponse) GetVResultCodeOk() (*string, bool)`

GetVResultCodeOk returns a tuple with the VResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVResultCode

`func (o *CardCancelResponse) SetVResultCode(v string)`

SetVResultCode sets VResultCode field to given value.

### HasVResultCode

`func (o *CardCancelResponse) HasVResultCode() bool`

HasVResultCode returns a boolean if a field has been set.

### GetMerrMsg

`func (o *CardCancelResponse) GetMerrMsg() string`

GetMerrMsg returns the MerrMsg field if non-nil, zero value otherwise.

### GetMerrMsgOk

`func (o *CardCancelResponse) GetMerrMsgOk() (*string, bool)`

GetMerrMsgOk returns a tuple with the MerrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerrMsg

`func (o *CardCancelResponse) SetMerrMsg(v string)`

SetMerrMsg sets MerrMsg field to given value.

### HasMerrMsg

`func (o *CardCancelResponse) HasMerrMsg() bool`

HasMerrMsg returns a boolean if a field has been set.

### GetMarchTxn

`func (o *CardCancelResponse) GetMarchTxn() string`

GetMarchTxn returns the MarchTxn field if non-nil, zero value otherwise.

### GetMarchTxnOk

`func (o *CardCancelResponse) GetMarchTxnOk() (*string, bool)`

GetMarchTxnOk returns a tuple with the MarchTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarchTxn

`func (o *CardCancelResponse) SetMarchTxn(v string)`

SetMarchTxn sets MarchTxn field to given value.

### HasMarchTxn

`func (o *CardCancelResponse) HasMarchTxn() bool`

HasMarchTxn returns a boolean if a field has been set.

### GetOrderId

`func (o *CardCancelResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardCancelResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardCancelResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CardCancelResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCustTxn

`func (o *CardCancelResponse) GetCustTxn() string`

GetCustTxn returns the CustTxn field if non-nil, zero value otherwise.

### GetCustTxnOk

`func (o *CardCancelResponse) GetCustTxnOk() (*string, bool)`

GetCustTxnOk returns a tuple with the CustTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustTxn

`func (o *CardCancelResponse) SetCustTxn(v string)`

SetCustTxn sets CustTxn field to given value.

### HasCustTxn

`func (o *CardCancelResponse) HasCustTxn() bool`

HasCustTxn returns a boolean if a field has been set.

### GetTxnVersion

`func (o *CardCancelResponse) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardCancelResponse) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardCancelResponse) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.

### HasTxnVersion

`func (o *CardCancelResponse) HasTxnVersion() bool`

HasTxnVersion returns a boolean if a field has been set.

### GetCardTransactiontype

`func (o *CardCancelResponse) GetCardTransactiontype() string`

GetCardTransactiontype returns the CardTransactiontype field if non-nil, zero value otherwise.

### GetCardTransactiontypeOk

`func (o *CardCancelResponse) GetCardTransactiontypeOk() (*string, bool)`

GetCardTransactiontypeOk returns a tuple with the CardTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTransactiontype

`func (o *CardCancelResponse) SetCardTransactiontype(v string)`

SetCardTransactiontype sets CardTransactiontype field to given value.

### HasCardTransactiontype

`func (o *CardCancelResponse) HasCardTransactiontype() bool`

HasCardTransactiontype returns a boolean if a field has been set.

### GetGatewayRequestDate

`func (o *CardCancelResponse) GetGatewayRequestDate() string`

GetGatewayRequestDate returns the GatewayRequestDate field if non-nil, zero value otherwise.

### GetGatewayRequestDateOk

`func (o *CardCancelResponse) GetGatewayRequestDateOk() (*string, bool)`

GetGatewayRequestDateOk returns a tuple with the GatewayRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRequestDate

`func (o *CardCancelResponse) SetGatewayRequestDate(v string)`

SetGatewayRequestDate sets GatewayRequestDate field to given value.

### HasGatewayRequestDate

`func (o *CardCancelResponse) HasGatewayRequestDate() bool`

HasGatewayRequestDate returns a boolean if a field has been set.

### GetGatewayResponseDate

`func (o *CardCancelResponse) GetGatewayResponseDate() string`

GetGatewayResponseDate returns the GatewayResponseDate field if non-nil, zero value otherwise.

### GetGatewayResponseDateOk

`func (o *CardCancelResponse) GetGatewayResponseDateOk() (*string, bool)`

GetGatewayResponseDateOk returns a tuple with the GatewayResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseDate

`func (o *CardCancelResponse) SetGatewayResponseDate(v string)`

SetGatewayResponseDate sets GatewayResponseDate field to given value.

### HasGatewayResponseDate

`func (o *CardCancelResponse) HasGatewayResponseDate() bool`

HasGatewayResponseDate returns a boolean if a field has been set.

### GetCenterRequestDate

`func (o *CardCancelResponse) GetCenterRequestDate() string`

GetCenterRequestDate returns the CenterRequestDate field if non-nil, zero value otherwise.

### GetCenterRequestDateOk

`func (o *CardCancelResponse) GetCenterRequestDateOk() (*string, bool)`

GetCenterRequestDateOk returns a tuple with the CenterRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestDate

`func (o *CardCancelResponse) SetCenterRequestDate(v string)`

SetCenterRequestDate sets CenterRequestDate field to given value.

### HasCenterRequestDate

`func (o *CardCancelResponse) HasCenterRequestDate() bool`

HasCenterRequestDate returns a boolean if a field has been set.

### GetCenterResponseDate

`func (o *CardCancelResponse) GetCenterResponseDate() string`

GetCenterResponseDate returns the CenterResponseDate field if non-nil, zero value otherwise.

### GetCenterResponseDateOk

`func (o *CardCancelResponse) GetCenterResponseDateOk() (*string, bool)`

GetCenterResponseDateOk returns a tuple with the CenterResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterResponseDate

`func (o *CardCancelResponse) SetCenterResponseDate(v string)`

SetCenterResponseDate sets CenterResponseDate field to given value.

### HasCenterResponseDate

`func (o *CardCancelResponse) HasCenterResponseDate() bool`

HasCenterResponseDate returns a boolean if a field has been set.

### GetPending

`func (o *CardCancelResponse) GetPending() string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CardCancelResponse) GetPendingOk() (*string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CardCancelResponse) SetPending(v string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *CardCancelResponse) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetLoopback

`func (o *CardCancelResponse) GetLoopback() string`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *CardCancelResponse) GetLoopbackOk() (*string, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *CardCancelResponse) SetLoopback(v string)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *CardCancelResponse) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetConnectedCenterId

`func (o *CardCancelResponse) GetConnectedCenterId() string`

GetConnectedCenterId returns the ConnectedCenterId field if non-nil, zero value otherwise.

### GetConnectedCenterIdOk

`func (o *CardCancelResponse) GetConnectedCenterIdOk() (*string, bool)`

GetConnectedCenterIdOk returns a tuple with the ConnectedCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedCenterId

`func (o *CardCancelResponse) SetConnectedCenterId(v string)`

SetConnectedCenterId sets ConnectedCenterId field to given value.

### HasConnectedCenterId

`func (o *CardCancelResponse) HasConnectedCenterId() bool`

HasConnectedCenterId returns a boolean if a field has been set.

### GetCenterRequestNumber

`func (o *CardCancelResponse) GetCenterRequestNumber() string`

GetCenterRequestNumber returns the CenterRequestNumber field if non-nil, zero value otherwise.

### GetCenterRequestNumberOk

`func (o *CardCancelResponse) GetCenterRequestNumberOk() (*string, bool)`

GetCenterRequestNumberOk returns a tuple with the CenterRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestNumber

`func (o *CardCancelResponse) SetCenterRequestNumber(v string)`

SetCenterRequestNumber sets CenterRequestNumber field to given value.

### HasCenterRequestNumber

`func (o *CardCancelResponse) HasCenterRequestNumber() bool`

HasCenterRequestNumber returns a boolean if a field has been set.

### GetCenterReferenceNumber

`func (o *CardCancelResponse) GetCenterReferenceNumber() string`

GetCenterReferenceNumber returns the CenterReferenceNumber field if non-nil, zero value otherwise.

### GetCenterReferenceNumberOk

`func (o *CardCancelResponse) GetCenterReferenceNumberOk() (*string, bool)`

GetCenterReferenceNumberOk returns a tuple with the CenterReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterReferenceNumber

`func (o *CardCancelResponse) SetCenterReferenceNumber(v string)`

SetCenterReferenceNumber sets CenterReferenceNumber field to given value.

### HasCenterReferenceNumber

`func (o *CardCancelResponse) HasCenterReferenceNumber() bool`

HasCenterReferenceNumber returns a boolean if a field has been set.

### GetResReturnReferenceNumber

`func (o *CardCancelResponse) GetResReturnReferenceNumber() string`

GetResReturnReferenceNumber returns the ResReturnReferenceNumber field if non-nil, zero value otherwise.

### GetResReturnReferenceNumberOk

`func (o *CardCancelResponse) GetResReturnReferenceNumberOk() (*string, bool)`

GetResReturnReferenceNumberOk returns a tuple with the ResReturnReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResReturnReferenceNumber

`func (o *CardCancelResponse) SetResReturnReferenceNumber(v string)`

SetResReturnReferenceNumber sets ResReturnReferenceNumber field to given value.

### HasResReturnReferenceNumber

`func (o *CardCancelResponse) HasResReturnReferenceNumber() bool`

HasResReturnReferenceNumber returns a boolean if a field has been set.

### GetResAuthCode

`func (o *CardCancelResponse) GetResAuthCode() string`

GetResAuthCode returns the ResAuthCode field if non-nil, zero value otherwise.

### GetResAuthCodeOk

`func (o *CardCancelResponse) GetResAuthCodeOk() (*string, bool)`

GetResAuthCodeOk returns a tuple with the ResAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResAuthCode

`func (o *CardCancelResponse) SetResAuthCode(v string)`

SetResAuthCode sets ResAuthCode field to given value.

### HasResAuthCode

`func (o *CardCancelResponse) HasResAuthCode() bool`

HasResAuthCode returns a boolean if a field has been set.

### GetResActionCode

`func (o *CardCancelResponse) GetResActionCode() string`

GetResActionCode returns the ResActionCode field if non-nil, zero value otherwise.

### GetResActionCodeOk

`func (o *CardCancelResponse) GetResActionCodeOk() (*string, bool)`

GetResActionCodeOk returns a tuple with the ResActionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResActionCode

`func (o *CardCancelResponse) SetResActionCode(v string)`

SetResActionCode sets ResActionCode field to given value.

### HasResActionCode

`func (o *CardCancelResponse) HasResActionCode() bool`

HasResActionCode returns a boolean if a field has been set.

### GetResCenterErrorCode

`func (o *CardCancelResponse) GetResCenterErrorCode() string`

GetResCenterErrorCode returns the ResCenterErrorCode field if non-nil, zero value otherwise.

### GetResCenterErrorCodeOk

`func (o *CardCancelResponse) GetResCenterErrorCodeOk() (*string, bool)`

GetResCenterErrorCodeOk returns a tuple with the ResCenterErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResCenterErrorCode

`func (o *CardCancelResponse) SetResCenterErrorCode(v string)`

SetResCenterErrorCode sets ResCenterErrorCode field to given value.

### HasResCenterErrorCode

`func (o *CardCancelResponse) HasResCenterErrorCode() bool`

HasResCenterErrorCode returns a boolean if a field has been set.

### GetAcquirerCode

`func (o *CardCancelResponse) GetAcquirerCode() string`

GetAcquirerCode returns the AcquirerCode field if non-nil, zero value otherwise.

### GetAcquirerCodeOk

`func (o *CardCancelResponse) GetAcquirerCodeOk() (*string, bool)`

GetAcquirerCodeOk returns a tuple with the AcquirerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerCode

`func (o *CardCancelResponse) SetAcquirerCode(v string)`

SetAcquirerCode sets AcquirerCode field to given value.

### HasAcquirerCode

`func (o *CardCancelResponse) HasAcquirerCode() bool`

HasAcquirerCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


