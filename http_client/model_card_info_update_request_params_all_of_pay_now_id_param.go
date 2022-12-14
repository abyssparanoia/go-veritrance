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

// CardInfoUpdateRequestParamsAllOfPayNowIdParam struct for CardInfoUpdateRequestParamsAllOfPayNowIdParam
type CardInfoUpdateRequestParamsAllOfPayNowIdParam struct {
	AccountParam CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParam `json:"accountParam"`
}

// NewCardInfoUpdateRequestParamsAllOfPayNowIdParam instantiates a new CardInfoUpdateRequestParamsAllOfPayNowIdParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoUpdateRequestParamsAllOfPayNowIdParam(accountParam CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParam) *CardInfoUpdateRequestParamsAllOfPayNowIdParam {
	this := CardInfoUpdateRequestParamsAllOfPayNowIdParam{}
	this.AccountParam = accountParam
	return &this
}

// NewCardInfoUpdateRequestParamsAllOfPayNowIdParamWithDefaults instantiates a new CardInfoUpdateRequestParamsAllOfPayNowIdParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoUpdateRequestParamsAllOfPayNowIdParamWithDefaults() *CardInfoUpdateRequestParamsAllOfPayNowIdParam {
	this := CardInfoUpdateRequestParamsAllOfPayNowIdParam{}
	return &this
}

// GetAccountParam returns the AccountParam field value
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParam) GetAccountParam() CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParam {
	if o == nil {
		var ret CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParam
		return ret
	}

	return o.AccountParam
}

// GetAccountParamOk returns a tuple with the AccountParam field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParam) GetAccountParamOk() (*CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountParam, true
}

// SetAccountParam sets field value
func (o *CardInfoUpdateRequestParamsAllOfPayNowIdParam) SetAccountParam(v CardInfoUpdateRequestParamsAllOfPayNowIdParamAccountParam) {
	o.AccountParam = v
}

func (o CardInfoUpdateRequestParamsAllOfPayNowIdParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountParam"] = o.AccountParam
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam struct {
	value *CardInfoUpdateRequestParamsAllOfPayNowIdParam
	isSet bool
}

func (v NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam) Get() *CardInfoUpdateRequestParamsAllOfPayNowIdParam {
	return v.value
}

func (v *NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam) Set(val *CardInfoUpdateRequestParamsAllOfPayNowIdParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoUpdateRequestParamsAllOfPayNowIdParam(val *CardInfoUpdateRequestParamsAllOfPayNowIdParam) *NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam {
	return &NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam{value: val, isSet: true}
}

func (v NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoUpdateRequestParamsAllOfPayNowIdParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


