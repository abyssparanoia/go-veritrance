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

// AccountDeleteRequest struct for AccountDeleteRequest
type AccountDeleteRequest struct {
	TxnVersion string `json:"txnVersion"`
	DummyRequest string `json:"dummyRequest"`
	MerchantCcid string `json:"merchantCcid"`
	AccountId string `json:"accountId"`
}

// NewAccountDeleteRequest instantiates a new AccountDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDeleteRequest(txnVersion string, dummyRequest string, merchantCcid string, accountId string) *AccountDeleteRequest {
	this := AccountDeleteRequest{}
	this.TxnVersion = txnVersion
	this.DummyRequest = dummyRequest
	this.MerchantCcid = merchantCcid
	this.AccountId = accountId
	return &this
}

// NewAccountDeleteRequestWithDefaults instantiates a new AccountDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDeleteRequestWithDefaults() *AccountDeleteRequest {
	this := AccountDeleteRequest{}
	return &this
}

// GetTxnVersion returns the TxnVersion field value
func (o *AccountDeleteRequest) GetTxnVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnVersion
}

// GetTxnVersionOk returns a tuple with the TxnVersion field value
// and a boolean to check if the value has been set.
func (o *AccountDeleteRequest) GetTxnVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnVersion, true
}

// SetTxnVersion sets field value
func (o *AccountDeleteRequest) SetTxnVersion(v string) {
	o.TxnVersion = v
}

// GetDummyRequest returns the DummyRequest field value
func (o *AccountDeleteRequest) GetDummyRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DummyRequest
}

// GetDummyRequestOk returns a tuple with the DummyRequest field value
// and a boolean to check if the value has been set.
func (o *AccountDeleteRequest) GetDummyRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DummyRequest, true
}

// SetDummyRequest sets field value
func (o *AccountDeleteRequest) SetDummyRequest(v string) {
	o.DummyRequest = v
}

// GetMerchantCcid returns the MerchantCcid field value
func (o *AccountDeleteRequest) GetMerchantCcid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantCcid
}

// GetMerchantCcidOk returns a tuple with the MerchantCcid field value
// and a boolean to check if the value has been set.
func (o *AccountDeleteRequest) GetMerchantCcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantCcid, true
}

// SetMerchantCcid sets field value
func (o *AccountDeleteRequest) SetMerchantCcid(v string) {
	o.MerchantCcid = v
}

// GetAccountId returns the AccountId field value
func (o *AccountDeleteRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountDeleteRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountDeleteRequest) SetAccountId(v string) {
	o.AccountId = v
}

func (o AccountDeleteRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	return json.Marshal(toSerialize)
}

type NullableAccountDeleteRequest struct {
	value *AccountDeleteRequest
	isSet bool
}

func (v NullableAccountDeleteRequest) Get() *AccountDeleteRequest {
	return v.value
}

func (v *NullableAccountDeleteRequest) Set(val *AccountDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDeleteRequest(val *AccountDeleteRequest) *NullableAccountDeleteRequest {
	return &NullableAccountDeleteRequest{value: val, isSet: true}
}

func (v NullableAccountDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

