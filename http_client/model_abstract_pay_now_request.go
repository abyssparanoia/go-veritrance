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

// AbstractPayNowRequest struct for AbstractPayNowRequest
type AbstractPayNowRequest struct {
	TxnVersion string `json:"txnVersion"`
	DummyRequest string `json:"dummyRequest"`
	MerchantCcid string `json:"merchantCcid"`
}

// NewAbstractPayNowRequest instantiates a new AbstractPayNowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractPayNowRequest(txnVersion string, dummyRequest string, merchantCcid string) *AbstractPayNowRequest {
	this := AbstractPayNowRequest{}
	this.TxnVersion = txnVersion
	this.DummyRequest = dummyRequest
	this.MerchantCcid = merchantCcid
	return &this
}

// NewAbstractPayNowRequestWithDefaults instantiates a new AbstractPayNowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractPayNowRequestWithDefaults() *AbstractPayNowRequest {
	this := AbstractPayNowRequest{}
	return &this
}

// GetTxnVersion returns the TxnVersion field value
func (o *AbstractPayNowRequest) GetTxnVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnVersion
}

// GetTxnVersionOk returns a tuple with the TxnVersion field value
// and a boolean to check if the value has been set.
func (o *AbstractPayNowRequest) GetTxnVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnVersion, true
}

// SetTxnVersion sets field value
func (o *AbstractPayNowRequest) SetTxnVersion(v string) {
	o.TxnVersion = v
}

// GetDummyRequest returns the DummyRequest field value
func (o *AbstractPayNowRequest) GetDummyRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DummyRequest
}

// GetDummyRequestOk returns a tuple with the DummyRequest field value
// and a boolean to check if the value has been set.
func (o *AbstractPayNowRequest) GetDummyRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DummyRequest, true
}

// SetDummyRequest sets field value
func (o *AbstractPayNowRequest) SetDummyRequest(v string) {
	o.DummyRequest = v
}

// GetMerchantCcid returns the MerchantCcid field value
func (o *AbstractPayNowRequest) GetMerchantCcid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantCcid
}

// GetMerchantCcidOk returns a tuple with the MerchantCcid field value
// and a boolean to check if the value has been set.
func (o *AbstractPayNowRequest) GetMerchantCcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantCcid, true
}

// SetMerchantCcid sets field value
func (o *AbstractPayNowRequest) SetMerchantCcid(v string) {
	o.MerchantCcid = v
}

func (o AbstractPayNowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["txnVersion"] = o.TxnVersion
	}
	if true {
		toSerialize["dummyRequest"] = o.DummyRequest
	}
	if true {
		toSerialize["merchantCcid"] = o.MerchantCcid
	}
	return json.Marshal(toSerialize)
}

type NullableAbstractPayNowRequest struct {
	value *AbstractPayNowRequest
	isSet bool
}

func (v NullableAbstractPayNowRequest) Get() *AbstractPayNowRequest {
	return v.value
}

func (v *NullableAbstractPayNowRequest) Set(val *AbstractPayNowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractPayNowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractPayNowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractPayNowRequest(val *AbstractPayNowRequest) *NullableAbstractPayNowRequest {
	return &NullableAbstractPayNowRequest{value: val, isSet: true}
}

func (v NullableAbstractPayNowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractPayNowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

