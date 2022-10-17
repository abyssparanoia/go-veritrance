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

// CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam struct for CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam
type CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam struct {
	CardId string `json:"cardId"`
	Token string `json:"token"`
}

// NewCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam instantiates a new CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam(cardId string, token string) *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam {
	this := CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam{}
	this.CardId = cardId
	this.Token = token
	return &this
}

// NewCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParamWithDefaults instantiates a new CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParamWithDefaults() *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam {
	this := CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam{}
	return &this
}

// GetCardId returns the CardId field value
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) GetCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) GetCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardId, true
}

// SetCardId sets field value
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) SetCardId(v string) {
	o.CardId = v
}

// GetToken returns the Token field value
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) SetToken(v string) {
	o.Token = v
}

func (o CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cardId"] = o.CardId
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam struct {
	value *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam
	isSet bool
}

func (v NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) Get() *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam {
	return v.value
}

func (v *NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) Set(val *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam(val *CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) *NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam {
	return &NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam{value: val, isSet: true}
}

func (v NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParamCardParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

