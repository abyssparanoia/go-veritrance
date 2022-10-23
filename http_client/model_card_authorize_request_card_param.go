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

// CardAuthorizeRequestCardParam struct for CardAuthorizeRequestCardParam
type CardAuthorizeRequestCardParam struct {
	CardId string `json:"cardId"`
}

// NewCardAuthorizeRequestCardParam instantiates a new CardAuthorizeRequestCardParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardAuthorizeRequestCardParam(cardId string) *CardAuthorizeRequestCardParam {
	this := CardAuthorizeRequestCardParam{}
	this.CardId = cardId
	return &this
}

// NewCardAuthorizeRequestCardParamWithDefaults instantiates a new CardAuthorizeRequestCardParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardAuthorizeRequestCardParamWithDefaults() *CardAuthorizeRequestCardParam {
	this := CardAuthorizeRequestCardParam{}
	return &this
}

// GetCardId returns the CardId field value
func (o *CardAuthorizeRequestCardParam) GetCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value
// and a boolean to check if the value has been set.
func (o *CardAuthorizeRequestCardParam) GetCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardId, true
}

// SetCardId sets field value
func (o *CardAuthorizeRequestCardParam) SetCardId(v string) {
	o.CardId = v
}

func (o CardAuthorizeRequestCardParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cardId"] = o.CardId
	}
	return json.Marshal(toSerialize)
}

type NullableCardAuthorizeRequestCardParam struct {
	value *CardAuthorizeRequestCardParam
	isSet bool
}

func (v NullableCardAuthorizeRequestCardParam) Get() *CardAuthorizeRequestCardParam {
	return v.value
}

func (v *NullableCardAuthorizeRequestCardParam) Set(val *CardAuthorizeRequestCardParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCardAuthorizeRequestCardParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCardAuthorizeRequestCardParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardAuthorizeRequestCardParam(val *CardAuthorizeRequestCardParam) *NullableCardAuthorizeRequestCardParam {
	return &NullableCardAuthorizeRequestCardParam{value: val, isSet: true}
}

func (v NullableCardAuthorizeRequestCardParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardAuthorizeRequestCardParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

