# CardAuthorizeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayNowIdResponse** | Pointer to [**AccountResponsePayNowIdResponse**](AccountResponsePayNowIdResponse.md) |  | [optional] 
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

### NewCardAuthorizeResponse

`func NewCardAuthorizeResponse(serviceType string, vResultCode string, mstatus Status, merrMsg string, marchTxn string, orderId string, custTxn string, txnVersion string, ) *CardAuthorizeResponse`

NewCardAuthorizeResponse instantiates a new CardAuthorizeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuthorizeResponseWithDefaults

`func NewCardAuthorizeResponseWithDefaults() *CardAuthorizeResponse`

NewCardAuthorizeResponseWithDefaults instantiates a new CardAuthorizeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayNowIdResponse

`func (o *CardAuthorizeResponse) GetPayNowIdResponse() AccountResponsePayNowIdResponse`

GetPayNowIdResponse returns the PayNowIdResponse field if non-nil, zero value otherwise.

### GetPayNowIdResponseOk

`func (o *CardAuthorizeResponse) GetPayNowIdResponseOk() (*AccountResponsePayNowIdResponse, bool)`

GetPayNowIdResponseOk returns a tuple with the PayNowIdResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowIdResponse

`func (o *CardAuthorizeResponse) SetPayNowIdResponse(v AccountResponsePayNowIdResponse)`

SetPayNowIdResponse sets PayNowIdResponse field to given value.

### HasPayNowIdResponse

`func (o *CardAuthorizeResponse) HasPayNowIdResponse() bool`

HasPayNowIdResponse returns a boolean if a field has been set.

### GetServiceType

`func (o *CardAuthorizeResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CardAuthorizeResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CardAuthorizeResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetStatus

`func (o *CardAuthorizeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardAuthorizeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardAuthorizeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardAuthorizeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVResultCode

`func (o *CardAuthorizeResponse) GetVResultCode() string`

GetVResultCode returns the VResultCode field if non-nil, zero value otherwise.

### GetVResultCodeOk

`func (o *CardAuthorizeResponse) GetVResultCodeOk() (*string, bool)`

GetVResultCodeOk returns a tuple with the VResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVResultCode

`func (o *CardAuthorizeResponse) SetVResultCode(v string)`

SetVResultCode sets VResultCode field to given value.


### GetMstatus

`func (o *CardAuthorizeResponse) GetMstatus() Status`

GetMstatus returns the Mstatus field if non-nil, zero value otherwise.

### GetMstatusOk

`func (o *CardAuthorizeResponse) GetMstatusOk() (*Status, bool)`

GetMstatusOk returns a tuple with the Mstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstatus

`func (o *CardAuthorizeResponse) SetMstatus(v Status)`

SetMstatus sets Mstatus field to given value.


### GetMerrMsg

`func (o *CardAuthorizeResponse) GetMerrMsg() string`

GetMerrMsg returns the MerrMsg field if non-nil, zero value otherwise.

### GetMerrMsgOk

`func (o *CardAuthorizeResponse) GetMerrMsgOk() (*string, bool)`

GetMerrMsgOk returns a tuple with the MerrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerrMsg

`func (o *CardAuthorizeResponse) SetMerrMsg(v string)`

SetMerrMsg sets MerrMsg field to given value.


### GetMarchTxn

`func (o *CardAuthorizeResponse) GetMarchTxn() string`

GetMarchTxn returns the MarchTxn field if non-nil, zero value otherwise.

### GetMarchTxnOk

`func (o *CardAuthorizeResponse) GetMarchTxnOk() (*string, bool)`

GetMarchTxnOk returns a tuple with the MarchTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarchTxn

`func (o *CardAuthorizeResponse) SetMarchTxn(v string)`

SetMarchTxn sets MarchTxn field to given value.


### GetOrderId

`func (o *CardAuthorizeResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardAuthorizeResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardAuthorizeResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetCustTxn

`func (o *CardAuthorizeResponse) GetCustTxn() string`

GetCustTxn returns the CustTxn field if non-nil, zero value otherwise.

### GetCustTxnOk

`func (o *CardAuthorizeResponse) GetCustTxnOk() (*string, bool)`

GetCustTxnOk returns a tuple with the CustTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustTxn

`func (o *CardAuthorizeResponse) SetCustTxn(v string)`

SetCustTxn sets CustTxn field to given value.


### GetTxnVersion

`func (o *CardAuthorizeResponse) GetTxnVersion() string`

GetTxnVersion returns the TxnVersion field if non-nil, zero value otherwise.

### GetTxnVersionOk

`func (o *CardAuthorizeResponse) GetTxnVersionOk() (*string, bool)`

GetTxnVersionOk returns a tuple with the TxnVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnVersion

`func (o *CardAuthorizeResponse) SetTxnVersion(v string)`

SetTxnVersion sets TxnVersion field to given value.


### GetCardTransactiontype

`func (o *CardAuthorizeResponse) GetCardTransactiontype() string`

GetCardTransactiontype returns the CardTransactiontype field if non-nil, zero value otherwise.

### GetCardTransactiontypeOk

`func (o *CardAuthorizeResponse) GetCardTransactiontypeOk() (*string, bool)`

GetCardTransactiontypeOk returns a tuple with the CardTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTransactiontype

`func (o *CardAuthorizeResponse) SetCardTransactiontype(v string)`

SetCardTransactiontype sets CardTransactiontype field to given value.

### HasCardTransactiontype

`func (o *CardAuthorizeResponse) HasCardTransactiontype() bool`

HasCardTransactiontype returns a boolean if a field has been set.

### GetGatewayRequestDate

`func (o *CardAuthorizeResponse) GetGatewayRequestDate() time.Time`

GetGatewayRequestDate returns the GatewayRequestDate field if non-nil, zero value otherwise.

### GetGatewayRequestDateOk

`func (o *CardAuthorizeResponse) GetGatewayRequestDateOk() (*time.Time, bool)`

GetGatewayRequestDateOk returns a tuple with the GatewayRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRequestDate

`func (o *CardAuthorizeResponse) SetGatewayRequestDate(v time.Time)`

SetGatewayRequestDate sets GatewayRequestDate field to given value.

### HasGatewayRequestDate

`func (o *CardAuthorizeResponse) HasGatewayRequestDate() bool`

HasGatewayRequestDate returns a boolean if a field has been set.

### GetGatewayResponseDate

`func (o *CardAuthorizeResponse) GetGatewayResponseDate() time.Time`

GetGatewayResponseDate returns the GatewayResponseDate field if non-nil, zero value otherwise.

### GetGatewayResponseDateOk

`func (o *CardAuthorizeResponse) GetGatewayResponseDateOk() (*time.Time, bool)`

GetGatewayResponseDateOk returns a tuple with the GatewayResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseDate

`func (o *CardAuthorizeResponse) SetGatewayResponseDate(v time.Time)`

SetGatewayResponseDate sets GatewayResponseDate field to given value.

### HasGatewayResponseDate

`func (o *CardAuthorizeResponse) HasGatewayResponseDate() bool`

HasGatewayResponseDate returns a boolean if a field has been set.

### GetCenterRequestDate

`func (o *CardAuthorizeResponse) GetCenterRequestDate() time.Time`

GetCenterRequestDate returns the CenterRequestDate field if non-nil, zero value otherwise.

### GetCenterRequestDateOk

`func (o *CardAuthorizeResponse) GetCenterRequestDateOk() (*time.Time, bool)`

GetCenterRequestDateOk returns a tuple with the CenterRequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestDate

`func (o *CardAuthorizeResponse) SetCenterRequestDate(v time.Time)`

SetCenterRequestDate sets CenterRequestDate field to given value.

### HasCenterRequestDate

`func (o *CardAuthorizeResponse) HasCenterRequestDate() bool`

HasCenterRequestDate returns a boolean if a field has been set.

### GetCenterResponseDate

`func (o *CardAuthorizeResponse) GetCenterResponseDate() time.Time`

GetCenterResponseDate returns the CenterResponseDate field if non-nil, zero value otherwise.

### GetCenterResponseDateOk

`func (o *CardAuthorizeResponse) GetCenterResponseDateOk() (*time.Time, bool)`

GetCenterResponseDateOk returns a tuple with the CenterResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterResponseDate

`func (o *CardAuthorizeResponse) SetCenterResponseDate(v time.Time)`

SetCenterResponseDate sets CenterResponseDate field to given value.

### HasCenterResponseDate

`func (o *CardAuthorizeResponse) HasCenterResponseDate() bool`

HasCenterResponseDate returns a boolean if a field has been set.

### GetPending

`func (o *CardAuthorizeResponse) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CardAuthorizeResponse) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CardAuthorizeResponse) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *CardAuthorizeResponse) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetLoopback

`func (o *CardAuthorizeResponse) GetLoopback() int32`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *CardAuthorizeResponse) GetLoopbackOk() (*int32, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *CardAuthorizeResponse) SetLoopback(v int32)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *CardAuthorizeResponse) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetConnectedCenterId

`func (o *CardAuthorizeResponse) GetConnectedCenterId() string`

GetConnectedCenterId returns the ConnectedCenterId field if non-nil, zero value otherwise.

### GetConnectedCenterIdOk

`func (o *CardAuthorizeResponse) GetConnectedCenterIdOk() (*string, bool)`

GetConnectedCenterIdOk returns a tuple with the ConnectedCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedCenterId

`func (o *CardAuthorizeResponse) SetConnectedCenterId(v string)`

SetConnectedCenterId sets ConnectedCenterId field to given value.

### HasConnectedCenterId

`func (o *CardAuthorizeResponse) HasConnectedCenterId() bool`

HasConnectedCenterId returns a boolean if a field has been set.

### GetCenterRequestNumber

`func (o *CardAuthorizeResponse) GetCenterRequestNumber() string`

GetCenterRequestNumber returns the CenterRequestNumber field if non-nil, zero value otherwise.

### GetCenterRequestNumberOk

`func (o *CardAuthorizeResponse) GetCenterRequestNumberOk() (*string, bool)`

GetCenterRequestNumberOk returns a tuple with the CenterRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterRequestNumber

`func (o *CardAuthorizeResponse) SetCenterRequestNumber(v string)`

SetCenterRequestNumber sets CenterRequestNumber field to given value.

### HasCenterRequestNumber

`func (o *CardAuthorizeResponse) HasCenterRequestNumber() bool`

HasCenterRequestNumber returns a boolean if a field has been set.

### GetCenterReferenceNumber

`func (o *CardAuthorizeResponse) GetCenterReferenceNumber() string`

GetCenterReferenceNumber returns the CenterReferenceNumber field if non-nil, zero value otherwise.

### GetCenterReferenceNumberOk

`func (o *CardAuthorizeResponse) GetCenterReferenceNumberOk() (*string, bool)`

GetCenterReferenceNumberOk returns a tuple with the CenterReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterReferenceNumber

`func (o *CardAuthorizeResponse) SetCenterReferenceNumber(v string)`

SetCenterReferenceNumber sets CenterReferenceNumber field to given value.

### HasCenterReferenceNumber

`func (o *CardAuthorizeResponse) HasCenterReferenceNumber() bool`

HasCenterReferenceNumber returns a boolean if a field has been set.

### GetReqAmount

`func (o *CardAuthorizeResponse) GetReqAmount() string`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *CardAuthorizeResponse) GetReqAmountOk() (*string, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *CardAuthorizeResponse) SetReqAmount(v string)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *CardAuthorizeResponse) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetResReturnReferenceNumber

`func (o *CardAuthorizeResponse) GetResReturnReferenceNumber() string`

GetResReturnReferenceNumber returns the ResReturnReferenceNumber field if non-nil, zero value otherwise.

### GetResReturnReferenceNumberOk

`func (o *CardAuthorizeResponse) GetResReturnReferenceNumberOk() (*string, bool)`

GetResReturnReferenceNumberOk returns a tuple with the ResReturnReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResReturnReferenceNumber

`func (o *CardAuthorizeResponse) SetResReturnReferenceNumber(v string)`

SetResReturnReferenceNumber sets ResReturnReferenceNumber field to given value.

### HasResReturnReferenceNumber

`func (o *CardAuthorizeResponse) HasResReturnReferenceNumber() bool`

HasResReturnReferenceNumber returns a boolean if a field has been set.

### GetResAuthCode

`func (o *CardAuthorizeResponse) GetResAuthCode() string`

GetResAuthCode returns the ResAuthCode field if non-nil, zero value otherwise.

### GetResAuthCodeOk

`func (o *CardAuthorizeResponse) GetResAuthCodeOk() (*string, bool)`

GetResAuthCodeOk returns a tuple with the ResAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResAuthCode

`func (o *CardAuthorizeResponse) SetResAuthCode(v string)`

SetResAuthCode sets ResAuthCode field to given value.

### HasResAuthCode

`func (o *CardAuthorizeResponse) HasResAuthCode() bool`

HasResAuthCode returns a boolean if a field has been set.

### GetResActionCode

`func (o *CardAuthorizeResponse) GetResActionCode() string`

GetResActionCode returns the ResActionCode field if non-nil, zero value otherwise.

### GetResActionCodeOk

`func (o *CardAuthorizeResponse) GetResActionCodeOk() (*string, bool)`

GetResActionCodeOk returns a tuple with the ResActionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResActionCode

`func (o *CardAuthorizeResponse) SetResActionCode(v string)`

SetResActionCode sets ResActionCode field to given value.

### HasResActionCode

`func (o *CardAuthorizeResponse) HasResActionCode() bool`

HasResActionCode returns a boolean if a field has been set.

### GetResCenterErrorCode

`func (o *CardAuthorizeResponse) GetResCenterErrorCode() string`

GetResCenterErrorCode returns the ResCenterErrorCode field if non-nil, zero value otherwise.

### GetResCenterErrorCodeOk

`func (o *CardAuthorizeResponse) GetResCenterErrorCodeOk() (*string, bool)`

GetResCenterErrorCodeOk returns a tuple with the ResCenterErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResCenterErrorCode

`func (o *CardAuthorizeResponse) SetResCenterErrorCode(v string)`

SetResCenterErrorCode sets ResCenterErrorCode field to given value.

### HasResCenterErrorCode

`func (o *CardAuthorizeResponse) HasResCenterErrorCode() bool`

HasResCenterErrorCode returns a boolean if a field has been set.

### GetAcquirerCode

`func (o *CardAuthorizeResponse) GetAcquirerCode() string`

GetAcquirerCode returns the AcquirerCode field if non-nil, zero value otherwise.

### GetAcquirerCodeOk

`func (o *CardAuthorizeResponse) GetAcquirerCodeOk() (*string, bool)`

GetAcquirerCodeOk returns a tuple with the AcquirerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerCode

`func (o *CardAuthorizeResponse) SetAcquirerCode(v string)`

SetAcquirerCode sets AcquirerCode field to given value.

### HasAcquirerCode

`func (o *CardAuthorizeResponse) HasAcquirerCode() bool`

HasAcquirerCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


