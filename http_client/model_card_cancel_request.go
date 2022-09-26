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

// CardCancelRequest struct for CardCancelRequest
type CardCancelRequest struct {
	OrderId string `json:"orderId"`
	Amount string `json:"amount"`
}

// NewCardCancelRequest instantiates a new CardCancelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardCancelRequest(orderId string, amount string) *CardCancelRequest {
	this := CardCancelRequest{}
	this.OrderId = orderId
	this.Amount = amount
	return &this
}

// NewCardCancelRequestWithDefaults instantiates a new CardCancelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardCancelRequestWithDefaults() *CardCancelRequest {
	this := CardCancelRequest{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *CardCancelRequest) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CardCancelRequest) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CardCancelRequest) SetOrderId(v string) {
	o.OrderId = v
}

// GetAmount returns the Amount field value
func (o *CardCancelRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CardCancelRequest) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CardCancelRequest) SetAmount(v string) {
	o.Amount = v
}

func (o CardCancelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableCardCancelRequest struct {
	value *CardCancelRequest
	isSet bool
}

func (v NullableCardCancelRequest) Get() *CardCancelRequest {
	return v.value
}

func (v *NullableCardCancelRequest) Set(val *CardCancelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardCancelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardCancelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardCancelRequest(val *CardCancelRequest) *NullableCardCancelRequest {
	return &NullableCardCancelRequest{value: val, isSet: true}
}

func (v NullableCardCancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardCancelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


