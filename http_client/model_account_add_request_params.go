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

// AccountAddRequestParams struct for AccountAddRequestParams
type AccountAddRequestParams struct {
	TxnVersion string `json:"txnVersion"`
	DummyRequest string `json:"dummyRequest"`
	MerchantCcid string `json:"merchantCcid"`
	PayNowIdParam AccountAddRequestParamsAllOfPayNowIdParam `json:"payNowIdParam"`
}

// NewAccountAddRequestParams instantiates a new AccountAddRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAddRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam AccountAddRequestParamsAllOfPayNowIdParam) *AccountAddRequestParams {
	this := AccountAddRequestParams{}
	this.TxnVersion = txnVersion
	this.DummyRequest = dummyRequest
	this.MerchantCcid = merchantCcid
	this.PayNowIdParam = payNowIdParam
	return &this
}

// NewAccountAddRequestParamsWithDefaults instantiates a new AccountAddRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAddRequestParamsWithDefaults() *AccountAddRequestParams {
	this := AccountAddRequestParams{}
	return &this
}

// GetTxnVersion returns the TxnVersion field value
func (o *AccountAddRequestParams) GetTxnVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnVersion
}

// GetTxnVersionOk returns a tuple with the TxnVersion field value
// and a boolean to check if the value has been set.
func (o *AccountAddRequestParams) GetTxnVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnVersion, true
}

// SetTxnVersion sets field value
func (o *AccountAddRequestParams) SetTxnVersion(v string) {
	o.TxnVersion = v
}

// GetDummyRequest returns the DummyRequest field value
func (o *AccountAddRequestParams) GetDummyRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DummyRequest
}

// GetDummyRequestOk returns a tuple with the DummyRequest field value
// and a boolean to check if the value has been set.
func (o *AccountAddRequestParams) GetDummyRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DummyRequest, true
}

// SetDummyRequest sets field value
func (o *AccountAddRequestParams) SetDummyRequest(v string) {
	o.DummyRequest = v
}

// GetMerchantCcid returns the MerchantCcid field value
func (o *AccountAddRequestParams) GetMerchantCcid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantCcid
}

// GetMerchantCcidOk returns a tuple with the MerchantCcid field value
// and a boolean to check if the value has been set.
func (o *AccountAddRequestParams) GetMerchantCcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantCcid, true
}

// SetMerchantCcid sets field value
func (o *AccountAddRequestParams) SetMerchantCcid(v string) {
	o.MerchantCcid = v
}

// GetPayNowIdParam returns the PayNowIdParam field value
func (o *AccountAddRequestParams) GetPayNowIdParam() AccountAddRequestParamsAllOfPayNowIdParam {
	if o == nil {
		var ret AccountAddRequestParamsAllOfPayNowIdParam
		return ret
	}

	return o.PayNowIdParam
}

// GetPayNowIdParamOk returns a tuple with the PayNowIdParam field value
// and a boolean to check if the value has been set.
func (o *AccountAddRequestParams) GetPayNowIdParamOk() (*AccountAddRequestParamsAllOfPayNowIdParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayNowIdParam, true
}

// SetPayNowIdParam sets field value
func (o *AccountAddRequestParams) SetPayNowIdParam(v AccountAddRequestParamsAllOfPayNowIdParam) {
	o.PayNowIdParam = v
}

func (o AccountAddRequestParams) MarshalJSON() ([]byte, error) {
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
		toSerialize["payNowIdParam"] = o.PayNowIdParam
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAddRequestParams struct {
	value *AccountAddRequestParams
	isSet bool
}

func (v NullableAccountAddRequestParams) Get() *AccountAddRequestParams {
	return v.value
}

func (v *NullableAccountAddRequestParams) Set(val *AccountAddRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAddRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAddRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAddRequestParams(val *AccountAddRequestParams) *NullableAccountAddRequestParams {
	return &NullableAccountAddRequestParams{value: val, isSet: true}
}

func (v NullableAccountAddRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAddRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


