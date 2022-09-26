/*
Veritrance API

Veritrance API 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package veritrance_paynow

import (
	"encoding/json"
)

// AccountUpdateRequest struct for AccountUpdateRequest
type AccountUpdateRequest struct {
	AccountId string `json:"accountId"`
}

// NewAccountUpdateRequest instantiates a new AccountUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateRequest(accountId string) *AccountUpdateRequest {
	this := AccountUpdateRequest{}
	this.AccountId = accountId
	return &this
}

// NewAccountUpdateRequestWithDefaults instantiates a new AccountUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateRequestWithDefaults() *AccountUpdateRequest {
	this := AccountUpdateRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountUpdateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountUpdateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountUpdateRequest) SetAccountId(v string) {
	o.AccountId = v
}

func (o AccountUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	return json.Marshal(toSerialize)
}

type NullableAccountUpdateRequest struct {
	value *AccountUpdateRequest
	isSet bool
}

func (v NullableAccountUpdateRequest) Get() *AccountUpdateRequest {
	return v.value
}

func (v *NullableAccountUpdateRequest) Set(val *AccountUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateRequest(val *AccountUpdateRequest) *NullableAccountUpdateRequest {
	return &NullableAccountUpdateRequest{value: val, isSet: true}
}

func (v NullableAccountUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


