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

// AccountDeleteRequest struct for AccountDeleteRequest
type AccountDeleteRequest struct {
	AccountId string `json:"accountId"`
}

// NewAccountDeleteRequest instantiates a new AccountDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDeleteRequest(accountId string) *AccountDeleteRequest {
	this := AccountDeleteRequest{}
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


