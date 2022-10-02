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

// CardInfoDeleteRequestParams struct for CardInfoDeleteRequestParams
type CardInfoDeleteRequestParams struct {
	TxnVersion string `json:"txnVersion"`
	DummyRequest string `json:"dummyRequest"`
	MerchantCcid string `json:"merchantCcid"`
	PayNowIdParam CardInfoDeleteRequestParamsAllOfPayNowIdParam `json:"payNowIdParam"`
}

// NewCardInfoDeleteRequestParams instantiates a new CardInfoDeleteRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoDeleteRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam CardInfoDeleteRequestParamsAllOfPayNowIdParam) *CardInfoDeleteRequestParams {
	this := CardInfoDeleteRequestParams{}
	this.TxnVersion = txnVersion
	this.DummyRequest = dummyRequest
	this.MerchantCcid = merchantCcid
	this.PayNowIdParam = payNowIdParam
	return &this
}

// NewCardInfoDeleteRequestParamsWithDefaults instantiates a new CardInfoDeleteRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoDeleteRequestParamsWithDefaults() *CardInfoDeleteRequestParams {
	this := CardInfoDeleteRequestParams{}
	return &this
}

// GetTxnVersion returns the TxnVersion field value
func (o *CardInfoDeleteRequestParams) GetTxnVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnVersion
}

// GetTxnVersionOk returns a tuple with the TxnVersion field value
// and a boolean to check if the value has been set.
func (o *CardInfoDeleteRequestParams) GetTxnVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnVersion, true
}

// SetTxnVersion sets field value
func (o *CardInfoDeleteRequestParams) SetTxnVersion(v string) {
	o.TxnVersion = v
}

// GetDummyRequest returns the DummyRequest field value
func (o *CardInfoDeleteRequestParams) GetDummyRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DummyRequest
}

// GetDummyRequestOk returns a tuple with the DummyRequest field value
// and a boolean to check if the value has been set.
func (o *CardInfoDeleteRequestParams) GetDummyRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DummyRequest, true
}

// SetDummyRequest sets field value
func (o *CardInfoDeleteRequestParams) SetDummyRequest(v string) {
	o.DummyRequest = v
}

// GetMerchantCcid returns the MerchantCcid field value
func (o *CardInfoDeleteRequestParams) GetMerchantCcid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantCcid
}

// GetMerchantCcidOk returns a tuple with the MerchantCcid field value
// and a boolean to check if the value has been set.
func (o *CardInfoDeleteRequestParams) GetMerchantCcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantCcid, true
}

// SetMerchantCcid sets field value
func (o *CardInfoDeleteRequestParams) SetMerchantCcid(v string) {
	o.MerchantCcid = v
}

// GetPayNowIdParam returns the PayNowIdParam field value
func (o *CardInfoDeleteRequestParams) GetPayNowIdParam() CardInfoDeleteRequestParamsAllOfPayNowIdParam {
	if o == nil {
		var ret CardInfoDeleteRequestParamsAllOfPayNowIdParam
		return ret
	}

	return o.PayNowIdParam
}

// GetPayNowIdParamOk returns a tuple with the PayNowIdParam field value
// and a boolean to check if the value has been set.
func (o *CardInfoDeleteRequestParams) GetPayNowIdParamOk() (*CardInfoDeleteRequestParamsAllOfPayNowIdParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayNowIdParam, true
}

// SetPayNowIdParam sets field value
func (o *CardInfoDeleteRequestParams) SetPayNowIdParam(v CardInfoDeleteRequestParamsAllOfPayNowIdParam) {
	o.PayNowIdParam = v
}

func (o CardInfoDeleteRequestParams) MarshalJSON() ([]byte, error) {
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

type NullableCardInfoDeleteRequestParams struct {
	value *CardInfoDeleteRequestParams
	isSet bool
}

func (v NullableCardInfoDeleteRequestParams) Get() *CardInfoDeleteRequestParams {
	return v.value
}

func (v *NullableCardInfoDeleteRequestParams) Set(val *CardInfoDeleteRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoDeleteRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoDeleteRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoDeleteRequestParams(val *CardInfoDeleteRequestParams) *NullableCardInfoDeleteRequestParams {
	return &NullableCardInfoDeleteRequestParams{value: val, isSet: true}
}

func (v NullableCardInfoDeleteRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoDeleteRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


