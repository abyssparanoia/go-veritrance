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

// CardInfoV1 カード情報
type CardInfoV1 struct {
	// 決済サーバーにて付与されたカード ID が返戻されます。
	CardId string `json:"cardId"`
	// カード番号の先頭 6 桁と下 2 桁のみ数字表記され、その他は \"*\"（アスタリスク）に変換されます。
	CardNumber string `json:"cardNumber"`
	CardExpire string `json:"cardExpire"`
	// カード情報を明示的に指定せず決済する場合に使用するカードか否かを示すフラグ \"1\"：決済時、カード情報を明示的に指定しない場合に使用されるカード \"0\"：決済時、カード情報を明示的に指定しない場合に使用されるカードではない 
	DefaultCard int32 `json:"defaultCard"`
	// カードの名義人
	CardholderName *string `json:"cardholderName,omitempty"`
}

// NewCardInfoV1 instantiates a new CardInfoV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoV1(cardId string, cardNumber string, cardExpire string, defaultCard int32) *CardInfoV1 {
	this := CardInfoV1{}
	this.CardId = cardId
	this.CardNumber = cardNumber
	this.CardExpire = cardExpire
	this.DefaultCard = defaultCard
	return &this
}

// NewCardInfoV1WithDefaults instantiates a new CardInfoV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoV1WithDefaults() *CardInfoV1 {
	this := CardInfoV1{}
	return &this
}

// GetCardId returns the CardId field value
func (o *CardInfoV1) GetCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value
// and a boolean to check if the value has been set.
func (o *CardInfoV1) GetCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardId, true
}

// SetCardId sets field value
func (o *CardInfoV1) SetCardId(v string) {
	o.CardId = v
}

// GetCardNumber returns the CardNumber field value
func (o *CardInfoV1) GetCardNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value
// and a boolean to check if the value has been set.
func (o *CardInfoV1) GetCardNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardNumber, true
}

// SetCardNumber sets field value
func (o *CardInfoV1) SetCardNumber(v string) {
	o.CardNumber = v
}

// GetCardExpire returns the CardExpire field value
func (o *CardInfoV1) GetCardExpire() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardExpire
}

// GetCardExpireOk returns a tuple with the CardExpire field value
// and a boolean to check if the value has been set.
func (o *CardInfoV1) GetCardExpireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardExpire, true
}

// SetCardExpire sets field value
func (o *CardInfoV1) SetCardExpire(v string) {
	o.CardExpire = v
}

// GetDefaultCard returns the DefaultCard field value
func (o *CardInfoV1) GetDefaultCard() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefaultCard
}

// GetDefaultCardOk returns a tuple with the DefaultCard field value
// and a boolean to check if the value has been set.
func (o *CardInfoV1) GetDefaultCardOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultCard, true
}

// SetDefaultCard sets field value
func (o *CardInfoV1) SetDefaultCard(v int32) {
	o.DefaultCard = v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise.
func (o *CardInfoV1) GetCardholderName() string {
	if o == nil || o.CardholderName == nil {
		var ret string
		return ret
	}
	return *o.CardholderName
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfoV1) GetCardholderNameOk() (*string, bool) {
	if o == nil || o.CardholderName == nil {
		return nil, false
	}
	return o.CardholderName, true
}

// HasCardholderName returns a boolean if a field has been set.
func (o *CardInfoV1) HasCardholderName() bool {
	if o != nil && o.CardholderName != nil {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given string and assigns it to the CardholderName field.
func (o *CardInfoV1) SetCardholderName(v string) {
	o.CardholderName = &v
}

func (o CardInfoV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cardId"] = o.CardId
	}
	if true {
		toSerialize["cardNumber"] = o.CardNumber
	}
	if true {
		toSerialize["cardExpire"] = o.CardExpire
	}
	if true {
		toSerialize["defaultCard"] = o.DefaultCard
	}
	if o.CardholderName != nil {
		toSerialize["cardholderName"] = o.CardholderName
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoV1 struct {
	value *CardInfoV1
	isSet bool
}

func (v NullableCardInfoV1) Get() *CardInfoV1 {
	return v.value
}

func (v *NullableCardInfoV1) Set(val *CardInfoV1) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoV1) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoV1(val *CardInfoV1) *NullableCardInfoV1 {
	return &NullableCardInfoV1{value: val, isSet: true}
}

func (v NullableCardInfoV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

