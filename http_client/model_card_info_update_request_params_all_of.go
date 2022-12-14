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

// CardInfoUpdateRequestParamsAllOf struct for CardInfoUpdateRequestParamsAllOf
type CardInfoUpdateRequestParamsAllOf struct {
	PayNowIdParam *CardInfoUpdateRequestParamsAllOfPayNowIdParam `json:"payNowIdParam,omitempty"`
}

// NewCardInfoUpdateRequestParamsAllOf instantiates a new CardInfoUpdateRequestParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoUpdateRequestParamsAllOf() *CardInfoUpdateRequestParamsAllOf {
	this := CardInfoUpdateRequestParamsAllOf{}
	return &this
}

// NewCardInfoUpdateRequestParamsAllOfWithDefaults instantiates a new CardInfoUpdateRequestParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoUpdateRequestParamsAllOfWithDefaults() *CardInfoUpdateRequestParamsAllOf {
	this := CardInfoUpdateRequestParamsAllOf{}
	return &this
}

// GetPayNowIdParam returns the PayNowIdParam field value if set, zero value otherwise.
func (o *CardInfoUpdateRequestParamsAllOf) GetPayNowIdParam() CardInfoUpdateRequestParamsAllOfPayNowIdParam {
	if o == nil || o.PayNowIdParam == nil {
		var ret CardInfoUpdateRequestParamsAllOfPayNowIdParam
		return ret
	}
	return *o.PayNowIdParam
}

// GetPayNowIdParamOk returns a tuple with the PayNowIdParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequestParamsAllOf) GetPayNowIdParamOk() (*CardInfoUpdateRequestParamsAllOfPayNowIdParam, bool) {
	if o == nil || o.PayNowIdParam == nil {
		return nil, false
	}
	return o.PayNowIdParam, true
}

// HasPayNowIdParam returns a boolean if a field has been set.
func (o *CardInfoUpdateRequestParamsAllOf) HasPayNowIdParam() bool {
	if o != nil && o.PayNowIdParam != nil {
		return true
	}

	return false
}

// SetPayNowIdParam gets a reference to the given CardInfoUpdateRequestParamsAllOfPayNowIdParam and assigns it to the PayNowIdParam field.
func (o *CardInfoUpdateRequestParamsAllOf) SetPayNowIdParam(v CardInfoUpdateRequestParamsAllOfPayNowIdParam) {
	o.PayNowIdParam = &v
}

func (o CardInfoUpdateRequestParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayNowIdParam != nil {
		toSerialize["payNowIdParam"] = o.PayNowIdParam
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoUpdateRequestParamsAllOf struct {
	value *CardInfoUpdateRequestParamsAllOf
	isSet bool
}

func (v NullableCardInfoUpdateRequestParamsAllOf) Get() *CardInfoUpdateRequestParamsAllOf {
	return v.value
}

func (v *NullableCardInfoUpdateRequestParamsAllOf) Set(val *CardInfoUpdateRequestParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoUpdateRequestParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoUpdateRequestParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoUpdateRequestParamsAllOf(val *CardInfoUpdateRequestParamsAllOf) *NullableCardInfoUpdateRequestParamsAllOf {
	return &NullableCardInfoUpdateRequestParamsAllOf{value: val, isSet: true}
}

func (v NullableCardInfoUpdateRequestParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoUpdateRequestParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


