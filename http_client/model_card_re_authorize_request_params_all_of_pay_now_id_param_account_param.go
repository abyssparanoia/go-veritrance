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

// CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam struct for CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam
type CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam struct {
	AccountId string `json:"accountId"`
	OrderId string `json:"orderId"`
	OriginalOrderId string `json:"originalOrderId"`
	Amount *string `json:"amount,omitempty"`
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

// NewCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam instantiates a new CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam(accountId string, orderId string, originalOrderId string) *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam {
	this := CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam{}
	this.AccountId = accountId
	this.OrderId = orderId
	this.OriginalOrderId = originalOrderId
	return &this
}

// NewCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParamWithDefaults instantiates a new CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParamWithDefaults() *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam {
	this := CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetAccountId(v string) {
	o.AccountId = v
}

// GetOrderId returns the OrderId field value
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetOrderId(v string) {
	o.OrderId = v
}

// GetOriginalOrderId returns the OriginalOrderId field value
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetOriginalOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalOrderId
}

// GetOriginalOrderIdOk returns a tuple with the OriginalOrderId field value
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetOriginalOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalOrderId, true
}

// SetOriginalOrderId sets field value
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetOriginalOrderId(v string) {
	o.OriginalOrderId = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetAmount(v string) {
	o.Amount = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetToken(v string) {
	o.Token = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetCardNumber() string {
	if o == nil || o.CardNumber == nil {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetCardNumberOk() (*string, bool) {
	if o == nil || o.CardNumber == nil {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) HasCardNumber() bool {
	if o != nil && o.CardNumber != nil {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetCardExpire returns the CardExpire field value if set, zero value otherwise.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetCardExpire() string {
	if o == nil || o.CardExpire == nil {
		var ret string
		return ret
	}
	return *o.CardExpire
}

// GetCardExpireOk returns a tuple with the CardExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetCardExpireOk() (*string, bool) {
	if o == nil || o.CardExpire == nil {
		return nil, false
	}
	return o.CardExpire, true
}

// HasCardExpire returns a boolean if a field has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) HasCardExpire() bool {
	if o != nil && o.CardExpire != nil {
		return true
	}

	return false
}

// SetCardExpire gets a reference to the given string and assigns it to the CardExpire field.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetCardExpire(v string) {
	o.CardExpire = &v
}

// GetSecurityCode returns the SecurityCode field value if set, zero value otherwise.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetSecurityCode() string {
	if o == nil || o.SecurityCode == nil {
		var ret string
		return ret
	}
	return *o.SecurityCode
}

// GetSecurityCodeOk returns a tuple with the SecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetSecurityCodeOk() (*string, bool) {
	if o == nil || o.SecurityCode == nil {
		return nil, false
	}
	return o.SecurityCode, true
}

// HasSecurityCode returns a boolean if a field has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) HasSecurityCode() bool {
	if o != nil && o.SecurityCode != nil {
		return true
	}

	return false
}

// SetSecurityCode gets a reference to the given string and assigns it to the SecurityCode field.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetSecurityCode(v string) {
	o.SecurityCode = &v
}

// GetCardOptionType returns the CardOptionType field value if set, zero value otherwise.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetCardOptionType() string {
	if o == nil || o.CardOptionType == nil {
		var ret string
		return ret
	}
	return *o.CardOptionType
}

// GetCardOptionTypeOk returns a tuple with the CardOptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetCardOptionTypeOk() (*string, bool) {
	if o == nil || o.CardOptionType == nil {
		return nil, false
	}
	return o.CardOptionType, true
}

// HasCardOptionType returns a boolean if a field has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) HasCardOptionType() bool {
	if o != nil && o.CardOptionType != nil {
		return true
	}

	return false
}

// SetCardOptionType gets a reference to the given string and assigns it to the CardOptionType field.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetCardOptionType(v string) {
	o.CardOptionType = &v
}

// GetJpo returns the Jpo field value if set, zero value otherwise.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetJpo() string {
	if o == nil || o.Jpo == nil {
		var ret string
		return ret
	}
	return *o.Jpo
}

// GetJpoOk returns a tuple with the Jpo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetJpoOk() (*string, bool) {
	if o == nil || o.Jpo == nil {
		return nil, false
	}
	return o.Jpo, true
}

// HasJpo returns a boolean if a field has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) HasJpo() bool {
	if o != nil && o.Jpo != nil {
		return true
	}

	return false
}

// SetJpo gets a reference to the given string and assigns it to the Jpo field.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetJpo(v string) {
	o.Jpo = &v
}

// GetWithCapture returns the WithCapture field value if set, zero value otherwise.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetWithCapture() bool {
	if o == nil || o.WithCapture == nil {
		var ret bool
		return ret
	}
	return *o.WithCapture
}

// GetWithCaptureOk returns a tuple with the WithCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) GetWithCaptureOk() (*bool, bool) {
	if o == nil || o.WithCapture == nil {
		return nil, false
	}
	return o.WithCapture, true
}

// HasWithCapture returns a boolean if a field has been set.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) HasWithCapture() bool {
	if o != nil && o.WithCapture != nil {
		return true
	}

	return false
}

// SetWithCapture gets a reference to the given bool and assigns it to the WithCapture field.
func (o *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) SetWithCapture(v bool) {
	o.WithCapture = &v
}

func (o CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["originalOrderId"] = o.OriginalOrderId
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
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

type NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam struct {
	value *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam
	isSet bool
}

func (v NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) Get() *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam {
	return v.value
}

func (v *NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) Set(val *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam(val *CardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) *NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam {
	return &NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam{value: val, isSet: true}
}

func (v NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardReAuthorizeRequestParamsAllOfPayNowIdParamAccountParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


