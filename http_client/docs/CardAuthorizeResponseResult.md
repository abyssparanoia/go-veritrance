# CardAuthorizeResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceType** | **string** | 要求電文を送信した決済サービスタイプ | 
**Status** | Pointer to **string** | 処理結果コード | [optional] 
**VResultCode** | **string** | 処理の結果を詳細に表すコード 4 桁ずつ 4 つのブロックで構成され、各ブロックでサービス毎の処理結果を表します。  | 
**Mstatus** | [**Status**](Status.md) |  | 
**MerrMsg** | **string** | 処理結果を日本語で表示します。 | 
**MarchTxn** | **string** | 決済サーバーにて決済処理電文（内部処理も含む）毎に付与する ID １つの取引 ID に対して、複数の ID が付与されます。  | 
**OrderId** | **string** | 決済要求時に店舗様にて任意に採番し送信された取引 ID | 
**CustTxn** | **string** | 決済サーバーがオーダー（取引 ID）と紐付ける為に採番する ID | 
**TxnVersion** | **string** | 電文のバージョン | 
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
**ReqAmount** | Pointer to **string** |  | [optional] 
**ResReturnReferenceNumber** | Pointer to **string** |  | [optional] 
**ResAuthCode** | Pointer to **string** |  | [optional] 
**ResActionCode** | Pointer to **string** |  | [optional] 
**ResCenterErrorCode** | Pointer to **string** |  | [optional] 
**AcquirerCode** | Pointer to **string** |  | [optional] 

## Methods

### NewCardAuthorizeResponseResult

`func NewCardAuthorizeResponseResult(serviceType string, vResultCode string, mstatus Status, merrMsg string, marchTxn string, orderId string, custTxn string, txnVersion string, ) *CardAuthorizeResponseResult`

NewCardAuthorizeResponseResult instantiates a new CardAuthorizeResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuthorizeResponseResultWithDefaults

`func NewCardAuthorizeResponseResultWithDefaults() *CardAuthorizeResponseResult`

NewCardAuthorizeResponseResultWithDefaults instantiates a new CardAuthorizeResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceType

`func (o *CardAuthorizeResponseResult) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CardAuthorizeResponseResult) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CardAuthorizeResponseResult) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetStatus

`func (o *CardAuthorizeResponseResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardAuthorizeResponseResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardAuthorizeResponseResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardAuthorizeResponseResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVResultCode

`func (o *CardAuthorizeResponseResult) GetVResultCode() string`

GetVResultCode returns the VResultCode field if non-nil, zero value otherwise.

### GetVResultCodeOk

`func (o *CardAuthorizeResponseResult) GetVResultCodeOk() (*string, bool)`

GetVResultCodeOk returns a tuple with the VResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVResultCode

`func (o *CardAuthorizeResponseResult) SetVResultCode(v string)`

SetVResultCode sets VResultCode field to given value.


### GetMstatus

`func (o *CardAuthorizeResponseResult) GetMstatus() Status`

GetMstatus returns the Mstatus field if non-nil, zero value otherwise.

### GetMstatusOk

`func (o *CardAuthorizeResponseResult) GetMstatusOk() (*Status, bool)`

GetMstatusOk returns a tuple with the Mstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstatus

`func (o *CardAuthorizeResponseResult) SetMstatus(v Status)`

SetMstatus sets Mstatus field to given value.


### GetMerrMsg

`func (o *CardAuthorizeResponseResult) GetMerrMsg() string`

GetMerrMsg returns the MerrMsg field if non-nil, zero value otherwise.

### GetMerrMsgOk

`func (o *CardAuthorizeResponseResult) GetMerrMsgOk() (*string, bool)`

GetMerrMsgOk returns a tuple with the MerrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerrMsg

`func (o *CardAuthorizeResponseResult) SetMerrMsg(v string)`

SetMerrMsg sets MerrMsg field to given value.


### GetMarchTxn

`func (o *CardAuthorizeResponseResult) GetMarchTxn() string`

GetMarchTxn returns the MarchTxn field if non-nil, zero value otherwise.

### GetMarchTxnOk

`func (o *CardAuthorizeResponseResult) GetMarchTxnOk() (*string, bool)`

GetMarchTxnOk returns a tuple with the MarchTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarchTxn

`func (o *CardAuthorizeResponseResult) SetMarchTxn(v string)`

SetMarchTxn sets MarchTxn field to given value.


### GetOrderId

`func (o *CardAuthorizeResponseResult) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardAuthorizeResponseResult) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardAuthorizeResponseResult) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetCustTxn

`func (o *CardAuthorizeResponseResult) GetCustTxn() string`

GetCustTxn returns the CustTxn field if non-nil, zero value otherwise.

### GetCustTxnOk

`func (o *CardAuthorizeResponseResult) GetCustTxnOk() (*string, bool)`

GetCustTxnOk returns a tuple with the CustTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustTxn

`func (o *CardAuthorizeResponseResult) SetCustTxn(v string)`

SetCustTxn sets CustTxn field to given value.


### GetTxnVersion

`func (o *CardAuthorizeResponseResult) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardAuthorizeResponseResult) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardAuthorizeResponseResult) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetCardTransactiontype

`func (o *CardAuthorizeResponseResult) GetCardTransactiontype() string`

GetCardTransactiontype returns the CardTransactiontype field if non-nil, zero value otherwise.

### GetCardTransactiontypeOk

`func (o *CardAuthorizeResponseResult) GetCardTransactiontypeOk() (*string, bool)`

GetCardTransactiontypeOk returns a tuple with the CardTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTransactiontype

`func (o *CardAuthorizeResponseResult) SetCardTransactiontype(v string)`

SetCardTransactiontype sets CardTransactiontype field to given value.

### HasCardTransactiontype

`func (o *CardAuthorizeResponseResult) HasCardTransactiontype() bool`

HasCardTransactiontype returns a boolean if a field has been set.

### GetGatewayRequestDate

`func (o *CardAuthorizeResponseResult) GetGatewayRequestDate() time.Time`

GetGatewayRequestDate returns the GatewayRequestDate field if non-nil, zero value otherwise.

### GetGatewayRequestDateOk

`func (o *CardAuthorizeResponseResult) GetGatewayRequestDateOk() (*time.Time, bool)`

GetGatewayRequestDateOk returns a tuple with the GatewayRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRequestDate

`func (o *CardAuthorizeResponseResult) SetGatewayRequestDate(v time.Time)`

SetGatewayRequestDate sets GatewayRequestDate field to given value.

### HasGatewayRequestDate

`func (o *CardAuthorizeResponseResult) HasGatewayRequestDate() bool`

HasGatewayRequestDate returns a boolean if a field has been set.

### GetGatewayResponseDate

`func (o *CardAuthorizeResponseResult) GetGatewayResponseDate() time.Time`

GetGatewayResponseDate returns the GatewayResponseDate field if non-nil, zero value otherwise.

### GetGatewayResponseDateOk

`func (o *CardAuthorizeResponseResult) GetGatewayResponseDateOk() (*time.Time, bool)`

GetGatewayResponseDateOk returns a tuple with the GatewayResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseDate

`func (o *CardAuthorizeResponseResult) SetGatewayResponseDate(v time.Time)`

SetGatewayResponseDate sets GatewayResponseDate field to given value.

### HasGatewayResponseDate

`func (o *CardAuthorizeResponseResult) HasGatewayResponseDate() bool`

HasGatewayResponseDate returns a boolean if a field has been set.

### GetCenterRequestDate

`func (o *CardAuthorizeResponseResult) GetCenterRequestDate() time.Time`

GetCenterRequestDate returns the CenterRequestDate field if non-nil, zero value otherwise.

### GetCenterRequestDateOk

`func (o *CardAuthorizeResponseResult) GetCenterRequestDateOk() (*time.Time, bool)`

GetCenterRequestDateOk returns a tuple with the CenterRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestDate

`func (o *CardAuthorizeResponseResult) SetCenterRequestDate(v time.Time)`

SetCenterRequestDate sets CenterRequestDate field to given value.

### HasCenterRequestDate

`func (o *CardAuthorizeResponseResult) HasCenterRequestDate() bool`

HasCenterRequestDate returns a boolean if a field has been set.

### GetCenterResponseDate

`func (o *CardAuthorizeResponseResult) GetCenterResponseDate() time.Time`

GetCenterResponseDate returns the CenterResponseDate field if non-nil, zero value otherwise.

### GetCenterResponseDateOk

`func (o *CardAuthorizeResponseResult) GetCenterResponseDateOk() (*time.Time, bool)`

GetCenterResponseDateOk returns a tuple with the CenterResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterResponseDate

`func (o *CardAuthorizeResponseResult) SetCenterResponseDate(v time.Time)`

SetCenterResponseDate sets CenterResponseDate field to given value.

### HasCenterResponseDate

`func (o *CardAuthorizeResponseResult) HasCenterResponseDate() bool`

HasCenterResponseDate returns a boolean if a field has been set.

### GetPending

`func (o *CardAuthorizeResponseResult) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CardAuthorizeResponseResult) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CardAuthorizeResponseResult) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *CardAuthorizeResponseResult) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetLoopback

`func (o *CardAuthorizeResponseResult) GetLoopback() int32`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *CardAuthorizeResponseResult) GetLoopbackOk() (*int32, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *CardAuthorizeResponseResult) SetLoopback(v int32)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *CardAuthorizeResponseResult) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetConnectedCenterId

`func (o *CardAuthorizeResponseResult) GetConnectedCenterId() string`

GetConnectedCenterId returns the ConnectedCenterId field if non-nil, zero value otherwise.

### GetConnectedCenterIdOk

`func (o *CardAuthorizeResponseResult) GetConnectedCenterIdOk() (*string, bool)`

GetConnectedCenterIdOk returns a tuple with the ConnectedCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedCenterId

`func (o *CardAuthorizeResponseResult) SetConnectedCenterId(v string)`

SetConnectedCenterId sets ConnectedCenterId field to given value.

### HasConnectedCenterId

`func (o *CardAuthorizeResponseResult) HasConnectedCenterId() bool`

HasConnectedCenterId returns a boolean if a field has been set.

### GetCenterRequestNumber

`func (o *CardAuthorizeResponseResult) GetCenterRequestNumber() string`

GetCenterRequestNumber returns the CenterRequestNumber field if non-nil, zero value otherwise.

### GetCenterRequestNumberOk

`func (o *CardAuthorizeResponseResult) GetCenterRequestNumberOk() (*string, bool)`

GetCenterRequestNumberOk returns a tuple with the CenterRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestNumber

`func (o *CardAuthorizeResponseResult) SetCenterRequestNumber(v string)`

SetCenterRequestNumber sets CenterRequestNumber field to given value.

### HasCenterRequestNumber

`func (o *CardAuthorizeResponseResult) HasCenterRequestNumber() bool`

HasCenterRequestNumber returns a boolean if a field has been set.

### GetCenterReferenceNumber

`func (o *CardAuthorizeResponseResult) GetCenterReferenceNumber() string`

GetCenterReferenceNumber returns the CenterReferenceNumber field if non-nil, zero value otherwise.

### GetCenterReferenceNumberOk

`func (o *CardAuthorizeResponseResult) GetCenterReferenceNumberOk() (*string, bool)`

GetCenterReferenceNumberOk returns a tuple with the CenterReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterReferenceNumber

`func (o *CardAuthorizeResponseResult) SetCenterReferenceNumber(v string)`

SetCenterReferenceNumber sets CenterReferenceNumber field to given value.

### HasCenterReferenceNumber

`func (o *CardAuthorizeResponseResult) HasCenterReferenceNumber() bool`

HasCenterReferenceNumber returns a boolean if a field has been set.

### GetReqAmount

`func (o *CardAuthorizeResponseResult) GetReqAmount() string`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *CardAuthorizeResponseResult) GetReqAmountOk() (*string, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *CardAuthorizeResponseResult) SetReqAmount(v string)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *CardAuthorizeResponseResult) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetResReturnReferenceNumber

`func (o *CardAuthorizeResponseResult) GetResReturnReferenceNumber() string`

GetResReturnReferenceNumber returns the ResReturnReferenceNumber field if non-nil, zero value otherwise.

### GetResReturnReferenceNumberOk

`func (o *CardAuthorizeResponseResult) GetResReturnReferenceNumberOk() (*string, bool)`

GetResReturnReferenceNumberOk returns a tuple with the ResReturnReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResReturnReferenceNumber

`func (o *CardAuthorizeResponseResult) SetResReturnReferenceNumber(v string)`

SetResReturnReferenceNumber sets ResReturnReferenceNumber field to given value.

### HasResReturnReferenceNumber

`func (o *CardAuthorizeResponseResult) HasResReturnReferenceNumber() bool`

HasResReturnReferenceNumber returns a boolean if a field has been set.

### GetResAuthCode

`func (o *CardAuthorizeResponseResult) GetResAuthCode() string`

GetResAuthCode returns the ResAuthCode field if non-nil, zero value otherwise.

### GetResAuthCodeOk

`func (o *CardAuthorizeResponseResult) GetResAuthCodeOk() (*string, bool)`

GetResAuthCodeOk returns a tuple with the ResAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResAuthCode

`func (o *CardAuthorizeResponseResult) SetResAuthCode(v string)`

SetResAuthCode sets ResAuthCode field to given value.

### HasResAuthCode

`func (o *CardAuthorizeResponseResult) HasResAuthCode() bool`

HasResAuthCode returns a boolean if a field has been set.

### GetResActionCode

`func (o *CardAuthorizeResponseResult) GetResActionCode() string`

GetResActionCode returns the ResActionCode field if non-nil, zero value otherwise.

### GetResActionCodeOk

`func (o *CardAuthorizeResponseResult) GetResActionCodeOk() (*string, bool)`

GetResActionCodeOk returns a tuple with the ResActionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResActionCode

`func (o *CardAuthorizeResponseResult) SetResActionCode(v string)`

SetResActionCode sets ResActionCode field to given value.

### HasResActionCode

`func (o *CardAuthorizeResponseResult) HasResActionCode() bool`

HasResActionCode returns a boolean if a field has been set.

### GetResCenterErrorCode

`func (o *CardAuthorizeResponseResult) GetResCenterErrorCode() string`

GetResCenterErrorCode returns the ResCenterErrorCode field if non-nil, zero value otherwise.

### GetResCenterErrorCodeOk

`func (o *CardAuthorizeResponseResult) GetResCenterErrorCodeOk() (*string, bool)`

GetResCenterErrorCodeOk returns a tuple with the ResCenterErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResCenterErrorCode

`func (o *CardAuthorizeResponseResult) SetResCenterErrorCode(v string)`

SetResCenterErrorCode sets ResCenterErrorCode field to given value.

### HasResCenterErrorCode

`func (o *CardAuthorizeResponseResult) HasResCenterErrorCode() bool`

HasResCenterErrorCode returns a boolean if a field has been set.

### GetAcquirerCode

`func (o *CardAuthorizeResponseResult) GetAcquirerCode() string`

GetAcquirerCode returns the AcquirerCode field if non-nil, zero value otherwise.

### GetAcquirerCodeOk

`func (o *CardAuthorizeResponseResult) GetAcquirerCodeOk() (*string, bool)`

GetAcquirerCodeOk returns a tuple with the AcquirerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerCode

`func (o *CardAuthorizeResponseResult) SetAcquirerCode(v string)`

SetAcquirerCode sets AcquirerCode field to given value.

### HasAcquirerCode

`func (o *CardAuthorizeResponseResult) HasAcquirerCode() bool`

HasAcquirerCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


