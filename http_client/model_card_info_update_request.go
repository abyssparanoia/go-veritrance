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

// CardInfoUpdateRequest struct for CardInfoUpdateRequest
type CardInfoUpdateRequest struct {
	TxnVersion string `json:"txnVersion"`
	DummyRequest string `json:"dummyRequest"`
	MerchantCcid string `json:"merchantCcid"`
	AccountId string `json:"accountId"`
	CardId string `json:"cardId"`
	CardNumber string `json:"cardNumber"`
	CardExpire string `json:"cardExpire"`
	SecurityCode string `json:"securityCode"`
}

// NewCardInfoUpdateRequest instantiates a new CardInfoUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfoUpdateRequest(txnVersion string, dummyRequest string, merchantCcid string, accountId string, cardId string, cardNumber string, cardExpire string, securityCode string) *CardInfoUpdateRequest {
	this := CardInfoUpdateRequest{}
	this.TxnVersion = txnVersion
	this.DummyRequest = dummyRequest
	this.MerchantCcid = merchantCcid
	this.AccountId = accountId
	this.CardId = cardId
	this.CardNumber = cardNumber
	this.CardExpire = cardExpire
	this.SecurityCode = securityCode
	return &this
}

// NewCardInfoUpdateRequestWithDefaults instantiates a new CardInfoUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoUpdateRequestWithDefaults() *CardInfoUpdateRequest {
	this := CardInfoUpdateRequest{}
	return &this
}

// GetTxnVersion returns the TxnVersion field value
func (o *CardInfoUpdateRequest) GetTxnVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnVersion
}

// GetTxnVersionOk returns a tuple with the TxnVersion field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequest) GetTxnVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnVersion, true
}

// SetTxnVersion sets field value
func (o *CardInfoUpdateRequest) SetTxnVersion(v string) {
	o.TxnVersion = v
}

// GetDummyRequest returns the DummyRequest field value
func (o *CardInfoUpdateRequest) GetDummyRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DummyRequest
}

// GetDummyRequestOk returns a tuple with the DummyRequest field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequest) GetDummyRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DummyRequest, true
}

// SetDummyRequest sets field value
func (o *CardInfoUpdateRequest) SetDummyRequest(v string) {
	o.DummyRequest = v
}

// GetMerchantCcid returns the MerchantCcid field value
func (o *CardInfoUpdateRequest) GetMerchantCcid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantCcid
}

// GetMerchantCcidOk returns a tuple with the MerchantCcid field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequest) GetMerchantCcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantCcid, true
}

// SetMerchantCcid sets field value
func (o *CardInfoUpdateRequest) SetMerchantCcid(v string) {
	o.MerchantCcid = v
}

// GetAccountId returns the AccountId field value
func (o *CardInfoUpdateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CardInfoUpdateRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetCardId returns the CardId field value
func (o *CardInfoUpdateRequest) GetCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequest) GetCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardId, true
}

// SetCardId sets field value
func (o *CardInfoUpdateRequest) SetCardId(v string) {
	o.CardId = v
}

// GetCardNumber returns the CardNumber field value
func (o *CardInfoUpdateRequest) GetCardNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequest) GetCardNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardNumber, true
}

// SetCardNumber sets field value
func (o *CardInfoUpdateRequest) SetCardNumber(v string) {
	o.CardNumber = v
}

// GetCardExpire returns the CardExpire field value
func (o *CardInfoUpdateRequest) GetCardExpire() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardExpire
}

// GetCardExpireOk returns a tuple with the CardExpire field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequest) GetCardExpireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardExpire, true
}

// SetCardExpire sets field value
func (o *CardInfoUpdateRequest) SetCardExpire(v string) {
	o.CardExpire = v
}

// GetSecurityCode returns the SecurityCode field value
func (o *CardInfoUpdateRequest) GetSecurityCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityCode
}

// GetSecurityCodeOk returns a tuple with the SecurityCode field value
// and a boolean to check if the value has been set.
func (o *CardInfoUpdateRequest) GetSecurityCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityCode, true
}

// SetSecurityCode sets field value
func (o *CardInfoUpdateRequest) SetSecurityCode(v string) {
	o.SecurityCode = v
}

func (o CardInfoUpdateRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["cardId"] = o.CardId
	}
	if true {
		toSerialize["cardNumber"] = o.CardNumber
	}
	if true {
		toSerialize["cardExpire"] = o.CardExpire
	}
	if true {
		toSerialize["securityCode"] = o.SecurityCode
	}
	return json.Marshal(toSerialize)
}

type NullableCardInfoUpdateRequest struct {
	value *CardInfoUpdateRequest
	isSet bool
}

func (v NullableCardInfoUpdateRequest) Get() *CardInfoUpdateRequest {
	return v.value
}

func (v *NullableCardInfoUpdateRequest) Set(val *CardInfoUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfoUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfoUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfoUpdateRequest(val *CardInfoUpdateRequest) *NullableCardInfoUpdateRequest {
	return &NullableCardInfoUpdateRequest{value: val, isSet: true}
}

func (v NullableCardInfoUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfoUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


