/*
Veritrance API

Veritrance API 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package veritrance_paynow

import (
	"encoding/json"
)

// CardAuthorizeRequest struct for CardAuthorizeRequest
type CardAuthorizeRequest struct {
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

// NewCardAuthorizeRequest instantiates a new CardAuthorizeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardAuthorizeRequest(orderId string, originalOrderId string) *CardAuthorizeRequest {
	this := CardAuthorizeRequest{}
	this.OrderId = orderId
	this.OriginalOrderId = originalOrderId
	return &this
}

// NewCardAuthorizeRequestWithDefaults instantiates a new CardAuthorizeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardAuthorizeRequestWithDefaults() *CardAuthorizeRequest {
	this := CardAuthorizeRequest{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *CardAuthorizeRequest) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CardAuthorizeRequest) SetOrderId(v string) {
	o.OrderId = v
}

// GetOriginalOrderId returns the OriginalOrderId field value
func (o *CardAuthorizeRequest) GetOriginalOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalOrderId
}

// GetOriginalOrderIdOk returns a tuple with the OriginalOrderId field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetOriginalOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalOrderId, true
}

// SetOriginalOrderId sets field value
func (o *CardAuthorizeRequest) SetOriginalOrderId(v string) {
	o.OriginalOrderId = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CardAuthorizeRequest) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CardAuthorizeRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CardAuthorizeRequest) SetAmount(v string) {
	o.Amount = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CardAuthorizeRequest) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CardAuthorizeRequest) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CardAuthorizeRequest) SetToken(v string) {
	o.Token = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *CardAuthorizeRequest) GetCardNumber() string {
	if o == nil || o.CardNumber == nil {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetCardNumberOk() (*string, bool) {
	if o == nil || o.CardNumber == nil {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *CardAuthorizeRequest) HasCardNumber() bool {
	if o != nil && o.CardNumber != nil {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *CardAuthorizeRequest) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetCardExpire returns the CardExpire field value if set, zero value otherwise.
func (o *CardAuthorizeRequest) GetCardExpire() string {
	if o == nil || o.CardExpire == nil {
		var ret string
		return ret
	}
	return *o.CardExpire
}

// GetCardExpireOk returns a tuple with the CardExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetCardExpireOk() (*string, bool) {
	if o == nil || o.CardExpire == nil {
		return nil, false
	}
	return o.CardExpire, true
}

// HasCardExpire returns a boolean if a field has been set.
func (o *CardAuthorizeRequest) HasCardExpire() bool {
	if o != nil && o.CardExpire != nil {
		return true
	}

	return false
}

// SetCardExpire gets a reference to the given string and assigns it to the CardExpire field.
func (o *CardAuthorizeRequest) SetCardExpire(v string) {
	o.CardExpire = &v
}

// GetSecurityCode returns the SecurityCode field value if set, zero value otherwise.
func (o *CardAuthorizeRequest) GetSecurityCode() string {
	if o == nil || o.SecurityCode == nil {
		var ret string
		return ret
	}
	return *o.SecurityCode
}

// GetSecurityCodeOk returns a tuple with the SecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetSecurityCodeOk() (*string, bool) {
	if o == nil || o.SecurityCode == nil {
		return nil, false
	}
	return o.SecurityCode, true
}

// HasSecurityCode returns a boolean if a field has been set.
func (o *CardAuthorizeRequest) HasSecurityCode() bool {
	if o != nil && o.SecurityCode != nil {
		return true
	}

	return false
}

// SetSecurityCode gets a reference to the given string and assigns it to the SecurityCode field.
func (o *CardAuthorizeRequest) SetSecurityCode(v string) {
	o.SecurityCode = &v
}

// GetCardOptionType returns the CardOptionType field value if set, zero value otherwise.
func (o *CardAuthorizeRequest) GetCardOptionType() string {
	if o == nil || o.CardOptionType == nil {
		var ret string
		return ret
	}
	return *o.CardOptionType
}

// GetCardOptionTypeOk returns a tuple with the CardOptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetCardOptionTypeOk() (*string, bool) {
	if o == nil || o.CardOptionType == nil {
		return nil, false
	}
	return o.CardOptionType, true
}

// HasCardOptionType returns a boolean if a field has been set.
func (o *CardAuthorizeRequest) HasCardOptionType() bool {
	if o != nil && o.CardOptionType != nil {
		return true
	}

	return false
}

// SetCardOptionType gets a reference to the given string and assigns it to the CardOptionType field.
func (o *CardAuthorizeRequest) SetCardOptionType(v string) {
	o.CardOptionType = &v
}

// GetJpo returns the Jpo field value if set, zero value otherwise.
func (o *CardAuthorizeRequest) GetJpo() string {
	if o == nil || o.Jpo == nil {
		var ret string
		return ret
	}
	return *o.Jpo
}

// GetJpoOk returns a tuple with the Jpo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetJpoOk() (*string, bool) {
	if o == nil || o.Jpo == nil {
		return nil, false
	}
	return o.Jpo, true
}

// HasJpo returns a boolean if a field has been set.
func (o *CardAuthorizeRequest) HasJpo() bool {
	if o != nil && o.Jpo != nil {
		return true
	}

	return false
}

// SetJpo gets a reference to the given string and assigns it to the Jpo field.
func (o *CardAuthorizeRequest) SetJpo(v string) {
	o.Jpo = &v
}

// GetWithCapture returns the WithCapture field value if set, zero value otherwise.
func (o *CardAuthorizeRequest) GetWithCapture() bool {
	if o == nil || o.WithCapture == nil {
		var ret bool
		return ret
	}
	return *o.WithCapture
}

// GetWithCaptureOk returns a tuple with the WithCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequest) GetWithCaptureOk() (*bool, bool) {
	if o == nil || o.WithCapture == nil {
		return nil, false
	}
	return o.WithCapture, true
}

// HasWithCapture returns a boolean if a field has been set.
func (o *CardAuthorizeRequest) HasWithCapture() bool {
	if o != nil && o.WithCapture != nil {
		return true
	}

	return false
}

// SetWithCapture gets a reference to the given bool and assigns it to the WithCapture field.
func (o *CardAuthorizeRequest) SetWithCapture(v bool) {
	o.WithCapture = &v
}

func (o CardAuthorizeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableCardAuthorizeRequest struct {
	value *CardAuthorizeRequest
	isSet bool
}

func (v NullableCardAuthorizeRequest) Get() *CardAuthorizeRequest {
	return v.value
}

func (v *NullableCardAuthorizeRequest) Set(val *CardAuthorizeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardAuthorizeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardAuthorizeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardAuthorizeRequest(val *CardAuthorizeRequest) *NullableCardAuthorizeRequest {
	return &NullableCardAuthorizeRequest{value: val, isSet: true}
}

func (v NullableCardAuthorizeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardAuthorizeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


