/*
Veritrance API

Veritrance API Requestの形式はcomponentsを参照。文字列にしたものを各Requestのparamsに用いる 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package veritrance_paynow

import (
	"encoding/json"
)

// CardAuthorizeResponseResult struct for CardAuthorizeResponseResult
type CardAuthorizeResponseResult struct {
	// 要求電文を送信した決済サービスタイプ
	ServiceType string `json:"serviceType"`
	// 処理結果コード
	Status *string `json:"status,omitempty"`
	// 処理の結果を詳細に表すコード 4 桁ずつ 4 つのブロックで構成され、各ブロックでサービス毎の処理結果を表します。 
	VResultCode string `json:"vResultCode"`
	Mstatus Status `json:"mstatus"`
	// 処理結果を日本語で表示します。
	MerrMsg string `json:"merrMsg"`
	// 決済サーバーにて決済処理電文（内部処理も含む）毎に付与する ID １つの取引 ID に対して、複数の ID が付与されます。 
	MarchTxn string `json:"marchTxn"`
	// 決済要求時に店舗様にて任意に採番し送信された取引 ID
	OrderId string `json:"orderId"`
	// 決済サーバーがオーダー（取引 ID）と紐付ける為に採番する ID
	CustTxn string `json:"custTxn"`
	// 電文のバージョン
	TxnVersion string `json:"txnVersion"`
	CardTransactiontype *string `json:"cardTransactiontype,omitempty"`
	GatewayRequestDate *string `json:"gatewayRequestDate,omitempty"`
	GatewayResponseDate *string `json:"gatewayResponseDate,omitempty"`
	CenterRequestDate *string `json:"centerRequestDate,omitempty"`
	CenterResponseDate *string `json:"centerResponseDate,omitempty"`
	Pending *int32 `json:"pending,omitempty"`
	Loopback *int32 `json:"loopback,omitempty"`
	ConnectedCenterId *string `json:"connectedCenterId,omitempty"`
	CenterRequestNumber *string `json:"centerRequestNumber,omitempty"`
	CenterReferenceNumber *string `json:"centerReferenceNumber,omitempty"`
	ReqAmount *string `json:"reqAmount,omitempty"`
	ResReturnReferenceNumber *string `json:"resReturnReferenceNumber,omitempty"`
	ResAuthCode *string `json:"resAuthCode,omitempty"`
	ResActionCode *string `json:"resActionCode,omitempty"`
	ResCenterErrorCode *string `json:"resCenterErrorCode,omitempty"`
	AcquirerCode *string `json:"acquirerCode,omitempty"`
}

// NewCardAuthorizeResponseResult instantiates a new CardAuthorizeResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardAuthorizeResponseResult(serviceType string, vResultCode string, mstatus Status, merrMsg string, marchTxn string, orderId string, custTxn string, txnVersion string) *CardAuthorizeResponseResult {
	this := CardAuthorizeResponseResult{}
	this.ServiceType = serviceType
	this.VResultCode = vResultCode
	this.Mstatus = mstatus
	this.MerrMsg = merrMsg
	this.MarchTxn = marchTxn
	this.OrderId = orderId
	this.CustTxn = custTxn
	this.TxnVersion = txnVersion
	return &this
}

// NewCardAuthorizeResponseResultWithDefaults instantiates a new CardAuthorizeResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardAuthorizeResponseResultWithDefaults() *CardAuthorizeResponseResult {
	this := CardAuthorizeResponseResult{}
	return &this
}

// GetServiceType returns the ServiceType field value
func (o *CardAuthorizeResponseResult) GetServiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetServiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *CardAuthorizeResponseResult) SetServiceType(v string) {
	o.ServiceType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CardAuthorizeResponseResult) SetStatus(v string) {
	o.Status = &v
}

// GetVResultCode returns the VResultCode field value
func (o *CardAuthorizeResponseResult) GetVResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VResultCode
}

// GetVResultCodeOk returns a tuple with the VResultCode field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetVResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VResultCode, true
}

// SetVResultCode sets field value
func (o *CardAuthorizeResponseResult) SetVResultCode(v string) {
	o.VResultCode = v
}

// GetMstatus returns the Mstatus field value
func (o *CardAuthorizeResponseResult) GetMstatus() Status {
	if o == nil {
		var ret Status
		return ret
	}

	return o.Mstatus
}

// GetMstatusOk returns a tuple with the Mstatus field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetMstatusOk() (*Status, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mstatus, true
}

// SetMstatus sets field value
func (o *CardAuthorizeResponseResult) SetMstatus(v Status) {
	o.Mstatus = v
}

// GetMerrMsg returns the MerrMsg field value
func (o *CardAuthorizeResponseResult) GetMerrMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerrMsg
}

// GetMerrMsgOk returns a tuple with the MerrMsg field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetMerrMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerrMsg, true
}

// SetMerrMsg sets field value
func (o *CardAuthorizeResponseResult) SetMerrMsg(v string) {
	o.MerrMsg = v
}

// GetMarchTxn returns the MarchTxn field value
func (o *CardAuthorizeResponseResult) GetMarchTxn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarchTxn
}

// GetMarchTxnOk returns a tuple with the MarchTxn field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetMarchTxnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarchTxn, true
}

// SetMarchTxn sets field value
func (o *CardAuthorizeResponseResult) SetMarchTxn(v string) {
	o.MarchTxn = v
}

// GetOrderId returns the OrderId field value
func (o *CardAuthorizeResponseResult) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CardAuthorizeResponseResult) SetOrderId(v string) {
	o.OrderId = v
}

// GetCustTxn returns the CustTxn field value
func (o *CardAuthorizeResponseResult) GetCustTxn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustTxn
}

// GetCustTxnOk returns a tuple with the CustTxn field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetCustTxnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustTxn, true
}

// SetCustTxn sets field value
func (o *CardAuthorizeResponseResult) SetCustTxn(v string) {
	o.CustTxn = v
}

// GetTxnVersion returns the TxnVersion field value
func (o *CardAuthorizeResponseResult) GetTxnVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnVersion
}

// GetTxnVersionOk returns a tuple with the TxnVersion field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetTxnVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnVersion, true
}

// SetTxnVersion sets field value
func (o *CardAuthorizeResponseResult) SetTxnVersion(v string) {
	o.TxnVersion = v
}

// GetCardTransactiontype returns the CardTransactiontype field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetCardTransactiontype() string {
	if o == nil || o.CardTransactiontype == nil {
		var ret string
		return ret
	}
	return *o.CardTransactiontype
}

// GetCardTransactiontypeOk returns a tuple with the CardTransactiontype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetCardTransactiontypeOk() (*string, bool) {
	if o == nil || o.CardTransactiontype == nil {
		return nil, false
	}
	return o.CardTransactiontype, true
}

// HasCardTransactiontype returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasCardTransactiontype() bool {
	if o != nil && o.CardTransactiontype != nil {
		return true
	}

	return false
}

// SetCardTransactiontype gets a reference to the given string and assigns it to the CardTransactiontype field.
func (o *CardAuthorizeResponseResult) SetCardTransactiontype(v string) {
	o.CardTransactiontype = &v
}

// GetGatewayRequestDate returns the GatewayRequestDate field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetGatewayRequestDate() string {
	if o == nil || o.GatewayRequestDate == nil {
		var ret string
		return ret
	}
	return *o.GatewayRequestDate
}

// GetGatewayRequestDateOk returns a tuple with the GatewayRequestDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetGatewayRequestDateOk() (*string, bool) {
	if o == nil || o.GatewayRequestDate == nil {
		return nil, false
	}
	return o.GatewayRequestDate, true
}

// HasGatewayRequestDate returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasGatewayRequestDate() bool {
	if o != nil && o.GatewayRequestDate != nil {
		return true
	}

	return false
}

// SetGatewayRequestDate gets a reference to the given string and assigns it to the GatewayRequestDate field.
func (o *CardAuthorizeResponseResult) SetGatewayRequestDate(v string) {
	o.GatewayRequestDate = &v
}

// GetGatewayResponseDate returns the GatewayResponseDate field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetGatewayResponseDate() string {
	if o == nil || o.GatewayResponseDate == nil {
		var ret string
		return ret
	}
	return *o.GatewayResponseDate
}

// GetGatewayResponseDateOk returns a tuple with the GatewayResponseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetGatewayResponseDateOk() (*string, bool) {
	if o == nil || o.GatewayResponseDate == nil {
		return nil, false
	}
	return o.GatewayResponseDate, true
}

// HasGatewayResponseDate returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasGatewayResponseDate() bool {
	if o != nil && o.GatewayResponseDate != nil {
		return true
	}

	return false
}

// SetGatewayResponseDate gets a reference to the given string and assigns it to the GatewayResponseDate field.
func (o *CardAuthorizeResponseResult) SetGatewayResponseDate(v string) {
	o.GatewayResponseDate = &v
}

// GetCenterRequestDate returns the CenterRequestDate field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetCenterRequestDate() string {
	if o == nil || o.CenterRequestDate == nil {
		var ret string
		return ret
	}
	return *o.CenterRequestDate
}

// GetCenterRequestDateOk returns a tuple with the CenterRequestDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetCenterRequestDateOk() (*string, bool) {
	if o == nil || o.CenterRequestDate == nil {
		return nil, false
	}
	return o.CenterRequestDate, true
}

// HasCenterRequestDate returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasCenterRequestDate() bool {
	if o != nil && o.CenterRequestDate != nil {
		return true
	}

	return false
}

// SetCenterRequestDate gets a reference to the given string and assigns it to the CenterRequestDate field.
func (o *CardAuthorizeResponseResult) SetCenterRequestDate(v string) {
	o.CenterRequestDate = &v
}

// GetCenterResponseDate returns the CenterResponseDate field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetCenterResponseDate() string {
	if o == nil || o.CenterResponseDate == nil {
		var ret string
		return ret
	}
	return *o.CenterResponseDate
}

// GetCenterResponseDateOk returns a tuple with the CenterResponseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetCenterResponseDateOk() (*string, bool) {
	if o == nil || o.CenterResponseDate == nil {
		return nil, false
	}
	return o.CenterResponseDate, true
}

// HasCenterResponseDate returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasCenterResponseDate() bool {
	if o != nil && o.CenterResponseDate != nil {
		return true
	}

	return false
}

// SetCenterResponseDate gets a reference to the given string and assigns it to the CenterResponseDate field.
func (o *CardAuthorizeResponseResult) SetCenterResponseDate(v string) {
	o.CenterResponseDate = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetPending() int32 {
	if o == nil || o.Pending == nil {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetPendingOk() (*int32, bool) {
	if o == nil || o.Pending == nil {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasPending() bool {
	if o != nil && o.Pending != nil {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *CardAuthorizeResponseResult) SetPending(v int32) {
	o.Pending = &v
}

// GetLoopback returns the Loopback field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetLoopback() int32 {
	if o == nil || o.Loopback == nil {
		var ret int32
		return ret
	}
	return *o.Loopback
}

// GetLoopbackOk returns a tuple with the Loopback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetLoopbackOk() (*int32, bool) {
	if o == nil || o.Loopback == nil {
		return nil, false
	}
	return o.Loopback, true
}

// HasLoopback returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasLoopback() bool {
	if o != nil && o.Loopback != nil {
		return true
	}

	return false
}

// SetLoopback gets a reference to the given int32 and assigns it to the Loopback field.
func (o *CardAuthorizeResponseResult) SetLoopback(v int32) {
	o.Loopback = &v
}

// GetConnectedCenterId returns the ConnectedCenterId field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetConnectedCenterId() string {
	if o == nil || o.ConnectedCenterId == nil {
		var ret string
		return ret
	}
	return *o.ConnectedCenterId
}

// GetConnectedCenterIdOk returns a tuple with the ConnectedCenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetConnectedCenterIdOk() (*string, bool) {
	if o == nil || o.ConnectedCenterId == nil {
		return nil, false
	}
	return o.ConnectedCenterId, true
}

// HasConnectedCenterId returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasConnectedCenterId() bool {
	if o != nil && o.ConnectedCenterId != nil {
		return true
	}

	return false
}

// SetConnectedCenterId gets a reference to the given string and assigns it to the ConnectedCenterId field.
func (o *CardAuthorizeResponseResult) SetConnectedCenterId(v string) {
	o.ConnectedCenterId = &v
}

// GetCenterRequestNumber returns the CenterRequestNumber field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetCenterRequestNumber() string {
	if o == nil || o.CenterRequestNumber == nil {
		var ret string
		return ret
	}
	return *o.CenterRequestNumber
}

// GetCenterRequestNumberOk returns a tuple with the CenterRequestNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetCenterRequestNumberOk() (*string, bool) {
	if o == nil || o.CenterRequestNumber == nil {
		return nil, false
	}
	return o.CenterRequestNumber, true
}

// HasCenterRequestNumber returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasCenterRequestNumber() bool {
	if o != nil && o.CenterRequestNumber != nil {
		return true
	}

	return false
}

// SetCenterRequestNumber gets a reference to the given string and assigns it to the CenterRequestNumber field.
func (o *CardAuthorizeResponseResult) SetCenterRequestNumber(v string) {
	o.CenterRequestNumber = &v
}

// GetCenterReferenceNumber returns the CenterReferenceNumber field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetCenterReferenceNumber() string {
	if o == nil || o.CenterReferenceNumber == nil {
		var ret string
		return ret
	}
	return *o.CenterReferenceNumber
}

// GetCenterReferenceNumberOk returns a tuple with the CenterReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetCenterReferenceNumberOk() (*string, bool) {
	if o == nil || o.CenterReferenceNumber == nil {
		return nil, false
	}
	return o.CenterReferenceNumber, true
}

// HasCenterReferenceNumber returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasCenterReferenceNumber() bool {
	if o != nil && o.CenterReferenceNumber != nil {
		return true
	}

	return false
}

// SetCenterReferenceNumber gets a reference to the given string and assigns it to the CenterReferenceNumber field.
func (o *CardAuthorizeResponseResult) SetCenterReferenceNumber(v string) {
	o.CenterReferenceNumber = &v
}

// GetReqAmount returns the ReqAmount field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetReqAmount() string {
	if o == nil || o.ReqAmount == nil {
		var ret string
		return ret
	}
	return *o.ReqAmount
}

// GetReqAmountOk returns a tuple with the ReqAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetReqAmountOk() (*string, bool) {
	if o == nil || o.ReqAmount == nil {
		return nil, false
	}
	return o.ReqAmount, true
}

// HasReqAmount returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasReqAmount() bool {
	if o != nil && o.ReqAmount != nil {
		return true
	}

	return false
}

// SetReqAmount gets a reference to the given string and assigns it to the ReqAmount field.
func (o *CardAuthorizeResponseResult) SetReqAmount(v string) {
	o.ReqAmount = &v
}

// GetResReturnReferenceNumber returns the ResReturnReferenceNumber field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetResReturnReferenceNumber() string {
	if o == nil || o.ResReturnReferenceNumber == nil {
		var ret string
		return ret
	}
	return *o.ResReturnReferenceNumber
}

// GetResReturnReferenceNumberOk returns a tuple with the ResReturnReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetResReturnReferenceNumberOk() (*string, bool) {
	if o == nil || o.ResReturnReferenceNumber == nil {
		return nil, false
	}
	return o.ResReturnReferenceNumber, true
}

// HasResReturnReferenceNumber returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasResReturnReferenceNumber() bool {
	if o != nil && o.ResReturnReferenceNumber != nil {
		return true
	}

	return false
}

// SetResReturnReferenceNumber gets a reference to the given string and assigns it to the ResReturnReferenceNumber field.
func (o *CardAuthorizeResponseResult) SetResReturnReferenceNumber(v string) {
	o.ResReturnReferenceNumber = &v
}

// GetResAuthCode returns the ResAuthCode field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetResAuthCode() string {
	if o == nil || o.ResAuthCode == nil {
		var ret string
		return ret
	}
	return *o.ResAuthCode
}

// GetResAuthCodeOk returns a tuple with the ResAuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetResAuthCodeOk() (*string, bool) {
	if o == nil || o.ResAuthCode == nil {
		return nil, false
	}
	return o.ResAuthCode, true
}

// HasResAuthCode returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasResAuthCode() bool {
	if o != nil && o.ResAuthCode != nil {
		return true
	}

	return false
}

// SetResAuthCode gets a reference to the given string and assigns it to the ResAuthCode field.
func (o *CardAuthorizeResponseResult) SetResAuthCode(v string) {
	o.ResAuthCode = &v
}

// GetResActionCode returns the ResActionCode field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetResActionCode() string {
	if o == nil || o.ResActionCode == nil {
		var ret string
		return ret
	}
	return *o.ResActionCode
}

// GetResActionCodeOk returns a tuple with the ResActionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetResActionCodeOk() (*string, bool) {
	if o == nil || o.ResActionCode == nil {
		return nil, false
	}
	return o.ResActionCode, true
}

// HasResActionCode returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasResActionCode() bool {
	if o != nil && o.ResActionCode != nil {
		return true
	}

	return false
}

// SetResActionCode gets a reference to the given string and assigns it to the ResActionCode field.
func (o *CardAuthorizeResponseResult) SetResActionCode(v string) {
	o.ResActionCode = &v
}

// GetResCenterErrorCode returns the ResCenterErrorCode field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetResCenterErrorCode() string {
	if o == nil || o.ResCenterErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ResCenterErrorCode
}

// GetResCenterErrorCodeOk returns a tuple with the ResCenterErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetResCenterErrorCodeOk() (*string, bool) {
	if o == nil || o.ResCenterErrorCode == nil {
		return nil, false
	}
	return o.ResCenterErrorCode, true
}

// HasResCenterErrorCode returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasResCenterErrorCode() bool {
	if o != nil && o.ResCenterErrorCode != nil {
		return true
	}

	return false
}

// SetResCenterErrorCode gets a reference to the given string and assigns it to the ResCenterErrorCode field.
func (o *CardAuthorizeResponseResult) SetResCenterErrorCode(v string) {
	o.ResCenterErrorCode = &v
}

// GetAcquirerCode returns the AcquirerCode field value if set, zero value otherwise.
func (o *CardAuthorizeResponseResult) GetAcquirerCode() string {
	if o == nil || o.AcquirerCode == nil {
		var ret string
		return ret
	}
	return *o.AcquirerCode
}

// GetAcquirerCodeOk returns a tuple with the AcquirerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeResponseResult) GetAcquirerCodeOk() (*string, bool) {
	if o == nil || o.AcquirerCode == nil {
		return nil, false
	}
	return o.AcquirerCode, true
}

// HasAcquirerCode returns a boolean if a field has been set.
func (o *CardAuthorizeResponseResult) HasAcquirerCode() bool {
	if o != nil && o.AcquirerCode != nil {
		return true
	}

	return false
}

// SetAcquirerCode gets a reference to the given string and assigns it to the AcquirerCode field.
func (o *CardAuthorizeResponseResult) SetAcquirerCode(v string) {
	o.AcquirerCode = &v
}

func (o CardAuthorizeResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["vResultCode"] = o.VResultCode
	}
	if true {
		toSerialize["mstatus"] = o.Mstatus
	}
	if true {
		toSerialize["merrMsg"] = o.MerrMsg
	}
	if true {
		toSerialize["marchTxn"] = o.MarchTxn
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["custTxn"] = o.CustTxn
	}
	if true {
		toSerialize["txnVersion"] = o.TxnVersion
	}
	if o.CardTransactiontype != nil {
		toSerialize["cardTransactiontype"] = o.CardTransactiontype
	}
	if o.GatewayRequestDate != nil {
		toSerialize["gatewayRequestDate"] = o.GatewayRequestDate
	}
	if o.GatewayResponseDate != nil {
		toSerialize["gatewayResponseDate"] = o.GatewayResponseDate
	}
	if o.CenterRequestDate != nil {
		toSerialize["centerRequestDate"] = o.CenterRequestDate
	}
	if o.CenterResponseDate != nil {
		toSerialize["centerResponseDate"] = o.CenterResponseDate
	}
	if o.Pending != nil {
		toSerialize["pending"] = o.Pending
	}
	if o.Loopback != nil {
		toSerialize["loopback"] = o.Loopback
	}
	if o.ConnectedCenterId != nil {
		toSerialize["connectedCenterId"] = o.ConnectedCenterId
	}
	if o.CenterRequestNumber != nil {
		toSerialize["centerRequestNumber"] = o.CenterRequestNumber
	}
	if o.CenterReferenceNumber != nil {
		toSerialize["centerReferenceNumber"] = o.CenterReferenceNumber
	}
	if o.ReqAmount != nil {
		toSerialize["reqAmount"] = o.ReqAmount
	}
	if o.ResReturnReferenceNumber != nil {
		toSerialize["resReturnReferenceNumber"] = o.ResReturnReferenceNumber
	}
	if o.ResAuthCode != nil {
		toSerialize["resAuthCode"] = o.ResAuthCode
	}
	if o.ResActionCode != nil {
		toSerialize["resActionCode"] = o.ResActionCode
	}
	if o.ResCenterErrorCode != nil {
		toSerialize["resCenterErrorCode"] = o.ResCenterErrorCode
	}
	if o.AcquirerCode != nil {
		toSerialize["acquirerCode"] = o.AcquirerCode
	}
	return json.Marshal(toSerialize)
}

type NullableCardAuthorizeResponseResult struct {
	value *CardAuthorizeResponseResult
	isSet bool
}

func (v NullableCardAuthorizeResponseResult) Get() *CardAuthorizeResponseResult {
	return v.value
}

func (v *NullableCardAuthorizeResponseResult) Set(val *CardAuthorizeResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCardAuthorizeResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCardAuthorizeResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardAuthorizeResponseResult(val *CardAuthorizeResponseResult) *NullableCardAuthorizeResponseResult {
	return &NullableCardAuthorizeResponseResult{value: val, isSet: true}
}

func (v NullableCardAuthorizeResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardAuthorizeResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


