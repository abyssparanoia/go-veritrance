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

// CardInfoDeleteRequestParamsAllOf struct for CardInfoDeleteRequestParamsAllOf
type CardInfoDeleteRequestParamsAllOf struct {
	PayNowIdParam *CardInfoDeleteRequestParamsAllOfPayNowIdParam `json:"payNowIdParam,omitempty"`
}

// NewCardInfoDeleteRequestParamsAllOf instantiates a new CardInfoDeleteRequestParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoDeleteRequestParamsAllOf() *CardInfoDeleteRequestParamsAllOf {
	this := CardInfoDeleteRequestParamsAllOf{}
	return &this
}

// NewCardInfoDeleteRequestParamsAllOfWithDefaults instantiates a new CardInfoDeleteRequestParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoDeleteRequestParamsAllOfWithDefaults() *CardInfoDeleteRequestParamsAllOf {
	this := CardInfoDeleteRequestParamsAllOf{}
	return &this
}

// GetPayNowIdParam returns the PayNowIdParam field value if set, zero value otherwise.
func (o *CardInfoDeleteRequestParamsAllOf) GetPayNowIdParam() CardInfoDeleteRequestParamsAllOfPayNowIdParam {
	if o == nil || o.PayNowIdParam == nil {
		var ret CardInfoDeleteRequestParamsAllOfPayNowIdParam
		return ret
	}
	return *o.PayNowIdParam
}

// GetPayNowIdParamOk returns a tuple with the PayNowIdParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfoDeleteRequestParamsAllOf) GetPayNowIdParamOk() (*CardInfoDeleteRequestParamsAllOfPayNowIdParam, bool) {
	if o == nil || o.PayNowIdParam == nil {
		return nil, false
	}
	return o.PayNowIdParam, true
}

// HasPayNowIdParam returns a boolean if a field has been set.
func (o *CardInfoDeleteRequestParamsAllOf) HasPayNowIdParam() bool {
	if o != nil && o.PayNowIdParam != nil {
		return true
	}

	return false
}

// SetPayNowIdParam gets a reference to the given CardInfoDeleteRequestParamsAllOfPayNowIdParam and assigns it to the PayNowIdParam field.
func (o *CardInfoDeleteRequestParamsAllOf) SetPayNowIdParam(v CardInfoDeleteRequestParamsAllOfPayNowIdParam) {
	o.PayNowIdParam = &v
}

func (o CardInfoDeleteRequestParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayNowIdParam != nil {
		toSerialize["payNowIdParam"] = o.PayNowIdParam
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoDeleteRequestParamsAllOf struct {
	value *CardInfoDeleteRequestParamsAllOf
	isSet bool
}

func (v NullableCardInfoDeleteRequestParamsAllOf) Get() *CardInfoDeleteRequestParamsAllOf {
	return v.value
}

func (v *NullableCardInfoDeleteRequestParamsAllOf) Set(val *CardInfoDeleteRequestParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoDeleteRequestParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoDeleteRequestParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoDeleteRequestParamsAllOf(val *CardInfoDeleteRequestParamsAllOf) *NullableCardInfoDeleteRequestParamsAllOf {
	return &NullableCardInfoDeleteRequestParamsAllOf{value: val, isSet: true}
}

func (v NullableCardInfoDeleteRequestParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoDeleteRequestParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


