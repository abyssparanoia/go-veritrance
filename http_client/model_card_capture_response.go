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

// CardCaptureResponse struct for CardCaptureResponse
type CardCaptureResponse struct {
	PayNowIdResponse *AccountResponsePayNowIdResponse `json:"payNowIdResponse,omitempty"`
	Result *CardCaptureResponseResult `json:"result,omitempty"`
}

// NewCardCaptureResponse instantiates a new CardCaptureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardCaptureResponse() *CardCaptureResponse {
	this := CardCaptureResponse{}
	return &this
}

// NewCardCaptureResponseWithDefaults instantiates a new CardCaptureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardCaptureResponseWithDefaults() *CardCaptureResponse {
	this := CardCaptureResponse{}
	return &this
}

// GetPayNowIdResponse returns the PayNowIdResponse field value if set, zero value otherwise.
func (o *CardCaptureResponse) GetPayNowIdResponse() AccountResponsePayNowIdResponse {
	if o == nil || o.PayNowIdResponse == nil {
		var ret AccountResponsePayNowIdResponse
		return ret
	}
	return *o.PayNowIdResponse
}

// GetPayNowIdResponseOk returns a tuple with the PayNowIdResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardCaptureResponse) GetPayNowIdResponseOk() (*AccountResponsePayNowIdResponse, bool) {
	if o == nil || o.PayNowIdResponse == nil {
		return nil, false
	}
	return o.PayNowIdResponse, true
}

// HasPayNowIdResponse returns a boolean if a field has been set.
func (o *CardCaptureResponse) HasPayNowIdResponse() bool {
	if o != nil && o.PayNowIdResponse != nil {
		return true
	}

	return false
}

// SetPayNowIdResponse gets a reference to the given AccountResponsePayNowIdResponse and assigns it to the PayNowIdResponse field.
func (o *CardCaptureResponse) SetPayNowIdResponse(v AccountResponsePayNowIdResponse) {
	o.PayNowIdResponse = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CardCaptureResponse) GetResult() CardCaptureResponseResult {
	if o == nil || o.Result == nil {
		var ret CardCaptureResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardCaptureResponse) GetResultOk() (*CardCaptureResponseResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CardCaptureResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given CardCaptureResponseResult and assigns it to the Result field.
func (o *CardCaptureResponse) SetResult(v CardCaptureResponseResult) {
	o.Result = &v
}

func (o CardCaptureResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayNowIdResponse != nil {
		toSerialize["payNowIdResponse"] = o.PayNowIdResponse
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableCardCaptureResponse struct {
	value *CardCaptureResponse
	isSet bool
}

func (v NullableCardCaptureResponse) Get() *CardCaptureResponse {
	return v.value
}

func (v *NullableCardCaptureResponse) Set(val *CardCaptureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardCaptureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardCaptureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardCaptureResponse(val *CardCaptureResponse) *NullableCardCaptureResponse {
	return &NullableCardCaptureResponse{value: val, isSet: true}
}

func (v NullableCardCaptureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardCaptureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


