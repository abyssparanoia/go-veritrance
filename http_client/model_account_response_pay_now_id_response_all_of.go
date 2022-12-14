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

// AccountResponsePayNowIdResponseAllOf struct for AccountResponsePayNowIdResponseAllOf
type AccountResponsePayNowIdResponseAllOf struct {
	Account *AccountV1 `json:"account,omitempty"`
}

// NewAccountResponsePayNowIdResponseAllOf instantiates a new AccountResponsePayNowIdResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResponsePayNowIdResponseAllOf() *AccountResponsePayNowIdResponseAllOf {
	this := AccountResponsePayNowIdResponseAllOf{}
	return &this
}

// NewAccountResponsePayNowIdResponseAllOfWithDefaults instantiates a new AccountResponsePayNowIdResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResponsePayNowIdResponseAllOfWithDefaults() *AccountResponsePayNowIdResponseAllOf {
	this := AccountResponsePayNowIdResponseAllOf{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AccountResponsePayNowIdResponseAllOf) GetAccount() AccountV1 {
	if o == nil || o.Account == nil {
		var ret AccountV1
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponsePayNowIdResponseAllOf) GetAccountOk() (*AccountV1, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AccountResponsePayNowIdResponseAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccountV1 and assigns it to the Account field.
func (o *AccountResponsePayNowIdResponseAllOf) SetAccount(v AccountV1) {
	o.Account = &v
}

func (o AccountResponsePayNowIdResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableAccountResponsePayNowIdResponseAllOf struct {
	value *AccountResponsePayNowIdResponseAllOf
	isSet bool
}

func (v NullableAccountResponsePayNowIdResponseAllOf) Get() *AccountResponsePayNowIdResponseAllOf {
	return v.value
}

func (v *NullableAccountResponsePayNowIdResponseAllOf) Set(val *AccountResponsePayNowIdResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResponsePayNowIdResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResponsePayNowIdResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResponsePayNowIdResponseAllOf(val *AccountResponsePayNowIdResponseAllOf) *NullableAccountResponsePayNowIdResponseAllOf {
	return &NullableAccountResponsePayNowIdResponseAllOf{value: val, isSet: true}
}

func (v NullableAccountResponsePayNowIdResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResponsePayNowIdResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


