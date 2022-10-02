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

// CardInfoGetRequestParams struct for CardInfoGetRequestParams
type CardInfoGetRequestParams struct {
	TxnVersion string `json:"txnVersion"`
	DummyRequest string `json:"dummyRequest"`
	MerchantCcid string `json:"merchantCcid"`
	PayNowIdParam CardInfoGetRequestParamsAllOfPayNowIdParam `json:"payNowIdParam"`
}

// NewCardInfoGetRequestParams instantiates a new CardInfoGetRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoGetRequestParams(txnVersion string, dummyRequest string, merchantCcid string, payNowIdParam CardInfoGetRequestParamsAllOfPayNowIdParam) *CardInfoGetRequestParams {
	this := CardInfoGetRequestParams{}
	this.TxnVersion = txnVersion
	this.DummyRequest = dummyRequest
	this.MerchantCcid = merchantCcid
	this.PayNowIdParam = payNowIdParam
	return &this
}

// NewCardInfoGetRequestParamsWithDefaults instantiates a new CardInfoGetRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoGetRequestParamsWithDefaults() *CardInfoGetRequestParams {
	this := CardInfoGetRequestParams{}
	return &this
}

// GetTxnVersion returns the TxnVersion field value
func (o *CardInfoGetRequestParams) GetTxnVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnVersion
}

// GetTxnVersionOk returns a tuple with the TxnVersion field value
// and a boolean to check if the value has been set.
func (o *CardInfoGetRequestParams) GetTxnVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnVersion, true
}

// SetTxnVersion sets field value
func (o *CardInfoGetRequestParams) SetTxnVersion(v string) {
	o.TxnVersion = v
}

// GetDummyRequest returns the DummyRequest field value
func (o *CardInfoGetRequestParams) GetDummyRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DummyRequest
}

// GetDummyRequestOk returns a tuple with the DummyRequest field value
// and a boolean to check if the value has been set.
func (o *CardInfoGetRequestParams) GetDummyRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DummyRequest, true
}

// SetDummyRequest sets field value
func (o *CardInfoGetRequestParams) SetDummyRequest(v string) {
	o.DummyRequest = v
}

// GetMerchantCcid returns the MerchantCcid field value
func (o *CardInfoGetRequestParams) GetMerchantCcid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantCcid
}

// GetMerchantCcidOk returns a tuple with the MerchantCcid field value
// and a boolean to check if the value has been set.
func (o *CardInfoGetRequestParams) GetMerchantCcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantCcid, true
}

// SetMerchantCcid sets field value
func (o *CardInfoGetRequestParams) SetMerchantCcid(v string) {
	o.MerchantCcid = v
}

// GetPayNowIdParam returns the PayNowIdParam field value
func (o *CardInfoGetRequestParams) GetPayNowIdParam() CardInfoGetRequestParamsAllOfPayNowIdParam {
	if o == nil {
		var ret CardInfoGetRequestParamsAllOfPayNowIdParam
		return ret
	}

	return o.PayNowIdParam
}

// GetPayNowIdParamOk returns a tuple with the PayNowIdParam field value
// and a boolean to check if the value has been set.
func (o *CardInfoGetRequestParams) GetPayNowIdParamOk() (*CardInfoGetRequestParamsAllOfPayNowIdParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayNowIdParam, true
}

// SetPayNowIdParam sets field value
func (o *CardInfoGetRequestParams) SetPayNowIdParam(v CardInfoGetRequestParamsAllOfPayNowIdParam) {
	o.PayNowIdParam = v
}

func (o CardInfoGetRequestParams) MarshalJSON() ([]byte, error) {
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

type NullableCardInfoGetRequestParams struct {
	value *CardInfoGetRequestParams
	isSet bool
}

func (v NullableCardInfoGetRequestParams) Get() *CardInfoGetRequestParams {
	return v.value
}

func (v *NullableCardInfoGetRequestParams) Set(val *CardInfoGetRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoGetRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoGetRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoGetRequestParams(val *CardInfoGetRequestParams) *NullableCardInfoGetRequestParams {
	return &NullableCardInfoGetRequestParams{value: val, isSet: true}
}

func (v NullableCardInfoGetRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoGetRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


