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

// CardAuthorizeRequestParams struct for CardAuthorizeRequestParams
type CardAuthorizeRequestParams struct {
	TxnVersion string `json:"txnVersion"`
	DummyRequest string `json:"dummyRequest"`
	MerchantCcid string `json:"merchantCcid"`
	PayNowIdParam CardAuthorizeRequestParamsAllOfPayNowIdParam `json:"payNowIdParam"`
	OrderId string `json:"orderId"`
	Amount string `json:"amount"`
	CurrencyUnit CurrencyUnit `json:"currencyUnit"`
	// トークンサーバーが発行した、クレジットカード情報の識別に用いるトークンの値
	Token *string `json:"token,omitempty"`
	// （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。
	CardNumber *string `json:"cardNumber,omitempty"`
	// （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。
	CardExpire *string `json:"cardExpire,omitempty"`
	// （重要）カード情報の非保持（非通過、非保持）への対応のため、通常は設定しないでください。
	SecurityCode *string `json:"securityCode,omitempty"`
	// カードオプションタイプ （MPI 有り/無し）
	CardOptionType *string `json:"cardOptionType,omitempty"`
	// 支払種別 \"10\"： 一括払い \"21\"： ボーナス一括 \"61Cxx\"： 分割払い、xx に分割回数指定 “80”： リボルビング払い ※指定が無い場合は、\"10\"（一括払い）が適用されます。 
	Jpo *string `json:"jpo,omitempty"`
	// 売上フラグ \"true\"： 与信・売上 \"false\"： 与信のみ 
	WithCapture *bool `json:"withCapture,omitempty"`
}

// NewCardAuthorizeRequestParams instantiates a new CardAuthorizeRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardAuthorizeRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam CardAuthorizeRequestParamsAllOfPayNowIdParam, orderId string, amount string, currencyUnit CurrencyUnit) *CardAuthorizeRequestParams {
	this := CardAuthorizeRequestParams{}
	this.TxnVersion = txnVersion
	this.DummyRequest = dummyRequest
	this.MerchantCcid = merchantCcid
	this.PayNowIdParam = payNowIdParam
	this.OrderId = orderId
	this.Amount = amount
	this.CurrencyUnit = currencyUnit
	return &this
}

// NewCardAuthorizeRequestParamsWithDefaults instantiates a new CardAuthorizeRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardAuthorizeRequestParamsWithDefaults() *CardAuthorizeRequestParams {
	this := CardAuthorizeRequestParams{}
	return &this
}

// GetTxnVersion returns the TxnVersion field value
func (o *CardAuthorizeRequestParams) GetTxnVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnVersion
}

// GetTxnVersionOk returns a tuple with the TxnVersion field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetTxnVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnVersion, true
}

// SetTxnVersion sets field value
func (o *CardAuthorizeRequestParams) SetTxnVersion(v string) {
	o.TxnVersion = v
}

// GetDummyRequest returns the DummyRequest field value
func (o *CardAuthorizeRequestParams) GetDummyRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DummyRequest
}

// GetDummyRequestOk returns a tuple with the DummyRequest field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetDummyRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DummyRequest, true
}

// SetDummyRequest sets field value
func (o *CardAuthorizeRequestParams) SetDummyRequest(v string) {
	o.DummyRequest = v
}

// GetMerchantCcid returns the MerchantCcid field value
func (o *CardAuthorizeRequestParams) GetMerchantCcid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantCcid
}

// GetMerchantCcidOk returns a tuple with the MerchantCcid field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetMerchantCcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantCcid, true
}

// SetMerchantCcid sets field value
func (o *CardAuthorizeRequestParams) SetMerchantCcid(v string) {
	o.MerchantCcid = v
}

// GetPayNowIdParam returns the PayNowIdParam field value
func (o *CardAuthorizeRequestParams) GetPayNowIdParam() CardAuthorizeRequestParamsAllOfPayNowIdParam {
	if o == nil {
		var ret CardAuthorizeRequestParamsAllOfPayNowIdParam
		return ret
	}

	return o.PayNowIdParam
}

// GetPayNowIdParamOk returns a tuple with the PayNowIdParam field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetPayNowIdParamOk() (*CardAuthorizeRequestParamsAllOfPayNowIdParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayNowIdParam, true
}

// SetPayNowIdParam sets field value
func (o *CardAuthorizeRequestParams) SetPayNowIdParam(v CardAuthorizeRequestParamsAllOfPayNowIdParam) {
	o.PayNowIdParam = v
}

// GetOrderId returns the OrderId field value
func (o *CardAuthorizeRequestParams) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CardAuthorizeRequestParams) SetOrderId(v string) {
	o.OrderId = v
}

// GetAmount returns the Amount field value
func (o *CardAuthorizeRequestParams) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CardAuthorizeRequestParams) SetAmount(v string) {
	o.Amount = v
}

// GetCurrencyUnit returns the CurrencyUnit field value
func (o *CardAuthorizeRequestParams) GetCurrencyUnit() CurrencyUnit {
	if o == nil {
		var ret CurrencyUnit
		return ret
	}

	return o.CurrencyUnit
}

// GetCurrencyUnitOk returns a tuple with the CurrencyUnit field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetCurrencyUnitOk() (*CurrencyUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyUnit, true
}

// SetCurrencyUnit sets field value
func (o *CardAuthorizeRequestParams) SetCurrencyUnit(v CurrencyUnit) {
	o.CurrencyUnit = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CardAuthorizeRequestParams) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CardAuthorizeRequestParams) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CardAuthorizeRequestParams) SetToken(v string) {
	o.Token = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *CardAuthorizeRequestParams) GetCardNumber() string {
	if o == nil || o.CardNumber == nil {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetCardNumberOk() (*string, bool) {
	if o == nil || o.CardNumber == nil {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *CardAuthorizeRequestParams) HasCardNumber() bool {
	if o != nil && o.CardNumber != nil {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *CardAuthorizeRequestParams) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetCardExpire returns the CardExpire field value if set, zero value otherwise.
func (o *CardAuthorizeRequestParams) GetCardExpire() string {
	if o == nil || o.CardExpire == nil {
		var ret string
		return ret
	}
	return *o.CardExpire
}

// GetCardExpireOk returns a tuple with the CardExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetCardExpireOk() (*string, bool) {
	if o == nil || o.CardExpire == nil {
		return nil, false
	}
	return o.CardExpire, true
}

// HasCardExpire returns a boolean if a field has been set.
func (o *CardAuthorizeRequestParams) HasCardExpire() bool {
	if o != nil && o.CardExpire != nil {
		return true
	}

	return false
}

// SetCardExpire gets a reference to the given string and assigns it to the CardExpire field.
func (o *CardAuthorizeRequestParams) SetCardExpire(v string) {
	o.CardExpire = &v
}

// GetSecurityCode returns the SecurityCode field value if set, zero value otherwise.
func (o *CardAuthorizeRequestParams) GetSecurityCode() string {
	if o == nil || o.SecurityCode == nil {
		var ret string
		return ret
	}
	return *o.SecurityCode
}

// GetSecurityCodeOk returns a tuple with the SecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetSecurityCodeOk() (*string, bool) {
	if o == nil || o.SecurityCode == nil {
		return nil, false
	}
	return o.SecurityCode, true
}

// HasSecurityCode returns a boolean if a field has been set.
func (o *CardAuthorizeRequestParams) HasSecurityCode() bool {
	if o != nil && o.SecurityCode != nil {
		return true
	}

	return false
}

// SetSecurityCode gets a reference to the given string and assigns it to the SecurityCode field.
func (o *CardAuthorizeRequestParams) SetSecurityCode(v string) {
	o.SecurityCode = &v
}

// GetCardOptionType returns the CardOptionType field value if set, zero value otherwise.
func (o *CardAuthorizeRequestParams) GetCardOptionType() string {
	if o == nil || o.CardOptionType == nil {
		var ret string
		return ret
	}
	return *o.CardOptionType
}

// GetCardOptionTypeOk returns a tuple with the CardOptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetCardOptionTypeOk() (*string, bool) {
	if o == nil || o.CardOptionType == nil {
		return nil, false
	}
	return o.CardOptionType, true
}

// HasCardOptionType returns a boolean if a field has been set.
func (o *CardAuthorizeRequestParams) HasCardOptionType() bool {
	if o != nil && o.CardOptionType != nil {
		return true
	}

	return false
}

// SetCardOptionType gets a reference to the given string and assigns it to the CardOptionType field.
func (o *CardAuthorizeRequestParams) SetCardOptionType(v string) {
	o.CardOptionType = &v
}

// GetJpo returns the Jpo field value if set, zero value otherwise.
func (o *CardAuthorizeRequestParams) GetJpo() string {
	if o == nil || o.Jpo == nil {
		var ret string
		return ret
	}
	return *o.Jpo
}

// GetJpoOk returns a tuple with the Jpo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetJpoOk() (*string, bool) {
	if o == nil || o.Jpo == nil {
		return nil, false
	}
	return o.Jpo, true
}

// HasJpo returns a boolean if a field has been set.
func (o *CardAuthorizeRequestParams) HasJpo() bool {
	if o != nil && o.Jpo != nil {
		return true
	}

	return false
}

// SetJpo gets a reference to the given string and assigns it to the Jpo field.
func (o *CardAuthorizeRequestParams) SetJpo(v string) {
	o.Jpo = &v
}

// GetWithCapture returns the WithCapture field value if set, zero value otherwise.
func (o *CardAuthorizeRequestParams) GetWithCapture() bool {
	if o == nil || o.WithCapture == nil {
		var ret bool
		return ret
	}
	return *o.WithCapture
}

// GetWithCaptureOk returns a tuple with the WithCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestParams) GetWithCaptureOk() (*bool, bool) {
	if o == nil || o.WithCapture == nil {
		return nil, false
	}
	return o.WithCapture, true
}

// HasWithCapture returns a boolean if a field has been set.
func (o *CardAuthorizeRequestParams) HasWithCapture() bool {
	if o != nil && o.WithCapture != nil {
		return true
	}

	return false
}

// SetWithCapture gets a reference to the given bool and assigns it to the WithCapture field.
func (o *CardAuthorizeRequestParams) SetWithCapture(v bool) {
	o.WithCapture = &v
}

func (o CardAuthorizeRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["txnVersion"] = o.TxnVersion
	}
	if true {
		toSerialize["dummyRequest"] = o.DummyRequest
	}
	if true {
		toSerialize["merchantCcid"] = o.MerchantCcid
	}
	if true {
		toSerialize["payNowIdParam"] = o.PayNowIdParam
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currencyUnit"] = o.CurrencyUnit
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.CardNumber != nil {
		toSerialize["cardNumber"] = o.CardNumber
	}
	if o.CardExpire != nil {
		toSerialize["cardExpire"] = o.CardExpire
	}
	if o.SecurityCode != nil {
		toSerialize["securityCode"] = o.SecurityCode
	}
	if o.CardOptionType != nil {
		toSerialize["cardOptionType"] = o.CardOptionType
	}
	if o.Jpo != nil {
		toSerialize["jpo"] = o.Jpo
	}
	if o.WithCapture != nil {
		toSerialize["withCapture"] = o.WithCapture
	}
	return json.Marshal(toSerialize)
}

type NullableCardAuthorizeRequestParams struct {
	value *CardAuthorizeRequestParams
	isSet bool
}

func (v NullableCardAuthorizeRequestParams) Get() *CardAuthorizeRequestParams {
	return v.value
}

func (v *NullableCardAuthorizeRequestParams) Set(val *CardAuthorizeRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCardAuthorizeRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCardAuthorizeRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardAuthorizeRequestParams(val *CardAuthorizeRequestParams) *NullableCardAuthorizeRequestParams {
	return &NullableCardAuthorizeRequestParams{value: val, isSet: true}
}

func (v NullableCardAuthorizeRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardAuthorizeRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


