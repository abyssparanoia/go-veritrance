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

// AccountResponse struct for AccountResponse
type AccountResponse struct {
	PayNowIdResponse *AccountResponsePayNowIdResponse `json:"payNowIdResponse,omitempty"`
}

// NewAccountResponse instantiates a new AccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResponse() *AccountResponse {
	this := AccountResponse{}
	return &this
}

// NewAccountResponseWithDefaults instantiates a new AccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResponseWithDefaults() *AccountResponse {
	this := AccountResponse{}
	return &this
}

// GetPayNowIdResponse returns the PayNowIdResponse field value if set, zero value otherwise.
func (o *AccountResponse) GetPayNowIdResponse() AccountResponsePayNowIdResponse {
	if o == nil || o.PayNowIdResponse == nil {
		var ret AccountResponsePayNowIdResponse
		return ret
	}
	return *o.PayNowIdResponse
}

// GetPayNowIdResponseOk returns a tuple with the PayNowIdResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetPayNowIdResponseOk() (*AccountResponsePayNowIdResponse, bool) {
	if o == nil || o.PayNowIdResponse == nil {
		return nil, false
	}
	return o.PayNowIdResponse, true
}

// HasPayNowIdResponse returns a boolean if a field has been set.
func (o *AccountResponse) HasPayNowIdResponse() bool {
	if o != nil && o.PayNowIdResponse != nil {
		return true
	}

	return false
}

// SetPayNowIdResponse gets a reference to the given AccountResponsePayNowIdResponse and assigns it to the PayNowIdResponse field.
func (o *AccountResponse) SetPayNowIdResponse(v AccountResponsePayNowIdResponse) {
	o.PayNowIdResponse = &v
}

func (o AccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayNowIdResponse != nil {
		toSerialize["payNowIdResponse"] = o.PayNowIdResponse
	}
	return json.Marshal(toSerialize)
}

type NullableAccountResponse struct {
	value *AccountResponse
	isSet bool
}

func (v NullableAccountResponse) Get() *AccountResponse {
	return v.value
}

func (v *NullableAccountResponse) Set(val *AccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResponse(val *AccountResponse) *NullableAccountResponse {
	return &NullableAccountResponse{value: val, isSet: true}
}

func (v NullableAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


