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

// PayNowIdResponse struct for PayNowIdResponse
type PayNowIdResponse struct {
	PayNowIdResponse *AbstractPayNowIdResponse `json:"payNowIdResponse,omitempty"`
}

// NewPayNowIdResponse instantiates a new PayNowIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayNowIdResponse() *PayNowIdResponse {
	this := PayNowIdResponse{}
	return &this
}

// NewPayNowIdResponseWithDefaults instantiates a new PayNowIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayNowIdResponseWithDefaults() *PayNowIdResponse {
	this := PayNowIdResponse{}
	return &this
}

// GetPayNowIdResponse returns the PayNowIdResponse field value if set, zero value otherwise.
func (o *PayNowIdResponse) GetPayNowIdResponse() AbstractPayNowIdResponse {
	if o == nil || o.PayNowIdResponse == nil {
		var ret AbstractPayNowIdResponse
		return ret
	}
	return *o.PayNowIdResponse
}

// GetPayNowIdResponseOk returns a tuple with the PayNowIdResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayNowIdResponse) GetPayNowIdResponseOk() (*AbstractPayNowIdResponse, bool) {
	if o == nil || o.PayNowIdResponse == nil {
		return nil, false
	}
	return o.PayNowIdResponse, true
}

// HasPayNowIdResponse returns a boolean if a field has been set.
func (o *PayNowIdResponse) HasPayNowIdResponse() bool {
	if o != nil && o.PayNowIdResponse != nil {
		return true
	}

	return false
}

// SetPayNowIdResponse gets a reference to the given AbstractPayNowIdResponse and assigns it to the PayNowIdResponse field.
func (o *PayNowIdResponse) SetPayNowIdResponse(v AbstractPayNowIdResponse) {
	o.PayNowIdResponse = &v
}

func (o PayNowIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayNowIdResponse != nil {
		toSerialize["payNowIdResponse"] = o.PayNowIdResponse
	}
	return json.Marshal(toSerialize)
}

type NullablePayNowIdResponse struct {
	value *PayNowIdResponse
	isSet bool
}

func (v NullablePayNowIdResponse) Get() *PayNowIdResponse {
	return v.value
}

func (v *NullablePayNowIdResponse) Set(val *PayNowIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePayNowIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePayNowIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayNowIdResponse(val *PayNowIdResponse) *NullablePayNowIdResponse {
	return &NullablePayNowIdResponse{value: val, isSet: true}
}

func (v NullablePayNowIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayNowIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


