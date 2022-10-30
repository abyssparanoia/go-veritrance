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

// CardCaptureRequest struct for CardCaptureRequest
type CardCaptureRequest struct {
	Params CardCaptureRequestParams `json:"params"`
	AuthHash string `json:"authHash"`
}

// NewCardCaptureRequest instantiates a new CardCaptureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardCaptureRequest(params CardCaptureRequestParams, authHash string) *CardCaptureRequest {
	this := CardCaptureRequest{}
	this.Params = params
	this.AuthHash = authHash
	return &this
}

// NewCardCaptureRequestWithDefaults instantiates a new CardCaptureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardCaptureRequestWithDefaults() *CardCaptureRequest {
	this := CardCaptureRequest{}
	return &this
}

// GetParams returns the Params field value
func (o *CardCaptureRequest) GetParams() CardCaptureRequestParams {
	if o == nil {
		var ret CardCaptureRequestParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CardCaptureRequest) GetParamsOk() (*CardCaptureRequestParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CardCaptureRequest) SetParams(v CardCaptureRequestParams) {
	o.Params = v
}

// GetAuthHash returns the AuthHash field value
func (o *CardCaptureRequest) GetAuthHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthHash
}

// GetAuthHashOk returns a tuple with the AuthHash field value
// and a boolean to check if the value has been set.
func (o *CardCaptureRequest) GetAuthHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthHash, true
}

// SetAuthHash sets field value
func (o *CardCaptureRequest) SetAuthHash(v string) {
	o.AuthHash = v
}

func (o CardCaptureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["params"] = o.Params
	}
	if true {
		toSerialize["authHash"] = o.AuthHash
	}
	return json.Marshal(toSerialize)
}

type NullableCardCaptureRequest struct {
	value *CardCaptureRequest
	isSet bool
}

func (v NullableCardCaptureRequest) Get() *CardCaptureRequest {
	return v.value
}

func (v *NullableCardCaptureRequest) Set(val *CardCaptureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardCaptureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardCaptureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardCaptureRequest(val *CardCaptureRequest) *NullableCardCaptureRequest {
	return &NullableCardCaptureRequest{value: val, isSet: true}
}

func (v NullableCardCaptureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardCaptureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


