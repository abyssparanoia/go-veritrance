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

// CardInfoDeleteRequestParamsAllOfPayNowIdParam struct for CardInfoDeleteRequestParamsAllOfPayNowIdParam
type CardInfoDeleteRequestParamsAllOfPayNowIdParam struct {
	AccountParam CardInfoDeleteRequestParamsAllOfPayNowIdParamAccountParam `json:"accountParam"`
}

// NewCardInfoDeleteRequestParamsAllOfPayNowIdParam instantiates a new CardInfoDeleteRequestParamsAllOfPayNowIdParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoDeleteRequestParamsAllOfPayNowIdParam(accountParam CardInfoDeleteRequestParamsAllOfPayNowIdParamAccountParam) *CardInfoDeleteRequestParamsAllOfPayNowIdParam {
	this := CardInfoDeleteRequestParamsAllOfPayNowIdParam{}
	this.AccountParam = accountParam
	return &this
}

// NewCardInfoDeleteRequestParamsAllOfPayNowIdParamWithDefaults instantiates a new CardInfoDeleteRequestParamsAllOfPayNowIdParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoDeleteRequestParamsAllOfPayNowIdParamWithDefaults() *CardInfoDeleteRequestParamsAllOfPayNowIdParam {
	this := CardInfoDeleteRequestParamsAllOfPayNowIdParam{}
	return &this
}

// GetAccountParam returns the AccountParam field value
func (o *CardInfoDeleteRequestParamsAllOfPayNowIdParam) GetAccountParam() CardInfoDeleteRequestParamsAllOfPayNowIdParamAccountParam {
	if o == nil {
		var ret CardInfoDeleteRequestParamsAllOfPayNowIdParamAccountParam
		return ret
	}

	return o.AccountParam
}

// GetAccountParamOk returns a tuple with the AccountParam field value
// and a boolean to check if the value has been set.
func (o *CardInfoDeleteRequestParamsAllOfPayNowIdParam) GetAccountParamOk() (*CardInfoDeleteRequestParamsAllOfPayNowIdParamAccountParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountParam, true
}

// SetAccountParam sets field value
func (o *CardInfoDeleteRequestParamsAllOfPayNowIdParam) SetAccountParam(v CardInfoDeleteRequestParamsAllOfPayNowIdParamAccountParam) {
	o.AccountParam = v
}

func (o CardInfoDeleteRequestParamsAllOfPayNowIdParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountParam"] = o.AccountParam
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam struct {
	value *CardInfoDeleteRequestParamsAllOfPayNowIdParam
	isSet bool
}

func (v NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam) Get() *CardInfoDeleteRequestParamsAllOfPayNowIdParam {
	return v.value
}

func (v *NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam) Set(val *CardInfoDeleteRequestParamsAllOfPayNowIdParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoDeleteRequestParamsAllOfPayNowIdParam(val *CardInfoDeleteRequestParamsAllOfPayNowIdParam) *NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam {
	return &NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam{value: val, isSet: true}
}

func (v NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoDeleteRequestParamsAllOfPayNowIdParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


