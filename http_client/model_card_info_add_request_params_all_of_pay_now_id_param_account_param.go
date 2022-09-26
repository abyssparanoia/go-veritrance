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

// CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam struct for CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam
type CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam struct {
	AccountId string `json:"accountId"`
	CardParam CardInfoAddRequestParamsAllOfPayNowIdParamAccountParamCardParam `json:"cardParam"`
}

// NewCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam instantiates a new CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam(accountId string, cardParam CardInfoAddRequestParamsAllOfPayNowIdParamAccountParamCardParam) *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam {
	this := CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam{}
	this.AccountId = accountId
	this.CardParam = cardParam
	return &this
}

// NewCardInfoAddRequestParamsAllOfPayNowIdParamAccountParamWithDefaults instantiates a new CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoAddRequestParamsAllOfPayNowIdParamAccountParamWithDefaults() *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam {
	this := CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) SetAccountId(v string) {
	o.AccountId = v
}

// GetCardParam returns the CardParam field value
func (o *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) GetCardParam() CardInfoAddRequestParamsAllOfPayNowIdParamAccountParamCardParam {
	if o == nil {
		var ret CardInfoAddRequestParamsAllOfPayNowIdParamAccountParamCardParam
		return ret
	}

	return o.CardParam
}

// GetCardParamOk returns a tuple with the CardParam field value
// and a boolean to check if the value has been set.
func (o *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) GetCardParamOk() (*CardInfoAddRequestParamsAllOfPayNowIdParamAccountParamCardParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardParam, true
}

// SetCardParam sets field value
func (o *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) SetCardParam(v CardInfoAddRequestParamsAllOfPayNowIdParamAccountParamCardParam) {
	o.CardParam = v
}

func (o CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["cardParam"] = o.CardParam
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam struct {
	value *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam
	isSet bool
}

func (v NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) Get() *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam {
	return v.value
}

func (v *NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) Set(val *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam(val *CardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) *NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam {
	return &NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam{value: val, isSet: true}
}

func (v NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoAddRequestParamsAllOfPayNowIdParamAccountParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

